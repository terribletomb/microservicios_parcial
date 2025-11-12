package repository

import (
	"context"
	"update-service/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type LibroRepository struct {
	Collection *mongo.Collection
}

func (r *LibroRepository) ActualizarLibro(ctx context.Context, id string, datos model.Libro) error {
	update := bson.M{"$set": datos}
	_, err := r.Collection.UpdateOne(ctx, bson.M{"_id": id}, update)
	return err
}
