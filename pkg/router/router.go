package router

import (
	"net/http"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	mux "github.com/gorilla/mux"
)

type RouterInstance struct {
	Router *mux.Router
}

type RequestHandler interface {
	Handle(client mqtt.Client, msg mqtt.Message)
}

func NewRouterInstance() *RouterInstance {
	return &RouterInstance{mux.NewRouter().StrictSlash(true)}
}

func (a *RouterInstance) RegisterHandler(Path string, Handler func(w http.ResponseWriter, r *http.Request), method string) {
	a.Router.HandleFunc(Path, Handler).Methods(method)
}

func (a *RouterInstance) Start() {
	http.ListenAndServe(":8080", a.Router)
}
