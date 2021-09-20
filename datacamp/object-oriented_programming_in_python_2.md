---
title: Object-Oriented Programming in Python
tags: python
url: https://campus.datacamp.com/courses/object-oriented-programming-in-python
---

# 1. OOP Fundamentals
## OOP termininology
```txt
## True
Methods encode behavior of an object and are represented by functions.
Attributes encode the state of an object and are represented by variables.
Encapsulation is a software design practice of bundling the data and the methods that operate on that data.

## False
`.columns` is an example of a method of a DataFrame object.
Object and class are different terms describing the same concept.
Object is an abstract template describing the general states and behaviors.
A programming language can be either object-oriented or procedural, but not both.
```

## Exploring object interface
```python
##
In: type(mystery)
Out: __main__.Employee

##
# Print the mystery employee's name
print(mystery.name)

# Print the mystery employee's salary
print(mystery.salary)

# Give the mystery employee a raise of $2500
mystery.give_raise(2500)

# Print the salary again
print(mystery.salary)
```

## Understanding class definitions
```python
class MyCounter:
  def set_count(self, n):
    self.count = n

mc = MyCounter()
mc.set_count(5)
mc.count = mc.count + 1
print(mc.count)
```

## Create your first class
```python
##
# Create an empty class Employee
class Employee:
    pass

# Create an object emp of class Employee  
emp = Employee()

##
class Employee:
  
  def set_name(self, new_name):
    self.name = new_name
  
  # Add set_salary() method
  def set_salary(self, new_salary):
    self.salary = new_salary
  
  
# Create an object emp of class Employee  
emp = Employee()

# Use set_name to set the name of emp to 'Korel Rossi'
emp.set_name('Korel Rossi')

# Set the salary of emp to 50000
emp.set_salary(50000)
```

## Using attributes in class definition
```python
##
class Employee:
    def set_name(self, new_name):
        self.name = new_name

    def set_salary(self, new_salary):
        self.salary = new_salary 
  
emp = Employee()
emp.set_name('Korel Rossi')
emp.set_salary(50000)

# Print the salary attribute of emp
print(emp.salary)

# Increase salary of emp by 1500
emp.salary = 1500 + emp.salary

# Print the salary attribute of emp again
print(emp.salary)

##
class Employee:
    def set_name(self, new_name):
        self.name = new_name

    def set_salary(self, new_salary):
        self.salary = new_salary 

    # Add a give_raise() method with raise amount as a parameter
    def give_raise(self, raise_amount):
        self.salary = self.salary + raise_amount

emp = Employee()
emp.set_name('Korel Rossi')
emp.set_salary(50000)

print(emp.salary)
emp.give_raise(1500)
print(emp.salary)

##
class Employee:
    def set_name(self, new_name):
        self.name = new_name

    def set_salary(self, new_salary):
        self.salary = new_salary 

    def give_raise(self, amount):
        self.salary = self.salary + amount

    # Add monthly_salary method that returns 1/12th of salary attribute
    def monthly_salary(self):
        return self.salary / 12

    
emp = Employee()
emp.set_name('Korel Rossi')
emp.set_salary(50000)

# Get monthly salary of emp and assign to mon_sal
mon_sal = emp.monthly_salary()

# Print mon_sal
print(mon_sal)
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

## Write a class from scratch
```python

```




# 2. Inheritance and Polymorphism
## Instance and class data
```python

```

## Class-level attributes
```python

```

## Changing class attributes
```python

```

## Alternative constructors
```python

```

## Class inheritance
```python

```

## Understanding inheritance
```python

```

## Create a subclass
```python

```

## Customizing functionality via inheritance
```python

```

## Method inheritance
```python

```

## Inheritance of class attributes
```python

```

## Customizing a DataFrame
```python

```




# 3. Integrating with Standard Python
## Operator overloading: comparison
```python

```

## Overloading equality
```python

```

## Checking class equality
```python

```

## Comparison and inheritance
```python

```

## Operator overloading: string representation
```python

```

## String formatting review
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

## Handling exception hierarchies
```python

```




# 4. Best Practices of Class Design
## Designing for inheritance and polymorphism
```python

```

## Polymorphic methods
```python

```

## Square and rectangle
```python

```

## Managing data access: private attributes
```python

```

## Attribute naming conventions
```python

```

## Using internal attributes
```python

```

## Properties
```python

```

## What do properties do?
```python

```

## Create and set properties
```python

```

## Read-only properties
```python

```

## Congratulations!
```python

```

