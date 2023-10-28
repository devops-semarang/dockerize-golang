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

In this step, you must generate PAT (Personal Access Token) for authentification proces. Create a new personal access token (classic) with the appropriate scopes for the tasks you want to accomplish. If your organization requires SSO, you must enable SSO for your new token.
- Select the read:packages scope to download container images and read their metadata.
- Select the write:packages scope to download and upload container images and read and write their metadata.
- Select the delete:packages scope to delete container images.

```console
docker ps # to check docker daemon running or not
docker login ghcr.io -u USERNAME # to login to GitHub Container Registry
```

## How to build and push Dockerfile to Docker Image

```console
docker build . -t ghcr.io/your-username/dockerize-golang:latest
docker push ghcr.io/your-username/dockerize-golang:latest
```