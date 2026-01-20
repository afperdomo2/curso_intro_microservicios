# Curso de introducción a microservicios

## Descripción

Este proyecto es una API RESTful construida con Go, Gin Framework y GORM como parte del curso de introducción a microservicios. Se conecta a una base de datos PostgreSQL.

## Configuración de Base de Datos

El proyecto espera una base de datos PostgreSQL corriendo en localhost.

- **Host**: localhost
- **Puerto**: 5433
- **Usuario**: devuser
- **Contraseña**: devpassword123
- **Base de datos**: intro_microservicios

Las tablas `adults` y `children` se migrarán automáticamente al iniciar la aplicación.

## Modelos de Datos

### Adult / Child
- **id**: UUID (Generado automáticamente)
- **name**: String
- **last_name**: String
- **birth_year**: Int
- **image_url**: String

## Endpoints implementados

La API incluye los siguientes endpoints:

- **GET** `/Adults`: Obtiene la lista de todos los adultos registrados.
- **GET** `/Children`: Obtiene la lista de todos los niños registrados.
- **GET** `/Adults/{id}`: Obtiene un adulto por su ID.
- **GET** `/Children/{id}`: Obtiene un niño por su ID.
- **POST** `/Add/Adults`: Agrega un nuevo adulto a la base de datos.
- **POST** `/Add/Children`: Agrega un nuevo niño a la base de datos.

## Ejecución

Para ejecutar el proyecto:

1. Asegúrate de que la base de datos PostgreSQL esté corriendo y las credenciales sean correctas.
2. Instala las dependencias:
   ```bash
   go mod tidy
   ```
3. Ejecuta la aplicación:
   ```bash
   go run main.go
   ```

## Documentación API (Swagger)

La documentación interactiva de la API está disponible a través de Swagger.

1. Inicia el servidor.
2. Abre tu navegador y ve a:

<http://localhost:8080/swagger/index.html>

Allí podrás ver los ejemplos de requests y probar los endpoints.

