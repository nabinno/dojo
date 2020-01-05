---
title: Software Engineering for Data Scientists in Python
tags: python
url: https://www.datacamp.com/courses/software-engineering-for-data-scientists-in-python
---

# 1. Software Engineering & Data Science
## Python modularity in the wild
```python
# import the numpy package
import numpy as np

# create an array class object
arr = np.array([8, 6, 7, 5, 3, 0, 9])

# use the sort method
arr.sort()

# print the sorted array
print(arr)
```

## Leveraging documentation
```python
# load the Counter function into our environment
from collections import Counter

# View the documentation for Counter.most_common
help(Counter.most_common)

# use Counter to find the top 5 most common words
top_5_words = Counter(words).most_common(5)

# display the top 5 most common words
print(top_5_words)
```

## Using pycodestyle
```python
# Import needed package
import pycodestyle

# Create a StyleGuide instance
style_checker = pycodestyle.StyleGuide()

# Run PEP 8 check on multiple files
result = style_checker.check_files(['nay_pep8.py', 'yay_pep8.py'])

# Print result of PEP 8 style check
print(result.messages)
```

## Conforming to PEP 8
```python
# Assign data to x
x = [8, 3, 4]

# Print the data
print(x)
```

## PEP 8 in documentation
```python
def print_phrase(phrase, polite=True, shout=False):
    if polite:  # It's generally polite to say please
        phrase = 'Please ' + phrase

    if shout:  # All caps looks like a written shout
        phrase = phrase.upper() + '!!'

    print(phrase)


# Politely ask for help
print_phrase('help me', polite=True)
# Shout about a discovery
print_phrase('eureka', shout=True)
```

# 2. Writing a Python Module
## Writing your first package
```python

```

## Minimal package requirements
```python

```

## Naming packages
```python

```

## Recognizing packages
```python

```

## Adding functionality to packages
```python

```

## Adding functionality to your package
```python

```

## Using your package's new functionality
```python

```

## Making your package portable
```python

```

## Writing requirements.txt
```python

```

## Installing package requirements
```python

```

## Creating setup.py
```python

```

## Listing requirements in setup.py
```python

```

# 3. Utilizing Classes
## Adding classes to a package
```python

```

## Writing a class for your package
```python

```

## Using your package's class
```python

```

## Adding functionality to classes
```python

```

## Writing a non-public method
```python

```

## Using your class's functionality
```python

```

## Classes and the DRY principle
```python

```

## Using inheritance to create a class
```python

```

## Adding functionality to a child class
```python

```

## Using your child class
```python

```

## Multilevel inheritance
```python

```

## Exploring with dir and help
```python

```

## Creating a grandchild class
```python

```

## Using inherited methods
```python

```

# 4. Maintainability
## Documentation
```python

```

## Identifying good comments
```python

```

## Identifying proper docstrings
```python

```

## Writing docstrings
```python

```

## Readability counts
```python

```

## Using good function names
```python

```

## Using good variable names
```python

```

## Refactoring for readability
```python

```

## Unit testing
```python

```

## Using doctest
```python

```

## Using pytest
```python

```

## Documentation & testing in practice
```python

```

## Documenting classes for Sphinx
```python

```

## Identifying tools
```python

```

## Final Thoughts
```python

```
