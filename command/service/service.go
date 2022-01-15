package service

import "gopkg.in/alecthomas/kingpin.v2"

func Register(app *kingpin.Application) {
	cmd := app.Command("service", "manages the runner service")
	registerInstall(cmd)
	registerRun(cmd)
	registerStart(cmd)
	registerStop(cmd)
	registerUninstall(cmd)
}
