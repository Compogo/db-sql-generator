package db_sql_generator

import (
	"github.com/Compogo/db-client/driver"
	"github.com/Compogo/types/linker"
)

var (
	// aliases stores dialect aliases for each registered database driver.
	// The linker associates each Driver with its corresponding goqu dialect string.
	aliases = linker.NewLinker[driver.Driver, string]()
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
func Registration(d driver.Driver, alias string) {
	aliases.Add(d, alias)
}
