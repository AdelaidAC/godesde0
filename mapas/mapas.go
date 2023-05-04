package mapas

import "fmt"

func MostrarMapas() {

	//Mapa[Clave]Valor - Clave tipo String y Valor tipo String

	paises := make(map[string]string) //Creando Mapa con Make

	paises["Mexico"] = "D.F."
	paises["Argentina"] = "Buenos Aires"

	fmt.Println(paises)
	fmt.Println(paises["Argentina"])

	campeonato := map[string]int{ //Creando Mapa de forma Tradicional
		"Barcelona":    39,
		"Real Madrid":  38,
		"Chivas":       37,
		"Boca Juniors": 30,
	}

	fmt.Println(campeonato)

	// for LLAVE, VALOR := range campeonato
	// range se usa para recorrer toda la coleccion del mapa y siempre acompanado del for, es una especie de foreach
	for equipo, puntaje := range campeonato {
		fmt.Printf("%s : %d \n", equipo, puntaje)
	}

	// Delete esta especialmente hecho para los mapas
	// (Nombre del Mapa, Llave)
	delete(campeonato, "Real Madrid")
	fmt.Println(campeonato)

	// Cuando consultamos un mapa se nos devuelven dos valores
	// El valor que tiene la clave y un booleano que indica si la clave existe o no
	puntaje, existe := campeonato["Juventus"]
	fmt.Printf("El puntaje capturado es %d y el equipo existe = %t \n", puntaje, existe)

}
