package messaging

type (
	response struct {
		report bool
		err    error
		data   interface{}
	}

	Response interface {
		Error() error
		Data() interface{}
		Report() bool
		IsError() bool
	}
)

func (e response) Error() error {
	return e.err
}

func (e response) Data() interface{} {
	return e.data
}

func (e response) Report() bool {
	return e.report
}

func (e response) IsError() bool {
	return e.err != nil
}

func ReportError(err error, data interface{}) response {
	return response{
		report: true,
		err:    err,
		data:   data,
	}
}

func ExpectError(err error, data interface{}) response {
	return response{
		err:  err,
		data: data,
	}
}

func Done(data interface{}) response {
	return response{
		data: data,
	}
}
