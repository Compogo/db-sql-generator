package db_sql_generator

import (
	"fmt"

	"github.com/Compogo/compogo/logger"
	"github.com/doug-martin/goqu/v9"
)

// NewGenerator creates a new goqu dialect wrapper for the configured database driver.
// It performs the following steps:
//  1. Looks up the dialect alias for the configured driver
//  2. Returns a goqu.DialectWrapper ready for building SQL queries
//
// The generator should be used in repositories to construct type-safe SQL queries.
//
// Example:
//
//	type UserRepository struct {
//	    db        db_client.Client
//	    generator *goqu.DialectWrapper
//	}
//
//	func (r *UserRepository) GetUsers(ctx context.Context) ([]User, error) {
//	    query := r.generator.From("users").Where(goqu.C("active").IsTrue())
//	    sql, _, _ := query.ToSQL()
//	    return r.db.QueryContext(ctx, sql)
//	}
func NewGenerator(config *Config, informer logger.Informer) (*goqu.DialectWrapper, error) {
	alias, err := aliases.Get(config.Driver)
	if err != nil {
		return nil, fmt.Errorf("[db-sql-generator] get driver '%s' dialect alias failed '': %w", config.Driver, err)
	}

	informer.Infof("[db-sql-generator] usage dialect '%s' for driver '%s'", alias, config.Driver.String())

	wrapper := goqu.Dialect(alias)

	return &wrapper, nil
}
