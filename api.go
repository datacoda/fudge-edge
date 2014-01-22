package main

import (
	"net/http"
	"encoding/json"
	"github.com/codegangsta/martini"
)



func GetRoutes(db DB) string {
	data := db.GetRoutes()
	b, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func PurgeRoutes(db DB) string {
	db.PurgeRoutes()
	return "All purged"
}

func CreateRoute(r *http.Request, db DB) string {
	hostname := r.FormValue("hostname")
	db.CreateRoute(&RouteInfo{ Hostname:hostname, Sticky:false })
	return "Route created " + r.FormValue("hostname")
}

func AddBackend(params martini.Params, r *http.Request, db DB) string {
	hostname := params["hostname"]
	backend := r.FormValue("backend")

	db.AddBackend(hostname, backend)
	return "Backend " + backend + " added for " + hostname
}

func RemoveBackend(params martini.Params, r *http.Request, db DB) string {
	hostname := params["hostname"]
	backend := r.FormValue("backend")

	db.RemoveBackend(hostname, backend)
	return "Backend " + backend + " removed for " + hostname
}
