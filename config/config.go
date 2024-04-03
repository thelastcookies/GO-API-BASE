package config

import (
	"time"
)

var (
	// Conf AppConf App 配置
	Conf *Config
)

// Config AppConfig App global config
// nolint
type Config struct {
	Name              string
	Version           string
	Mode              string
	PprofPort         string
	URL               string
	JwtSecret         string
	JwtTimeout        int
	SSL               bool
	CtxDefaultTimeout time.Duration
	CSRF              bool
	Debug             bool
	EnableTrace       bool
	EnablePprof       bool
	HTTP              ServerConfig
	GRPC              ServerConfig
}

// ServerConfig server config.
type ServerConfig struct {
	Network      string
	Addr         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type MysqlConfig struct {
	DBName          string
	Addr            string
	UserName        string
	Password        string
	ShowLog         bool
	MaxIdleConn     int
	MaxOpenConn     int
	ConnMaxLifeTime time.Duration
	SlowThreshold   time.Duration
}
