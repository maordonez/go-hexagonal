package manejador

import (
	"github.com/maordonez/go-hexagonal/aplicacion/comando"
	"github.com/maordonez/go-hexagonal/aplicacion/transformador"
	"github.com/maordonez/go-hexagonal/dominio/servicios"
)

type ManejadorRegistrarUsuario struct {
	transformadorUsuario     *transformador.TransformadorUsuario
	servicioRegistrarUsuario *servicios.ServicioRegistrarUsuario
}

func (t *ManejadorRegistrarUsuario) Ejecutar(comando *comando.ComandoUsuario) error {
	err, usuario := t.transformadorUsuario.CrearUsuario(comando)

	if err != nil {
		return err
	}

	return t.servicioRegistrarUsuario.Ejecutar(usuario)
}

func NewManejadorRegistrarUsuario(
	transformadorUsuario *transformador.TransformadorUsuario,
	servicioRegistrarUsuario *servicios.ServicioRegistrarUsuario) *ManejadorRegistrarUsuario {
	return &ManejadorRegistrarUsuario{
		transformadorUsuario:     transformadorUsuario,
		servicioRegistrarUsuario: servicioRegistrarUsuario,
	}
}
