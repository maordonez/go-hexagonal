package excepciones

import (
	"fmt"
)

const DATA_INVALIDA = "datos invalidos"

type ErrorCustom struct {
	codigo      int
	categoria   string
	mensaje     string
	operacional bool
}

func (e *ErrorCustom) Error() string {
	return fmt.Sprintf("%t %s %s", e.operacional, e.categoria, e.mensaje)
}

func DataInvalida(mensaje string) error {
	return &ErrorCustom{
		mensaje:     mensaje,
		categoria:   DATA_INVALIDA,
		codigo:      400,
		operacional: false,
	}
}
