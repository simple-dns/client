// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE file.

package service

import (
	"log"

	"dns-client/daemon/service"

	"gopkg.in/alecthomas/kingpin.v2"
)

type uninstallCommand struct {
	config service.Config
}

func (c *uninstallCommand) run(*kingpin.ParseContext) error {
	log.Printf("uninstalling service %s\n", c.config.Name)
	s, err := service.New(c.config)
	if err != nil {
		return err
	}
	return s.Uninstall()
}

func registerUninstall(cmd *kingpin.CmdClause) {
	c := new(uninstallCommand)
	s := cmd.Command("uninstall", "uninstall the service").
		Action(c.run)

	s.Flag("name", "service name").
		Default(service.DefaultName).
		StringVar(&c.config.Name)

	s.Flag("desc", "service description").
		Default(service.DefaultDesc).
		StringVar(&c.config.Desc)

	s.Flag("username", "windows account username").
		Default("").
		StringVar(&c.config.Username)

	s.Flag("password", "windows account password").
		Default("").
		StringVar(&c.config.Password)

}
