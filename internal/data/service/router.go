package service

import (
	"context"

	"github.com/cifra-city/cifractx"
	"github.com/cifra-city/httpkit"
	"github.com/cifra-city/location-storage/internal/config"
	"github.com/cifra-city/location-storage/internal/data/service/handlers"
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
					r.Post("/country", handlers.CreateCountry)
					r.Post("/city", handlers.CreateCity)
					r.Post("/districts", handlers.CreateDistrict)
					r.Post("/streets", handlers.CreateStreet)
				})
				r.Route("/update", func(r chi.Router) {
					r.Put("/country", handlers.UpdateCountry)
					r.Put("/city", handlers.UpdateCity)
					r.Put("/districts", handlers.UpdateDistrict)
					r.Put("/streets", handlers.UpdateStreet)
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
