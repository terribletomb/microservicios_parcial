package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Libro struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Titulo    string             `bson:"titulo" json:"titulo"`
	Autor     string             `bson:"autor" json:"autor"`
	Editorial string             `bson:"editorial" json:"editorial"`
	Anio      int                `bson:"anio" json:"anio"`
	Genero    string             `bson:"genero" json:"genero"`
	ISBN      string             `bson:"isbn" json:"isbn"`
}
