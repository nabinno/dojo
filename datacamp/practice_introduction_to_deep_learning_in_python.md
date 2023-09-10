---
title: Practice: Introduction to Deep Learning in Python
tags: deep-learning, machine-learning, python
url: https://practice.datacamp.com/p/26
---

## 1. Which import statement below will allow you to reopen/use models you create?
```python
from keras.models import load_model
```

## 2.
```python
model = Sequential()
add_layers(model)
model.compile(optimizer='adam',
    loss='categorical_crossentropy')
print(model.layers)

#=> [<keras.layers.core.Dense object at 0x7f607703ee48>, <keras.layers.core.Dense object at 0x7f60770424e0>]
```

## 3.
```python
model = Sequential()
add_layers(model)
model.compile(optimizer='adam',
    loss='categorical_crossentropy',
    metrics=['accuracy'])
print(model.layers)
```

## 4. What is the correct way to initialize a neural netwrok model?
```python
[x] model=Sequential()
[ ] model=Sequential
[ ] model=sequential
[ ] model=sequential()
```

## 5.
```txt
Which of the following options describes a logical flow when starting to fit a neural network?
[x]Start with small network, get validation score, if validation score is still getting better, increase model capacity
[ ]Start with large network, get validation score, if validation score is still getting better, increase model capacity
[ ]Start with large network, get validation score, if validation score is still getting better, decrease model capacity
[ ]Start with small network, get validation score, if validation score is still getting better, decrease model capacity
```
