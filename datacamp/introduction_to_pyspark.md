---
title: Introduction to PySpark
tags: python,database
url: https://www.datacamp.com/courses/introduction-to-pyspark
---

# 1. Getting to know PySpark
## Examining The SparkContext
```python
# Verify SparkContext
print(sc)

# Print Spark version
print(sc.version)
```

## Creating a SparkSession
```python
# Import SparkSession from pyspark.sql
from pyspark.sql import SparkSession

# Create my_spark
my_spark = SparkSession.builder.getOrCreate()

# Print my_spark
print(my_spark)
```

## Viewing tables
```python
# Print the tables in the catalog
print(spark.catalog.listTables())
```

## Are you query-ious?
```python
# Don't change this query
query = "FROM flights SELECT * LIMIT 10"

# Get the first 10 rows of flights
flights10 = spark.sql(query)

# Show the results
flights10.show()
```

## Pandafy a Spark DataFrame
```python
# Don't change this query
query = "SELECT origin, dest, COUNT(*) as N FROM flights GROUP BY origin, dest"

# Run the query
flight_counts = spark.sql(query)

# Convert the results to a pandas DataFrame
pd_counts = flight_counts.toPandas()

# Print the head of pd_counts
print(pd_counts.head())
```

## Put some Spark in your data
```python
# Create pd_temp
pd_temp = pd.DataFrame(np.random.random(10))

# Create spark_temp from pd_temp
spark_temp = spark.createDataFrame(pd_temp)

# Examine the tables in the catalog
print(spark.catalog.listTables())

# Add spark_temp to the catalog
spark_temp.createOrReplaceTempView('temp')

# Examine the tables in the catalog again
print(spark_temp)
```

## Dropping the middle man
```python
# Don't change this file path
file_path = "/usr/local/share/datasets/airports.csv"

# Read in the airports data
airports = spark.read.csv(file_path, header=True)

# Show the data
airports.show()
```

## Creating columns
```python
# Create the DataFrame flights
flights = spark.table("flights")

# Show the head
flights.show()

# Add duration_hrs
flights = flights.withColumn("duration_hrs", flights.air_time/60)
```

## Filtering Data
```python
# Filter flights by passing a string
long_flights1 = flights.filter("distance > 1000")

# Filter flights by passing a column of boolean values
long_flights2 = flights.filter(flights.distance > 1000)

# Print the data to check they're equal
long_flights1.show()
long_flights2.show()
```

## Selecting
```python
# Select the first set of columns
selected1 = flights.select('tailnum', 'origin', 'dest')

# Select the second set of columns
temp = flights.select(flights.origin, flights.dest, flights.carrier)

# Define first filter
filterA = flights.origin == "SEA"

# Define second filter
filterB = flights.dest == "PDX"

# Filter the data, first by filterA then by filterB
selected2 = temp\
    .filter(filterA)\
    .filter(filterB)
```

## Selecting II
```python
# Define avg_speed
avg_speed = (flights.distance/(flights.air_time/60)).alias("avg_speed")

# Select the correct columns
speed1 = flights.select("origin", "dest", "tailnum", avg_speed)

# Create the same table using a SQL expression
speed2 = flights\
    .selectExpr(
        "origin",
        "dest",
        "tailnum",
        "distance/(air_time/60) as avg_speed"
    )
```

# 2. Manipulating data
## Aggregating
```python
# Find the shortest flight from PDX in terms of distance
flights\
    .filter(flights.origin == "PDX")\
	.groupBy()\
	.min("distance")\
	.show()

# Find the longest flight from SEA in terms of air time
flights\
    .filter(flights.origin == "SEA")\
	.groupBy()\
	.max("air_time")\
	.show()
```

## Aggregating II
```python
# Average duration of Delta flights
flights\
    .filter(flights.carrier == "DL")\
    .filter(flights.origin == "SEA")\
    .groupBy()\
    .avg("air_time")\
    .show()

# Total hours in the air
flights\
    .withColumn("duration_hrs", flights.air_time/60)\
    .groupBy()\
    .sum("duration_hrs")\
    .show()
```

## Grouping and aggregating I
```python
# Group by tailnum
by_plane = flights.groupBy("tailnum")

# Number of flights each plane made
by_plane.count().show()

# Group by origin
by_origin = flights.groupBy("origin")

# Average duration of flights from PDX and SEA
by_origin.avg("air_time").show()
```

## Grouping and aggregating II
```python
# Import pyspark.sql.functions as F
import pyspark.sql.functions as F

# Group by month and dest
by_month_dest = flights.groupBy('month', 'dest')

# Average departure delay by month and destination
by_month_dest.avg('dep_delay').show()

# Standard deviation of departure delay
by_month_dest.agg(F.stddev('dep_delay')).show()
```

## Joining II
```python
# Examine the data
print(airports.show())

# Rename the faa column
airports = airports.withColumnRenamed("faa", "dest")

# Join the DataFrames
flights_with_airports = flights.join(airports, on='dest', how='leftouter')

# Examine the new DataFrame
print(flights_with_airports.show())
```

# 3. Getting started with machine learning pipelines
## Join the DataFrame
```python
# Rename year column
planes = planes.withColumnRenamed('year', 'plane_year')

# Join the DataFrames
model_data = flights.join(planes, on='tailnum', how="leftouter")
```

## String to integer
```python
# Cast the columns to integers
model_data = model_data.withColumn("arr_delay", model_data.arr_delay.cast('integer'))
model_data = model_data.withColumn("air_time", model_data.air_time.cast('integer'))
model_data = model_data.withColumn("month", model_data.month.cast('integer'))
model_data = model_data.withColumn("plane_year", model_data.plane_year.cast('integer'))
```

## Create a new column
```python
# Create the column plane_age
model_data = model_data.withColumn("plane_age", model_data.year - model_data.plane_year)
```

## Making a Boolean
```python
# Create is_late
model_data = model_data.withColumn("is_late", model_data.arr_delay > 0)

# Convert to an integer
model_data = model_data.withColumn("label", model_data.is_late.cast('integer'))

# Remove missing values
model_data = model_data.filter("\
arr_delay IS NOT NULL AND \
dep_delay IS NOT NULL AND \
air_time IS NOT NULL AND \
plane_year IS NOT NULL\
")
```

## Carrier
```python
# Create a StringIndexer
carr_indexer = StringIndexer(inputCol='carrier', outputCol='carrier_index')

# Create a OneHotEncoder
carr_encoder = OneHotEncoder(inputCol='carrier_index', outputCol='carrier_fact')
```

## Destination
```python
# Create a StringIndexer
dest_indexer = StringIndexer(inputCol='dest', outputCol='dest_index')

# Create a OneHotEncoder
dest_encoder = OneHotEncoder(inputCol='dest_index', outputCol='dest_fact')
```

## Assemble a vector
```python
# Make a VectorAssembler
vec_assembler = VectorAssembler(inputCols=["month", "air_time", "carrier_fact", "dest_fact", "plane_age"], outputCol='features')
```

## Create the pipeline
```python
# Import Pipeline
from pyspark.ml import Pipeline

# Make the pipeline
flights_pipe = Pipeline(stages=[
    dest_indexer,
    dest_encoder,
    carr_indexer,
    carr_encoder,
    vec_assembler
])
```

## Transform the data
```python
# Fit and transform the data
piped_data = flights_pipe.fit(model_data).transform(model_data)
```

## Split the data
```python
# Split the data into training and test sets
training, test = piped_data.randomSplit([.6, .4])
```

# 4. Model tuning and selection
## Create the modeler
```python
# Import LogisticRegression
from pyspark.ml.classification import LogisticRegression

# Create a LogisticRegression Estimator
lr = LogisticRegression()
```

## Create the evaluator
```python
# Import the evaluation submodule
import pyspark.ml.evaluation as evals

# Create a BinaryClassificationEvaluator
evaluator = evals.BinaryClassificationEvaluator(metricName='areaUnderROC')
```

## Make a grid
```python
# Import the tuning submodule
import pyspark.ml.tuning as tune

# Create the parameter grid
grid = tune.ParamGridBuilder()

# Add the hyperparameter
grid = grid.addGrid(lr.regParam, np.arange(0, .1, .01))
grid = grid.addGrid(lr.elasticNetParam, [0, 1])

# Build the grid
grid = grid.build()
```

## Make the validator
```python
# Create the CrossValidator
cv = tune.CrossValidator(
    estimator=lr,
    estimatorParamMaps=grid,
    evaluator=evaluator
)
```

## Fit the model(s)
```python
# Call lr.fit()
best_lr = lr.fit(training)

# Print best_lr
print(best_lr)
```

## Evaluate the model
```python
# Use the model to predict the test set
test_results = best_lr.transform(test)

# Evaluate the predictions
print(evaluator.evaluate(test_results))
```
