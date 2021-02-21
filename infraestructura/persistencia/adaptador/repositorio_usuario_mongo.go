package adaptador

import (
	"context"
	"errors"

	"github.com/maordonez/go-hexagonal/dominio/entidades"
	"github.com/maordonez/go-hexagonal/dominio/puertos/repositorios"
	"github.com/maordonez/go-hexagonal/infraestructura/persistencia/mongoEntidades"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RepositorioUsuarioMongo struct {
	cliente         *mongo.Client
	nombreDb        string
	nombreColeccion string
}

func (repo *RepositorioUsuarioMongo) Registrar(usuario *entidades.Usuario) error {
	coleccion := repo.cliente.Database(repo.nombreColeccion).Collection(repo.nombreColeccion)

	usuarioMongo := mongoEntidades.MongoUsuario{
		ID:        usuario.ID,
		Documento: usuario.Documento,
		Nombre:    usuario.Nombre,
		Apellido:  usuario.Apellido,
		Edad:      int(usuario.Edad),
	}

	result, err := coleccion.InsertOne(context.TODO(), usuarioMongo)

	if result.InsertedID == nil {
		return errors.New("error al registrar")
	}

	return err
}

func (repo *RepositorioUsuarioMongo) Actualizar(usuario *entidades.Usuario) error {

	usuarioMongo := mongoEntidades.MongoUsuario{
		ID:        usuario.ID,
		Documento: usuario.Documento,
		Nombre:    usuario.Nombre,
		Apellido:  usuario.Apellido,
		Edad:      int(usuario.Edad),
	}

	filter := bson.D{primitive.E{Key: "ID", Value: usuarioMongo.ID}}

	coleccion := repo.cliente.Database(repo.nombreColeccion).Collection(repo.nombreColeccion)
	result, err := coleccion.UpdateOne(context.TODO(), filter, usuarioMongo)

	if result.ModifiedCount == 0 {
		return errors.New("error al actualizar")
	}

	return err
}

func NewRepositorioUsuario(cliente *mongo.Client) repositorios.RepositorioUsuario {
	return &RepositorioUsuarioMongo{cliente: cliente, nombreColeccion: "usuarios", nombreDb: "biblioteca"}
}
