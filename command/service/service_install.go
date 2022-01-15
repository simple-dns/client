package service

import (
	"dns-client/command/path"
	"dns-client/daemon/service"
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
	"log"
	"os"
)

type installCommand struct {
	config service.Config
}

func (c *installCommand) run(*kingpin.ParseContext) error {
	log.Printf("read configuration %s\n", c.config.ConfigFile)
	log.Printf("installing service %s\n", c.config.Name)
	if _, err := os.Stat(c.config.ConfigFile); err != nil {
		return fmt.Errorf("cannot read configuration")
	}
	s, err := service.New(c.config)
	if err != nil {
		return err
	}
	return s.Install()
}

func registerInstall(cmd *kingpin.CmdClause) {
	c := new(installCommand)
	s := cmd.Command("install", "install the service").
		Action(c.run)

	s.Flag("name", "service name").
		Default(service.DefaultName).
		StringVar(&c.config.Name)

	s.Flag("desc", "service description").
		Short('d').
		Default(service.DefaultDesc).
		StringVar(&c.config.Desc)

	s.Flag("username", "windows account username").
		Short('u').
		Default("").
		StringVar(&c.config.Username)

	s.Flag("password", "windows account password").
		Short('p').
		Default("").
		StringVar(&c.config.Password)

	s.Flag("config", "service configuration file").
		Short('c').
		Default(path.DefaultConfigPath()).
		StringVar(&c.config.ConfigFile)

}
