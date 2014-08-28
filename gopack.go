// GoPack is a package manger written in Go Language
// Author: Jesper Brodersen 2014
// BSD-License applies: http://opensource.org/licenses/bsd-license.php

// !! For experimental use only, not at all in a state of usefulness !!

package main

import (
	"github.com/broeman/gopack/cmd" // using CLI command args
	"github.com/codegangsta/cli"
	"os"
	"runtime"
)

const APP_VER = "0.1 Alpha"

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())

}

func main() {
	app := cli.NewApp()
	app.Name = "gopack"
	app.Usage = "Go Package Manager"
	app.Version = APP_VER
	app.Commands = []cli.Command{
		cmd.Install,   // install a package
		cmd.Installed, // shows current installed packages
	}

	app.Run(os.Args)
}
