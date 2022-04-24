---
title: Data Manipulation with pandas
tags: pandas, python, analytics
url: https://campus.datacamp.com/courses/data-manipulation-with-pandas
---

# 1. Transforming DataFrames
## Inspecting a DataFrame
```python
# Print the head of the homelessness data
print(homelessness.head())

# Print information about homelessness
print(homelessness.info())

# Print the shape of homelessness
print(homelessness.shape)

# Print a description of homelessness
print(homelessness.describe())
```

## Parts of a DataFrame
```python
# Import pandas using the alias pd
import pandas as pd

# Print the values of homelessness
print(homelessness.values)

# Print the column index of homelessness
print(homelessness.columns)

# Print the row index of homelessness
print(homelessness.index)
```

## Sorting rows
```python
##
# Sort homelessness by individuals
homelessness_ind = homelessness.sort_values("individuals")

# Print the top few rows
print(homelessness_ind.head())

##
# Sort homelessness by descending family members
homelessness_fam = homelessness.sort_values("family_members", ascending=False)

# Print the top few rows
print(homelessness_fam.head())

##
# Sort homelessness by region, then descending family members
homelessness_reg_fam = homelessness.sort_values(["region", "family_members"], ascending=[True, False])

# Print the top few rows
print(homelessness_reg_fam.head())
```

## Subsetting columns
```python
##
# Select the individuals column
individuals = homelessness["individuals"]

# Print the head of the result
print(individuals.head())

##
# Select the state and family_members columns
state_fam = homelessness[["state", "family_members"]]

# Print the head of the result
print(state_fam.head())

##
# Select only the individuals and state columns, in that order
ind_state = homelessness[["individuals", "state"]]

# Print the head of the result
print(ind_state.head())
```

## Subsetting rows
```python
##
# Filter for rows where individuals is greater than 10000
ind_gt_10k = homelessness[homelessness["individuals"] > 10000]

# See the result
print(ind_gt_10k)

##
# Filter for rows where region is Mountain
mountain_reg = homelessness[homelessness["region"] == "Mountain"]

# See the result
print(mountain_reg)

##
# Filter for rows where family_members is less than 1000 
# and region is Pacific
fam_lt_1k_pac = homelessness[(homelessness["family_members"] < 1000) & (homelessness["region"] == "Pacific")]

# See the result
print(fam_lt_1k_pac)
```

## Subsetting rows by categorical variables
```python
##
# Subset for rows in South Atlantic or Mid-Atlantic regions
south_mid_atlantic = homelessness[(homelessness["region"] == "South Atlantic") | (homelessness["region"] == "Mid-Atlantic")]

# See the result
print(south_mid_atlantic)

##
# The Mojave Desert states
canu = ["California", "Arizona", "Nevada", "Utah"]

# Filter for rows in the Mojave Desert states
mojave_homelessness = homelessness[homelessness["state"].isin(canu)]

# See the result
print(mojave_homelessness)
```

## Adding new columns
```python
# Add total col as sum of individuals and family_members
homelessness["total"] = homelessness["individuals"] + homelessness["family_members"]

# Add p_individuals col as proportion of total that are individuals
homelessness["p_individuals"] = homelessness["individuals"] / homelessness["total"]

# See the result
print(homelessness)
```

## Combo-attack!
```python

```





# 2. Aggregating DataFrames
## Summary statistics
```python

```

## Mean and median
```python

```

## Summarizing dates
```python

```

## Efficient summaries
```python

```

## Cumulative statistics
```python

```

## Counting
```python

```

## Dropping duplicates
```python

```

## Counting categorical variables
```python

```

## Grouped summary statistics
```python

```

## What percent of sales occurred at each store type?
```python

```

## Calculations with .groupby()
```python

```

## Multiple grouped summaries
```python

```

## Pivot tables
```python

```

## Pivoting on one variable
```python

```

## Fill in missing values and sum values with pivot tables
```python

```




# 3. Slicing and Indexing DataFrames
## Explicit indexes
```python

```

## Setting and removing indexes
```python

```

## Subsetting with .loc[]
```python

```

## Setting multi-level indexes
```python

```

## Sorting by index values
```python

```

## Slicing and subsetting with .loc and .iloc
```python

```

## Slicing index values
```python

```

## Slicing in both directions
```python

```

## Slicing time series
```python

```

## Subsetting by row/column number
```python

```

## Working with pivot tables
```python

```

## Pivot temperature by city and year
```python

```

## Subsetting pivot tables
```python

```

## Calculating on a pivot table
```python

```




# 4. Creating and Visualizing DataFrames
## Visualizing your data
```python

```

## Which avocado size is most popular?
```python

```

## Changes in sales over time
```python

```

## Avocado supply and demand
```python

```

## Price of conventional vs. organic avocados
```python

```

## Missing values
```python

```

## Finding missing values
```python

```

## Removing missing values
```python

```

## Replacing missing values
```python

```

## Creating DataFrames
```python

```

## List of dictionaries
```python

```

## Dictionary of lists
```python

```

## Reading and writing CSVs
```python

```

## CSV to DataFrame
```python

```

## DataFrame to CSV
```python

```

## Wrap-up
```python

```

