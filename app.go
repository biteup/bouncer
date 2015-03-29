package main

import (
	"flag"
	"fmt"
	"net/http"
	"log"
	"os"
	controllers "github.com/gobbl/bouncer/controllers"
	"github.com/gorilla/mux"
)

var port int

func init() {
	const (
		defaultPort = 8080
		portUsage   = "port number"
	)

	port := os.Getenv("PORT")
    if port == "" {
		flag.IntVar(&port, "p", defaultPort, portUsage)
    }
}

func setupServer() *mux.Router {
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/", controllers.GetHomeHandler).Methods("GET").Schemes("https")
	r.HandleFunc("/status", controllers.GetStatusHandler).Methods("GET")

	signup := r.Path("/signup").Subrouter()
	signup.Methods("POST").HandlerFunc(signHandler)
	/*
		signup := r.Path("/signup").Subrouter()
		signup.Methods("POST").HandlerFunc(signHandler)

		// Posts collection
		signup := r.Path("/signup").Subrouter()
		signup.Methods("GET").HandlerFunc(PostsIndexHandler)
		signup.Methods("POST").HandlerFunc(PostsCreateHandler)

		// Posts singular
		login := r.PathPrefix("/login/{id}").Subrouter()
		login.Methods("GET").Path("/edit").HandlerFunc(PostEditHandler).Schemes("http")
		login.Methods("GET").HandlerFunc(PostShowHandler)
		login.Methods("PUT", "POST").HandlerFunc(PostUpdateHandler)
		login.Methods("DELETE").HandlerFunc(PostDeleteHandler)
	*/
	return r

}

func main() {

	flag.Parse()

	r := setupServer()

	fmt.Printf("Starting server on port:%d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}

func signHandler(rw http.ResponseWriter, r *http.Request) {
	controllers.SignupHandler(rw)
}
