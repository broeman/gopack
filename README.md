gopack
======

** Package Manager in Go Language **
BSD-License applies, see LICENSE for more information.

### Current Version: 0.1a
Running in alpha, and not that useful. At this stage it is just me tinkering. Haven't even acted professional yet.

### Purpose
This is a library for pm-get and pm-tools, which are there to let users install packages.More ideas is on the table, but I'll take one step at the time. For this time being, it is just me wanting to do something productive and learning to make something useful for myself.

### Overview
- [Documentation (WIP)](http://godoc.org/github.com/broeman/gopack)

### Features
**HEAD (0.1 alpha)**
- Installs a package
- Shows installed packages
- Version/dependencies model
- Uninstalls a package
- SQLite3 implementation

**WIP (road to 0.1b)**
- Updates packages
- Linkage to install script
- Running a source installation: configure, make, make tests

**WIP (road to 0.1)**
- Package Manager Management Tools: CRUD, settings

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
