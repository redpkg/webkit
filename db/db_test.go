package db_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/redpkg/webkit/db"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	assert := assert.New(t)

	db, err := db.New(db.Config{
		Writer: db.ConfigNode{
			Host:     "localhost",
			Port:     3306,
			Username: "root",
			Password: "password",
		},
		Reader: db.ConfigNode{
			Host:     "localhost",
			Port:     3306,
			Username: "root",
			Password: "password",
		},
		Database:        "test",
		Timezone:        "UTC",
		ConnMaxIdleTime: time.Minute * 5,
		ConnMaxLifetime: time.Hour,
		MaxIdleConns:    5,
		MaxOpenConns:    10,
	})
	if !assert.NoError(err) {
		return
	}

	fmt.Printf("%+v\n", db)
}
