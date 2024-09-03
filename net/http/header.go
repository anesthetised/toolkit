package http

import (
	"net"
	"net/http"

	"github.com/anesthetised/toolkit/ds"
)

// GetIP returns the IP address of request.
func GetIP(r *http.Request) string {
	headers := []string{"CF-Connecting-IP", "X-Forwarded-For"}

	for _, header := range headers {
		if value := r.Header.Get(header); value != "" {
			return value
		}
	}

	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	return ip
}

func WriteHeaders(w http.ResponseWriter, code int, headers ...ds.Pair[string, string]) {
	for _, h := range headers {
		w.Header().Set(h.Values())
	}

	w.WriteHeader(code)
}
