package ejercicios

import "strconv"

// Solucion Propia
func EjercicioUno(texto string) (int, string, error) {

	entero, err := strconv.Atoi(texto)

	return entero, texto, err
}

// Solucion del Maestro
func ConvNumerico(texto string) (int, string) {

	// _ Indica que el argumento que nos esten devolviendo no lo vamos a asignar a ninguna variable (no me interesa el error)
	//num, _ := strconv.Atoi(texto)

	num, err := strconv.Atoi(texto)

	if err != nil {
		return 0, "Hubo en error " + err.Error() // Imprime el error en string
	}
	if num > 100 {
		return num, "Es mayor a 100"
	} else {
		return num, "Es menor a 100"
	}
}
