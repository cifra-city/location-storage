package handlers

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/cifra-city/comtools/cifractx"
	"github.com/cifra-city/comtools/httpkit"
	"github.com/cifra-city/comtools/httpkit/problems"
	"github.com/cifra-city/location-storage/internal/config"
	"github.com/cifra-city/location-storage/internal/service/requests"
	"github.com/cifra-city/tokens"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

var geoRegex = regexp.MustCompile(`^(POINT\((-?\d+(\.\d+)?\s-?\d+(\.\d+)?)\)|LINESTRING\((-?\d+(\.\d+)?\s-?\d+(\.\d+)?,?\s?)+\)|POLYGON\(\((-?\d+(\.\d+)?\s-?\d+(\.\d+)?,?\s?)+\)\))$`)

func validateGeoString(geo string) error {
	if !geoRegex.MatchString(geo) {
		return errors.New("invalid geometry format")
	}
	return nil
}

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
	location := req.Data.Attributes.Location

	err = validateGeoString(location)
	if err != nil {
		log.Warnf("Invalid location: %v", err)
		httpkit.RenderErr(w, problems.BadRequest(errors.New("invalid location"))...)
		return
	}

	if name == "" {
		log.Warn("City name is required")
		httpkit.RenderErr(w, problems.BadRequest(errors.New("city name is required"))...)
		return
	}

	_, err = server.Databaser.Cities.Create(r.Context(), name, location)
	if err != nil {
		log.Errorf("Failed to create city: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	httpkit.Render(w, http.StatusCreated)
}
