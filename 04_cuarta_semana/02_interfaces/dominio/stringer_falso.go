package dominio

// StringerFalso es una interfaz que creamos
// para mostrar cómo podemos estar implementando
// más de una interfaz a la vez sólo con implementar
// los métodos necesarios
type StringerFalso interface {
	String() string
	StringAlReves() string
}
