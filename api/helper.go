package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/axelspringer/swerve/log"
	"github.com/gorilla/mux"
)

var (
	uiDomain = strings.TrimSpace(os.Getenv("API_UI_URL"))
)

func sendJSON(r *http.Request, w http.ResponseWriter, obj interface{}, code int) {
	jsonBytes := []byte{}
	_, ok := obj.(string)
	if ok {
		jsonBytes, _ = json.Marshal(struct {
			Data interface{} `json:"data"`
		}{
			Data: obj,
		})
	} else {
		jsonBytes = obj.([]byte)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", uiDomain)
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.WriteHeader(code)
	w.Write(jsonBytes)
}

func sendJSONMessage(r *http.Request, w http.ResponseWriter, msg string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", uiDomain)
	w.Header().Set("Access-Control-Allow-Credentials", "true")
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
