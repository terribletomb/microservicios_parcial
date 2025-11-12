package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Libro representa un libro en la base de datos
type Libro struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Titulo        string             `bson:"titulo" json:"titulo"`
	Autor         string             `bson:"autor" json:"autor"`
	Anio          int                `bson:"anio" json:"anio"`
	Paginas       int                `bson:"paginas" json:"paginas"`
	Editorial     string             `bson:"editorial" json:"editorial"`
	CreadoEn      int64              `bson:"creado_en,omitempty" json:"creado_en,omitempty"`
	ActualizadoEn int64              `bson:"actualizado_en,omitempty" json:"actualizado_en,omitempty"`
}
