---
title: Introduction to Snowflake
tags: snowflake,analytics,data-engineering
url: https://campus.datacamp.com/courses/introduction-to-snowflake/
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

## Essentials of query optimization
```
True
- In Snowflake, the more time a query takes, the more resources are consumed, which can lead to higher costs.
- Using the `TOP` or `LIMIT` clause can make your query more efficient by restricting the number of rows returned.
- Using `UNION ALL` instead of `UNION` can speed up your query if you are sure there are no duplicate records.

False
- Filtering data using the `WHERE` clause should be done after performing joins for optimized performance.
- Joining tables without precise conditions, often called 'exploding join,' is a recommended practice for efficient querying.
- It's a good practice to use `SELECT *` to fetch all columns for large tables.
```

## Early filtering
```
WITH filtered_orders AS (
    -- Select order_id, order_date and filter records where order_date is greater than November 1, 2015.
    SELECT order_id, order_date
    FROM orders WHERE order_date > TO_DATE('2015-11-01')
), filtered_pizza_type AS (
    -- Select name, pizza_type_id and filter the pizzas which has Veggie category
    SELECT name, pizza_type_id
    FROM pizza_type WHERE category = 'Veggie'
)
SELECT o.order_id, o.order_date, pt.name, od.quantity
-- Get the details from filtered_orders CTE
FROM filtered_orders AS o
JOIN order_details AS od ON o.order_id = od.order_id
JOIN pizzas AS p ON od.pizza_id = p.pizza_id
-- JOIN CTE filtered_pizza_type on common column
JOIN filtered_pizza_type AS pt ON p.pizza_type_id = pt.pizza_type_id
```

## Query history
```
[ ]SELECT query_id, start_time, query_text, warehouse_name FROM snowflake.information_schema.query_history
[ ]SELECT query_id, start_time, query_text, warehouse_name FROM snowflake.account.query_history
[x]SELECT query_id, start_time, end_time, query_text FROM snowflake.account_usage.query_history
[ ]SELECT query_id, start_time, query_text, warehouse_name FROM snowflake.account_usage.access_history
```

## `PARSE_JSON` & `OBJECT_CONSTRUCT`
```
True:
- This is a valid syntax:
    PARSE_JSON('{
        "business_id": 10,
        "name": "restaurant"
    }')
- `PARSE_JSON(<expr>) returns a valid JSON object.
- `OBJECT_CONSTRUCT` returns a JSON object.

False:
- This is a valid syntax:
    OBJECT_CONSTRUCT('{
        "business_id": 8,
        "name": "restaurant"
    }')
- `PARSE_JSON` returns key-value pairs separated by commas.
- The input to `PARSE_JSON()` is key-value pairs data separated by commas.
```

## Querying JSON data
```
SELECT
    name,
    categories,
    -- Select WheelchairAccessible from attributes converting it to STRING
    attributes:WheelchairAccessible::STRING AS wheelchair_accessible,
    -- Select Saturday, Sunday from hours converting it to STRING
    (hours:Saturday::STRING IS NOT NULL OR hours:Sunday::STRING IS NOT NULL) AS open_on_weekend
FROM yelp_business_data
WHERE
    -- Filter where wheelchair_accessible is 'True' and open_on_weekend is 'true'
    wheelchair_accessible = 'True' AND open_on_weekend = 'true'
    -- Filter further where categories is having Italian in it
    AND categories LIKE '%Italian%'
```

## JSONified
```
WITH dogs_allowed AS (
    SELECT *
    FROM yelp_business_data
    WHERE attributes:DogsAllowed::STRING  NOT ILIKE '%None%'
        AND attributes:DogsAllowed::STRING = 'True'
)

, touristy_places AS (
    SELECT *
    FROM yelp_business_data
    WHERE attributes:Ambience NOT ILIKE '%None%'
        AND attributes:Ambience IS NOT NULL
        AND attributes:Ambience NOT ILIKE '%u''%'
        -- Convert Ambience attribute in the attributes columns into valid JSON using PARSE_JSON.
        -- From Valid JSON, fetch the touristy attribute and check if it is true when casted to BOOLEAN.
        AND PARSE_JSON(attributes:Ambience):touristy::BOOLEAN = TRUE
)

SELECT
    d.business_id,
    d.name
FROM dogs_allowed d
    JOIN touristy_places t
        ON d.business_id = t.business_id
```

---

# 1 Introduction to Snowflake UI
## Navigating to a database and schema
```
What are the schemas in the `COURSE_39090` database?
[ ]INFORM_SCHEME, STUDENT_DB
[x]INFORMATION_SCHEMA, NOTEBOOKS, STUDENT_STORE
[ ]INFORMATION_SCHEMA, DATACAMP_DB
```

## Navigate to a table
```
Snowflake > Data > Databases > COURSE_39090_DB > STUDENT_STORE
Tables:
- SALES_TRANSACTIONS
    - data type of TRANSACTION_ID column: Number
- VIDEO_DATASET
```

## Analyzing the sales amount
```
1. Click button [+ Create > SQL Worksheet] in Snowflake

2. Select STUDENT_STORE schema:
select *
from COURSE_39090_DB_E6831F36335248DFBAE60ECC4F554BC7.STUDENT_STORE.SALES_TRANSACTIONS
limit 5;

3. Calculate the total sales using the `AMOUNT` column
select sum(AMOUNT)
from COURSE_39090_DB_E6831F36335248DFBAE60ECC4F554BC7.STUDENT_STORE.SALES_TRANSACTIONS
limit 5;
-- => 15326.00
```

## Monitoring Query and Copy History
```
1. Create SQL worksheet "Exploratory Analysis" in [Snowflakes > Worksheets]
2. Create SQL worksheet "Product Category Analysis" in [Snowflakes > Worksheets]
3. Calculate the total sales by PRODUCT_CATEGORY.
select
    t.PRODUCT_CATEGORY,
    SUM(t.AMOUNT)
from COURSE_39090_DB_65386DDFBC424F698A2B798C914A8446.STUDENT_STORE.SALES_TRANSACTIONS t
GROUP BY t.PRODUCT_CATEGORY;
-- Result:
PRODUCT_CATEGORY	SUM(T.AMOUNT)
Books	8670.00
Paper	318.00
Electronics	5208.00
Utilities	1130.00
```

## Viewing query Details
```
1. Create new worksheet in [Snowflake > Worksheets]

2. Execute the following query:
SELECT product, amount
FROM COURSE_39090_DB_BAF2B903E9F84297BA8BCD37D8D6BDCE.STUDENT_STORE.SALES_TRANSACTIONS
GROUP BY product, amount
ORDER BY amount DESC;

3. Check the query history in [Snowflake > Monitoring > Query History]
```

## Exploring Python Worksheets
```python
# 1. Navigate to Worksheets in the Projects section to create a Python worksheet.

# 2. Execute the Python code and investigate the results.
import snowflake.snowpark as snowpark
from snowflake.snowpark.functions import col

def main(session: snowpark.Session):
    # Your code goes here, inside the "main" handler.
    tableName = 'COURSE_39090_DB_508DF5A8C30E4D88AB9B0031A6DB2EB7.information_schema.packages'
    dataframe = session.table(tableName).filter(col("language") == 'python')

    # Print a sample of the dataframe to standard output.
    dataframe.show()

    # Return value will appear in the Results tab.
    return dataframe
```




# 2 Snowsight in Action: Managing Data & Access
## Loading Data from Snowflake Marketplace
```

```

## Exploring the Snowflake marketplace
```

```

## Uploading CSV Data
```

```

## Loading Data from a cloud provider
```

```

## Creating a table from an external stage
```

```

## Creating new roles, users, and warehouses
```

```

## Exploring user roles
```

```

## Creating your first notebook
```

```

## Course Wrap-Up
```

```
