package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"update-service/model"
	"update-service/service"

	"github.com/gorilla/mux"
)

type LibroController struct {
	Service *service.LibroService
}

func (c *LibroController) ActualizarLibro(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var libro model.Libro
	json.NewDecoder(r.Body).Decode(&libro)

	if err := c.Service.ActualizarLibro(context.Background(), id, libro); err != nil {
		http.Error(w, "Error al actualizar libro", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Libro actualizado correctamente"))
}
