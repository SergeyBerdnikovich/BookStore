package main

import (
	"net/http"
	"time"

	"github.com/bmizerany/pat"
	"github.com/gtforge/services_common_go/gett-config"
	"github.com/gtforge/services_common_go/gett-ops"
	"github.com/urfave/negroni"
)

// newServer - setup http server with negroni()
func newServer() *http.Server {
	n := negroni.New()
	if gettConfig.Settings.Environment.GetBool("log") {
		n.Use(negroni.NewLogger())
	}
	// Middlewares
	for _, middleware := range gettOps.GetAllMiddlewares() {
		n.UseFunc(middleware)
	}
	// Add the router action
	n.UseHandler(setupRouter())

	Server = &http.Server{
		Addr:           ":" + gettConfig.Settings.Environment.GetString("server.port"),
		Handler:        n,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return Server
}

func setupRouter() *pat.PatternServeMux {
	router := pat.New()
	//TODO ADD Route here
	for path, handler := range gettOps.GetAllRouter() {
		router.Get(path, handler)
	}
	return router
}
