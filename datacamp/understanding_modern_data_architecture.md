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

## What is the landing zone?
```
[ ]The final storage location for processed data.
[ ]A dedicated area for data quality checks and cleansing.
[ ]A component that analyzes data patterns and generates insights.
[x]The initial storage area for incoming data.
```

## Ingesting new data
```
Your company plans to build a cloud-based data platform. For this purpose, they want to replicate data from various sources, including cloud-native apps and on-premises SQL databases. The data science team requested a copy of the data in the cloud for analytical purposes. Once the data is in the cloud, your company wants to also be able to consume it directly from its serving data warehouse. It is acceptable for the data to have a delay of up to 24 hours if it reduces costs.

How would you design your data platform ingestion to fulfill your company’s requirements?

[ ]Batch job to ingest from databases and store in the serving data warehouse.
[ ]Streaming job to listen to changes and replicate to the serving data warehouse.
[x]Batch job to ingest from databases and store in the landing zone.
[ ]Streaming job to listen to changes and replicate to the landing zone.
```

## Batch vs. Streaming
```
Batch:
- Schedule
- Fixed amount of data

Streaming:
- Runs 24/7
- Real-time insights
- Windows are needed to perform aggregations
```

## What is a window in streaming?
```
[ ]Temporary storage areas where real-time data is buffered before processing.
[x]Time-based or size-based partitions of the continuous data stream for performing computations and aggregations.
[ ]Divisions of data into fixed intervals for efficient data transmission.
```

## Processing paradigms
```
You are an engineer working for a data processing company that receives files from an external provider. These files contain critical data that needs to be validated for integrity before being accepted into the system. The challenge you are facing is that the arrival pattern of these files is unpredictable. They could either arrive in large batches during the morning or night, or sporadically throughout the day. Additionally, the file sizes are relatively small, and you won't get more than a couple of thousand per day.

Which processing strategy would you choose that minimizes costs?

[ ]Streaming in a cluster
[ ]Streaming in a serverless service
[x]Function computing
```

## Serving technologies
```
Data Warehouse:
- You have structured data, and need to create dashboards.
- You have structured and semi-structured data, and need to perform analytical querying.

Blob Storage:
- You have structured data, and need to archive it to the long term without much querying.
- You have unstructured data, and would like to train machine learning models.

RDBMS:
- You have structured data, and need to expose it via an API with single record lookup capability.
```

## Sort the layers
```
1. Ingestion
2. Landing zone
3. Processing
4. Serving
5. Consume
```




# 3. Transversal Components of Data Architectures
## What is data governance?
```
[ ]The process of analyzing and interpreting data to make informed business decisions.
[ ]A technique used to secure data from unauthorized access and protect it from cyber threats.
[x]A set of rules and policies that dictate the creation, collection, storage, usage, and disposal of an organization's data.
```

## Data governance roles
```
Governors:
- Classify data into categories like PII
- Provide or restrict access to the data

Users:
- Create machine learning models from data
- Use dashboards to make informed decisions

Ancillary:
- Userstand and inform legal requirements from industry regulations
- Funding the data governance strategy
```

## The role of metadata
```
[ ]Metadata enables data encryption and security.
[ ]Metadata is crucial for the development of machine learning models.
[ ]Metadata is primarily used to increase data storage capacity.
[x]Metadata assists in organizing, accessing, and understanding data.
```

## Metadata types
```
Technical:
- Movies dataset schema
- Tables foreign keys

Business:
- Owner for moviews dataset
- Meaning of each column in movies dataset

Operational:
- Last time moviews dataset was processed
- Percentage of null values in the movie's genre
```

## The business glossary is composed of...
```
[x]Data owner
[ ]Tables relationships
[x]Business definitions
[ ]Timestamps of data processing
```

## What is a data breach?
```
[ ]A data analyst can't access the necessary data due to insufficient permissions.
[ ]A sales representative can't see customer names in a report because of data masking.
[x]An unauthorized user obtains sensitive customer information.
[ ]Data is inaccessible due to a network outage caused by a malicious actor.
```

## Virtual Private Clouds (VPC)
```
[ ]VPCs encrypt your data at rest.
[ ]VPCs are used to manage user access permissions.
[ ]VPCs manage the lifecycle of your cryptographic keys.
[x]VPCs allow you to create isolated networks, thus restricting access if needed.
```

## Effective data security
```
[x]Implementing strong access controls
[x]Using data encryption at rest
[x]Regularly updating firewall rules
[ ]Keeping all data in a single location for ease of access
[x]Applying data masking to sensitive fields
```

## Classifying Data Security Measures
```
Cloud services:
- VPC
- IAM services

Data protection:
- Masking sensitive data
- Encrypting data in transit
- Using a Customer Managed Key

Network security:
- Defining firewall rules
```

## Troubleshooting the source of a problem
```
You've noted a particular transaction that's been unusually slow and want to analyze it. Which aspect of observability would help you understand exactly where the slowdown is happening in the transaction's path?

[ ]Monitoring to keep track of the system's overall performance.
[ ]Logging to keep a record of when the transaction occurred.
[x]Tracing to follow the transaction's journey through the system.
```

## Data security measures
```
Easier to implement:
- Data encryption
- Access control
- Audit -logging and monitoring

Harder to implement:
- Security testing and validation
- Network security
```




# 4. Putting it All Together
## Orchestration vs. Scheduling
```
Orchestration:
- Managing a complex machine learning pipeline with conditional executions.
- Coordinating a multi-step ETL process with dependencies between them

Scheduling:
- Running a nightly backup of a database.
- Launching a series of independent tasks at various intervals throughout the day.
- Triggering on email report to be sent every Monday morning.
```

## Orchestration sensors
```
Within the context of Apache Airflow, which of the following best describes a scenario where the use of a sensor operator would be appropriate?
[ ]Executing a SQL query.
[ ]Triggering an email notification once a particular task within a DAG has been completed.
[ ]Running a Bash command to copy files from one directory to another.
[x]Pausing a workflow until a particular file lands in a specific location, and then continuing the subsequent tasks.
```

## Reserved capacity in cloud services
```
[ ]Reduced costs for unexpected spikes in usage.
[ ]Flexibility to change the resources without any commitments.
[x]Discounts for committing to certain usage levels over an extended period.
[ ]No charges for networking or data transfer between regions.
```

## Storage classes
```
Standard:
- Storing financial reports that are accessed daily.
- Storing real-time analytics data for immediate processing.

Glacier/Coldline:
- Keeping backups of old project files that are rarely accessed.
- Archiving emails that are over 5 years old.
```

## Asking the right questions
```
Select which question would be better to ask based on the business case reviewed in the video.
[ ]What programming language was used to create the API for the model?
[x]What is the level of expertise of the engineering team in the cloud provider being use?
[ ]What is the separator of the CSV files generated by the robots?
```

## Ingesting the application data
```
[ ]Load it directly to BigQuery from the app database.
[ ]Export deltas of the data during low load traffic periods to the landing zone.
[ ]Use a message queue like Pub/Sub to replicate each event in the application, and read the data from the queue.
[x]All of the above
```

## Why a quarantine zone?
```
[ ]To reduce the amount of data to be processed by further stages.
[ ]To discard corrupted records.
[x]To later analyze the errors and reprocess them after fixing them.
```

## Designing a modern data platform
```

```

## Wrap-up
```

```
