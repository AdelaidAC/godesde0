package arreglos_slices

import "fmt"

var tablaS []int = []int{2, 5, 4}
var arreglo [10]int = [10]int{6, 98, 21, 676, 18, 36, 78, 9}

func MuestroSlice() {

	fmt.Println(tablaS)

	porcion1 := arreglo[3:]  // Slice creado a partir de un vector desde la posicion 3 hasta el final
	porcion2 := arreglo[:5]  // Slice desde la posicion 0 hasta la 5
	porcion3 := arreglo[6:7] // Slice desde la posicion 6 hasta la 7

	fmt.Println(porcion1)
	fmt.Println(porcion2)
	fmt.Println(porcion3)

}

// make es para indicarle al compilador la capacidad de memoria que se le asignara a un slice
func Capacidad() {

	//Slice de tipo int de 5 elementos y capacidad de 20 elementos
	elementos := make([]int, 5, 20)
	fmt.Printf("Largo %d, Capacidad %d", len(elementos), cap(elementos))

	nums := make([]int, 0, 0)

	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}

	fmt.Printf("Largo %d, Capacidad %d", len(nums), cap(nums))

}
