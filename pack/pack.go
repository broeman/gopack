package pack

type Package struct {
	name      string
	version   Version
	installed bool
}

type Version struct {
	id           string
	description  string
	dependencies []Package
}

// name getter
func (p Package) Name() string {
	return p.name
}

// version getter
func (p Package) Version() string {
	return p.version.id
}

// installed getter
func (p Package) Installed() bool {
	return p.installed
}

// description getter
func (p Package) Description() string {
	return p.version.description
}

// dependencies getter
func (p Package) Dependencies() string {
	var result string
	dep := p.version.dependencies
	for i := range dep {
		result += dep[i].Name()
	}
	return result
}

// package constructor
func NewPackage(name string, version Version, installed bool) Package {
	return Package{name, version, installed}
}

// installed setter
func (p *Package) SetInstalled(setting bool) {
	p.installed = setting
}

// Slices of packages
func PackageDB() (packages []Package) {
	// TODO: get access to database
	return packages
}

// version constructor
func NewVersion(id string, description string, dependencies []Package) Version {
	return Version{id, description, dependencies}
}
