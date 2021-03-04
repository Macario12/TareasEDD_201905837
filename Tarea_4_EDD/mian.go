package main

import (
	"fmt"
)

func main() {
	user0 := Usuarios{
		ID:     8,
		Nombre: "user0"}

	user1 := Usuarios{
		ID:     78,
		Nombre: "user1"}

	user2 := Usuarios{
		ID:     1,
		Nombre: "user2"}

	user3 := Usuarios{
		ID:     45,
		Nombre: "user3"}

	user4 := Usuarios{
		ID:     100,
		Nombre: "user04"}

	arbolito := New_Arbol()
	Insertar_Arbol(arbolito, user0)
	Insertar_Arbol(arbolito, user1)
	Insertar_Arbol(arbolito, user2)
	Insertar_Arbol(arbolito, user3)
	Insertar_Arbol(arbolito, user4)
	fmt.Println("Recorrido en PreOrden")
	preOrden(&arbolito.Raiz)

	fmt.Println("Recorrido en InOrden")
	inOrden(&arbolito.Raiz)

	fmt.Println("Recorrido en PostOrden")
	postOrden(&arbolito.Raiz)
}

type Usuarios struct {
	ID     int
	Nombre string
}

type Nodo struct {
	User      Usuarios
	Izquierda *Nodo
	Derecha   *Nodo
}

func New_Nodo(user Usuarios) *Nodo {
	return &Nodo{
		user,
		nil,
		nil}
}

type Arbol struct {
	Raiz *Nodo
}

func New_Arbol() *Arbol {
	return &Arbol{nil}
}

func Insertar_Arbol(arbol *Arbol, user Usuarios) {
	Insertar_Nodo(&arbol.Raiz, user)
}

func Insertar_Nodo(nodo **Nodo, user Usuarios) bool {

	if *nodo == nil {
		*nodo = New_Nodo(user)
		return true
	} else if user.ID < (*nodo).User.ID {
		return Insertar_Nodo(&(*nodo).Izquierda, user)
	} else if user.ID > (*nodo).User.ID {
		return Insertar_Nodo(&(*nodo).Derecha, user)
	}
	return false
}

func preOrden(nodr **Nodo) {
	if *nodr != nil {
		fmt.Println((*nodr).User.Nombre)
		preOrden(&(*nodr).Izquierda)
		preOrden(&(*nodr).Derecha)

	}

}

func postOrden(nodr **Nodo) {
	if *nodr != nil {
		preOrden(&(*nodr).Izquierda)
		preOrden(&(*nodr).Derecha)

		fmt.Println((*nodr).User.Nombre)

	}

}

func inOrden(nodr **Nodo) {
	if *nodr != nil {
		preOrden(&(*nodr).Izquierda)

		fmt.Println((*nodr).User.Nombre)
		preOrden(&(*nodr).Derecha)

	}

}
