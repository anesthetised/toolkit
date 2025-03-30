package database

import "database/sql/driver"

type Exec func(stmt string, args []driver.Value) (driver.Result, error)
