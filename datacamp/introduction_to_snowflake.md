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
-- Retrieve order_id, pizza_id and sum of quantity
SELECT order_id, pizza_id, SUM(quantity) AS total_quantity
FROM order_details
-- GROUP the orders using group by all having total_quantity greater than 3
GROUP BY ALL
HAVING total_quantity > 3
-- ORDER BY order id and total quantity in descending order
ORDER BY order_id, total_quantity DESC;
```

## DATE & TIME
```
-- Select the current date, current time
-- Concatenate and convert it to TIMESTAMP
SELECT *,
    -- Extract month and alias to concat_month
    EXTRACT(month FROM CONCAT(CURRENT_DATE, ' ', CURRENT_TIME)::TIMESTAMP) AS concat_month
-- Use table uber_request_data where request_timestamp's month is greater than concat_month
FROM uber_request_data
WHERE EXTRACT(month FROM request_timestamp) >= concat_month;
```



# 3. Advance Snowflake SQL Concepts
## NATURAL JOIN
```
SELECT
    pt.category,
    -- Calculate total_revenue
    SUM(p.price * od.quantity) total_revenue
FROM order_details AS od
    -- NATURAL JOIN all tables
    NATURAL JOIN pizzas AS p
    NATURAL JOIN pizza_type AS pt
-- GROUP the records by category from pizza_type table
GROUP BY pt.category
-- ORDER by total_revenue and limit the records
ORDER BY total_revenue DESC
LIMIT 1;
```

## The world of JOINS
```
SELECT COUNT(o.order_id) AS total_orders,
    AVG(p.price) AS average_price,
    -- Calculate total revenue
    SUM(p.price * od.quantity) AS total_revenue,
    -- Get the name from pizza_type table
    pt.name AS pizza_name
FROM orders AS o
    -- Use appropriate JOIN
    LEFT JOIN order_details AS od
        ON o.order_id = od.order_id
    -- Use appropriate JOIN with pizzas table
    RIGHT JOIN pizzas p
        ON od.pizza_id = p.pizza_id
    -- Natural join pizza_type table
    NATURAL LEFT JOIN pizza_type pt
GROUP BY pt.name, pt.category
ORDER BY total_revenue desc, total_orders desc;
```

## LATERAL JOIN
```
SELECT pt.name,
    pt.category,
    o.order_date,
    -- Get max quantity from lateral query
    x.max_quantity
FROM pizzas AS pz
    JOIN pizza_type AS pt ON pz.pizza_type_id = pt.pizza_type_id
    JOIN order_details AS od ON pz.pizza_id = od.pizza_id
    -- Join with orders table
    JOIN orders AS o ON od.order_id = o.order_id,
LATERAL (
    -- Select max of order_details quantity
    SELECT MAX(od2.quantity) AS max_quantity
    FROM order_details AS od2
        -- Join with pizzas table
        JOIN pizzas AS pz2
        ON od2.pizza_id = pz2.pizza_id
    -- Filtering condition for the subquery
    WHERE pz2.pizza_type_id = pz.pizza_type_id
) AS x
WHERE od.quantity = x.max_quantity
GROUP BY pt.name, pt.category, o.order_date, x.max_quantity
ORDER BY pt.name;
```

## Subqueries
```
SELECT
    pt.name,
    pt.category,
    SUM(od.quantity) AS total_quantity
FROM pizza_type AS pt
    -- Join pizzas and order_details table
    JOIN pizzas AS pz
        ON pt.pizza_type_id = pz.pizza_type_id
    JOIN order_details AS od
        ON pz.pizza_id = od.pizza_id
    -- Group by name and category
GROUP BY pt.name, pt.category
HAVING SUM(od.quantity) < (
    -- Calculate AVG of total_quantity
    SELECT AVG(total_quantity)
    FROM (
        SELECT SUM(od2.quantity) AS total_quantity
        FROM pizzas AS pz2
            JOIN order_details AS od2
                ON pz2.pizza_id = od2.pizza_id
        GROUP BY pz2.pizza_type_id
    ) AS sub
)
-- Order  by total_quantity in ascending order
ORDER BY total_quantity ASC
```

## Understanding CTE
```
WITH size_sales_cte AS (
    SELECT
        p.pizza_size AS size,
        COUNT(*) AS total_sales
    FROM pizzas AS p
        JOIN order_details AS od
            ON p.pizza_id = od.pizza_id
    GROUP BY p.pizza_size
)
SELECT
    size,
    total_sales
FROM size_sales_cte
ORDER BY total_sales DESC
LIMIT 1;
```

## CTEs
```
-- Create a CTE named most_ordered and limit the results
WITH most_ordered AS (
    SELECT pizza_id, SUM(quantity) AS total_qty
    FROM order_details GROUP BY pizza_id ORDER BY total_qty DESC
    LIMIT 1
)
-- Create CTE cheapest_pizza where price is equals to min price from pizzas table
, cheapest_pizza AS (
    SELECT pizza_id, price
    FROM pizzas
    WHERE price = (SELECT MIN(price) FROM pizzas)
    LIMIT 1
)
-- Select pizza_id and total_qty aliased as metric from first cte most_ordered
SELECT pizza_id, 'Most Ordered' AS Description, total_qty AS metric
FROM most_ordered
UNION ALL
-- Select pizza_id and price aliased as metric from second cte cheapest_pizza
SELECT pizza_id, 'Cheapest' AS Description, price AS metric
FROM cheapest_pizza
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
