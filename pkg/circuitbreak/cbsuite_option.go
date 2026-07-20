package circuitbreak

type CBSuiteConfig struct {
	serviceGetErrorTypeFunc  GetErrorTypeFunc
	instanceGetErrorTypeFunc GetErrorTypeFunc
}

type CBSuiteOption func(s *CBSuiteConfig)

func WithServiceGetErrorType(customFunc GetErrorTypeFunc) CBSuiteOption {
	_ = "STUB: not implemented"
	return *new(CBSuiteOption)
}

func WithWrappedServiceGetErrorType(customFunc GetErrorTypeFunc) CBSuiteOption {
	_ = "STUB: not implemented"
	return *new(CBSuiteOption)
}

func WithInstanceGetErrorType(f GetErrorTypeFunc) CBSuiteOption {
	_ = "STUB: not implemented"
	return *new(CBSuiteOption)
}

func WithWrappedInstanceGetErrorType(f GetErrorTypeFunc) CBSuiteOption {
	_ = "STUB: not implemented"
	return *new(CBSuiteOption)
}

func WrapErrorTypeFunc(customFunc, originalFunc GetErrorTypeFunc) GetErrorTypeFunc {
	_ = "STUB: not implemented"
	return *new(GetErrorTypeFunc)
}
