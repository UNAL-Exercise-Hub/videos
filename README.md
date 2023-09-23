# Videos
Repository for microservice with Golang and mySQL

## Despliegue de la Base de Datos
Para desplegar la base de datos, sigue estos pasos desde la carpeta db_dockerfile:

1. **Construye la imagen de MySQL:**

   ```bash
    docker build -t mysql_videos_db .
   
2. **Ejecuta el contenedor de MySQL:**

     ```bash
    docker run -d -t -i -p 3306:3306 --name mysql_videos_db mysql_videos_db
     
3. **Opcional: Ejecuta phpMyAdmin para administrar la base de datos:**

    ```bash
    docker run --name db_client_videos -d --link mysql_videos_db:db -p 8081:80 phpmyadmin
    
El contenedor MySQL estará en funcionamiento, y si lo deseas, puedes usar phpMyAdmin en http://localhost:8081 para gestionar la base de datos.

## Despliegue de microservicio
Para desplegar el microservicio, sigue estos pasos desde la carpeta principal:

1. **Construye la imagen del ms:**

   ```bash
   docker build -t golang_videos_ms .

2. **Ejecuta el contenedor del ms:**

   ```bash
   docker run -d -p 8080:8080 --name golang_videos_ms golang_videos_ms


Ahora podras usar la aplicación desde la terminal del contenedor
