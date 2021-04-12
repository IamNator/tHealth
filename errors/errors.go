package errors

type errors struct {
	Code    int         `json:"code"`
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
}

//New returns a new error object
func New(code int, status bool, message string, err interface{}) errors {
	return errors{
		Code:    code,
		Status:  status,
		Message: message,
		Error:   err,
	}
}

//Default creates a errors object with status defined
func Default(message string) errors {
	return errors{
		Status:  false,
		Message: message,
		Error:   message,
	}
}

//GetStatus returns status of the error
func (e errors) GetCode() int {
	return int(e.Code)
}

//SetStatus sets the status of the error
func (e *errors) SetCode(s bool) bool {
	e.Status = s
	return e.Status
}



//GetStatus returns status of the error
func (e errors) GetStatus() bool {
	return e.Status
}

//SetStatus sets the status of the error
func (e *errors) SetStatus(s bool) bool {
	e.Status = s
	return e.Status
}