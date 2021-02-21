package repositorios

import "github.com/maordonez/go-hexagonal/dominio/entidades"

type RepositorioPelicula interface {
	Registrar(usuario *entidades.Pelicula) error
	Actualizar(usuario *entidades.Pelicula) error
}
