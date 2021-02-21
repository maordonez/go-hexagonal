package comando

type ComandoPelicula struct {
	ID        string `json: "id"`
	Documento string `json: "documento"`
	Nombre    string `json: "nombre"`
	Apellido  string `json: "apellido"`
	Edad      int8   `json: "edad"`
}
