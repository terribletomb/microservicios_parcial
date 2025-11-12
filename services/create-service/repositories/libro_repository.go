package repositories

import (
	"context"
	"errors"
	"time"

	"create-service/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Interfaz pública del repositorio
type LibroRepositorio interface {
	Insertar(ctx context.Context, libro *models.Libro) (*models.Libro, error)
	BuscarPorID(ctx context.Context, id string) (*models.Libro, error)
	BuscarTodos(ctx context.Context) ([]*models.Libro, error)
	Actualizar(ctx context.Context, id string, datos bson.M) (*models.Libro, error)
	Eliminar(ctx context.Context, id string) error
}

// Implementación MongoDB
type mongoLibroRepo struct {
	coleccion *mongo.Collection
}

// Constructor del repositorio
func NuevoLibroRepositorio(coleccion *mongo.Collection) LibroRepositorio {
	return &mongoLibroRepo{coleccion: coleccion}
}

func (r *mongoLibroRepo) Insertar(ctx context.Context, libro *models.Libro) (*models.Libro, error) {
	libro.ID = primitive.NewObjectID()
	libro.CreadoEn = time.Now().Unix()
	_, err := r.coleccion.InsertOne(ctx, libro)
	if err != nil {
		return nil, err
	}
	return libro, nil
}

func (r *mongoLibroRepo) BuscarPorID(ctx context.Context, id string) (*models.Libro, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("ID inválido")
	}
	var libro models.Libro
	if err := r.coleccion.FindOne(ctx, bson.M{"_id": oid}).Decode(&libro); err != nil {
		return nil, err
	}
	return &libro, nil
}

func (r *mongoLibroRepo) BuscarTodos(ctx context.Context) ([]*models.Libro, error) {
	cur, err := r.coleccion.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var libros []*models.Libro
	for cur.Next(ctx) {
		var l models.Libro
		if err := cur.Decode(&l); err != nil {
			return nil, err
		}
		libros = append(libros, &l)
	}
	return libros, nil
}

func (r *mongoLibroRepo) Actualizar(ctx context.Context, id string, datos bson.M) (*models.Libro, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("ID inválido")
	}
	datos["actualizado_en"] = time.Now().Unix()
	res := r.coleccion.FindOneAndUpdate(ctx, bson.M{"_id": oid}, bson.M{"$set": datos})
	var actualizado models.Libro
	if err := res.Decode(&actualizado); err != nil {
		return nil, err
	}
	return &actualizado, nil
}

func (r *mongoLibroRepo) Eliminar(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("ID inválido")
	}
	res, err := r.coleccion.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return errors.New("no se eliminó ningún documento")
	}
	return nil
}
