package db_sql_generator

import (
	"github.com/Compogo/compogo"
)

// Component — компонент SQL-генератора для Compogo.
// Регистрирует конфигурацию и генератор в DI-контейнере.
//
// Пример подключения:
//
//	app.AddComponents(&db_sql_generator.Component)
//
//	var gen *goqu.DialectWrapper
//	container.Invoke(func(g *goqu.DialectWrapper) { gen = g })
//
//	// Генерация запроса
//	query, args, _ := gen.From("users").Where(goqu.C("id").Eq(1)).ToSQL()
//	rows, err := client.Query(query, args...)
var Component = compogo.Component{
	Init: compogo.StepFunc(func(container compogo.Container) error {
		return container.Provides(
			NewConfig,
			NewGenerator,
		)
	}),
	Configuration: compogo.StepFunc(func(container compogo.Container) error {
		return container.Invoke(Configuration)
	}),
}
