---
title: Writing Functions in Python
tags: python
url: https://www.datacamp.com/courses/writing-functions-in-python
---

# 1. Best Practices
## Crafting a docstring
```python
def count_letter(content, letter):
  """Count the number of times `letter` appears in `content`.

  Args:
    content (str): The string to search.
    letter (str): The letter to search for.

  Returns:
    int

  # Add a section detailing what errors might be raised
  Raises:
    ValueError: If `letter` is not a one-character string.
  """
  if (not isinstance(letter, str)) or len(letter) != 1:
    raise ValueError('`letter` must be a single character string.')
  return len([char for char in content if char == letter])
```

## Retrieving docstrings
```python
##
# Get the docstring with an attribute of count_letter()
docstring = count_letter.__doc__

border = '#' * 28
print('{}\n{}\n{}'.format(border, docstring, border))

##
import inspect

# Get the docstring with a function from the inspect module
docstring = inspect.getdoc(count_letter)

border = '#' * 28
print('{}\n{}\n{}'.format(border, docstring, border))

##
def build_tooltip(function):
  """Create a tooltip for any function that shows the 
  function's docstring.
  
  Args:
    function (callable): The function we want a tooltip for.
    
  Returns:
    str
  """
  # Use 'inspect' to get the docstring
  docstring = inspect.getdoc(function)
  border = '#' * 28
  return '{}\n{}\n{}'.format(border, docstring, border)

print(build_tooltip(count_letter))
print(build_tooltip(range))
print(build_tooltip(print))
```

## Docstrings to the rescue!
```python

```

## DRY and "Do One Thing"
```python

```

## Extract a function
```python

```

## Split up a function
```python

```

## Pass by assignment
```python

```

## Mutable or immutable?
```python

```

## Best practice for default arguments
```python

```

# 2. Context Managers
## Using context managers
```python

```

## The number of cats
```python

```

## The speed of cats
```python

```

## Writing context managers
```python

```

## The timer() context manager
```python

```

## A read-only open() context manager
```python

```

## Advanced topics
```python

```

## Context manager use cases
```python

```

## Scraping the NASDAQ
```python

```

## Changing the working directory
```python

```

# 3. Decorators
## Functions are objects
```python

```

## Building a command line data app
```python

```

## Reviewing your co-worker's code
```python

```

## Returning functions for a math game
```python

```

## Scope
```python

```

## Understanding scope
```python

```

## Modifying variables outside local scope
```python

```

## Closures
```python

```

## Checking for closure
```python

```

## Closures keep your values safe
```python

```

## Decorators
```python

```

## Using decorator syntax
```python

```

## Defining a decorator
```python

```

# 4. More on Decorators
## Real-world examples
```python

```

## Print the return type
```python

```

## Counter
```python

```

## Decorators and metadata
```python

```

## Preserving docstrings when decorating functions
```python

```

## Measuring decorator overhead
```python

```

## Decorators that take arguments
```python

```

## Run_n_times()
```python

```

## HTML Generator
```python

```

## Timeout(): a real world example
```python

```

## Tag your functions
```python

```

## Check the return type
```python

```

## Great job!
```python

```


