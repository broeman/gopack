// Copyright 2014 Jesper Brodersen. All rights reserved.
// This code is BSD-licensed, see LICENSE file.

// Package structure with Versions
package pack

type Package struct {
	name         string
	versions     []Version
	versionRegEx string
	installed    bool
}

type Version struct {
	id           string
	description  string
	state        string
	dependencies []Package
}

// name getter
func (p Package) Name() string {
	return p.name
}

// Currentversion getter: TODO: use Regex for multiple stable
func (p Package) CurrentVersion() Version {
	if len(p.versions) > 0 {
		for version := range p.versions {
			if p.versions[version].state == "stable" {
				return p.versions[version]
			}
		}
	}
	return Version{}
}

// version getter
func (p Package) Version() string {
	return p.CurrentVersion().id + ", " + p.CurrentVersion().state
}

// installed getter
func (p Package) Installed() bool {
	return p.installed
}

// description getter
func (p Package) Description() string {
	return p.CurrentVersion().description
}

// dependencies printer
func (p Package) Dependencies() string {
	var result string
	dep := p.CurrentVersion().dependencies
	for i := range dep {
		result += dep[i].Name()
	}
	return result
}

/* Package
 *
 */

// package constructor
func NewPackage(name string, versions []Version, versionRegEx string, installed bool) Package {
	// TODO: insert into database
	return Package{name, versions, versionRegEx, installed}
}

// TODO: get from from database
func RetrievePackage(name string) {
}

// TODO: delete from database
func (p *Package) Delete() {
}

// TODO: update database
func (p *Package) UpdatePackage() {
}

// installed setter
func (p *Package) SetInstalled(setting bool) {
	p.installed = setting
}

/* Version
 *
 */

// version constructor
func NewVersion(id string, description string, state string, dependencies []Package) Version {
	// TODO: insert into database
	return Version{id, description, state, dependencies}
}

// TODO: get from database
func RetrieveVersion(id string) {
}

/* Placeholders
 * To be removed when database is implemented
 */

// Slices of packages, placeholder until database is up running
func PackageDB() (packages []Package) {
	return packages
}

// Slices of versions, placeholder until database is up running
func NewVersions() (versions []Version) {
	return versions
}
