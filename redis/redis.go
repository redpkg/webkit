package redis

import (
	"fmt"

	vredis "github.com/redis/go-redis/v9"
)

type Config struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

// New redis instance
func New(conf Config) (*vredis.Client, error) {
	return vredis.NewClient(&vredis.Options{
		Addr:     buildAddress(conf.Host, conf.Port),
		Password: conf.Password,
		DB:       conf.DB,
	}), nil
}

func buildAddress(host string, port int) string {
	return fmt.Sprintf("%s:%d", host, port)
}
