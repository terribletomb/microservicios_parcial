package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"read-service/controller"
	"read-service/repository"
	"read-service/service"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	mongoURI := os.Getenv("MONGO_URI")
	dbName := os.Getenv("MONGO_DB")
	collectionName := os.Getenv("MONGO_COLLECTION")

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal("Error conectando a MongoDB:", err)
	}

	db := client.Database(dbName)
	repo := &repository.LibroRepository{Collection: db.Collection(collectionName)}
	svc := &service.LibroService{Repo: repo}
	ctrl := &controller.LibroController{Service: svc}

	http.HandleFunc("/libros", ctrl.ObtenerLibros)

	fmt.Println("📘 Servicio READ corriendo en puerto 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
