package httphandler

import (
	"capstone/pkg/database"
	"capstone/pkg/models"
	"log"
)

func (m *HTTPHandlers) extractInformation(retrieved *database.RetrievedData) []models.ViolationData {
	var (
		time            string
		totalViolations int
		class           string
	)
	var response []models.ViolationData
	for retrieved.Data.Next() {
		var each models.ViolationData
		err := retrieved.Data.Scan(&time, &totalViolations, &class)
		if err != nil {
			log.Println("Error retrieving sql  ", err.Error())
		}
		each.Time = time
		each.TotalViolations = totalViolations
		each.Class = class
		response = append(response, each)
	}
	return response

}
