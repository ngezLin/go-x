package asynq

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/hibiken/asynq"
	"github.com/super-saga/go-x/dtask"
	"github.com/super-saga/go-x/graceful"
)

const (
	DefaultTimeout   = time.Second * 10
	DefaultMaxRetry  = 10
	DefaultRetention = time.Hour * 2
)

type (
	publisher struct {
		client   *asynq.Client
		redisOpt *asynq.RedisClientOpt
		cfg      *PublisherConfig
	}
	PublisherConfig struct {
		MaxRetry  int
		Timeout   time.Duration
		Retention time.Duration
	}
)

var (
	ErrTaskNotFound = errors.New("asynq: task not found")
)

func NewPublisher(ctx context.Context, redisOpt *asynq.RedisClientOpt, cfg *PublisherConfig) (*publisher, graceful.ProcessStopper, error) {
	client := asynq.NewClient(redisOpt)

	if cfg.Timeout == 0 {
		cfg.Timeout = DefaultTimeout
	}
	if cfg.MaxRetry == 0 {
		cfg.MaxRetry = DefaultMaxRetry
	}
	if cfg.Retention == 0 {
		cfg.Retention = DefaultRetention
	}

	p := &publisher{
		client:   client,
		redisOpt: redisOpt,
		cfg:      cfg,
	}
	stopper := func(ctx context.Context) error {
		return client.Close()
	}
	return p, stopper, nil
}

func (pub *publisher) Publish(ctx context.Context, queueName, id string, data interface{}, deadline time.Time) (info *dtask.TaskInfo, err error) {
	jData, err := json.Marshal(data)
	if err != nil {
		return
	}

	task := asynq.NewTask(queueName, jData, asynq.MaxRetry(pub.cfg.MaxRetry), asynq.Timeout(pub.cfg.Timeout), asynq.Deadline(deadline), asynq.Retention(pub.cfg.Retention), asynq.Queue(queueName))
	asynqTaskInfo, err := pub.client.Enqueue(task, asynq.TaskID(id))
	if err != nil {
		return
	}
	info = &dtask.TaskInfo{
		ID:            asynqTaskInfo.ID,
		Queue:         asynqTaskInfo.Queue,
		Type:          asynqTaskInfo.Type,
		Payload:       asynqTaskInfo.Payload,
		State:         asynqTaskInfo.State.String(),
		MaxRetry:      asynqTaskInfo.MaxRetry,
		Retried:       asynqTaskInfo.Retried,
		LastErr:       asynqTaskInfo.LastErr,
		LastFailedAt:  asynqTaskInfo.LastFailedAt,
		Timeout:       asynqTaskInfo.Timeout,
		Deadline:      asynqTaskInfo.Deadline,
		Group:         asynqTaskInfo.Group,
		NextProcessAt: asynqTaskInfo.NextProcessAt,
		IsOrphaned:    asynqTaskInfo.IsOrphaned,
		Retention:     asynqTaskInfo.Retention,
		CompletedAt:   asynqTaskInfo.CompletedAt,
		Result:        asynqTaskInfo.Result,
	}

	return
}

func (pub *publisher) PublishWithSchedule(ctx context.Context, queueName, id string, data interface{}, processAt time.Time, deadline time.Time) (info *dtask.TaskInfo, err error) {
	jData, err := json.Marshal(data)
	if err != nil {
		return
	}

	task := asynq.NewTask(queueName, jData, asynq.ProcessAt(processAt), asynq.Timeout(pub.cfg.Timeout), asynq.Deadline(deadline), asynq.MaxRetry(pub.cfg.MaxRetry), asynq.Retention(pub.cfg.Retention), asynq.Queue(queueName))
	asynqTaskInfo, err := pub.client.Enqueue(task, asynq.TaskID(id))
	if err != nil {
		return
	}
	info = &dtask.TaskInfo{
		ID:            asynqTaskInfo.ID,
		Queue:         asynqTaskInfo.Queue,
		Type:          asynqTaskInfo.Type,
		Payload:       asynqTaskInfo.Payload,
		State:         asynqTaskInfo.State.String(),
		MaxRetry:      asynqTaskInfo.MaxRetry,
		Retried:       asynqTaskInfo.Retried,
		LastErr:       asynqTaskInfo.LastErr,
		LastFailedAt:  asynqTaskInfo.LastFailedAt,
		Timeout:       asynqTaskInfo.Timeout,
		Deadline:      asynqTaskInfo.Deadline,
		Group:         asynqTaskInfo.Group,
		NextProcessAt: asynqTaskInfo.NextProcessAt,
		IsOrphaned:    asynqTaskInfo.IsOrphaned,
		Retention:     asynqTaskInfo.Retention,
		CompletedAt:   asynqTaskInfo.CompletedAt,
		Result:        asynqTaskInfo.Result,
	}

	return
}

func (pub *publisher) PublishWithLevel(ctx context.Context, queueName, id string, data interface{}, deadline time.Time, level dtask.Level) (info *dtask.TaskInfo, err error) {
	jData, err := json.Marshal(data)
	if err != nil {
		return
	}

	task := asynq.NewTask(queueName, jData, asynq.Queue(string(level)), asynq.MaxRetry(pub.cfg.MaxRetry), asynq.Retention(pub.cfg.Retention), asynq.Timeout(pub.cfg.Timeout), asynq.Deadline(deadline), asynq.Queue(queueName))
	asynqTaskInfo, err := pub.client.Enqueue(task, asynq.TaskID(id))
	if err != nil {
		return
	}
	info = &dtask.TaskInfo{
		ID:            asynqTaskInfo.ID,
		Queue:         asynqTaskInfo.Queue,
		Type:          asynqTaskInfo.Type,
		Payload:       asynqTaskInfo.Payload,
		State:         asynqTaskInfo.State.String(),
		MaxRetry:      asynqTaskInfo.MaxRetry,
		Retried:       asynqTaskInfo.Retried,
		LastErr:       asynqTaskInfo.LastErr,
		LastFailedAt:  asynqTaskInfo.LastFailedAt,
		Timeout:       asynqTaskInfo.Timeout,
		Deadline:      asynqTaskInfo.Deadline,
		Group:         asynqTaskInfo.Group,
		NextProcessAt: asynqTaskInfo.NextProcessAt,
		IsOrphaned:    asynqTaskInfo.IsOrphaned,
		Retention:     asynqTaskInfo.Retention,
		CompletedAt:   asynqTaskInfo.CompletedAt,
		Result:        asynqTaskInfo.Result,
	}

	return
}

func (pub *publisher) Delete(ctx context.Context, queueName, id string) (err error) {
	inspector := asynq.NewInspector(pub.redisOpt)
	defer inspector.Close()

	err = inspector.DeleteTask(queueName, id)
	if err != nil {
		return
	}
	return
}

func (pub *publisher) GetTaskInfo(ctx context.Context, queueName string, ids []string) (res []dtask.TaskInfo, err error) {
	inspector := asynq.NewInspector(pub.redisOpt)
	defer inspector.Close()

	for _, id := range ids {
		asynqTaskInfo, err := inspector.GetTaskInfo(queueName, id)
		if err != nil {
			if err.Error() == ErrTaskNotFound.Error() {
				res = append(res, dtask.TaskInfo{
					ID:    id,
					State: "not found",
				})

				continue
			}

			return res, err
		}

		res = append(res, dtask.TaskInfo{
			ID:            asynqTaskInfo.ID,
			Queue:         asynqTaskInfo.Queue,
			Type:          asynqTaskInfo.Type,
			Payload:       asynqTaskInfo.Payload,
			State:         asynqTaskInfo.State.String(),
			MaxRetry:      asynqTaskInfo.MaxRetry,
			Retried:       asynqTaskInfo.Retried,
			LastErr:       asynqTaskInfo.LastErr,
			LastFailedAt:  asynqTaskInfo.LastFailedAt,
			Timeout:       asynqTaskInfo.Timeout,
			Deadline:      asynqTaskInfo.Deadline,
			Group:         asynqTaskInfo.Group,
			NextProcessAt: asynqTaskInfo.NextProcessAt,
			IsOrphaned:    asynqTaskInfo.IsOrphaned,
			Retention:     asynqTaskInfo.Retention,
			CompletedAt:   asynqTaskInfo.CompletedAt,
			Result:        asynqTaskInfo.Result,
		})
	}

	return
}

func (pub *publisher) GetAllArchivedTasks(ctx context.Context, queueName string) (res []string, err error) {
	inspector := asynq.NewInspector(pub.redisOpt)
	defer inspector.Close()

	tasks, err := inspector.ListArchivedTasks(queueName)
	if err != nil {
		return
	}

	for _, task := range tasks {
		res = append(res, task.ID)
	}

	return
}

func (pub *publisher) RunAllRetryTasks(ctx context.Context, queueName string) (int, error) {
	var (
		count int
		err   error
	)
	inspector := asynq.NewInspector(pub.redisOpt)
	defer inspector.Close()

	count, err = inspector.RunAllRetryTasks(queueName)
	if err != nil {
		return count, err
	}
	return count, nil
}
