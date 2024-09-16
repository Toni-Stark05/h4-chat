package http_server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/samber/do"
	"h4-chat/internal/middleware/logger"
	"log/slog"
)

func NewRouter(i *do.Injector) (*chi.Mux, error) {
	log := do.MustInvoke[*slog.Logger](i)
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(logger.New(log))
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)

	return r, nil
}
