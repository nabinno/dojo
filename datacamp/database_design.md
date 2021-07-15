---
title: Database Design
tags: database,database-design
url: https://www.datacamp.com/courses/database-design
---

# 1. Processing, Storing, and Organizing Data
## OLAP vs. OLTP
```
OLAP
- Queries a larger ammount of data
- Helps businesses with decision making and problem solving
- Typically uses a data warehouse

OLTP
- Most likely to have data from the past hour
- Data is inserted and updated more often
- Typically uses an operational database
```

## Which is better?
```sql
SELECT * FROM Potholes;
```

## Name that data type!
```
Unstructured
- To-do notes in a text editor
- Images in your photo library
- Zip file of all text messages ever received

Semi-Structured
- JSON object of tweets outputted in real-time by the Twitter API
- XML
- CSVs of open data downloaded from your local government websites

Structured
- A relational database with latest withdrawals and deposits made by clients
```

## Ordering ETL Tasks
```
- eCommerce API outputs real time data of transactions
- Python script drops null rows and clean data into pre-determined columns
- Resulting dataframe is written into an AWS Redshift Warehouse
```

## Classifying data models
```
Conceptual Data Model
- Entities, attributes, and relationships
- Gathers business requirements

Logical Data Model
- Relational model
- Determining tables and columns

Physical Data Model
- File structure of data storage
```

## Deciding fact and dimension tables
```
-- Create a route dimension table
CREATE TABLE route(
	route_id INTEGER PRIMARY KEY,
    park_name VARCHAR(160) NOT NULL,
    city_name VARCHAR(160) NOT NULL,
    distance_km FLOAT NOT NULL,
    route_name VARCHAR(160) NOT NULL
);
-- Create a week dimension table
CREATE TABLE week(
	week_id INTEGER PRIMARY KEY,
    week INTEGER NOT NULL,
    month VARCHAR(160) NOT NULL,
    year INTEGER NOT NULL
);
```

## Querying the dimensional model
```
SELECT 
	-- Get the total duration of all runs
	SUM(duration_mins)
FROM 
	runs_fact
-- Get all the week_id's that are from July, 2019
INNER JOIN week_dim ON week_dim.week_id = runs_fact.week_id
WHERE week_dim.month = 'July' and week_dim.year = '2019';
```



# 2. Database Schemas and Normalization
## Adding foreign keys
```
-- Add the book_id foreign key
ALTER TABLE fact_booksales ADD CONSTRAINT sales_book
    FOREIGN KEY (book_id) REFERENCES dim_book_star (book_id);
    
-- Add the time_id foreign key
ALTER TABLE fact_booksales ADD CONSTRAINT sales_time
    FOREIGN KEY (time_id) REFERENCES dim_time_star (time_id);
    
-- Add the store_id foreign key
ALTER TABLE fact_booksales ADD CONSTRAINT sales_store
    FOREIGN KEY (store_id) REFERENCES dim_store_star (store_id);
```

## Extending the book dimension
```

```

## Normalized and denormalized databases
```

```

## Querying the star schema
```

```

## Querying the snowflake schema
```

```

## Updating countries
```

```

## Extending the snowflake schema
```

```

## Normal forms
```

```

## Converting to 1NF
```

```

## Converting to 2NF
```

```

## Converting to 3NF
```

```




# 3. Database Views
## Database views
```

```

## Tables vs. views
```

```

## Viewing views
```

```

## Creating and querying a view
```

```

## Managing views
```

```

## Creating a view from other views
```

```

## Granting and revoking access
```

```

## Updatable views
```

```

## Redefining a view
```

```

## Materialized views
```

```

## Materialized versus non-materialized
```

```

## Creating and refreshing a materialized view
```

```

## Managing materialized views
```

```




# 4. Database Mangement
## Database roles and access control
```

```

## Create a role
```

```

## GRANT privileges and ALTER attributes
```

```

## Add a user role to a group role
```

```

## Table partitioning
```

```

## Reasons to partition
```

```

## Partitioning and normalization
```

```

## Creating vertical partitions
```

```

## Creating horizontal partitions
```

```

## Data integration
```

```

## Data integration do's and dont's
```

```

## Analyzing a data integration plan
```

```

## Picking a Database Management System (DBMS)
```

```

## SQL versus NoSQL
```

```

## Choosing the right DBMS
```

```


