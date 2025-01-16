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

## Adding an alternative constructor
```python
class Person:
  CURRENT_YEAR = 2024
  def __init__(self, name, age):
    self.name = name
    self.age = age

  # Add a class method decorator
  @classmethod
  # Define the from_birth_year method
  def from_birth_year(cls, name, birth_year):
    # Create age
    age = cls.CURRENT_YEAR - birth_year
    # Return the name and age
    return cls(name, age)

bob = Person.from_birth_year("Bob", 1990)
```

## Building a BetterDate Class
```python
class BetterDate:
  def __init__(self, year, month, day):
    self.year, self.month, self.day = year, month, day

  # Define a class method from_str
  @classmethod
  def from_str(cls, datestr):
    # Split the string at "-"
    parts = datestr.split("-")
    year, month, day = int(parts[0]), int(parts[1]), int(parts[2])
    # Return the class instance
    return cls(year, month, day)

# Create the xmas object
xmas = BetterDate.from_str("2024-12-25")
print(xmas.year)
print(xmas.month)
print(xmas.day)
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

# Define a new class Manager inheriting from Employee
class Manager(Employee):
  # Add a keyword to leave this class empty
  pass

# Define a Manager object
mng = Manager("Debbie Lashko", 86500)

# Print mng's name
print(mng.name)
```

## Understanding inheritance
```
True:
- If `ind` is an `Indexxer` object, then `isinstance(ind, Counter)` will return `True`.
- Running `ind = Indexer()` will cause an error.

False:
- If `ind` is an `Indexer` object,  then `ind.and_counts(5)` will cause an error.
- Inheritance can be used to add some of the parts of one class to another class.
- Class `Indexer` inherits from `Counter`.
- Every `Counter` object is an `Indexxer` object.
```

## Customize a subclass
```python
class Employee:
  def __init__(self, name, salary=30000):
    self.name = name
    self.salary = salary

  def give_raise(self, amount):
    self.salary += amount

class Manager(Employee):
  # Add a constructor
  def __init__(self, name, salary=50000, project=None):

    # Call the parent's constructor
    Employee.__init__(self, name, salary)

    # Assign project attribute
    self.project = project

  def display(self):
    print("Manager ", self.name)
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
    amount = amount * bonus
    Employee.give_raise(self, amount)

    mngr = Manager("Ashta Dunbar", 78500)
    mngr.give_raise(2000, bonus=1.03)
    print(mngr.salary)
```

## Inheritance of class attributes
```python
class Player:
  MAX_POSITION = 10

  def __init__(self):
    self.position = 0

  def move(self, steps):
    if self.position + steps < Player.MAX_POSITION:
      self.position += steps
    else:
      self.position = Player.MAX_POSITION

# Create a Racer class inheriting from Player
class Racer(Player):
  # Create MAX_POSITION with a value of 15
  MAX_POSITION = 15

# Create a Player and a Racer objects
p = Player()
r = Racer()

print("p.MAX_POSITION = ", p.MAX_POSITION)
print("r.MAX_POSITION = ", r.MAX_POSITION)
# => <script.py> output:
#     p.MAX_POSITION =  10
#     r.MAX_POSITION =  15

# [ ] Class attributes CANNOT be inherited, but new class attributes of the same name CAN be created in a child class.
# [ ] Class attributes CANNOT be inherited, and new class attributes of the same name CANNOT be created in a child class.
# [x] Class attributes CAN be inherited, and the value of class attributes CAN be overwritten in the child class
# [ ] Class attributes CAN be inherited,  and the value of class attributes CANNOT be overwritten in the child class
```




# 3. Integrating with Standard Python
## Overloading equality
```python
class BankAccount:
  # Modify to initialize a number attribute
  def __init__(self, number, balance=0):
    self.balance = balance
    self.number = number

  def withdraw(self, amount):
    self.balance -= amount

  # Define __eq__ that returns True if the number attributes are equal
  def __eq__(self, other):
    return self.number == other.number

# Create accounts and compare them
acct1 = BankAccount(123, 1000)
acct2 = BankAccount(123, 1000)
acct3 = BankAccount(456, 1000)
print(acct1 == acct2)
print(acct1 == acct3)
```

## Checking class equality
```python
class Phone:
  def __init__(self, number):
    self.number = number

  def __eq__(self, other):
    return (self.number == other.number) & (type(self) == type(other))

class BankAccount:
  def __init__(self, number, balance=0):
    self.number, self.balance = number, balance

  def withdraw(self, amount):
    self.balance -= amount

  # Modify to add a check for the class type
  def __eq__(self, other):
    return (self.number == other.number) & (type(self) == type(other))

acct = BankAccount(873555333)
pn = Phone(873555333)

# Check if the two objects are equal
print(acct == pn)
```

## Object representation
```python
__str__():
- Informal, mostly used by end-users
- Returns an object in string format

__repr__():
- Shows the object in a format representing how it was made
- Formal, mostly for use by developers
```

## Comparison and inheritance
```python
class Parent:
  def __eq__(self, other):
    print("Parent's __eq__() called")
    return True

class Child(Parent):
  def __eq__(self, other):
    print("Child's __eq__() called")
    return True

p = Parent()
c = Child()

p == c
# Child's __eq__() called
# True
```

## String representation of objects
```python
class Employee:
  def __init__(self, name, salary=30000):
    self.name, self.salary = name, salary

  # Add the __repr__() method
  def __repr__(self):
    emp_str = f"Employee('{self.name}', {self.salary})"
    return emp_str

emp1 = Employee("Amar Howard", 30000)
print(repr(emp1))
emp2 = Employee("Carolyn Ramirez", 35000)
print(repr(emp2))
# <script.py> output:
#   Employee('Amar Howard', 30000)
#   Employee('Carolyn Ramirez', 35000)

class Employee:
  def __init__(self, name, salary=30000):
    self.name, self.salary = name, salary

  # Add the __str__() method
  def __str__(self):
    emp_str = f"""Employee name: {self.name}
    Employee salary: {self.salary}"""
    return emp_str

emp1 = Employee("Amar Howard", 30000)
print(emp1)
emp2 = Employee("Carolyn Ramirez", 35000)
print(emp2)
# <script.py> output:
#   Employee name: Amar Howard
#   Employee salary: 30000
#   Employee name: Carolyn Ramirez
#   Employee salary: 35000")""""
```

## Exceptions
```python
# Modify the function to catch exceptions
def invert_at_index(x, ind):
  try:
    return 1/x[ind]
  except ZeroDivisionError:
    print("Cannot divide by zero!")
  except IndexError:
    print("Index out of range!")

a_list = [5,6,0,7]

# Works okay
print(invert_at_index(a_list, 1))

# Potential ZeroDivisionError
print(invert_at_index(a_list, 2))

# Potential IndexError
print(invert_at_index(a_list, 5))
```

## Catching exceptions
```python
class SalaryError(ValueError):
    pass
class BonusError(SalaryError):
  pass

class Employee:
  MIN_SALARY = 30000
  MAX_BONUS = 5000

  def __init__(self, name, salary = 30000):
    self.name = name
    if salary < Employee.MIN_SALARY:
      raise SalaryError("Salary is too low!")
    self.salary = salary

    # Raise exceptions
    def give_bonus(self, amount):
      if amount > Employee.MAX_BONUS:
        raise BonusError("The bonus amount is too high!")

      elif self.salary + amount <  Employee.MIN_SALARY:
        raise SalaryError("The salary after bonus is too low!")

      self.salary += amount
```

## Custom exceptions
```python

```
