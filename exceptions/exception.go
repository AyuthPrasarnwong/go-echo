package exceptions

import "reflect"

type (
	// ErrorException set error data
	ErrorException struct {
		Message  string
		Detail   string
		Code     int64
		ErrorKey string
		Errors   interface{}
		kind     reflect.Kind
		typ      reflect.Type
	}

	// ErrorInterface setup get error data
	ErrorInterface interface {
		GetMessage() string
		GetDetail() string
		GetErrors() interface{}
		GetCode() int64
		// Kind returns the Field's reflect Kind
		//
		// eg. time.Time's kind is a struct
		Kind() reflect.Kind

		// Type returns the Field's reflect Type
		//
		// eg. time.Time's type is time.Time
		Type() reflect.Type
	}
)

// compile time interface checks
var _ ErrorInterface = new(ErrorException)
var _ error = new(ErrorException)

func (fe *ErrorException) Error() string {
	return fe.GetMessage()
}

// GetMessage get error message
func (fe *ErrorException) GetMessage() string {
	return fe.Message
}

// GetDetail get error detail
func (fe *ErrorException) GetDetail() string {
	return fe.Detail
}

// GetCode get response code
func (fe *ErrorException) GetCode() int64 {
	return fe.Code
}

// GetErrors get error list
func (fe *ErrorException) GetErrors() interface{} {
	return fe.Errors
}

// GetErrorKey get error message
func (fe *ErrorException) GetErrorKey() string {
	if fe.ErrorKey == "" {
		return "internal-server-error"
	}
	return fe.ErrorKey
}

// Kind returns the Field's reflect Kind
func (fe *ErrorException) Kind() reflect.Kind {
	return fe.kind
}

// Type returns the Field's reflect Type
func (fe *ErrorException) Type() reflect.Type {
	return fe.typ
}

// func (ve ErrorException) Error() string {

// 	buff := bytes.NewBufferString("")

// 	var fe *ErrorException

// 	buff.WriteString(fe.Error())
// 	buff.WriteString("\n")

// 	return strings.TrimSpace(buff.String())
// }
