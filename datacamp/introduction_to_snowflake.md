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
[ ]Snowflake Staging is an area where data transformations occur before they're loaded into Snowflake tables.
[x]Snowflake Staging is an intermediary storage area used for storing data files before they are loaded into Snowflake tables.
[ ]Snowflake Staging is the Snowflake's backup mechanism for saving database states in case of system failures.
```

## Loading data
```
[ ]LOAD DATA INTO pizza_type FROM <//your/local/file/path>
[ ]INSERT INTO pizza_type from @my_local_stage
[ ]UPLOAD DATA INTO pizza_type from @my_local_stage
[x]COPY INTO pizza_type FROM @my_local_stage
```

## DESCRIBE & SHOW
```
Correct Syntax:
- DESCRIBE TABLE orders
- DESCRIBE DATABASE pizza
- SHOW TABLES LIKE '%ORDERS%' IN DATABASE pizza

Incorrect Syntax:
- SHOW TABLES FOR DATABASE pizza
- SHOW COLUMNS OF TABLE orders
- DESCRIBE IN VIEW ORDERS_VIEW
```

## Data types
```
True:
- `TIMESTAMP` is a combination of `DATE` and `TIME`.
- The default `DATE` format in Snowflake is `YYYY-MM-DD`.
- The default format for `TIME` is `HH:MI:SS`.

False:
- Postgres `VARCHAR` can store more characters than Snowflake `VARCHAR`.
- `TIMESTAMP` in Snowflake doesn't store information about the minutes in Time.
```

## Datatype conversion
```
SELECT
    -- Convert request_id to VARCHAR using CAST method and alias to request_id_string
    CAST(request_id AS VARCHAR) AS request_id_string,
    -- Convert request_timestamp to DATE using TO_DATE and alias as request_date
    TO_DATE(request_timestamp) AS request_date,
    -- Convert drop_timestamp column to TIME using :: operator and alias to drop_time
    CAST(drop_timestamp AS TIME) AS drop_time
FROM uber_request_data
-- Filter the records where request_date is greater than '2016-06-01' and drop_time is less than 6 AM.
WHERE request_date > '2016-06-01'
    AND drop_time < '06:00';
```

## String functions
```
-- Convert status to lowercase
SELECT LOWER(status) FROM uber_request_data;

-- Convert pickup_point to uppercase
SELECT UPPER(pickup_point) FROM uber_request_data;

-- Complete the CONCAT function for columns pickup_point and status
SELECT CONCAT('Trip from ', pickup_point, ' was completed with status: ', status) AS trip_comment
FROM uber_request_data;
```

## Functions & Grouping
```

```


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
