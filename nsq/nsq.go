package nsq

import (
	"fmt"
	"time"

	vnsq "github.com/nsqio/go-nsq"
)

type Config struct {
	Producer ConfigProducer `mapstructure:"producer"`
	Consumer ConfigConsumer `mapstructure:"consumer"`
}

type ConfigProducer struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type ConfigConsumer struct {
	Host                string        `mapstructure:"host"`
	Port                int           `mapstructure:"port"`
	MaxAttempts         uint16        `mapstructure:"max_attempts"`
	MaxInFlight         int           `mapstructure:"max_in_flight"`
	MaxRequeueDelay     time.Duration `mapstructure:"max_requeue_delay"`
	DefaultRequeueDelay time.Duration `mapstructure:"default_requeue_delay"`
}

func NewProducer(conf ConfigProducer) (*vnsq.Producer, error) {
	producer, err := vnsq.NewProducer(buildAddress(conf.Host, conf.Port), vnsq.NewConfig())
	if err != nil {
		return nil, err
	}

	return producer, nil
}

func NewConsumer(conf ConfigConsumer, topic, channel string, handler vnsq.Handler) (*vnsq.Consumer, error) {
	c := vnsq.NewConfig()
	c.MaxAttempts = conf.MaxAttempts
	c.MaxInFlight = conf.MaxInFlight
	c.MaxRequeueDelay = conf.MaxRequeueDelay
	c.DefaultRequeueDelay = conf.DefaultRequeueDelay

	consumer, err := vnsq.NewConsumer(topic, channel, c)
	if err != nil {
		return nil, err
	}

	consumer.AddHandler(handler)

	if err = consumer.ConnectToNSQLookupd(buildAddress(conf.Host, conf.Port)); err != nil {
		return nil, err
	}

	return consumer, nil
}

func buildAddress(host string, port int) string {
	return fmt.Sprintf("%s:%d", host, port)
}
