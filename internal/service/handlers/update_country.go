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

func UpdateCountry(w http.ResponseWriter, r *http.Request) {
	server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVER)
	if err != nil {
		logrus.Errorf("Failed to retrieve service configuration: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}
	log := server.Logger

	req, err := requests.NewUpdateCountry(r)
	if err != nil {
		log.Debugf("error decoding request: %v", err)
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	newCountryName := strings.ToLower(req.Data.Attributes.NewName)
	newCountryId := req.Data.Attributes.CountryId

	countryId, err := uuid.Parse(newCountryId)
	if err != nil {
		log.Debugf("error parsing city id: %v", err)
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	_, err = server.Databaser.Countries.Update(r.Context(), countryId, newCountryName)
	if err != nil {
		log.Errorf("Failed to update country: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	httpkit.Render(w, http.StatusOK)
}
