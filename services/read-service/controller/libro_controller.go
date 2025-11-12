package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"read-service/service"
)

type LibroController struct {
	Service *service.LibroService
}

func (c *LibroController) ObtenerLibros(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	libros, err := c.Service.ObtenerLibros(ctx)
	if err != nil {
		http.Error(w, "Error al obtener libros", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(libros)
}
