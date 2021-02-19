package dominio

// Implementador es una estructura que usamos para mostrar
// cómo se implementa la interfaz fmt.Stringer
type Implementador struct {
	A string
	B string
}

// String imprime los campos concatenados
func (e *Implementador) String() string {
	return e.A + e.B
}

// StringAlReves imprime los campos al revés
func (e *Implementador) StringAlReves() string {
	return e.B + e.A
}
