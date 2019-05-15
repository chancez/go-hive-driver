package hive

import (
	"database/sql"
	"database/sql/driver"
	"errors"
)

const defaultBatchSize = int64(1000)

var (
	ErrNoPassword = errors.New("hive: passwd is required")

	_ driver.Driver = &Driver{}
)

func init() {
	sql.Register("hive", &Driver{})
}

type Driver struct{}

func (*Driver) Open(dsn string) (driver.Conn, error) {
	return Open(dsn)
}
