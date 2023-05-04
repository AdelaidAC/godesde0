package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Solucion Propia

func TablaDeMultiplicar() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Ingrese un numero:")

	if scanner.Scan() {

		numero, err := strconv.Atoi(scanner.Text())

		if err != nil {
			fmt.Println("Ocurrio un error al convertir: " + err.Error())
		}

		for i := 1; i <= 10; i++ {
			fmt.Printf("%d x %d = %d \n", numero, i, numero*i)
		}
	}
}

// Solucion del Profesor

// func TabladeMultiplicar() {

// 	var numero int
// 	var err error

// 	scanner := bufio.NewScanner(os.Stdin)

// 	for {
// 		fmt.Println("Ingrese un numero: ")
// 		if scanner.Scan() {
// 			numero, err = strconv.Atoi(scanner.Text())
// 			if err != nil {
// 				continue
// 			} else {
// 				break
// 			}
// 		}
// 	}

// 	for i := 1; i <= 10; i++ {
// 		fmt.Printf("%d x %d = %d \n", numero, i, numero*i)
// 	}

// }

func TabladeMultiplicar() string {

	var numero int
	var err error
	var texto string

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Ingrese un numero: ")
		if scanner.Scan() {
			numero, err = strconv.Atoi(scanner.Text())
			if err != nil {
				continue
			} else {
				break
			}
		}
	}

	for i := 1; i <= 10; i++ {
		texto += fmt.Sprintf("%d x %d = %d \n", numero, i, numero*i)
	}

	return texto

}
