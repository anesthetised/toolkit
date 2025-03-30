package sqlite

import (
	"database/sql"
	"database/sql/driver"
	"fmt"

	tdatabase "github.com/anesthetised/toolkit/database"
	"modernc.org/sqlite"
)

type sqliteDriver struct {
	*sqlite.Driver
}

func (d sqliteDriver) Open(dsn string) (driver.Conn, error) {
	conn, err := d.Driver.Open(dsn)
	if err != nil {
		return conn, err
	}

	c := conn.(tdatabase.Executor)

	if _, err = c.Exec("PRAGMA foreign_keys = on;", nil); err != nil {
		_ = conn.Close()
		return nil, fmt.Errorf("enable foreign keys: %w", err)
	}

	return conn, nil
}

func init() {
	sql.Register("sqlite3", sqliteDriver{Driver: &sqlite.Driver{}})
}
