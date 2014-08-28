gopack
======

Package Manager in Go Language
BSD-License applies, see LICENSE for more information.

##### Current Version: 0.1 Aplha
Running in Alpha, and not at all useful. At this stage it is just me tinkering. Haven't even acted professional yet.

### Purpose
The idea is to get it to install source packages, defined in a database, with install-scripts based on LinuxFromScratch.

### Overview
- Documentation (WIP): http://godoc.org/github.com/broeman/gopack

### Features
**HEAD (0.1 alpha)**
- Installs a package
- Shows installed packages
- Version/dependencies model

**WIP (road to 0.1 beta)**
- Uninstalls a package
- SQL / PostgresSQL implementation
- Updates packages

**WIP (road to 0.1)**
- Linkage to install script
- Running a source installation: configure, make, make tests

### System Requirements
- A system that wants system packages, like *NIX
- Go Language

### Installation
You shouldn't, but do it the go way:

```
$ go get github.com/broeman/gopack
```

Make sure that you have set $GOPATH and source it in your PATH (e.g. ~/.bashrc):
```
export GOPATH="$HOME/go"
export PATH=$PATH:$GOPATH/bin
```

### About
Go Packge is written by [Jesper Brodersen](http://jesperbrodersen.dk)