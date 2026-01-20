# Curso de introducción a microservicios

## Descripción

Este proyecto es una API simple construida con Go y Gin Framework como parte del curso de introducción a microservicios.

## Endpoints implementados

La API incluye los siguientes endpoints:

- **GET** `/Adults`: Obtiene la lista de adultos.
- **GET** `/Children`: Obtiene la lista de niños.
- **GET** `/Adults/{id}`: Obtiene un adulto por ID.
- **GET** `/Children/{id}`: Obtiene un niño por ID.
- **POST** `/Add/Adults`: Agrega un nuevo adulto.
- **POST** `/Add/Children`: Agrega un nuevo niño.

## Ejecución

Para ejecutar el proyecto, asegúrate de tener Go instalado y corre el siguiente comando:

```bash
go run main.go
```

## Documentación API (Swagger)

La documentación de la API está disponible a través de Swagger. Una vez que el servidor esté corriendo, puedes acceder a ella en:

<http://localhost:8080/swagger/index.html>
