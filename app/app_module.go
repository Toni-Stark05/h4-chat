package app

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/samber/do"
	"h4-chat/internal/config"
	"h4-chat/internal/http-server"
	"h4-chat/internal/logger"
	"h4-chat/internal/modules/auth"
	"log"
	"log/slog"
	"net/http"
)

var modules = []ModulesFunc{
	auth.NewAuthModule,
}

type App struct {
	injector *do.Injector
}

func NewApp() *App {
	return &App{injector: do.New()}
}

type ModulesFunc func(i *do.Injector)

func (a App) Run() {
	do.Provide(a.injector, logger.NewLogger)
	do.Provide(a.injector, config.NewConfig)
	do.Provide(a.injector, http_server.NewRouter)

	a.registrationModule()

	a.buildApp()
}

func (a App) registrationModule() {
	for _, module := range modules {
		module(a.injector)
	}
}

func (a App) buildApp() {
	r, err := do.Invoke[*chi.Mux](a.injector)
	if err != nil {
		fmt.Println(err)
		return
	}
	logger, err := do.Invoke[*slog.Logger](a.injector)
	if err != nil {
		log.Fatal("Build App error:", err)
	}

	logger.Info("Server running on http://localhost:8080")

	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("ListenAndServe error", err)
	}
}
