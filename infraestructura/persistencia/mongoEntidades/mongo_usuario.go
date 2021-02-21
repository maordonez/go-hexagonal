package mongoEntidades

type MongoUsuario struct {
	ID        string `bson:"-"`
	Documento string `bson:"documento"`
	Nombre    string `bson:"nombre"`
	Apellido  string `bson:"apellido"`
	Edad      int    `bson:"edad"`
}
