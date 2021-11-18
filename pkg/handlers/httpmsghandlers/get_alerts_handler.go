package httphandler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func (h *HTTPHandlers) GetAllViolations(w http.ResponseWriter, r *http.Request) {
	retrievedData, err := h.Database.RetrieveData("SELECT * FROM violations")
	if err != nil {
		h.Logger.ErrorLogger.Println("Error reading database: ", err.Error())
	}
	response := h.extractInformation(retrievedData)
	json.NewEncoder(w).Encode(response)
}

func (h *HTTPHandlers) GetViolationsOfClass(w http.ResponseWriter, r *http.Request) {
	request := mux.Vars(r)
	class := request["class"]
	retrievedData, err := h.Database.RetrieveData(fmt.Sprintf("SELECT * FROM violations WHERE class = '%s'", class))
	if err != nil {
		h.Logger.ErrorLogger.Println("Error reading database: ", err.Error())
	}
	response := h.extractInformation(retrievedData)
	json.NewEncoder(w).Encode(response)

}
