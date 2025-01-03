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

func UpdateCity(w http.ResponseWriter, r *http.Request) {
	server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVER)
	if err != nil {
		logrus.Errorf("Failed to retrieve service configuration: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}
	log := server.Logger

	req, err := requests.NewUpdateCity(r)
	if err != nil {
		log.Debugf("error decoding request: %v", err)
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	city := req.Data.Id
	newCityName := req.Data.Attributes.NewName
	newLocation := req.Data.Attributes.NewLocation

	cityId, err := uuid.Parse(city)
	if err != nil {
		log.Debugf("error parsing city id: %v", err)
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	if newCityName != nil {
		_, err = server.Databaser.Cities.UpdateName(r.Context(), cityId, strings.ToLower(*newCityName))
		if err != nil {
			log.Errorf("Failed to update city name: %v", err)
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

		_, err = server.Databaser.Cities.UpdateLocation(r.Context(), cityId, *newLocation)
		if err != nil {
			log.Errorf("Failed to update city country: %v", err)
			httpkit.RenderErr(w, problems.InternalError())
			return
		}
	}

	httpkit.Render(w, http.StatusOK)
}
