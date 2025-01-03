package service

import (
	"context"

	"github.com/cifra-city/comtools/cifractx"
	"github.com/cifra-city/comtools/httpkit"
	"github.com/cifra-city/location-storage/internal/config"
	"github.com/cifra-city/location-storage/internal/service/handlers"
	"github.com/cifra-city/tokens"
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
	_ = service.TokenManager.AuthMiddleware(service.Config.JWT.AccessToken.SecretKey)
	moderatorGrants := service.TokenManager.RoleGrant(service.Config.JWT.AccessToken.SecretKey, tokens.AdminRole, tokens.ModeratorRole)

	r.Route("/location-storage", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Route("/private", func(r chi.Router) {
				r.Use(moderatorGrants)
				r.Route("/create", func(r chi.Router) {
					r.Post("/city", handlers.CreateCity)
					r.Post("/streets", handlers.CreateStreet)
				})
				r.Route("/update", func(r chi.Router) {
					r.Put("/city", handlers.UpdateCity)
					r.Put("/streets", handlers.UpdateStreet)
				})
			})

			r.Route("/public", func(r chi.Router) { //PUBLIC
				r.Route("/get_data/{id}", func(r chi.Router) {
					r.Get("/city", handlers.DataByCity)
					r.Get("/streets", handlers.DataByStreet)
				})
				r.Route("/check", func(r chi.Router) {
					r.Route("/access", func(r chi.Router) {
						r.Get("/{city}/{street}", handlers.CheckCredibilityAddress)
						r.Get("/text/{city}/{street}", handlers.CheckCredibilityAddressByText)
					})
				})
			})
		})
	})

	server := httpkit.StartServer(ctx, service.Config.Server.Port, r, service.Logger)
	<-ctx.Done()
	httpkit.StopServer(context.Background(), server, service.Logger)
}
