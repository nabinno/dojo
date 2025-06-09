---
title: Introduction to Docker
tags: docker
url: https://campus.datacamp.com/courses/introduction-to-docker
---

# 1 Using Docker Containers
## Running your first container
```sh
$ docker run hello-world
```

## Running a container in the background
```sh
$ docker run -d postgres
$ docker ps
```

## An interactive container
```sh
$ docker run -it ubuntu
ubuntu> exit
```

## Helping a colleague
```sh
$ docker run --name colleague_project -d my_project
$ docker ps -f name=colleague_project
$ docker logs colleague_project
```

## Cleaning up containers
```
$ docker stop colleague_project
$ docker rm colleague_project
```

## Pulling your first image
```
$ docker pull hello-world
```

## Pulling a specific tag
```
$ docker images ubuntu
$ docker pull ubuntu:22.04
```

## Cleaning up images
```
$ docker rmi ubuntu
$ docker container prune -f ubuntu
$ docker image prune -a -f
```





# 2 Writing Your Own Docker Images
## Sharing your work using a Docker registry
```
$ docker tag spam:v1 docker.mycompany.com/spam:v1
$ docker push docker.mycompany.com/spam:v1
```

## Saving an image to a file
```
docker save -o spam_updated.tar spam:v2
```

## Receiving Docker Images
```
$ docker pull docker.mycompany.com/spam_alice:v3
$ docker run docker.mycompany.com/spam_alice:v3
$ docker load -i spam_bob.tar
$ docker run spam_bob:v3
```

## Building your first image
```
$ cat Dockerfile
FROM ubuntu:22.04

$ docker build .
$ docker build . -t my_first_image
```

## Working in the command-line
```
$ touch Dockerfile
$ nano Dockerfile
FROM ubuntu
$ echo "RUN apt-get update" >>Dockerfile
$ cat Dockerfile
```

## Editing a Dockerfile
```
$ nano Dockerfile
FROM ubuntu:22.04
RUN mkdir my_app

$ docker build --tag my_app .
```

## Creating your own Dockerfile
```
$ touch Dockerfile

$ nano Dockerfile
FROM ubuntu
RUN apt-get update
RUN apt-get install -y python3

$ docker build . --tag my_python_image
```

## Copying files into an image
```
$ nano Dockerfile
FROM ubuntu:22.04
RUN apt-get update
RUN apt-get -y install python3
COPY ./pipeline.py /app/pipeline.py
```

## Copying folders
```
$ nano Dockerfile
FROM ubuntu:22.04
RUN apt-get update
RUN apt-get -y install python3
COPY ./pipeline_v3/ /app/

$ docker build . --tag pipeline_v3
```

## Working with downloaded files
```

```

## Choosing a start command for your Docker image
```

```

## Overriding the default command
```

```

## Pulling a specific tag
```

```

## Adding a CMD
```

```





# 3 Creating Secure Docker Images
## Introduction to Docker caching
```

```

## Docker caching
```

```

## Ordering Dockerfile instructions
```

```

## Changing users and working directory
```

```

## WORKDIR and USER
```

```

## Setting the user
```

```

## Setting the working directory
```

```

## Variables in Dockerfiles
```

```

## Understanding ARG and ENV
```

```

## Overriding ARG in a build
```

```

## Changing behavior when starting a container
```

```

## Creating Secure Docker Images
```

```

## Security best practices
```

```

## Keeping packages up-to-date
```

```

## Be safe, don't use root
```

```
