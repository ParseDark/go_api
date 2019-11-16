package config

import (
	"log"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	Name string
}

func Init(cfg string) error {
	c := Config {
		Name: cfg,
	}

	// 初始化配置
	if err := c.initConfig(); err != nil {
		return err
	}

	// watch file modify and hot load
	c.watchConfig()

	return nil
}

func (c *Config) initConfig() error {
	if c.Name != "" {
		viper.SetConfigFile(c.Name) // if setting config.file then paser config file
	} else {
		viper.AddConfigPath("conf") // setting default dir paht
		viper.SetConfigName("config") // setting default file 
	}

	// setting default file type
	viper.SetConfigType("yaml")
	// setting read setting config varible
	viper.AutomaticEnv()
	// // 读取环境变量的前缀为APISERVER
	viper.SetEnvPrefix("API")

	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	// viper解析配置文件
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}
// 监控配置文件变化并热加载程序
func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func (e fsnotify.Event) {
		log.Printf("Config file change: %s", e.Name)
	})
}