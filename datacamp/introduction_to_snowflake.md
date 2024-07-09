---
title: Introduction to Snowflake
tags: snowflake,analytics,data-engineering
url: https://campus.datacamp.com/courses/introduction-to-snowflake/introduction-to-snowflake-architecture-competitors-and-snowflakesql
---

# 1. Introduction to Snowflake: Architecture, Competitors, and SnowflakeSQL
## Traditional vs. cloud data warehouse
```
Traditional data warehouse
- Limited accessibility
- Limited scalability
- High upfront hardware and software costs

Cloud data warehouse
- Highly scalable and flexible
- Accessible from anywhere
- Utilizes cloud infrastructure
```

## Row versus column oriented database
```
Row-oriented database
- Recording and processing individual sales transactions in an online retail system.
- Updating and managing inventory lvels for individual products.

Column-oriented database
- Running complex queries to track inventory levels and sales performance in a retail system.
- Creating real-time reports for website traffic and user interactions on an e-commerce platform.
- Analyzing and summarizing large volumes of historical sales data for trend analysis and forecasting in a retail chain.
```

## Snowflake use cases
```
[x]Data analytics
[ ]Content creation for online platforms.
[x]To gain insights through advanced analytics
[ ]To increase social media followers
```

## Introduction to Snowflake SQL
```
-- Select pizza_type_id, pizza_size and price from pizzas table
SELECT pizza_type_id, pizza_size, price
FROM pizzas;
```

## Decoupling Compute & Storage
```
[ ]Combining compute and storage resources to optimize performance.
[ ]Storing data and compute resources on separate servers but linked closely.
[x]Separating the compute resources from the data storage allows to scale them independently.
```

## Snowflake Architecture Layers
```
Storage Layer
- Data comperession
- Cloud Storage
- Columnar Storage

Compute Layer
- Virtual Warehouse
- Query Processing

Cloud Services Layer
- Authentication
- Optimizer
```

## Virtual Warehouse
```
What is the role of the "Virtual Warehouse" in Snowflake's architecture?

[x]Handling query processing.
[ ]Storing and organizing data in Snowflake.
[x]Managing compute resources.
[ ]Managing metadata and access control.
```

## Data warehousing platforms
```
[x]Amazon Redshift
[x]Google BigQuery
[ ]Postgres
[ ]Azure Data Lake Storage
```

## Features: Snowflake & its competitors
```
True
- SnowflakeSQL has many similarities with PostgreSQL
- Databricks provides autoscaling
- Postgres can be hosted on Cloud Platforms

False
- Snowflake uses a single pricing model that charges for bath compute and storage.
- Snowflake doesn't support JSON
- Snowflake offers coupled Storage and Compute
```

## Snowflake SQL: Using SELECT and WHERE in Snowflake
```
-- Count all pizza entries
SELECT COUNT(*) AS count_all_pizzas
FROM pizza_type
-- Apply filter on category for Classic pizza types
WHERE category = 'Classic'
    -- Additional condition to filter where name has Cheese in it
    AND name LIKE '%Cheese%';
```




# 2. Snowflake SQL and key concepts
## Snowflake connections and DDL commands
```
True
- SnowSQL is a command-line client that provides a direct way to interact with Snowflake
- Snowflake offers worksheets to write and execute queries in Snowflake Web Interface
- `DROP TABLE` command deleted the data and table structure

False
- We can't connect to Snowflake through JDBC/ODBC drivers
- There is no difference in adding comments in Snowflake versus Postgres
- In Snowflake, you can use 'VALIDATE TABLE' to ensure an operation proceeds if a specific table exists
```

## Snowflake Staging
```

```

## Snowflake database structures and DML
```

```

## Loading data
```

```

## DESCRIBE & SHOW
```

```

## Snowflake data type and data type conversion
```

```

## Data types
```

```

## Datatype conversion
```

```

## Functions, sorting, and grouping
```

```

## String functions
```

```

## Functions & Grouping
```

```

DATE & TIME
100 XPanalytics



# 3. Advance Snowflake SQL Concepts
## Joining in Snowflake
```

```

## NATURAL JOIN
```

```

## The world of JOINS
```

```

## LATERAL JOIN
```

```

## Subquerying and Common Table Expressions
```

```

## Subqueries
```

```

## Understanding CTE
```

```

## CTEs
```

```

## Snowflake Query Optimization
```

```

## Essentials of query optimization
```

```

## Early filtering
```

```

## Query history
```

```

## Handling semi-structured data
```

```

## PARSE_JSON & OBJECT_CONSTRUCT
```

```

## Querying JSON data
```

```

## JSONified
```

```

## Wrap-up
```

```
