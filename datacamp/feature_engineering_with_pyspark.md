---
title: Feature Engineering with PySpark
tags: python,database
url: https://www.datacamp.com/courses/feature-engineering-with-pyspark
---

# 1. Exploratory Data Analysis
## Check Version
```python
# Return spark version
print(spark.version)

# Return python version
import sys
print(sys.version_info)
```

## Load in the data
```python
# Read the file into a dataframe
df = spark.read.parquet('Real_Estate.parq')
# Print columns in dataframe
print(df.columns)
```

## What are we predicting?
```python
# Select our dependent variable
Y_df = df.select(['SALESCLOSEPRICE'])

# Display summary statistics
Y_df.describe().show()
```

## Verifying Data Load
```python
def check_load(df, num_records, num_columns):
  # Takes a dataframe and compares record and column counts to input
  # Message to return if the critera below aren't met
  message = 'Validation Failed'
  # Check number of records
  if num_records == df.count():
    # Check number of columns
    if num_columns == len(df.columns):
      # Success message
      message = 'Validation Passed'
  return message

# Print the data validation message
print(check_load(df, 5000, 74))
```

## Verifying DataTypes
```python
# create list of actual dtypes to check
actual_dtypes_list = df.dtypes
print(actual_dtypes_list)

# Iterate through the list of actual dtypes tuples
for attribute_tuple in actual_dtypes_list:
  
  # Check if column name is dictionary of expected dtypes
  col_name = attribute_tuple[0]
  if col_name in validation_dict.keys():

    # Compare attribute types
    col_type = attribute_tuple[1]
    if col_type == validation_dict[col_name]:
      print(col_name + ' has expected dtype.')
```

## Using Corr()
```python
# Name and value of col with max corr
corr_max = 0
corr_max_col = columns[0]

# Loop to check all columns contained in list
for col in columns:
    # Check the correlation of a pair of columns
    corr_val = df.corr(col, 'SALESCLOSEPRICE')
    # Logic to compare corr_max with current corr_val
    if corr_val > corr_max:
        # Update the column name and corr value
        corr_max = corr_val
        corr_max_col = col

print(corr_max_col)
```

## Using Visualizations: distplot
```python
# Select a single column and sample and convert to pandas
sample_df = df.select(['LISTPRICE']).sample(False, 0.5, 42)
pandas_df = sample_df.toPandas()

# Plot distribution of pandas_df and display plot
sns.distplot(pandas_df)
plt.show()

# Import skewness function
from pyspark.sql.functions import skewness

# Compute and print skewness of LISTPRICE
print(df.agg({'LISTPRICE': 'skewness'}).collect())
```

## Using Visualizations: lmplot
```python
# Select a the relevant columns and sample
sample_df = df\
    .select(['SALESCLOSEPRICE', 'LIVINGAREA'])\
    .sample(False, 0.5, 42)

# Convert to pandas dataframe
pandas_df = sample_df.toPandas()

# Linear model plot of pandas_df
sns.lmplot(x='LIVINGAREA', y='SALESCLOSEPRICE', data=pandas_df)
plt.show()
```

# 2. Wrangling with Spark Functions
## Dropping data
```python

```

## Dropping a list of columns
```python

```

## Using text filters to remove records
```python

```

## Filtering numeric fields conditionally
```python

```

## Adjusting Data
```python

```

## Custom Percentage Scaling
```python

```

## Scaling your scalers
```python

```

## Correcting Right Skew Data
```python

```

## Working with Missing Data
```python

```

## Visualizing Missing Data
```python

```

## Imputing Missing Data
```python

```

## Calculate Missing Percents
```python

```

## Getting More Data
```python

```

## A Dangerous Join
```python

```

## Spark SQL Join
```python

```

## Checking for Bad Joins
```python

```

# 3. Feature Engineering
## Feature Generation
```python

```

## Differences
```python

```

## Ratios
```python

```

## Deeper Features
```python

```

## Time Features
```python

```

## Time Components
```python

```

## Joining On Time Components
```python

```

## Date Math
```python

```

## Extracting Features
```python

```

## Extracting Text to New Features
```python

```

## Splitting & Exploding
```python

```

## Pivot & Join
```python

```

## Binarizing, Bucketing & Encoding
```python

```

## Binarizing Day of Week
```python

```

## Bucketing
```python

```

## One Hot Encoding
```python

```

# 4. Building a Model
## Choosing the Algorithm
```python

```

## Which MLlib Module?
```python

```

## Creating Time Splits
```python

```

## Adjusting Time Features
```python

```

## Feature Engineering Assumptions for RFR
```python

```

## Feature Engineering For Random Forests
```python

```

## Dropping Columns with Low Observations
```python

```

## Naively Handling Missing and Categorical Values
```python

```

## Building a Model
```python

```

## Building a Regression Model
```python

```

## Evaluating & Comparing Algorithms
```python

```

## Understanding Metrics
```python

```

## Interpreting, Saving & Loading
```python

```

## Interpreting Results
```python

```

## Saving & Loading Models
```python

```

## Final Thoughts
```python

```




