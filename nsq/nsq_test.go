package nsq_test

import (
	"fmt"
	"testing"
	"time"

	vnsq "github.com/nsqio/go-nsq"
	"github.com/redpkg/webkit/nsq"
	"github.com/stretchr/testify/assert"
)

func TestNewProducer(t *testing.T) {
	assert := assert.New(t)

	conf := nsq.Config{
		Producer: nsq.ConfigProducer{
			Host: "localhost",
			Port: 4150,
		},
		LogLevel: "warn",
	}

	producer, err := nsq.NewProducer(conf)
	if !assert.NoError(err) {
		return
	}

	fmt.Printf("%+v\n", producer)
}

func TestNewConsumer(t *testing.T) {
	assert := assert.New(t)

	conf := nsq.Config{
		Consumer: nsq.ConfigConsumer{
			Host:                "localhost",
			Port:                4161,
			MaxAttempts:         5,
			MaxInFlight:         1,
			MaxRequeueDelay:     15 * time.Minute,
			DefaultRequeueDelay: 90 * time.Second,
		},
		LogLevel: "warn",
	}

	consumer, err := nsq.NewConsumer(conf, "test_topic", "test_channel", &handler{})
	if !assert.NoError(err) {
		return
	}

	fmt.Printf("%+v\n", consumer)
}

type handler struct{}

func (h *handler) HandleMessage(message *vnsq.Message) error {
	return nil
}
