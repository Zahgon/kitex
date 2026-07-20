package transmeta

import "github.com/bytedance/gopkg/cloud/metainfo"

const (
	MeshVersion uint16 = iota
	TransportType
	LogID
	FromService
	FromCluster
	FromIDC
	ToService
	ToCluster
	ToIDC
	ToMethod
	Env
	DestAddress
	RPCTimeout
	ReadTimeout
	RingHashKey
	DDPTag
	WithMeshHeader
	ConnectTimeout
	SpanContext
	ShortConnection
	FromMethod
	StressTag
	MsgType
	HTTPContentType
	RawRingHashKey
	LBType
)

const (
	HeaderIDLServiceName      = "isn"
	HeaderTransRemoteAddr     = "rip"
	HeaderTransToCluster      = "tc"
	HeaderTransToIDC          = "ti"
	HeaderTransPerfTConnStart = "pcs"
	HeaderTransPerfTConnEnd   = "pce"
	HeaderTransPerfTSendStart = "pss"
	HeaderTransPerfTRecvStart = "prs"
	HeaderTransPerfTRecvEnd   = "pre"

	HeaderConnectionReadyToReset = "crrst"
	HeaderProcessAtTime          = "K_ProcessAtTime"

	HeaderCRC32C = "crc32c"
)

const (
	GDPRToken = metainfo.PrefixTransient + "gdpr-token"
)
