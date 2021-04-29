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
# Complete the SELECT statement
data = pd.read_sql("""
SELECT * FROM "Customer"
INNER JOIN "Order"
ON "Order"."customer_id"="Customer"."id"
""", db_engine)

# Show the id column of data
print(data.id)
```

## From task to subtasks
```python
# Function to apply a function over multiple cores
@print_timing
def parallel_apply(apply_func, groups, nb_cores):
    with Pool(nb_cores) as p:
        results = p.map(apply_func, groups)
    return pd.concat(results)

# Parallel apply using 1 core
parallel_apply(take_mean_age, athlete_events.groupby('Year'), 1)

# Parallel apply using 2 cores
parallel_apply(take_mean_age, athlete_events.groupby('Year'), 2)

# Parallel apply using 4 cores
parallel_apply(take_mean_age, athlete_events.groupby('Year'), 4)
```

## Using a DataFrame
```python
import dask.dataframe as dd

# Set the number of pratitions
athlete_events_dask = dd.from_pandas(athlete_events, npartitions = 4)

# Calculate the mean Age per Year
print(athlete_events_dask.groupby('Year').Age.mean().compute())
```

## A PySpark groupby
```python
# Print the type of athlete_events_spark
print(type(athlete_events_spark))

# Print the schema of athlete_events_spark
print(athlete_events_spark.printSchema())

# Group by the Year, and find the mean Age
print(athlete_events_spark.groupBy('Year').mean('Age'))

# Group by the Year, and find the mean Age
print(athlete_events_spark.groupBy('Year').mean('Age').show())
```

## Running PySpark files
```python
from pyspark.sql import SparkSession

if __name__ == "__main__":
    spark = SparkSession.builder.getOrCreate()
    athlete_events_spark = (spark
        .read
        .csv("/home/repl/datasets/athlete_events.csv",
             header=True,
             inferSchema=True,
             escape='"'))

    athlete_events_spark = (athlete_events_spark
        .withColumn("Height",
                    athlete_events_spark.Height.cast("integer")))

    print(athlete_events_spark
        .groupBy('Year')
        .mean('Height')
        .orderBy('Year')
        .show())
```

## Airflow DAGs
```python
# Create the DAG object
dag = DAG(dag_id="car_factory_simulation",
          default_args={"owner": "airflow","start_date": airflow.utils.dates.days_ago(2)},
          schedule_interval="0 * * * *")

# Task definitions
assemble_frame = BashOperator(task_id="assemble_frame", bash_command='echo "Assembling frame"', dag=dag)
place_tires = BashOperator(task_id="place_tires", bash_command='echo "Placing tires"', dag=dag)
assemble_body = BashOperator(task_id="assemble_body", bash_command='echo "Assembling body"', dag=dag)
apply_paint = BashOperator(task_id="apply_paint", bash_command='echo "Applying paint"', dag=dag)

# Complete the downstream flow
assemble_frame.set_downstream(place_tires)
assemble_frame.set_downstream(assemble_body)
assemble_body.set_downstream(apply_paint)
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

