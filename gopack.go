// GoPack is a package manger written in Go Language
// Author: Jesper Brodersen 2014
// BSD-License applies: http://opensource.org/licenses/bsd-license.php

// !! For experimental use only, not at all in a state of usefulness !!

package main

import (
	. "./pack" // using Package struct
	"flag"
	"fmt"
	"os"
)

var packages = PackageDB()

func main() {
	supdeb := PackageDB()
	supr := NewPackage("Super", NewVersion("1.0-1", "SuperTest", supdeb), true)
	supdeb = append(supdeb, supr)
	packages = append(packages, supr)
	tester := NewPackage("Test", NewVersion("0.1-alpha", "Testpackage", supdeb), true)
	packages = append(packages, tester)

	flag.Usage = usage
	flag.Parse()

	args := flag.Args()
	if len(args) > 0 {
		switch os.Args[1] {
		case "install":
			if len(args) > 1 {
				curpackage := os.Args[2]
				notfound := true
				for i := range packages {
					item := packages[i]
					if curpackage == item.Name() {
						if item.Installed() {
							fmt.Println("Package is already installed")
						} else {
							item.SetInstalled(true)
							fmt.Println("Installed package:", curpackage)
							notfound = false
							break
						}
					}
				}
				if notfound {
					fmt.Println(curpackage, "was not found")
				}

			} else {
				fmt.Println("Please refer which package you want to install")
			}
		case "installed":
			for i := range packages {
				item := packages[i]
				if item.Installed() {
					fmt.Println("Package:", item.Name(), "\nVersion:", item.Version(), "\nDescription:", item.Description(), "\nDependencies:", item.Dependencies(), "\n")
				}
			}
		default:
		}
	} else {
		usage()
	}
}

func usage() {
	fmt.Println("GoPack 0.1: Package Manager in Go")
	fmt.Println("Usage: gopack <command> <options>")
	fmt.Println("Commands:")
	fmt.Println("install - install a package")
	fmt.Println("installed - show installed packages")
}
