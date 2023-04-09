---
title: Introduction to TensorFlow in Python
tags: python,machine-learning
url: https://campus.datacamp.com/courses/introduction-to-tensorflow-in-python
---

# 1. Introduction to TensorFlow
## Defining data as constants
```python
# Import constant from TensorFlow
from tensorflow import constant

# Convert the credit_numpy array into a tensorflow constant
credit_constant = constant(credit_numpy)

# Print constant datatype
print('\n The datatype is:', credit_constant.dtype)

# Print constant shape
print('\n The shape is:', credit_constant.shape)
```

## Defining variables
```python
# Define the 1-dimensional variable A1
A1 = Variable([1, 2, 3, 4])

# Print the variable A1
print('\n A1: ', A1)

# Convert A1 to a numpy array and assign it to B1
B1 = A1.numpy()

# Print B1
print('\n B1: ', B1)
```

## Basic operations
```python

```

## Performing element-wise multiplication
```python

```

## Making predictions with matrix multiplication
```python

```

## Summing over tensor dimensions
```python

```

## Advanced operations
```python

```

## Reshaping tensors
```python

```

## Optimizing with gradients
```python

```

## Working with image data
```python

```




# 2. Linear models
## Input data
```python

```

## Load data using pandas
```python

```

## Setting the data type
```python

```

## Loss functions
```python

```

## Loss functions in TensorFlow
```python

```

## Modifying the loss function
```python

```

## Linear regression
```python

```

## Set up a linear regression
```python

```

## Train a linear model
```python

```

## Multiple linear regression
```python

```

## Batch training
```python

```

## Preparing to batch train
```python

```

## Training a linear model in batches
```python

```




# 3. Neural Networks
## Dense layers
```python

```

## The linear algebra of dense layers
```python

```

## The low-level approach with multiple examples
```python

```

## Using the dense layer operation
```python

```

## Activation functions
```python

```

## Binary classification problems
```python

```

## Multiclass classification problems
```python

```

## Optimizers
```python

```

## The dangers of local minima
```python

```

## Avoiding local minima
```python

```

## Training a network in TensorFlow
```python

```

## Initialization in TensorFlow
```python

```

## Defining the model and loss function
```python

```

## Training neural networks with TensorFlow
```python

```




# 4. High Level APIs
## Defining neural networks with Keras
```python

```

## The sequential model in Keras
```python

```

## Compiling a sequential model
```python

```

## Defining a multiple input model
```python

```

## Training and validation with Keras
```python

```

## Training with Keras
```python

```

## Metrics and validation with Keras
```python

```

## Overfitting detection
```python

```

## Evaluating models
```python

```

## Training models with the Estimators API
```python

```

## Preparing to train with Estimators
```python

```

## Defining Estimators
```python

```

## Congratulations!
```python

```


