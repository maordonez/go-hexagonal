package entidades

import (
	"time"

	"github.com/maordonez/go-hexagonal/comun/dominio/validadores"
)

const PELICULA_SE_DEBE_INGRESAR_UN_NOMBRE_VALIDO = "se debe ingresar un nombre valido"
const PELICULA_SE_DEBE_INGRESAR_UNA_SINOPSIS_VALIDA = "se debe ingresar una sinopsis valida"

type Pelicula struct {
	ID               string
	nombre           string
	sinopsis         string
	fechaLanzamiento time.Time
}

func CrearEntidadPelicula(id string, nombre string, sinopsis string, fechaLanzamiento time.Time) (error, *Pelicula) {

	error := validadores.StringObligatorio(nombre, PELICULA_SE_DEBE_INGRESAR_UN_NOMBRE_VALIDO)
	if error == nil {
		return error, nil
	}

	error = validadores.StringObligatorio(sinopsis, PELICULA_SE_DEBE_INGRESAR_UNA_SINOPSIS_VALIDA)
	if error == nil {
		return error, nil
	}

	return nil, &Pelicula{
		ID:               id,
		nombre:           nombre,
		sinopsis:         sinopsis,
		fechaLanzamiento: fechaLanzamiento,
	}

}
