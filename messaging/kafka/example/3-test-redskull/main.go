package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/super-saga/go-x/messaging/kafka"
)

const KafkaHeaderKey = "redskullHeader"

// type MessageError struct {
// 	MessageStr string `json:"MessageStr"`
// 	StackTrace string `json:"StackTrace"`
// }
// type MetaRedSkull struct {
// 	RedSkullID  string       `json:"RedSkullID"`
// 	RefID       string       `json:"RefID"`
// 	Attempt     int          `json:"Attempt"`
// 	MaxAttempt  int          `json:"MaxAttempt"`
// 	DestTopic   string       `json:"DestTopic"`
// 	OriginTopic string       `json:"OriginTopic"`
// 	ClusterID   string       `json:"ClusterID"`
// 	Error       MessageError `json:"Error"`
// }

type MessageError struct {
	MessageStr string
	StackTrace string
}
type MetaRedSkull struct {
	RedSkullID  string
	RefID       string
	Attempt     int
	MaxAttempt  int
	DestTopic   string
	OriginTopic string
	ClusterID   string
	Error       MessageError
}

type User struct {
	ID       string
	Username string
}

func main() {
	fmt.Println("hello world")
	pub()
}

func pub() {
	p, err := kafka.NewPublisher([]string{"kafka01-bootstrap.dev.amartha.id:9094"}, kafka.WithOrigin("local-suhar"))
	if err != nil {
		log.Fatal(err)
	}

	defer p.ClosePublisher()

	user := User{
		ID:       time.Now().Format(time.RFC3339),
		Username: "user" + time.Now().Format(time.RFC3339),
	}

	meta := MetaRedSkull{
		RefID:       user.ID,
		MaxAttempt:  3,
		DestTopic:   "test-topic-dlq-retry",
		OriginTopic: "test-topic-dlq",
		ClusterID:   "main",
	}

	err = p.PublishSyncWithKeyAndHeader(context.Background(), "redskulld-dlq", "", user, map[string]interface{}{
		KafkaHeaderKey: meta,
	})

	if err != nil {
		log.Fatal(err)
	}

	log.Println("done")

}
