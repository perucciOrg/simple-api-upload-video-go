package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go_2/pkg/routes"
)

func main(){
	r := mux.NewRouter()
	routes.RegisterVideoStore(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}


