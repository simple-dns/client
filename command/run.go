package command

import (
	"context"
	"dns-client/command/path"
	"dns-client/daemon"
	"gopkg.in/alecthomas/kingpin.v2"
)

type runConfig struct {
	configFile string
}

func (c *runConfig) run(*kingpin.ParseContext) error {
	path.SetConfigPath(c.configFile)
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()
	return daemon.Run(ctx)
}

func registryRun(app *kingpin.Application) {
	c := new(runConfig)
	cmd := app.Command("run", "run dns config").Action(c.run)
	cmd.Flag("config", "config file").Short('c').Default(path.DefaultConfigPath()).StringVar(&c.configFile)
}
