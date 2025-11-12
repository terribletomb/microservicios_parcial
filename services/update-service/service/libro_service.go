package service

import (
	"context"
	"update-service/model"
	"update-service/repository"
)

type LibroService struct {
	Repo *repository.LibroRepository
}

func (s *LibroService) ActualizarLibro(ctx context.Context, id string, libro model.Libro) error {
	return s.Repo.ActualizarLibro(ctx, id, libro)
}
