package handlers

import (
	"net/http"
	"strings"

	"github.com/cifra-city/comtools/cifractx"
	"github.com/cifra-city/comtools/httpkit"
	"github.com/cifra-city/comtools/httpkit/problems"
	"github.com/cifra-city/location-storage/internal/config"
	"github.com/cifra-city/location-storage/internal/service/requests"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func UpdateStreet(w http.ResponseWriter, r *http.Request) {
	server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVER)
	if err != nil {
		logrus.Errorf("Failed to retrieve service configuration: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}
	log := server.Logger

	req, err := requests.NewUpdateStreet(r)
	if err != nil {
		log.Debugf("error decoding request: %v", err)
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	street := req.Data.Id
	newName := req.Data.Attributes.NewName
	newLocation := req.Data.Attributes.NewLocation
	newCity := req.Data.Attributes.NewCity

	streetId, err := uuid.Parse(street)
	if err != nil {
		log.Debugf("error parsing street id: %v", err)
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	if newName != nil {
		_, err = server.Databaser.Streets.UpdateName(r.Context(), streetId, strings.ToLower(*newName))
		if err != nil {
			log.Errorf("Failed to update street name: %v", err)
			httpkit.RenderErr(w, problems.InternalError())
			return
		}
	}

	if newLocation != nil {
		err = validateGeoString(*newLocation)
		if err != nil {
			log.Warnf("Invalid location: %v", err)
			httpkit.RenderErr(w, problems.BadRequest(errors.New("invalid location"))...)
			return
		}

		_, err = server.Databaser.Streets.UpdateLocation(r.Context(), streetId, *newLocation)
		if err != nil {
			log.Errorf("Failed to update street district: %v", err)
			httpkit.RenderErr(w, problems.InternalError())
			return
		}
	}

	if newCity != nil {
		newCityId, err := uuid.Parse(*newCity)
		if err != nil {
			log.Debugf("error parsing city id: %v", err)
			httpkit.RenderErr(w, problems.BadRequest(err)...)
			return
		}
		_, err = server.Databaser.Streets.UpdateCity(r.Context(), streetId, newCityId)
		if err != nil {
			log.Errorf("Failed to update street city: %v", err)
			httpkit.RenderErr(w, problems.InternalError())
			return
		}
	}

	httpkit.Render(w, http.StatusOK)
}
