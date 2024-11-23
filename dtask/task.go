package dtask

type Task interface {
	Bind(data any) error
	TaskID() string
	Type() string
}
