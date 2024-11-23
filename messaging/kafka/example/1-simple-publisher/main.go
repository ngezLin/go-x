package main

import (
	"context"
	"log"

	"github.com/super-saga/go-x/messaging"
	"github.com/super-saga/go-x/messaging/kafka"
)

type UserService struct {
	// Create publisher using interface.
	publisher messaging.PublisherWithKeyAndHeader
}

// Submit is only test function
func (u *UserService) Submit(username string) error {
	// Do whatever.
	userData := struct {
		Username string
	}{
		Username: username,
	}

	// Publish to kafka in sync with partition key and header.
	err := u.publisher.PublishSyncWithKeyAndHeader(context.Background(), "test-push", username, userData, map[string]interface{}{
		"custom-header": "header value",
	})

	// Async example.
	// To do async you can use this function "PublishAsyncWithKeyAndHeader" or "PublishAsyncWithKey"
	// The async will return *Promise, you may need to handle the success/error callback by using Promise.Then method.
	done := make(chan struct{})
	promise, err := u.publisher.PublishAsyncWithKeyAndHeader(context.Background(), "test-push", username, userData, map[string]interface{}{
		"custom-header": "header value",
	})
	promise.Then(func(errCb error) {
		if errCb != nil {
			err = errCb
		}
		close(done)
	})
	<-done

	// Do whatever.

	return err
}

func main() {
	// Create new instance of kafka that implement all the function of the interface messaging.PublisherWithKeyAndHeader.
	pub, err := kafka.NewPublisher([]string{"127.0.0.1:9093"}, kafka.WithOrigin("testservice"))
	if err != nil {
		log.Fatal(err)
	}
	defer pub.ClosePublisher()

	userService := UserService{
		publisher: pub,
	}

	err = userService.Submit("test-user")
	if err != nil {
		log.Fatal(err)
	}
}
