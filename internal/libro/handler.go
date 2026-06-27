package libro

import (
	"biblioteca/pkg/response"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

func (repo *MemoriaLibroRepo) HandlerCrear(w http.ResponseWriter, r *http.Request) {
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(r.Body)

	var nuevoLibro Libro

	err := json.NewDecoder(r.Body).Decode(&nuevoLibro)
	if err != nil {
		res := map[string]string{"Error": "Json mal formado"}
		response.ResponderJSON(w, http.StatusBadRequest, res)
		return
	}

	if err := repo.ValidarLibro(nuevoLibro.Titulo, nuevoLibro.Autor); err != nil {
		res := map[string]string{"Error": "El libro ya existe"}
		response.ResponderJSON(w, http.StatusBadRequest, res)
		return
	}

	if err := repo.Crear(&nuevoLibro); err != nil {
		res := map[string]string{"Error": "No se pudo crear el libro"}
		response.ResponderJSON(w, http.StatusBadRequest, res)
	}

	res := map[string]string{"Message": "Libro creado exitosamente!", "Titulo": nuevoLibro.Titulo}
	response.ResponderJSON(w, http.StatusCreated, res)
}

func (repo *MemoriaLibroRepo) HandleObtenerPorID(w http.ResponseWriter, r *http.Request) {

	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		res := map[string]string{"error": "El ID provisto debe ser un número entero válido"}
		response.ResponderJSON(w, http.StatusBadRequest, res)
		return
	}

	libro, err := repo.ObtenerPorID(id)
	if err != nil {
		res := map[string]string{"error": "El libro con el ID solicitado no existe"}
		response.ResponderJSON(w, http.StatusBadRequest, res)
		return
	}
	response.ResponderJSON(w, http.StatusOK, libro)
}

func (repo *MemoriaLibroRepo) HandleActualizar(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "ID inválido"}`))
		return
	}

	var libroEditado Libro
	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&libroEditado); err != nil {
		res := map[string]string{"Error": "JSON invalido"}
		response.ResponderJSON(w, http.StatusBadRequest, res)
		return
	}
	libroEditado.ID = id
	if err := repo.Actualizar(libroEditado); err != nil {
		res := map[string]string{"Error": err.Error()}
		response.ResponderJSON(w, http.StatusNotFound, res)
		return
	}

	response.ResponderJSON(w, http.StatusOK, libroEditado)
}

func (repo *MemoriaLibroRepo) HandleEliminar(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodDelete {
		res := map[string]string{"Error": "Method no permitido"}
		response.ResponderJSON(w, http.StatusBadRequest, res)
		return
	}

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "ID inválido"}`))
		return
	}

	if err := repo.Eliminar(id); err != nil {
		response.ResponderJSON(w, http.StatusBadRequest, err.Error())
		return
	}
	response.ResponderJSON(w, http.StatusNoContent, "...")
}

func (repo *MemoriaLibroRepo) HandleObtenerTodos(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		res := map[string]string{"Error": "Method no permitido"}
		response.ResponderJSON(w, http.StatusBadRequest, res)
		return
	}

	libros := repo.ObtenerTodos()
	response.ResponderJSON(w, http.StatusOK, libros)
}
