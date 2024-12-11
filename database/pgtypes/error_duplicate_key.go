package pgtypes

import (
	"errors"
	"strconv"

	"github.com/lib/pq"
)

const (
	ErrorCodeDuplicateKey = 23505
)

func MatchErrorCode(err error, code int) bool {
	var pqErr *pq.Error

	if !errors.As(err, &pqErr) {
		return false
	}

	return pqErr.Code == pq.ErrorCode(strconv.Itoa(code))
}

func IsDuplicateError(err error) bool { return MatchErrorCode(err, ErrorCodeDuplicateKey) }
