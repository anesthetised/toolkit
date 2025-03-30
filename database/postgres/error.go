package postgres

import (
	"errors"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
)

const (
	ErrorCodeDuplicateKey = pgerrcode.UniqueViolation
)

func MatchErrorCode(err error, code string) bool {
	var pgErr *pgconn.PgError

	if !errors.As(err, &pgErr) {
		return false
	}

	return pgErr.Code == code
}

func IsDuplicateError(err error) bool { return MatchErrorCode(err, ErrorCodeDuplicateKey) }
