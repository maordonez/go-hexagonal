package repositorios

import "github.com/maordonez/go-hexagonal/dominio/entidades"

type RepositorioUsuario interface {
	Registrar(usuario *entidades.Usuario) error
	Actualizar(usuario *entidades.Usuario) error
}
