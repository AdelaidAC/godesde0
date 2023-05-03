package ejercicios

import "strconv"

func EjercicioUno(texto string) (int, string, error) {

	entero, err := strconv.Atoi(texto)

	return entero, texto, err
}
