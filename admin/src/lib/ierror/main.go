package ierror

var (
	Validate *ValidateErr
	Unknown  *UnknownErr
)

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

type UnknownErr struct {
	msg string
}

func (e *UnknownErr) Error() string {
	return e.msg
}

func NewUnknownErr(s string) error {
	return &UnknownErr{
		msg: s,
	}
}
