package servicios

import (
	"github.com/maordonez/go-hexagonal/dominio/entidades"
	"github.com/maordonez/go-hexagonal/dominio/puertos/repositorios"
)

type ServicioRegistrarUsuario struct {
	repositorioUsuario repositorios.RepositorioUsuario
}

func (s *ServicioRegistrarUsuario) Ejecutar(usuario *entidades.Usuario) error {
	return s.repositorioUsuario.Registrar(usuario)
}

func NewServicioRegistrarUsuario(repositorioUsuario repositorios.RepositorioUsuario) *ServicioRegistrarUsuario {
	return &ServicioRegistrarUsuario{repositorioUsuario: repositorioUsuario}
}
