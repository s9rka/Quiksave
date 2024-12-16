package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"notas/routes"
)

func main() {
	fmt.Println("Helloo there!")
	r := mux.NewRouter()
	r.HandleFunc("/", routes.HomeHandler)
	
}