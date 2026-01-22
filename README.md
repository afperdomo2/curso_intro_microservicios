# 🚀 Curso de introducción a microservicios

## 📋 Descripción

Este proyecto es una arquitectura de **microservicios event-driven** construida con Go, Gin Framework, Apache Kafka y GORM. La arquitectura implementa patrones SOLID y separa cada responsabilidad en su propio servicio independiente y containerizado con Docker.

**Modelo Arquitectónico:**

- ✅ Microservicios desacoplados
- ✅ Event-driven con Apache Kafka
- ✅ API Gateway (Nginx)
- ✅ Base de datos PostgreSQL
- ✅ Arquitectura SOLID (Single Responsibility Principle)

## 📂 Estructura del Proyecto

```
proyecto/
├── pkg/                          # Paquetes compartidos
│   ├── database/                # Conexión a BD
│   └── models/                  # Modelos compartidos
├── services/                     # Microservicios independientes
│   ├── add-member/              # Punto de entrada HTTP (Kafka Producer)
│   ├── pick-age/                # Clasificador de edad (Kafka Consumer → Producer)
│   ├── add-adult/               # Persiste adultos (Kafka Consumer)
│   ├── add-child/               # Persiste menores (Kafka Consumer)
│   ├── get-adults/              # Consulta adultos (HTTP)
│   ├── get-children/            # Consulta menores (HTTP)
│   ├── get-adult-by-id/         # Consulta adulto por ID (HTTP)
│   └── get-child-by-id/         # Consulta menor por ID (HTTP)
├── nginx/                        # API Gateway
├── docker-compose.yml            # Orquestación de servicios
├── requests.http                 # Requests de prueba (VS Code)
└── intro_microservicios_collection.json  # Colección Postman
```

## 🗄️ Configuración de Base de Datos

El proyecto incluye un contenedor de PostgreSQL configurado automáticamente en el `docker-compose.yml`.

- 👤 **Credenciales**: Usuario `devuser`, Contraseña `devpassword123`
- 📊 **Database**: `intro_microservicios`
- 🏠 **Host**: `postgres:5432` (en Docker)

## 🐳 Ejecución con Docker Compose

La forma recomendada de levantar todo el entorno es usando Docker Compose:

```bash
docker-compose up --build
```

Esto iniciará:

- 1x PostgreSQL (puerto 5433)
- 1x Apache Kafka (puerto 9092)
- 1x Nginx API Gateway (puerto 8000)
- 8x Microservicios Go

Una vez desplegado, los servicios HTTP estarán disponibles a través del **API Gateway (Nginx)** en el puerto **8000**.

## 📡 Endpoints Disponibles

| Servicio | Ruta (Gateway) | Método | Descripción |
|----------|----------------|--------|-------------|
| **📖 GET** | | | |
| GetAdults | `/Adults` | GET | Retorna lista de adultos |
| GetChildren | `/Children` | GET | Retorna lista de menores |
| GetAdultById | `/Adults/:id` | GET | Retorna adulto por UUID |
| GetChildById | `/Children/:id` | GET | Retorna menor por UUID |
| **📝 POST** | | | |
| AddMember | `/Add/Member` | POST | ⭐ Inicia flujo Kafka |

### 🔴 Endpoints Deprecados

Los siguientes endpoints **ya no existen** y han sido reemplazados por el flujo event-driven:

```
❌ POST /Add/Adults      → Use POST /Add/Member (event-driven)
❌ POST /Add/Children    → Use POST /Add/Member (event-driven)
❌ GET  /PickAge         → Servicio internal (sin HTTP)
```

## 🧪 Pruebas

### 📮 Opción 1: Postman

1. Importa el archivo [intro_microservicios_collection.json](intro_microservicios_collection.json)
2. Ejecuta las requests contra `http://localhost:8000`

### 💻 Opción 2: VS Code REST Client

1. Abre el archivo [requests.http](requests.http)
2. Haz clic en "Send Request" sobre cada endpoint

### 📝 Ejemplo de Request

```bash
curl -X POST http://localhost:8000/Add/Member \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Carlos",
    "last_name": "Rodriguez",
    "birth_year": 1990,
    "image_url": "https://example.com/carlos.jpg"
  }'
```

---

## 🏗️ Arquitectura Event-Driven con Kafka

El sistema implementa un flujo completamente **event-driven** donde cada microservicio tiene una responsabilidad única y se comunica mediante Apache Kafka:

### 📊 Diagrama de Flujo

```
┌─────────────────────────────────────────────────────────────────┐
│                       Cliente HTTP                              │
└────────────────────────────┬────────────────────────────────────┘
                             │
                    POST /Add/Member
                             │
         ┌───────────────────▼────────────────────┐
         │   🟢 ADD-MEMBER (Kafka Producer)      │
         │                                        │
         │  - Handler HTTP                        │
         │  - Valida datos                        │
         │  - Produce evento a Kafka              │
         └───────────────────┬────────────────────┘
                             │
            Topic: members.registration.fct.member.received
            {name, last_name, birth_year, image_url, timestamp}
                             │
         ┌───────────────────▼────────────────────┐
         │  🟡 PICK-AGE (Consumer → Producer)    │
         │                                        │
         │  - Consume evento de miembro           │
         │  - Classifier: Calcula edad            │
         │  - Produce a 2 topics según edad       │
         └───────┬────────────────────┬───────────┘
                 │                    │
        edad >= 18│                   │edad < 18
                 │                    │
    members.classification.fct. │  members.classification.fct.
    adult.validated             │  child.validated
                 │                    │
    ┌────────────▼──────────┐  ┌──────▼──────────────┐
    │ 🔵 ADD-ADULT         │  │ 🟣 ADD-CHILD       │
    │ (Kafka Consumer)     │  │ (Kafka Consumer)   │
    │                      │  │                    │
    │ - Repository BD      │  │ - Repository BD    │
    │ - Tabla: adults      │  │ - Tabla: children  │
    └──────────────────────┘  └────────────────────┘
```

### 📨 Servicio AddMember (HTTP Producer)

**Única entrada HTTP al sistema.** Recibe solicitudes de clientes y las publica a Kafka.

**Endpoints HTTP:**

- `POST /Add/Member` → Publica evento de miembro registrado

**Estructura SOLID:**

- `handlers/add_member_handler.go`: Lógica de negocio (SRP)
- `kafka/producer.go`: Publicación a Kafka (SRP)
- `config/config.go`: Configuración (SRP)
- `models/member.go`: Modelo de datos (SRP)

**Topics producidos:**

- `members.registration.fct.member.received`

**Payload de entrada:**

```json
{
  "name": "string (required)",
  "last_name": "string (required)",
  "birth_year": "integer (required)",
  "image_url": "string (optional)"
}
```

**Payload en Kafka:**

```json
{
  "name": "string",
  "last_name": "string",
  "birth_year": "integer",
  "image_url": "string",
  "timestamp": "RFC3339"
}
```

---

### 🔄 Servicio PickAge (Kafka Consumer → Producer)

**Cerebro de la clasificación.** Consume miembros registrados, calcula edad y los clasifica por tipo.

**Arquitectura SOLID:**

- `classifier/classifier.go`: Lógica de clasificación (SRP)
  - `Classify()`: Determina si es adulto (≥18) o menor (<18)
- `kafka/consumer.go`: Lectura de eventos (SRP)
  - `Start()`: Escucha continuamente el topic
- `kafka/producer.go`: Publicación de eventos clasificados (SRP)
  - `PublishClassification()`: Envía a topic específico según clasificación

**Topics consumidos:**

- `members.registration.fct.member.received`

**Topics producidos:**

- `members.classification.fct.adult.validated` (edad >= 18)
- `members.classification.fct.child.validated` (edad < 18)

**GroupID:** `pick-age-service`

**Payload de salida (Adultos):**

```json
{
  "name": "string",
  "last_name": "string",
  "birth_year": "integer",
  "image_url": "string",
  "age": "integer",
  "published_at": "RFC3339"
}
```

---

### 👤 Servicio AddAdult (Kafka Consumer)

**Persiste adultos en la base de datos.** Consumer puro sin endpoint HTTP.

**Arquitectura SOLID:**

- `repository/adult_repository.go`: Acceso a datos (SRP)
  - `SaveAdult()`: Inserta adulto en tabla `adults`
- `kafka/consumer.go`: Lectura de eventos (SRP)
  - `Start()`: Escucha topic de adultos clasificados
  - `processAdult()`: Orquesta guardado en BD
- `config/config.go`: Configuración (SRP)

**Topics consumidos:**

- `members.classification.fct.adult.validated`

**Base de datos:**

- Tabla: `adults`
- Campos: `id` (UUID), `name`, `last_name`, `birth_year`, `image_url`

**GroupID:** `add-adult-service`

---

### 👶 Servicio AddChild (Kafka Consumer)

**Persiste menores en la base de datos.** Consumer puro sin endpoint HTTP.

**Arquitectura SOLID:**

- `repository/child_repository.go`: Acceso a datos (SRP)
  - `SaveChild()`: Inserta menor en tabla `children`
- `kafka/consumer.go`: Lectura de eventos (SRP)
  - `Start()`: Escucha topic de menores clasificados
  - `processChild()`: Orquesta guardado en BD
- `config/config.go`: Configuración (SRP)

**Topics consumidos:**

- `members.classification.fct.child.validated`

**Base de datos:**

- Tabla: `children`
- Campos: `id` (UUID), `name`, `last_name`, `birth_year`, `image_url`

**GroupID:** `add-child-service`

---

### 📖 Servicios de Consulta (HTTP Readers)

Los servicios GET son **read-only** sin lógica de evento:

- `GetAdults`: Consulta tabla `adults`
- `GetChildren`: Consulta tabla `children`
- `GetAdultById`: Consulta adulto por UUID
- `GetChildById`: Consulta menor por UUID

**Protocolo:** HTTP REST puro
**Método:** GET
**Autenticación:** Ninguna (desarrollo)

---

## 🧠 Principios SOLID Aplicados

### Single Responsibility Principle (SRP)

Cada archivo tiene una única responsabilidad:

- `config/config.go`: Solo configuración
- `kafka/consumer.go`: Solo lectura de Kafka
- `kafka/producer.go`: Solo escritura a Kafka
- `repository/repository.go`: Solo acceso a datos
- `classifier/classifier.go`: Solo lógica de negocio

### Open/Closed Principle (OCP)

Fácil extender sin modificar:

- Agregar nuevas clasificaciones en `Classifier`
- Agregar nuevos servicios consumers

### Dependency Injection (DI)

Inyección de dependencias explícita:

- `NewConsumer(topic, brokerAddr, repo)`
- `NewProducer(brokerAddr)`
- `NewAddMemberHandler(kafkaProducer)`

---

## 🔧 Configuración de Kafka

### 🧪 Para Desarrollo (Docker Compose)

- **Brokers:** 1
- **Particiones:** 1 por topic
- **Replication Factor:** 1
- **Retención:** 24 horas
- **Auto-creación de topics:** Habilitada
- **Sincronización:** RequireOne (solo líder)

```bash
docker-compose up
```

### 🏢 Para Producción

Se recomienda ajustar en `docker-compose.yml`:

- **Brokers:** 3+
- **Particiones:** 3+ (paralelismo)
- **Replication Factor:** 3 (tolerancia a fallos)
- **Retención:** 72 horas (3 días)
- **Auto-creación de topics:** Deshabilitada
- **Sincronización:** RequireAll (todas las replicas)

---

## 📊 Flujo Completo - Ejemplo Práctico

**Request:** Agregar nuevo miembro

```bash
curl -X POST http://localhost:8000/Add/Member \
  -H "Content-Type: application/json" \
  -d '{"name":"Maria","last_name":"Garcia","birth_year":2005,"image_url":"https://example.com/maria.jpg"}'
```

**Paso 1:** AddMember Handler valida y produce

```
HTTP 200 OK
Response: {"message":"Miembro Maria Garcia agregado correctamente",...}
```

**Paso 2:** PickAge consume y clasifica

```
Calcula: 2026 - 2005 = 21 años
Clasifica: ADULTO (>= 18)
Produce a: members.classification.fct.adult.validated
```

**Paso 3:** AddAdult consume y persiste

```
Consumer recibe: {name:"Maria", last_name:"Garcia", birth_year:2005, age:21, ...}
INSERT INTO adults (id, name, last_name, birth_year, image_url)
VALUES (uuid, 'Maria', 'Garcia', 2005, 'https://example.com/maria.jpg')
```

**Paso 4:** Consultar en GetAdults

```bash
curl http://localhost:8000/Adults
# Retorna: [{"id":"uuid...","name":"Maria","last_name":"Garcia",...}]
```

---

## 🚀 Ejecución Completa

### 1. Iniciar servicios

```bash
cd curso_intro_microservicios
docker-compose up --build
```

### 2. Verificar que todos están corriendo

```bash
docker-compose ps
```

### 3. Hacer un POST a Add/Member

```bash
# Opción A: cURL
curl -X POST http://localhost:8000/Add/Member \
  -H "Content-Type: application/json" \
  -d '{"name":"Test","last_name":"User","birth_year":1990,"image_url":"https://example.com/test.jpg"}'

# Opción B: Postman (importa intro_microservicios_collection.json)
# Opción C: VS Code REST Client (abre requests.http)
```

### 4. Verificar en logs

```bash
docker-compose logs -f pick_age    # Ver clasificación
docker-compose logs -f add_adult   # Ver persistencia
docker-compose logs -f add_child   # Ver persistencia
```

### 5. Consultar datos

```bash
curl http://localhost:8000/Adults   # Listar adultos
curl http://localhost:8000/Children # Listar menores
```

### 6. Detener servicios

```bash
docker-compose down
```

---

## 📝 Notas Importantes

- **Sin autenticación:** Este es un proyecto educativo sin seguridad
- **Modo desarrollo:** Las configuraciones están optimizadas para desarrollo local
- **Kafka debe estar running:** Sin Kafka, los consumers fallarán
- **PostgreSQL debe estar running:** Sin BD, no se pueden guardar datos
- **Nginx como gateway:** Todos los requests HTTP van a través de Nginx en puerto 8000
