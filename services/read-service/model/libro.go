package model

type Libro struct {
	ID          string `bson:"_id,omitempty" json:"id"`
	Titulo      string `bson:"titulo" json:"titulo"`
	Autor       string `bson:"autor" json:"autor"`
	Anio        int    `bson:"anio" json:"anio"`
	Editorial   string `bson:"editorial" json:"editorial"`
}
