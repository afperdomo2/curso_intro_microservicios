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
| AddMember | `/Add/Member` | POST | `http://localhost:8000/Add/Member` |
| AddAdult | N/A (Kafka Consumer) | - | Event-driven desde PickAge |
| AddChild | N/A (Kafka Consumer) | - | Event-driven desde PickAge |
| PickAge | N/A (Kafka Consumer/Producer) | - | Event-driven desde AddMember |

> **Nota**: Los servicios `AddAdult`, `AddChild` y `PickAge` son **Kafka Consumers** sin endpoint HTTP. Funcionan de forma event-driven en el flujo de mensajería.

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

## 📊 Servicio PickAge (Kafka Consumer Simple)

El servicio `pick-age` es un **consumidor de Kafka** que escucha el topic `pickage` y loguea si los miembros son adultos (18+) o menores de edad.

### 🔄 Flujo de Proceso

1. **Recibe**: Escucha mensajes del topic `pickage` (publicados por `add-member`)
2. **Procesa**: Calcula edad basándose en el año de nacimiento actual
3. **Loguea**: 
   - `👤 ADULTO: [Nombre] [Apellido] - Nacido en [Año] (edad: [Años] años)` si tiene 18+
   - `👶 MENOR: [Nombre] [Apellido] - Nacido en [Año] (edad: [Años] años)` si es menor

### 📁 Estructura del Servicio

```
services/pick-age/
├── cmd/main.go              # Punto de entrada
├── config/config.go         # Gestión de configuración
├── kafka/consumer.go        # Consumer que procesa directamente
├── models/member.go         # Modelo de datos
├── .env.example             # Plantilla de variables
└── Dockerfile               # Configuración de contenedor
```

### ✨ Características

- 👂 **Consumer de Kafka**: Escucha el topic `pickage` continuamente
- 🔍 **Análisis de Edad**: Calcula edad en tiempo real
- 📝 **Logging Simple**: Loguea adultos vs menores
- 🎯 **Sin Complejidades**: No usa producers, handlers, ni base de datos
- 🚫 **Sin HTTP**: Puramente event-driven

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

---

## 🏗️ Flujo de Microservicios con Kafka

El sistema implementa un flujo event-driven utilizando Apache Kafka como bus de mensajería:

### 📊 Diagrama de Flujo

```
[Cliente HTTP]
      ↓
[POST /Add/Member] → add-member service
      ↓
  [Produce] → Topic: "members.registration.fct.member.received"
      ↓
[pick-age service]
      ├─→ Classifica por edad
      ├─→ Si edad >= 18: [Produce] → "members.classification.fct.adult.validated"
      └─→ Si edad < 18: [Produce] → "members.classification.fct.child.validated"
      ↓
[Consumers]
├─→ add-adult service (consume "members.classification.fct.adult.validated")
│   └─→ Guarda en tabla "adults"
│
└─→ add-child service (consume "members.classification.fct.child.validated")
    └─→ Guarda en tabla "children"
```

### 📨 Servicio AddMember (Kafka Producer)

El servicio `add-member` sigue siendo el punto de entrada HTTP. Recibe un miembro y lo publica al topic `members.registration.fct.member.received`.

**Endpoints:**
- `POST /Add/Member` → Publica evento de miembro registrado

**Topics producidos:**
- `members.registration.fct.member.received`

---

### 🔄 Servicio PickAge (Kafka Consumer → Producer)

El servicio `pick-age` consume miembros registrados, calcula su edad y los clasifica.

**Arquitectura SOLID:**
- `classifier/classifier.go`: Lógica de clasificación (SRP)
- `kafka/consumer.go`: Lectura de eventos (SRP)
- `kafka/producer.go`: Publicación de eventos clasificados (SRP)

**Topics consumidos:**
- `members.registration.fct.member.received`

**Topics producidos:**
- `members.classification.fct.adult.validated` (edad >= 18)
- `members.classification.fct.child.validated` (edad < 18)

**GroupID:** `pick-age-service`

---

### 👤 Servicio AddAdult (Kafka Consumer)

El servicio `add-adult` consume adultos clasificados y los guarda en la base de datos.

**Ya NO tiene endpoint HTTP** - Es puramente event-driven.

**Arquitectura SOLID:**
- `repository/adult_repository.go`: Acceso a datos (SRP)
- `kafka/consumer.go`: Lectura de eventos (SRP)
- `config/config.go`: Gestión de configuración (SRP)

**Topics consumidos:**
- `members.classification.fct.adult.validated`

**Base de datos:**
- Tabla: `adults` (crea automáticamente adultos)

**GroupID:** `add-adult-service`

---

### 👶 Servicio AddChild (Kafka Consumer)

Similar a `add-adult`, este servicio consume menores clasificados y los guarda en la base de datos.

**Ya NO tiene endpoint HTTP** - Es puramente event-driven.

**Arquitectura SOLID:**
- `repository/child_repository.go`: Acceso a datos (SRP)
- `kafka/consumer.go`: Lectura de eventos (SRP)
- `config/config.go`: Gestión de configuración (SRP)

**Topics consumidos:**
- `members.classification.fct.child.validated`

**Base de datos:**
- Tabla: `children` (crea automáticamente menores)

**GroupID:** `add-child-service`
