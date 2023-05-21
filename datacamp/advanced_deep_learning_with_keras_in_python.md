---
title: Advanced Deep Learning with Keras in Python
tags: python,machine-learning
url: https://campus.datacamp.com/courses/advanced-deep-learning-with-keras/the-keras-functional-api
---

# 1. The Keras Functional API
## Input layers
```python
# Import Input from tensorflow.keras.layers
from tensorflow.keras.layers import Input

# Create an input layer of shape 1
input_tensor = Input(shape=(1,))
```

## Dense layers
```python
# Load layers
from tensorflow.keras.layers import Input, Dense

# Input layer
input_tensor = Input(shape=(1,))

# Dense layer
output_layer = Dense(1)

# Connect the dense layer to the input_tensor
output_tensor = output_layer(input_tensor)
```

## Output layers
```python
# Load layers
from tensorflow.keras.layers import Input, Dense

# Input layer
input_tensor = Input(shape=(1,))

# Create a dense layer and connect the dense layer to the input_tensor in one step
# Note that we did this in 2 steps in the previous exercise, but are doing it in one step now
output_tensor = Dense(1)(input_tensor)
```

## Build a model
```python
# Input/dense/output layers
from tensorflow.keras.layers import Input, Dense
input_tensor = Input(shape=(1,))
output_tensor = Dense(1)(input_tensor)

# Build the model
from tensorflow.keras.models import Model
model = Model(input_tensor, output_tensor)
```

## Compile a model
```python

```

## Visualize a model
```python

```

## Fit and evaluate a model
```python

```

## Fit the model to the tournament basketball data
```python

```

## Evaluate the model on a test set
```python

```




# 2. Two Input Networks Using Categorical Embeddings, Shared Layers, and Merge Layers
## Category embeddings
```python

```

## Define team lookup
```python

```

## Define team model
```python

```

## Shared layers
```python

```

## Defining two inputs
```python

```

## Lookup both inputs in the same model
```python

```

## Merge layers
```python

```

## Output layer using shared layer
```python

```

## Model using two inputs and one output
```python

```

## Predict from your model
```python

```

## Fit the model to the regular season training data
```python

```

## Evaluate the model on the tournament test data
```python

```




# 3. Multiple Inputs: 3 Inputs (and Beyond!)
## Three-input models
```python

```

## Make an input layer for home vs. away
```python

```

## Make a model and compile it
```python

```

## Fit the model and evaluate
```python

```

## Summarizing and plotting models
```python

```

## Model summaries
```python

```

## Plotting models
```python

```

## Stacking models
```python

```

## Add the model predictions to the tournament data
```python

```

## Create an input layer with multiple columns
```python

```

## Fit the model
```python

```

## Evaluate the model
```python

```




# 4. Multiple Outputs
## Two-output models
```python

```

## Simple two-output model
```python

```

## Fit a model with two outputs
```python

```

## Inspect the model (I)
```python

```

## Evaluate the model
```python

```

## Single model for classification and regression
```python

```

## Classification and regression in one model
```python

```

## Compile and fit the model
```python

```

## Inspect the model (II)
```python

```

## Evaluate on new data with two metrics
```python

```

## Wrap-up
```python

```




