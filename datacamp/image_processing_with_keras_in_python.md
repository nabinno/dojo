---
title: Image Processing with Keras in Python
tags: keras,image-processing
url: https://www.datacamp.com/courses/convolutional-neural-networks-for-image-processing
---

# 1. Image Processing With Neural Networks
## Images as data: visualizations
```python
# Import matplotlib
import matplotlib.pyplot as plt

# Load the image
data = plt.imread('bricks.png')

# Display the image
plt.imshow(data)
plt.show()
```

## Images as data: changing images
```python
# Set the red channel in this part of the image to 1
data[:10, :10, 0] = 1

# Set the green channel in this part of the image to 0
data[:10, :10, 1] = 0

# Set the blue channel in this part of the image to 0
data[:10, :10, 2] = 0

# Visualize the result
plt.imshow(data)
plt.show()
```

## Using one-hot encoding to represent images
```python
# The number of image categories
n_categories = 3

# The unique values of categories in the data
categories = np.array(["shirt", "dress", "shoe"])

# Initialize ohe_labels as all zeros
ohe_labels = np.zeros((len(labels), n_categories))

# Loop over the labels
for ii in range(len(labels)):
    # Find the location of this label in the categories variable
    jj = np.where(categories == labels[ii])
    # Set the corresponding zero to one
    ohe_labels[ii, jj] = 1
```

## Evaluating a classifier
```python
# Calculate the number of correct predictions
number_correct = (test_labels * predictions).sum()
print(number_correct)

# Calculate the proportion of correct predictions
proportion_correct = number_correct/len(predictions)
print(proportion_correct)
```

## Build a neural network
```python
# Imports components from Keras
from keras.models import Sequential
from keras.layers import Dense

# Initializes a sequential model
model = Sequential()

# First layer
model.add(Dense(10, activation='relu', input_shape=(784,)))

# Second layer
model.add(Dense(10, activation='relu'))

# Output layer
model.add(Dense(3, activation='softmax'))
```

## Compile a neural network
```python
# Compile the model
model.compile(optimizer='adam',
           loss='categorical_crossentropy', 
           metrics=['accuracy'])
```

## Fitting a neural network model to clothing data
```python
# Reshape the data to two-dimensional array
train_data = train_data.reshape(50, 784)

# Fit the model
model.fit(train_data, train_labels, validation_split=0.2, epochs=3)
```

## Cross-validation for neural network evaluation
```python
# Reshape test data
test_data = test_data.reshape(10, 28 * 28)

# Evaluate the model
model.evaluate(test_data, test_labels)
```

# 2. Using Convolutions
## Convolutions
```python

```

## One dimensional convolutions
```python

```

## Image convolutions
```python

```

## Defining image convolution kernels
```python

```

## Implementing image convolutions in Keras
```python

```

## Convolutional network for image classification
```python

```

## Training a CNN to classify clothing types
```python

```

## Evaluating a CNN with test data
```python

```

## Tweaking your convolutions
```python

```

## Add padding to a CNN
```python

```

## Add strides to a convolutional network
```python

```

## Calculate the size of convolutional layer output
```python

```

# 3. Going Deeper
## Going deeper
```python

```

## Creating a deep learning network
```python

```

## Train a deep CNN to classify clothing images
```python

```

## What is special about a deep network?
```python

```

## How many parameters?
```python

```

## How many parameters in a CNN?
```python

```

## How many parameters in a deep CNN?
```python

```

## Pooling operations
```python

```

## Write your own pooling operation
```python

```

## Keras pooling layers
```python

```

## Train a deep CNN with pooling to classify images
```python

```

# 4. Understanding and Improving Deep Convolutional Networks
## Tracking learning
```python

```

## Plot the learning curves
```python

```

## Using stored weights to predict in a test set
```python

```

## Regularization
```python

```

## Adding dropout to your network
```python

```

## Add batch normalization to your network
```python

```

## Interpreting the model
```python

```

## Extracting a kernel from a trained network
```python

```

## Shape of the weights
```python

```

## Visualizing kernel responses
```python

```

## Next steps
```python

```
