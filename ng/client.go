package ng

import (
	"bytes"
	"errors"
	"net"
	"net/url"
	"time"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/google/uuid"
	"github.com/stefanovazzocell/bencode"
)

const Offer = "offer"
const Answer = "answer"
const Delete = "delete"
const Query = "query"
const Ping = "ping"
const StartRecording = "start recording"
const StopRecording = "stop recording"
const BlockDMTF = "block DTMF"
const UnblockDMTF = "unblock DTMF"
const PlayDMTF = "play DTMF"
const PlayMedia = "play media"
const StopMedia = "stop media"
const BlockMedia = "block media"
const UnblockMedia = "unblock media"
const SilenceMedia = "silence media"
const UnsilenceMedia = "unsilence media"
const StartForwarding = "start forwarding"
const StopForwarding = "stop forwarding"
const Publish = "publish"
const SubscribeRequest = "subscribe request"
const SubscribeAnswer = "subscribe answer"
const Unsubscribe = "unsubscribe"

type Client struct {
	conn    net.Conn
	logger  log.Logger
	URL     *url.URL
	Timeout time.Duration
}

func NewClient(rtpengineURI string, timeout time.Duration, logger log.Logger) (*Client, error) {
	rtpengineURL, err := url.Parse(rtpengineURI)
	if err != nil {
		return nil, err
	}

	conn, err := net.Dial(rtpengineURL.Scheme, rtpengineURL.Host)
	if err != nil {
		return nil, err
	}

	return &Client{
		conn:    conn,
		URL:     rtpengineURL,
		Timeout: timeout,
		logger:  logger,
	}, nil
}

func (c Client) Close() {
	c.conn.Close()
}

func (c Client) RunCommand(command string) (map[string]interface{}, error) {
	cookie := uuid.New().String()
	err := c.sendCommand(cookie, command)
	if err != nil {
		return nil, err
	}

	decodedMsg, err := c.getResponse(cookie)
	if err != nil {
		return nil, err
	}
	return decodedMsg, nil
}

func (c Client) sendCommand(cookie string, command string) error {
	message, err := encodeCommand(cookie, command)
	if err != nil {
		return err
	}
	level.Debug(c.logger).Log("cookie", cookie, "command", command, "message", string(message))
	if _, err := c.conn.Write(message); err != nil {
		return err
	}
	return nil
}

func (c Client) getResponse(cookie string) (map[string]interface{}, error) {
	c.conn.SetReadDeadline(time.Now().Add(c.Timeout))
	response := make([]byte, 65536)
	_, err := c.conn.Read(response)
	if err != nil {
		return nil, err
	}

	level.Debug(c.logger).Log("cookie", cookie, "response", string(response))
	decodedMsg, err := decodeResponse(cookie, response)
	if err != nil {
		return nil, err
	}
	return c.validateResponse(decodedMsg)
}

func (c Client) validateResponse(decodedMsg map[string]interface{}) (map[string]interface{}, error) {
	result, ok := decodedMsg["result"]
	if !ok {
		return nil, errors.New("No result returned")
	}
	if result == "ok" || result == "pong" {
		if warning, ok := decodedMsg["warning"]; ok {
			level.Warn(c.logger).Log("warning", warning)
		}
		return decodedMsg, nil
	}
	if result == "error" {
		if reason, ok := decodedMsg["error-reason"]; ok {
			return nil, errors.New(reason.(string))
		}
	}
	return nil, errors.New("Unknown error")
}

func decodeResponse(cookie string, response []byte) (map[string]interface{}, error) {
	cookieIndex := bytes.IndexAny(response, " ")
	if cookieIndex != len(cookie) {
		return nil, errors.New("Error parsing message")
	}

	cookieResponse := string(response[:cookieIndex])
	if cookieResponse != cookie {
		return nil, errors.New("Expected cookie did not match")
	}

	encodedData := string(response[cookieIndex+1:])
	decodedData, err := bencode.NewParserFromString(encodedData).AsDict()
	if err != nil {
		return nil, err
	}
	return decodedData, nil
}

func encodeCommand(cookie string, command string) ([]byte, error) {
	var dict interface{} = map[string]interface{}{
		"command": command,
	}
	bdata, err := bencode.NewEncoderFromInterface(dict)
	if err != nil {
		return nil, err
	}
	bid := []byte(cookie + " ")
	return append(bid, bdata.Bytes()...), nil
}
