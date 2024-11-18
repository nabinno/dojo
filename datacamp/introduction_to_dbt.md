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

## dbt model statements
```

```

## Creating a dbt model
```

```

## Updating dbt models
```

```

## Config files
```

```

## Updating a dbt model
```

```

## Hierarchical models in dbt
```

```

## No hierarchy model creation
```

```

## Hierarchical model creation
```

```

## Updating model hierarchies
```

```

## Model troubleshooting
```

```

## Error classification
```

```

## Process of troubleshooting
```

```

## Troubleshooting model errors
```

```




# 3 Testing & Documentation
## Introduction to testing in dbt
```

```

## Built-in tests
```

```

## Defining tests on a model
```

```

## Finding bad data
```

```

## Creating singular tests
```

```

## Steps to develop a singular test
```

```

## Verifying trip duration
```

```

## Verifying test queries
```

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
