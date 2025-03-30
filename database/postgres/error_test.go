package postgres

import (
	"testing"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/stretchr/testify/assert"
)

func TestMatchErrorCode(t *testing.T) {
	pgErr := &pgconn.PgError{Code: ErrorCodeDuplicateKey}
	assert.True(t, IsDuplicateError(pgErr))
	assert.False(t, MatchErrorCode(pgErr, pgerrcode.NoDataFound))
}
