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
# Create an empty class Employee
class Employee:
      pass
# Create an object emp of class Employee
emp = Employee()

# Create new class Employee
class Employee:
  # Include a set_name method
  def set_name(self, new_name):
    self.name = new_name
emp = Employee()
# Use set_name() on emp to set the name of emp to 'Korel Rossi'
emp.set_name('Korel Rossi')
print(emp.name)
```

## Adding methods and attributes
```python
class Employee:
  def set_name(self, new_name):
    self.name = new_name

  # Add set_salary() method
  def set_salary(self, new_salary):
    self.salary = new_salary

emp = Employee()

# Use set_name to set the name of emp to 'Korel Rossi'
emp.set_name('Korel Rossi')

# Set the salary of emp to 50000
emp.set_salary(50000)

# Print the emp object's salary
print(emp.salary)
```

## Extending a class
```python
class Employee:
    def set_name(self, new_name):
      self.name = new_name

    def set_salary(self, new_salary):
      self.salary = new_salary

    # Add a give_raise() method with amount as an argument
    def give_raise(self, amount):
      self.salary = self.salary + amount

# Create the emp object
emp = Employee()
emp.set_name('Korel Rossi')
emp.set_salary(50000)

# Print the salary
print(emp.salary)

# Give emp a raise of 1500
emp.give_raise(1500)
print(emp.salary)
```

## Correct use of __init__
```python
class Counter:
    def __init__(self, count, name):
      self.count = count
      self.name = name

c = Counter(0, "My counter")
print(c.count)
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
