package libro

type Libro struct {
	ID         int    `json:"id"`
	Titulo     string `json:"titulo"`
	Autor      string `json:"autor"`
	Paginas    int    `json:"paginas"`
	Disponible bool   `json:"disponible"`
}

type Repositorio interface {
	Crear()
	ObtenerPorID()
	ObtenerTodos()
	Actualizar()
	Eliminar()
}
