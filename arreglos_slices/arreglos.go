package arreglos_slices

import "fmt"

var tabla [10]int = [10]int{10, 42, 52} // Vector o Arrego de tipo entero de 10 elementos
var matriz [20][30]int                  // Matriz de 20 filas y 30 columnas

/*
	La diferencia entre un vector y un slice es que
	el slice no tiene una cantidad definida de elementos (dimension)
	y un vector si
*/

func MuestroArreglos() {

	tabla[7] = 33
	tabla[2] = 54
	fmt.Println(tabla)

	tabla2 := [10]string{"Pablo", "Juan"}
	fmt.Println(tabla2)

	for i := 0; i < len(tabla); i++ {
		fmt.Println(tabla[i])
	}

	matriz[7][24] = 15

	fmt.Println(matriz)

}
