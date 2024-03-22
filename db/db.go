package db

import (
	"fmt"
	"net/url"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

type Config struct {
	Writer          ConfigNode    `mapstructure:"writer"`
	Reader          ConfigNode    `mapstructure:"reader"`
	Database        string        `mapstructure:"database"`
	Timezone        string        `mapstructure:"timezone"`
	ConnMaxIdleTime time.Duration `mapstructure:"conn_max_idle_time"`
	ConnMaxLifetime time.Duration `mapstructure:"conn_max_lifetime"`
	MaxIdleConns    int           `mapstructure:"max_idle_conns"`
	MaxOpenConns    int           `mapstructure:"max_open_conns"`
}

type ConfigNode struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

// New database instance
func New(conf Config) (*gorm.DB, error) {
	db, err := gorm.Open(newDialector(conf.Writer, conf), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		DisableAutomaticPing:   true,
		TranslateError:         true,
	})
	if err != nil {
		return nil, err
	}

	replicas := []gorm.Dialector{}
	if conf.Reader.Host != "" {
		replicas = append(replicas, newDialector(conf.Reader, conf))
	}

	resolver := dbresolver.Register(dbresolver.Config{
		Replicas: replicas,
	}).
		SetConnMaxIdleTime(conf.ConnMaxIdleTime).
		SetConnMaxLifetime(conf.ConnMaxLifetime).
		SetMaxIdleConns(conf.MaxIdleConns).
		SetMaxOpenConns(conf.MaxOpenConns)

	if err := db.Use(resolver); err != nil {
		return nil, err
	}

	return db, nil
}

func newDialector(confNode ConfigNode, conf Config) gorm.Dialector {
	return mysql.Open(fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=True&loc=%s",
		confNode.Username,
		confNode.Password,
		confNode.Host,
		confNode.Port,
		conf.Database,
		url.QueryEscape(conf.Timezone),
	))
}
