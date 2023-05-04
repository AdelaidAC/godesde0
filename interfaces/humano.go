package interfaces

/*
	Las interfaces son modelos de comportamiento que podemos asociar a nuestras
	estructuras, si nosotros tenemos una estructura de tipo usuario este usuario
	se podria decir que podria implementar una interfaz de tipo humano
*/

/*
	A diferencia de una estructura aqui no podemos colocar variables,
	propiedades, ni tipos de datos, solamente esto va a ser una
	definicion de funciones (que es lo que hace un humano, por ejemplo)
*/

type Humano interface {
	Respirar()
	Pensar()
	Comer()
	Sexo() string
	EstaVivo() bool
}
