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

## Data Preparation for Spark ALS
```python

```

## Correct format and distinct users
```python

```

## Assigning integer id's to movies
```python

```

## ALS Parameters and Hyperparameters
```python

```

## Build Out An ALS Model
```python

```

## Build RMSE Evaluator
```python

```

## Get RMSE
```python

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


