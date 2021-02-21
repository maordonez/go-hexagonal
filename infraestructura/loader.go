package infraestructura

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/maordonez/go-hexagonal/aplicacion/manejador"
	"github.com/maordonez/go-hexagonal/aplicacion/transformador"
	"github.com/maordonez/go-hexagonal/dominio/servicios"
	"github.com/maordonez/go-hexagonal/infraestructura/config"
	"github.com/maordonez/go-hexagonal/infraestructura/controlador"
	"github.com/maordonez/go-hexagonal/infraestructura/persistencia/adaptador"
)

// StartAplication metodo para iniciar aplicacion
func StartAplication() {

	err, cliente := config.IniciarDatabase()

	if err != nil {
		return
	}

	repositorioUsuario := adaptador.NewRepositorioUsuario(cliente)

	transformadorUsuario := transformador.CrearTransformadorUsuario()

	servicioRegistrarUsuario := servicios.NewServicioRegistrarUsuario(repositorioUsuario)
	servicioActualizarUsuario := servicios.NewServicioActualizarUsuario(repositorioUsuario)

	manejadorRegistrarUsuario := manejador.NewManejadorRegistrarUsuario(transformadorUsuario, servicioRegistrarUsuario)
	manejadorActualizarUsuario := manejador.NewManejadorActualizarUsuario(transformadorUsuario, servicioActualizarUsuario)

	controladorComandoUsuario := controlador.NewControladorComandoUsuario(manejadorRegistrarUsuario, manejadorActualizarUsuario)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	controladorComandoUsuario.Inicializar(r)

	http.ListenAndServe(":3000", r)

}
