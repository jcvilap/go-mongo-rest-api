package route

import (
	"net/http"
	"controller"
	"gopkg.in/mgo.v2"
)

func Route(s *mgo.Session) {
	// Controllers
	mainController := controller.NewMainController(s)

	// App Routes
	http.HandleFunc("/", mainController.Main)
}