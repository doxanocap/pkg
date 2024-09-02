package test_service

type Result struct {
	data []byte
	err  error
}

func (r Result) Data() []byte {
	return r.data
}

func (r Result) Err() error {
	return r.err
}
