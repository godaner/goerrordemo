package error

// common
var (
	ErrRequest  = &RequestErr{}
	ErrNotFound = &NotFoundErr{}
	ErrServer   = &ServerErr{}
)
