package libro

import (
	"errors"
	"sort"
	"strings"
	"sync"
)

type MemoriaLibroRepo struct {
	libros map[int]Libro
	sync.RWMutex
	ultimoID int
	Repositorio
}

func NewMemoriaLibroRepo() *MemoriaLibroRepo {
	repo := &MemoriaLibroRepo{
		libros: make(map[int]Libro),
	}

	repo.libros[1] = Libro{ID: 1, Titulo: "Cien años de soledad", Autor: "Gabriel García Márquez", Paginas: 432, Disponible: true}
	repo.libros[2] = Libro{ID: 2, Titulo: "1984", Autor: "George Orwell", Paginas: 328, Disponible: false}

	return repo
}

func (repo *MemoriaLibroRepo) Crear(nuevoLibro *Libro) error {
	repo.Lock()
	defer repo.Unlock()

	repo.ultimoID++
	nuevoLibro.ID = repo.ultimoID
	repo.libros[nuevoLibro.ID] = *nuevoLibro
	return nil
}

func (repo *MemoriaLibroRepo) ObtenerTodos() []Libro {
	repo.RLock()
	defer repo.RUnlock()

	lista := make([]Libro, 0, len(repo.libros))
	for _, libro := range repo.libros {
		lista = append(lista, libro)
	}
	sort.Slice(lista, func(i, j int) bool {
		return lista[i].ID < lista[j].ID
	})
	return lista
}

func (repo *MemoriaLibroRepo) ObtenerPorID(id int) (Libro, error) {
	repo.RLock()
	defer repo.RUnlock()

	libro, existe := repo.libros[id]

	if !existe {
		return Libro{}, errors.New("El libro no existe!")
	}
	return libro, nil
}

func (repo *MemoriaLibroRepo) Actualizar(l Libro) error {
	repo.Lock()
	defer repo.Unlock()

	if _, existe := repo.libros[l.ID]; !existe {
		return errors.New("No se puede actualizar! El libro no existe")
	}
	repo.libros[l.ID] = l
	return nil
}

func (repo *MemoriaLibroRepo) Eliminar(id int) error {
	repo.Lock()
	defer repo.Unlock()

	if _, existe := repo.libros[id]; !existe {
		return errors.New("No se puede eliminar! El libro no existe")
	}
	delete(repo.libros, id)
	return nil
}

func (repo *MemoriaLibroRepo) ValidarLibro(titulo string, autor string) error {
	for _, libro := range repo.libros {
		if strings.ToLower(libro.Titulo) == strings.ToLower(titulo) && strings.ToLower(autor) == strings.ToLower(libro.Autor) {
			return errors.New("El libro ya existe!")
		}
	}
	return nil
}
