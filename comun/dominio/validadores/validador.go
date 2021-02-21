package validadores

import (
	"github.com/maordonez/go-hexagonal/comun/dominio/excepciones"
)

func StringObligatorio(cadena string, mensaje string) error {
	if len(cadena) == 0 {
		return excepciones.DataInvalida(mensaje)
	}

	return nil
}
