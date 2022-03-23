package controller

import (
	"distancing-detect-backend/internal/controller/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func (c *Controller) GetViolationsOfClass(w http.ResponseWriter, r *http.Request) {
	violations, err := c.service.GetViolations()
	if err != nil {
		c.handleInternalServerError(w, err)
	}
	var encodedViolations []models.ViolationData
	for _, violation := range violations {
		encodedViolations = append(encodedViolations, models.BuildViolationData(violation))
	}
	json.NewEncoder(w).Encode(encodedViolations)
	return
}

func (c *Controller) GetAllViolations(w http.ResponseWriter, r *http.Request) {
	classroom := mux.Vars(r)["class"]
	violations, err := c.service.GetViolationsOfClass(classroom)
	if err != nil {
		c.handleInternalServerError(w, err)
	}
	var encodedViolations []models.ViolationData
	for _, violation := range violations {
		encodedViolations = append(encodedViolations, models.BuildViolationData(violation))
	}
	json.NewEncoder(w).Encode(encodedViolations)
	return
}

func (c *Controller) handleInternalServerError(w http.ResponseWriter, err error) {
	c.logger.ErrorLogger.Println(err.Error())
	if w != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
