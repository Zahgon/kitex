package utils

const (
	EnvConfDir  = "KITEX_CONF_DIR"
	EnvConfFile = "KITEX_CONF_FILE"
	EnvLogDir   = "KITEX_LOG_DIR"

	DefaultConfDir  = "conf"
	DefaultConfFile = "kitex.yml"
	DefaultLogDir   = "log"
)

func GetConfDir() string { _ = "STUB: not implemented"; return "" }

func GetConfFile() string { _ = "STUB: not implemented"; return "" }

func GetEnvLogDir() string { _ = "STUB: not implemented"; return "" }

func GetLogDir() string { _ = "STUB: not implemented"; return "" }
