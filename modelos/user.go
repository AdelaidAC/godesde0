package modelos

import "time"

type User struct {
	Id        int
	Name      string
	CreatedAt time.Time
	Status    bool
}

/*
	*User donde * es un puntero que referencia a la direccion de memoria de User
	func (A que estructura estamos referenciando) = func (this (nombre del parametro) User (nombre del modelo))
*/

func (u *User) AddUser(id int, name string, createdAt time.Time, status bool) {
	u.Id = id
	u.Name = name
	u.CreatedAt = createdAt
	u.Status = status
}
