package main

import (
	"log"
	"net/http"

	"github.com/husobee/vestigo"
)

func StartServer(port string) {
	// Router
	router := vestigo.NewRouter()
	vestigo.AllowTrace = true //Trace is on

	// Handlers
	router.HandleFunc("/", Index)

	router.Get("/graphql", GraphQLGETHandler)
	router.Post("/graphql", GraphQLPOSTHandler)

	// Serve
	log.Fatal(http.ListenAndServe(port, router))
}
