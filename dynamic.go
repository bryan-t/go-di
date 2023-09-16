package godi

// Singleton describes a singleton implementation of the interface
type Singleton struct {
	InterfacePkg string
	Interface    string
	ImplPkg      string
	Impl         string
}

// Config contains the configuration for godi service provider.
type Config struct {
	Singletons []Singleton
}
