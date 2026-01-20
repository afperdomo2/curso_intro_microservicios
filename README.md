# Curso de introducción a microservicios

## Descripción

Este proyecto es una arquitectura de **microservicios** construida con Go, Gin Framework y GORM. Cada endpoint ha sido desacoplado en su propio servicio independiente y "containerizado" con Docker.

## Estructura del Proyecto

- `pkg/`: Paquetes compartidos (conexión DB, modelos).
- `services/`: Código fuente de cada microservicio.
- `docker-compose.yml`: Orquestación de contenedores.
- `intro_microservicios_collection.json`: Colección de Postman para probar la API.

## Configuración de Base de Datos

El proyecto incluye un contenedor de PostgreSQL configurado automáticamente en el `docker-compose.yml`.

- **Credenciales**: Usuario `devuser`, Contraseña `devpassword123`.
- **Database**: `intro_microservicios`.

## Ejecución con Docker Compose

La forma recomendada de levantar todo el entorno es usando Docker Compose. Esto iniciará la base de datos y todos los microservicios, mapeando sus puertos para acceso local.

```bash
docker-compose up --build
```

Una vez desplegado, los servicios estarán disponibles en los siguientes puertos:

| Servicio | Endpoint | Método | URL Local |
|----------|----------|--------|-----------|
| GetAdults | `/Adults` | GET | `http://localhost:8081/Adults` |
| GetChildren | `/Children` | GET | `http://localhost:8082/Children` |
| GetAdultById | `/Adults/:id` | GET | `http://localhost:8083/Adults/{id}` |
| GetChildById | `/Children/:id` | GET | `http://localhost:8084/Children/{id}` |
| AddAdult | `/Add/Adults` | POST | `http://localhost:8085/Add/Adults` |
| AddChild | `/Add/Children` | POST | `http://localhost:8086/Add/Children` |
| PickAge | `/PickAge` | GET | `http://localhost:8087/PickAge` |
| AddMember | `/Add/Member` | POST | `http://localhost:8088/Add/Member` |

## Pruebas (Postman)

Se ha eliminado la documentación de Swagger en favor de Postman.

1. Abre Postman.
2. Importa el archivo `intro_microservicios_collection.json` ubicado en la raíz del proyecto.
3. Ejecuta las peticiones directamente contra el entorno local desplegado con Docker.
