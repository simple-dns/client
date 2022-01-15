package path

import (
	"os"
	"path/filepath"
)

const envName = "DnsConfigPath"

func SetConfigPath(path string) {
	os.Setenv(envName, path)

}
func ConfigPath() string {
	path := os.Getenv(envName)
	if path == "" {
		path = DefaultConfigPath()
	}
	return path
}

func ConfigDirPath() string {
	dir, _ := filepath.Split(ConfigPath())
	return dir
}

func DefaultConfigPath() string {
	return filepath.Join(DefaultConfigDirPath(), "config.toml")
}
