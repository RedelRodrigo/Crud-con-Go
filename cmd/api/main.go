package main

import (
	"biblioteca/internal/trasport/routers"
	"biblioteca/pkg/middleware"
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	routers.Health(mux)
	routers.Libros(mux)
	routers.Usuario(mux)

	fmt.Printf("Servidor corriendo en el http://localhost:8080 \n")
	err := http.ListenAndServe(":8080", middleware.Logging(mux))
	if err != nil {
		return
	}
}
