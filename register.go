package db_sql_generator

import (
	"github.com/Compogo/types/linker"
)

// aliases — хранилище соответствий между драйверами БД и диалектами goqu.
// Ключ — имя драйвера, значение — имя диалекта.
//
// Доступные диалекты goqu:
//   - "mysql"
//   - "postgres"
//   - "sqlite3"

var aliases = linker.NewLinker[string, string](linker.KeyStringNormalizer[string]())

// Registration регистрирует соответствие между драйвером БД и диалектом goqu.
// Должна вызываться в init() каждого пакета драйвера.
//
// Пример регистрации MySQL-драйвера:
//
//	func init() {
//	    db_sql_generator.Registration("mysql", "mysql")
//	}
//
// Пример регистрации PostgreSQL-драйвера:
//
//	func init() {
//	    db_sql_generator.Registration("postgres", "postgres")
//	}
func Registration(driverName string, alias string) {
	aliases.Add(driverName, alias)
}
