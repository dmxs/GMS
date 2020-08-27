package conf

import (
	"errors"
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var (
	Conf *Config
)

func InitConfig() error {
	appPath, _ := getCurrentPath()
	err := LoadGlobalConfig(appPath + string(os.PathSeparator) + "config.toml")
	return err
}

// LoadGlobalConfig 加载全局配置
func LoadGlobalConfig(fpath string) error {
	c, err := ParseConfig(fpath)
	if err != nil {
		return err
	}
	Conf = c
	return nil
}

// GetGlobalConfig 获取全局配置
func GetGlobalConfig() *Config {
	if Conf == nil {
		return &Config{}
	}
	return Conf
}

// ParseConfig 解析配置文件
func ParseConfig(fpath string) (*Config, error) {
	var c Config
	_, err := toml.DecodeFile(fpath, &c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

//Config .
type Config struct {
	MySQL    MySQL `toml:"mysql"`
	Etcd     Etcd  `toml:"etcd"`
	Micro    Micro `toml:"micro"`
	Http     Http  `toml:"http"`
}

type Http struct {
	Port int `toml:"port"`
}

//MySQL .
type MySQL struct {
	Debug      bool   `toml:"debug"`
	Host       string `toml:"host"`
	Port       int    `toml:"port"`
	User       string `toml:"user"`
	Password   string `toml:"password"`
	DBName     string `toml:"db_name"`
	Parameters string `toml:"parameters"`
}

//DSN .
func (a MySQL) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		a.User, a.Password, a.Host, a.Port, a.DBName, a.Parameters)
}


//Etcd .
type Etcd struct {
	Addrs            []string `toml:"addrs"`
	RegisterTTL      int      `toml:"register_ttl"`
	RegisterInterval int      `toml:"register_interval"`
}

//Micro .
type Micro struct {
	Name    string `toml:"name"`
	Version string `toml:"version"`
}

func getCurrentPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	path, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}
	i := strings.LastIndex(path, "/")
	if i < 0 {
		i = strings.LastIndex(path, "\\")
	}
	if i < 0 {
		return "", errors.New("error: Can't find  or ")
	}
	return string(path[0 : i+1]), nil
}
