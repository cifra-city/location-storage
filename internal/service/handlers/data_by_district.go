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

func DataByDistrict(w http.ResponseWriter, r *http.Request) {
	server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVER)
	if err != nil {
		logrus.Errorf("Failed to retrieve service configuration: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}
	log := server.Logger

	districtStr := chi.URLParam(r, "id")
	districtId, err := uuid.Parse(districtStr)
	if err != nil {
		httpkit.RenderErr(w, problems.BadRequest(errors.New("district id is required"))...)
		return
	}

	district, err := server.Databaser.Districts.Get(r.Context(), districtId)
	if err != nil {
		log.Errorf("Failed to get district: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	streets, err := server.Databaser.Streets.GetByDistrict(r.Context(), districtId)
	if err != nil {
		log.Errorf("Failed to get districts: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	httpkit.Render(w, NewDataByDistrictResponse(district, streets))
}

func NewDataByDistrictResponse(district dbcore.District, streets []dbcore.Street) resources.DataByDistrict {
	var streetsInners []resources.DataByDistrictDataAttributesStreetsInner
	for _, street := range streets {
		streetsInners = append(streetsInners, resources.DataByDistrictDataAttributesStreetsInner{
			Id:   street.ID.String(),
			Name: street.Name,
		})
	}
	return resources.DataByDistrict{
		Data: resources.DataByDistrictData{
			Id:   district.ID.String(),
			Type: resources.UpdateDistrictType,
			Attributes: resources.DataByDistrictDataAttributes{
				Name:    district.Name,
				Streets: streetsInners,
			},
		},
	}
}
