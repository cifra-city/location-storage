package handlers

import (
	"net/http"

	"github.com/cifra-city/comtools/cifractx"
	"github.com/cifra-city/comtools/httpkit"
	"github.com/cifra-city/comtools/httpkit/problems"
	"github.com/cifra-city/location-storage/internal/config"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func CheckCredibilityAddressByText(w http.ResponseWriter, r *http.Request) {
	server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVER)
	if err != nil {
		logrus.Errorf("Failed to retrieve service configuration: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}
	log := server.Logger

	cityUrl := chi.URLParam(r, "city")
	streetUrl := chi.URLParam(r, "street")

	cityId, err := uuid.Parse(cityUrl)
	if err != nil {
		httpkit.RenderErr(w, problems.BadRequest(errors.New("city id is required"))...)
		return
	}
	streetId, err := uuid.Parse(streetUrl)
	if err != nil {
		httpkit.RenderErr(w, problems.BadRequest(errors.New("street id is required"))...)
		return
	}

	street, err := server.Databaser.Streets.Get(r.Context(), streetId)
	if err != nil {
		log.Errorf("Failed to get street: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	city, err := server.Databaser.Cities.Get(r.Context(), cityId)
	if err != nil {
		log.Errorf("Failed to get city: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	if street.CityID != city.ID {
		httpkit.RenderErr(w, problems.Conflict())
		return
	}

	httpkit.Render(w, http.StatusFound)
}
