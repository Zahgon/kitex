package netpoll

import (
	"context"
	"net"
	"net/http"

	"github.com/cloudwego/kitex/pkg/remote"
	"github.com/cloudwego/kitex/pkg/remote/trans"
)

func newHTTPCliTransHandler(opt *remote.ClientOption, ext trans.Extension) (remote.ClientTransHandler, error) {
	_ = "STUB: not implemented"
	return *new(remote.ClientTransHandler), nil
}

type httpCliTransHandler struct {
	opt       *remote.ClientOption
	codec     remote.Codec
	transPipe *remote.TransPipeline
	ext       trans.Extension
}

func (t *httpCliTransHandler) Write(ctx context.Context, conn net.Conn, sendMsg remote.Message) (nctx context.Context, err error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func (t *httpCliTransHandler) Read(ctx context.Context, conn net.Conn, msg remote.Message) (nctx context.Context, err error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func (t *httpCliTransHandler) OnMessage(ctx context.Context, args, result remote.Message) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func (t *httpCliTransHandler) OnInactive(ctx context.Context, conn net.Conn) {
	_ = "STUB: not implemented"
	return
}

func (t *httpCliTransHandler) OnError(ctx context.Context, err error, conn net.Conn) {
	_ = "STUB: not implemented"
	return
}

func (t *httpCliTransHandler) SetPipeline(p *remote.TransPipeline) {
	_ = "STUB: not implemented"
	return
}

func addMetaInfo(msg remote.Message, h http.Header) error { _ = "STUB: not implemented"; return nil }

func readLine(buffer remote.ByteBuffer) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func skipLine(buffer remote.ByteBuffer) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func parseHTTPResponseHead(line string) (protoMajor, protoMinor, statusCodeInt int, err error) {
	_ = "STUB: not implemented"
	return 0, 0, 0, nil
}

func skipToBody(buffer remote.ByteBuffer) error { _ = "STUB: not implemented"; return nil }

func getBodyBufReader(buf remote.ByteBuffer) (remote.ByteBuffer, error) {
	_ = "STUB: not implemented"
	return *new(remote.ByteBuffer), nil
}
