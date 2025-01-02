package handlers

import (
	"net/http"

	"github.com/cifra-city/cifractx"
	"github.com/cifra-city/httpkit"
	"github.com/cifra-city/httpkit/problems"
	"github.com/cifra-city/location-storage/internal/config"
	"github.com/cifra-city/location-storage/internal/data/db/dbcore"
	"github.com/cifra-city/location-storage/resources"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func DataByCountry(w http.ResponseWriter, r *http.Request) {
	server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVER)
	if err != nil {
		logrus.Errorf("Failed to retrieve service configuration: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}
	log := server.Logger

	countryStr := chi.URLParam(r, "id")
	countryId, err := uuid.Parse(countryStr)
	if err != nil {
		httpkit.RenderErr(w, problems.BadRequest(errors.New("country id is required"))...)
		return
	}

	country, err := server.Databaser.Countries.Get(r.Context(), countryId)
	if err != nil {
		log.Errorf("Failed to get country: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	cities, err := server.Databaser.Cities.GetByCountry(r.Context(), countryId)
	if err != nil {
		log.Errorf("Failed to get cities: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	httpkit.Render(w, NewDataByCountryResponse(country, cities))
}

func NewDataByCountryResponse(country dbcore.Country, cities []dbcore.City) resources.DataByCountry {
	var citiesInners []resources.DataByCountryDataAttributesCitiesInner
	for _, city := range cities {
		citiesInners = append(citiesInners, resources.DataByCountryDataAttributesCitiesInner{
			Id:   city.ID.String(),
			Name: city.Name,
		})
	}
	return resources.DataByCountry{
		Data: resources.DataByCountryData{
			Id:   country.ID.String(),
			Type: resources.DataByCountryType,
			Attributes: resources.DataByCountryDataAttributes{
				Name:   country.Name,
				Cities: citiesInners,
			},
		},
	}
}
