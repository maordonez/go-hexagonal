package transformador

import (
	"github.com/maordonez/go-hexagonal/aplicacion/comando"
	"github.com/maordonez/go-hexagonal/dominio/entidades"
)

type TranformadorPelicula struct {
}

func (t *TranformadorPelicula) CrearPelicula(comando *comando.ComandoPelicula) (error, *entidades.Pelicula) {
	return nil, nil
}
