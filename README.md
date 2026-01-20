# Curso de introducción a microservicios

## Descripción

Este proyecto es ahora una arquitectura de **microservicios** construida con Go, Gin Framework y GORM. Cada endpoint ha sido desacoplado en su propio servicio independiente.

## Estructura del Proyecto

El proyecto se ha reestructurado de la siguiente manera:

- `pkg/`: Paquetes compartidos.
  - `database/`: Lógica de conexión a base de datos (PostgreSQL).
  - `models/`: Definiciones de estructuras de datos (Adult, Child).
- `services/`: Contiene los microservicios individuales.
  - `GetAdults/`: Servicio para obtener todos los adultos (`GET /Adults`).
  - `GetChildren/`: Servicio para obtener todos los niños (`GET /Children`).
  - `GetAdultById/`: Servicio para obtener un adulto por ID (`GET /Adults/:id`).
  - `GetChildById/`: Servicio para obtener un niño por ID (`GET /Children/:id`).
  - `AddAdult/`: Servicio para agregar un adulto (`POST /Add/Adults`).
  - `AddChild/`: Servicio para agregar un niño (`POST /Add/Children`).

## Configuración de Base de Datos

Todos los servicios comparten la misma base de datos PostgreSQL:

- **Host**: localhost
- **Puerto**: 5433
- **Usuario**: devuser
- **Contraseña**: devpassword123
- **Base de datos**: intro_microservicios

## Ejecución de Microservicios

Cada servicio es una aplicación Go independiente. Para ejecutar un servicio específico, navega a su directorio o usa `go run`.

**Nota**: Todos los servicios están configurados para escuchar en el puerto `8080` por defecto. Para ejecutarlos simultáneamente, necesitarás configurar puertos diferentes o usar contenedores.

Ejemplo para ejecutar el servicio `GetAdults`:

```bash
go run services/GetAdults/main.go
```

Ejemplo para ejecutar el servicio `AddChild`:

```bash
go run services/AddChild/main.go
```

## Documentación API (Swagger)

Cada servicio mantiene su propia instancia de Swagger si se ejecuta independientemente.

<http://localhost:8080/swagger/index.html>
