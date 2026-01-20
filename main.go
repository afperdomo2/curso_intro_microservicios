package main

import "fmt"

func main() {
	fmt.Println("Este proyecto ha sido migrado a una arquitectura de microservicios.")
	fmt.Println("Por favor, ejecuta los servicios individualmente desde la carpeta 'services/'.")
	fmt.Println("Ejemplo: go run services/GetAdults/main.go")
	fmt.Println("O utiliza docker-compose para desplegar todo el entorno.")
}
