package auth

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/samber/do"
	"log/slog"
	"net/http"
)

type Controller struct {
	authService *Service
	logger      *slog.Logger
}

type response struct {
	Result string `json:"result,omitempty"`
}

func NewController(i *do.Injector) (*Controller, error) {
	authService := do.MustInvoke[*Service](i)
	logger := do.MustInvoke[*slog.Logger](i)

	controller := &Controller{authService, logger}
	return controller, nil
}

func (c *Controller) registerRoutes(r *chi.Mux) {
	r.Route("/auth", func(r chi.Router) {
		r.Get("/login", c.login)
	})

	c.logger.Info("registered routes '/auth'")
}

func (c *Controller) login(w http.ResponseWriter, r *http.Request) {
	str, err := c.authService.Login()
	if err != nil {
		return
	}

	render.JSON(w, r, response{Result: str})
}
