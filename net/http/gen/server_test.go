package gen

import (
	"encoding/json"
	"io"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServer(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	assert.NotNil(t, NewServer[json.RawMessage](logger, EncodingJSON{}))
}
