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

## Testing, testing, testing
```
Reusable:
- accepted_values
- `tests/generic/check_columns_equal.sql`
- `select * from {{ ref('order') }}` where fulfilled = 1`

Singular:
- `tests/assert_orderdate_aft`
- A test file with `{% enttest %}` as the last line.
```

## Implementing a reusable test
```
$ cat models/taxi_rides/model_properties.yml
version: 2

# Replace any entry of ____ with the appropriate content
models:
- name: taxi_rides_raw
  columns:
    - name: fare_amount
      tests:
        - not_null
    - name: total_amount
      tests:
        - check_gt_0

$ cat tests/generic/check_gt_0.sql
{% test check_gt_0(model, column_name) %}
select *
from {{ model }}
where {{ column_name }} <= 0
{% endtest %}

$ dbt clear
$ dbt run
$ dbt test --select taxi_rides_raw.sql
```

## Updating from singular to reusable test
```
$ cat tests/generic/columns_equal.sql
{% test columns_equal(model, column_name, column_name2) %}
select *
from {{ model }}
where {{ column_name }} = {{ column_name2 }}
{% endtest %}

$ cat models/taxi_rides/model_properties.yml
version: 2

models:
- name: taxi_rides_raw
  columns:
    - name: tpep_pickup_datetime
      tests:
        - columns_equal:
            column_name2: tpep_dropoff_datetime

$ dbt run
$ dbt test --select taxi_ride_raw.sql
```

## dbt docs Command Options
```
[ ]dbt docs generate
[ ]dbt docs -h
[ ]dbt docs serve
[x]dbt docs create
```

## dbt documentation flow
```
Add documentation to models, test, etc
Use `dbt run` to generate models
Use `dbt docs generate` to generate the documentation files
Copy content to hosting service
Access documentation via web browser
```

## Creating dbt documentation
```
$ cat mdels/model_properties.yml
version: 2

models:
- name: taxi_rides_raw
  description: Initial import of the NYC Yellow Taxi trip data from Parquet source
  access: public

$ dbt docs generate
$ dbt docs serve
```




# 4 Implementing dbt in production
## Orderly YML
```
sources:
- name: raw
  tables:
    - name: phone_orders
      columns:
        - name: id
          tests:
            - unique
```

## Models, sources, or both?
```
Only dbt models:
- Uses a `models:` section
- Uses the `{{ ref() }}` function

Common to models and sources:
- Could be defined in `models/model_properties.yml`
- Data lineage is applied
- Can run tests

Only dbt sources:
- Uses the `sources:` key
- Called via `{{ source(source_name, table_name) }}`
```

## Adding a source
```
$ cat datacheck
#!/usr/bin/env python3
import duckdb
con = duckdb.connect('dbt.duckdb', read_only=True)
print(con.sql('select * from taxi_rides_raw'))
print(con.sql('select count(*) from taxi_rides_raw'))
if (con.execute('select count(*) from taxi_rides_raw').fetchall()[0][0] == 300000):
  with open('/home/repl/workspace/successful_data_check', 'w') as f:
    f.write('300000')

$ cat models/model_properties.yml
version: 2

sources:
- name: raw
  tables:
    - name: taxi_rides

$ cat models/taxi_rides/taxi_rides_raw.sql
{{ config(materialized='view')}}

select * from {{ source('raw', 'taxi_rides') }}

$ dbt run
$ ./datacheck
[..]
┌──────────────┐
│ count_star() │
│    int64     │
├──────────────┤
│       300000 │
└──────────────┘
```

## Kernels of truth
```
[x]Seeds are (mostly) CSV files
[x]A user can manually define column datatypes on a seed
[ ]Seeds are accessed via the `source()` function
[ ]Seeds should contain regularly updated information
```

## ZIP is the code
```
$ mv nynj_zipcodes.csv seeds/

$ cat seeds/properties.yml
version: 2

seeds:
  - name: nynj_zipcodes.csv
    config:
      column_types:
        zipcode: varchar(5)

$ dbt seed

$ cat datacheck
#!/usr/bin/env python3
import duckdb
import os.path

try:
  con = duckdb.connect('dbt.duckdb', read_only=True)
  print(con.sql('select * from nynj_zipcodes'))
  if (con.execute('select count(*) as total_zipcodes from nynj_zipcodes').fetchall()[0][0] == 2877):
    print("Your data counts are correct!")
    with open('/home/repl/workspace/successful_data_check', 'w') as f:
      f.write('2877')
  else:
    print("There appears to be an issue with your data counts.")
except duckdb.CatalogException:
  if not os.path.isfile('/home/repl/workspace/nyc_yellow_taxi/seeds/nynj_zipcodes.csv'):
    print("It looks like the nynj_zipcodes.csv file is\n not present in the seeds/ directory.\n\nPlease move the file and try again.\n")
  else:
    print("It looks like your data warehouse doesn't exist yet\n\n Please run the dbt seed command and try again.\n")

$ ./datacheck
┌─────────┬─────────────────┬────────────┐
│ zipcode │      place      │   state    │
│  int32  │     varchar     │  varchar   │
├─────────┼─────────────────┼────────────┤
│    8037 │ Hammonton       │ New Jersey │
│    8201 │ Absecon         │ New Jersey │
│    8203 │ Brigantine      │ New Jersey │
│    8205 │ Absecon         │ New Jersey │
│    8213 │ Cologne         │ New Jersey │
│    8215 │ Egg Harbor City │ New Jersey │
│    8217 │ Elwood          │ New Jersey │
│    8220 │ Leeds Point     │ New Jersey │
│    8221 │ Linwood         │ New Jersey │
│    8225 │ Northfield      │ New Jersey │
│      ·  │     ·           │    ·       │
│      ·  │     ·           │    ·       │
│      ·  │     ·           │    ·       │
│   14418 │ Branchport      │ New York   │
│   14441 │ Dresden         │ New York   │
│   14478 │ Keuka Park      │ New York   │
│   14507 │ Middlesex       │ New York   │
│   14527 │ Penn Yan        │ New York   │
│   14544 │ Rushville       │ New York   │
│   14837 │ Dundee          │ New York   │
│   14842 │ Himrod          │ New York   │
│   14857 │ Lakemont        │ New York   │
│   14878 │ Rock Stream     │ New York   │
├─────────┴─────────────────┴────────────┤
│ 2877 rows (20 shown)         3 columns │
└────────────────────────────────────────┘
```

## Snapshot process
```
1. Create a `.sql` file named for the snapshot
2. Add a `{% snapshot %}` directive
3. Select the `unique_key` column
4. Specify the `updated_at` column
5. Write the query
6. Run `dbt snapshot` on a regular basis
```

## Snapshot issue
```
[ ]Tehy should split the update into smaller chunks.
[ ]Your colleague did not provide a `refresh:` option to the `snapshot` YAML.
[x]The snapshot should be run hourly, after the update process completes.
```

## Adding a snapshot
```
$ cat snapshots/vehicle_list_snapshot.sql
{% snapshot vehicle_list_snapshot %}

{{
  config(
    target_schema='main',
    unique_key='license_id',
    strategy='timestamp',
    updated_at='last_updated'
  )
}}

select * from {{ source('raw', 'vehicle_list') }}
{% endsnapshot %}

$ dbt snapshot
```

## What can't dbt build do?
```
dbt build executes this step:
- dbt seed
- dbt snapshot
- dbt run
- dbt test

dbt build does *not* run this step:
- dbt docs
- dbt compile
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
