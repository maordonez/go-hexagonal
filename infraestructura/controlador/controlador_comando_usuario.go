package controlador

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/maordonez/go-hexagonal/aplicacion/comando"
	"github.com/maordonez/go-hexagonal/aplicacion/manejador"
)

type ControladorComandoUsuario struct {
	manejadorRegistrarUsuario  *manejador.ManejadorRegistrarUsuario
	manejadorActualizarUsuario *manejador.ManejadorActualizarUsuario
}

func (c *ControladorComandoUsuario) registrar(w http.ResponseWriter, r *http.Request) {
	var comando comando.ComandoUsuario

	json.NewDecoder(r.Body).Decode(&comando)

	c.manejadorRegistrarUsuario.Ejecutar(&comando)

	w.Write([]byte("registro de usuario exitoso"))
	w.WriteHeader(201)
}

func (c *ControladorComandoUsuario) actualizar(w http.ResponseWriter, r *http.Request) {
	var comando comando.ComandoUsuario
	json.NewDecoder(r.Body).Decode(&comando)

	c.manejadorActualizarUsuario.Ejecutar(&comando)

	w.Write([]byte("actualizacion de usuario exitosa"))
	w.WriteHeader(200)
}

func (c *ControladorComandoUsuario) Inicializar(router *chi.Mux) {
	router.Post("/usuarios", c.registrar)
	router.Put("/usuarios/{idUsuario}", c.actualizar)

}

func NewControladorComandoUsuario(manejadorRegistrarUsuario *manejador.ManejadorRegistrarUsuario, manejadorActualizarUsuario *manejador.ManejadorActualizarUsuario) *ControladorComandoUsuario {
	return &ControladorComandoUsuario{
		manejadorRegistrarUsuario:  manejadorRegistrarUsuario,
		manejadorActualizarUsuario: manejadorActualizarUsuario,
	}
}
