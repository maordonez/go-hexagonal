package entidades

import (
	"errors"

	"github.com/maordonez/go-hexagonal/comun/dominio/validadores"
)

const SE_DEBE_INGRESAR_UN_NOMBRE_VALIDO = "Se debe ingresar un nombre valido"
const SE_DEBE_INGRESAR_UN_DOCUMENTO_VALIDO = "Se debe ingresar un documento valido"

type Usuario struct {
	ID        string
	Documento string
	Nombre    string
	Apellido  string
	Edad      int8
}

func CrearEntidadUsuario(id string, documento string, nombre string, apellido string, edad int8) (error, *Usuario) {

	error := validadores.StringObligatorio(documento, SE_DEBE_INGRESAR_UN_DOCUMENTO_VALIDO)

	if error != nil {
		return error, nil
	}

	error = validadores.StringObligatorio(nombre, SE_DEBE_INGRESAR_UN_NOMBRE_VALIDO)

	if error != nil {
		return error, nil
	}

	usuario := new(Usuario)

	usuario.ID = id
	usuario.Documento = documento
	usuario.Nombre = nombre
	usuario.Apellido = apellido
	usuario.Edad = edad

	return nil, usuario
}

func validarStringObligatorio(cadena string, mensaje string) error {
	if len(cadena) != 0 {
		return errors.New(mensaje)
	}

	return nil
}
