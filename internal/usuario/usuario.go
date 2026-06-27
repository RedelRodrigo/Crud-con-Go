package usuario

import "biblioteca/internal/libro"

type Usuario struct {
	ID       int           `json:"id"`
	Nombre   string        `json:"nombre"`
	Apellido string        `json:"apellido"`
	Email    string        `json:"email"`
	Libros   []libro.Libro `json:"libros"`
}

type Repositorio interface {
	Crear()
	ObtenerPorID()
	ObtenerTodos()
	Actualizar()
	Eliminar()
	AsignarLibro(usuarioID int, libro libro.Libro) error
}
