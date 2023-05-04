package modelos

type Hombre struct { // Es como la definicion de una clase con sus propiedades
	edad       int
	altura     float32
	peso       float32
	respirando bool
	pensando   bool
	comiendo   bool
	vivo       bool
}

func (h *Hombre) Respirar() { // Es como la definicion de los metodos de una clase
	h.respirando = true
}

func (h *Hombre) Comer() {
	h.comiendo = true
}

func (h *Hombre) Pensar() {
	h.pensando = true
}

func (h *Hombre) Sexo() string {
	return "Hombre"
}
