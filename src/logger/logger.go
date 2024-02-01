package logger

type logger struct {
	Message string   `json:"message"`
	Err     string   `json:"error"`
	Code    int      `json:"code"`
	Causes  []Causes `json:"causes"`
}

type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func NewRestErr(message, err string, code int, causes []Causes) *logger {
	return &logger{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}
