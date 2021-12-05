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

## Creating in-memory DataFrames
```python
from datetime import date
from pyspark.sql import Row

Record = Row("country", "utm_campaign", "airtime_in_minutes", "start_date", "end_date")

# Create a tuple of records
data = (
  Record("USA", "DiapersFirst", 28, date(2017, 1, 20), date(2017, 1, 27)),
  Record("Germany", "WindelKind", 31, date(2017, 1, 25), None),
  Record("India", "CloseToCloth", 32, date(2017, 1, 25), date(2017, 2, 2))
)

# Create a DataFrame from these records
frame = spark.createDataFrame(data)
frame.show()
```

## Making a function more widely reusable
```python
## chinese_provinces_improved.py
from pyspark.sql import SparkSession
from pyspark.sql.functions import col, lower, sum

from .catalog import catalog


def extract_demographics(sparksession, catalog):
    return sparksession.read.parquet(catalog["clean/demographics"])


def store_chinese_demographics(frame, catalog):
    frame.write.parquet(catalog["business/chinese_demographics"])


# Improved aggregation function, grouped by country and province
def aggregate_inhabitants_by_province(frame):
    return (frame
            .groupBy("country", "province")
            .agg(sum(col("inhabitants")).alias("inhabitants"))
            )


def main():
    spark = SparkSession.builder.getOrCreate()
    frame = extract_demographics(spark, catalog)
    chinese_demographics = frame.filter(lower(col("country")) == "china")
    aggregated_demographics = aggregate_inhabitants_by_province(chinese_demographics)
    store_chinese_demographics(aggregated_demographics, catalog)


if __name__ == "__main__":
    main()
```

## A high-level view on CI/CD
```python
Check out your application from version control.
Install your Python application's dependecies.
Run the test suite of your application.
Create artifacts (jar, egg, wheel, documentation, ...).
Save the artifacts to a location accessible by your company's compute infrastructure.
```

## Improving style guide compliancy
```yml:config.yml
version: 2
jobs:
  build:
    working_directory: ~/data_scientists/optimal_diapers/
    docker:
      - image: gcr.io/my-companys-container-registry-on-google-cloud-123456/python:3.6.4
    steps:
      - checkout
      - run:
          command: |
            sudo pip install pipenv
            pipenv install
      - run:
          command: |
            pipenv run flake8 .
            pipenv run pytest .
      - store_test_results:
          path: test-results
      - store_artifacts:
          path: test-results
          destination: tr1
```


# 4. Managing and orchestrating a workflow
## Specifying the DAG schedule
```python
from datetime import datetime
from airflow import DAG

reporting_dag = DAG(
    dag_id="publish_EMEA_sales_report", 
    # Insert the cron expression
    schedule_interval="0 7 * * 1",
    start_date=datetime(2019, 11, 24),
    default_args={"owner": "sales"}
)
```

## Specifying operator dependencies
```python
# Specify direction using verbose method
prepare_crust.set_downstream(apply_tomato_sauce)

tasks_with_tomato_sauce_parent = [add_cheese, add_ham, add_olives, add_mushroom]
for task in tasks_with_tomato_sauce_parent:
    # Specify direction using verbose method on relevant task
    apply_tomato_sauce.set_downstream(task)

# Specify direction using bitshift operator
tasks_with_tomato_sauce_parent >> bake_pizza

# Specify direction using verbose method
bake_pizza.set_upstream(prepare_oven)
```

## Preparing a DAG for daily pipelines
```python
# Create a DAG object
dag = DAG(
  dag_id='optimize_diaper_purchases',
  default_args={
    # Don't email on failure
    'email_on_failure': False,
    # Specify when tasks should have started earliest
    'start_date': datetime(2019, 6, 25)
  },
  # Run the DAG daily
  schedule_interval='@daily')
```

## Scheduling bash scripts with Airflow
```python
config = os.path.join(os.environ["AIRFLOW_HOME"], 
                      "scripts",
                      "configs", 
                      "data_lake.conf")

ingest = BashOperator(
  # Assign a descriptive id
  task_id="ingest_data", 
  # Complete the ingestion pipeline
  bash_command="target-csv --config %s" % config,
  dag=dag)
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
spark_args = {"py_files": dependency_path,
              "conn_id": "spark_default"}
# Define ingest, clean and transform job.
with dag:
    ingest = BashOperator(task_id='Ingest_data', bash_command='tap-marketing-api | target-csv --config %s' % config)
    clean = SparkSubmitOperator(application=clean_path, task_id='clean_data', **spark_args)
    insight = SparkSubmitOperator(application=transform_path, task_id='show_report', **spark_args)
    
    # set triggering sequence
    ingest >> clean >> insight
```

## Recovering from deployed but broken DAGs
```python
"""
An Apache Airflow DAG used in course material.

"""
from datetime import datetime, timedelta

from airflow import DAG
from airflow.models import Variable
from airflow.operators.bash_operator import BashOperator
from airflow.operators.python_operator import PythonOperator

default_args = {
    #    "owner": "squad-a",
    "depends_on_past": False,
    "start_date": datetime(2019, 7, 5),
    "email": ["foo@bar.com"],
    "email_on_failure": False,
    "email_on_retry": False,
    "retries": 1,
    "retry_delay": timedelta(minutes=5),
}

dag = DAG(
    "cleaning",
    default_args=default_args,
    user_defined_macros={"env": Variable.get("environment")},
    schedule_interval="0 5 */2 * *"
)


def say(what):
    print(what)


with dag:
    say_hello = BashOperator(task_id="say-hello", bash_command="echo Hello,")
    say_world = BashOperator(task_id="say-world", bash_command="echo World")
    shout = PythonOperator(task_id="shout",
                           python_callable=say,
                           op_kwargs={'what': '!'})

    say_hello >> say_world >> shout
```

## Running tests on Airflow
```python

```

## Final thoughts
```python

```

