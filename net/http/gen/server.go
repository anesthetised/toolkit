package gen

import (
	"log/slog"
	"net/http"
)

type Server struct {
	log *slog.Logger
	enc ServerEncodeDecoder
}

func NewServer(l *slog.Logger, enc ServerEncodeDecoder) *Server {
	return &Server{
		log: l,
		enc: enc,
	}
}

func (s *Server) Log() *slog.Logger {
	return s.log
}

func (s *Server) Encoding() ServerEncodeDecoder {
	return s.enc
}

func (s *Server) Response(rw http.ResponseWriter, code int, v any) error {
	if code == 0 {
		code = http.StatusOK
	}

	return s.enc.Encode(rw, code, v)
}

func (s *Server) Error(rw http.ResponseWriter, code int, err error, logArgs ...any) error {
	if code == 0 {
		code = http.StatusInternalServerError
	}

	text := http.StatusText(code)
	if err != nil {
		text = err.Error()
		s.log.Error(text)
	}

	return s.enc.Encode(rw, code, text)
}
