package gen

import (
	"encoding/json"
	"net/http"
)

type EncodingJSON struct{}

func (e EncodingJSON) Encode(rw http.ResponseWriter, code int, payload any) error {
	rw.WriteHeader(code)

	if err := json.NewEncoder(rw).Encode(payload); err != nil {
		return err
	}

	return nil
}

func (e EncodingJSON) Decode(r *http.Request, payload any) error {
	return json.NewDecoder(r.Body).Decode(payload)
}
