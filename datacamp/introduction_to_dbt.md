---
title: Introduction to dbt
tags: dbt, data-engineering
url: https://campus.datacamp.com/courses/introduction-to-dbt/welcome-to-dbt
---

# 1 Welcome to dbt
## Users of dbt
```
Likely to use dbt:
- Data Engineer
- Analytics Engineer
- Data Analyst

Unlikely to use dbt:
- Data Scientist
- CIO
- ML Engineer
```

## Version of dbt
```
dbt
```

## dbt subcommands
```
Which of the following is not a valid dbt subcommand?
[ ]docs
[ ]init
[x]reeset
[ ]run
[ ]build
```

## Initializing a dbt project
```
Run the following commands in the terminal window:
$ dbt init

Note: You will need to specify the project name (nyc_yellow_taxi) and select the database type ([1] duckdb)

Once complete run:
$ cd nyc_yellow_taxi
$ ls
```

## Creating a project profile
```
$ cat profiles.yml
nyc_yellow_taxi:
  outputs:
    dev:
      type: duckdb
      path: dbt.duckdb
    target: dev

$ dbt debug
```

## dbt project workflow
```
1. `dbt init`
2. Create data destinations in the `profiles.yml` file
3. Define the models
4. `dbt run`
5. Verify contents in the data warehouse
```

## Running a project
```
$ dbt run

$ cat models/taxi_rides/taxi_rides_raw.sql
-- Modify the following line to change the materialization type
with source_data as (
    -- Add the query as described to generate the data model
    select * from read_parquet('yellow_tripdata_2023-01-partial.parquet')
)

select * from source_data

$ ./datacheck
```



# 2 dbt models
## Features of a data model
```
True of a data model:
- Classifies common attributes about data
- Facilitates collaboration
- Is conceptual

False of a data model:
- Is the same as a data structure.
- Are always written in SQL
```

## Creating a dbt model
```
$ cd nyc_yellow_taxi
$ touch models/taxi_rides/taxxi_rides_raw.sql
$ cat $_
select * from 'yellow_tripdata_2023-01-partial.parquet'
$ dbt run
```

## Config files
```
Project:
- profiles.yml
- dbt_project.yml

Model only:
- test_model.ql
- models/test_model.yml
- model_properties.yml
```

## Hierarchical models in dbt
```
$ dbt run

$ cat models/taxi_rides/total_creditcard_riders_by_day.sql
-- Update with SQL to return requested information
select
   date_part('day', tpep_pickup_datetime) as day,
   count(*) as total_rides
from taxi_rides_raw
where Payment_type = 1
group by day

$ tail -9 dbt_project.yml
# In this example config, we tell dbt to build all models in the example/
# directory as views. These settings can be overridden in the individual model
# files using the `{{ config(...) }}` macro.
models:
  nyc_yellow_taxi:
  # Config indicated by + and applies to all files under models/example/
    taxi_rides:
    +materialized: view

$ ./datacheck
```

## No hierarchy model creation
```
1. all_weekend_sales
2. new_products_purchased_this_week
3. sales_raw
4. total_sales_by_rep
```

## Hierarchical model creation
```
1. sales_raw
2. all_weekend_sales
3. new_products_purchased_this_week
4. total_sales_by_rep
```

## Updating model hierarchies
```
$ dbt run -f
[..]
23:02:58  Runtime Error in model creditcard_riders_by_day (models/taxi_rides/creditcard_riders_by_day.sql)
23:02:58    Catalog Error: Table with name taxi_rides_raw does not exist!
23:02:58    Did you mean "temp.information_schema.tables"?
23:02:58
23:02:58  Done. PASS=1 WARN=0 ERROR=1 SKIP=0 TOTAL=2

$ cat models/taxi_rides/creditcard_riders_by_day.sql
-- Update SQL to use Jinja reference
select
   date_part('day', tpep_pickup_datetime) as day,
   count(*) as total_riders
from {{ref('taxi_rides_raw')}}
where payment_type = 1
group by day

$ dbt run -f
[..]
23:08:41  Completed successfully
23:08:41
23:08:41  Done. PASS=2 WARN=0 ERROR=0 SKIP=0 TOTAL=2
```

## Error classification
```
Query syntax error:
SELECT * WHERE
start_date > 2023-05-01;
SELECT first_name
last_name date_of_birth
FROM users;

Query logic error:
-- Generate the Average transaction
-- per user purchase
SELECT user_id, SUM(transaction_total)
FROM users
GROUP BY user_id;
SELECT * FROM users
WHERE last_login > 2023-05-01 and last_login < 2023-04-30;

Invalid reference:
The message:
Compilation Error Model 'model.taxi.fare' (models/taxi/fare.sql) depends on a node named 'taxi_rw' which was not found
```

## Process of troubleshooting
```
1. `dbt run`
2. Review `logs/dbt.log` or `run_results.json`
3. View generated SQL
4. Running generated SQL manually
5. Verify fix
```

## Troubleshooting model errors
```
$ dbt run
23:29:47  Completed with 1 error and 0 warnings:
23:29:47
23:29:47  Runtime Error in model taxi_rides_raw (models/taxi_rides/taxi_rides_raw.sql)
23:29:47    Catalog Error: Table with name source_data does not exist!
23:29:47    Did you mean "temp.information_schema.schemata"?
23:29:47
23:29:47  Done. PASS=0 WARN=0 ERROR=1 SKIP=0 TOTAL=1

$ dbt run
23:32:02  Completed with 1 error and 0 warnings:
23:32:02
23:32:02  Runtime Error in model taxi_rides_raw (models/taxi_rides/taxi_rides_raw.sql)
23:32:02    IO Error: No files found that match the pattern "yellow_tripdata_2023-01-partial.parquet"
23:32:02
23:32:02  Done. PASS=0 WARN=0 ERROR=1 SKIP=0 TOTAL=1

$ dbt run -f
```




# 3 Testing & Documentation
## Built-in tests
```
Built-in:
- `not_null`
- `unique`
- `relationships`
- `accepted_values`

Not built-in:
- `is_upper`
- `less_than`
- `not_in`
```

## Defining tests on a model
```
$ cat models/taxi_rides/model_properties.yml
version: 2

models:
- name: taxi_rides_raw
  columns:
    - name: fare_amount
      tests:
        - not_null
    - name: payment_type
      tests:
        - not_null
        - accepted_values:
            values: [1, 2, 3, 4, 5, 6]

$ dbt run

$ dbt test
```

## Finding bad data
```
$ dbt run
$ dbt test
[..]
23:19:36  Completed with 1 error and 0 warnings:
23:19:36
23:19:36  Failure in test accepted_values_taxi_rides_raw_payment_type__1__2__3__4__5__6 (models/taxi_rides/model_properties.yml)
23:19:36    Got 1 result, configured to fail if != 0
23:19:36
23:19:36    compiled Code at target/compiled/taxi_project/models/taxi_rides/model_properties.yml/accepted_values_taxi_rides_raw_payment_type__1__2__3__4__5__6.sql
23:19:36
23:19:36  Done. PASS=2 WARN=0 ERROR=1 SKIP=0 TOTAL=3

$ cat taxi_rides_raw.sql
{{ config(materialized='view')}}

with source_data as (
select * from read_parquet('yellow_tripdata_2023-01-partial.parquet') where payment_type != 0
)

-- Add a filter / WHERE clause to the line below
select * from source_data

$ dbt run
$ dbt test
[..]
23:20:19  Completed successfully
23:20:19
23:20:19  Done. PASS=3 WARN=0 ERROR=0 SKIP=0 TOTAL=3
```

## Steps to develop a singular test
```
1. Determine the validation required
2. Create the SQL query
3. Save the query in a test `.sql` file
4. `dbt run`
5. `dbt test`
```

## Verifying trip duration
```
$ cat tests/assert_trip_duration_gt_0.sql
select *
from taxi_rides_raw
-- Complete the test on the following line
where tpep_pickup_datetime = tpep_dropoff_datetime

$ dbt test --select assert_trip_duration_gt_0.sql
```

## Verifying test queries
```
Valid test query:
- SELECT * FROM taxi_rides_raw WHERE tpep_pickup_datetime = tpep_dropoff_datetime
- SELECT * FROM taxi_rides_raw WHERE tpep_dropoff_datetime - tpep_pickup_datetime = 0

Invalid test query:
- SELECT * FROM taxi_rides_raw WHERE tpep_dropoff_datetime > tpep_pickup_datetime
- SELECT * FROM taxi_rides_raw WHERE tpep_dropoff_datetime - tpep_pickup_datetime >= 0
```

## Creating custom reusable tests
```

```

## Testing, testing, testing
```

```

## Implementing a reusable test
```

```

## Updating from singular to reusable test
```

```

## Creating and generating dbt documentation
```

```

## dbt docs Command Options
```

```

## dbt documentation flow
```

```

## Creating dbt documentation
```

```




# 4 Implementing dbt in production
## dbt sources
```

```

## Orderly YML
```

```

## Models, sources, or both?
```

```

## Adding a source
```

```

## dbt seeds
```

```

## Kernels of truth
```

```

## ZIP is the code
```

```

## SCD2 with dbt snapshots
```

```

## Snapshot process
```

```

## Snapshot issue
```

```

## Adding a snapshot
```

```

## Automating with dbt build
```

```

## What can't dbt build do?
```

```

## Helping the intern!
```

```

## Putting it all together
```

```

## Course review
```

```
