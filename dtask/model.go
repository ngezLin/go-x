package dtask

import "time"

type TaskInfo struct {
	ID            string
	Queue         string
	Type          string
	Payload       []byte
	State         string
	MaxRetry      int
	Retried       int
	LastErr       string
	LastFailedAt  time.Time
	Timeout       time.Duration
	Deadline      time.Time
	Group         string
	NextProcessAt time.Time
	IsOrphaned    bool
	Retention     time.Duration
	CompletedAt   time.Time
	Result        []byte
}
