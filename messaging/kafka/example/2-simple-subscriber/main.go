package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/super-saga/go-x/messaging"
	"github.com/super-saga/go-x/messaging/codec"
	"github.com/super-saga/go-x/messaging/kafka"

	_ "net/http/pprof"
)

func main() {
	reg := prometheus.DefaultRegisterer
	go func() {
		http.Handle("/metrics", promhttp.InstrumentMetricHandler(reg, promhttp.HandlerFor(prometheus.DefaultGatherer, promhttp.HandlerOpts{})))
		http.ListenAndServe("localhost:3000", nil)
	}()
	withDelay(reg)
	// withoutDelay()
}

func withDelay(reg prometheus.Registerer) {
	log.Println("starting")
	consumerGroup := "simple-subscriber-local"
	brokers := []string{"localhost:9092"}
	sub, err := kafka.NewSubscriber(
		brokers,
		consumerGroup,
		kafka.WithMiddleware(mid1),
		kafka.WithSubsciberEndCallback(end),
		kafka.WithSubsciberPreStartCallback(preStart),
		kafka.WithSubscriberErrorCallback(errcb),
		kafka.WithSubscriberKafkaVersion("3.3.1"),
		kafka.WithSubscriberGenericPromMetrics(reg, "go-payment-lib", "subscriber", 1*time.Second),
	)

	if err != nil {
		log.Fatalln("error creating subscriber:", err)
	}
	log.Println("subscriber created")

	defer func() {
		log.Println("closing subscriber")
		err := sub.CloseSubscriber()
		log.Println("error close bro:", err)
	}()

	wg := new(sync.WaitGroup)
	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Println("Ready?")
		err := sub.Subscribe(ctx, messaging.WithTopic("test-topic", codec.NewJson("v1"), handler, mid2))
		log.Println("stop subscriber:", err)
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("Blocking, press ctrl+c to continue...")
	<-done // Will block here until user hits ctrl+c
	cancel()
	wg.Wait()
	log.Println("Done gracefully")
}

func withoutDelay() {
	log.Println("starting")
	consumerGroup := "coretan-test-push"
	brokers := []string{"127.0.0.1:9093"}
	sub, err := kafka.NewSubscriber(
		brokers,
		consumerGroup,
		kafka.WithMiddleware(mid1),
		kafka.WithSubsciberEndCallback(end),
		kafka.WithSubsciberPreStartCallback(preStart),
		kafka.WithSubscriberErrorCallback(errcb),
		kafka.WithSubscriberKafkaVersion("3.3.1"),
	)

	if err != nil {
		log.Fatalln("error creating subscriber:", err)
	}
	log.Println("subscriber created")

	defer func() {
		log.Println("closing subscriber")
		err := sub.CloseSubscriber()
		log.Println("error close bro:", err)
	}()

	wg := new(sync.WaitGroup)
	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Println("Ready?")
		err := sub.Subscribe(ctx, messaging.WithDelayedTopic("test-push", codec.NewJson("v1"), 1*time.Minute, 10, handler, mid2))
		log.Println("stop subscriber:", err)
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("Blocking, press ctrl+c to continue...")
	<-done // Will block here until user hits ctrl+c
	cancel()
	wg.Wait()
	log.Println("Done gracefully")
}

func handler(message messaging.Message) messaging.Response {
	type data struct {
		ID       int
		Username string
	}
	log.Println("got message", message)

	var d data
	err := message.Bind(&d)
	log.Println("message data", d, d.ID, d.Username, err)

	log.Println("ctx 1: ", message.Context().Value("hello1"))
	log.Println("ctx 2: ", message.Context().Value("hello2"))

	// select {
	// case <-time.After(10 * time.Second):
	// 	log.Println("SELESAI WAITING")
	// case <-message.Context().Done():
	// 	log.Println("CTX DONE")
	// }

	// time.Sleep(1 * time.Second)
	// log.Println("wakeup :)")

	return messaging.Done(d)
}

func mid1(next messaging.SubscriptionHandler) messaging.SubscriptionHandler {
	return func(message messaging.Message) messaging.Response {
		log.Println("Masuk ke middleware 1", message, message.Context().Value("hello2"))
		message = message.WithContext(context.WithValue(message.Context(), "hello1", "mid1"))
		return next(message)
	}
}

func mid2(next messaging.SubscriptionHandler) messaging.SubscriptionHandler {
	return func(message messaging.Message) messaging.Response {
		log.Println("Masuk ke middleware 2", message, message.Context().Value("hello1"))
		message = message.WithContext(context.WithValue(message.Context(), "hello2", "mid2"))
		return next(message)
	}
}

func end(ctx context.Context, claims map[string][]int32) error {
	log.Println("Masuk ke end", claims)
	return nil
}

func preStart(ctx context.Context, claims map[string][]int32) error {
	log.Println("Masuk ke preStart", claims)
	return nil
}

func errcb(err error) {
	log.Println("Masuk ke error", err)
}
