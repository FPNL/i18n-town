package ierror

import "fmt"

var Validate *ValidateErr

type ValidateErr struct {
	msg string
}

func (e *ValidateErr) Error() string {
	return e.msg
}

func NewValidateErr(s string, a ...any) error {
	return &ValidateErr{
		msg: fmt.Sprintf(s, a),
	}
}

var Server *ServerErr

type ServerErr struct {
	msg string
}

func (e *ServerErr) Error() string {
	return e.msg
}

func NewServerErr(s string, a ...any) error {
	return &ServerErr{
		msg: fmt.Sprintf(s, a),
	}
}
