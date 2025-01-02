package handlers

import (
	"net/http"

	"github.com/cifra-city/cifractx"
	"github.com/cifra-city/httpkit"
	"github.com/cifra-city/httpkit/problems"
	"github.com/cifra-city/location-storage/internal/config"
	"github.com/cifra-city/location-storage/internal/service/requests"
	"github.com/google/uuid"
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

	newName := req.Data.Attributes.NewName
	street := req.Data.Attributes.StreetId
	newDistrict := req.Data.Attributes.DistrictId

	streetId, err := uuid.Parse(street)
	if err != nil {
		log.Debugf("error parsing street id: %v", err)
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	if newName != nil {
		_, err = server.Databaser.Streets.UpdateName(r.Context(), streetId, *newName)
		if err != nil {
			log.Errorf("Failed to update street name: %v", err)
			httpkit.RenderErr(w, problems.InternalError())
			return
		}
	}

	if newDistrict != nil {
		newDistrictId, err := uuid.Parse(*newDistrict)
		if err != nil {
			log.Debugf("error parsing district id: %v", err)
			httpkit.RenderErr(w, problems.BadRequest(err)...)
			return
		}
		_, err = server.Databaser.Streets.UpdateDistrict(r.Context(), streetId, newDistrictId)
		if err != nil {
			log.Errorf("Failed to update street district: %v", err)
			httpkit.RenderErr(w, problems.InternalError())
			return
		}
	}

	httpkit.Render(w, http.StatusOK)
}
