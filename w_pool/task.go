package w_pool

type Task struct {
	Error error
	Data  interface{}
	f     func(interface{}) error
}

func NewTask(data interface{}, f func(interface{}) error) *Task {
	return &Task{
		Error: nil,
		Data:  data,
		f:     f,
	}
}

func (t *Task) Run(workerId int) {
	//fmt.Println("task running on " + strconv.Itoa(workerId) + " worker...")
	t.Error = t.f(t.Data)
}
