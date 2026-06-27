package routers

import (
	"biblioteca/internal/usuario"
	"biblioteca/pkg/response"
	"net/http"
)

func Usuario(mux *http.ServeMux) {
	repo := usuario.NewMemoriaUsuarioRepo()

	mux.HandleFunc("/usuarios", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			repo.HandlerObtenerTodos(w, r)
		case http.MethodPost:
			repo.HandlerCrear(w, r)
		default:
			res := map[string]string{"Error": "Method not found"}
			response.ResponderJSON(w, http.StatusBadRequest, res)
		}
	})
	mux.HandleFunc("/usuarios/{id}", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			repo.HandlerObtenerPorID(w, r)
		case http.MethodPut:
			repo.HandlerActualizar(w, r)
		case http.MethodDelete:
			repo.HandlerEliminar(w, r)
		}

	})
	mux.HandleFunc("/usuarios/libro", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			repo.HandlerAsignarLibro(w, r)
		default:
			res := map[string]string{"Error": "Method not found"}
			response.ResponderJSON(w, http.StatusBadRequest, res)
		}
	})
}
