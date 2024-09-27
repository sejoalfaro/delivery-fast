package domain

// Deploy is a struct that represents a deploy
type Deploy struct {
	ID          string
	Name        string
	Application string
	Environment string
	Domain      string
	TraefikRule string
	Version     string
}

// NewDeploy is a function that creates a new deploy
func NewDeploy(name, application, environment, domain, traefikRule, version string) *Deploy {
	return &Deploy{
		Name:        name,
		Application: application,
		Environment: environment,
		Domain:      domain,
		TraefikRule: traefikRule,
		Version:     version,
	}
}
