---
title: Containerization and Virtualization Concepts
tags: data-engineering, os-level-virtualization
url: https://campus.datacamp.com/courses/containerization-and-virtualization-concepts/foundations-of-containerization-and-virtualization
---

# 1. Foundations of Containerization and Virtualization
## Limitations of physical machines
```
[x]Physical machines are costly.
[ ]Physical machines are not very customizable.
[x]Physical machines are complex to scale.
[x]Physical machines need maintenance.
```

## Defining virtual machine
```
[ ]A physical computer inside another physical computer.
[x]A simulated computer inside another computer.
[ ]A program that creates a visual representation of a computer screen within the display of another computer.
```

## Benefits of virtualization
```
True:
- VMs are platform-independent, and different operating systems can run seamlessly on the same host machine.
- Virtual machines offer greater scalability and flexibility than physical machines.
- Virtualization maximizes resource utilization, resulting in cost efficiencies and sustainable usage

False:
- Applications running in different virtual machines are less isolated than applications running on the same physical machine.
- Virtual machines typically cause more downtime than physical machines.
```

## Definition of containerization
```
[x]The process of packaging an application and its dependencies into a container, allowing it to be deployed and executed consistently across different environments.
[ ]The process of creating a virtual version of computers or hardware components, such as a server, storage device, or network resource, that can be accessed and utilized remotely.
[x]Containerization is virtualization at the operating system level.
```

## Benefits of containers
```
True:
- Containers start up quickly because there is no need to load the entire operating system.
- Applications are easily portable and reproducible in different environments because containers contain everything needed to run the application.
- Containers on the same host are isolated, so a crash in one container will not cause a crash in another.

False:
- Applications with different hardware requirements can run on the same host machine because the hardware can be customized for each container.
- Applications can run on different operating systems on the same host machine because containerization is virtualization at the OS-level.
```

## Virtual machines vs. containers
```
Containers:
- Run applications in virtualized, isolated user spaces (OS-level virtualization).
- Use software tools such as Docker and Kubernetes.
- Share the host systems OS kernel.

Virtual machines:
- Virtualize on entire computer (full virtualization).
- Managed by software tools such as VMware and VirtualBox.
- Have their own operating system that they do not have to share.
```

## Use cases of containerization & virtualization
```
[x]Microservices architecture: Break down large applications into smaller services that can be deployed independently.
[ ]Legacy application support: Maintain compatibility with older operating systems or hardware dependencies.
[x]Orchestration: Automate the deployment, scaling, and management of applications.
[ ]Server consolidation: Run multiple virtual machines on a single physical server to improve resource utilization and reduce hardware costs.
```

## Benefits of containers vs. virtual machines
```
Containers:
- Rapid scalability: Quickly scale applications up and down based on demand.
- Reduced over head costs: Reduce the need for expensive IT infrastructure and resources.
- Portability: Move seamlessly between different environments with minimal configuration changes.

Virtual machines:
- Multiple environments: Multiple, different operating systems can run on the same host machine.
- Legacy application support: Provision of legacy applications by emulating older hardware/software.
- Strong & secure isolation: High level of security with a completely isolated environment.
```




# 2. Applications of Containerization
## The Docker platform
```
[ ]Tools for building virtual machines
[x]Tools for building, distributing, and running containers
[ ]Tools for developing mobile applications
[ ]Tools for creating graphical user interfaces
```

## Actions with Docker objects
```

```

## Docker's client-server architecture
```

```

## Container orchestration
```

```

## Definition of container orchestration
```

```

## Benefits of container orchestration
```

```

## Application of container orchestration
```

```

## Container orchestration with Kubernetes
```

```

## The Kubernetes platform
```

```

## Important Kubernetes components
```

```

## Docker and Kubernetes
```

```

## Reading Dockerfiles and running containers
```

```

## Understanding a Dockerfile
```

```

## Building and running a Docker container
```

```

## Wrap-up
```

```
