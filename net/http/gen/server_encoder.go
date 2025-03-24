package gen

import (
	"net/http"
)

type ServerEncodeDecoder interface {
	Encode(rw http.ResponseWriter, code int, payload any) error
	Decode(r *http.Request, payload any) error
}
