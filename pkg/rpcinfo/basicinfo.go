package rpcinfo

type EndpointBasicInfo struct {
	ServiceName string
	Method      string
	Tags        map[string]string
}

const (
	ConnResetTag     = "crrst"
	RetryTag         = "retry"
	RetryLastCostTag = "last_cost"
	RetryPrevInstTag = "prev_inst"
	ShmIPCTag        = "shmipc"
	RemoteClosedTag  = "remote_closed"
)

const (
	HTTPURL = "http_url"

	HTTPHost = "http_host"

	HTTPHeader = "http_header"
)
