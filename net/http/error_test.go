package http

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnexpectedStatusCode(t *testing.T) {
	code := http.StatusInternalServerError
	err := UnexpectedStatusCode(code)
	assert.True(t, StatusCodeIs(err, code))
}
