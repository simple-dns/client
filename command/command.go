package command

import (
	"dns-client/command/service"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
)

var version = "0.0.0"

func Command() {
	app := kingpin.New("dns-client", "dns report client")
	registryGenerate(app)
	registryRun(app)
	service.Register(app)
	kingpin.Version(version)
	kingpin.MustParse(app.Parse(os.Args[1:]))
}
