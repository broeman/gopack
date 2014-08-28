package cmd

import (
	"fmt"
	"github.com/broeman/gopack/pack" // using Package struct
	"github.com/codegangsta/cli"
	"os"
)

var Install = cli.Command{
	Name:        "install",
	Usage:       "Installs a package",
	Description: `Installs a package, that isn't already installed`,
	Action:      runInstall,
	Flags:       []cli.Flag{},
}

var Installed = cli.Command{
	Name:        "installed",
	Usage:       "Shows installed packages",
	Description: `Shows all currently installed packages`,
	Action:      runInstalled,
	Flags:       []cli.Flag{},
}

// placeholders
var packages = pack.PackageDB()

func runPlaceholder() {
	dependency := pack.PackageDB()    // empty depencies
	supversions := pack.NewVersions() // empty versions
	supversions = append(supversions, pack.NewVersion("0.1b", "SuperTest 0.1 alpha", "alpha", dependency))
	supversions = append(supversions, pack.NewVersion("1.0", "SuperTest 1.0", "stable", dependency))

	supr := pack.NewPackage("Super", supversions, "[0-9].[0-9][a-z]", true)
	packages = append(packages, supr)

	dependency = append(dependency, supr) // adding dependency for Test
	testversions := pack.NewVersions()    // empty versions
	testversions = append(testversions, pack.NewVersion("0.1a", "TestPackage 0.1a", "alpha", dependency))
	testversions = append(testversions, pack.NewVersion("0.2", "TestPackage 0.2 Useful", "stable", dependency))

	tester := pack.NewPackage("Test", testversions, "[0-9].[0-9][a-z]", true)
	packages = append(packages, tester)
}

func runInstall(ctx *cli.Context) {
	runPlaceholder()
	if len(ctx.Args()) != 1 {
		fmt.Println("You need to specify which package you want to install")
		os.Exit(2)
	}

	curpackage := ctx.Args().First()
	notfound := true
	for i := range packages {
		item := packages[i]
		if curpackage == item.Name() {
			if item.Installed() {
				fmt.Println("Package is already installed")
				notfound = false
				break
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
}

func runInstalled(*cli.Context) {
	runPlaceholder()
	for i := range packages {
		item := packages[i]
		if item.Installed() {
			fmt.Println("Package:", item.Name(), "\nVersion:", item.Version(), "\nDescription:", item.Description(), "\nDependencies:", item.Dependencies(), "\n")
		}
	}
}
