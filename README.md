Markdown# 📚 Biblioteca CRUD API con Go

Una API RESTful construida con Go para gestionar el catálogo de una biblioteca. Este proyecto demuestra la implementación de un sistema CRUD (Crear, Leer, Actualizar, Eliminar) robusto y eficiente para el manejo de libros.

## 🚀 Características

* **C**reate: Añadir nuevos libros al catálogo con su información (título, autor, año, etc.).
* **R**ead: Obtener la lista completa de libros o buscar un libro específico por su ID.
* **U**pdate: Modificar la información de un libro existente.
* **D**elete: Eliminar registros del sistema.

## 🛠️ Tecnologías Utilizadas

* **Lenguaje:** [Go (Golang)]
* **Base de Datos:** [PostgreSQL]


## ⚙️ Instalación y Configuración

Sigue estos pasos para correr el proyecto en tu entorno local:

1. **Clonar el repositorio:**
   ```bash
   git clone [https://github.com/RedelRodrigo/Crud-con-Go.git](https://github.com/RedelRodrigo/Crud-con-Go.git)
   cd Crud-con-Go
   
Instalar dependencias:Bash
go mod tidy

Configurar variables de entorno:Crea un archivo .env en la raíz del proyecto (puedes basarte en un .env.example si lo hay) y configura tu conexión a la base de datos:Fragmento de códigoDB_HOST=localhost
DB_USER=root
DB_PASSWORD=tu_password
DB_NAME=biblioteca
DB_PORT=5432

Ejecutar la aplicación:Bash
go run main.go

La API estará disponible en http://localhost:8080.📡 Endpoints PrincipalesMétodoEndpointDescripción
GET/api/librosObtiene todos los libros
GET/api/libros/{id}Obtiene un libro por su ID
POST/api/librosCrea un nuevo registro de libro
PUT/api/libros/{id}Actualiza la información de un libro
DELETE/api/libros/{id}Elimina un libro del catálogo💡 
Ejemplos de PeticionesCrear un libro (POST /api/libros)
JSON{
  "titulo": "El Aleph",
  "autor": "Jorge Luis Borges",
  "año_publicacion": 1949,
  "genero": "Ficción"
}
👨‍💻 AutorRodrigo Redel - LinkedIn - GitHub
