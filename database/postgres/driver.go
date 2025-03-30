package postgres

// import (
// 	"context"
// 	"database/sql"
// 	"database/sql/driver"

// 	pgxstd "github.com/jackc/pgx/v5/stdlib"
// )

// type postgresDriver struct {
// 	*pgxstd.Driver
// }

// func (postgresDriver) Open(dsn string) (driver.Conn, error) {
// 	db, err := sql.Open("pgx", dsn)
// 	if err != nil {
// 		return nil, err
// 	}

// 	conn, err := db.Conn(context.Background())
// 	if err != nil {
// 		return nil, err
// 	}

// 	err = conn.Raw(func(driverConn any) error {
// 		conn := driverConn.(*pgxstd.Conn).Conn()

// 		return nil
// 	})

// 	return nil, err
// }
