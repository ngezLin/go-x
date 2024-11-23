package dtask

import (
	"context"
	"time"
)

type Level string

var (
	Critical = "critical"
	Default  = "default"
	Low      = "low"
)

type (
	Publisher interface {
		PublishWithSchedule(ctx context.Context, queueName, id string, data interface{}, processAt time.Time, deadline time.Time) (info *TaskInfo, err error)
		Publish(ctx context.Context, queueName, id string, data interface{}, deadline time.Time) (info *TaskInfo, err error)
		PublishWithLevel(ctx context.Context, queueName, id string, data interface{}, deadline time.Time, level Level) (info *TaskInfo, err error)
		Delete(ctx context.Context, queueName, id string) (err error)
		GetTaskInfo(ctx context.Context, queueName string, ids []string) (res []TaskInfo, err error)
		GetAllArchivedTasks(ctx context.Context, queueName string) (res []string, err error)
		RunAllRetryTasks(ctx context.Context, queueName string) (int, error)
	}
)
