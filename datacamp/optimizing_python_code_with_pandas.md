---
title: Optimizing Python Code with pandas
tags: python,pandas
url: https://www.datacamp.com/courses/optimizing-python-code-with-pandas
---

# 1. Selecting columns and rows efficiently
## Measuring time I
```python
# Calculate the result of the problem using formula() and print the time required
N = 1000000
fm_start_time = time.time()
first_method = formula(N)
print("Time using formula: {} sec".format(time.time() - fm_start_time))

# Calculate the result of the problem using brute_force() and print the time required
sm_start_time = time.time()
second_method = brute_force(N)
print("Time using the brute force: {} sec".format(time.time() - sm_start_time))
```

## Measuring time II
```python
# Store the time before the execution
start_time = time.time()

# Execute the operation
letlist = [wrd for wrd in words if wrd.startswith('b')]

# Store and print the difference between the start and the current time
total_time_lc = time.time() - start_time
print('Time using list comprehension: {} sec'.format(total_time_lc))

# Store the time before the execution
start_time = time.time()

# Execute the operation
letlist = []
for wrd in words:
    if wrd.startswith('b'):
        letlist.append(wrd)
        
# Print the difference between the start and the current time
total_time_fl = time.time() - start_time
print('Time using for loop: {} sec'.format(total_time_fl))
```

## Row selection: loc[] vs iloc[]
```python

```

## Column selection: .iloc[] vs by name
```python

```

## Select random rows
```python

```

## Random row selection
```python

```

## Random column selection
```python

```

# 2. Replacing values in a DataFrame
## Replace scalar values using .replace()
```python

```

## Replacing scalar values I
```python

```

## Replace scalar values II
```python

```

## Replace values using lists
```python

```

## Replace multiple values I
```python

```

## Replace multiple values II
```python

```

## Replace values using dictionaries
```python

```

## Replace single values I
```python

```

## Replace single values II
```python

```

## Replace multiple values III
```python

```

## Most efficient method for scalar replacement
```python

```

# 3. Efficient iterating
## Looping using the .iterrows() function
```python

```

## Create a generator for a pandas DataFrame
```python

```

## The iterrows() function for looping
```python

```

## Looping using the .apply() function
```python

```

## .apply() function in every cell
```python

```

## .apply() for rows iteration
```python

```

## Vectorization over pandas series
```python

```

## Why vectorization in pandas is so fast?
```python

```

## pandas vectorization in action
```python

```

## Vectorization with NumPy arrays using .values()
```python

```

## Best method of vectorization
```python

```

## Vectorization methods for looping a DataFrame
```python

```

# 4. Data manipulation using .groupby()
## Data transformation using .groupby().transform
```python

```

## The min-max normalization using .transform()
```python

```

## Transforming values to probabilities
```python

```

## Validation of normalization
```python

```

## When to use transform()?
```python

```

## Missing value imputation using transform()
```python

```

## Identifying missing values
```python

```

## Missing value imputation
```python

```

## Data filtration using the filter() function
```python

```

## When to use filtration?
```python

```

## Data filtration
```python

```

## Congratulations!
```python

```
