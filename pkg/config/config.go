package config

import (
	"errors"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
)

var (
	FileTypeYaml = "yaml"
	FileTypeJson = "json"
	FileTypeToml = "toml"
)

type Config struct {
	env        string
	configDir  string
	configType string // file type, eg: yaml, json, toml, default is yaml
	val        map[string]*viper.Viper
}

type Option func(*Config)

// WithFileType config file type
func WithFileType(fileType string) Option {
	return func(c *Config) {
		c.configType = fileType
	}
}

// WithEnv env var
func WithEnv(name string) Option {
	return func(c *Config) {
		c.env = name
	}
}

func New(cfgDir string, opts ...Option) *Config {
	// must set config dir
	if cfgDir == "" {
		panic("config dir is not set")
	}
	c := Config{
		configDir:  cfgDir,
		configType: FileTypeYaml,
		val:        make(map[string]*viper.Viper),
	}
	for _, opt := range opts {
		opt(&c)
	}

	return &c
}

func (c *Config) Load(filename string, cfgType string, val interface{}) error {
	env := GetEnvString("APP_ENV", "")
	path := filepath.Join(c.configDir, env)
	if c.env != "" {
		path = filepath.Join(c.configDir, c.env)
	}

	v := viper.New()
	// 设置配置文件的路径
	v.AddConfigPath(path)
	// 设置配置文件的名字
	v.SetConfigName(filename)
	// 设置配置文件的类型
	v.SetConfigType(c.configType)
	if cfgType != "" {
		v.SetConfigType(cfgType)
	}

	// 寻找配置文件并读取
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return errors.New("config file not found")
		}
		return err
	}

	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("Config file changed: %s", e.Name)
	})

	c.val[filename] = v

	err := v.Unmarshal(&val)
	if err != nil {
		return err
	}

	return nil
}

func GetEnvString(key string, defaultValue string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}
	return val
}
