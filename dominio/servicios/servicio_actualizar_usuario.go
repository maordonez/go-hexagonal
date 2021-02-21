package servicios

import (
	"github.com/maordonez/go-hexagonal/dominio/entidades"
	"github.com/maordonez/go-hexagonal/dominio/puertos/repositorios"
)

type ServicioActualizarUsuario struct {
	repositorioUsuario repositorios.RepositorioUsuario
}

func (s *ServicioActualizarUsuario) Ejecutar(usuario *entidades.Usuario) error {
	return s.repositorioUsuario.Actualizar(usuario)
}

func NewServicioActualizarUsuario(repositorioUsuario repositorios.RepositorioUsuario) *ServicioActualizarUsuario {
	return &ServicioActualizarUsuario{repositorioUsuario: repositorioUsuario}
}
