---
title: Introduction to Deep Learning in Python
tags: deep-learning, machine-learning, python
url: https://www.datacamp.com/courses/introduction-to-deep-learning-in-python
---

# 1. Basics of deep learning and neural networks
## Coding the forward propagation algorithm
```python
# Calculate node 0 value: node_0_value
node_0_value = (input_data * weights['node_0']).sum()

# Calculate node 1 value: node_1_value
node_1_value = (input_data * weights['node_1']).sum()

# Put node values into array: hidden_layer_outputs
hidden_layer_outputs = np.array([node_0_value, node_1_value])

# Calculate output: output
output = (hidden_layer_outputs * weights['output']).sum()

# Print output
print(output)
```

## The Rectified Linear Activation Function
```python
def relu(input):
    '''Define your relu activation function here'''
    # Calculate the value for the output of the relu function: output
    output = max(0, input)
    
    # Return the value just calculated
    return(output)

# Calculate node 0 value: node_0_output
node_0_input = (input_data * weights['node_0']).sum()
node_0_output = relu(node_0_input)

# Calculate node 1 value: node_1_output
node_1_input = (input_data * weights['node_1']).sum()
node_1_output = relu(node_1_input)

# Put node values into array: hidden_layer_outputs
hidden_layer_outputs = np.array([node_0_output, node_1_output])

# Calculate model output (do not apply relu)
model_output = (hidden_layer_outputs * weights['output']).sum()

# Print model output
print(model_output)
```

## Applying the network to many observations/rows of data
```python
# Define predict_with_network()
def predict_with_network(input_data_row, weights):

    # Calculate node 0 value
    node_0_input = (input_data_row * weights['node_0']).sum()
    node_0_output = relu(node_0_input)

    # Calculate node 1 value
    node_1_input = (input_data_row * weights['node_1']).sum()
    node_1_output = relu(node_1_input)

    # Put node values into array: hidden_layer_outputs
    hidden_layer_outputs = np.array([node_0_output, node_1_output])
    
    # Calculate model output
    input_to_final_layer = (hidden_layer_outputs * weights['output']).sum()
    model_output = relu(input_to_final_layer)
    
    # Return model output
    return(model_output)


# Create empty list to store prediction results
results = []
for input_data_row in input_data:
    # Append prediction to results
    results.append(predict_with_network(input_data_row, weights))

# Print results
print(results)
```

## Deeper networks
```python

```

## Forward propagation in a deeper network
```python

```

## Multi-layer neural networks
```python

```

## Representations are learned
```python

```

## Levels of representation
```python

```




# 2. Optimizing a neural network with backward propagation
## The need for optimization
```python

```

## Calculating model errors
```python

```

## Understanding how weights change model accuracy
```python

```

## Coding how weight changes affect accuracy
```python

```

## Scaling up to multiple data points
```python

```

## Gradient descent
```python

```

## Calculating slopes
```python

```

## Improving model weights
```python

```

## Making multiple updates to weights
```python

```

## Backpropagation
```python

```

## The relationship between forward and backward propagation
```python

```

## Thinking about backward propagation
```python

```

## Backpropagation in practice
```python

```

## A round of backpropagation
```python

```




# 3. Building deep learning models with keras
## Creating a keras model
```python

```

## Understanding your data
```python

```

## Specifying a model
```python

```

## Compiling and fitting a model
```python

```

## Compiling the model
```python

```

## Fitting the model
```python

```

## Classification models
```python

```

## Understanding your classification data
```python

```

## Last steps in classification models
```python

```

## Using models
```python

```

## Making predictions
```python

```




# 4. Fine-tuning keras models
## Understanding model optimization
```python

```

## Diagnosing optimization problems
```python

```

## Changing optimization parameters
```python

```

## Model validation
```python

```

## Evaluating model accuracy on validation dataset
```python

```

## Early stopping: Optimizing the optimization
```python

```

## Experimenting with wider networks
```python

```

## Adding layers to a network
```python

```

## Thinking about model capacity
```python

```

## Experimenting with model structures
```python

```

## Stepping up to images
```python

```

## Building your own digit recognition model
```python

```

## Final thoughts
```python

```

