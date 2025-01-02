package handlers

import (
	"net/http"
	"strings"

	"github.com/cifra-city/cifractx"
	"github.com/cifra-city/httpkit"
	"github.com/cifra-city/httpkit/problems"
	"github.com/cifra-city/location-storage/internal/config"
	"github.com/cifra-city/location-storage/internal/service/requests"
	"github.com/cifra-city/tokens"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func CreateCity(w http.ResponseWriter, r *http.Request) {
	server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVER)
	if err != nil {
		logrus.Errorf("Failed to retrieve service configuration: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}
	log := server.Logger

	req, err := requests.NewCreateCity(r)
	if err != nil {
		log.Debugf("error decoding request: %v", err)
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	role, ok := r.Context().Value(tokens.RoleKey).(string)
	if !ok {
		log.Warn("UserID not found in context")
		httpkit.RenderErr(w, problems.Unauthorized("User not authenticated"))
		return
	}

	if role != tokens.AdminRole && role != tokens.ModeratorRole {
		log.Warn("User is not authorized to create city")
		httpkit.RenderErr(w, problems.Unauthorized("User is not authorized to create city"))
		return
	}

	name := strings.ToLower(req.Data.Attributes.Name)
	country := req.Data.Attributes.CountryId

	if name == "" {
		log.Warn("City name is required")
		httpkit.RenderErr(w, problems.BadRequest(errors.New("city name is required"))...)
		return
	}

	countryId, err := uuid.Parse(*country)
	if err != nil {
		log.Warn("Country ID is invalid")
		httpkit.RenderErr(w, problems.BadRequest(errors.New("country ID is invalid"))...)
		return
	}

	_, err = server.Databaser.Cities.Create(r.Context(), name, countryId)
	if err != nil {
		log.Errorf("Failed to create city: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	httpkit.Render(w, http.StatusCreated)
}
