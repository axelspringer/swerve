package api

import (
	"fmt"
	"net/http"

	"github.com/axelspringer/swerve/log"
	"github.com/gorilla/mux"
)

func sendJSON(w http.ResponseWriter, data []byte, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write([]byte(fmt.Sprintf("{\"data\":%s}", string(data))))
}

func sendJSONMessage(w http.ResponseWriter, msg string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write([]byte(fmt.Sprintf("{\"code\":%d,\"message\":\"%s\"}", code, msg)))
}

func sendTextMessage(w http.ResponseWriter, msg string, code int) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(code)
	w.Write([]byte(fmt.Sprintf(msg)))
}

func walkRoutes(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
	tpl, err := route.GetPathTemplate()
	if err != nil {
		log.Warn(err.Error())
	}
	met, err := route.GetMethods()
	if err != nil {
		log.Warn(err.Error())
	}
	log.Infof("API route: %+v - %+v", tpl, met)
	return nil
}
