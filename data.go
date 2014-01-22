package main

import (
	"errors"
	"sync"
	"log"
	"github.com/garyburd/redigo/redis"
	"github.com/garyburd/redigo/redisx"
)

var (
	ErrAlreadyExists = errors.New("route already exists")
)

type DB interface {
	PurgeRoutes()
	GetRoutes() []string
	CreateRoute(route *RouteInfo) bool
	AddBackend(hostname string, backend string) bool
	RemoveBackend(hostname string, backend string) bool
}

// Thread-safe in-memory map of albums.
type routesDB struct {
	sync.RWMutex
	m map[int]*Route
	seq int
	rclient redis.Conn
}


// The one and only database instance.
var db DB

func init() {
	rc, err := redis.Dial("tcp", ":6379")
	if err != nil {
		log.Fatal(err)
	}

	db = &routesDB{
		m: make(map[int]*Route),
		rclient: rc,
	}

	initExample()
}

func initExample() {
	// Fill the database
	db.CreateRoute(&RouteInfo{ Hostname: "foobar.com", Sticky:true })
	db.AddBackend("foobar.com", "http://127.0.0.1:8081")
	db.AddBackend("foobar.com", "http://127.0.0.1:8082")
	db.AddBackend("foobar.com", "127.1.1.55")
	db.RemoveBackend("foobar.com", "127.1.1.55")

	db.CreateRoute(&RouteInfo{ Hostname: "example.com", Sticky:true })
	db.AddBackend("example.com", "http://127.0.0.1:8083")

	db.CreateRoute(&RouteInfo{ Hostname: "*.example.com", Sticky:false })
	db.AddBackend("*.example.com", "127.1.1.45")
}

func (db *routesDB) PurgeRoutes() {
	_, err := db.rclient.Do("DEL", "index:hostnames")
	if err != nil {
		log.Fatal(err)
	}
}

func (db *routesDB) GetRoutes() []string {
	values, err := redis.Values(db.rclient.Do("SMEMBERS", "index:hostnames"))
	if err != nil {
		log.Fatal(err)
	}

	var hostnames []string
	for len(values) > 0 {
		var hostname string
		values, err = redis.Scan(values, &hostname)
		if err != nil {
			panic(err)
		}
		hostnames = append(hostnames, hostname)
	}
	return hostnames
}

// Add creates a new album and returns its id, or an error.
func (db *routesDB) CreateRoute(route *RouteInfo) bool {
	key := "routeinfo:" + route.Hostname
	_, err := db.rclient.Do("HMSET", redisx.AppendStruct([]interface{}{key}, route)...)
	if err != nil {
		log.Fatal(err)
	}

	// Add to index
	_, err = db.rclient.Do("SADD", "index:hostnames", route.Hostname)
	if err != nil {
		log.Fatal(err)
	}
	return true
}

func (db *routesDB) AddBackend(hostname string, backend string) bool {
	_, err := db.rclient.Do("ZADD", "backend:"+hostname, 1, backend)
	if err != nil {
		log.Fatal(err)
	}
	return true
}

func (db *routesDB) RemoveBackend(hostname string, backend string) bool {
	_, err := db.rclient.Do("ZREM", "backend:"+hostname, backend)
	if err != nil {
		log.Fatal(err)
	}
	return true
}

type RouteInfo struct {
	Hostname string 		`json:"hostname"`
	Sticky bool				`json:"sticky"`
	Varnish bool			`json:"varnish"`
	VarnishHandler string	`json:"varnish_handler"`
}

type Route struct {
	Id      int      		`json:"id"`
	Hostname string 		`json:"hostname"`
}
