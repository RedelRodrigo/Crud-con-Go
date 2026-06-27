package routers

import (
	"biblioteca/internal/libro"
	"biblioteca/pkg/response"
	"net/http"
)

func Health(mux *http.ServeMux) {

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			res := map[string]string{"message": "OK"}
			response.ResponderJSON(w, http.StatusOK, res)
		default:
			w.Write([]byte("Method is not GET"))
		}
	})
}

func Libros(mux *http.ServeMux) {
	repo := libro.NewMemoriaLibroRepo()

	mux.HandleFunc("/libros", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			repo.HandleObtenerTodos(w, r)
		case http.MethodPost:
			repo.HandlerCrear(w, r)
		default:
			res := map[string]string{"Error": "Method not found"}
			response.ResponderJSON(w, http.StatusBadRequest, res)
		}
	})

	mux.HandleFunc("/libros/{id}", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			repo.HandleObtenerPorID(w, r)
		case http.MethodPut:
			repo.HandleActualizar(w, r)
		case http.MethodDelete:
			repo.HandleEliminar(w, r)
		}

	})

}
