package main

import (
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
)

type nodo struct {
	nombre, apellido, apodo, favoritos string
	Siguiente, Anterior                *nodo
}

type lista struct {
	cabeza *nodo
	cola   *nodo
}

func (this *lista) Insertar(nuevo *nodo) {
	if this.cabeza == nil {
		this.cabeza = nuevo
		this.cola = nuevo
	} else {
		this.cola.Siguiente = nuevo
		nuevo.Anterior = this.cola
		this.cola = nuevo
	}
}

func graficarLista(l lista) {
	var contador int
	archivo, _ := os.Create("grafico.dot")
	_, _ = archivo.WriteString("digraph grafico{" + "\n")
	_, _ = archivo.WriteString("compound=true;" + "\n")
	_, _ = archivo.WriteString("subgraph cluster0{" + "\n")
	_, _ = archivo.WriteString("color=orangered" + "\n")
	_, _ = archivo.WriteString("edge[color=\"greenyellow\",minlen=0.1, dir=both]" + "\n")

	_, _ = archivo.WriteString("edge[minlen=1, dir=both]" + "\n")
	aux := l.cabeza
	for aux != nil {
		_, _ = archivo.WriteString("struct" + strconv.Itoa(contador) + "[shape=record,color=\".7 .3 1.0\"label=\"{" + aux.nombre + "|" + aux.apellido + "|" + aux.apodo + "|" + aux.favoritos + "}\"];" + "\n")
		contador++
		aux = aux.Siguiente
	}

	for i := 0; i < contador; i++ {
		if i+1 <= contador-1 {
			_, _ = archivo.WriteString("struct" + strconv.Itoa(i) + "->" + "struct" + strconv.Itoa(i+1) + "\n")
		}
	}

	_, _ = archivo.WriteString("}" + "\n")
	_, _ = archivo.WriteString("}" + "\n")
	archivo.Close()
	path, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(path, "-Tpdf", "./grafico"+".dot").Output()
	mode := 0777
	_ = ioutil.WriteFile("grafica.pdf", cmd, os.FileMode(mode))
}

func main() {
	li := lista{nil, nil}
	a := nodo{"Marvin", "Martinez", "Marvin25ronal", "Jugar apex", nil, nil}
	b := nodo{"Yaiza", "Pineda", "Bambi", "Patinar", nil, nil}
	c := nodo{"Jonathan", "Lopez", "Pancho", "Comer", nil, nil}
	d := nodo{"usuario1", "bla", "bla", "Jugar apex", nil, nil}
	e := nodo{"usuario2", "bla", "bla", "Jugar apex", nil, nil}
	f := nodo{"usuario3", "sale edd", "vamos con todo", "100 en la fase 1", nil, nil}
	li.Insertar(&a)
	li.Insertar(&b)
	li.Insertar(&c)
	li.Insertar(&d)
	li.Insertar(&e)
	li.Insertar(&f)
	graficarLista(li)
}
