# Dockerize Go-Lang Application
Dockerizing your Go application allows for easy deployment and scalability while ensuring consistency across different environments. Follow the steps outlined here to package your Go application into a Docker container.

## MySQL in Docker

```console
docker run --name mysql \
        -e MYSQL_ROOT_PASSWORD=f214e666b9ededb8acdf0780cc796cf1370a41de47f97810f09ddfbbc237ea3f \
        -e MYSQL_DATABASE=portal_web \
        -e MYSQL_USER=administrator \
        -e MYSQL_PASSWORD=f214e666b9ededb8acdf0780cc796cf1370a41de47f97810f09ddfbbc237ea3f \
        -p 3306:3306 \
        -d mysql:5.7 \
        --character-set-server=utf8mb4 --collation-server=utf8mb4_unicode_ci
```

## Sign in to the Container registry service at ghcr.io.

```console
docker ps
docker login ghcr.io -u USERNAME --password-stdin
```