---
title: Introduction to Object-Oriented Programming in Python
tags: python,object-oriented-programming
url: https://campus.datacamp.com/courses/introduction-to-object-oriented-programming-in-python
---

# 1. OOP Fundamentals
## OOP terminology
```
True:
- Methods encode the behavior of an object and are represented by funcitons.
- Attributes encode the state of an object.
- A key benefit of OOP is the bundling of data and methods.

False:
- `.head()` is an example of an attribute of a DataFrame object.
- Object is an obstract template describing the general states and behaviors.
- A class is an instance of an object.
```

## Exploring objects and classes
```python
ratio = 12 / 8

# List all attributes and methods for the ratio object
print(dir(ratio))

# List all attributes and methods for the float class
print(dir(float))
```

## Understanding class definitions
```python
class MyCounter:
  def set_count(self, n):
      self.n = n

mc = MyCounter()
mc.set_count(5)
print(mc.n)
```

## Create your first class
```python

```

## Adding methods and attributes
```python

```

## Extending a class
```python

```

## Class anatomy: the __init__ constructor
```python

```

## Correct use of __init__
```python

```

## Add a class constructor
```python

```

## Building a class from scratch
```python

```




# 2. Inheritance and Polymorphism
## Class vs. instance attributes
```python

```

## Class-level attributes
```python

```

## Implementing logic for attributes
```python

```

## Changing class attributes
```python

```

## Class methods
```python

```

## Adding an alternative constructor
```python

```

## Building a BetterDate Class
```python

```

## Class inheritance
```python

```

## Create a subclass
```python

```

## Understanding inheritance
```python

```

## Customizing functionality via inheritance
```python

```

## Customize a subclass
```python

```

## Method inheritance
```python

```

## Inheritance of class attributes
```python

```




# 3. Integrating with Standard Python
## Operator overloading: comparing objects
```python

```

## Overloading equality
```python

```

## Checking class equality
```python

```

## Inheritance comparison and string representation
```python

```

## Object representation
```python

```

## Comparison and inheritance
```python

```

## String representation of objects
```python

```

## Exceptions
```python

```

## Catching exceptions
```python

```

## Custom exceptions
```python

```
