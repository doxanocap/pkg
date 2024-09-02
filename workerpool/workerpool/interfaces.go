package workerpool

type ITask interface {
	ID() uint64
}
