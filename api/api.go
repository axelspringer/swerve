package api

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/axelspringer/swerve/helper"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/axelspringer/swerve/log"
	"github.com/gorilla/mux"
)

var githubHash string

// NewAPIServer creates a new instance
func NewAPIServer(mod ModelAdapter, conf Config, wrapHandler func(string, http.Handler) http.Handler) *API {
	api := &API{
		Model:  mod,
		Config: conf,
	}

	router := mux.NewRouter()
	router.Use(helper.LoggingMiddleware)

	router.HandleFunc("/metrics", prometheusHandler).
		Methods(http.MethodGet)
	router.HandleFunc("/health", api.health).
		Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/version", api.version).
		Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/login", api.login).
		Methods(http.MethodPost, http.MethodOptions)

	// TODO conf.Version can be "" this cause to an error, remove or make it mandetory
	auth := router.PathPrefix("/" + conf.Version + "/redirects").Subrouter()
	auth.HandleFunc("/export", api.exportRedirects).
		Methods(http.MethodGet, http.MethodOptions)
	auth.HandleFunc("/import", api.importRedirects).
		Methods(http.MethodPost, http.MethodOptions)
	auth.HandleFunc("/{"+pathParamName+"}", api.getRedirect).
		Methods(http.MethodGet, http.MethodOptions)
	auth.HandleFunc("/", api.createRedirect).
		Methods(http.MethodPost, http.MethodOptions)
	auth.HandleFunc("/", api.getRedirectsPaginated).
		Methods(http.MethodGet)
	auth.HandleFunc("/{"+pathParamName+"}", api.deleteRedirect).
		Methods(http.MethodDelete, http.MethodOptions)
	auth.HandleFunc("/{"+pathParamName+"}", api.updateRedirect).
		Methods(http.MethodPut, http.MethodOptions)
	router.Use(api.corsMiddlewear)
	auth.Use(api.corsMiddlewear)
	auth.Use(api.authMiddlewear)
	router.Walk(walkRoutes)

	addr := ":" + strconv.Itoa(api.Config.Listener)
	api.server = &http.Server{
		Addr:    addr,
		Handler: router,
	}

	return api
}

// Listen to api
func (api *API) Listen() error {
	log.Infof("API listening to %d", api.Config.Listener)
	return api.server.ListenAndServe()
}

func (api *API) health(w http.ResponseWriter, r *http.Request) {
	sendJSONMessage(r, w, "OK", http.StatusOK)
}

func (api *API) version(w http.ResponseWriter, r *http.Request) {
	versionSuffix := ""
	if githubHash != "" {
		versionSuffix = fmt.Sprintf("-%s", githubHash)
	}
	sendJSONMessage(r, w, api.Config.Version+versionSuffix, http.StatusOK)
}

func (api *API) exportRedirects(w http.ResponseWriter, r *http.Request) {
	data, err := api.Model.ExportRedirectsAsJSON()
	if err != nil {
		log.Debugf(ErrRedirectsExport+": %s", err.Error())
		sendJSONMessage(r, w, ErrRedirectsExport, http.StatusInternalServerError)
	}
	modtime := time.Now()
	name := "redirects" + modtime.Format("2006-01-02") + ".json"
	reader := bytes.NewReader(data)
	w.Header().Set("Content-Disposition", "attachment; filename=\""+name+"\"")
	w.Header().Set("Content-Type", "application/json")
	http.ServeContent(w, r, name, modtime, reader)
}

func (api *API) getRedirectsPaginated(w http.ResponseWriter, r *http.Request) {
	queryVars := r.URL.Query()
	cursorParams, ok := queryVars[queryParamCursor]
	cursor := ""
	if ok && len(cursorParams) != 0 {
		cursor = cursorParams[0]
		return
	}
	redirects, newCursor, err := api.Model.GetRedirectsPaginatedAsJSON(cursor)
	if err != nil {
		log.Debugf(ErrRedirectsFetch+": %s", err.Error())
		sendJSONMessage(r, w, ErrRedirectsFetch, http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("{\"data\":%s,\"cursor\":\"%s\"}", string(redirects), newCursor)))
}

func (api *API) importRedirects(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		sendJSONMessage(r, w, "Please provide a file", http.StatusBadRequest)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		sendJSONMessage(r, w, "File could not be read", http.StatusBadRequest)
		return
	}
	err = api.Model.ImportRedirectsFromJSON(data)
	if err != nil {
		log.Debugf(ErrRedirectsImport+": %s", err.Error())
		sendJSONMessage(r, w, ErrRedirectsImport, http.StatusInternalServerError)
		return
	}
	sendJSONMessage(r, w, "Success", http.StatusOK)
}

func (api *API) getRedirect(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	domain, ok := vars[pathParamName]
	if !ok {
		sendJSONMessage(r, w, "Please provide a domain name", http.StatusBadRequest)
		return
	}
	redirect, err := api.Model.GetRedirectByDomainAsJSON(domain)
	if err != nil {
		log.Debugf(ErrRedirectNotFound+": %s", err.Error())
		sendJSONMessage(r, w, ErrRedirectNotFound, http.StatusNotFound)
		return
	}
	sendJSON(r, w, redirect, http.StatusOK)
}

func (api *API) createRedirect(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		sendJSONMessage(r, w, "Body is invalid", http.StatusBadRequest)
		return
	}
	err = api.Model.CreateRedirectFromJSON(data)
	if err != nil {
		log.Debugf(ErrRedirectCreate+": %s", err.Error())
		sendJSONMessage(r, w, ErrRedirectCreate, http.StatusInternalServerError)
		return
	}
	sendJSONMessage(r, w, "Success", http.StatusOK)
}

func (api *API) deleteRedirect(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	domain, ok := vars[pathParamName]
	if !ok {
		sendJSONMessage(r, w, "Please provide a domain name", http.StatusBadRequest)
		return
	}
	err := api.Model.DeleteRedirectByDomain(domain)
	if err != nil {
		log.Debugf(ErrRedirectDelete+": %s", err.Error())
		sendJSONMessage(r, w, ErrRedirectDelete, http.StatusInternalServerError)
		return
	}
	sendJSONMessage(r, w, "Success", http.StatusOK)
}

func (api *API) updateRedirect(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	domain, ok := vars[pathParamName]
	if !ok {
		sendJSONMessage(r, w, "Please provide a domain name", http.StatusBadRequest)
		return
	}
	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		sendJSONMessage(r, w, "Body is invalid", http.StatusBadRequest)
		return
	}
	err = api.Model.UpdateRedirectByDomainWithJSON(domain, data)
	if err != nil {
		log.Debugf(ErrRedirectUpdate+": %s", err.Error())
		sendJSONMessage(r, w, ErrRedirectUpdate, http.StatusInternalServerError)
		return
	}
	sendJSONMessage(r, w, "Success", http.StatusOK)
}

func (api *API) login(w http.ResponseWriter, r *http.Request) {
	log.Info("here")
	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		sendJSONMessage(r, w, "Body is invalid", http.StatusBadRequest)
		return
	}
	tokenString, expirationTime, err := api.Model.CheckPasswordFromJSON(data, api.Config.Secret)
	if err != nil {
		log.Error(err)
		sendJSONMessage(r, w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	log.Debug("set cookie")

	http.SetCookie(w, &http.Cookie{
		Name:    cookieName,
		Value:   tokenString,
		Expires: time.Unix(expirationTime, 0),
	})

	sendJSONMessage(r, w, "Success", http.StatusOK)
}

func prometheusHandler(w http.ResponseWriter, r *http.Request) {
	h := promhttp.Handler()
	h.ServeHTTP(w, r)
}
