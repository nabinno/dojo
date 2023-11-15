---
title: Introduction to Databricks
tags: databricks, data-science, data-engineering, analytics
url: https://app.datacamp.com/learn/courses/introduction-to-databricks
---

# 1. Introduction to Databricks
## Why pick a Lakehouse?
```
Which of the following is not a reason to implement a lakehouse architecture?

[ ]You want great performance for both BI and machine learning workloads.
[ ]You want a single architecture to support all of your data workloads.
[x]You want to store all of your data in a proprietory data format.
[ ]You want the most cost-effective architectural solution for your organization.
```

## Benefits of the Databricks Lakehouse
As part of your presentation to the CIO, you would like to list why you should want to switch to the lakehouse architecture. Which of the following are key benefits of the Databricks platform? Select all that apply.
```
[x]Built on open-source technologies
[x]Databricks can run on any major cloud provider.
[x]Databricks enables data teams to be more collaborative.
[ ]Only teams that heavily use Python will benefit from Databricks
[ ]Database administrators can separate their data warehousing workloads away from other languages and personas.
```

## Architectural Decisions
```
## Data Warehouse
Typically very expensive
Performant for structured data only

## Data Lake
Data becomes very messy
Sub-optimal performance for different workloads

## Lakehouse
Great performance across all workloads
Very cost-effective for any data workload
```

## Why Delta?
```
The benefits of storing data in the Delta Lake table format include:

1. Fully ACID-compliant transactions: Delta Lake provides support for ACID (Atomicity, Consistency, Isolation, Durability) transactions, which ensures data integrity and consistency even in the face of concurrent data updates and failures.
2. Unified batch and streaming data sources: Delta Lake is designed to handle both batch and streaming data sources, making it suitable for real-time data processing scenarios.
3. Time travel and table history: Delta Lake allows you to access and query previous versions of the data (table history) through time travel capabilities, which can be useful for auditing, data recovery, and analysis purposes.

So, the correct answers are:

[x]Fully ACID-compliant transactions
[x]Unified batch and streaming data sources
[ ]Rigid table schema limitations
[x]Time travel and table history
```

## Databricks for different personas
```
One of the concerns of your CIO is that your future data architecture must satisfy all the different teams who work with your company's data. While you feel confident that Databricks can, you want to provide a list that clearly outlines supported personas.
Which of the following is not a dedicated persona in the Databricks Lakehouse platform?

[ ]Data engineer
[x]Software developer
[ ]Data scientist
[ ]SQL analyst
```

## Capabilities for each data persona
```
## SQL analyst
Dashboards
SQL Editor
SQL Warehouses

## Data Engineer
Delta Live Tables

## Data Scientist
Model Registry
Model Serving
```

## Administering a Databricks workspace
```

```

## Managing and adding users
```

```

## Setting up a Databricks workspace example
```

```

## Control Plane vs. Data Plane
```

```

## Configure your Databricks workspace
```

```




# 2. Data Engineering
## Getting started with Databricks
```

```

## Data Explorer capabilities
```

```

## Setting up a notebook
```

```

## Cluster configurations
```

```

## Data Engineering foundations in Databricks
```

```

## DataFrames
```

```

## Reading from a database
```

```

## Write an external table
```

```

## Data transformations in Databricks
```

```

## Loading in hosted files
```

```

## Selecting the right language
```

```

## Data pipeline steps
```

```

## Data orchestration in Databricks
```

```

## Possible automations in Databricks
```

```

## Benefits of Delta Live Tables
```

```

## End-to-end data pipeline example in Databricks
```

```

## Data pipeline steps
```

```




# 3. Databricks SQL and Data Warehousing
## Overview of Databricks SQL
```

```

## Benefits of Databricks SQL
```

```

## Databricks SQL in the data workflow
```

```

## Databricks SQL vs. other databases
```

```

## Getting started with Databricks SQL
```

```

## Choosing your SQL warehouse
```

```

## SQL Editor vs. notebooks
```

```

## Creating the usSales table
```

```

## Databricks SQL queries and dashboards
```

```

## Understanding Databricks SQL assets
```

```

## Using parameters in queries
```

```

## Creating a Databricks SQL Dashboard
```

```

## Create a user review query
```

```




# 4. Databricks for Large-scale Applications and Machine Learning
## Overview of Lakehouse AI
```

```

## Lakehouse benefits to ML
```

```

## MLOps tasks in Databricks
```

```

## Using Databricks for machine learning
```

```

## EDA in Databricks
```

```

## Why the ML Runtime?
```

```

## Exploring data in a notebook
```

```

## Model training with MLFlow in Databricks
```

```

## Single node vs. multi node ML
```

```

## Databricks for citizen data scientists
```

```

## Using MLFlow for Tracking
```

```

## Deploying a model in Databricks
```

```

## Models and the Model Registry
```

```

## Why Databricks for model deployment?
```

```

## Example end-to-end machine learning pipeline
```

```

## End-to-end ML pipeline
```

```

## Wrap Up
```

```
