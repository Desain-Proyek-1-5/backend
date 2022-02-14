package httphandler

import (
	"distancing-detect-backend/pkg/database"
	"distancing-detect-backend/pkg/models"
	"log"
)

func (m *HTTPHandlers) extractInformation(retrieved *database.RetrievedData) []models.ViolationData {
	var response []models.ViolationData
	for retrieved.Data.Next() {
		var each models.ViolationData
		err := retrieved.Data.Scan(&each.ImageLink, &each.TotalViolations,
			&each.Class, &each.Time)
		if err != nil {
			log.Println("Error retrieving sql  ", err.Error())
		}
		response = append(response, each)
	}
	return response

}
