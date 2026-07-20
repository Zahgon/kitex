package kerrors

var ErrStreamingProtocol = &basicError{"streaming protocol error"}

var ErrStreamingCanceled = &basicError{"streaming canceled"}

var ErrStreamingTimeout = &basicError{"streaming timeout"}
