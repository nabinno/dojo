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

```

## Running an ingestion pipeline with Singer
```python

```

## Properly propagating state
```python

```

## Communicating with an API
```python

```

## Streaming records
```python

```

## Chain taps and targets
```python

```


# 2. Creating a data transformation pipeline with PySpark
## Basic introduction to PySpark
```python

```

## Reading a CSV file
```python

```

## Defining a schema
```python

```

## Cleaning data
```python

```

## Sensible data types
```python

```

## Removing invalid rows
```python

```

## Filling unknown data
```python

```

## Conditionally replacing values
```python

```

## Transforming data with Spark
```python

```

## Selecting and renaming columns
```python

```

## Grouping and aggregating data
```python

```

## Packaging your application
```python

```

## Creating a deployable artifact
```python

```

## Submitting your Spark job
```python

```

## Debugging simple errors
```python

```

## Verifying your pipeline’s output
```python

```


# 3. Testing your data pipeline
## On the importance of tests
```python

```

## Regression errors
```python

```

## Characteristics of tests
```python

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

