package handlers

import (
	"net/http"
	"strings"

	"github.com/cifra-city/comtools/cifractx"
	"github.com/cifra-city/comtools/httpkit"
	"github.com/cifra-city/comtools/httpkit/problems"
	"github.com/cifra-city/location-storage/internal/config"
	"github.com/cifra-city/location-storage/internal/service/requests"
	"github.com/cifra-city/tokens"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func CreateStreet(w http.ResponseWriter, r *http.Request) {
	server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVER)
	if err != nil {
		logrus.Errorf("Failed to retrieve service configuration: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}
	log := server.Logger

	req, err := requests.NewCreateStreet(r)
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
		log.Warn("User is not authorized to create country")
		httpkit.RenderErr(w, problems.Unauthorized("User is not authorized to create country"))
		return
	}

	name := strings.ToLower(req.Data.Attributes.Name)
	district := req.Data.Attributes.DistrictId

	if name == "" {
		log.Warn("City name is required")
		httpkit.RenderErr(w, problems.BadRequest(errors.New("city name is required"))...)
		return
	}

	districtId, err := uuid.Parse(district)
	if err != nil {
		log.Warn("Country ID is invalid")
		httpkit.RenderErr(w, problems.BadRequest(errors.New("country ID is invalid"))...)
		return
	}

	_, err = server.Databaser.Streets.Create(r.Context(), name, districtId)
	if err != nil {
		log.Errorf("Failed to create country: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	httpkit.Render(w, http.StatusCreated)

}
