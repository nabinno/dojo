---
title: Introduction to Airflow in Python
tags: python, apache-airflow
url: https://campus.datacamp.com/courses/object-oriented-programming-in-pythonhttps://campus.datacamp.com/courses/introduction-to-airflow-in-python/
---

# Intro to Airflow
## Running a task in Airflow
```python
$ airflow run etl_pipeline download_file 2020-01-08
```

## Defining a simple DAG
```python
# Import the DAG object
from airflow.models import DAG

# Define the default_args dictionary
default_args = {
  'owner': 'dsmith',
  'start_date': datetime(2020, 1, 14),
  'retries': 2
}

# Instantiate the DAG object
etl_dag = DAG('example_etl', default_args=default_args)
```

## Working with DAGs and the Airflow shell
```python

```

## Troubleshooting DAG creation
```python

```

## Airflow web interface
```python

```

## Starting the Airflow webserver
```python

```

## Navigating the Airflow UI
```python

```

## Examining DAGs with the Airflow UI
```python

```




# Implementing Airflow DAGs
## Airflow operators
```python

```

## Defining a BashOperator task
```python

```

## Multiple BashOperators
```python

```

## Airflow tasks
```python

```

## Define order of BashOperators
```python

```

## Determining the order of tasks
```python

```

## Troubleshooting DAG dependencies
```python

```

## Additional operators
```python

```

## Using the PythonOperator
```python

```

## More PythonOperators
```python

```

## EmailOperator and dependencies
```python

```

## Airflow scheduling
```python

```

## Schedule a DAG via Python
```python

```

## Deciphering Airflow schedules
```python

```

## Troubleshooting DAG runs
```python

```




# Maintaining and monitoring Airflow workflows
## Airflow sensors
```python

```

## Sensors vs operators
```python

```

## Sensory deprivation
```python

```

## Airflow executors
```python

```

## Determining the executor
```python

```

## Executor implications
```python

```

## Debugging and troubleshooting in Airflow
```python

```

## DAGs in the bag
```python

```

## Missing DAG
```python

```

## SLAs and reporting in Airflow
```python

```

## Defining an SLA
```python

```

## Defining a task SLA
```python

```

## Generate and email a report
```python

```

## Adding status emails
```python

```




# Building production pipelines in Airflow
## Working with templates
```python

```

## Creating a templated BashOperator
```python

```

## Templates with multiple arguments
```python

```

## More templates
```python

```

## Using lists with templates
```python

```

## Understanding parameter options
```python

```

## Sending templated emails
```python

```

## Branching
```python

```

## Define a BranchPythonOperator
```python

```

## Branch troubleshooting
```python

```

## Creating a production pipeline
```python

```

## Creating a production pipeline #1
```python

```

## Creating a production pipeline #2
```python

```

## Adding the final changes to your pipeline
```python

```

## Congratulations!
```python

```


