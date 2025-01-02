package handlers

import (
	"net/http"
	"strings"

	"github.com/cifra-city/cifractx"
	"github.com/cifra-city/httpkit"
	"github.com/cifra-city/httpkit/problems"
	"github.com/cifra-city/location-storage/internal/config"
	"github.com/cifra-city/location-storage/internal/service/requests"
	"github.com/google/uuid"
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

	city := req.Data.Attributes.CityId
	newCityName := req.Data.Attributes.NewName
	newCountryId := req.Data.Attributes.CountryId

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

	if newCountryId != nil {
		countryId, err := uuid.Parse(*newCountryId)
		if err != nil {
			log.Debugf("error parsing country id: %v", err)
			httpkit.RenderErr(w, problems.BadRequest(err)...)
			return
		}
		_, err = server.Databaser.Cities.UpdateCountry(r.Context(), cityId, countryId)
		if err != nil {
			log.Errorf("Failed to update city country: %v", err)
			httpkit.RenderErr(w, problems.InternalError())
			return
		}
	}

	httpkit.Render(w, http.StatusOK)
}
