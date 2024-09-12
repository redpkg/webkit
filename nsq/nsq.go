package nsq

import (
	"fmt"
	"strings"
	"time"

	vnsq "github.com/nsqio/go-nsq"
)

type Config struct {
	Producer ConfigProducer `mapstructure:"producer"`
	Consumer ConfigConsumer `mapstructure:"consumer"`
	LogLevel string         `mapstructure:"log_level"`
}

func (c Config) level() vnsq.LogLevel {
	switch strings.ToLower(c.LogLevel) {
	case "debug":
		return vnsq.LogLevelDebug
	case "info":
		return vnsq.LogLevelInfo
	case "warn":
		return vnsq.LogLevelWarning
	case "error":
		return vnsq.LogLevelError
	default:
		return vnsq.LogLevelWarning
	}
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

func NewProducer(conf Config) (*vnsq.Producer, error) {
	producer, err := vnsq.NewProducer(buildAddress(conf.Producer.Host, conf.Producer.Port), vnsq.NewConfig())
	if err != nil {
		return nil, err
	}

	producer.SetLoggerLevel(conf.level())

	return producer, nil
}

func NewConsumer(conf Config, topic, channel string, handler vnsq.Handler) (*vnsq.Consumer, error) {
	c := vnsq.NewConfig()
	c.MaxAttempts = conf.Consumer.MaxAttempts
	c.MaxInFlight = conf.Consumer.MaxInFlight
	c.MaxRequeueDelay = conf.Consumer.MaxRequeueDelay
	c.DefaultRequeueDelay = conf.Consumer.DefaultRequeueDelay

	consumer, err := vnsq.NewConsumer(topic, channel, c)
	if err != nil {
		return nil, err
	}

	consumer.SetLoggerLevel(conf.level())
	consumer.AddHandler(handler)

	if err = consumer.ConnectToNSQLookupd(buildAddress(conf.Consumer.Host, conf.Consumer.Port)); err != nil {
		return nil, err
	}

	return consumer, nil
}

func buildAddress(host string, port int) string {
	return fmt.Sprintf("%s:%d", host, port)
}
