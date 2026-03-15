package db_sql_generator

import (
	"github.com/Compogo/types/linker"
)

var (
	// aliases stores dialect aliases for each registered database driver.
	// The linker associates each Driver with its corresponding goqu dialect string.
	aliases = linker.NewLinker[Driver, string]()
)

// Registration registers a new database driver and its goqu dialect alias.
// This function should be called during driver package initialization.
// The dialect will then be available for use via the generator component.
//
// Example (in postgres driver):
//
//	func init() {
//	    db_sql_generator.Registration(Postgres, "postgres")
//	}
func Registration(d Driver, alias string) {
	aliases.Add(d, alias)
}

// Driver represents a database driver identifier (e.g., "postgres", "mysql").
// It should match the driver name used in db-client and db-migrator to ensure consistency.
type Driver string

// String returns the driver name as a string.
func (d Driver) String() string {
	return string(d)
}
