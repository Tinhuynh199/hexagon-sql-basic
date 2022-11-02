package app

import (
	"log"
	"net/http"
)

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
	PATCH = "PATCH"
)

// // Set all required routers
func (a *App) setRouters() {
	// Routing for handling the projects

	userPath := "/users"
	a.Get(userPath+"/search", a.User.Search)
	a.Get(userPath+"/{id}", a.User.Load)
	a.Post(userPath, a.User.Create)
	a.Put(userPath+"/{id}", a.User.Update)
	a.Patch(userPath+"/{id}", a.User.Patch)
	a.Delete(userPath+"/{id}", a.User.Delete)

}

// // Wrap  the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods(GET)
}

// Wrap the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods(POST)
}

// Wrap the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods(PUT)
}

// Wrap the router for PATCH method
func (a *App) Patch(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods(PATCH)
}

// Wrap the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods(DELETE)
}

// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}
