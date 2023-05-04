package modelos

type Mujer struct {
	Hombre  //Hereda sus propiedades de Hombre, tambien podria tener propiedades propias, agregandolas
	EsMadre bool
}

func (m *Mujer) Respirar() { m.respirando = true }

func (m *Mujer) Comer() { m.comiendo = true }

func (m *Mujer) Pensar() { m.pensando = true }

func (m *Mujer) Sexo() string { return "Mujer" }
