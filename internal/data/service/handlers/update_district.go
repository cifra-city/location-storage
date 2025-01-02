package handlers

import (
	"net/http"

	"github.com/cifra-city/cifractx"
	"github.com/cifra-city/httpkit"
	"github.com/cifra-city/httpkit/problems"
	"github.com/cifra-city/location-storage/internal/config"
	"github.com/cifra-city/location-storage/internal/data/service/requests"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func UpdateDistrict(w http.ResponseWriter, r *http.Request) {
	server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVER)
	if err != nil {
		logrus.Errorf("Failed to retrieve service configuration: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}
	log := server.Logger

	req, err := requests.NewUpdateDistrict(r)
	if err != nil {
		log.Debugf("error decoding request: %v", err)
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	newName := req.Data.Attributes.NewName
	district := req.Data.Attributes.DistrictId
	newCity := req.Data.Attributes.CityId

	districtId, err := uuid.Parse(district)
	if err != nil {
		log.Debugf("error parsing district id: %v", err)
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	if newName != nil {
		_, err = server.Databaser.Districts.UpdateName(r.Context(), districtId, *newName)
		if err != nil {
			log.Errorf("Failed to update district name: %v", err)
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
		_, err = server.Databaser.Districts.UpdateCity(r.Context(), districtId, newCityId)
		if err != nil {
			log.Errorf("Failed to update district city: %v", err)
			httpkit.RenderErr(w, problems.InternalError())
			return
		}
	}

	httpkit.Render(w, http.StatusOK)
}
