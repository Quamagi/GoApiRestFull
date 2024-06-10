![logo](https://raw.githubusercontent.com/Quamagi/GoApiRestFull/main/logo.jpg?token=GHSAT0AAAAAACSTN3V4AHGEMGG2M3XGLB6YZTE72ZA)

Para agregar tu código de Ko-fi al README de tu repositorio de GitHub, puedes incluirlo en una sección dedicada a las donaciones o soporte. Sin embargo, ten en cuenta que el código HTML y JavaScript no se ejecutará directamente en un archivo README.md en GitHub, ya que Markdown no soporta la ejecución de scripts. Puedes proporcionar el código como referencia para que los visitantes lo agreguen a su propio sitio web.

Aquí tienes un ejemplo de cómo puedes hacerlo en tu README.md:

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


## Apóyame en Ko-fi

Si encuentras útil este proyecto y deseas apoyarme, puedes hacerlo a través de Ko-fi:

<a href='https://ko-fi.com/quamagi' target='_blank'>
<img height='36' style='border:0px;height:36px;' src='https://storage.ko-fi.com/cdn/kofi3.png?v=3' border='0' alt='Buy Me a Coffee at ko-fi.com' />
</a>

## Licencia

Este proyecto está bajo la Licencia MIT. Consulta el archivo LICENSE para más detalles.
