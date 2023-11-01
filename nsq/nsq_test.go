package nsq_test

import (
	"fmt"
	"testing"

	vnsq "github.com/nsqio/go-nsq"
	"github.com/redpkg/webkit/nsq"
	"github.com/stretchr/testify/assert"
)

func TestNewProducer(t *testing.T) {
	assert := assert.New(t)

	producer, err := nsq.NewProducer(nsq.ConfigProducer{
		Host: "localhost",
		Port: 4150,
	})
	if !assert.NoError(err) {
		return
	}

	fmt.Printf("%+v\n", producer)
}

func TestNewConsumer(t *testing.T) {
	assert := assert.New(t)

	consumer, err := nsq.NewConsumer(nsq.ConfigConsumer{
		Host:        "localhost",
		Port:        4161,
		MaxAttempts: 5,
		MaxInFlight: 1,
	}, "test_topic", "test_channel", &handler{})
	if !assert.NoError(err) {
		return
	}

	fmt.Printf("%+v\n", consumer)
}

type handler struct{}

func (h *handler) HandleMessage(message *vnsq.Message) error {
	return nil
}
