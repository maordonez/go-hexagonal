package servicios

import (
	"github.com/maordonez/go-hexagonal/dominio/entidades"
	"github.com/maordonez/go-hexagonal/dominio/puertos/repositorios"
)

type ServicioRegistrarPelicula struct {
	repositorioPelicula repositorios.RepositorioPelicula
}

func (s *ServicioRegistrarPelicula) Ejecutar(pelicula *entidades.Pelicula) error {
	return s.repositorioPelicula.Actualizar(pelicula)
}

func NewServicioRegistrarPelicula(repositorioPelicula repositorios.RepositorioPelicula) *ServicioRegistrarPelicula {
	return &ServicioRegistrarPelicula{
		repositorioPelicula: repositorioPelicula,
	}
}
