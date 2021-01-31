package error

type CustomerError struct {
	msg string
	err error
}

func (c CustomerError) Error() string {
	return "CustomerError:" + c.msg
}
