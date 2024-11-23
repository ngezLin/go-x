package asynq

import (
	"encoding/json"

	"github.com/hibiken/asynq"
)

type Task struct {
	task *asynq.Task
}

func NewTask(t *asynq.Task) *Task {
	return &Task{
		task: t,
	}
}

func (t *Task) Bind(d any) (err error) {
	err = json.Unmarshal(t.task.Payload(), d)
	if err != nil {
		return
	}

	return
}

func (t *Task) TaskID() string {
	return t.task.ResultWriter().TaskID()
}

func (t *Task) Type() string {
	return t.task.Type()
}
