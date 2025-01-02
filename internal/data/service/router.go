package service

import (
	"context"

	"github.com/cifra-city/cifractx"
	"github.com/cifra-city/httpkit"
	"github.com/cifra-city/location-storage/internal/config"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

func Run(ctx context.Context) {
	r := chi.NewRouter()

	service, err := cifractx.GetValue[*config.Service](ctx, config.SERVER)
	if err != nil {
		logrus.Fatalf("failed to get server from context: %v", err)
	}

	r.Use(cifractx.MiddlewareWithContext(config.SERVER, service))
	authMW := service.TokenManager.Middleware(service.Config.JWT.AccessToken.SecretKey)

	r.Route("/location-storage", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Route("/private", func(r chi.Router) {
				r.Use(authMW)
				r.Route("/create", func(r chi.Router) {
					r.Post("/country", nil)
					r.Post("/city", nil)
					r.Post("/districts", nil)
					r.Post("/streets", nil)
				})
				r.Route("/update", func(r chi.Router) {
					r.Put("/country", nil)
					r.Put("/city", nil)
					r.Put("/districts", nil)
					r.Put("/streets", nil)
				})
			})
			r.Route("/public", func(r chi.Router) {
				r.Route("/get", func(r chi.Router) {
				})
			})
		})
	})

	server := httpkit.StartServer(ctx, service.Config.Server.Port, r, service.Logger)
	<-ctx.Done()
	httpkit.StopServer(context.Background(), server, service.Logger)
}
