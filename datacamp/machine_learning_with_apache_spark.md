---
title: Machine Learning with Apache Spark
tags: python,machine-learning
url: https://campus.datacamp.com/courses/machine-learning-with-apache-spark
---

# 1. Introduction
## Creating a SparkSession
```python
# Import the PySpark module
from pyspark.sql import SparkSession

# Create SparkSession object
spark = SparkSession\
    .builder\
    .master('local[*]')\
    .appName('test')\
    .getOrCreate()

# What version of Spark?
print(spark.version)

# Terminate the cluster
spark.stop()
```

## Loading flights data
```python
# Read data from CSV file
flights = spark.read.csv(
    'flights.csv',
    sep=',',
    header=True,
    inferSchema=True,
    nullValue='NA'
)

# Get number of records
print("The data contain %d records." % flights.count())

# View the first five records
flights.show(5)

# Check column data types
flights.dtypes
```

## Loading SMS spam data
```python
from pyspark.sql.types import StructType, StructField, IntegerType, StringType

# Specify column names and types
schema = StructType([
    StructField("id", IntegerType()),
    StructField("text", StringType()),
    StructField("label", IntegerType())
])

# Load data from a delimited file
sms = spark.read.csv('sms.csv', sep=';', header=False, schema=schema)

# Print schema of DataFrame
sms.printSchema()
```

# 2. Classification
## Data Preparation
```python

```

## Removing columns and rows
```python

```

## Column manipulation
```python

```

## Categorical columns
```python

```

## Assembling columns
```python

```

## Decision Tree
```python

```

## Train/test split
```python

```

## Build a Decision Tree
```python

```

## Evaluate the Decision Tree
```python

```

## Logistic Regression
```python

```

## Build a Logistic Regression model
```python

```

## Evaluate the Logistic Regression model
```python

```

## Turning Text into Tables
```python

```

## Punctuation, numbers and tokens
```python

```

## Stop words and hashing
```python

```

## Training a spam classifier
```python

```

# 3. Regression
## One-Hot Encoding
```python

```

## Encoding flight origin
```python

```

## Encoding shirt sizes
```python

```

## Regression
```python

```

## Flight duration model: Just distance
```python

```

## Interpreting the coefficients
```python

```

## Flight duration model: Adding origin airport
```python

```

## Interpreting coefficients
```python

```

## Bucketing & Engineering
```python

```

## Bucketing departure time
```python

```

## Flight duration model: Adding departure time
```python

```

## Regularization
```python

```

## Flight duration model: More features!
```python

```

## Flight duration model: Regularisation!
```python

```

# 4. Ensembles & Pipelines
## Pipeline
```python

```

## Flight duration model: Pipeline stages
```python

```

## Flight duration model: Pipeline model
```python

```

## SMS spam pipeline
```python

```

## Cross-Validation
```python

```

## Cross validating simple flight duration model
```python

```

## Cross validating flight duration model pipeline
```python

```

## Grid Search
```python

```

## Optimizing flights linear regression
```python

```

## Dissecting the best flight duration model
```python

```

## SMS spam optimised
```python

```

## How many models for grid search?
```python

```

## Ensemble
```python

```

## Delayed flights with Gradient-Boosted Trees
```python

```

## Delayed flights with a Random Forest
```python

```

## Evaluating Random Forest
```python

```

## Closing thoughts
```python

```
