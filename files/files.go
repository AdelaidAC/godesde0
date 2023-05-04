package files

import (
	"bufio"
	"fmt"
	"os"

	"github.com/AdelaidAC/godesde0/ejercicios"
)

var fileName string = "./files/txt/tabla.txt"

func GrabaTabla() {

	var texto string = ejercicios.TabladeMultiplicar()

	archivo, err := os.Create(fileName)

	if err != nil {
		fmt.Println("Error al crear el archivo " + err.Error())
		return
	}

	fmt.Fprintln(archivo, texto)

	archivo.Close()

}

func SumaTabla() {

	var texto string = ejercicios.TabladeMultiplicar()

	if !Append(fileName, texto) {
		fmt.Println("Error al concatenar contenido")
	}

}

func Append(fileName string, texto string) bool {

	archivo, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		fmt.Println("Error durante el Append " + err.Error())
		return false
	}

	_, err = archivo.WriteString(texto)

	if err != nil {
		fmt.Println("Error durante el WriteString " + err.Error())
		return false
	}

	archivo.Close()
	return true
}

// func LeoArchivo() { DEPRECADO

// 	archivo, err := ioutil.ReadFile(fileName)

// 	if err != nil {
// 		fmt.Println("Error al leer archivo " + err.Error())
// 		return
// 	}

// 	fmt.Println(string(archivo))
// }

func LeoArchivo() {

	archivo, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Error al leer archivo " + err.Error())
		return
	}

	scanner := bufio.NewScanner(archivo)

	for scanner.Scan() { // Mientras existan lineas en el archivo itera
		registro := scanner.Text() // Devuelve cada linea del archivo
		fmt.Println("> " + registro)
	}

	archivo.Close()

}
