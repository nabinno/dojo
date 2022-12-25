---
title: Cleaning Data in Python (2)
tags: python,data-pre-processing
url: https://www.datacamp.com/courses/cleaning-data-in-python
---

# 1. Common data problems
## Common data types
```python
## Numeric data types
Salary earned monthly
Number of items bought in a basket

## Text
Shipping address of a customer
First name
City of residence

## Dates
Order date of a product
Birthdates of clients
```

## Numeric data or ... ?
```python
# Print the information of ride_sharing
print(ride_sharing.info())

# Print summary statistics of user_type column
print(ride_sharing['user_type'].describe())

# Convert user_type from integer to category
ride_sharing['user_type_cat'] = ride_sharing['user_type'].astype('category')

# Write an assert statement confirming the change
assert ride_sharing['user_type_cat'].dtype == 'category'

# Print new summary statistics 
print(ride_sharing['user_type_cat'].describe())
```

## Summing strings and concatenating numbers
```python

```

## Data range constraints
```python

```

## Tire size constraints
```python

```

## Back to the future
```python

```

## Uniqueness constraints
```python

```

## How big is your subset?
```python

```

## Finding duplicates
```python

```

## Treating duplicates
```python

```



# 2. Text and categorical data problems
## Membership constraints
```python

```

## Members only
```python

```

## Finding consistency
```python

```

## Categorical variables
```python

```

## Categories of errors
```python

```

## Inconsistent categories
```python

```

## Remapping categories
```python

```

## Cleaning text data
```python

```

## Removing titles and taking names
```python

```

## Keeping it descriptive
```python

```



# 3. Advanced data problems
## Uniformity
```python

```

## Ambiguous dates
```python

```

## Uniform currencies
```python

```

## Uniform dates
```python

```

## Cross field validation
```python

```

## Cross field or no cross field?
```python

```

## How's our data integrity?
```python

```

## Completeness
```python

```

## Is this missing at random?
```python

```

## Missing investors
```python

```

## Follow the money
```python

```



# 4. Record linkage
## Comparing strings
```python

```

## Minimum edit distance
```python

```

## The cutoff point
```python

```

## Remapping categories II
```python

```

## Generating pairs
```python

```

## To link or not to link?
```python

```

## Pairs of restaurants
```python

```

## Similar restaurants
```python

```

## Linking DataFrames
```python

```

## Getting the right index
```python

```

## Linking them together!
```python

```

## Congratulations!
```python

```

