package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"

	thttp "github.com/anesthetised/toolkit/net/http"
	"github.com/anesthetised/toolkit/net/http/gen"
	"golang.org/x/sync/errgroup"
)

func run(ctx context.Context, logger *slog.Logger) error {
	s := gen.NewServer(logger, gen.EncodingJSON{})

	r := http.NewServeMux()
	r.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		var test struct {
			String  string `json:"string"`
			Number  int    `json:"number"`
			Boolean bool   `json:"boolean"`
		}

		test.String = "test"
		test.Number = 123
		test.Boolean = true

		_ = s.Response(w, http.StatusTeapot, &test)
	})

	hs := &http.Server{
		Handler: r,
		Addr:    ":8080",
	}

	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error { return thttp.ListenAndServe(ctx, hs) })

	logger.Info("example server is running", "addr", hs.Addr)

	return g.Wait()
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	err := run(ctx, logger)
	if errors.Is(err, context.Canceled) {
		logger.Info("canceled")
		return
	}
	if err != nil {
		logger.Error("operation is failed", "err", err)
		return
	}
}
