package remote

type RPCRole int

const (
	Client RPCRole = iota
	Server
)
