package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const DATABASE_HOST = "DATABASE_HOST"
const DATABASE_PORT = "DATABASE_PORT"
const DATABASE_USERNAME = "DATABASE_USERNAME"
const DATABASE_PASSWORD = "DATABASE_PASSWORD"
const DATABASE_NAME = "DATABASE_NAME"

func IniciarDatabase() (error, *mongo.Client) {
	servidor := os.Getenv(DATABASE_HOST)
	puerto := os.Getenv(DATABASE_PORT)
	dbNombre := os.Getenv(DATABASE_NAME)
	usuario := os.Getenv(DATABASE_USERNAME)
	clave := os.Getenv(DATABASE_USERNAME)

	clientOpts := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s:%s/%s", usuario, clave, servidor, puerto, dbNombre))
	client, err := mongo.Connect(context.TODO(), clientOpts)

	if err != nil {
		log.Fatal(err)
		return err, nil
	}

	// Check the connections
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return err, nil
	}

	fmt.Println("Congratulations, you're already connected to MongoDB!")
	return err, client
}
