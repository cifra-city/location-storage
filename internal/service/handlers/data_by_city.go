package handlers

import (
	"net/http"

	"github.com/cifra-city/comtools/cifractx"
	"github.com/cifra-city/comtools/httpkit"
	"github.com/cifra-city/comtools/httpkit/problems"
	"github.com/cifra-city/location-storage/internal/config"
	"github.com/cifra-city/location-storage/internal/data/db/dbcore"
	"github.com/cifra-city/location-storage/resources"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func DataByCity(w http.ResponseWriter, r *http.Request) {
	server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVER)
	if err != nil {
		logrus.Errorf("Failed to retrieve service configuration: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}
	log := server.Logger

	cityStr := chi.URLParam(r, "id")
	cityId, err := uuid.Parse(cityStr)
	if err != nil {
		httpkit.RenderErr(w, problems.BadRequest(errors.New("city id is required"))...)
		return
	}

	city, err := server.Databaser.Cities.Get(r.Context(), cityId)
	if err != nil {
		log.Errorf("Failed to get city: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	streets, err := server.Databaser.Streets.ListStreetsByCity(r.Context(), cityId)
	if err != nil {
		log.Errorf("Failed to get districts: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	httpkit.Render(w, NewDataByCityResponse(city, streets))
}

func NewDataByCityResponse(city dbcore.City, districts []dbcore.Street) resources.DataByCity {
	var streetsInners []resources.DataByCityDataAttributesStreetsInner
	for _, district := range districts {
		streetsInners = append(streetsInners, resources.DataByCityDataAttributesStreetsInner{
			Id:   district.ID.String(),
			Name: district.Name,
		})
	}
	return resources.DataByCity{
		Data: resources.DataByCityData{
			Id:   city.ID.String(),
			Type: resources.DataByCityType,
			Attributes: resources.DataByCityDataAttributes{
				Name:    city.Name,
				Streets: streetsInners,
			},
		},
	}
}
