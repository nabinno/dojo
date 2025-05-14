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
```python
day_of_week = datetime.now().isoweekday()

def get_unique_values(lst):
    return list(set(lst))

condition_string = 'day_of_week == 6'
# Add the conditional skip marker and the string here
@pytest.mark.skipif(condition_string)
def test_function():
    # Complete the assertion tests here
    assert get_unique_values([1,2,3]) == [1,2,3]
    assert get_unique_values([1,2,3,1]) == [1,2,3]
```





# 2 Pytest Fixtures
## Getting familiar with fixtures
```
True:
- Fixtures help to make the test setup more modular.
- One has to declare the `@pytest.fixture` decorator along with a custom fixture function to implement a `pytest` fixture.
- A fixture is a prepared environment that is used in tests.

False:
- A fixture is a testing type.
- A `pytest` fixture can be implemented as a regular function without decorators.
```

## Data preparation
```python
# Import the pytest library
import pytest

# Define the fixture decorator
@pytest.fixture
# Name the fixture function
def prepare_data():
    return [i for i in range(10)]

# Create the tests
def test_elements(prepare_data):
    assert 9 in prepare_data
    assert 10 not in prepare_data
```

## Run with a fixture
```sh
$ cd /home/repl/workspace
$ echo run_with_fixtures.py
import pytest

@pytest.fixture
def prepare_data():
    return [i for i in range(10)]

def test_elements(prepare_data):
    assert 9 in prepare_data
    assert 10 not in prepare_data
    $ pytest run_with_fixtures.py
```

## Chain this out
```python
import pytest

@pytest.fixture
def setup_data():
    return [1, 2, 3, 4, 5]

@pytest.fixture
def process_data(setup_data):
    return [el * 2 for el in setup_data]

def test_process_data(process_data):
    assert process_data == [2, 4, 6, 8, 10]
```

## List with a custom length
```
$ echo list_custom_len.py
import pytest

# Define the fixture for returning the length
@pytest.fixture
def list_length():
    return 10

# Define the fixture for a list preparation
@pytest.fixture
def prepare_list(list_length):
    return [i for i in range(list_length)]

def test_9(prepare_list):
    assert 9 in prepare_list
    assert 10 not in prepare_list

$ pytest list_custom_len.py
```

## autouse statements
```
Select true statements about the autouse argument.

[ ]`autouse` is an argument of a test function.
[x]`autouse` helps to simplify the code.
[x]`autouse` is an argument of a fixture decorator.
[x]It might be convinient to apply `autouse` to use the same data for all tests.
[ ]`autouse` can make the code faster.
[x]It might be convenient to apply `autouse` to set the same parameters for all tests.
[ ]One can use variables declared inside of the `autouse` fixture and use it outside of it by default.
```

## Auto add numbers
```python
import pytest

@pytest.fixture
def init_list():
    return []

# Add autouse here
@pytest.fixture(autouse=True)
def add_numbers_to_list(init_list):
    init_list.extend([i for i in range(10)])

# Complete the tests
def test_elements(init_list):
    assert 1 in init_list
    assert 9 in init_list
```

## Data with teardown
```
$ echo data_w_teardown.py
import pytest

@pytest.fixture
def prepare_data():
    data = [i for i in range(10)]
    # Return the data with the special keyword
    yield data
    # Clear the data list
    data.clear()
    # Delete the data variable
    del data

def test_elements(prepare_data):
    assert 9 in prepare_data
    assert 10 not in prepare_data

$ pytest data_w_teardown.py
```

## Read data with teardown
```
$ echo read_df_teardown.py
import pytest
import pandas as pd

@pytest.fixture
def data():
    df = pd.read_csv('/usr/local/share/games.csv')
    # Return df with the special keyword
    yield df
    # Remove all rows in df
    df.drop(df.index, inplace=True)
    # Delete the df variable
    del df

def test_type(data):
    assert type(data) == pd.DataFrame

def test_shape(data):
    assert data.shape[0] == 1512

$ pytest read_df_teardown.py
```





# 3 Basic Testing Types
## Unit testing terms
```
True:
- Creating test case is creative task.
- A test case is a set of unit inputs and expected outputs summarizing a partifcular problem.
- Unit test is a type of test that verifies a software unit.

False:
- Unit test is a type of test that verifies software as a whole.
- A test case is an implemented test.
```

## Cover more test cases
```
What is the best set of tests (that covers more test cases) to validate the following division function?
[ ]division(47,-100), division(0, 16), division(740, 512), division(-345, -10), division(-10, 589)
[ ]division(10,0), division(0, 55), division(0, 0), division(740, 512), division(10, 10), division(100, 5)
[x]division(11,0), division(0, 60), division(0, 0), division(876, 34), division(-10, -10), division(-10, 5)
```

## Factorial of number
```python
def factorial(n):
    if n == 0: return 1
    elif (type(n) == int):
        return n * factorial(n-1)
    else: return -1

# Test case: expected input
def test_regular():
    assert factorial(5) == 120

# Test case: zero input
def test_zero():
    assert factorial(0) == 1

# Test case: input of a wrong type
def test_str():
    assert factorial('5') == -1
    print('Test passed')
```

## Run factorial
```python
import pytest

def factorial(n):
    if n == 0: return 1
    elif (type(n) == int):
        return n * factorial(n-1)
    else: return -1

# Test case: expected input
def test_regular():
    assert factorial(5) == 120

# Test case: zero input
def test_zero():
    assert factorial(0) == 1

# Test case: input of a wrong type
def test_str():
    assert factorial('5') == -1
```

## Feature or unit testing
```
Feature test:
- Verifying that a browser can open a website on a computer.
- Ensuring maps navigation app can build a route from A to B.
- Testing that a user can log in to your website.

Unit test:
- Checking a function of sorting an array in a relational database.
- Checking that Python properly decodes a symbol from ASCII.
```

## Aggregate with sum
```
$ echo agg_with_sum.py
# Don't forget to run
# pytest agg_with_sum.py
# in the CLI to actually run the test!

import pandas as pd
import pytest

# Fixture to prepare the data
@pytest.fixture
def get_df():
    return pd.read_csv('https://assets.datacamp.com/production/repositories/6253/datasets/757c6cb769f7effc5f5496050ea4d73e4586c2dd/laptops_train.csv')

# Aggregation feature
def agg_with_sum(data, group_by_column, aggregate_column):
    return data.groupby(group_by_column)[aggregate_column].sum()

# Test function
def test_agg_feature(get_df):
    # Aggregate preparation
    aggregated = agg_with_sum(get_df, 'Manufacturer', 'Price')
    # Test the type of the aggregated
    assert isinstance(aggregated, pd.Series)
    # Test the number of rows of the aggregated
    assert aggregated.shape[0] > 0
    # Test the data type of the aggregated
    assert aggregated.dtype in (int, float)

$ pytest agg_with_sum.py
```

## Integration test or not
```
Integration test:
- Checking that ping requests can reach the destination and return.
- Ensuring that the power cable is connected to the device.
- Verifying the connection between a cloud server and a database.

Not an integration test:
- Ensuring the tires of your car have sufficient pressure.
- Confirmaing that the "addition" button of a calculator works properly.
- Checking that the dataset concatenation works.
```

## Read the file
```
$ echo reading_test.py
import pandas as pd
import pytest

# Fixture to read the dataframe
@pytest.fixture
def get_df():
    return pd.read_csv('https://assets.datacamp.com/production/repositories/6253/datasets/757c6cb769f7effc5f5496050ea4d73e4586c2dd/laptops_train.csv')

# Integration test function
def test_get_df(get_df):
    # Check the type
    assert isinstance(get_df, pd.DataFrame)
    # Check the number of rows
    assert get_df.shape[0] > 0

$ pytest reading_test.py
```

## What is performance testing?
```
[ ]To scrutinize the correctness of a unit
[ ]To verify the interactions between different modules of a system
[x]To measure software performance
```

## Finding an element
```python
def create_list():
    return [i for i in range(1000)]
def create_set():
    return set([i for i in range(1000)])
def find(it, el=50):
    return el in it

# Write the performance test for a list
def test_list(benchmark):
    benchmark(find, it=create_list())

# Write the performance test for a set
def test_set(benchmark):
    benchmark(find, it=create_set())
```

## Speed of loops
```python
def test_list(benchmark):
    # Add decorator here
    @benchmark
    def iterate_list():
        # Complete the loop here
        for iterate_ in [i for i in range(1000)]:
            pass

def test_set(benchmark):
    # Add decorator here
    @benchmark
    def iterate_set():
        # Complete the loop here
        for iterate_ in {i for i in range(1000)}:
            pass
```




# 4 Writing tests with unittest
## Factorial with unittest
```python
def func_factorial(number):
    if number < 0:
        raise ValueError('Factorial is not defined for negative values')
    factorial = 1
    while number > 1:
        factorial = factorial * number
        number = number - 1
        return factorial

class TestFactorial(unittest.TestCase):
    def test_positives(self):
        # Add the test for testing positives here
        self.assertEqual(func_factorial(5), 120)

    def test_zero(self):
        # Add the test for testing zero here
        self.assertEqual(func_factorial(0), 1)

    def test_negatives(self):
        # Add the test for testing negatives here
        with self.assertRaises(ValueError):
            func_factorial(-5)
```

## Is prime or not
```python
def is_prime(num):
    if num == 1: return False
    up_limit = int(math.sqrt(num)) + 1
    for i in range(2, up_limit):
        if num % i == 0:
            return False
    return True

class TestSuite(unittest.TestCase):
    def test_is_prime(self):
        # Check that 17 is prime
        self.assertTrue(is_prime(17))

    def test_is_prime(self):
        # Check that 6 is not prime
        self.assertFalse(is_prime(6))

    def test_is_prime(self):
        # Check that 1 is not prime
        self.assertFalse(is_prime(1))
```

## Run factorial with unittest
```sh
$ python3 -m unittest factorial_unittest.py
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
