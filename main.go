package main

import (
	"fmt"     //Imprimir texto
	"runtime" // Contiene toda la informacion del equipo sobre el cual corre nuestro sistema

	"github.com/AdelaidAC/godesde0/variables"
)

func main() {

	//variables.MuestroEnteros()
	//variables.RestoVariables()

	estado, texto := variables.ConviertoATexto(1000) // Se capturan ambos parametros de retorno de la func ConviertoATexto
	fmt.Println(estado)
	fmt.Println(texto)

	// Asignacion ; Pregunta y Condicion para Evaluarla
	if os := runtime.GOOS; os == "Linux." || os == "OS X." {
		fmt.Println("Esto no es Windows, es ", os)
	} else {
		fmt.Println("Esto es Windows")
	}

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Esto es Linux")
	case "darwin":
		fmt.Println("Esto es Darwnin")
	default: // Cualquier otro SO
		fmt.Printf("%s \n", os)
	}
}
