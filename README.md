![logo](https://raw.githubusercontent.com/Quamagi/GoApiRestFull/main/logo.jpg?token=GHSAT0AAAAAACSTN3V4AHGEMGG2M3XGLB6YZTE72ZA)

```markdown
# GoApiRestFull

## API REST en Go con autenticación JWT

Esta es una API REST construida en Go con autenticación utilizando JSON Web Tokens (JWT). La API permite realizar operaciones CRUD (Crear, Leer, Actualizar y Eliminar) en usuarios, e incluye una función de inicio de sesión para autenticar a los usuarios.

## Características

- Gestión de usuarios (Crear, Leer, Actualizar y Eliminar)
- Autenticación de usuarios mediante JWT
- Paginación de usuarios
- Almacenamiento de datos en una base de datos SQLite
- Uso del paquete `gorm.io/gorm` para interactuar con la base de datos
- Enrutamiento con el paquete `gorilla/mux`

## Requisitos

- Go 1.16 o superior
- SQLite

## Instalación

1. Clona este repositorio en tu máquina local:

    ```bash
    git clone https://github.com/tu-usuario/api-rest-go.git
    ```

2. Navega al directorio del proyecto:

    ```bash
    cd api-rest-go
    ```

3. Instala las dependencias del proyecto:

    ```bash
    go get ./...
    ```

4. Ejecuta la aplicación:

    ```bash
    go run main.go
    ```

    La aplicación estará disponible en `http://localhost:8080`.

## Uso

### Endpoints

La API proporciona los siguientes endpoints:

- `POST /users`: Crea un nuevo usuario
- `GET /users`: Obtiene una lista de todos los usuarios (requiere autenticación)
- `GET /users/:id`: Obtiene un usuario por su ID (requiere autenticación)
- `PUT /users/:id`: Actualiza un usuario por su ID (requiere autenticación)
- `DELETE /users/:id`: Elimina un usuario por su ID (requiere autenticación)
- `POST /login`: Inicia sesión y obtiene un token JWT
- `GET /paginate/users`: Obtiene una lista paginada de usuarios (requiere autenticación)

### Paginación de Usuarios

Para obtener una lista paginada de usuarios, usa los parámetros `limit` y `cursor` en la solicitud `GET /paginate/users`. Por ejemplo:

```bash
curl -X GET "http://localhost:8080/paginate/users?limit=10&cursor=1" -H "Authorization: Bearer <token_jwt>"
```

### Ejemplo de solicitudes

**Crear un usuario**

```bash
curl -X POST "http://localhost:8080/users" -H "Content-Type: application/json" -d '{
    "name": "John Doe",
    "email": "johndoe@example.com",
    "password": "password123"
}'
```

**Iniciar sesión**

```bash
curl -X POST "http://localhost:8080/login" -H "Content-Type: application/json" -d '{
    "email": "johndoe@example.com",
    "password": "password123"
}'
```

**Obtener lista de usuarios**

```bash
curl -X GET "http://localhost:8080/users" -H "Authorization: Bearer <token_jwt>"
```

**Obtener lista paginada de usuarios**

```bash
curl -X GET "http://localhost:8080/paginate/users?limit=10&cursor=1" -H "Authorization: Bearer <token_jwt>"
```

**Obtener un usuario por ID**

```bash
curl -X GET "http://localhost:8080/users/1" -H "Authorization: Bearer <token_jwt>"
```

**Actualizar un usuario por ID**

```bash
curl -X PUT "http://localhost:8080/users/1" -H "Content-Type: application/json" -H "Authorization: Bearer <token_jwt>" -d '{
    "name": "John Doe Updated",
    "email": "johnupdated@example.com",
    "password": "newpassword123"
}'
```

**Eliminar un usuario por ID**

```bash
curl -X DELETE "http://localhost:8080/users/1" -H "Authorization: Bearer <token_jwt>"
```

## Contribución

Si deseas contribuir a este proyecto, por favor realiza un fork del repositorio y envía un pull request con tus cambios.

## Licencia

Este proyecto está bajo la Licencia MIT. Consulta el archivo LICENSE para más detalles.

## Apóyame en Ko-fi

Si encuentras útil este proyecto y deseas apoyarme, puedes hacerlo a través de Ko-fi:

```html
<script src='https://storage.ko-fi.com/cdn/scripts/overlay-widget.js'></script>
<script>
  kofiWidgetOverlay.draw('quamagi', {
    'type': 'floating-chat',
    'floating-chat.donateButton.text': 'Support me',
    'floating-chat.donateButton.background-color': '#794bc4',
    'floating-chat.donateButton.text-color': '#fff'
  });
</script>

```

De esta forma, los visitantes de tu repositorio verán el código de Ko-fi y podrán incluirlo en sus propios sitios web si así lo desean.
--- 
La Licencia MIT es una de las licencias de código abierto más permisivas y populares. Permite el uso, modificación y distribución del software de forma gratuita, con la única condición de incluir la notificación de derechos de autor y la exención de responsabilidad en todas las copias o porciones sustanciales del software. Esta licencia es adecuada para proyectos de código abierto que no tienen restricciones específicas y que permiten el uso comercial, la modificación y la distribución del código fuente sin requerimientos adicionales.
