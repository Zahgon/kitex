package codec

import (
	"context"
	"hash/crc32"
	"sync"

	"github.com/cloudwego/kitex/pkg/remote"
)

const (
	PayloadValidatorPrefix   = "PV_"
	maxPayloadChecksumLength = 4096
)

type PayloadValidator interface {
	Key(ctx context.Context) string

	Generate(ctx context.Context, outboundPayload []byte) (need bool, checksum string, err error)

	Validate(ctx context.Context, expectedValue string, inboundPayload []byte) (pass bool, err error)
}

func getValidatorKey(ctx context.Context, p PayloadValidator) string {
	_ = "STUB: not implemented"
	return ""
}

func payloadChecksumGenerate(ctx context.Context, pv PayloadValidator, outboundPayload []byte, message remote.Message) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func payloadChecksumValidate(ctx context.Context, pv PayloadValidator, in remote.ByteBuffer, message remote.Message) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func fillRPCInfoBeforeValidate(ctx context.Context, message remote.Message) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func NewCRC32PayloadValidator() PayloadValidator {
	_ = "STUB: not implemented"
	return *new(PayloadValidator)
}

type crcPayloadValidator struct{}

var _ PayloadValidator = &crcPayloadValidator{}

func (p *crcPayloadValidator) Key(ctx context.Context) string { _ = "STUB: not implemented"; return "" }

func (p *crcPayloadValidator) Generate(ctx context.Context, outPayload []byte) (need bool, value string, err error) {
	_ = "STUB: not implemented"
	return false, "", nil
}

func (p *crcPayloadValidator) Validate(ctx context.Context, expectedValue string, inputPayload []byte) (pass bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

var (
	crc32cTable    *crc32.Table
	crc32TableOnce sync.Once
)

func getCRC32C(payload []byte) string { _ = "STUB: not implemented"; return "" }
