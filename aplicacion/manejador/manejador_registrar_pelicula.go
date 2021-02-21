package manejador

import (
	"github.com/maordonez/go-hexagonal/aplicacion/comando"
	"github.com/maordonez/go-hexagonal/aplicacion/transformador"
	"github.com/maordonez/go-hexagonal/dominio/servicios"
)

type ManejadorRegistrarPelicula struct {
	transformadorPelicula     *transformador.TranformadorPelicula
	servicioRegistrarPelicula *servicios.ServicioRegistrarPelicula
}

func (m *ManejadorRegistrarPelicula) Ejecutar(comando *comando.ComandoPelicula) error {
	err, pelicula := m.transformadorPelicula.CrearPelicula(comando)

	if err != nil {
		return err
	}

	return m.servicioRegistrarPelicula.Ejecutar(pelicula)
}

func NewManejadorRegistrarPelicula(
	transformadorPelicula *transformador.TranformadorPelicula,
	servicioRegistrarPelicula *servicios.ServicioRegistrarPelicula) *ManejadorRegistrarPelicula {
	return &ManejadorRegistrarPelicula{
		transformadorPelicula:     transformadorPelicula,
		servicioRegistrarPelicula: servicioRegistrarPelicula,
	}
}
