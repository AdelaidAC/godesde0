package variables

import (
	"fmt"
	"strconv"
	"time"
)

// var nombre string - la variable es vista por todas las funciones del archivo resto.go
// var Nombre string - la variable es vista por todas las funciones pertenecientes al paquete variables y por los archivos que importen el paquete

var Nombre string
var Estado bool
var Sueldo float32
var Fecha time.Time

func RestoVariables() { // funcion publica - puede ser vista fuera del paquete

	Nombre = "Pedro"
	Estado = true
	Sueldo = 1577.66
	Fecha = time.Now()

	fmt.Println(Nombre)
	fmt.Println(Estado)
	fmt.Println(Sueldo)
	fmt.Println(Fecha)
}

// Primer grupo de parentesis: parametros de entrada
// Segundo grupo de parentesis: parametros de salida (respuesta)
func ConviertoATexto(numero int) (bool, string) {
	texto := strconv.Itoa(numero)
	return true, texto
}
