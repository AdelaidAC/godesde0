package users

import (
	"fmt"
	"time"

	"github.com/AdelaidAC/godesde0/modelos"
)

/*
En POO las definiciones no son objetos hasta que uno los crea con NEW
*/
func AltaUsuario() {

	u := new(modelos.User) // Aqui estamos creando un objeto de tipo modelo User
	u.AddUser(10, "Pablo", time.Now(), true)
	fmt.Println(u) // &{10 Pablo 2023-05-04 14:19:54.9308174 -0700 PDT m=+0.008814801 true} | el & indica que es una estructura

}
