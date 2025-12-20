# 🚀 ESTSoftware - API REST

API REST desarrollada con Go (Golang) para la creacion de cursos.

## 📋 Tabla de Contenidos

- [Características](#características)
- [Tecnologías](#tecnologías)
- [Requisitos Previos](#requisitos-previos)
- [Instalación](#instalación)
- [Configuración](#configuración)
- [Uso](#uso)
- [Estructura del Proyecto](#estructura-del-proyecto)
- [API Endpoints](#api-endpoints)
- [Testing](#testing)
- [Deployment](#deployment)
- [Contribuir](#contribuir)
- [Licencia](#licencia)

## ✨ Características

- ✅ API RESTful con Go
- ✅ Arquitectura limpia (Clean Architecture, Vertical Slice, Hexagonal)
- ✅ Autenticación JWT
- ✅ Integración con Cloudinary
- ✅ Dockerizado
- ✅ CI/CD con GitHub Actions
- ✅ Tests automatizados

## 🛠️ Tecnologías

- **Go** 1.24+
- **Gin** - Framework web
- **JWT** - Autenticación
- **Cloudinary** - Gestión de imágenes
- **Docker** - Containerización
- **GitHub Actions** - CI/CD

## 📦 Requisitos Previos

Antes de comenzar, asegúrate de tener instalado:

- [Go](https://golang.org/dl/) 1.24 o superior
- [Docker](https://www.docker.com/get-started) (opcional)
- [Git](https://git-scm.com/)

## 🚀 Instalación

### 1. Clonar el repositorio
```bash
git clone https://github.com/Ameth-Toledo/ESTSoftwareOficial.git
cd ESTSoftwareOficial
```

### 2. Instalar dependencias
```bash
go mod download
```

### 3. Configurar variables de entorno
```bash
cp .env.example .env
# Edita .env con tus credenciales
```

### 4. Ejecutar la aplicación
```bash
go run main.go
```

La API estará disponible en `http://localhost:8080`

## ⚙️ Configuración

Crea un archivo `.env` basado en `.env.example`:
```env
# Credenciales de BD
DB_HOST=solicitar_credenciales a QA
DB_PORT=solicitar_credenciales a QA
DB_NAME=solicitar_credenciales a QA
DB_USER=solicitar_credenciales a QA
DB_PASSWORD=solicitar_credenciales a QA
DB_SSL=false

# JWT Configuration
JWT_SECRET=solicitar_credenciales a QA

# Cloudinary Configuration
CLOUDINARY_NAME=solicitar_credenciales a QA
API_KEY=solicitar_credenciales a QA
API_SECRET=solicitar_credenciales a QA

# Frontend URL 
FRONTEND_URL=

# Google OAuth Configuration
GOOGLE_CLIENT_ID=solicitar_client_id a QA
GOOGLE_CLIENT_SECRET=solicitar_client_secret a QA
GOOGLE_REDIRECT_URL=solicitar_redirect_url a QA

# GitHub OAuth Configuration
GITHUB_CLIENT_ID=solicitar_client_id a QA
GITHUB_CLIENT_SECRET=solicitar_client_secret a QA
GITHUB_REDIRECT_URL=solicitar_redirect_url a QA
```

## 🎯 Uso

### Con Go
```bash
# Desarrollo
go run main.go

# Build
go build -o app
./app
```

### Con Docker
```bash
# Build
docker build -t estsoftware .

# Run
docker run -p 8080:8080 --env-file .env estsoftware
```

### Con Docker Compose
```bash
docker-compose up -d
```

## 📁 Estructura del Proyecto
```
estsoftwareoficial/
├── .github/
│   └── workflows/
│       └── go-ci.yml          # GitHub Actions CI/CD
├── src/
│   ├── core/
│   │   ├── cloudinary/        # Servicio de Cloudinary
│   │   └── security/          # Auth, JWT, Hash
│   └── users/
│       ├── application/       # Casos de uso
│       ├── domain/           # Lógica de negocio
│       │   ├── dto/          # Data Transfer Objects
│       │   └── entities/     # Entidades del dominio
│       └── infrastructure/   # Implementaciones técnicas
│           ├── adapters/     # Adaptadores externos
│           ├── controllers/  # Controladores HTTP
│           └── routes/       # Rutas
├── .env.example              # Template de variables de entorno
├── .gitignore
├── Dockerfile
├── go.mod
├── go.sum
├── main.go
├── CONTRIBUTING.md           # Guía de contribución
└── README.md
```

## 🔌 API Endpoints

### Autenticación
```http
POST   /api/v1/auth/register    # Registrar usuario
POST   /api/v1/auth/login       # Iniciar sesión
POST   /api/v1/auth/refresh     # Refrescar token
```

### Usuarios
```http
GET    /api/v1/users            # Listar usuarios
GET    /api/v1/users/:id        # Obtener usuario
POST   /api/v1/users            # Crear usuario
PUT    /api/v1/users/:id        # Actualizar usuario
DELETE /api/v1/users/:id        # Eliminar usuario
```

### Ejemplo de uso
```bash
# Registro
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "securepassword"
  }'

# Login
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "securepassword"
  }'
```

## 🧪 Testing

### Ejecutar tests
```bash
# Todos los tests
go test ./...

# Con cobertura
go test -cover ./...

# Con detalles
go test -v ./...

# Cobertura detallada
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

### Tests automáticos (CI/CD)

Los tests se ejecutan automáticamente en cada Pull Request mediante GitHub Actions.

## 🚢 Deployment

### Requisitos

- Servidor con Docker instalado
- Variables de entorno configuradas

### Pasos

1. **Build de la imagen:**
```bash
   docker build -t estsoftware:latest .
```

2. **Subir a registry (opcional):**
```bash
   docker tag estsoftware:latest registry.example.com/estsoftware:latest
   docker push registry.example.com/estsoftware:latest
```

3. **Deploy:**
```bash
   docker run -d \
     -p 8080:8080 \
     --env-file .env \
     --name estsoftware \
     estsoftware:latest
```

## 🤝 Contribuir

Por favor, lee [CONTRIBUTING.md](CONTRIBUTING.md) para conocer el proceso.

### Flujo de trabajo

1. Fork el proyecto
2. Crea una rama feature, ejemplo: (`git checkout -b feature/amethdev-nueva-funcionalidad`)
3. Commit tus cambios, agregar el prefijo de commit, ejemplo: (`git commit -m 'feat: add nueva funcionalidad'`)
4. Push a la rama (`git push origin feature/amethdev-nueva-funcionalidad`)
5. Abre un Pull Request

### Commits

Usamos [Conventional Commits](https://www.conventionalcommits.org/):

- `feat:` Nueva funcionalidad
- `fix:` Corrección de bugs
- `docs:` Documentación
- `test:` Tests
- `refactor:` Refactorización
- `chore:` Mantenimiento

## 👥 Equipo

- **Ameth Toledo** - *Developer* - [@Ameth-Toledo](https://github.com/Ameth-Toledo)
- **Fabricio Perez** - *Developer* - [@FabricioPRZ](https://github.com/FabricioPRZ)
- **Eddy Jordan** - *Developer* - [@JORED666](https://github.com/JORED666)
- **Ivan** - *Developer* - [@ivanGG23](https://github.com/ivanGG23)

## 📄 Licencia

Este proyecto está bajo la Licencia MIT - ver el archivo [LICENSE](LICENSE) para más detalles.

## 📞 Contacto

- **Email:** shakerzest@gmail.com
- **GitHub:** [@Ameth-Toledo](https://github.com/Ameth-Toledo)
- **LinkedIn:** [Ameth Toledo](https://www.linkedin.com/in/ameth-de-jes%C3%BAs-m%C3%A9ndez-toledo/)

---

⭐️ Si te gusta este proyecto, dale una estrella en GitHub!
