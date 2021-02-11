package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Datos struct {
	Vector []Mensaje `json:"Mensajes"`
}

type Mensaje struct {
	Origen  string        `json:"Origen"`
	Destino string        `json:"Destino"`
	Ms      []DatoMensaje `json:"Msg"`
}

type DatoMensaje struct {
	Fecha string `json:"Fecha"`
	Texto string `json:"Texto"`
}

func iniciando(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Funcionando en el puerto 3000")
}

var ms Datos

func agregarMensaje(w http.ResponseWriter, r *http.Request) {

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error al insertar")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.Unmarshal(reqBody, &ms)
	json.NewEncoder(w).Encode(ms)
}

func getMensajes(w http.ResponseWriter, r *http.Request) {
	for i := 0; i < len(ms.Vector); i++ {
		fmt.Println("Indice: " + ms.Vector[i].Destino)
		fmt.Println("Indice: " + ms.Vector[i].Origen)
		for j := 0; j < len(ms.Vector[i].Ms); j++ {
			fmt.Println("Departamento: " + ms.Vector[i].Ms[j].Fecha)
			fmt.Println("Departamento: " + ms.Vector[i].Ms[j].Texto)
		}
	}
	json.NewEncoder(w).Encode(ms)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", iniciando).Methods("GET")
	router.HandleFunc("/agregarMSG", agregarMensaje).Methods("POST")
	router.HandleFunc("/Mensajes", getMensajes).Methods("GET")
	log.Fatal(http.ListenAndServe(":3000", router))
}
