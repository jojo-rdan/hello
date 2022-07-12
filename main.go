package main

import (
	"fmt"

	"github.com/jojo-rdan/hello/greet"
)

func main() {
	// para exportar identificadores siempre utilizamos
	// como prefijo, el nombre del paquete
	// en este caso fmt es el paquete donde está
	// incluída la función "Println" la cuál tiene
	// la primera letra en mayúscula
	fmt.Println(greet.English())

}
