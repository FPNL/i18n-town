package ierror

var Validate *ValidateErr

type ValidateErr struct {
	msg string
}

func (e *ValidateErr) Error() string {
	return e.msg
}

func NewValidateErr(s string) error {
	return &ValidateErr{
		msg: s,
	}
}
