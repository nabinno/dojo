---
title: Introduction to Spark SQL in Python
tags: python,database,apache-spark
url: https://campus.datacamp.com/courses/introduction-to-spark-sql-in-python
---

# 1. Pyspark SQL
## Create a SQL table from a dataframe
```python
# Load trainsched.txt
df = spark.read.csv("trainsched.txt", header=True)

# Create temporary table called table1
df.createOrReplaceTempView("table1")
```

## Determine the column names of a table
```python
# Inspect the columns in the table df
spark.sql("DESCRIBE schedule").show()
```

## Running sums using window function SQL
```python
# Add col running_total that sums diff_min col in each group
query = """
SELECT train_id, station, time, diff_min,
SUM(diff_min) OVER (PARTITION BY train_id ORDER BY time) AS running_total
FROM schedule
"""

# Run the query and display the result
spark.sql(query).show()
```

## Fix the broken query
```python
query = """
SELECT 
ROW_NUMBER() OVER (ORDER BY time) AS row,
train_id, 
station, 
time, 
LEAD(time,1) OVER (ORDER BY time) AS time_next 
FROM schedule
"""
spark.sql(query).show()

# Give the number of the bad row as an integer
bad_row = 7

# Provide the missing clause, SQL keywords in upper case
clause = 'PARTITION BY train_id'
```

## Aggregation, step by step
```python
# Give the identical result in each command
spark.sql('SELECT train_id, MIN(time) AS start FROM schedule GROUP BY train_id').show()
df.groupBy('train_id').agg({'time':'min'}).withColumnRenamed('MIN(time)', 'start').show()

# Print the second column of the result
spark.sql('SELECT train_id, MIN(time), MAX(time) FROM schedule GROUP BY train_id').show()
result = df.groupBy('train_id').agg({'time':'min', 'time':'max'})
result.show()
print(result.columns[0])
```

## Aggregating the same column twice
```python
# Write a SQL query giving a result identical to dot_df
query = "SELECT train_id, MIN(time) AS start, MAX(time) AS end FROM schedule GROUP BY train_id"
sql_df = spark.sql(query)
sql_df.show()
```

## Aggregate dot SQL
```python
# Obtain the identical result using dot notation 
dot_df = df.withColumn('time_next', lead('time', 1)
        .over(Window.partitionBy('train_id')
        .orderBy('time')))
```

## Convert window function from dot notation to SQL
```python
# Create a SQL query to obtain an identical result to dot_df
query = """
SELECT *, 
(UNIX_TIMESTAMP(LEAD(time, 1) OVER (PARTITION BY train_id ORDER BY time),'H:m') 
 - UNIX_TIMESTAMP(time, 'H:m'))/60 AS diff_min 
FROM schedule 
"""
sql_df = spark.sql(query)
sql_df.show()
```



# 2. Using window function sql for natural language processing
## Loading a dataframe from a parquet file
```python
# Load the dataframe
df = spark.read.load('sherlock_sentences.parquet')

# Filter and show the first 5 rows
df.where('id > 70').show(5, truncate=False)
```

## Split and explode a text column
```python
# Split the clause column into a column called words 
split_df = clauses_df.select(split('clause', ' ').alias('words'))
split_df.show(5, truncate=False)

# Explode the words column into a column called word 
exploded_df = split_df.select(explode('words').alias('word'))
exploded_df.show(10)

# Count the resulting number of rows in exploded_df
print("\nNumber of rows: ", exploded_df.count())
```

## Using monotonically_increasing_id()
```python
pyspark.sql.function.monotonically_increasing_id()
```

## Moving window analysis
```python

```

## Creating context window feature data
```python

```

## Repartitioning the data
```python

```

## Common word sequences
```python

```

## What type of data is this
```python

```

## Finding common word sequences
```python

```

## Unique 5-tuples in sorted order
```python

```

## Most frequent 3-tuples per chapter
```python

```





# 3. Caching, Logging, and the Spark UI
## Caching
```python

```

## Practicing caching: part 1
```python

```

## Practicing caching: the SQL
```python

```

## Practicing caching: putting it all together
```python

```

## Caching and uncaching tables
```python

```

## The Spark UI
```python

```

## Spark UI storage tab
```python

```

## Inspecting cache in the Spark UI
```python

```

## Logging
```python

```

## Practice logging
```python

```

## Practice logging 2
```python

```

## Query plans
```python

```

## Practice query plans
```python

```

## Practice reading query plans 2
```python

```




# 4. Text classification
## Extract Transform Select
```python

```

## Practicing creating a UDF
```python

```

## Practicing array column
```python

```

## Creating feature data for classification
```python

```

## Creating a UDF for vector data
```python

```

## Applying a UDF to vector data
```python

```

## Transforming text to vector format
```python

```

## Text Classification
```python

```

## Label the data
```python

```

## Split the data
```python

```

## Train the classifier
```python

```

## Predicting and evaluating
```python

```

## Evaluate the classifier
```python

```

## Predict test data
```python

```

## Recap
```python

```


