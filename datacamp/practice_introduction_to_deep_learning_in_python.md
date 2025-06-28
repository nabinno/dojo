---
title: Practice: Introduction to Deep Learning in Python
tags: deep-learning, machine-learning, python
url: https://practice.datacamp.com/p/26
---

## A-1. Which import statement below will allow you to reopen/use models you create?
```python
from keras.models import load_model
```

## A-2.
```python
model = Sequential()
add_layers(model)
model.compile(optimizer='adam',
    loss='categorical_crossentropy')
print(model.layers)

#=> [<keras.layers.core.Dense object at 0x7f607703ee48>, <keras.layers.core.Dense object at 0x7f60770424e0>]
```

## A-3.
```python
model = Sequential()
add_layers(model)
model.compile(optimizer='adam',
    loss='categorical_crossentropy',
    metrics=['accuracy'])
print(model.layers)
```

## A-4. What is the correct way to initialize a neural netwrok model?
```python
[x] model=Sequential()
[ ] model=Sequential
[ ] model=sequential
[ ] model=sequential()
```

## A-5. Which of the following options describes a logical flow when starting to fit a neural network?
```txt
[x]Start with small network, get validation score, if validation score is still getting better, increase model capacity
[ ]Start with large network, get validation score, if validation score is still getting better, increase model capacity
[ ]Start with large network, get validation score, if validation score is still getting better, decrease model capacity
[ ]Start with small network, get validation score, if validation score is still getting better, decrease model capacity
```

## A-6. Which import statement below will allow you to reopen/use models you create?
```python
[ ]from keras.models import open_model
[x]from keras.models import load_model
[ ]from keras.models import use_model
```

## A-7.
```python
model = Sequential()
add_layers(model)
model.compile(
    optimizer='adam',
    loss='categorical_crossentropy'
)
print(model.layers)
```

## B-1. Which layer below could be creating the output layer of a regression model?
```python
[ ]model.add(Dense(10))
[ ]model.add(Dense(1, activation = 'relu'))
[ ]model.add(Dense(10, activation = 'relu'))
[x]model.add(Dense(1))
```

## B-2. What is the correct way to initialize a neural network model?
```python
[ ]model=sequential
[ ]model=sequential()
[x]model=Sequential()
[ ]model=Sequential
```

## B-3. Which of the following options describes a logical flow when starting to fit a neural netwrok?
```txt
[ ]Start with small network, get validation score, if validation score is still getting better, decrease model capacity
[x]Start with small network, get validation score, if validation score is still getting better, increase model capacity
[ ]Start with large network, get validation score, if validation score is still getting better, decrease model capacity
[ ]Start with large network, get validation score, if validation score is still getting better, increase model capacity
```

## B-4. Create the first layer of the model with 100 nodes and ReLU as the activation function
```python
model = Sequential()
model.add(Dense(100, activation = 'relu',
    input_shape = (n_cols,)))
print(model.layers)

#=> [<deras.layers.core.Dense object at 0x7f60753653c8>]
```

## B-5. Create the first layer of the model.
```python
model = Sequential()
model.add(Dense(100,
        activation='relu',
        input_shape = (n_cols,)))
print(model.layers[0].name)

#=> dense_1
```

## B-6. Which of the following options describes a logical flow when starting to fit a neural network?
```txt
[ ]Start with large network, get validation score, if validation score is still getting better, decrease model capacity
[x]Start with small network, get validation score, if validation score is still getting better, increase model capacity
[ ]Start with small network, get validation score, if validation score is still getting better, decrease model capacity
[ ]Start with large network, get validation score, if validation score is still getting better, increase model capacity
```

## C-1. `weights`, `input_data`, `preds` and `error` have already been loaded/calculated. Like in backpropagation, calculate the slope of the weights using the loss function, `mean_squared_error`.
```python
slope = 2 * input_data * error
print(slope)

#=> [-84 -56 -70]
```

## C-2. Fill in the code to find the value at `n_0` given `data`
```python
val = (data * weights['n_0']).sum()
print(val)

#=> 7
```

## C-3. You use the ReLU activation function on the two input data points -3, and 7. What is the output?
```txt
[ ]3,7
[ ]0,-7
[x]0,7
[ ]-3,0
```

## C-4. Ensure that `30%` of the initial data is held out of training for validation. Note: The function `make_model` initiates and  compiles a 2-layered classification model.
```python
model = Sequential()
make_model(model)
model.fit(predictors, target,
    verbose = 0, validation_split=0.3)

#=> <keras.callbacks.History at 0x7f6076480dd8>
```

## C-5. Fill in the code to have the same behavior as the ReLU activation function.
```python
def relu(input):
    output = max(0, input)
    return output
output = relu(-2)
print(output)

#=> 0
```

## C-6. Ensure that `30%` of the initial data is held out of training for validation. Note: The function `make_model` initiates and compiles a 2-layered classification model.
```python
model = Sequential()
make_model(model)
model.fit(predictors, target,
    verbose = 0, validation_split=0.3)

#=> <keras.callbacks.History at 0x7f6076480dd8>
```

