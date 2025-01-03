package handlers

import (
	"net/http"

	"github.com/cifra-city/comtools/cifractx"
	"github.com/cifra-city/comtools/httpkit"
	"github.com/cifra-city/comtools/httpkit/problems"
	"github.com/cifra-city/location-storage/internal/config"
	"github.com/cifra-city/location-storage/resources"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func DataByStreet(w http.ResponseWriter, r *http.Request) {
	server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVER)
	if err != nil {
		logrus.Errorf("Failed to retrieve service configuration: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}
	log := server.Logger

	streetStr := chi.URLParam(r, "id")
	streetId, err := uuid.Parse(streetStr)
	if err != nil {
		httpkit.RenderErr(w, problems.BadRequest(errors.New("country id is required"))...)
		return
	}

	street, err := server.Databaser.Streets.Get(r.Context(), streetId)
	if err != nil {
		log.Errorf("Failed to get street: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	district, err := server.Databaser.Districts.Get(r.Context(), street.DistrictID)
	if err != nil {
		log.Errorf("Failed to get district: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	httpkit.Render(w, resources.DataByStreet{
		Data: resources.DataByStreetData{
			Id:   street.ID.String(),
			Type: resources.DataByStreetType,
			Attributes: resources.DataByStreetDataAttributes{
				Name:     street.Name,
				District: district.ID.String(),
				City:     district.CityID.String(),
			},
		},
	})
}
