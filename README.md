# Introducción

Esta es una solución simple para un manejador (CRUD) de productos. Se usa una arquitectura hexagonal donde los componentes se encuentran aislados y cada uno puede ser probado y usado por separado.
Las tecnologías utilizadas son:
- GIN para el servicio de la API
- GORM para la comunicación con la base de datos

Para la conexión con la base de datos se utiliza el patron de diseño "Singleton" generando solo una instancia a la base de datos en la vida del artefacto.

# Ejecución

Es necesario contar con MySQL instalado para generar la conección a la base de datos.

```bash
//Primero clonar o descargar el repositorio
git clone https://github.com/Kurozzu/test-agrak.git

//Descargar e instalar los paquetes de golang
go mod tidy

//Cambiar los datos de conexión a la base de datos en el archivo
test-Agrak/libs/mysql.go

//Iniciar directamente desde el proyecto
go run main.go

//También se puede hacer un build
go build .

//Si presenta algún problema al generar el build de la app, recuerde cambiar el GOENV

```
# Endpoints

```bash
NAME          | METHOD        | URL                                         |
------------- | ------------- | ------------------------------------------- |
Insert        | POST          | http://localhost:8080/product               |
Get All       | GET           | http://localhost:8080/product               |
Get           | GET           | http://localhost:8080/product/FAL-881952283 |
Update        | PUT           | http://localhost:8080/product               |
Delate        | DELETE        | http://localhost:8080/product/FAL-881952283 |
```

## Crear producto

Permite crear nuevos productos uno a uno. Lista de ejemplo para agregar 3 productos.
```bash
{
    "sku": "FAL-881952283",
    "name": "Bicicleta Baltoro Aro 29",
    "brand": "Jeep",
    "size": "ST",
    "price": 399990.00,
    "mainimage": "https://ruta update 2",
    "otherimage": "https://other 2"
}

// {
//     "sku": "FAL-8818985 02",
//     "name": "Camisa Manga Corta Hombre",
//     "brand": "Basement",
//     "size": "M",
//     "price": 24990.00,
//     "mainimage": "https://ruta update 3",
//     "otherimage": "https://other 3"
// }


// {
//     "sku": "FAL-8406270",
//     "name": "500 Zapatilla Urbana Mujer",
//     "brand": "New Balance",
//     "size": "37",
//     "price": 42990.00,
//     "mainimage": "https://ruta update 1",
//     "otherimage": "https://other 1"
// }
```

## Postman

El archivo postman contiene los endpoint con datos listos para poder hacer pruebas una vez la API se encuentra corriendo.

# Introduction

This is a simple manager solution (CRUD) of products. It's used a hexagonal architecture where the components are isolated and each can be tested and deployed separately.
Technologies used are:
- GIN for services the API
- GORM for database communication

For database connection is used "Singleton" design pattern that allow to generate one instance to database

# Execution

Is necessary  to have MySQL installed to generate the database connection

```bash
//Firt, clone or download repository
git clone https://github.com/Kurozzu/test-agrak.git

//Download or install the packages of golang
go mod tidy

//Change the database connection data in the file file
test-Agrak/libs/mysql.go

//Started in the same project
go run main.go

//Also it is possible to make a build
go build .

//If you have some problem when app is building, please remember to change the GOENV

```
# Endpoints

```bash
NAME          | METHOD        | URL                                         |
------------- | ------------- | ------------------------------------------- |
Insert        | POST          | http://localhost:8080/product               |
Get All       | GET           | http://localhost:8080/product               |
Get           | GET           | http://localhost:8080/product/FAL-881952283 |
Update        | PUT           | http://localhost:8080/product               |
Delate        | DELETE        | http://localhost:8080/product/FAL-881952283 |
```

## Crear producto

Allows to create new products one by one. List with examples to add 3 products
```bash
{
    "sku": "FAL-881952283",
    "name": "Bicicleta Baltoro Aro 29",
    "brand": "Jeep",
    "size": "ST",
    "price": 399990.00,
    "mainimage": "https://ruta update 2",
    "otherimage": "https://other 2"
}

// {
//     "sku": "FAL-8818985 02",
//     "name": "Camisa Manga Corta Hombre",
//     "brand": "Basement",
//     "size": "M",
//     "price": 24990.00,
//     "mainimage": "https://ruta update 3",
//     "otherimage": "https://other 3"
// }


// {
//     "sku": "FAL-8406270",
//     "name": "500 Zapatilla Urbana Mujer",
//     "brand": "New Balance",
//     "size": "37",
//     "price": 42990.00,
//     "mainimage": "https://ruta update 1",
//     "otherimage": "https://other 1"
// }
```
## Postman

The postman file contains the endpoints with data and it is ready for testing once the API is running.
