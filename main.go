package main

import (
	"fmt"

	"github.com/AdelaidAC/godesde0/variables"
)

func main() {
	//variables.MuestroEnteros()
	//variables.RestoVariables()

	estado, texto := variables.ConviertoATexto(1000) // Se capturan ambos parametros de retorno de la func ConviertoATexto
	fmt.Println(estado)
	fmt.Println(texto)
}
