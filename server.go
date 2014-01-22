package main

import (
	"net/http"
	"log"
	"github.com/codegangsta/martini"
)


// The one and only martini instance.
var m *martini.Martini

func init() {
	m := martini.Classic()

	// Setup routes
	m.Get(`/routes`, GetRoutes)
	m.Post(`/routes`, CreateRoute)
	m.Delete(`/routes`, PurgeRoutes)

	m.Post("/routes/:hostname/backends", AddBackend)
	m.Delete("/routes/:hostname/backends", RemoveBackend)
//	r.Get("/routes/:hostname", func(params martini.Params) string {
//			panic("not implemented")
//	})
//
//	r.Post("/routes/:hostname/backends", func(params martini.Params) string {
//			panic("not implemented")
//	})
//
//	r.Get("/routes/:hostname/sslcert", func(params martini.Params) string {
//			panic("not implemented")
//	})
//
//	r.Post("/routes/:hostname/privatekey", func(params martini.Params) string {
//			panic("not implemented")
//	})

	// Inject database
	m.MapTo(db, (*DB)(nil))

	// Add the router action
	m.Run()
}


func main() {
	print ("starting server")
	if err := http.ListenAndServe(":3000", m); err != nil {
		log.Fatal(err)
	}
}
