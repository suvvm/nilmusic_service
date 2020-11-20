package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var (
	AppConfig Config
)

type Config struct {
	DBConfig MysqlConfig `yaml:"DbConf"`
}

type MysqlConfig struct {
	DSNTemplate  string `yaml:"DSNTemplate"` // template for dsn, all parameter must be set except ip
	Username     string `yaml:"Username"`
	Password     string `yaml:"Password"`
	DBName       string `yaml:"DBName"`
	Hostname	 string `yaml:"Hostname"`
	Port		 string `yaml:"Port"`
	ConsulName   string `yaml:"ConsulName"` // consul name
	Timeout      string `yaml:"Timeout"`    // connect timeout
	ReadTimeout  string `yaml:"ReadTimeout"`
	WriteTimeout string `yaml:"WriteTimeout"`
	MaxIdle      int    `yaml:"MaxIdle"`
	MaxOpen      int    `yaml:"MaxOpen"`
}

// GetDSN 根据mysql配置信息构造DSN
func (config *MysqlConfig) GetDSN() string {
	dsn := fmt.Sprintf(config.DSNTemplate, config.Username, config.Password, config.Hostname,
		config.Port, config.DBName, config.Timeout, config.ReadTimeout, config.WriteTimeout)
	return dsn
}

// Init 根据文件名读取配置文件到内存
//
// 入参
//	filename string	// 配置文件名
// 返回
//	*Config //	配置结构指针
func Init(filename string) *Config {
	log.Printf("init config file: %s", filename)
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(fmt.Sprintf("read config file error. configFile=%s err=%s", filename, err))
	}
	err = yaml.Unmarshal(buf, &AppConfig)
	if err != nil {
		panic(fmt.Sprintf("unmarshal config err=%s", err))
	}
	log.Printf("Config = %#v", AppConfig)
	return &AppConfig
}
