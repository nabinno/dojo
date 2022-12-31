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
# Strip duration of minutes
ride_sharing['duration_trim'] = ride_sharing['duration'].str.strip('minutes')

# Convert duration to integer
ride_sharing['duration_time'] = ride_sharing['duration_trim'].astype('int')

# Write an assert statement making sure of conversion
assert ride_sharing['duration_time'].dtype == 'int'

# Print formed columns and calculate average ride duration 
print(ride_sharing[['duration','duration_trim','duration_time']])
print(ride_sharing['duration_time'].mean())
```

## Tire size constraints
```python
# Convert tire_sizes to integer
ride_sharing['tire_sizes'] = ride_sharing['tire_sizes'].astype('int')

# Set all values above 27 to 27
ride_sharing.loc[ride_sharing['tire_sizes'] > 27, 'tire_sizes'] = 27

# Reconvert tire_sizes back to categorical
ride_sharing['tire_sizes'] = ride_sharing['tire_sizes'].astype('category')

# Print tire size description
print(ride_sharing['tire_sizes'].describe())
```

## Back to the future
```python
# Convert ride_date to date
ride_sharing['ride_dt'] = pd.to_datetime(ride_sharing['ride_date']).dt.date

# Save today's date
today = dt.date.today()

# Set all in the future to today's date
ride_sharing.loc[ride_sharing['ride_dt'] > today, 'ride_dt'] = today

# Print maximum of ride_dt column
print(ride_sharing['ride_dt'].max())
```

## Finding duplicates
```python
# Find duplicates
duplicates = ride_sharing.duplicated(subset = ['ride_id'], keep = False)

# Sort your duplicated rides
duplicated_rides = ride_sharing[duplicates].sort_values(['ride_id'])

# Print relevant columns of duplicated_rides
print(duplicated_rides[['ride_id','duration','user_birth_year']])
```

## Treating duplicates
```python
# Drop complete duplicates from ride_sharing
ride_dup = ride_sharing.drop_duplicates()

# Create statistics dictionary for aggregation function
statistics = {'user_birth_year': 'min', 'duration': 'mean'}

# Group by ride_id and compute new statistics
ride_unique = ride_dup.groupby('ride_id').agg(statistics).reset_index()

# Find duplicated values again
duplicates = ride_unique.duplicated(subset = 'ride_id', keep = False)
duplicated_rides = ride_unique[duplicates == True]

# Assert duplicates are processed
assert duplicated_rides.shape[0] == 0
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

