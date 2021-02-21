package manejador

import (
	"github.com/maordonez/go-hexagonal/aplicacion/comando"
	"github.com/maordonez/go-hexagonal/aplicacion/transformador"
	"github.com/maordonez/go-hexagonal/dominio/servicios"
)

type ManejadorActualizarUsuario struct {
	transformadorUsuario      *transformador.TransformadorUsuario
	servicioActualizarUsuario *servicios.ServicioActualizarUsuario
}

func (m ManejadorActualizarUsuario) Ejecutar(comando *comando.ComandoUsuario) error {
	err, usuario := m.transformadorUsuario.CrearUsuario(comando)

	if err != nil {
		return err
	}

	return m.servicioActualizarUsuario.Ejecutar(usuario)

}

func NewManejadorActualizarUsuario(transformadorUsuario *transformador.TransformadorUsuario, servicioActualizarUsuario *servicios.ServicioActualizarUsuario) *ManejadorActualizarUsuario {
	return &ManejadorActualizarUsuario{
		transformadorUsuario:      transformadorUsuario,
		servicioActualizarUsuario: servicioActualizarUsuario,
	}
}
