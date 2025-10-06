---
title: Cleaning an Orders Dataset with PySpark
tags: analytics, python
url: https://app.datacamp.com/learn/projects/2355
---

```python
from pyspark.sql import (
    SparkSession,
    types,
    functions as F,
)

# Initiate a Spark session
spark = (
    SparkSession
    .builder
    .appName('cleaning_orders_dataset_with_pyspark')
    .getOrCreate()
)

# IMPORT DATA

# Read data from the parquet file
orders_data = spark.read.parquet('orders_data.parquet')

# DATA CLEANING AND PREPROCESSING

orders_data = (
    orders_data
    # Create a new column time_of_day
    .withColumn(
        'time_of_day',
        # When/otherwise (similar to case/when/else) statements extracting hour from timestamp
        F.when((F.hour('order_date') >= 0) & (F.hour('order_date') <= 5), 'night')
         .when((F.hour('order_date') >= 6) & (F.hour('order_date') <= 11), 'morning')
         .when((F.hour('order_date') >= 12) & (F.hour('order_date') <= 17), 'afternoon')
         .when((F.hour('order_date') >= 18) & (F.hour('order_date') <= 23), 'evening')
        # You can keep the otherwise statement as None to validate whether the conditions are exhaustive
         .otherwise(None)
    )
    # Filter by time of day
    .filter(
        F.col('time_of_day') != 'night'
    )
    # Cast order_date to date as it is originally a timestamp
    .withColumn(
        'order_date',
        F.col('order_date').cast(types.DateType())
    )
)


orders_data = (
    orders_data
    # Make product and category columns lowercase
    .withColumn(
        'product',
        F.lower('product')
    )
    .withColumn(
        'category',
        F.lower('category')
    )
    # Remove rows where product column contains "tv" (as you have already made it lowercase)
    .filter(
        ~F.col('product').contains('tv')
    )
)


orders_data = (
    orders_data
    # First you split the purchase address by space (" ")
    .withColumn(
        'address_split',
        F.split('purchase_address', ' ')
    )
    # If you look at the address lines, you can see that the state abbreviation is always at the 2nd last position
    .withColumn(
        'purchase_state',
        F.col('address_split').getItem(F.size('address_split') - 2)
    )
    # Dropping address_split columns as it is a temporary technical column
    .drop('address_split')
)

# Use distinct and count to calculate the number of unique values
n_states = (
    orders_data
    .select('purchase_state')
    .distinct()
    .count()
)


# EXPORT

# Export the resulting table to parquet format with the new name
(
    orders_data
    .write
    .parquet(
        'orders_data_clean.parquet',
        mode='overwrite',
    )
)
```
