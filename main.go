package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/mickelsonm/goauthsecure/controllers/middleware"
)

var (
	listenAddr = flag.Int("port", 8080, "http listen port")
)

func main() {
	flag.Parse()

	r := mux.NewRouter()

	//notice that this route is wide open
	r.HandleFunc("/hamster", HamsterHandler).Methods("GET")

	//this route is setup to require authentication
	r.HandleFunc("/hamsterdance", middleware.Route(HamsterDanceHandler, &middleware.Config{RequireAuth: true})).Methods("GET")

	//pretty standard out of the box endpoints
	r.HandleFunc("/healthstatus", HealthCheckHandler)
	r.HandleFunc("/", RootHandler)

	n := negroni.New(negroni.NewRecovery())
	n.Use(negroni.HandlerFunc(middleware.Middleware))
	n.UseHandler(r)
	n.Run(fmt.Sprintf(":%d", *listenAddr))
}

func RootHandler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Dude, it's an API."))
}

func HealthCheckHandler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK: It's all good."))
}

func HamsterHandler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hamster Info!"))
}

func HamsterDanceHandler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hamster Dance"))
}
