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
## Dropping a list of columns
```python
# Show top 30 records
df.show(30)

# List of columns to remove from dataset
cols_to_drop = ['STREETNUMBERNUMERIC', 'LOTSIZEDIMENSIONS']

# Drop columns in list
df = df.drop(*cols_to_drop)
```

## Using text filters to remove records
```python
# Inspect unique values in the column 'ASSUMABLEMORTGAGE'
df.select(['ASSUMABLEMORTGAGE']).distinct().show()

# List of possible values containing 'yes'
yes_values = ['Yes w/ Qualifying', 'Yes w/No Qualifying']

# Filter the text values out of df but keep null values
text_filter = ~df['ASSUMABLEMORTGAGE'].isin(yes_values) | df['ASSUMABLEMORTGAGE'].isNull()
df = df.where(text_filter)

# print count of remaining records
print(df.count())
```

## Filtering numeric fields conditionally
```python
from pyspark.sql.functions import mean, stddev

# Calculate values used for outlier filtering
mean_val = df.agg({'log_SalesClosePrice': 'mean'}).collect()[0][0]
stddev_val = df.agg({'log_SalesClosePrice': 'stddev'}).collect()[0][0]

# Create three standard deviation (n ± 3d) lower and upper bounds for data
low_bound = mean_val - (3 * stddev_val)
hi_bound = mean_val + (3 * stddev_val)

# Filter the data to fit between the lower and upper bounds
df = df.where((df['log_SalesClosePrice'] < hi_bound) & (df['log_SalesClosePrice'] > low_bound))
```

## Custom Percentage Scaling
```python
# Define max and min values and collect them
max_days = df.agg({'DAYSONMARKET': 'max'}).collect()[0][0]
min_days = df.agg({'DAYSONMARKET': 'min'}).collect()[0][0]

# Create a new column based off the scaled data
df = df.withColumn(
    'percentage_scaled_days', 
    round((df['DAYSONMARKET'] - min_days) / (max_days - min_days)) * 100
)

# Calc max and min for new column
print(df.agg({'percentage_scaled_days': 'max'}).collect())
print(df.agg({'percentage_scaled_days': 'min'}).collect())
```

## Scaling your scalers
```python
def min_max_scaler(df, cols_to_scale):
    # Takes a dataframe and list of columns to minmax scale. Returns a dataframe.
    for col in cols_to_scale:
        # Define min and max values and collect them
        max_days = df.agg({col: 'max'}).collect()[0][0]
        min_days = df.agg({col: 'min'}).collect()[0][0]
        new_column_name = 'scaled_' + col
        # Create a new column based off the scaled data
        df = df.withColumn(
            new_column_name, 
            (df[col] - min_days) / (max_days - min_days)
        )
    return df

df = min_max_scaler(df, cols_to_scale)
# Show that our data is now between 0 and 1
df[['DAYSONMARKET', 'scaled_DAYSONMARKET']].show()
```

## Correcting Right Skew Data
```python
from pyspark.sql.functions import log

# Compute the skewness
print(df.agg({'YEARBUILT': 'skewness'}).collect())

# Calculate the max year
max_year = df.agg({'YEARBUILT': 'max'}).collect()[0][0]

# Create a new column of reflected data
df = df.withColumn('Reflect_YearBuilt', (max_year + 1) - df['YEARBUILT'])

# Create a new column based reflected data
df = df.withColumn('adj_yearbuilt', 1 / log(df['Reflect_YearBuilt']))
```

## Visualizing Missing Data
```python
# Sample the dataframe and convert to Pandas
sample_df = df\
    .select(columns)\
    .sample(False, 0.1, 42)
pandas_df = sample_df.toPandas()

# Convert all values to T/F
tf_df = pandas_df.isnull()

# Plot it
sns.heatmap(data=tf_df)
plt.xticks(rotation=30, fontsize=10)
plt.yticks(rotation=0, fontsize=10)
plt.show()

# Set the answer to the column with the most missing data
answer = 'BACKONMARKETDATE'
```

## Imputing Missing Data
```python
# Count missing rows
missing = df\
    .where(df['PDOM'].isNull())\
    .count()

# Calculate the mean value
col_mean = df\
    .agg({'PDOM': 'mean'})\
    .collect()[0][0]

# Replacing with the mean value for that column
df.fillna(20.792025646163093, subset=['PDOM'])
```

## Calculate Missing Percents
```python
def column_dropper(df, threshold):
    # Takes a dataframe and threshold for missing values. Returns a dataframe.
    total_records = df.count()
    for col in df.columns:
        # Calculate the percentage of missing values
        missing = df.where(df[col].isNull()).count()
        missing_percent = missing / total_records
        # Drop column if percent of missing is more than threshold
        if missing_percent > threshold:
            df = df.drop(col)
    return df

# Drop columns that are more than 60% missing
df = column_dropper(df, .6)
```

## A Dangerous Join
```python
# Cast data types
walk_df = walk_df\
    .withColumn('longitude', walk_df['longitude'].cast('double'))
walk_df = walk_df\
    .withColumn('latitude', walk_df['latitude'].cast('double'))

# Round precision
df = df.withColumn('latitude', round('latitude', 5))
df = df.withColumn('longitude', round('longitude', 5))

# Create join condition
condition = [
    walk_df['latitude'] == df['latitude'],
    walk_df['longitude'] == df['longitude']
]

# Join the dataframes together
join_df = df.join(walk_df, on=condition, how='left')
# Count non-null records from new field
print(join_df.where(~join_df['walkscore'].isNull()).count())
```

## Spark SQL Join
```python
# Register dataframes as tables
df.createOrReplaceTempView("df")
walk_df.createOrReplaceTempView("walk_df")

# SQL to join dataframes
join_sql = 	"""
			SELECT 
				*
			FROM df
			LEFT JOIN walk_df
			ON df.longitude = walk_df.longitude
			AND df.latitude = walk_df.latitude
			"""
# Perform sql join
joined_df = spark.sql(join_sql)
```

## Checking for Bad Joins
```python
# Join on mismatched keys precision 
wrong_prec_cond = [
	df_orig.longitude == walk_df.longitude,
	df_orig.latitude == walk_df.latitude
]
wrong_prec_df = df_orig.join(walk_df, on=wrong_prec_cond, how='left')

# Compare bad join to the correct one
print(
	wrong_prec_df\
		.where(wrong_prec_df['walkscore'].isNull())\
		.count()
)
print(
	correct_join_df\
		.where(correct_join_df['walkscore'].isNull())\
		.count()
)

# Create a join on too few keys
few_keys_cond = [df['longitude'] == walk_df['longitude']]
few_keys_df = df.join(walk_df, on=few_keys_cond, how='left')

# Compare bad join to the correct one
print("Record Count of the Too Few Keys Join Example: " + str(few_keys_df.count()))
print("Record Count of the Correct Join Example: " + str(correct_join_df.count()))
```

# 3. Feature Engineering
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




