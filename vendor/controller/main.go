package controller

import (
	"net/http"
	"gopkg.in/mgo.v2"
	"golang.org/x/tools/go/gcimporter15/testdata"
)

type (
	MainController struct {
		session *mgo.Session
	}
)

func NewMainController(s *mgo.Session) *MainController {
	return &MainController{s}
}

func (m *MainController) Main(w http.ResponseWriter, req *http.Request) {
	// Controller logic here...

	// Get Table
	err := m.session.DB("go-mongo-db").C("table")
	if err {
		w.WriteHeader(404)
		return
	}


	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello World!"))
}