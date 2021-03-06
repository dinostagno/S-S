package main

import (
	"log"
	"net/http"
	"text/template"
)

var plantillas = template.Must(template.ParseGlob("plantillas/*"))

func main() {
	http.HandleFunc("/", Inicio)
	log.Println("servidor corriendo. . .")
	http.ListenAndServe(":8080", nil)
}
func Inicio(w http.ResponseWriter, r *http.Request) {
	plantillas.ExecuteTemplate(w, "inicio", nil)
}
