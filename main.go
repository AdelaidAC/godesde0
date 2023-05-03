package main

import (
	"fmt"     //Imprimir texto
	"runtime" // Contiene toda la informacion del equipo sobre el cual corre nuestro sistema

	"github.com/AdelaidAC/godesde0/ejercicios"
	"github.com/AdelaidAC/godesde0/teclado"
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

	// EJERCICIO 1 (solucion propia)

	entero, texto, error := ejercicios.EjercicioUno("1000")

	if error != nil {
		fmt.Println("Error durante la conversion")
		return
	} else if entero > 100 {
		fmt.Printf("%s es mayor a 100 \n", texto)
	} else {
		fmt.Printf("%s es menor a 100 \n", texto)
	}

	// EJERCICIO 1 (solucion del maestro)

	num, txt := ejercicios.ConvNumerico("1000")
	fmt.Println(num, txt)

	teclado.IngresoNumeros()
}
