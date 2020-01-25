---
title: Unit Testing for Data Science in Python
tags: unit-testing,software-testing,python
url: https://www.datacamp.com/courses/unit-testing-for-data-science-in-python
---

# 1. Unit testing basics
## Your first unit test using pytest
```python
# Import the pytest package
import pytest

# Import the function convert_to_int()
from preprocessing_helpers import convert_to_int

# Complete the unit test name by adding a prefix
def test_on_string_with_one_comma():
  # Complete the assert statement
  assert convert_to_int("2,081") == 2081
```

## Running unit tests
```python
!pytest test_convert_to_int.py
```

## Spotting and fixing bugs
```python
def convert_to_int(string_with_comma):
    # Fix this line so that it returns an int, not a str
    return int(string_with_comma.replace(",", ""))
```

```python
import pytest

from preprocessing_helpers import convert_to_int

def test_on_string_with_one_comma():
  assert convert_to_int("2,081") == 2081
```

## Benefits of unit testing
The benefits that unit testing might bring.

- Time savings, leading to faster development of new features.
- Improved documentation, which will help new colleagues understand the code base better.
- More user trust in the software product.
- Better user experience due to reduced downtime.

## Unit tests as documentation
```python
import numpy as np
import pytest

from mystery_function import mystery_function

def test_on_clean_data():
    assert np.array_equal(
		mystery_function("example_clean_data.txt", num_columns=2),
		np.array([[2081.0, 314942.0], [1059.0, 186606.0]]))
```

# 2. Intermediate unit testing
## Mastering assert statements
```python

```

## Write an informative test failure message
```python

```

## Testing float return values
```python

```

## Testing with multiple assert statements
```python

```

## Testing for exceptions instead of return values
```python

```

## Practice the context manager
```python

```

## Unit test a ValueError
```python

```

## The well tested function
```python

```

## Testing well: Boundary values
```python

```

## Testing well: Values triggering special logic
```python

```

## Testing well: Normal arguments
```python

```

## Test Driven Development (TDD)
```python

```

## TDD: Tests for normal arguments
```python

```

## TDD: Requirement collection
```python

```

## TDD: Implement the function
```python

```

# 3. Test Organization and Execution
## How to organize a growing set of tests?
```python

```

## Place test modules at the correct location
```python

```

## Create a test class
```python

```

## Mastering test execution
```python

```

## One command to run them all
```python

```

## Running test classes
```python

```

## Expected failures and conditional skipping
```python

```

## Mark a test class as expected to fail
```python

```

## Mark a test as conditionally skipped
```python

```

## Reasoning in the test result report
```python

```

## Continuous integration and code coverage
```python

```

## Build failing
```python

```

## What does code coverage mean?
```python

```

# 4. Testing Models, Plots and Much More
## Beyond assertion: setup and teardown
```python

```

## Use a fixture for a clean data file
```python

```

## Write a fixture for an empty data file
```python

```

## Fixture chaining using tmpdir
```python

```

## Mocking
```python

```

## Program a bug-free dependency
```python

```

## Mock a dependency
```python

```

## Testing models
```python

```

## Testing on linear data
```python

```

## Testing on circular data
```python

```

## Testing plots
```python

```

## Generate the baseline image
```python

```

## Run the tests for the plotting function
```python

```

## Fix the plotting function
```python

```

## Congratulations
```python

```
