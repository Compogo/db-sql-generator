# Compogo DB SQL Generator 🛠️

**Compogo DB SQL Generator** — компонент для типобезопасного построения SQL-запросов, построенный на базе [goqu](https://github.com/doug-martin/goqu). Автоматически использует диалект, соответствующий драйверу, выбранному в `db-client`.

## 🚀 Установка

```bash
go get github.com/Compogo/db-sql-generator
```

### 📦 Быстрый старт

```go
package main

import (
    "github.com/Compogo/compogo"
    "github.com/Compogo/db-client"
    "github.com/Compogo/db-sql-generator"
    _ "github.com/Compogo/postgres" // ваш драйвер БД
)

func main() {
    app := compogo.NewApp("myapp",
        compogo.WithOsSignalCloser(),
        db_client.Component,           // выбираем драйвер через --db.driver
        db_sql_generator.Component,    // генератор запросов
        compogo.WithComponents(
            userRepositoryComponent,
        ),
    )

    if err := app.Serve(); err != nil {
        panic(err)
    }
}

// Репозиторий, использующий генератор
var userRepositoryComponent = &component.Component{
    Dependencies: component.Components{
        db_client.Component,
        db_sql_generator.Component,
    },
    Execute: component.StepFunc(func(c container.Container) error {
        return c.Invoke(func(db db_client.Client, gen *goqu.DialectWrapper) {
            repo := &UserRepository{db: db, generator: gen}
            // ... регистрация репозитория
        })
    }),
}

type UserRepository struct {
    db        db_client.Client
    generator *goqu.DialectWrapper
}

func (r *UserRepository) GetUsers(ctx context.Context) ([]User, error) {
    // Строим типобезопасный запрос
    query := r.generator.From("users").
        Where(goqu.C("active").IsTrue()).
        Order(goqu.C("created_at").Desc()).
        Limit(20)

    sql, _, _ := query.ToSQL()
    rows, err := r.db.QueryContext(ctx, sql)
    // ...
}
```

### ✨ Возможности

#### 🔌 Плагинная архитектура диалектов

Драйверы БД сами регистрируют свой диалект для goqu:

```go
// В драйвере postgres
func init() {
    db_sql_generator.Registration(Postgres, "postgres")
}

// В драйвере mysql
func init() {
    db_sql_generator.Registration(MySQL, "mysql")
}
```
