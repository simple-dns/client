package daemon

import (
	"dns-client/command/path"
	"fmt"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"sync"
)

type Config struct {
	Server server `toml:"server"`
	Domain domain `toml:"domain"`
}

type server struct {
	Proto   string `toml:"proto"`
	Host    string `toml:"host"`
	Port    int    `toml:"Port"`
	Address string `toml:"address"`
	Token   string `toml:"token"`
}
type domain struct {
	Suffix     string `toml:"suffix"`
	DomainFile string `toml:"domainFile"`
}

var configCache struct {
	ConfigSingle *Config
	Lock         sync.Mutex
}

func loadFile(path string) error {
	defer configCache.Lock.Unlock()
	configCache.Lock.Lock()
	if configCache.ConfigSingle == nil {
		configCache.ConfigSingle = &Config{}
	}
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	_, err = toml.Decode(string(file), configCache.ConfigSingle)
	if configCache.ConfigSingle.Domain.DomainFile == "" {
		configCache.ConfigSingle.Domain.DomainFile = "domain"
	}
	if configCache.ConfigSingle.Server.Port == 0 {
		configCache.ConfigSingle.Server.Port = 80
	}
	if configCache.ConfigSingle.Domain.Suffix != "" {
		configCache.ConfigSingle.Domain.Suffix = "." + configCache.ConfigSingle.Domain.Suffix
	}
	configCache.ConfigSingle.Server.Address = fmt.Sprintf(
		"%s://%s:%d",
		configCache.ConfigSingle.Server.Proto,
		configCache.ConfigSingle.Server.Host,
		configCache.ConfigSingle.Server.Port,
	)
	return err
}

func getConfig() (*Config, error) {
	if configCache.ConfigSingle == nil {
		err := loadFile(path.ConfigPath())
		if err != nil {
			return nil, err
		}
	}
	return configCache.ConfigSingle, nil
}

func ConfigTemplate() string {
	temp :=
		`[server]
#server address
proto=""
host=""
port=
token=""
[domain]
#report domain name add suffix 
suffix=""
#file path
domainFile=""
`
	return temp
}
