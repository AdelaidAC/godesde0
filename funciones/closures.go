package funciones

import "fmt"

/*
	Funcion que recibe como parametro un entero
	Y devuelve una funcion anonima que retorna un entero
*/

func tabla(valor int) func() int {

	numero := valor
	secuencia := 0

	return func() int {
		secuencia++
		return numero * secuencia
	}

}

func LlamarClosure() {

	tabladel := 2
	MiTabla := tabla(tabladel)

	for i := 1; i <= 10; i++ {
		fmt.Println(MiTabla())
	}

}
