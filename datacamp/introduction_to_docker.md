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
$ touch Dockerfile
$ nano Dockerfile
FROM ubuntu
RUN apt-get update
RUN apt-get install -y python3 curl unzip

RUN curl https://assets.datacamp.com/production/repositories/6082/datasets/31a5052c6a5424cbb8d939a7a6eff9311957e7d0/pipeline_final.zip -o /pipeline_final.zip
RUN unzip /pipeline_final.zip
RUN rm /pipeline_final.zip

$ docker build . --tag pipeline
```

## Overriding the default command
```
$ docker run -it postgres bash
```

## Pulling a specific tag
```
$ docker images
$ docker pull ubuntu:22.04
```

## Adding a CMD
```
$ nano Dockerfile
FROM ubuntu:22.04
RUN apt-get update
RUN apt-get -y install python3
CMD python3

$ docker build . --tag pipeline_debug

$ docker run pipeline_debug
```





# 3 Creating Secure Docker Images
## Docker caching
```
[ ]Docker builds Dockerfiles into images; an image is composed of layers that correspond to specific Dockerfile instructions. A layer can be re-used for Dockerfiles with identical instructions.
[ ]When we build an image from a Dockerfile, every Dockerfile instruction is run, and the changes it makes to the file system are saved. The bundle of these changes to the file system is called a layer.
[ ]Image layer caching can be complex, but it allows us to understand how to greatly increase the speed with which we can iterate on, i.e., improve or fix bugs in our images.
[x]All of the above.
```

## Ordering Dockerfile instructions
```
FROM docker.io/library/ubuntu
RUN apt-get update
RUN apt-get install -y python3
COPY /app/requirements.txt /app/requirements.txt
COPY /app/pipeline.py /app/pipeline.py
```

## WORKDIR and USER
```
[ ]After using `WORKDIR` in our Dockerfile, no instructions after `WORKDIR` can use any other path than the one we set with `WORKDIR`, until the workdir is changed again.
[x]`WORKDIR` allows us to change the path in which the command of the `CMD` instruction is run.
[x]After using `USER` in our Dockerfile, no instructions after `USER` can use any other user than the one we set with `USER`, until the user is changed again.
[x]`USER` allows us to change the user with which the command of the `CMD` instruction is run.
```

## Setting the user
```
$ nano DockerfileFROM ubuntu:22.04
RUN useradd -m repl
USER repl
RUN mkdir /home/repl/projects/pipeline_final
COPY /home/repl/project /home/repl/projects/pipeline_final
```

## Setting the working directory
```
FROM ubuntu:22.04
RUN useradd -m replUSER repl
WORKDIR /home/repl
RUN mkdir projects/pipeline_final
COPY /home/repl/project projects/pipeline_final
```

## Understanding ARG and ENV
```
[ ]Variables set in a Dockerfile using the ARG instruction are not accessible after the image is built. This means it is safe to use ARG to store secrets in a Dockerfile.
[x]Variables set using ENV can be used in containers starting from your image, making it a good way to set configuration using a runtime.
[x]It is possible to override variables set with ARG during the build, allowing us to configure images at build-time.
[x]Every user starting a container from our image can select a different value for any ENV variables we set in our image.
```

## Overriding ARG in a build
```
$ cat Dockerfile
FROM ubuntu:22.04
ARG WELCOME_TEXT=Hello!
RUN echo $WELCOME_TEXT
CMD echo $WELCOME_TEXT

$ docker build . --build-arg WELCOME_TEXT=Welcome!
```

## Changing behavior when starting a container
```
$ cat Dockerfile
FROM ubuntu:22.04
ENV NAME=Tim
CMD echo "Hello, my name is $NAME"

$ docker build . -t hello_image

$ docker run -e NAME=nabinno hello_image
```

## Security best practices
```
[x]Using a container is a good way to run an executable or open an archive from an untrusted source because you greatly decrease the chance of a malicious actor accessing your computer.
[ ]If I don't follow all security precautions, I might as well not follow any.
[ ]Because of isolation between a container and the host OS, nothing in the container can ever affect the host environment.
[x]There is no safer application than one we haven't installed.
[x]When creating an image ourselves, we can increase the security by changing the Linux user to something other than the root user.
```
## Be safe, don't use root
```
$ cat Dockerfile
FROM ubuntu:22.04
RUN useradd -m repl
USER repl
CMD apt-get install python3

$ docker run repl_try_install
```
