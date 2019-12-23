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
## Data Preparation: Removing columns and rows
```python
# Remove the 'flight' column
flights = flights.drop('flight')

# Number of records with missing 'delay' values
flights.filter('delay IS NULL').count()

# Remove records with missing 'delay' values
flights = flights.filter('delay IS NOT NULL')

# Remove records with missing values in any column and get the number of remaining rows
flights = flights.dropna()
print(flights.count())
```

## Data Preparation: Column manipulation
```python
# Import the required function
from pyspark.sql.functions import round

# Convert 'mile' to 'km' and drop 'mile' column
flights_km = flights\
    .withColumn('km', round(flights.mile *  1.60934, 0))\
    .drop('mile')

# Create 'label' column indicating whether flight delayed (1) or not (0)
flights_km = flights_km\
    .withColumn('label', (flights_km.delay >= 15).cast('integer'))

# Check first five records
flights_km.show(5)
```

## Data Preparation: Categorical columns
```python
from pyspark.ml.feature import StringIndexer

# Create an indexer
indexer = StringIndexer(inputCol='carrier', outputCol='carrier_idx')

# Indexer identifies categories in the data
indexer_model = indexer.fit(flights)

# Indexer creates a new column with numeric index values
flights_indexed = indexer_model.transform(flights)

# Repeat the process for the other categorical feature
flights_indexed = StringIndexer(inputCol='org', outputCol='org_idx')\
    .fit(flights_indexed)\
    .transform(flights_indexed)
```

## Data Preparation: Assembling columns
```python
# Import the necessary class
from pyspark.ml.feature import VectorAssembler

# Create an assembler object
assembler = VectorAssembler(
    inputCols=['mon', 'dom', 'dow', 'carrier_idx', 'org_idx', 'km', 'depart', 'duration'],
    outputCol='features'
)

# Consolidate predictor columns
flights_assembled = assembler.transform(flights)

# Check the resulting column
flights_assembled.select('features', 'delay').show(5, truncate=False)
```

## Decision Tree: Train/test split
```python
# Split into training and testing sets in a 80:20 ratio
flights_train, flights_test = flights.randomSplit([0.8, 0.2], 17)

# Check that training set has around 80% of records
training_ratio = flights_train.count() / flights.count()
print(training_ratio)
```

## Decision Tree: Build a Decision Tree
```python
# Import the Decision Tree Classifier class
from pyspark.ml.classification import DecisionTreeClassifier

# Create a classifier object and fit to the training data
tree = DecisionTreeClassifier()
tree_model = tree.fit(flights_train)

# Create predictions for the testing data and take a look at the predictions
prediction = tree_model.transform(flights_test)
prediction.select('label', 'prediction', 'probability').show(5, False)
```

## Decision Tree: Evaluate the Decision Tree
```python
# Create a confusion matrix
prediction.groupBy('label', 'prediction').count().show()

# Calculate the elements of the confusion matrix
TN = prediction.filter('prediction = 0 AND label = prediction').count()
TP = prediction.filter('prediction = 1 AND label = prediction').count()
FN = prediction.filter('prediction = 0 AND label != prediction').count()
FP = prediction.filter('prediction = 1 AND label != prediction').count()

# Accuracy measures the proportion of correct predictions
accuracy = (TN + TP) / (TN + TP + FN + FP)
print(accuracy)
```

## Logistic Regression: Build a Logistic Regression model
```python
# Import the logistic regression class
from pyspark.ml.classification import LogisticRegression

# Create a classifier object and train on training data
logistic = LogisticRegression().fit(flights_train)

# Create predictions for the testing data and show confusion matrix
prediction = logistic.transform(flights_test)
prediction.groupBy('label', 'prediction').count().show()
```

## Logistic Regression: Evaluate the Logistic Regression model
```python
from pyspark.ml.evaluation import MulticlassClassificationEvaluator, BinaryClassificationEvaluator

# Calculate precision and recall
precision = TP / (TP + FP)
recall = TP / (TP + FN)
print('precision = {:.2f}\nrecall    = {:.2f}'.format(precision, recall))

# Find weighted precision
multi_evaluator = MulticlassClassificationEvaluator()
weighted_precision = multi_evaluator.evaluate(prediction, {multi_evaluator.metricName: "weightedPrecision"})

# Find AUC
binary_evaluator = BinaryClassificationEvaluator()
auc = binary_evaluator.evaluate(prediction, {binary_evaluator.metricName: "areaUnderROC"})
```

## Turning Text into Tables: Punctuation, numbers and tokens
```python
# Import the necessary functions
from pyspark.sql.functions import regexp_replace
from pyspark.ml.feature import Tokenizer

# Remove punctuation (REGEX provided) and numbers
wrangled = sms.withColumn('text', regexp_replace(sms.text, '[_():;,.!?\\-]', ' '))
wrangled = wrangled.withColumn('text', regexp_replace(wrangled.text, '[0-9]', ' '))

# Merge multiple spaces
wrangled = wrangled.withColumn('text', regexp_replace(wrangled.text, ' +', ' '))

# Split the text into words
wrangled = Tokenizer(inputCol='text', outputCol='words').transform(wrangled)

wrangled.show(4, truncate=False)
```

## Turning Text into Tables: Stop words and hashing
```python
from pyspark.ml.feature import StopWordsRemover, HashingTF, IDF

# Remove stop words.
wrangled = StopWordsRemover(inputCol='words', outputCol='terms')\
    .transform(sms)

# Apply the hashing trick
wrangled = HashingTF(inputCol='terms', outputCol='hash', numFeatures=1024)\
    .transform(wrangled)

# Convert hashed symbols to TF-IDF
tf_idf = IDF(inputCol='hash', outputCol='features')\
    .fit(wrangled)\
    .transform(wrangled)
      
tf_idf.select('terms', 'features').show(4, truncate=False)
```

## Turning Text into Tables: Training a spam classifier
```python
# Split the data into training and testing sets
sms_train, sms_test = sms.randomSplit([0.8, 0.2], 13)

# Fit a Logistic Regression model to the training data
logistic = LogisticRegression(regParam=0.2).fit(sms_train)

# Make predictions on the testing data
prediction = logistic.transform(sms_test)

# Create a confusion matrix, comparing predictions to known labels
prediction\
    .groupBy('label', 'prediction')\
    .count()\
    .show()
```

# 3. Regression
## One-Hot Encoding: Encoding flight origin
```python
# Import the one hot encoder class
from pyspark.ml.feature import OneHotEncoderEstimator

# Create an instance of the one hot encoder
onehot = OneHotEncoderEstimator(inputCols=['org_idx'], outputCols=['org_dummy'])

# Apply the one hot encoder to the flights data
onehot = onehot.fit(flights)
flights_onehot = onehot.transform(flights)

# Check the results
flights_onehot\
	.select('org', 'org_idx', 'org_dummy')\
	.distinct()\
	.sort('org_idx')\
	.show()
```

## Regression: Flight duration model: Just distance
```python
from pyspark.ml.regression import LinearRegression
from pyspark.ml.evaluation import RegressionEvaluator

# Create a regression object and train on training data
regression = LinearRegression(labelCol='duration').fit(flights_train)

# Create predictions for the testing data and take a look at the predictions
predictions = regression.transform(flights_test)
predictions.select('duration', 'prediction').show(5, False)

# Calculate the RMSE
RegressionEvaluator(labelCol='duration').evaluate(predictions)
```

## Regression: Interpreting the coefficients
```python

```

## Regression: Flight duration model: Adding origin airport
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
