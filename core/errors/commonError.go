package errors

type CommonError struct {
	Msg string
}

func (c *CommonError) Error() string {
	return c.Msg
}
