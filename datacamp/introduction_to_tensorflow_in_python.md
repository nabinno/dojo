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

## Performing element-wise multiplication
```python
# Define tensors A1 and A23 as constants
A1 = constant([1, 2, 3, 4])
A23 = constant([[1, 2, 3], [1, 6, 4]])

# Define B1 and B23 to have the correct shape
B1 = ones_like(A1)
B23 = ones_like(A23)

# Perform element-wise multiplication
C1 = multiply(A1, B1)
C23 = multiply(A23, B23)

# Print the tensors C1 and C23
print('\n C1: {}'.format(C1.numpy()))
print('\n C23: {}'.format(C23.numpy()))
```

## Making predictions with matrix multiplication
```python
# Define features, params, and bill as constants
features = constant([[2, 24], [2, 26], [2, 57], [1, 37]])
params = constant([[1000], [150]])
bill = constant([[3913], [2682], [8617], [64400]])

# Compute billpred using features and params
billpred = matmul(features, params)

# Compute and print the error
error = bill - billpred
print(error.numpy())
```

## Summing over tensor dimensions
```python
import tensorflow as tf

wealth = tf.constant([[11,50],[7,2],[4,60],[3,0],[25,10]])

# The individual in the first row has the highest total wealth (i.e. stocks + bonds).
total_wealth = tf.reduce_sum(wealth, axis=1)
print(total_wealth.numpy()) # [61  9 64  3 35]
print(total_wealth[0].numpy() > tf.reduce_max(total_wealth[1:]).numpy()) # False

# Combined, the 5 individuals hold $50,000 in stocks.
total_stocks = tf.reduce_sum(wealth[:,1])
print(total_stocks.numpy()) # 122
print(total_stocks.numpy() == 50) # False

# Combined, the 5 individuals hold $50,000 in bonds.
total_bonds = tf.reduce_sum(wealth[:,0])
print(total_bonds.numpy()) # 50
print(total_bonds.numpy() == 50) # True

# The individual in the second row has the lowest total wealth (i.e. stocks + bonds).
print(total_wealth[1].numpy() < tf.reduce_min(total_wealth).numpy()) # False
```

## Reshaping tensors
```python
# Reshape the grayscale image tensor into a vector
gray_vector = reshape(gray_tensor, (28*28, 1))

# Reshape the color image tensor into a vector
color_vector = reshape(color_tensor, (28*28*3, 1))
```

## Optimizing with gradients
```python
def compute_gradient(x0):
  	# Define x as a variable with an initial value of x0
	x = Variable(x0)
	with GradientTape() as tape:
		tape.watch(x)
        # Define y using the multiply operation
		y = multiply(x, x)
    # Return the gradient of y with respect to x
	return tape.gradient(y, x).numpy()

# Compute and print gradients at x = -1, 1, and 0
print(compute_gradient(-1.0))
print(compute_gradient(1.0))
print(compute_gradient(0.0))
```

## Working with image data
```python
# Reshape model from a 1x3 to a 3x1 tensor
model = reshape(model, (3, 1))

# Multiply letter by model
output = matmul(letter, model)

# Sum over output and print prediction using the numpy method
prediction = reduce_sum(output)
print(prediction.numpy())
```




# 2. Linear models
## Load data using pandas
```python
# Import pandas under the alias pd
import pandas as pd

# Assign the path to a string variable named data_path
data_path = 'kc_house_data.csv'

# Load the dataset as a dataframe named housing
housing = pd.read_csv(data_path)

# Print the price column of housing
print(housing['price'])
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


