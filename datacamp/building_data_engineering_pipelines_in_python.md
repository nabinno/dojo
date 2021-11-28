---
title: Building Data Engineering Pipelines in Python
tags: python,data-engineering
url: https://www.datacamp.com/courses/building-data-engineering-pipelines-in-python
---

# 1. Ingesting Data
## The data catalog
```python
In: type(catalog['diaper_reviews'].read())
Out: pyspark.sql.dataframe.DataFrame
```

## Working with JSON
```python
# Import json
import json

database_address = {
  "host": "10.0.0.5",
  "port": 8456
}

# Open the configuration file in writable mode
with open("database_config.json", "w") as fh:
  # Serialize the object in this file handle
  json.dump(obj=database_address, fp=fh)
```

## Specifying the schema of the data
```python
# Complete the JSON schema
schema = {'properties': {
    'brand': {'type': 'string'},
    'model': {'type': 'string'},
    'price': {'type': 'number'},
    'currency': {'type': 'string'},
    'quantity': {'type': 'number', 'minimum': 1},  
    'date': {'type': 'string', 'format': 'date'},
    'countrycode': {'type': 'string', 'pattern': "^[A-Z]{2}$"}, 
    'store_name': {'type': 'string'}}}

# Write the schema
singer.write_schema(stream_name='products', schema=schema, key_properties=[])
```

## Communicating with an API
```python
endpoint = "http://localhost:5000"

# Fill in the correct API key
api_key = "scientist007"

# Create the web API’s URL
authenticated_endpoint = "{}/{}".format(endpoint, api_key)

# Get the web API’s reply to the endpoint
api_response = requests.get(authenticated_endpoint).json()
pprint.pprint(api_response)

# Create the API’s endpoint for the shops
shops_endpoint = "{}/{}/{}/{}".format(endpoint, api_key, "diaper/api/v1.0", "shops")
shops = requests.get(shops_endpoint).json()
print(shops)

# Create the API’s endpoint for items of the shop starting with a "D"
items_of_specific_shop_URL = "{}/{}/{}/{}/{}".format(endpoint, api_key, "diaper/api/v1.0", "items", "DM")
products_of_shop = requests.get(items_of_specific_shop_URL).json()
pprint.pprint(products_of_shop)
```

## Streaming records
```python
# Use the convenience function to query the API
tesco_items = retrieve_products("Tesco")

singer.write_schema(stream_name="products", schema=schema,
                    key_properties=[])

# Write a single record to the stream, that adheres to the schema
singer.write_record(stream_name="products", 
                    record={**tesco_items[0], "store_name": "Tesco"})

for shop in requests.get(SHOPS_URL).json()["shops"]:
    # Write all of the records that you retrieve from the API
    singer.write_records(
      stream_name="products", # Use the same stream name that you used in the schema
      records=({**item, "store_name": shop}
               for item in retrieve_products(shop))
    )
```

## Chain taps and targets
```sh
tap-marketing-api | target-csv --config ingest/data_lake.conf
```


# 2. Creating a data transformation pipeline with PySpark
## Reading a CSV file
```python
# Read a csv file and set the headers
df = (spark.read
      .options(header=True)
      .csv("/home/repl/workspace/mnt/data_lake/landing/ratings.csv"))

df.show()
```

## Defining a schema
```python
# Define the schema
schema = StructType([
  StructField("brand", StringType(), nullable=False),
  StructField("model", StringType(), nullable=False),
  StructField("absorption_rate", IntegerType(), nullable=True),
  StructField("comfort", IntegerType(), nullable=True)
])

better_df = (spark
             .read
             .options(header="true")
             # Pass the predefined schema to the Reader
             .schema(schema)
             .csv("/home/repl/workspace/mnt/data_lake/landing/ratings.csv"))
pprint(better_df.dtypes)
```

## Removing invalid rows
```python
# Specify the option to drop invalid rows
ratings = (spark
           .read
           .options(header=True, mode="DROPMALFORMED")
           .csv("/home/repl/workspace/mnt/data_lake/landing/ratings_with_invalid_rows.csv"))
ratings.show()
```

## Filling unknown data
```python
print("BEFORE")
ratings.show()

print("AFTER")
# Replace nulls with arbitrary value on column subset
ratings = ratings.fillna(4, subset=["comfort"])
ratings.show()
```

## Conditionally replacing values
```python
from pyspark.sql.functions import col, when

# Add/relabel the column
categorized_ratings = ratings.withColumn(
    "comfort",
    # Express the condition in terms of column operations
    when(col("comfort") > 3, "sufficient").otherwise("insufficient"))

categorized_ratings.show()
```

## Selecting and renaming columns
```python
from pyspark.sql.functions import col

# Select the columns and rename the "absorption_rate" column
result = ratings.select([col("brand"),
                       col("model"),
                       col("absorption_rate").alias("absorbency")])

# Show only unique values
result.distinct().show()
```

## Grouping and aggregating data
```python
from pyspark.sql.functions import col, avg, stddev_samp, max as sfmax

aggregated = (purchased
              # Group rows by 'Country'
              .groupBy(col('Country'))
              .agg(
                # Calculate the average salary per group and rename
                avg('Salary').alias('average_salary'),
                # Calculate the standard deviation per group
                stddev_samp('Salary'),
                # Retain the highest salary per group and rename
                sfmax('Salary').alias('highest_salary')
              )
             )

aggregated.show()
```

## Creating a deployable artifact
```python
zip --recurse-paths pydiaper.zip pydiaper
```

## Submitting your Spark job
```python
# You need to run this command in the terminal
spark-submit --py-files spark_pipelines/pydiaper/pydiaper.zip ./spark_pipelines/pydiaper/pydiaper/cleaning/clean_ratings.py
```


# 3. Testing your data pipeline
## Characteristics of tests
```python
## Unit Test
Can typically be written in a few minutes.
They typically run in milliseconds.
Lowercasing a column in a table is an example of this.
You typically have most of these.
You should be able to run these on your local laptop without configuration, aside from installing a test framework.

## Service Test
Requires a bit of setup to have each service configured in the desired state.
Tests the interaction between a few services.
A database writing out to a file is an example.

## End-to-end Test
Most difficult to debug.
Combine the services of many systems.
Subscribing for a newsletter and receiving e-mail confirmation from this is an example.
Is closest to what an end-user experiences.
```

## Writing unit tests for PySpark
```python

```

## Creating in-memory DataFrames
```python

```

## Making a function more widely reusable
```python

```

## Continuous testing
```python

```

## A high-level view on CI/CD
```python

```

## Understanding the output of pytest
```python

```

## Improving style guide compliancy
```python

```


# 4. Managing and orchestrating a workflow
## Modern day workflow management
```python

```

## Specifying the DAG schedule
```python

```

## Setting up daily tasks
```python

```

## Specifying operator dependencies
```python

```

## Building a data pipeline with Airflow
```python

```

## Preparing a DAG for daily pipelines
```python

```

## Scheduling bash scripts with Airflow
```python

```

## Scheduling Spark jobs with Airflow
```python
# Import the operator
from airflow.contrib.operators.spark_submit_operator import SparkSubmitOperator

# Set the path for our files.
entry_point = os.path.join(os.environ["AIRFLOW_HOME"], "scripts", "clean_ratings.py")
dependency_path = os.path.join(os.environ["AIRFLOW_HOME"], "dependencies", "pydiaper.zip")

with DAG('data_pipeline', start_date=datetime(2019, 6, 25),
         schedule_interval='@daily') as dag:
  	# Define task clean, running a cleaning job.
    clean_data = SparkSubmitOperator(
        application=entry_point, 
        py_files=dependency_path,
        task_id='clean_data',
        conn_id='spark_default')
```

## Scheduling the full data pipeline with Airflow
```python

```

## Deploying Airflow
```python

```

## Airflow’s executors
```python

```

## Recovering from deployed but broken DAGs
```python

```

## Running tests on Airflow
```python

```

## Final thoughts
```python

```

