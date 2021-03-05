// Cuando se trata de paquetes de tests, se puede utilizar un
// "_test" después del nombre del paquete con el que esté trabajando
// en ese directorio o se puede usar el mismo paquete.
// El nombre del archivo sí tiene que tener un sufijo "_test".
package porcentaje_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	porcentaje "github.com/magdagc/CursoGOFebrero2021/01_primera_semana/06_importar_paquetes_externos"
)

// TestPorcentaje usa el paquete que viene para testing en golang.
func TestPorcentaje(t *testing.T) {
	var porcentaje float32 = porcentaje.Porcentaje(50, 10)

	if porcentaje != 5 {
		t.Error("El 10% de 50 debería ser 5")
	}
}

// TestPorcentajeConTestify usa el paquete assert de testify para probar,
// esta dependencia la manejamos con go modules.
func TestPorcentajeConTestify(t *testing.T) {
	var porcentaje float32 = porcentaje.Porcentaje(50, 10)
	assert.Equal(t, float32(5), porcentaje, "El 10% de 50 debería ser 5")
}
