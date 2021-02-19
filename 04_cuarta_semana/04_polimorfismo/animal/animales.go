package animal

// Perro representa un perro
type Perro struct{}

// Gato representa un gato
type Gato struct{}

// EmitirSonido devuelve el sonido que hacen los perros
func (p *Perro) EmitirSonido() string {
	return "guau"
}

// EmitirSonido devuelve el sonido que hacen los gatos
func (g *Gato) EmitirSonido() string {
	return "miau"
}
