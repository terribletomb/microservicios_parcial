package controllers

import (
	"create-service/models"
	"create-service/services"
	"encoding/json"
	"net/http"
)

type ControladorCrear struct {
	servicio services.ServicioCrear
}

func NuevoControladorCrear(s services.ServicioCrear) *ControladorCrear {
	return &ControladorCrear{servicio: s}
}

func (c *ControladorCrear) CrearLibro(w http.ResponseWriter, r *http.Request) {
	var libro models.Libro
	if err := json.NewDecoder(r.Body).Decode(&libro); err != nil {
		http.Error(w, "Error al decodificar el cuerpo", http.StatusBadRequest)
		return
	}
	creado, err := c.servicio.CrearLibro(r.Context(), &libro)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(creado)
}
