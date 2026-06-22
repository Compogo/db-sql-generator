package db_sql_generator

import (
	"fmt"

	"github.com/Compogo/compogo"
	"github.com/doug-martin/goqu/v9"
)

// NewGenerator создаёт новый SQL-генератор для указанного драйвера.
// Использует goqu.Dialect для генерации SQL-запросов в нужном диалекте.
//
// Пример использования:
//
//	var gen *goqu.DialectWrapper
//	container.Invoke(func(g *goqu.DialectWrapper) { gen = g })
//
//	// Генерация SELECT
//	sql, args, _ := gen.From("users").Where(goqu.C("age").Gt(18)).ToSQL()
//	// SELECT * FROM users WHERE age > 18
//
//	// Генерация INSERT
//	sql, args, _ := gen.Insert("users").Rows(goqu.Record{"name": "John", "age": 30}).ToSQL()
//	// INSERT INTO users (name, age) VALUES ('John', 30)
//
//	// Генерация UPDATE
//	sql, args, _ := gen.Update("users").Set(goqu.Record{"age": 31}).Where(goqu.C("id").Eq(1)).ToSQL()
//	// UPDATE users SET age = 31 WHERE id = 1
func NewGenerator(config *Config, logger compogo.Logger) (*goqu.DialectWrapper, error) {
	alias, err := aliases.Get(config.Driver)
	if err != nil {
		return nil, fmt.Errorf("[Database][sql-generator] get driver '%s' dialect alias failed '': %w", config.Driver, err)
	}

	logger.GetLogger("Database").GetLogger("sql-generator").Infof("usage dialect '%s' for driver '%s'", alias, config.Driver)

	return new(goqu.Dialect(alias)), nil
}
