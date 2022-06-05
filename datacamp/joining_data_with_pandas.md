---
tags: pandas, python
title: Joining Data with pandas
url: https://campus.datacamp.com/courses/joining-data-with-pandas/data-merging-basics
---

# Data Merging Basics
## What column to merge on?
```python
taxi_owners.merge(taxi_veh, on='vid')
```

## Your first inner join
```python
##
# Merge the taxi_owners and taxi_veh tables
taxi_own_veh = taxi_owners.merge(taxi_veh, on="vid")

# Print the column names of the taxi_own_veh
print(taxi_own_veh.columns)

##
# Merge the taxi_owners and taxi_veh tables setting a suffix
taxi_own_veh = taxi_owners.merge(taxi_veh, on='vid', suffixes=('_own', '_veh'))

# Print the column names of taxi_own_veh
print(taxi_own_veh.columns)

##
# Merge the taxi_owners and taxi_veh tables setting a suffix
taxi_own_veh = taxi_owners.merge(taxi_veh, on='vid', suffixes=('_own','_veh'))

# Print the value_counts to find the most popular fuel_type
print(taxi_own_veh['fuel_type'].value_counts())
```

## Inner joins and number of rows returned
```python
##
# Merge the wards and census tables on the ward column
wards_census = wards.merge(census, on="ward")

# Print the shape of wards_census
print('wards_census table shape:', wards_census.shape)

##
# Print the first few rows of the wards_altered table to view the change 
print(wards_altered[['ward']].head())

# Merge the wards_altered and census tables on the ward column
wards_altered_census = wards_altered.merge(census, on="ward")

# Print the shape of wards_altered_census
print('wards_altered_census table shape:', wards_altered_census.shape)

##
# Print the first few rows of the census_altered table to view the change 
print(census_altered[['ward']].head())

# Merge the wards and census_altered tables on the ward column
wards_census_altered = wards.merge(census_altered, on="ward")

# Print the shape of wards_census_altered
print('wards_census_altered table shape:', wards_census_altered.shape)
```

## One-to-many relationships
```python

```

## One-to-many classification
```python

```

## One-to-many merge
```python

```

## Merging multiple DataFrames
```python

```

## Total riders in a month
```python

```

## Three table merge
```python

```

## One-to-many merge with multiple tables
```python

```




# Merging Tables With Different Join Types
## Left join
```python

```

## Counting missing rows with left join
```python

```

## Enriching a dataset
```python

```

## How many rows with a left join?
```python

```

## Other joins
```python

```

## Right join to find unique movies
```python

```

## Popular genres with right join
```python

```

## Using outer join to select actors
```python

```

## Merging a table to itself
```python

```

## Self join
```python

```

## How does pandas handle self joins?
```python

```

## Merging on indexes
```python

```

## Index merge for movie ratings
```python

```

## Do sequels earn more?
```python

```




# Advanced Merging and Concatenating
## Filtering joins
```python

```

## Steps of a semi join
```python

```

## Performing an anti join
```python

```

## Performing a semi join
```python

```

## Concatenate DataFrames together vertically
```python

```

## Concatenation basics
```python

```

## Concatenating with keys
```python

```

## Using the append method
```python

```

## Verifying integrity
```python

```

## Validating a merge
```python

```

## Concatenate and merge to find common songs
```python

```




# Merging Ordered and Time-Series Data
## Using merge_ordered()
```python

```

## Correlation between GDP and S&P500
```python

```

## Phillips curve using merge_ordered()
```python

```

## merge_ordered() caution, multiple columns
```python

```

## Using merge_asof()
```python

```

## Using merge_asof() to study stocks
```python

```

## Using merge_asof() to create dataset
```python

```

## merge_asof() and merge_ordered() differences
```python

```

## Selecting data with .query()
```python

```

## Explore financials with .query()
```python

```

## Subsetting rows with .query()
```python

```

## Reshaping data with .melt()
```python

```

## Select the right .melt() arguments
```python

```

## Using .melt() to reshape government data
```python

```

## Using .melt() for stocks vs bond performance
```python

```

## Course wrap-up
```python

```

