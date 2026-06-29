package usuario

import (
	"biblioteca/pkg/response"
	"encoding/json"
	"net/http"
	"strconv"
)

func (repo *MemoriaUsuarioRepo) HandlerCrear(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var nuevoUsuario Usuario

	if err := json.NewDecoder(r.Body).Decode(&nuevoUsuario); err != nil {
		res := map[string]string{"Error": "JSON invalido"}
		response.ResponderJSON(w, http.StatusBadRequest, res)
		return
	}

	if err := repo.ValidarUsuario(nuevoUsuario); err != nil {
		res := map[string]string{"Error": err.Error()}
		response.ResponderJSON(w, http.StatusBadRequest, res)
		return
	}

	if err := repo.Crear(&nuevoUsuario); err != nil {
		res := map[string]string{"Error": err.Error()}
		response.ResponderJSON(w, http.StatusBadRequest, res)
		return
	}

	response.ResponderJSON(w, http.StatusCreated, nuevoUsuario)
}

func (repo *MemoriaUsuarioRepo) HandlerActualizar(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	defer r.Body.Close()
	var nuevoUsuario Usuario

	if err := json.NewDecoder(r.Body).Decode(&nuevoUsuario); err != nil {
		res := map[string]string{"Error": "JSON invalido"}
		response.ResponderJSON(w, http.StatusBadRequest, res)
		return
	}

	nuevoUsuario.ID = id
	if err := repo.Actualizar(&nuevoUsuario); err != nil {
		res := map[string]string{"Error": err.Error()}
		response.ResponderJSON(w, http.StatusBadRequest, res)
		return
	}

	response.ResponderJSON(w, http.StatusCreated, nuevoUsuario)
}

func (repo *MemoriaUsuarioRepo) HandlerObtenerPorID(w http.ResponseWriter, r *http.Request) {

	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		res := map[string]string{"error": "El ID provisto debe ser un número entero válido"}
		response.ResponderJSON(w, http.StatusBadRequest, res)
		return
	}

	usuario, err := repo.ObtenerPorID(id)
	if err != nil {
		res := map[string]string{"error": "El usuario con el ID solicitado no existe"}
		response.ResponderJSON(w, http.StatusBadRequest, res)
		return
	}
	response.ResponderJSON(w, http.StatusOK, usuario)
}

func (repo *MemoriaUsuarioRepo) HandlerEliminar(w http.ResponseWriter, r *http.Request) {
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
	response.ResponderJSON(w, http.StatusNoContent, id)
}

func (repo *MemoriaUsuarioRepo) HandlerObtenerTodos(w http.ResponseWriter, r *http.Request) {
	usuarios := repo.ObtenerTodos()
	response.ResponderJSON(w, http.StatusOK, usuarios)
}

func (repo *MemoriaUsuarioRepo) HandlerAsignarLibro(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var requestData struct {
		UsuarioID int `json:"usuario_id"`
		LibroID   int `json:"libro_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		res := map[string]string{"Error": "JSON invalido"}
		response.ResponderJSON(w, http.StatusBadRequest, res)
		return
	}

	err := repo.AsignarLibro(requestData.UsuarioID, requestData.LibroID)
	if err != nil {
		res := map[string]string{"Error": err.Error()}
		response.ResponderJSON(w, http.StatusBadRequest, res)
		return
	}

	res := map[string]string{"Mensaje": "Libro asignado correctamente al usuario"}
	response.ResponderJSON(w, http.StatusOK, res)
}
