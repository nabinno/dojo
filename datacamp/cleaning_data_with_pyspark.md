---
title: Cleaning Data with PySpark
tags: python,data-pre-processing
url: https://www.datacamp.com/courses/cleaning-data-with-apache-spark-in-python
---

# 1. DataFrame details
## Data cleaning review
```python
# Import the pyspark.sql.types library
from pyspark.sql.types import *

# Define a new schema using the StructType method
people_schema = StructType([
  # Define a StructField for each field
  StructField('name', StringType(), False),
  StructField('age', IntegerType(), False),
  StructField('city', StringType(), False)
])
```

## Using lazy processing
```python
# Load the CSV file
aa_dfw_df = spark.read.format('csv').options(Header=True).load('AA_DFW_2018.csv.gz')

# Add the airport column using the F.lower() method
aa_dfw_df = aa_dfw_df.withColumn('airport', F.lower(aa_dfw_df['Destination Airport']))

# Drop the Destination Airport column
aa_dfw_df = aa_dfw_df.drop(aa_dfw_df['Destination Airport'])

# Show the DataFrame
aa_dfw_df.show()
```

## Saving a DataFrame in Parquet format
```python
# View the row count of df1 and df2
print("df1 Count: %d" % df1.count())
print("df2 Count: %d" % df2.count())

# Combine the DataFrames into one
df3 = df1.union(df2)

# Save the df3 DataFrame in Parquet format
df3.write.parquet('AA_DFW_ALL.parquet', mode='overwrite')

# Read the Parquet file into a new DataFrame and run a count
print(spark.read.parquet('AA_DFW_ALL.parquet').count())
```

## SQL and Parquet
```python
# Read the Parquet file into flights_df
flights_df = spark.read.parquet('AA_DFW_ALL.parquet')

# Register the temp table
flights_df.createOrReplaceTempView('flights')

# Run a SQL query of the average flight duration
avg_duration = spark\
    .sql("SELECT avg(flight_duration) from flights")\
    .collect()[0]
print('The average flight time is: %d' % avg_duration)
```

# 2. Manipulating DataFrames in the real wold
## DataFrame column operations
```python
# Show the distinct VOTER_NAME entries
voter_df\
    .select('VOTER_NAME')\
    .distinct()\
    .show(40, truncate=False)

# Filter voter_df where the VOTER_NAME is 1-20 characters in length
voter_df = voter_df.filter('length(VOTER_NAME) > 0 and length(VOTER_NAME) < 20')

# Filter out voter_df where the VOTER_NAME contains an underscore
voter_df = voter_df.filter(~ F.col('VOTER_NAME').contains('_'))

# Show the distinct VOTER_NAME entries again
voter_df\
    .select('VOTER_NAME')\
    .distinct()\
    .show(40, truncate=False)
```

## Modifying DataFrame columns
```python
# Add a new column called splits separated on whitespace
voter_df = voter_df.withColumn('splits', F.split(voter_df.VOTER_NAME, '\s+'))

# Create a new column called first_name based on the first item in splits
voter_df = voter_df.withColumn('first_name', voter_df.splits.getItem(0))

# Get the last entry of the splits list and create a column called last_name
voter_df = voter_df.withColumn('last_name', voter_df.splits.getItem(F.size('splits') - 1))

# Drop the splits column
voter_df = voter_df.drop('splits')

# Show the voter_df DataFrame
voter_df.show()
```

## when() example
```python
# Add a column to voter_df for any voter with the title **Councilmember**
voter_df = voter_df\
    .withColumn(
        'random_val',
        when(voter_df.TITLE == 'Councilmember', F.rand())
    )

# Show some of the DataFrame rows, noting whether the when clause worked
voter_df.show()
```

## When / Otherwise
```python
# Add a column to voter_df for a voter based on their position
voter_df = voter_df\
    .withColumn(
        'random_val',
        when(voter_df.TITLE == 'Councilmember', F.rand())
        .when(voter_df.TITLE == 'Mayor', 2)
        .otherwise(0)
    )

# Show some of the DataFrame rows
voter_df.show()

# Use the .filter() clause with random_val
voter_df.filter(voter_df.random_val == 0).show()
```

## Using user defined functions in Spark
```python
def getFirstAndMiddle(names):
  # Return a space separated string of names
  return ' '.join(names)

# Define the method as a UDF
udfFirstAndMiddle = F.udf(getFirstAndMiddle, StringType())

# Create a new column using your UDF
voter_df = voter_df.withColumn('first_and_middle_name', udfFirstAndMiddle(voter_df.splits))

# Drop the unecessary columns then show the DataFrame
voter_df = voter_df.drop('first_name')
voter_df = voter_df.drop('splits')
voter_df.show()
```

## Adding an ID Field
```python
# Select all the unique council voters
voter_df = df.select(df["VOTER NAME"]).distinct()

# Count the rows in voter_df
print("\nThere are %d rows in the voter_df DataFrame.\n" % voter_df.count())

# Add a ROW_ID
voter_df = voter_df.withColumn('ROW_ID', F.monotonically_increasing_id())

# Show the rows with 10 highest IDs in the set
voter_df.orderBy(voter_df.ROW_ID.desc()).show(10)
```

## IDs with different partitions
```python
# Print the number of partitions in each DataFrame
print("\nThere are %d partitions in the voter_df DataFrame.\n" % voter_df.rdd.getNumPartitions())
print("\nThere are %d partitions in the voter_df_single DataFrame.\n" % voter_df_single.rdd.getNumPartitions())

# Add a ROW_ID field to each DataFrame
voter_df = voter_df.withColumn('ROW_ID', F.monotonically_increasing_id())
voter_df_single = voter_df_single.withColumn('ROW_ID', F.monotonically_increasing_id())

# Show the top 10 IDs in each DataFrame 
voter_df.orderBy(voter_df.ROW_ID.desc()).show(10)
voter_df_single.orderBy(voter_df_single.ROW_ID.desc()).show(10)
```

## More ID tricks
```python
# Determine the highest ROW_ID and save it in previous_max_ID
previous_max_ID = voter_df_march\
    .select('ROW_ID')\
    .rdd\
    .max()[0]

# Add a ROW_ID column to voter_df_april starting at the desired value
voter_df_april = voter_df_april\
    .withColumn(
        'ROW_ID',
        F.monotonically_increasing_id() + previous_max_ID
    )

# Show the ROW_ID from both DataFrames and compare
voter_df_march.select('ROW_ID').show()
voter_df_april.select('ROW_ID').show()
```

# 3. Improving Performance
## Caching
```python

```

## Caching a DataFrame
```python

```

## Removing a DataFrame from cache
```python

```

## Improve import performance
```python

```

## File size optimization
```python

```

## File import performance
```python

```

## Cluster configurations
```python

```

## Reading Spark configurations
```python

```

## Writing Spark configurations
```python

```

## Performance improvements
```python

```

## Normal joins
```python

```

## Using broadcasting on Spark joins
```python

```

## Comparing broadcast vs normal joins
```python

```

# 4. Complex processing and data pipelines
## Introduction to data pipelines
```python

```

## Quick pipeline
```python

```

## Pipeline data issue
```python

```

## Data handling techniques
```python

```

## Removing commented lines
```python

```

## Removing invalid rows
```python

```

## Splitting into columns
```python

```

## Further parsing
```python

```

## Data validation
```python

```

## Validate rows via join
```python

```

## Examining invalid rows
```python

```

## Final analysis and delivery
```python

```

## Dog parsing
```python

```

## Per image count
```python

```

## Percentage dog pixels
```python

```

## Congratulations and next steps
```python

```


