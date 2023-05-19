package redis_test

import (
	"fmt"
	"testing"

	"github.com/redpkg/webkit/redis"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	assert := assert.New(t)

	redis, err := redis.New(redis.Config{
		Host:     "localhost",
		Port:     6379,
		Password: "",
		DB:       0,
	})
	if !assert.NoError(err) {
		return
	}

	fmt.Printf("%+v\n", redis)
}
