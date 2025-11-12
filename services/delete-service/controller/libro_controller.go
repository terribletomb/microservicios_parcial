package controller

import (
	"context"
	"delete-service/service"
	"net/http"

	"github.com/gorilla/mux"
)

type LibroController struct {
	Service *service.LibroService
}

func (c *LibroController) EliminarLibro(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if err := c.Service.EliminarLibro(context.Background(), id); err != nil {
		http.Error(w, "Error al eliminar libro", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Libro eliminado correctamente"))
}
