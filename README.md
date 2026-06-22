# Compogo DB SQL Generator

SQL-генератор для фреймворка [Compogo](https://github.com/Compogo/compogo).

На основе [goqu](https://github.com/doug-martin/goqu) предоставляет:

* Генерацию SQL-запросов для различных СУБД (MySQL, PostgreSQL, SQLite)
* Типобезопасное построение запросов
* Автоматический выбор диалекта по драйверу БД
* Плагинную систему диалектов

## Установка

```shell
go get github.com/Compogo/db-sql-generator
```

## Быстрый старт

```go
package main

import (
    "github.com/Compogo/compogo"
    "github.com/Compogo/db-sql-generator"
	"github.com/doug-martin/goqu"
)

func main() {
    app := compogo.NewApp("myapp",
        compogo.WithComponents(&db_sql_generator.Component),
    )

    app.AddComponents(&compogo.Component{
        Name: "user_repository",
        Init: compogo.StepFunc(func(container compogo.Container) error {
            return container.Invoke(func(gen *goqu.DialectWrapper) error {
                // Генерация SELECT
                sql, args, _ := gen.From("users").Where(goqu.C("age").Gt(18)).ToSQL()
                // SELECT * FROM users WHERE age > 18

                // Генерация INSERT
                sql, args, _ := gen.Insert("users").Rows(goqu.Record{
                    "name": "John",
                    "age":  30,
                }).ToSQL()
                // INSERT INTO users (name, age) VALUES ('John', 30)

                return nil
            })
        }),
    })

    if err := app.Serve(); err != nil {
        panic(err)
    }
}
```

## Регистрация диалектов

```go
import "github.com/Compogo/db-sql-generator"

func init() {
    db_sql_generator.Registration("mysql", "mysql")
}
```

## Зависимости

* [Compogo](https://github.com/Compogo/compogo) — основной фреймворк
* [goqu](https://github.com/doug-martin/goqu) — SQL-генератор

## Лицензия

```plantuml
MIT License

Copyright (c) 2026 Compogo

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

```
