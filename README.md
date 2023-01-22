# Online Store
Online Store Backend side using Golang (Gin, Gorm, JWT, MySQL, and MongoDB)

# ER Digaram
Implement the ER Diagram. ERD is a diagram that describes the relationship between data objects that have relationships
between relationships.

![alt text](/asset/erd.png)


# Flow Apps
![alt text](/asset/flow-online-store-simple.drawio.png)


# How to use it 
## Using local
- Clone source code ``` git clone https://github.com/aldiset/online_store ```
- Intall mysql and create database (set environment according your config database)
- Install MongoDB
- Run App with ``` go run main.go ```
- You can access ```localhost:8080 ``` or ```127.0.0.1:8080``` as base path.

## Using Docker
- Pull image from ``` https://hub.docker.com/r/akarkode/gin-online-store ``` with command ``` docker pull akarkode/gin-online-store ```
- Intall mysql and create database (set environment according your config database)
- Install MongoDB
- Running docker container with ```Environment``` MySQL and MongoDB config.
- You can access ```localhost:8080 ``` or ```127.0.0.1:8080``` as base path.

## Using Docker Compose
- Clone source code ``` git clone https://github.com/aldiset/online_store ```
- Run command ``` docker-compose up ``` or ``` docker-compose up -d``` (make sure docker-compose is installed).
- You can access ```localhost:8080 ``` or ```127.0.0.1:8080``` as base path.

# Access Swagger
- You can access swagger from ```http://localhost:8080/api/docs/index.html#/```
[asset_access-swagger.webm](https://user-images.githubusercontent.com/52232543/213930130-dd45d11c-4af4-4382-b8bc-23baf28d3262.webm)
