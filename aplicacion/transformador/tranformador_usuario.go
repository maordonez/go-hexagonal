package transformador

import (
	"github.com/maordonez/go-hexagonal/aplicacion/comando"
	"github.com/maordonez/go-hexagonal/dominio/entidades"
)

type TransformadorUsuario struct {
}

func (t *TransformadorUsuario) CrearUsuario(comando *comando.ComandoUsuario) (error, *entidades.Usuario) {
	return entidades.CrearEntidadUsuario(
		comando.ID,
		comando.Documento,
		comando.Nombre,
		comando.Apellido,
		comando.Edad,
	)
}

func CrearTransformadorUsuario() *TransformadorUsuario {
	return &TransformadorUsuario{}
}
