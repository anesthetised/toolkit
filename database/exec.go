package database

import "database/sql/driver"

type Executor interface {
	Exec(stmt string, args []driver.Value) (driver.Result, error)
}
