---
title: Understanding Modern Data Architecture
tags: data-architecture, data-management
url: https://campus.datacamp.com/courses/understanding-modern-data-architecture/introduction-to-modern-data-architecture
---

# 1. Introduction to Modern Data Architecture
## What is a modern data architecture?
```
Which of the following is an appropriate definition of modern data architecture?

[ ]A secure, flexible, and scalable system for organizing and storing digital files.
[ ]A secure, flexible, and scalable system for real-time data processing and analysis.
[x]A secure, flexible, and scalable system for ingesting, managing, and analyzing large and diverse datasets.
```

## Modern vs traditional architectures
```
Modern data architecture;
- Self-service analytics tools
- Serverless architecture minimizing infrastructure management
- Auto-scaling capabilities to handle dynamic workloads
- Real-time processing

Traditional data architecture:
- Scheduled batch processing
- Data is primarily sourced from structured databases
- Heavy Reliance on In-House Infrastructure
```

## Lambda architecture
```
[ ]A traditional batch processing system that handles data in large, periodic batches.
[x]A processing approach for batch and streaming data independently that merges them to provide a unified view of data.
[ ]A processing approach that separates data into two paths: batch and real-time.
```

## Lambda architecture layers
```
Batch layer:
- Master dataset
- Batch processing

Speed layer:
- Stream processing
- Stream views

Serving layer:
- Unified views
- Query interface
```

## Kappa architecture
```
[ ]An architecture that separates batch and real-time processing paths to handle high-velocity data streams.
[x]An architecture that treats all data as one stream, enabling a single stack for batch and streaming processing.
[ ]An architecture that processes only streaming data and does not handle history.
```

## Kappa: What about the batch layer?
```
[ ]Change your current architecture to Lambda, as you need to process historical data.
[ ]Create a new flow for extracting the new data, and wait a couple of months to collect enough history.
[x]Create a new flow for extracting the new data, and re-process the source of events to get the whole history.
```

## Data Mesh
```
[ ]A decentralized data processing approach.
[ ]A decentralized data storage system.
[x]A decentralized approach to data management.
[ ]A centralized approach to data management.
```

## Data Fabric
```
[ ]A data architecture that focuses on visualization to integrate diverse datasets.
[ ]A centralized data architecture with a great focus on metadata analysis.
[x]A data architecture with a great focus on metadata analysis.
```

## Data architecture & its use cases
```
Data mesh:
- A company with multiple business units needs to share data while maintaining their own data management and enabling self-service consumption.
- Decentralized datqa ownership and management for autonomous product teams.

Data fabric:
- Building a data platform leveraging metadata analysis to improve data management.
- Integrate disparate systems and data sources to provide an unified view of data assets.
```




# 2. Modern Data Architecture Components
## What are blob storages?
```
[ ]A data lake.
[x]A distributed storage with support of all types of data.
[ ]A single physical device used to store data in a blob format.
```

## When should you use blob storage?
```
[ ]Primarily work with structured data, and need to execute complex queries over it.
[x]Archive data for compliance purposes, but you won't access it normally.
[x]Implement a data lake, which will receive multiple types of data; including unstructured.
[ ]Store and serve semi-structured data with high demand.
```

## SQL vs. NoSQL
```
[ ]RDBMS. You could scale servers horizontally, and query with SQL.
[x]NoSQL. You could scale horizontally and be eventually consistent across the system.
[ ]NoSQL. You could scale vertically to meet the demand.
```

## Storage use cases
```
Blob storage:
- Archive data for the long-term
- Store large media files

NoSQL DB:
- Serve applications with high demand, and not necessarily a strong consistency
- Store semi-structured data

Data warehouse:
- Analyze huge amount of structured data, using SQL.
```

## Data Ingestion
```
Ingestion is:
- Collecting and importing data from various sources
- A continuous and ongoing activity
- Necessary for data analysis and decision-making

Ingestion is not:
- Limited to a specific type of data source or format.
- Responsible for data interpretation or drawing insights
```

## What is ingestion?
```

```

## What is the landing zone?
```

```

## Ingesting new data
```

```

## Data Processing
```

```

## Batch vs. Streaming
```

```

## What is a window in streaming?
```

```

## Processing paradigms
```

```

## Data serving
```

```

## Serving technologies
```

```

## Sort the layers
```

```




# 3. Transversal Components of Data Architectures
## Data governance
```

```

## What is data governance?
```

```

## Data governance roles
```

```

## Metadata Management
```

```

## The role of metadata
```

```

## Metadata types
```

```

## The business glossary is composed of...
```

```

## Data Security
```

```

## What is a data breach?
```

```

## Virtual Private Clouds (VPC)
```

```

## Effective data security
```

```

## Classifying Data Security Measures
```

```

## Observability
```

```

## Troubleshooting the source of a problem
```

```

## Data security measures
```

```




# 4. Putting it All Together
## Orchestration
```

```

## Orchestration vs. Scheduling
```

```

## Orchestration sensors
```

```

## Storage & processing costs best practices
```

```

## Reserved capacity in cloud services
```

```

## Storage classes
```

```

## Designing a modern data architecture
```

```

## Asking the right questions
```

```

## Ingesting the application data
```

```

## Evaluating modern data architecture solutions
```

```

## Why a quarantine zone?
```

```

## Designing a modern data platform
```

```

## Wrap-up
```

```
