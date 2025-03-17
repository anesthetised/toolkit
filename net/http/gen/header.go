package gen

import (
	"net"
	"net/http"

	"github.com/anesthetised/toolkit/ds"
)

// GetIP returns the IP address of request.
// Example: GetIP(r, "CF-Connecting-IP", "X-Forwarded-For")
func GetIP(r *http.Request, headers ...string) string {
	for _, header := range headers {
		if value := r.Header.Get(header); value != "" {
			return value
		}
	}

	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	return ip
}

// WriteHeaders writes pairs of headers to a given http.ResponseWriter, finishing with a status code.
func WriteHeaders(w http.ResponseWriter, code int, headers ...ds.Pair[string, string]) {
	for _, h := range headers {
		w.Header().Set(h.Values())
	}

	w.WriteHeader(code)
}
