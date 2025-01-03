package handlers

import (
	"net/http"
	"strings"

	"github.com/cifra-city/comtools/cifractx"
	"github.com/cifra-city/comtools/httpkit"
	"github.com/cifra-city/comtools/httpkit/problems"
	"github.com/cifra-city/location-storage/internal/config"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

func CheckCredibilityAddress(w http.ResponseWriter, r *http.Request) {
	server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVER)
	if err != nil {
		logrus.Errorf("Failed to retrieve service configuration: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}
	log := server.Logger

	cityUrl := strings.ToLower(chi.URLParam(r, "city"))
	streetUrl := strings.ToLower(chi.URLParam(r, "street"))

	street, err := server.Databaser.Streets.GetByName(r.Context(), streetUrl)
	if err != nil {
		log.Errorf("Failed to get street: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	city, err := server.Databaser.Cities.GetByName(r.Context(), cityUrl)
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
