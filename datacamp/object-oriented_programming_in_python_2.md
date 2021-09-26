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

## Correct use of __init__
```python
class Counter:
    def __init__(self, count, name):
	  self.count = 5
	  self.name = name

c = Counter(0, "My counter")
print(c.count)
```

## Add a class constructor
```python
# Import datetime from datetime
from datetime import datetime

class Employee:   
    def __init__(self, name, salary=0):
        self.name = name
        if salary > 0:
          self.salary = salary
        else:
          self.salary = 0
          print("Invalid salary!")
          
        # Add the hire_date attribute and set it to today's date
        self.hire_date = datetime.today()
        
    # From the previous lesson
    def give_raise(self, amount):
        self.salary += amount

    def monthly_salary(self):
        return self.salary/12
      
emp = Employee("Korel Rossi", -1000)
print(emp.name)
print(emp.salary)
```

## Write a class from scratch
```python
# For use of np.sqrt
import numpy as np

class Point:
    """ A point on a 2D plane
    
   Attributes
    ----------
    x : float, default 0.0. The x coordinate of the point        
    y : float, default 0.0. The y coordinate of the point
    """
    def __init__(self, x=0.0, y=0.0):
      self.x = x
      self.y = y
      
    def distance_to_origin(self):
      """Calculate distance from the point to the origin (0,0)"""
      return np.sqrt(self.x ** 2 + self.y ** 2)
    
    def reflect(self, axis):
      """Reflect the point with respect to x or y axis."""
      if axis == "x":
        self.y = - self.y
      elif axis == "y":
        self.x = - self.x
      else:
        print("The argument axis only accepts values 'x' and 'y'!")
```



# 2. Inheritance and Polymorphism
## Class-level attributes
```python
class Player:
    MAX_POSITION = 10
    
    def __init__(self):
        self.position = 0

    # Add a move() method with steps parameter     
    def move(self, steps):
        if self.position + steps < Player.MAX_POSITION:
           self.position = self.position + steps 
        else:
           self.position = Player.MAX_POSITION
    
    # This method provides a rudimentary visualization in the console    
    def draw(self):
        drawing = "-" * self.position + "|" +"-"*(Player.MAX_POSITION - self.position)
        print(drawing)

p = Player(); p.draw()
p.move(4); p.draw()
p.move(5); p.draw()
p.move(3); p.draw()
```

## Changing class attributes
```python
##
# Create Players p1 and p2
p1 = Player(); p2 = Player()

print("MAX_SPEED of p1 and p2 before assignment:")
# Print p1.MAX_SPEED and p2.MAX_SPEED
print(p1.MAX_SPEED)
print(p2.MAX_SPEED)

# Assign 7 to p1.MAX_SPEED
p1.MAX_SPEED = 7

print("MAX_SPEED of p1 and p2 after assignment:")
# Print p1.MAX_SPEED and p2.MAX_SPEED
print(p1.MAX_SPEED)
print(p2.MAX_SPEED)

print("MAX_SPEED of Player:")
# Print Player.MAX_SPEED
print(Player.MAX_SPEED)

##
# Create Players p1 and p2
p1, p2 = Player(), Player()

print("MAX_SPEED of p1 and p2 before assignment:")
# Print p1.MAX_SPEED and p2.MAX_SPEED
print(p1.MAX_SPEED)
print(p2.MAX_SPEED)

# ---MODIFY THIS LINE--- 
Player.MAX_SPEED = 7

print("MAX_SPEED of p1 and p2 after assignment:")
# Print p1.MAX_SPEED and p2.MAX_SPEED
print(p1.MAX_SPEED)
print(p2.MAX_SPEED)

print("MAX_SPEED of Player:")
# Print Player.MAX_SPEED
print(Player.MAX_SPEED)
```

## Alternative constructors
```python
# import datetime from datetime
from datetime import datetime

class BetterDate:
    def __init__(self, year, month, day):
      self.year, self.month, self.day = year, month, day
      
    @classmethod
    def from_str(cls, datestr):
        year, month, day = map(int, datestr.split("-"))
        return cls(year, month, day)
      
    # Define a class method from_datetime accepting a datetime object
    @classmethod
    def from_datetime(cls, datetime):
        year, month, day = datetime.year, datetime.month, datetime.day
        return cls(year, month, day)

# You should be able to run the code below with no errors: 
today = datetime.today()     
bd = BetterDate.from_datetime(today)   
print(bd.year)
print(bd.month)
print(bd.day)
```

## Understanding inheritance
```python
## True
Class `Indexer` is inherited from `Counter`.
Inheritance represents is-a relationship.
Running `ind = Indexer()` will cause an error.
If `ind` is an `Indexer` object, then `isinstance(ind, Counter)` will return `True`.

## False
Inheritance can be used to add some of the parts of one class to another class.
If `ind` is an `Indexer` object, then running `ind.add_counts(5)` will cause an error.
Every `Counter` object is an `Indexer` object.
```

## Create a subclass
```python
class Employee:
  MIN_SALARY = 30000

  def __init__(self, name, salary=MIN_SALARY):
      self.name = name
      if salary >= Employee.MIN_SALARY:
        self.salary = salary
      else:
        self.salary = Employee.MIN_SALARY
  def give_raise(self, amount):
    self.salary += amount
        
# MODIFY Manager class and add a display method
class Manager(Employee):
  def display(self):
    print("Manager", self.name)

mng = Manager("Debbie Lashko", 86500)
print(mng.name)

# Call mng.display()
mng.display()
```

## Method inheritance
```python
class Employee:
    def __init__(self, name, salary=30000):
        self.name = name
        self.salary = salary

    def give_raise(self, amount):
        self.salary += amount
        
class Manager(Employee):
    def display(self):
        print("Manager ", self.name)

    def __init__(self, name, salary=50000, project=None):
        Employee.__init__(self, name, salary)
        self.project = project

    # Add a give_raise method
    def give_raise(self, amount, bonus=1.05):
        Employee.give_raise(self, amount*bonus)

mngr = Manager("Ashta Dunbar", 78500)
mngr.give_raise(1000)
print(mngr.salary)
mngr.give_raise(2000, bonus=1.03)
print(mngr.salary)
```

## Inheritance of class attributes
```python
# Create a Racer class and set MAX_SPEED to 5
class Racer(Player):
    MAX_SPEED = 5
 
# Create a Player and a Racer objects
p = Player()
r = Racer()

print("p.MAX_SPEED = ", p.MAX_SPEED)
print("r.MAX_SPEED = ", r.MAX_SPEED)

print("p.MAX_POSITION = ", p.MAX_POSITION)
print("r.MAX_POSITION = ", r.MAX_POSITION)
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

