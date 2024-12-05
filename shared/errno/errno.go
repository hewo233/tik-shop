package errno

type Errno struct {
	ErrCode int64
	ErrMsg  string
}

func NewErrno(code int64, msg string) Errno {
	return Errno{
		ErrCode: code,
		ErrMsg:  msg,
	}
}

func (e Errno) WithMessage(msg string) Errno {
	e.ErrMsg = msg
	return e
}

const (
	SuccessCode    = 0
	NoRouteCode    = 1
	NoMethodCode   = 2
	BadRequestCode = 10000
)

var (
	Success    = NewErrno(SuccessCode, "success")
	NoRoute    = NewErrno(NoRouteCode, "no route")
	NoMethod   = NewErrno(NoMethodCode, "no method")
	BadRequest = NewErrno(BadRequestCode, "bad request")
)
