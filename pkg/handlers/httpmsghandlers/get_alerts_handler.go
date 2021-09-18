package httphandler

import (
	"net/http"

	"github.com/gorilla/mux"
)

func GetAllAlerts(w http.ResponseWriter, r *http.Request) {
	// retrieve all alerts
}

func GetAlertsOfClass(w http.ResponseWriter, r *http.Request) {
	request := mux.Vars(r)
	topicName := request["class"]
	// retrieve alert of class in db
}
