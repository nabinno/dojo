---
title: Introduction to Deep Learning with Keras
tags: python,machine-learning,deep-learning
url: https://campus.datacamp.com/courses/introduction-to-deep-learning-with-keras/introducing-keras
---

# 1. Introducing Keras
## Describing Keras
```python
False - Keras can work well on its own without using a backend, like TensorFlow.
=> Keras is a wrapper around a backend library, so a backend like TensorFlow, Theano, CNTK, etc must be provided.
```

## Would you use deep learning?
```python
Imagine you're building an app that allows you to take a picture of your clothes and then shows you a pair of shoes that would match well. This app needs a machine learning module that's in charge of identifying the type of clothes you are wearing, as well as their color and texture. Would you use deep learning to accomplish this task?
=> I'd use deep learning since we are dealing with unstructured data and neural networks work well with images.
```

## Hello nets!
```python
# Import the Sequential model and Dense layer
from tensorflow.keras.models import Sequential
from tensorflow.keras.layers import Dense

# Create a Sequential model
model = Sequential()

# Add an input layer and a hidden layer with 10 neurons
model.add(Dense(10, input_shape=(2,), activation="relu"))

# Add a 1-neuron output layer
model.add(Dense(1))

# Summarise your model
model.summary()
```

## Counting parameters
```python
# Instantiate a new Sequential model
model = Sequential()

# Add a Dense layer with five neurons and three inputs
model.add(Dense(5, input_shape=(3,), activation="relu"))

# Add a final Dense layer with one neuron and no activation
model.add(Dense(1))

# Summarize your model
model.summary()

<script.py> output:
    Model: "sequential"
    _________________________________________________________________
    Layer (type)                 Output Shape              Param #   
    =================================================================
    dense (Dense)                (None, 5)                 20        
    _________________________________________________________________
    dense_1 (Dense)              (None, 1)                 6         
    =================================================================
    Total params: 26
    Trainable params: 26
    Non-trainable params: 0
    _________________________________________________________________
```

## Build as shown!
```python
from tensorflow.keras.models import Sequential
from tensorflow.keras.layers import Dense

# Instantiate a Sequential model
model = Sequential()

# Build the input and hidden layer
model.add(Dense(3, input_shape=(2,)))

# Add the ouput layer
model.add(Dense(1))
```

## Specifying a model
```python
# Instantiate a Sequential model
model = Sequential()

# Add a Dense layer with 50 neurons and an input of 1 neuron
model.add(Dense(50, input_shape=(1,), activation='relu'))

# Add two Dense layers with 50 neurons and relu activation
model.add(Dense(50, activation='relu'))
model.add(Dense(50, activation='relu'))

# End your model with a Dense layer and no activation
model.add(Dense(1))
```

## Training
```python

```

## Predicting the orbit!
```python

```




# 2. Going Deeper
## Binary classification
```python

```

## Exploring dollar bills
```python

```

## A binary classification model
```python

```

## Is this dollar bill fake ?
```python

```

## Multi-class classification
```python

```

## A multi-class model
```python

```

## Prepare your dataset
```python

```

## Training on dart throwers
```python

```

## Softmax predictions
```python

```

## Multi-label classification
```python

```

## An irrigation machine
```python

```

## Training with multiple labels
```python

```

## Keras callbacks
```python

```

## The history callback
```python

```

## Early stopping your model
```python

```

## A combination of callbacks
```python

```




# 3. 
## Learning curves
```python

```

## Learning the digits
```python

```

## Is the model overfitting?
```python

```

## Do we need more data?
```python

```

## Activation functions
```python

```

## Different activation functions
```python

```

## Comparing activation functions
```python

```

## Comparing activation functions II
```python

```

## Batch size and batch normalization
```python

```

## Changing batch sizes
```python

```

## Batch normalizing a familiar model
```python

```

## Batch normalization effects
```python

```

## Hyperparameter tuning
```python

```

## Preparing a model for tuning
```python

```

## Tuning the model parameters
```python

```

## Training with cross-validation
```python

```




# 4. Advanced Model Architectures
## Tensors, layers, and autoencoders
```python

```

## It's a flow of tensors
```python

```

## Neural separation
```python

```

## Building an autoencoder
```python

```

## De-noising like an autoencoder
```python

```

## Intro to CNNs
```python

```

## Building a CNN model
```python

```

## Looking at convolutions
```python

```

## Preparing your input image
```python

```

## Using a real world model
```python

```

## Intro to LSTMs
```python

```

## Text prediction with LSTMs
```python

```

## Build your LSTM model
```python

```

## Decode your predictions
```python

```

## Test your model!
```python

```

## You're done!
```python

```

