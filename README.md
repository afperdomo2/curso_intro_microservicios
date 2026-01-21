# 🚀 Curso de introducción a microservicios

## 📋 Descripción

Este proyecto es una arquitectura de **microservicios** construida con Go, Gin Framework y GORM. Cada endpoint ha sido desacoplado en su propio servicio independiente y "containerizado" con Docker.

## 📂 Estructura del Proyecto

- 📦 `pkg/`: Paquetes compartidos (conexión DB, modelos).
- 🔧 `services/`: Código fuente de cada microservicio.
- 🐳 `docker-compose.yml`: Orquestación de contenedores.
- 📮 `intro_microservicios_collection.json`: Colección de Postman para probar la API.

## 🗄️ Configuración de Base de Datos

El proyecto incluye un contenedor de PostgreSQL configurado automáticamente en el `docker-compose.yml`.

- 👤 **Credenciales**: Usuario `devuser`, Contraseña `devpassword123`.
- 📊 **Database**: `intro_microservicios`.

## 🐳 Ejecución con Docker Compose

La forma recomendada de levantar todo el entorno es usando Docker Compose. Esto iniciará la base de datos y todos los microservicios, mapeando sus puertos para acceso local.

```bash
docker-compose up --build
```

Una vez desplegado, los servicios estarán disponibles a través del **API Gateway (Nginx)** en el puerto **8000**. Ya no es necesario acceder a puertos individuales.

| Servicio | Ruta (Gateway) | Método | URL Local |
|----------|----------------|--------|-----------|
| GetAdults | `/Adults` | GET | `http://localhost:8000/Adults` |
| GetChildren | `/Children` | GET | `http://localhost:8000/Children` |
| GetAdultById | `/Adults/:id` | GET | `http://localhost:8000/Adults/{id}` |
| GetChildById | `/Children/:id` | GET | `http://localhost:8000/Children/{id}` |
| AddAdult | `/Add/Adults` | POST | `http://localhost:8000/Add/Adults` |
| AddChild | `/Add/Children` | POST | `http://localhost:8000/Add/Children` |
| PickAge | `/PickAge` | GET | `http://localhost:8000/PickAge` |
| AddMember | `/Add/Member` | POST | `http://localhost:8000/Add/Member` |

## 📨 Servicio AddMember (Kafka)

El servicio `add-member` implementa integración con **Apache Kafka** para publicar eventos de miembros agregados.

### ✨ Características

- 📤 **Productor de Kafka**: Envía mensajes del nuevo miembro al topic `pickage`.
- 💉 **Inyección de Dependencias**: El broker de Kafka se inyecta desde la configuración.
- ⚙️ **Gestión de Configuración**: Carga variables desde `.env` usando `godotenv`.
- 🎯 **Manejo de Tópics**: Crea automáticamente el topic si no existe.

### ⚙️ Configuración

El servicio requiere la variable de entorno `KAFKA_BROKER`:

```bash
KAFKA_BROKER=kafka:9092  # Para desarrollo con Docker
KAFKA_BROKER=localhost:9093  # Para cliente externo
```

En Docker Compose, esta variable se define automáticamente a través de `KAFKA_BROKER: kafka:9092`.

### 📁 Estructura del Servicio

```
services/add-member/
├── cmd/main.go                  # Punto de entrada principal
├── config/config.go             # Gestión de configuración
├── handlers/add_member_handler.go # Lógica de negocio
├── kafka/producer.go            # Productor de Kafka
├── models/member.go             # Modelo de datos
├── .env.example                 # Plantilla de variables
└── Dockerfile                   # Configuración de contenedor
```

### 🔧 Optimizaciones de Kafka

**🧪 Para Pruebas**:

- ⏱️ Retención: 24 horas
- 🔄 Auto-creación de topics: Habilitada
- 🔗 Sincronización: RequireOne (solo el líder confirma)

**🏢 Para Producción** (8GB RAM, 15GB storage):

- ⏱️ Retención: 72 horas (3 días)
- 💾 Almacenamiento: Máximo 10GB
- 🔗 Sincronización: RequireAll (todas las replicas confirman)

Ver configuración en `docker-compose.yml` sección Kafka.

## 🧪 Pruebas

### 📮 Opción 1: Postman

1. Abre Postman.
2. Importa el archivo `intro_microservicios_collection.json` ubicado en la raíz del proyecto.
3. Ejecuta las peticiones directamente contra el entorno local desplegado con Docker.

### 💻 Opción 2: VS Code REST Client

Si utilizas la extensión **REST Client** en VS Code, puedes ejecutar las peticiones directamente desde el editor:

1. Abre el archivo `requests.http`.
2. Haz clic en "Send Request" sobre cada definición de endpoint.
