---
title: "Introduction to Docker"
tags: google-cloud-platform, docker, google-container-registry
url: https://www.qwiklabs.com/focuses/1029
---

# Goal
- How to build, run, and debug Docker containers
- How to pull Docker images from Docker Hub and Google Container Registry
- How to push Docker images to Google Container Registry

# Task
- [x] Hello World
- [x] Build
- [x] Run
- [x] Debug
- [x] Publish

# Supplement
## Hello World
```sh
docker run hello-world
docker images
docker run hello-world
docker ps
docker ps -a
```

## Build
```sh
mkdir test && cd test

cat > Dockerfile <<EOF
FROM node:6
WORKDIR /app
ADD . /app
EXPOSE 80
CMD ["node", "app.js"]
EOF

cat > app.js <<EOF
const http = require('http');

const hostname = '0.0.0.0';
const port = 80;

const server = http.createServer((req, res) => {
    res.statusCode = 200;
      res.setHeader('Content-Type', 'text/plain');
        res.end('Hello World\n');
});

server.listen(port, hostname, () => {
    console.log('Server running at http://%s:%s/', hostname, port);
});

process.on('SIGINT', function() {
    console.log('Caught interrupt signal and will exit');
    process.exit();
});
EOF

docker build -t node-app:0.1 .
docker images
```

## Run
```sh
docker run -p 4000:80 --name my-app node-app:0.1
```

```sh
curl http://localhost:4000
docker stop my-app && docker rm my-app
docker run -p 4000:80 --name my-app -d node-app:0.1
docker ps
docker logs 27af2576b2f3
cd test
docker build -t node-app:0.2 .
docker run -p 8080:80 --name my-app-2 -d node-app:0.2
docker ps
curl http://localhost:8080
curl http://localhost:4000
```

## Debug
```sh
docker logs 94891876f8b0
docker exec -it 94891876f8b0 bash
docker> exit
docker inspect 94891876f8b0
docker inspect --format='{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' 94891876f8b0
```

## Publish
```sh
gcloud config list project
docker tag node-app:0.2 gcr.io/qwiklabs-gcp-a6c1f1ff786cd82d/node-app:0.2
docker images
docker push gcr.io/qwiklabs-gcp-a6c1f1ff786cd82d/node-app:0.2
docker stop $(docker ps -q)
docker rm $(docker ps -aq)
docker rmi -f $(docker images -aq) # remove remaining images
docker images
docker pull gcr.io/qwiklabs-gcp-a6c1f1ff786cd82d/node-app:0.2
docker run -p 4000:80 -d gcr.io/qwiklabs-gcp-a6c1f1ff786cd82d/node-app:0.2
curl http://localhost:4000
```

## Reference
- https://docs.docker.com/engine/reference/commandline/exec/
- https://docs.docker.com/engine/reference/commandline/inspect/#examples
- https://docs.docker.com/engine/reference/builder/

