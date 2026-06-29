package usuario

import (
	"biblioteca/internal/libro"
	"errors"
	"sort"
	"strings"
	"sync"
)

type MemoriaUsuarioRepo struct {
	sync.RWMutex
	usuarios map[int]*Usuario
	ultimoID int
}

func NewMemoriaUsuarioRepo() *MemoriaUsuarioRepo {
	repo := &MemoriaUsuarioRepo{
		usuarios: make(map[int]*Usuario),
	}
	return repo
}

func (repo *MemoriaUsuarioRepo) Crear(nuevoUsuario *Usuario) error {
	repo.Lock()
	defer repo.Unlock()

	repo.ultimoID++
	nuevoUsuario.ID = repo.ultimoID
	repo.usuarios[nuevoUsuario.ID] = nuevoUsuario
	return nil
}

func (repo *MemoriaUsuarioRepo) ObtenerTodos() []*Usuario {
	repo.RLock()
	defer repo.RUnlock()

	lista := make([]*Usuario, 0, len(repo.usuarios))
	for _, usuario := range repo.usuarios {
		lista = append(lista, usuario)
	}
	sort.Slice(lista, func(i, j int) bool {
		return lista[i].ID < lista[j].ID
	})
	return lista
}

func (repo *MemoriaUsuarioRepo) ObtenerPorID(id int) (*Usuario, error) {
	repo.RLock()
	defer repo.RUnlock()

	usuario, existe := repo.usuarios[id]

	if !existe {
		return &Usuario{}, errors.New("El usuario no existe!")
	}
	return usuario, nil
}

func (repo *MemoriaUsuarioRepo) Actualizar(u *Usuario) error {
	repo.Lock()
	defer repo.Unlock()

	if _, existe := repo.usuarios[u.ID]; !existe {
		return errors.New("No se puede actualizar! El usuario no existe")
	}
	repo.usuarios[u.ID] = u
	return nil
}

func (repo *MemoriaUsuarioRepo) Eliminar(id int) error {
	repo.Lock()
	defer repo.Unlock()

	if _, existe := repo.usuarios[id]; !existe {
		return errors.New("No se puede eliminar! El usuario no existe")
	}
	delete(repo.usuarios, id)
	return nil
}

func (repo *MemoriaUsuarioRepo) AsignarLibro(usuarioID int, libroID int) error {

	libroEncontrado, err := libro.NewMemoriaLibroRepo().ObtenerPorID(libroID)
	if err != nil {
		return err
	}

	usuarioEncontrado, err := repo.ObtenerPorID(usuarioID)
	if err != nil {
		return err
	}
	for _, libroSearch := range usuarioEncontrado.Libros {
		if libroSearch.ID == libroID {
			return errors.New("El usuario ya tiene este libro!")
		}
	}

	usuarioEncontrado.Libros = append(usuarioEncontrado.Libros, libroEncontrado)

	return nil
}

func (repo *MemoriaUsuarioRepo) ValidarUsuario(usuario Usuario) error {

	emailLimpio := strings.ToLower(usuario.Email)

	for _, usuario := range repo.usuarios {
		if strings.ToLower(usuario.Email) == emailLimpio {
			return errors.New("Este email ya esta registrado!")
		}
	}
	return nil

}
