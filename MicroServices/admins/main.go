package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/matscus/Hamster/MicroServices/admins/handlers"
	"github.com/matscus/Hamster/Package/Middleware/middleware"
)

var (
	pemPath    string
	keyPath    string
	proto      string
	listenport string
)

func main() {
	flag.StringVar(&pemPath, "pem", os.Getenv("SERVERREM"), "path to pem file")
	flag.StringVar(&keyPath, "key", os.Getenv("SERVERKEY"), "path to key file")
	flag.StringVar(&listenport, "port", ":10005", "port to Listen")
	flag.StringVar(&proto, "proto", "https", "http or https")
	flag.Parse()
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/admins/getallusers", middleware.Middleware(handlers.GetAllUsers)).Methods("GET", "OPTIONS")
	http.Handle("/api/v1/admins/", r)
	log.Println("ListenAndServe: " + listenport)
	err := http.ListenAndServeTLS(listenport, pemPath, keyPath, context.ClearHandler(http.DefaultServeMux))
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}