package service

import (
	"context"
	"delete-service/repository"
)

type LibroService struct {
	Repo *repository.LibroRepository
}

func (s *LibroService) EliminarLibro(ctx context.Context, id string) error {
	return s.Repo.EliminarLibro(ctx, id)
}
