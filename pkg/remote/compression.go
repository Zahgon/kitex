package remote

import (
	"github.com/cloudwego/kitex/pkg/rpcinfo"
)

type CompressType int32

const (
	NoCompress CompressType = iota
	GZip
)

func SetRecvCompressor(ri rpcinfo.RPCInfo, compressorName string) {
	_ = "STUB: not implemented"
	return
}

func SetSendCompressor(ri rpcinfo.RPCInfo, compressorName string) {
	_ = "STUB: not implemented"
	return
}

func GetSendCompressor(ri rpcinfo.RPCInfo) string { _ = "STUB: not implemented"; return "" }

func GetRecvCompressor(ri rpcinfo.RPCInfo) string { _ = "STUB: not implemented"; return "" }
