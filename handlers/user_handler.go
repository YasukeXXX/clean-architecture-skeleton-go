package handlers

import (
	"fmt"
	"github.com/YasukeXXX/clean-architecture-skeleton-go/controllers"
	"log"
	"net/http"
)

type UserHandler struct {
	Controller controllers.UserController
}

func (h UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		res, err := h.Controller.Index()
		if err != nil {
			log.Print(err)
			w.WriteHeader(500)
			return
		}
		fmt.Fprintln(w, res)
	case http.MethodPost:
		res, err := h.Controller.Create(r.FormValue("name"), r.FormValue("email"), r.FormValue("password"))
		if err != nil {
			log.Print(err)
			w.WriteHeader(500)
			return
		}
		fmt.Fprintln(w, res)
	default:
		w.WriteHeader(404)
	}
}
