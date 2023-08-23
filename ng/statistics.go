package ng

import (
	"encoding/json"
)

const StatisticsCmd = "statistics"

type CurrentMetrics struct {
	SessionsOwn      int `json:"sessionsown"`
	SessionsForeign  int `json:"sessionsforeign"`
	SessionsTotal    int `json:"sessionstotal"`
	TranscodedMedia  int `json:"transcodedmedia"`
	PacketRateUser   int `json:"packetrate_user"`
	ByteRateUser     int `json:"byterate_user"`
	ErrorRateUser    int `json:"errorrate_user"`
	PacketRateKernel int `json:"packetrate_kernel"`
	ByteRateKernel   int `json:"byterate_kernel"`
	ErrorRateKernel  int `json:"errorrate_kernel"`
	PacketRate       int `json:"packetrate"`
	ByteRate         int `json:"byterate"`
	ErrorRate        int `json:"errorrate"`
	MediaUserspace   int `json:"media_userspace"`
	MediaKernel      int `json:"media_kernel"`
	MediaMixed       int `json:"media_mixed"`
}

type TotalMetrics struct {
	Uptime                    int     `json:"uptime,string"`
	ManagedSessions           int     `json:"managedsessions"`
	RejectedSessions          int     `json:"rejectedsessions"`
	TimeoutSessions           int     `json:"timeoutsessions"`
	SilentTimeoutSessions     int     `json:"silenttimeoutsessions"`
	FinalTimeoutSessions      int     `json:"finaltimeoutsessions"`
	OfferTimeoutSessions      int     `json:"offertimeoutsessions"`
	RegularTerminatedSessions int     `json:"regularterminatedsessions"`
	ForcedTerminatedSessions  int     `json:"forcedterminatedsessions"`
	RelayedPacketsUser        int     `json:"relayedpackets_user"`
	RelayedPacketErrorsUser   int     `json:"relayedpacketerrors_user"`
	RelayedBytesUser          int     `json:"relayedbytes_user"`
	RelayedPacketsKernel      int     `json:"relayedpackets_kernel"`
	RelayedPacketErrorsKernel int     `json:"relayedpacketerrors_kernel"`
	RelayedBytesKernel        int     `json:"relayedbytes_kernel"`
	RelayedPackets            int     `json:"relayedpackets"`
	RelayedPacketErrors       int     `json:"relayedpacketerrors"`
	RelayedBytes              int     `json:"relayedbytes"`
	ZerowayStreams            int     `json:"zerowaystreams"`
	OnewayStreams             int     `json:"onewaystreams"`
	AvgCallDuration           float64 `json:"avgcallduration,string"`
	TotalCallsDuration        float64 `json:"totalcallsduration,string"`
	TotalCallsDuration2       float64 `json:"totalcallsduration2,string"`
	TotalCallsDurationStddev  float64 `json:"totalcallsduration_stddev,string"`
}

type MosMetrics struct {
	MosTotal        float64 `json:"mos_total,string"`
	Mos2Total       float64 `json:"mos2_total,string"`
	MosSamplesTotal int     `json:"mos_samples_total"`
	MosAverage      float64 `json:"mos_average,string"`
	MosStddev       float64 `json:"mos_stddev,string"`
}

type VoipMetrics struct {
	JitterTotal                float64 `json:"jitter_total,string"`
	Jitter2Total               float64 `json:"jitter2_total,string"`
	JitterSamplesTotal         int     `json:"jitter_samples_total"`
	JitterAverage              float64 `json:"jitter_average,string"`
	JitterStddev               float64 `json:"jitter_stddev,string"`
	RttE2ETotal                float64 `json:"rtt_e2e_total,string"`
	RttE2E2Total               float64 `json:"rtt_e2e2_total,string"`
	RttE2ESamplesTotal         int     `json:"rtt_e2e_samples_total"`
	RttE2EAverage              float64 `json:"rtt_e2e_average,string"`
	RttE2EStddev               float64 `json:"rtt_e2e_stddev,string"`
	RttDsctTotal               float64 `json:"rtt_dsct_total,string"`
	RttDsct2Total              float64 `json:"rtt_dsct2_total,string"`
	RttDsctSamplesTotal        int     `json:"rtt_dsct_samples_total"`
	RttDsctAverage             float64 `json:"rtt_dsct_average,string"`
	RttDsctStddev              float64 `json:"rtt_dsct_stddev,string"`
	PacketLossTotal            float64 `json:"packetloss_total,string"`
	PacketLoss2Total           float64 `json:"packetloss2_total,string"`
	PacketLossSamplesTotal     int     `json:"packetloss_samples_total"`
	PacketLossAverage          float64 `json:"packetloss_average,string"`
	PacketLossStddev           float64 `json:"packetloss_stddev,string"`
	JitterMeasuredTotal        float64 `json:"jitter_measured_total,string"`
	JitterMeasured2Total       float64 `json:"jitter_measured2_total,string"`
	JitterMeasuredSamplesTotal int     `json:"jitter_measured_samples_total"`
	JitterMeasuredAverage      float64 `json:"jitter_measured_average,string"`
	JitterMeasuredStddev       float64 `json:"jitter_measured_stddev,string"`
	PacketsLost                int     `json:"packets_lost"`
	RtpDuplicates              int     `json:"rtp_duplicates"`
	RtpSkips                   int     `json:"rtp_skips"`
	RtpSeqResets               int     `json:"rtp_seq_resets"`
	RtpReordered               int     `json:"rtp_reordered"`
}

type ProxyMetrics struct {
	Proxy                  string  `json:"proxy"`
	PingCount              int     `json:"pingcount"`
	PingDuration           float64 `json:"pingduration,string"`
	OfferCount             int     `json:"offercount"`
	OfferDuration          float64 `json:"offerduration,string"`
	AnswerCount            int     `json:"answercount"`
	AnswerDuration         float64 `json:"answerduration,string"`
	DeleteCount            int     `json:"deletecount"`
	DeleteDuration         float64 `json:"deleteduration,string"`
	QueryCount             int     `json:"querycount"`
	QueryDuration          float64 `json:"queryduration,string"`
	ListCount              int     `json:"listcount"`
	ListDuration           float64 `json:"listduration,string"`
	StartRecordingCount    int     `json:"startreccount"`
	StartRecordingDuration float64 `json:"startrecduration,string"`
	StopRecordingCount     int     `json:"stopreccount"`
	StopRecordingDuration  float64 `json:"stoprecduration,string"`
	PauseRecordingCount    int     `json:"pausereccount"`
	PauseRecordingDuration float64 `json:"pauserecduration,string"`
	StartForwardCount      int     `json:"startfwdcount"`
	StartForwardDuration   float64 `json:"startfwdduration,string"`
	StopForwardCount       int     `json:"stopfwdcount"`
	StopForwardDuration    float64 `json:"stopfwdduration,string"`
	BlockDtmfCount         int     `json:"blkdtmfcount"`
	BlockDtmfDuration      float64 `json:"blkdtmfduration,string"`
	UnblockDtmfCount       int     `json:"unblkdtmfcount"`
	UnblockDtmfDuration    float64 `json:"unblkdtmfduration,string"`
	BlockMediaCount        int     `json:"blkmediacount"`
	BlockMediaDuration     float64 `json:"blkmediaduration,string"`
	UnblockMediaCount      int     `json:"unblkmediacount"`
	UnblockMediaDuration   float64 `json:"unblkmediaduration,string"`
	PlayMediaCount         int     `json:"playmediacount"`
	PlayMediaDuration      float64 `json:"playmediaduration,string"`
	StopMediaCount         int     `json:"stopmediacount"`
	StopMediaDuration      float64 `json:"stopmediaduration,string"`
	PlayDtmfCount          int     `json:"playdtmfcount"`
	PlayDtmfDuration       float64 `json:"playdtmfduration,string"`
	StatsCount             int     `json:"statscount"`
	StatsDuration          float64 `json:"statsduration,string"`
	SilenceMediaCount      int     `json:"slnmediacount"`
	SilenceMediaDuration   float64 `json:"slnmediaduration,string"`
	UnsilenceMediaCount    int     `json:"unslnmediacount"`
	UnsilenceMediaDuration float64 `json:"unslnmediaduration,string"`
	PubCount               int     `json:"pubcount"`
	PubDuration            float64 `json:"pubduration,string"`
	SubReqCount            int     `json:"subreqcount"`
	SubReqDuration         float64 `json:"subreqduration,string"`
	SubAnsCount            int     `json:"subanscount"`
	SubAnsDuration         float64 `json:"subansduration,string"`
	UnsubCount             int     `json:"unsubcount"`
	UnsubDuration          float64 `json:"unsubduration,string"`
	ErrorCount             int     `json:"errorcount"`
}

type ControlMetrics struct {
	Proxies                  []ProxyMetrics `json:"proxies"`
	TotalPingCount           int            `json:"totalpingcount"`
	TotalOfferCount          int            `json:"totaloffercount"`
	TotalAnswerCount         int            `json:"totalanswercount"`
	TotalDeleteCount         int            `json:"totaldeletecount"`
	TotalQueryCount          int            `json:"totalquerycount"`
	TotalListCount           int            `json:"totallistcount"`
	TotalStartRecordingCount int            `json:"totalstartreccount"`
	TotalStopRecordingCount  int            `json:"totalstopreccount"`
	TotalPauseRecordingCount int            `json:"totalpausereccount"`
	TotalStartForwardCount   int            `json:"totalstartfwdcount"`
	TotalStopForwardCount    int            `json:"totalstopfwdcount"`
	TotalBlockDtmfCount      int            `json:"totalblkdtmfcount"`
	TotalUnblockDtmfCount    int            `json:"totalunblkdtmfcount"`
	TotalBlockMediaCount     int            `json:"totalblkmediacount"`
	TotalUnblockMediaCount   int            `json:"totalunblkmediacount"`
	TotalPlayMediaCount      int            `json:"totalplaymediacount"`
	TotalStopMediaCount      int            `json:"totalstopmediacount"`
	TotalPlayDtmfCount       int            `json:"totalplaydtmfcount"`
	TotalStatsCount          int            `json:"totalstatscount"`
	TotalSlnmediaCount       int            `json:"totalslnmediacount"`
	TotalUnslnmediaCount     int            `json:"totalunslnmediacount"`
	TotalPubCount            int            `json:"totalpubcount"`
	TotalSubreqCount         int            `json:"totalsubreqcount"`
	TotalSubansCount         int            `json:"totalsubanscount"`
	TotalunsubCount          int            `json:"totalunsubcount"`
}

type InterfacesMetrics struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Ports   struct {
		Min     int    `json:"min"`
		Max     int    `json:"max"`
		Used    int    `json:"used"`
		UsedPct string `json:"used_pct"`
		Free    int    `json:"free"`
		Totals  int    `json:"totals"`
		Last    int    `json:"last"`
	} `json:"ports"`
	PacketsLost int `json:"packets_lost"`
	Duplicates  int `json:"duplicates"`
	VoipMetrics struct {
		MosTotal                   float64 `json:"mos_total,string"`
		Mos2Total                  float64 `json:"mos2_total,string"`
		MosSamplesTotal            int     `json:"mos_samples_total"`
		MosAverage                 float64 `json:"mos_average,string"`
		MosStddev                  float64 `json:"mos_stddev,string"`
		JitterTotal                float64 `json:"jitter_total,string"`
		Jitter2Total               float64 `json:"jitter2_total,string"`
		JitterSamplesTotal         int     `json:"jitter_samples_total"`
		JitterAverage              float64 `json:"jitter_average,string"`
		JitterStddev               float64 `json:"jitter_stddev,string"`
		RttE2ETotal                float64 `json:"rtt_e2e_total,string"`
		RttE2E2Total               float64 `json:"rtt_e2e2_total,string"`
		RttE2ESamplesTotal         int     `json:"rtt_e2e_samples_total"`
		RttE2EAverage              float64 `json:"rtt_e2e_average,string"`
		RttE2EStddev               float64 `json:"rtt_e2e_stddev,string"`
		RttDsctTotal               float64 `json:"rtt_dsct_total,string"`
		RttDsct2Total              float64 `json:"rtt_dsct2_total,string"`
		RttDsctSamplesTotal        int     `json:"rtt_dsct_samples_total"`
		RttDsctAverage             float64 `json:"rtt_dsct_average,string"`
		RttDsctStddev              float64 `json:"rtt_dsct_stddev,string"`
		PacketlossTotal            float64 `json:"packetloss_total,string"`
		Packetloss2Total           float64 `json:"packetloss2_total,string"`
		PacketlossSamplesTotal     int     `json:"packetloss_samples_total"`
		PacketlossAverage          float64 `json:"packetloss_average,string"`
		PacketlossStddev           float64 `json:"packetloss_stddev,string"`
		JitterMeasuredTotal        float64 `json:"jitter_measured_total,string"`
		JitterMeasured2Total       float64 `json:"jitter_measured2_total,string"`
		JitterMeasuredSamplesTotal int     `json:"jitter_measured_samples_total"`
		JitterMeasuredAverage      float64 `json:"jitter_measured_average,string"`
		JitterMeasuredStddev       float64 `json:"jitter_measured_stddev,string"`
	} `json:"voip_metrics"`
	Ingress struct {
		Packets int `json:"packets"`
		Bytes   int `json:"bytes"`
		Errors  int `json:"errors"`
	} `json:"ingress"`
	Egress struct {
		Packets int `json:"packets"`
		Bytes   int `json:"bytes"`
		Errors  int `json:"errors"`
	} `json:"egress"`
}

type Statistics struct {
	Current     CurrentMetrics      `json:"currentstatistics"`
	Total       TotalMetrics        `json:"totalstatistics"`
	Mos         MosMetrics          `json:"mos"`
	Voip        VoipMetrics         `json:"voip_metrics"`
	Control     ControlMetrics      `json:"controlstatistics"`
	Interfaces  []InterfacesMetrics `json:"interfaces"`
	Transcoders []any               `json:"transcoders"`
}

func (c Client) GetStatistics() (*Statistics, error) {
	decodedMsg, err := c.RunCommand(StatisticsCmd)
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(decodedMsg["statistics"])
	if err != nil {
		return nil, err
	}

	var stats Statistics
	err = json.Unmarshal(data, &stats)
	if err != nil {
		return nil, err
	}
	return &stats, nil
}
