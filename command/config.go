package command

import (
	"dns-client/daemon"
	"gopkg.in/alecthomas/kingpin.v2"
	"io/ioutil"
	"log"
)

type generateConfig struct {
	DirPath string
}

func (c *generateConfig) run(*kingpin.ParseContext) error {
	err := ioutil.WriteFile(c.DirPath+"config.toml", []byte(daemon.ConfigTemplate()), 0644)
	if err != nil {
		return err
	}
	log.Println("config file to generate,path: " + c.DirPath)
	return nil
}

func registryGenerate(app *kingpin.Application) {
	c := new(generateConfig)
	cmd := app.Command("gen", "generate config file").Action(c.run)
	cmd.Flag("dir", "config file path").Short('d').Default("").StringVar(&c.DirPath)
}
