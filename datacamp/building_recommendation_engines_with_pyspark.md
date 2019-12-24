---
title: Building Recommendation Engines with PySpark
tags: python,recommender-system,machine-learning
url: https://www.datacamp.com/courses/recommendation-engines-in-pyspark
---

# 1. Recommendations Are Everywhere
## See the power of a recommendation engine
```python
# View TJ_ratings
TJ_ratings.show()

# Generate recommendations for users
get_ALS_recs(["Jane","Taylor"]) 
```

## Implicit vs Explicit Data
```python
# Type "implicit" or "explicit"
answer = "implicit"
```

## Ratings data types
```python
# Group the data by "Genre"
markus_ratings.groupBy("Genre").sum().show()
```

## Confirm understanding of latent features
```python
# Examine matrix P using the .show() method
P.show()

# Examine matrix Pi using the .show() method
Pi.show()
```

# 2. How does ALS work?
## Matrix Multiplication
```python
# Use the .head() method to view the contents of matrices a and b
print("Matrix A:")
print (a.head())

print("Matrix B:")
print (b.head())

# Complete the matrix with the product of matrices a and b
product = np.array([[10,12], [15,18]])

# Run this validation to see how your estimate performs
product == np.dot(a,b)
```

## Matrix Multiplication Part II
```python
# Print the dimensions of C
print(C.shape)

# Print the dimensions of D
print(D.shape)

# Can C and D be multiplied together?
C_times_D = None
```

## Matrix Factorization
```python
# Take a look at Matrix G using the following print function
print("Matrix G:")
print(G)

# Take a look at the matrices H, I, and J and determine which pair of those matrices will produce G when multiplied together. 
print("Matrix H:")
print(H)
print("Matrix I:")
print(I)
print("Matrix J:")
print(J)

# Multiply the two matrices that are factors of the matrix G
prod = np.matmul(H, J)
print(G == prod)
```

## Non-Negative Matrix Factorization
```python
# View the L, U, W, and H matrices.
print("Matrices L and U:") 
print(L)
print(U)

print("Matrices W and H:")
print(W)
print(H)

# Calculate RMSE between LU and M
print("RMSE of LU: ", getRMSE(LU, M))

# Calculate RMSE between WH and M
print("RMSE of WH: ", getRMSE(WH, M))
```

## Estimating Recommendations
```python
# View left factor matrix
print(U)

# View left factor matrix
print(U)

# Multiply factor matrices
UP = np.matmul(U,P)

# Convert to pandas DataFrame
print(pd.DataFrame(UP, columns=P.columns, index=U.index))
```

## RMSE As ALS Alternates
```python
# Use getRMSE(preds, actuals) to calculate the RMSE of matrices T and F1.
getRMSE(T, F1)

# Create list of F2, F3, F4, F5, and F6
Fs = [F2, F3, F4, F5, F6]

# Calculate RMSE for F2, F3, F4, F5, and F6.
getRMSEs(Fs, T)
```

## Correct format and distinct users
```python
# Import monotonically_increasing_id and show R
from pyspark.sql.functions import monotonically_increasing_id
R.show()

# Use the to_long() function to convert the dataframe to the "long" format.
ratings = to_long(R)
ratings.show()

# Get unique users and repartition to 1 partition
users = ratings.select("User").distinct().coalesce(1)

# Create a new column of unique integers called "userId" in the users dataframe.
users = users\
    .withColumn("userId", monotonically_increasing_id())\
    .persist()
users.show()
```

## Assigning integer id's to movies
```python
# Extract the distinct movie id's
movies = ratings.select("Movie").distinct() 

# Repartition the data to have only one partition.
movies = movies.coalesce(1)

# Create a new column of movieId integers. 
movies = movies\
    .withColumn("movieId", monotonically_increasing_id())\
    .persist() 

# Join the ratings, users and movies dataframes
movie_ratings = ratings\
    .join(users, "User", "left")\
    .join(movies, "Movie", "left")
movie_ratings.show()
```

## Build Out An ALS Model
```python
# Split the ratings dataframe into training and test data
(training_data, test_data) = ratings.randomSplit([0.8, 0.2], seed=42)

# Set the ALS hyperparameters
from pyspark.ml.recommendation import ALS
als = ALS(userCol="userId", itemCol="movieId", ratingCol="rating", rank=10, maxIter =15, regParam=0.1,
          coldStartStrategy="drop", nonnegative=True, implicitPrefs=False)

# Fit the mdoel to the training_data
model = als.fit(training_data)

# Generate predictions on the test_data
test_predictions = model.transform(test_data)
test_predictions.show()
```

## Build RMSE Evaluator
```python
# Import RegressionEvaluator
from pyspark.ml.evaluation import RegressionEvaluator

# Complete the evaluator code
evaluator = RegressionEvaluator(metricName="rmse", labelCol="ratings", predictionCol="prediction")

# Extract the 3 parameters
print(evaluator.getMetricName())
print(evaluator.getLabelCol())
print(evaluator.getPredictionCol())
```

## Get RMSE
```python
# Evaluate the "test_predictions" dataframe
RMSE = evaluator.evaluate(test_predictions)

# Print the RMSE
print(RMSE)
```

# 3. Recommending Movies
## Introduction to the MovieLens dataset
```python

```

## Viewing the MovieLens Data
```python

```

## Calculate sparsity
```python

```

## The GroupBy and Filter Methods
```python

```

## MovieLens Summary Statistics
```python

```

## View Schema
```python

```

## ALS model buildout on MovieLens Data
```python

```

## Create test/train splits and build your ALS model
```python

```

## Tell Spark how to tune your ALS model
```python

```

## Build your cross validation pipeline
```python

```

## Best Model and Best Model Parameters
```python

```

## Model Performance Evaluation
```python

```

## Generate predictions and calculate RMSE
```python

```

## Interpreting the RMSE
```python

```

## Do Recommendations Make Sense
```python

```

# 4. What if you don't have customer ratings?
## Introduction to the Million Songs Dataset
```python

```

## Confirm understanding of implicit rating concepts
```python

```

## MSD summary statistics
```python

```

## Grouped summary statistics
```python

```

## Add Zeros
```python

```

## Evaluating Implicit Ratings Models
```python

```

## Specify ALS Hyperparameters
```python

```

## Build Implicit Models
```python

```

## Running a Cross-Validated Implicit ALS Model
```python

```

## Extracting Parameters
```python

```

## Overview of binary, implicit ratings
```python

```

## Binary Model Performance
```python

```

## Recommendations From Binary Data
```python

```

## Course Recap
```python

```


