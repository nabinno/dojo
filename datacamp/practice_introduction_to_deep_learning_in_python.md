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
