package http

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/anesthetised/toolkit/ds"
	"github.com/anesthetised/toolkit/validate"
)

func EncodeJSON(w http.ResponseWriter, code int, v any) error {
	WriteHeaders(w, code, ds.NewPair("Content-Type", "application/json; charset=utf-8"))
	return json.NewEncoder(w).Encode(v)
}

func DecodeJSON[T any](r io.Reader) (T, error) {
	var (
		zero = new(T)
		val  = *zero

		err error
	)

	if err = json.NewDecoder(r).Decode(&val); err != nil {
		return val, err
	}

	if v, ok := any(val).(validate.Validator); ok {
		if err = v.Validate(); err != nil {
			return val, fmt.Errorf("validate: %w", err)
		}
	}

	return val, err
}

func Response(w http.ResponseWriter, code int, v any) error {
	if code == 0 {
		code = http.StatusOK
	}

	return EncodeJSON(w, code, v)
}

func Error(w http.ResponseWriter, code int, err error, l *slog.Logger, a ...any) error {
	if code == 0 {
		code = http.StatusInternalServerError
	}

	text := http.StatusText(code)
	if err != nil {
		text = err.Error()
	}

	if l != nil {
		l.Error(text, a...)
	}

	type respErr struct {
		Error string `json:"error"`
	}

	return EncodeJSON(w, code, respErr{Error: text})
}
