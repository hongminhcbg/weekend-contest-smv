package erp

type RuntimeError struct {
	Code    int
	Message string
}

func NewRuntimeError(code int, m string) *RuntimeError {
	return &RuntimeError{
		Code:    code,
		Message: m,
	}
}

var ERR_BAD_REQUEST = NewRuntimeError(400000, "Yêu cầu không hợp lệ")
var ERR_INTENAL_SERVER = NewRuntimeError(500000, "Đã có lỗi hệ thống, vui lòng thử lại sau")
var ERR_RATE_LIMIT = NewRuntimeError(429000, "Hmmm, bạn thao tác quá nhanh rồi!!!!")
