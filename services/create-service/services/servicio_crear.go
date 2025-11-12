package services

import (
	"context"
	"create-service/models"
	"create-service/repositories"
	"errors"
)

type ServicioCrear interface {
	CrearLibro(ctx context.Context, libro *models.Libro) (*models.Libro, error)
}

type servicioCrear struct {
	repo repositories.LibroRepositorio
}

func NuevoServicioCrear(repo repositories.LibroRepositorio) ServicioCrear {
	return &servicioCrear{repo: repo}
}

func (s *servicioCrear) CrearLibro(ctx context.Context, libro *models.Libro) (*models.Libro, error) {
	if libro.Titulo == "" || libro.Autor == "" {
		return nil, errors.New("el título y el autor son obligatorios")
	}
	return s.repo.Insertar(ctx, libro)
}
