![logo]([https://i3.wp.com/raw.githubusercontent.com/Quamagi/GoApiRestFull/main/logo.jpg](https://raw.githubusercontent.com/Quamagi/GoApiRestFull/main/logo.jpg?token=GHSAT0AAAAAACSTN3V4AHGEMGG2M3XGLB6YZTE72ZA))
# GoApiRestFull

```markdown
# API REST en Go con autenticación JWT

Esta es una API REST construida en Go con autenticación utilizando JSON Web Tokens (JWT). La API permite realizar operaciones CRUD (Crear, Leer, Actualizar y Eliminar) en usuarios, y también incluye una función de inicio de sesión para autenticar a los usuarios.

## Características

- Gestión de usuarios (Crear, Leer, Actualizar y Eliminar)
- Autenticación de usuarios mediante JWT
- Almacenamiento de datos en una base de datos SQLite
- Uso del paquete `gorm.io/gorm` para interactuar con la base de datos
- Enrutamiento con el paquete `gorilla/mux`

## Requisitos

- Go 1.16 o superior
- SQLite

## Instalación

1. Clona este repositorio en tu máquina local:

```
git clone https://github.com/tu-usuario/api-rest-go.git
```

2. Navega al directorio del proyecto:

```
cd api-rest-go
```

3. Instala las dependencias del proyecto:

```
go get ./...
```

4. Ejecuta la aplicación:

```
go run main.go
```

La aplicación estará disponible en `http://localhost:8080`.

## Uso

La API proporciona los siguientes endpoints:

- `POST /users`: Crea un nuevo usuario
- `GET /users`: Obtiene una lista de todos los usuarios (requiere autenticación)
- `GET /users/:id`: Obtiene un usuario por su ID (requiere autenticación)
- `PUT /users/:id`: Actualiza un usuario por su ID (requiere autenticación)
- `DELETE /users/:id`: Elimina un usuario por su ID (requiere autenticación)
- `POST /login`: Inicia sesión y obtiene un token JWT

Para las solicitudes que requieren autenticación, debes incluir el token JWT en el encabezado `Authorization` de la siguiente manera:

```
Authorization: Bearer <token_jwt>
```

## Licencia

Este proyecto está bajo la Licencia [MIT](https://opensource.org/licenses/MIT). La Licencia MIT es una de las licencias de código abierto más permisivas y populares. Permite el uso, modificación y distribución del software de forma gratuita, con la única condición de incluir la notificación de derechos de autor y la exención de responsabilidad en todas las copias o porciones sustanciales del software.

Esta licencia es adecuada para proyectos de código abierto que no tienen restricciones específicas y que permiten el uso comercial, la modificación y la distribución del código fuente sin requerimientos adicionales.

```

Esta documentación en formato Markdown cubre los aspectos principales del proyecto, incluyendo una breve descripción, características, requisitos, instrucciones de instalación y uso, y la licencia recomendada (MIT). Puedes personalizar y ajustar el contenido según tus necesidades específicas.
