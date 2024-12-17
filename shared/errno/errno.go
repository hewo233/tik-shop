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
	SuccessCode                   = 0
	NoRouteCode                   = 1
	NoMethodCode                  = 2
	StatusOKCode                  = 20000
	StatusBadRequestCode          = 40000
	StatusUnauthorizedCode        = 40100
	ForbiddenCode                 = 40300
	StatusNotFoundCode            = 40400
	StatusNotAcceptableCode       = 40600
	StatusConflictCode            = 40900
	StatusInternalServerErrorCode = 50000
)

var (
	Success                   = NewErrno(SuccessCode, "success")
	NoRoute                   = NewErrno(NoRouteCode, "no route")
	NoMethod                  = NewErrno(NoMethodCode, "no method")
	StatusOK                  = NewErrno(StatusOKCode, "ok")
	StatusBadRequest          = NewErrno(StatusBadRequestCode, "bad request")
	StatusUnauthorized        = NewErrno(StatusUnauthorizedCode, "unauthorized")
	Forbidden                 = NewErrno(ForbiddenCode, "forbidden")
	StatusNotFound            = NewErrno(StatusNotFoundCode, "not found")
	StatusNotAcceptable       = NewErrno(StatusNotAcceptableCode, "not acceptable")
	StatusConflict            = NewErrno(StatusConflictCode, "conflict")
	StatusInternalServerError = NewErrno(StatusInternalServerErrorCode, "internal server error")
)
