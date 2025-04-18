---
title: Introduction to Testing in Python
tags: python,software-testing
url: https://campus.datacamp.com/courses/introduction-to-testing-in-python
---

# 1 1Creating Tests with pytest
## The first test suite
```python
def multiple_of_two(num):
    if num == 0:
    raise(ValueError)
    return num % 2 == 0

def test_numbers():
    assert multiple_of_two(2) is True
    # Write the "False" test below
    assert multiple_of_two(3) is False
```

## pytest.raises
```
def multiple_of_two(num):
    if num == 0:
        raise(ValueError)
    return num % 2 == 0

def test_zero():
    # Add a context for an exception test here
    with pytest.raises(ValueError):
        # Check zero input below
        multiple_of_two(0)
```

## Run the test!
```
$ echo run_the_test.py
# Import the pytest library
import pytest

def multiple_of_two(num):
    if num == 0:
        raise(ValueError)
    return num % 2 == 0

def test_numbers():
    assert multiple_of_two(2) is True
    assert multiple_of_two(3) is False

def test_zero():
    with pytest.raises(ValueError):
        multiple_of_two(0)

$ pytest run_the_test.py
```

## Run with the keyword
```
$ echo run_the_test.py
# Import the pytest library
import pytest

def multiple_of_two(num):
    if num == 0:
        raise(ValueError)
    return num % 2 == 0

def test_numbers():
    assert multiple_of_two(2) == True
    assert multiple_of_two(3) == False

def test_zero():
    with pytest.raises(ValueError):
        multiple_of_two(0)

$ pytest run_the_test.py "numbers"
```

## Markers use cases
```
@pytest.mark.skip:
- Skip the test no matter what.
- A test should be skipped indenfinitely until the mark is removed.

@pytest.mark.xfail:
- To verify that the test fail.
- To verify that the test checking 'a' + 'b' != 'ab' will fail.

@pytest.mark.skipif:
- To skip the test if Python version is less than `3.4.x`.
- To check the condition and skip the test if it is `True`.
```

## Failed tests with xfail
```python
def multiple_of_two(num):
    if num == 0:
        raise(ValueError)
    return num % 2 == 0

# Add the pytest marker decorator here
@pytest.mark.xfail
def test_fails():
    # Write any assert test that will fail
    assert multiple_of_two(0) is False
```

## Conditional skipping
```

```





# 2 Pytest Fixtures
## Introduction to fixtures
```

```

## Getting familiar with fixtures
```

```

## Data preparation
```

```

## Run with a fixture
```

```

## Chain Fixtures Requests
```

```

## Chain this out
```

```

## List with a custom length
```

```

## Fixtures autouse
```

```

## autouse statements
```

```

## Auto add numbers
```

```

## Fixtures Teardowns
```

```

## Data with teardown
```

```

## Read data with teardown
```

```





# 3 Basic Testing Types
## Unit testing with pytest
```

```

## Unit testing terms
```

```

## Cover more test cases
```

```

## Factorial of number
```

```

## Run factorial
```

```

## Feature testing with pytest
```

```

## Feature or unit testing
```

```

## Aggregate with sum
```

```

## Integration testing with pytest
```

```

## Integration test or not
```

```

## Read the file
```

```

## Performance testing with pytest
```

```

## What is performance testing?
```

```

## Finding an element
```

```

## Speed of loops
```

```




# 4 Writing tests with unittest
## Meeting the Unittest
```

```

## Factorial with unittest
```

```

## Is prime or not
```

```

## CLI Interface
```

```

## Run factorial with unittest
```

```

## Erroneouos factorial
```

```

## Unittest options
```

```

## Fixtures in unittest
```

```

## Test the string variable
```

```

## Palindrome check
```

```

## Practical examples
```

```

## Integration and unit tests
```

```

## Feature and performance tests
```

```

## Energy pipeline
```

```

## Congratulations!
```

```
