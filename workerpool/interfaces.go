package workerpool

type ITask interface {
	TaskID() string
}

type IResult interface {
	Data() []byte
	Err() error
}
