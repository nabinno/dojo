---
title: Introduction to Airflow in Python
tags: python, apache-airflow
url: https://campus.datacamp.com/courses/object-oriented-programming-in-pythonhttps://campus.datacamp.com/courses/introduction-to-airflow-in-python/
---

# 1. Intro to Airflow
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
$ airflow list_dags
```

## Troubleshooting DAG creation
```python
## refresh_data_workflow.py
from airflow.models import DAG
default_args = {
  'owner': 'jdoe',
  'email': 'jdoe@datacamp.com'
}
dag = DAG( 'refresh_data', default_args=default_args )
```

## Starting the Airflow webserver
```python
$ airflow webserver -p 9000
```




# 2. Implementing Airflow DAGs
## Defining a BashOperator task
```python
# Import the BashOperator
from airflow.operators.bash_operator import BashOperator

# Define the BashOperator 
cleanup = BashOperator(
    task_id='cleanup_task',
    # Define the bash_command
    bash_command='cleanup.sh',
    # Add the task to the dag
    dag=analytics_dag
)
```

## Multiple BashOperators
```python
# Define a second operator to run the `consolidate_data.sh` script
consolidate = BashOperator(
    task_id='consolidate_task',
    bash_command='consolidate_data.sh',
    dag=analytics_dag)

# Define a final operator to execute the `push_data.sh` script
push_data = BashOperator(
    task_id='pushdata_task',
    bash_command='push_data.sh',
    dag=analytics_dag)
```

## Define order of BashOperators
```python
# Define a new pull_sales task
pull_sales = BashOperator(
    task_id='pullsales_task',
    bash_command='wget https://salestracking/latestinfo?json',
    dag=analytics_dag
)

# Set pull_sales to run prior to cleanup
pull_sales >> cleanup

# Configure consolidate to run after cleanup
cleanup >> consolidate

# Set push_data to run last
consolidate >> push_data
```

## Using the PythonOperator
```python
def pull_file(URL, savepath):
    r = requests.get(URL)
    with open(savepath, 'wb') as f:
        f.write(r.content)   
    # Use the print method for logging
    print(f"File pulled from {URL} and saved to {savepath}")

from airflow.operators.python_operator import PythonOperator

# Create the task
pull_file_task = PythonOperator(
    task_id='pull_file',
    # Add the callable
    python_callable=pull_file,
    # Define the arguments
    op_kwargs={'URL':'http://dataserver/sales.json', 'savepath':'latestsales.json'},
    dag=process_sales_dag
)
```

## More PythonOperators
```python
# Add another Python task
parse_file_task = PythonOperator(
    task_id='parse_file',
    # Set the function to call
    python_callable=parse_file,
    # Add the arguments
    op_kwargs={'inputfile':'latestsales.json', 'outputfile':'parsedfile.json'},
    # Add the DAG
    dag=process_sales_dag
)
```

## EmailOperator and dependencies
```python
# Import the Operator
from airflow.operators.email_operator import EmailOperator

# Define the task
email_manager_task = EmailOperator(
    task_id='email_manager',
    to='manager@datacamp.com',
    subject='Latest sales JSON',
    html_content='Attached is the latest sales JSON file as requested.',
    files='parsedfile.json',
    dag=process_sales_dag
)

# Set the order of tasks
pull_file_task >> parse_file_task >> email_manager_task
```

## Schedule a DAG via Python
```python
# Update the scheduling arguments as defined
default_args = {
  'owner': 'Engineering',
  'start_date': datetime(2019, 11, 1),
  'email': ['airflowresults@datacamp.com'],
  'email_on_failure': False,
  'email_on_retry': False,
  'retries': 3,
  'retry_delay': timedelta(minutes=20)
}

dag = DAG('update_dataflows', default_args=default_args, schedule_interval='30 12 * * 3')
```

## Deciphering Airflow schedules
```python
* * * * *
timedelta(minutes=5)
@hourly
* 0,12 * * *
timedelta(days=1)
@weekly
```

## Troubleshooting DAG runs
```python
import requests
import json
from datetime import datetime
from airflow.models import DAG
from airflow.operators.python_operator import PythonOperator
from airflow.operators.email_operator import EmailOperator


default_args = {
    'owner':'sales_eng',
    'start_date': datetime(2020, 2, 15),
}

process_sales_dag = DAG(dag_id='process_sales', default_args=default_args, schedule_interval='@monthly')


def pull_file(URL, savepath):
    r = requests.get(URL)
    with open(savepath, 'w') as f:
        f.write(r.content)
    print(f"File pulled from {URL} and saved to {savepath}")
    

pull_file_task = PythonOperator(
    task_id='pull_file',
    # Add the callable
    python_callable=pull_file,
    # Define the arguments
    op_kwargs={'URL':'http://dataserver/sales.json', 'savepath':'latestsales.json'},
    dag=process_sales_dag
)

def parse_file(inputfile, outputfile):
    with open(inputfile) as infile:
      data=json.load(infile)
      with open(outputfile, 'w') as outfile:
        json.dump(data, outfile)
        
parse_file_task = PythonOperator(
    task_id='parse_file',
    # Set the function to call
    python_callable=parse_file,
    # Add the arguments
    op_kwargs={'inputfile':'latestsales.json', 'outputfile':'parsedfile.json'},
    # Add the DAG
    dag=process_sales_dag
)

email_manager_task = EmailOperator(
    task_id='email_manager',
    to='manager@datacamp.com',
    subject='Latest sales JSON',
    html_content='Attached is the latest sales JSON file as requested.',
    files='parsedfile.json',
    dag=process_sales_dag
)

pull_file_task >> parse_file_task >> email_manager_task
```




# 3. Maintaining and monitoring Airflow workflows
## Sensors vs operators
```python
# Sensors
Derives from `BaseSensorOperator`.
Has a `poke_interval` attribute.
`FileSensor`

# Both
Are assigned to DAGs.
Have a `task_id`.

# Operators
`BashOperator`
Only runs once per DAG run.
```

## Sensory deprivation
```python
from airflow.models import DAG
from airflow.operators.bash_operator import BashOperator
from airflow.operators.python_operator import PythonOperator
from airflow.operators.http_operator import SimpleHttpOperator
from airflow.contrib.sensors.file_sensor import FileSensor

dag = DAG(
   dag_id = 'update_state',
   default_args={"start_date": "2019-10-01"}
)

precheck = FileSensor(
   task_id='check_for_datafile',
   filepath='salesdata_ready.csv',
   dag=dag)

part1 = BashOperator(
   task_id='generate_random_number',
   bash_command='echo $RANDOM',
   dag=dag
)

import sys
def python_version():
   return sys.version

part2 = PythonOperator(
   task_id='get_python_version',
   python_callable=python_version,
   dag=dag)
   
part3 = SimpleHttpOperator(
   task_id='query_server_for_external_ip',
   endpoint='https://api.ipify.org',
   method='GET',
   dag=dag)
   
precheck >> part3 >> part2
```

## Executor implications
```python
from airflow.models import DAG
from airflow.operators.bash_operator import BashOperator
from airflow.contrib.sensors.file_sensor import FileSensor
from datetime import datetime

report_dag = DAG(
    dag_id = 'execute_report',
    schedule_interval = "0 0 * * *"
)

precheck = FileSensor(
    task_id='check_for_datafile',
    filepath='salesdata_ready.csv',
    start_date=datetime(2020,2,20),
    mode='reschedule',
    dag=report_dag
)

generate_report_task = BashOperator(
    task_id='generate_report',
    bash_command='generate_report.sh',
    start_date=datetime(2020,2,20),
    dag=report_dag
)

precheck >> generate_report_task
```

## Debugging and troubleshooting in Airflow
```python

```

## DAGs in the bag
```python
from airflow.models import DAG
from airflow.operators.bash_operator import BashOperator
from airflow.operators.python_operator import PythonOperator
from airflow.operators.http_operator import SimpleHttpOperator
from airflow.contrib.sensors.file_sensor import FileSensor

dag = DAG(
   dag_id = 'update_state',
   default_args={"start_date": "2019-10-01"}
)

precheck = FileSensor(
   task_id='check_for_datafile',
   filepath='salesdata_ready.csv',
   dag=dag)

part1 = BashOperator(
   task_id='generate_random_number',
   bash_command='echo $RANDOM',
   dag=dag
)

import sys
def python_version():
   return sys.version

part2 = PythonOperator(
   task_id='get_python_version',
   python_callable=python_version,
   dag=dag)

part3 = SimpleHttpOperator(
   task_id='query_server_for_external_ip',
   endpoint='https://api.ipify.org',
   method='GET',
   dag=dag)

precheck >> part3 >> part2
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




# 4. Building production pipelines in Airflow
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

