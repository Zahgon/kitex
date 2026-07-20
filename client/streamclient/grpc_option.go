package streamclient

import (
	"crypto/tls"

	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2/grpc"
)

func WithGRPCConnPoolSize(s uint32) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithGRPCWriteBufferSize(s uint32) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithGRPCReadBufferSize(s uint32) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithGRPCInitialWindowSize(s uint32) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithGRPCInitialConnWindowSize(s uint32) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithGRPCMaxHeaderListSize(s uint32) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithGRPCKeepaliveParams(kp grpc.ClientKeepalive) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithGRPCTLSConfig(tlsConfig *tls.Config) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}
