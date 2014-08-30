// Copyright 2014 Jesper Brodersen. All rights reserved.
// This code is BSD-licensed, see LICENSE file.

// GoPack is a package manger written in Go Language
// !! For experimental use only, not at all in a state of usefulness !!
package main

import (
	"github.com/broeman/gopack/cmd" // using CLI command args
	"github.com/codegangsta/cli"
	"os"
	"runtime"
)

const APP_VER = "0.1 Alpha"

// const MAN_VER = "0.01 Alpha"

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	app := cli.NewApp()
	app.Name = "pm"
	app.Usage = "Package Manager in Go"
	app.Version = APP_VER
	app.Commands = []cli.Command{
		cmd.Install,   // install a package
		cmd.UnInstall, // uninstall a package
		cmd.Show,      // show package
		cmd.Installed, // shows current installed packages, placeholder
		//cmd.Update,	// update packages
		cmd.Init, // placeholder initialization
	}
	app.Run(os.Args)

	// Future implementation of CRUD Management
	// appMan := cli.NewApp()
	// appMan.Name = "pm-tools"
	// app.Usage = "Go Package Manager: Management Tools"
	// app.Version = MAN_VER
	// app.Commands = []cli.Command{
	// 	// CRUD implementation
	//	cmd.Init,		   // initialize database if not exist
	// 	cmd.AddPackage,    // adding a package to the database
	// 	cmd.ShowPackage,   // showing a package from the database
	// 	cmd.EditPackage,   // editing a package from the database
	// 	cmd.RemovePackage, // removing a package from the database
	// }
	// appMan.Run(os.Args)

}
