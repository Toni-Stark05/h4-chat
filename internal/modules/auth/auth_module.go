package auth

import (
	"github.com/go-chi/chi/v5"
	"github.com/samber/do"
	"log"
)

func NewAuthModule(i *do.Injector) {
	do.Provide(i, NewController)
	do.Provide(i, NewService)

	controller, err := do.Invoke[*Controller](i)
	if err != nil {
		log.Fatal("Init AuthModule ERROR:", err)
	}

	router := do.MustInvoke[*chi.Mux](i)
	controller.registerRoutes(router)
}
