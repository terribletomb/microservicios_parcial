package services_test

import (
	"context"
	"testing"

	"create-service/models"
	"create-service/services"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type fakeRepo struct{}

func (f *fakeRepo) Insertar(ctx context.Context, libro *models.Libro) (*models.Libro, error) {
	libro.ID = primitive.NewObjectID()
	return libro, nil
}

func testCrearLibro(t *testing.T) {
	s := services.NuevoServicioCrear(&fakeRepo{})
	libro := &models.Libro{
		Titulo:    "El Principito",
		Autor:     "Antoine de Saint-Exupéry",
		Anio:      1943,
		Paginas:   96,
		Editorial: "Anagrama",
	}
	creado, err := s.CrearLibro(context.Background(), libro)
	assert.NoError(t, err)
	assert.NotNil(t, creado)
	assert.NotEmpty(t, creado.ID)
}

func TestCrearLibro(t *testing.T) {
	testCrearLibro(t)
}
