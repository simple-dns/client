// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE file.

//go:build !windows
// +build !windows

package path

import (
	"os"
	"os/user"
)

// function returns the current user
var getUser = user.Current

// function returns the current uid
var getuid = os.Getuid

// helper function returns the default configuration path
// for the drone configuration.
func DefaultConfigDirPath() string {
	u, err := getUser()
	if err != nil || getuid() == 0 {
		return "/etc/dns-client/"
	}
	return u.HomeDir + "config/"
}
