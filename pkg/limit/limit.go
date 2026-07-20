package limit

type Updater interface {
	UpdateLimit(opt *Option) (updated bool)
}

type Option struct {
	MaxConnections int
	MaxQPS         int

	UpdateControl func(u Updater)
}

func (lo *Option) Valid() bool { _ = "STUB: not implemented"; return false }
