package iteraciones

import "fmt"

func Iterar() {
	for i := 0; i < 100; i += 5 {
		if i == 6 {
			break //salir del ciclo for ( se aborta )
			// continue : muestra todos los numeros menos el 6 (se lo salta)
		}
		fmt.Println(i)
	}
}
