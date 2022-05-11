## Pruebas de conocimientos Golang.

###  Descripción:
Desarrollar una APIRest implementando la librería estándar de Go (golang) que permita hacer un
CRUD de Usuarios y sus respectivos Posts, como motor de base de datos se debe implementar
mongoDB utilizando la librería oficial de mongo.
Tanto la API como el motor de base de datos deben correr dentro de un contenedor de Docker,
la gestión de redes y puesta en marcha de los contenderos debe ser gestionada mediante Docker
compose.

### Entregables:
1. Código fuente, archivos de docker y docker compose en un mismo repositorio publico de
   gitlab/github con el respectivo README.md donde se describa las indicaciones o pasos
   para poner en marcha el proyecto.
### Que se va a tener en cuenta:
1. Estructuración del proyecto bajo el estándar propuesto por la comunidad en Go (Golang).
2. Patrones de diseño implementados en el código fuente.
3. Pruebas unitarias y cobertura superiores al 75% de cobertura.
### Que se va a calificar como plus:
1. Implementación de autenticación utilizando JWT y protección de los endpoints

## Pasos para probar por parte del desarrollador

* Ejecutar comando docker-compose build && docker-compose up -d
#### probar los siguientes endpoint 
- /api/v1/user get,post,put,delete
- /api/v1/users get
- /api/v1/post post
- /api/v1/posts get
#### Testing 
- make test
