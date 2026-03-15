package db_sql_generator

import (
	"github.com/Compogo/compogo/component"
	"github.com/Compogo/compogo/container"
)

// Component is a ready-to-use Compogo component that provides a SQL query builder.
// It automatically:
//   - Registers Config and Generator in the DI container
//   - Applies configuration during Configuration phase
//   - Makes the generator available for repositories and services
//
// Usage:
//
//	compogo.WithComponents(
//	    db_client.Component,           // database client (provides driver name)
//	    db_sql_generator.Component,    // SQL query builder
//	    // ... driver components (postgres, mysql, etc.)
//	)
//
// The driver name is automatically taken from db-client configuration
// and used to select the appropriate SQL dialect.
//
// In your repository component:
//
//	type UserRepository struct {
//	    db        db_client.Client
//	    generator *goqu.DialectWrapper
//	}
//
//	func NewUserRepository(container container.Container) (*UserRepository, error) {
//	    var repo UserRepository
//	    err := container.Invoke(func(db db_client.Client, gen *goqu.DialectWrapper) {
//	        repo = UserRepository{db: db, generator: gen}
//	    })
//	    return &repo, err
//	}
var Component = &component.Component{
	Init: component.StepFunc(func(container container.Container) error {
		return container.Provides(
			NewConfig,
			NewGenerator,
		)
	}),
	Configuration: component.StepFunc(func(container container.Container) error {
		return container.Invoke(Configuration)
	}),
}
