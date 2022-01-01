package httphandler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func (h *HTTPHandlers) GetAllViolations(w http.ResponseWriter, r *http.Request) {
	retrievedData, err := h.Database.RetrieveData("SELECT photolink,totalviolations,class,timeofdetection FROM violations ORDER BY id")
	if err != nil {
		h.Logger.ErrorLogger.Println("Error reading database: ", err.Error())
	}
	response := h.extractInformation(retrievedData)
	json.NewEncoder(w).Encode(response)
}

func (h *HTTPHandlers) GetViolationsOfClass(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	request := mux.Vars(r)
	class := request["class"]
	retrievedData, err := h.Database.RetrieveData(fmt.Sprintf("Select photolink,totalviolations,class,timeofdetection from violations WHERE class = '%s' ORDER BY id DESC LIMIT 1;", class))
	if err != nil {
		h.Logger.ErrorLogger.Println("Error reading database: ", err.Error())
	}
	response := h.extractInformation(retrievedData)
	json.NewEncoder(w).Encode(response)

}
