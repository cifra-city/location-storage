package handlers

import (
	"net/http"

	"github.com/cifra-city/cifractx"
	"github.com/cifra-city/httpkit"
	"github.com/cifra-city/httpkit/problems"
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
	districtUrl := chi.URLParam(r, "district")
	streetUrl := chi.URLParam(r, "street")

	cityId, err := uuid.Parse(cityUrl)
	if err != nil {
		httpkit.RenderErr(w, problems.BadRequest(errors.New("city id is required"))...)
		return
	}
	districtId, err := uuid.Parse(districtUrl)
	if err != nil {
		httpkit.RenderErr(w, problems.BadRequest(errors.New("district id is required"))...)
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

	if street.DistrictID != districtId {
		httpkit.RenderErr(w, problems.Conflict("street does not belong to the district"))
		return
	}

	district, err := server.Databaser.Districts.Get(r.Context(), districtId)
	if err != nil {
		log.Errorf("Failed to get district: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	if district.CityID != cityId {
		httpkit.RenderErr(w, problems.Conflict("district does not belong to the city"))
		return
	}

	httpkit.Render(w, http.StatusFound)
}
