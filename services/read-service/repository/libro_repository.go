package repository

import (
	"context"
	"log"

	"read-service/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type LibroRepository struct {
	Collection *mongo.Collection
}

func (r *LibroRepository) ObtenerTodos(ctx context.Context) ([]model.Libro, error) {
	cursor, err := r.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var libros []model.Libro
	for cursor.Next(ctx) {
		var libro model.Libro
		if err := cursor.Decode(&libro); err != nil {
			log.Println("Error al decodificar libro:", err)
			continue
		}
		libros = append(libros, libro)
	}

	return libros, nil
}
