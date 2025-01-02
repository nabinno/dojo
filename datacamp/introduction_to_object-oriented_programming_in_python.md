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
class Employee:
  def __init__(self, name, salary=0):
    self.name = name
    # Check if salary is positive
    if salary >= 0:
      self.salary = salary
    else:
      self.salary = 0
      print("Invalid salary!")

  def give_raise(self, amount):
    self.salary += amount

  def monthly_salary(self):
    return self.salary / 12

emp = Employee("Korel Rossi", -1000)
print(emp.name)
print(emp.salary)

# Define and initialize the Calculator class
class Calculator:
  def __init__(self, num_one, num_two):
    self.num_one = num_one
    self.num_two = num_two

  # Create the addition method
  def addition(self):
    return self.num_one + self.num_two

  # Create the subtraction method
  def subtraction(self):
    return self.num_one - self.num_two

  # Create the multiplication method
  def multiplication(self):
    return self.num_one * self.num_two
```




# 2. Inheritance and Polymorphism
## Class-level attributes
```python
# Create a Player class
class Player:
  # Create MAX_POSITION class attribute
  MAX_POSITION = 10

  # Add a constructor, setting position to zero
  def __init__(self):
    self.position = 0

# Create a player p and print its MAX_POSITION
p = Player()
print(p.MAX_POSITION)
```

## Implementing logic for attributes
```python
class Player:
  MAX_POSITION = 10

  # Define a constructor
  def __init__(self, position):

    # Check if position is less than the class-level attribute value
    if position <= Player.MAX_POSITION:
      self.position = position

    # If not, set equal to the class-level attribute
    else:
      self.position = Player.MAX_POSITION

# Create a Player object, p, and print its MAX_POSITITON
p = Player(6)
print(p.MAX_POSITION)
```

## Changing class attributes
```python
# Create Players p1 and p2
p1 = Player(9)
p2 = Player(5)

print("MAX_POSITION of p1 and p2 before assignment:")
# Print p1.MAX_POSITION and p2.MAX_POSITION
print(p1.MAX_POSITION)
print(p2.MAX_POSITION)

# Assign 7 to p1.MAX_POSITION
p1.MAX_POSITION = 7

print("MAX_POSITION of p1 and p2 after assignment:")
# Print p1.MAX_POSITION and p2.MAX_POSITION
print(p1.MAX_POSITION)
print(p2.MAX_POSITION)
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
