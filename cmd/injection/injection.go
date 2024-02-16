package injection

import (
	"go.uber.org/dig"
)

// container instance
var container *dig.Container

// user
func ProvideComponents() {
	// create a new container
	container = dig.New()
	// user provider injection
	provideUserComponents(container)
}

// init container
func InitComponents() error {
	return container.Invoke(
		initUserComponents(),
	)
}
