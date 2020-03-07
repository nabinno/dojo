---
title: Add Healthcheck for Containers
tags: docker,healthcheck
url: https://learning.oreilly.com/scenarios/add-healthchecks-to/9781492061540/
---

# 1. Creating Service
```Dockerfile
FROM katacoda/docker-http-server:health
HEALTHCHECK --timeout=1s --interval=1s --retries=3 \
  CMD curl -s --fail http://localhost:80/ || exit 1
```

```sh
Your Interactive Bash Terminal. A safe place to learn and execute commands.
$
$ docker build -t http .
Sending build context to Docker daemon  2.048kB
Step 1/2 : FROM katacoda/docker-http-server:health
health: Pulling from katacoda/docker-http-server
12b41071e6ce: Pull complete
fb1cef6edba2: Pull complete
1061ea2815dd: Pull complete
Digest: sha256:fee2132b14b4148ded82aacd8f06bdcb9efa535b4dfd2f1d88518996f4b2fb1d
Status: Downloaded newer image for katacoda/docker-http-server:health
 ---> 7f16ea0c8bd8
Step 2/2 : HEALTHCHECK --timeout=1s --interval=1s --retries=3   CMD curl -s --fail http://localhost:80/ || exit 1
 ---> Running in 70426f520fe1
Removing intermediate container 70426f520fe1
 ---> 8bfbf00d0a35
Successfully built 8bfbf00d0a35
Successfully tagged http:latest
$ docker run -d -p 80:80 --name srv http
8b576f16fec07acdf8090b7dfb00e5603ebc8881104c01f5d4ba87b64b70125d
```

# 2. Crash Service
```sh
$ docker ps
CONTAINER ID        IMAGE               COMMAND             CREATED             STATUS                    PORTS                NAMES
8b576f16fec0        http                "/app"              51 seconds ago      Up 50 seconds (healthy)   0.0.0.0:80->80/tcp   srv
$ curl http://docker/unhealthy
```

# 3. Verify Status
```sh
$ docker inspect --format "{{json .State.Health.Status }}" srv
"unhealthy"
$ docker inspect --format "{{json .State.Health }}" srv
{"Status":"unhealthy","FailingStreak":48,"Log":[{"Start":"2020-03-07T01:10:52.99480533Z","End":"2020-03-07T01:10:53.042402583Z","ExitCode":1,"Output":""},{"Start":"2020-03-07T01:10:54.048449606Z","End":"2020-03-07T01:10:54.097565081Z","ExitCode":1,"Output":""},{"Start":"2020-03-07T01:10:55.103429673Z","End":"2020-03-07T01:10:55.158515221Z","ExitCode":1,"Output":""},{"Start":"2020-03-07T01:10:56.166361404Z","End":"2020-03-07T01:10:56.21747501Z","ExitCode":1,"Output":""},{"Start":"2020-03-07T01:10:57.222655272Z","End":"2020-03-07T01:10:57.275237598Z","ExitCode":1,"Output":""}]}
$ docker ps
CONTAINER ID        IMAGE               COMMAND             CREATED              STATUS                          PORTS                NAMES
8b576f16fec0        http                "/app"              About a minute ago   Up About a minute (unhealthy)   0.0.0.0:80->80/tcp   srv
```

# 4. Fix Service
```sh
$ curl http://docker/healthy
$ docker ps
CONTAINER ID        IMAGE               COMMAND             CREATED             STATUS                   PORTS                NAMES
8b576f16fec0        http                "/app"              2 minutes ago       Up 2 minutes (healthy)   0.0.0.0:80->80/tcp   srv
$ docker inspect --format "{{json .State.Health.Status }}" srv
"healthy"
```

# 5. Healthchecks with Swarm
```sh
$ docker rm -f $(docker ps -qa);
8b576f16fec0
$ docker swarm init
Swarm initialized: current node (lx3abqwifwv2i129ct6eppc1p) is now a manager.

To add a worker to this swarm, run the following command:

    docker swarm join --token SWMTKN-1-021mth2nahpj4zdsxpoaxnxra33shdqms11zntlxzs4rjlh36n-2s7o71d4uwooxnn0n7z4bx8lc 172.17.0.25:2377

To add a manager to this swarm, run 'docker swarm join-token manager' and follow the instructions.

$ docker service create --name http --replicas 2 -p 80:80 http
image http:latest could not be accessed on a registry to record
its digest. Each node will access http:latest independently,
possibly leading to different nodes running different
versions of the image.

5a5a5lt2easekla5qumoi3nud
overall progress: 0 out of 2 tasks
overall progress: 2 out of 2 tasks
1/2: running   [==================================================>]
2/2: running   [==================================================>]
verify: Service converged
$ curl host01
<h1>A healthy request was processed by host: a1de73626776</h1>
$ curl host01/unhealthy
$ curl host01
<h1>A healthy request was processed by host: a1de73626776</h1>
$ docker ps
CONTAINER ID        IMAGE               COMMAND             CREATED             STATUS                    PORTS               NAMES
a1de73626776        http:latest         "/app"              19 seconds ago      Up 17 seconds (healthy)   80/tcp              http.2.myldxdr5rwuvebq1sm8xre8tx
$ curl host01
<h1>A healthy request was processed by host: a1de73626776</h1>
```

