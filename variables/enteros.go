package variables

import "fmt"

func MuestroEnteros() {
	var intcomun int     // Declarado variable de forma declarativa
	intde32 := int32(10) // Declarando variable de manera de asignacion
	intde64 := int64(100)

	fmt.Println("intcomun = ", intcomun)
	fmt.Println("intde32 = ", intde32)
	fmt.Println("intde64 = ", intde64)
}

// A diferencia de NodeJS no hay necesidad de exportar las funciones o variables
// Ya que con el simple hecho de que el nombre contenga una mayuscula ya lo exportara en automatico
// Como "MuestroEnteros"
