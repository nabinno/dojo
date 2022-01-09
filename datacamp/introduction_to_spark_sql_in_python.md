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

## Creating context window feature data
```python
# Word for each row, previous two and subsequent two words
query = """
SELECT
part,
LAG(word, 2) OVER(PARTITION BY part ORDER BY id) AS w1,
LAG(word, 1) OVER(PARTITION BY part ORDER BY id) AS w2,
word AS w3,
LEAD(word, 1) OVER(PARTITION BY part ORDER BY id) AS w4,
LEAD(word, 2) OVER(PARTITION BY part ORDER BY id) AS w5
FROM text
"""
spark.sql(query).where("part = 12").show(10)
```

## Repartitioning the data
```python
# Repartition text_df into 12 partitions on 'chapter' column
repart_df = text_df.repartition(12, 'chapter')

# Prove that repart_df has 12 partitions
repart_df.rdd.getNumPartitions()
```

## Finding common word sequences
```python
# Find the top 10 sequences of five words
query = """
SELECT w1, w2, w3, w4, w5, COUNT(*) AS count FROM (
   SELECT word AS w1,
   LEAD(word,1) OVER(PARTITION BY part ORDER BY id ) AS w2,
   LEAD(word,2) OVER(PARTITION BY part ORDER BY id ) AS w3,
   LEAD(word,3) OVER(PARTITION BY part ORDER BY id ) AS w4,
   LEAD(word,4) OVER(PARTITION BY part ORDER BY id ) AS w5
   FROM text
)
GROUP BY w1, w2, w3, w4, w5
ORDER BY count DESC
LIMIT 10
""" 
df = spark.sql(query)
df.show()
```

## Unique 5-tuples in sorted order
```python
# Unique 5-tuples sorted in descending order
query = """
SELECT DISTINCT w1, w2, w3, w4, w5 FROM (
   SELECT word AS w1,
   LEAD(word,1) OVER(PARTITION BY part ORDER BY id ) AS w2,
   LEAD(word,2) OVER(PARTITION BY part ORDER BY id ) AS w3,
   LEAD(word,3) OVER(PARTITION BY part ORDER BY id ) AS w4,
   LEAD(word,4) OVER(PARTITION BY part ORDER BY id ) AS w5
   FROM text
)
ORDER BY w1 DESC, w2 DESC, w3 DESC, w4 DESC, w5 DESC
LIMIT 10
"""
df = spark.sql(query)
df.show()
```

## Most frequent 3-tuples per chapter
```python
#   Most frequent 3-tuple per chapter
query = """
SELECT chapter, w1, w2, w3, count FROM
(
  SELECT
  chapter,
  ROW_NUMBER() OVER (PARTITION BY chapter ORDER BY count DESC) AS row,
  w1, w2, w3, count
  FROM ( %s )
)
WHERE row = 1
ORDER BY chapter ASC
""" % subquery

spark.sql(query).show()
```



# 3. Caching, Logging, and the Spark UI
## Practicing caching: part 1
```python
# Unpersists df1 and df2 and initializes a timer
prep(df1, df2) 

# Cache df1
df1.cache()

# Run actions on both dataframes
run(df1, "df1_1st") 
run(df1, "df1_2nd")
run(df2, "df2_1st")
run(df2, "df2_2nd", elapsed=True)

# Prove df1 is cached
print(df1.is_cached)
```

## Practicing caching: the SQL
```python
# Unpersist df1 and df2 and initializes a timer
prep(df1, df2) 

# Persist df2 using memory and disk storage level 
df2.persist(pyspark.StorageLevel.MEMORY_AND_DISK)

# Run actions both dataframes
run(df1, "df1_1st") 
run(df1, "df1_2nd") 
run(df2, "df2_1st") 
run(df2, "df2_2nd", elapsed=True)
```

## Caching and uncaching tables
```python
# List the tables
print("Tables:\n", spark.catalog.listTables())

# Cache table1 and Confirm that it is cached
spark.catalog.cacheTable('table1')
print("table1 is cached: ", spark.catalog.isCached('table1'))

# Uncache table1 and confirm that it is uncached
spark.catalog.uncacheTable('table1')
print("table1 is cached: ", spark.catalog.isCached('table1'))
```

## Inspecting cache in the Spark UI
```python
partitioned_df.count()
spark.sql("select count(*) from text")
```

## Practice logging
```python
# Log columns of text_df as debug message
logging.debug("text_df columns: %s", text_df.columns)

# Log whether table1 is cached as info message
logging.info("table1 is cached: %s", spark.catalog.isCached(tableName="table1"))

# Log first row of text_df as warning message
logging.warning("The first row of text_df:\n %s", text_df.first())

# Log selected columns of text_df as error message
logging.error("Selected columns: %s", text_df.select("id", "word"))
```

## Practice logging 2
```python
# Uncomment the 5 statements that do NOT trigger text_df
logging.debug("text_df columns: %s", text_df.columns)
logging.info("table1 is cached: %s", spark.catalog.isCached(tableName="table1"))
# logging.warning("The first row of text_df: %s", text_df.first())
logging.error("Selected columns: %s", text_df.select("id", "word"))
logging.info("Tables: %s", spark.sql("show tables").collect())
logging.debug("First row: %s", spark.sql("SELECT * FROM table1 limit 1"))
# logging.debug("Count: %s", spark.sql("SELECT COUNT(*) AS count FROM table1").collect())
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


