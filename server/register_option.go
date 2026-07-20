package server

import (
	internal_server "github.com/cloudwego/kitex/internal/server"
)

type RegisterOption = internal_server.RegisterOption

type RegisterOptions = internal_server.RegisterOptions

func WithFallbackService() RegisterOption { _ = "STUB: not implemented"; return *new(RegisterOption) }
