---
title: Introduction to Data Engineering
tags: python,data-engineering
url: https://www.datacamp.com/courses/introduction-to-data-engineering
---

# 1. Introduction to Data Engineering


# 2. Data engineering toolbox
## The database schema
```python
# Complete the SELECT statement
data = pd.read_sql("""
SELECT first_name, last_name FROM "Customer"
ORDER BY last_name, first_name
""", db_engine)

# Show the first 3 rows of the DataFrame
print(data.head(3))

# Show the info of the DataFrame
print(data.info())
```

## Joining on relations
```python

```

## Star schema diagram
```python

```

## What is parallel computing
```python

```

## Why parallel computing?
```python

```

## From task to subtasks
```python

```

## Using a DataFrame
```python

```

## Parallel computation frameworks
```python

```

## Spark, Hadoop and Hive
```python

```

## A PySpark groupby
```python

```

## Running PySpark files
```python

```

## Workflow scheduling frameworks
```python

```

## Airflow, Luigi and cron
```python

```

## Airflow DAGs
```python

```



# 3. Extract, Transform and Load (ETL)
## Extract
```python

```

## Data sources
```python

```

## Fetch from an API
```python

```

## Read from a database
```python

```

## Transform
```python

```

## Splitting the rental rate
```python

```

## Prepare for transformations
```python

```

## Joining with ratings
```python

```

## Loading
```python

```

## OLAP or OLTP
```python

```

## Writing to a file
```python

```

## Load into Postgres
```python

```

## Putting it all together
```python

```

## Defining a DAG
```python

```

## Setting up Airflow
```python

```

## Interpreting the DAG
```python

```



# 4. Case Study: DataCamp
## Course ratings
```python

```

## Exploring the schema
```python

```

## Querying the table
```python

```

## Average rating per course
```python

```

## From ratings to recommendations
```python

```

## Filter out corrupt data
```python

```

## Using the recommender transformation
```python

```

## Scheduling daily jobs
```python

```

## The target table
```python

```

## Defining the DAG
```python

```

## Enable the DAG
```python

```

## Querying the recommendations
```python

```

## Congratulations
```python

```

