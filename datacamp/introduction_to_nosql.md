---
title: Introduction to NoSQL
tags: nosql,structured-query-language
url: https://campus.datacamp.com/courses/introduction-to-nosql/introduction-to-nosql-databases
---

# 1. Introduction to NoSQL databases
## Querying NoSQL databases
```python
# Update the query to pull the team, year, event, and medal
# fields from the olympic_medals table
query = """
SELECT
team,
    year,
    event,
    medal
FROM olympic_medals;
"""
results = conn.cursor().execute(query).fetch_pandas_all()
print(results)
<script.py> output:
     TEAM   YEAR  EVENT                        MEDAL
  0  China  1992  Basketball Mens Basketball   None
  1  China  2012  Judo Mens Extra-Lightweight  None
  [..]
  [70000 rows x 4 columns]

# Select the review column from the nested_reviews table
query = """
SELECT
review
FROM nested_reviews;
"""
data = pd.read_sql(query, db_engine)
print(data)

# Set the name key-value pair
redis_conn.set("name", "Sarah")
# Retrieve and print the value stored in the "name" key
name = redis_conn.get("name")
print(name)
<script.py> output:
  review
  0  {'statement': 'If youve ever been to Disneyla...'
  1  {'statement': 'Its been a while since d last t...'
  [..]
  [10000 rows x 1 columns]
<script.py> output:
  Sarah
```

## NoSQL use-cases and applications
```
[x]Persisting and accessing large, analytical datasets containing census data
[ ]Storing student information with a  high-degree of structure and constraints
[x] Storing and querying loosely structured user-generated data whose schema changes frequently
[x]Capturing and delivering data in-memory, which a mobile ordering application
```

## Column-oriented databases
```
[x]Column-oriented databases allow for selective column read and retrieval, which reduces the amount of data that needs to be read from disk
[x]Its easier for column-oriented database to handle schema changes than traditional row-oriented databases
[ ]Column-oriented databases cannot be queried with SQL-like syntax, and require an entirely different query language to interact with
[x]Column-oriented database design allows for more efficient data storage, which enhances query performance
```

## Querying a column-oriented database
```
# Write a query to return all columns, limiting to 10 rows
query = "SELECT * FROM olympic_medals LIMIT 10;"

# Execute the query
results = conn.cursor().execute(query).fetch_pandas_all()

# Print the results of the query
print(results)
<script.py> output:
       ID                      NAME SEX  AGE  HEIGHT  ...  SEASON         CITY          SPORT                              EVENT  MEDAL
    0   1                 A Dijiang   M   24   180.0  ...  Summer    Barcelona     Basketball         Basketball Mens Basketball   None
    1   2                  A Lamusi   M   23   170.0  ...  Summer       London           Judo        Judo Mens Extra-Lightweight   None
    [..]
    [10 rows x 15 columns]
```

## Filtering a column-oriented database
```
# Return team, name, and year for all years greater than 2000
query = """
SELECT
team,
    name,
    year
FROM olympic_medals
WHERE year > 2000
;
"""
# Execute the query, print the results
results = conn.cursor().execute(query).fetch_pandas_all()
print(results)
```

## Identifying tabular and non-tabular NoSQL data stores
```
Tabular:
- Column-oriented databases

Non-tabular:
- Key-value databases
- Graph databases
- Document databases
```

## Querying document databases
```
# Update the query to select the review field
query = """
SELECT
review
    FROM nested_reviews;
"""

# Execute the query
data = pd.read_sql(query, db_engine)

# Print the first element of the DataFrame
print(data.iloc[0, 0])
```

## Querying nested documents with Postgres JSON
```
# Build the query to select the statement and location fields
query = """
SELECT
    review -> 'statement' AS statement,
    review -> 'location' AS location
FROM nested_reviews;
"""

# Execute the query, render results
data = pd.read_sql(query, db_engine)
print(data)
# <script.py> output:
# statement                                           location
#     0     If youve ever been to Disneyland anywhere you...  {'branch': 'Disneyland_HongKong', 'reviewer': ...
#     1     Its been a while since d last time we visit HK...  {'branch': 'Disneyland_HongKong', 'reviewer': ...
#     2     Thanks God it wasn   t too hot or too humid wh...  {'branch': 'Disneyland_HongKong', 'reviewer': ...
#     3     HK Disneyland is a great compact park. Unfortu...  {'branch': 'Disneyland_HongKong', 'reviewer': ...
#     4     the location is not in the city, took around 1...  {'branch': 'Disneyland_HongKong', 'reviewer': ...
#     [..]
#     [10000 rows x 2 columns]
```




# 2. Column-oriented Databases
## CREATE TABLE with Snowflake
```sql
CREATE TABLE olympic_athletes (
  age INT,
  country VARCHAR(64),
  is_first_games BOOLEAN
  last_name VARCHAR(64),
);
```

## Populating Snowflake tables
```
[ ]CREATE  TABLE ... AS: best-used when loading the very first data into a Snowflake data warehouse, where no other tables currently exist.
[ ]COPY INTO: typically used when writing the results of a Snowflake query to a new table.
[x]CREATE TABLE ... AS: Snowflakes syntax that allows for a table to be created from the results of another query, against existing tables or views.
[x]COPY INTO: used when loading data from a file to a Snowflake table.
```

## COPY INTO and CREATE TABLE ... AS with Snowflake
```
Correct:
- `COPY INTO Olympic_medals FROM "file://raw_olympic" FILE_FORMAT = (TYPE = 'CSV' FIELD_DELIMITER = ',');`
- `CREATE TABLE gold_medal_winners AS SELECT team, year, sport, event FROM olympic_medals WHERE medal = 'Gold';`

Incorrect:
- `CREATE TABLE olympic_medals_analysis USING FROM Olympic_medals FILE_FORMAT = (TYPE = 'CSV' FIELD_DELIMITER = ',');`
- `COPY INTO silver_meal_winners AS SELECT team, year, sport, event FROM olympic_medals WHERE medal = 'Silver';`
```

## Micro-partitioning and data clustering with Snowflake
```
[ ]Micro-partition: a small amount of data that a Data Analyst uses to source a dashboard.
[x]Micro-partition: a group of rows from a larger Snowflake table, stored in columnar format, containing metadata about the data it is storing.
[x]Data clustering: the grouping and sorting of data within a micro-partition that decreases the data read via query pruning.
[ ]Data clustering: Using the `GROUP BY` keywords to create a dataset later used for clustering.
```

## Query pruning with micro-partitioning
```
[ ]By only returning the first 100 records of the result set rather than all of the records.
[ ]By compressing the data specified by the query above as it is returned.
[x]By scanning the metadata for each partition and using that metadata to determine if the data in the micro-partition can be pruned.
```

## Leveraging micro-partitions and data clustering
```python
# Leverage the existing micro-partitions and data clustering
query = """
  SELECT
    team,
    year,
    sport,
    event,
    medal
  FROM olympic_medals
  WHERE year >= 2000;
  """

# Execute the query, print the results
results = conn.cursor().execute(query).fetch_pandas_all()
print(results)
```

## Common table expressions with Snowflake
```sql
WITH summer_metals AS (
  -- Return only summer results
  SELECT
    year,
    team,
    event,
    medal
  FROM olympic_medals
  WHWRE season = 'Summer'
)
SELECT * FROM summer_medals;
```

## Building an analytics workflow with Snowflake
```
[x]Common table expressions are used to temporarily store the result of a query in a named object, for later use in a query.
[x]When executed, a statement starting with CREATE OR REPLACE MATERIALIZED VIEW stores the results of a query as a table, to be queried later using the name provided.
[ ]Common table expressions group common tables together into a single schema in Snowflake.
[ ]Materialized and non-materialized VIEWs both store the results of a query in a table, to be later queried or referenced.
```

## Materialized vs. non-materialized VIEWs
```
Non-materialized VIEW:
- Does not store data in a table when defined, rather, stores a "named definition" or a query.
- Defined using the `CREATE [OR REPLACE] VIEW` keywords.
- Help more with query organization, rather than performance.

Materialized VIEW:
- Stores the results of a query in a table upon definition of the `VIEW`, at the cost of data recency.
- Defined using the `CREATE [OR REPLACE] MATERIALIZED VIEW` keywords.
- Help with both query performance, as well as modularity and maintainability.
```

## Semi-structured Snowflake data types
```
[x]ARRAY
[x]OBJECT
[ ]JSON
[x]VARIANT
```

## Querying semi-structured data in Snowflake
```python
# Build a query to pull city and country names
query = """
  SELECT
    city_meta:city,
    city_meta:country
  FROM host_cities;
  """

# Execute query and output results
results = conn.cursor().execute(query).fetch_pandas_all()
print(results)
```

## Querying nested semi-structured data
```

```





# 3. Document Databases
## Understanding JSON data in Postgres
```

```

## JSON and JSONB data in Postgres
```

```

## Executing queries with sqlalchemy and pandas
```

```

## Querying JSON and JSONB data from Postgres
```

```

## Storing JSON data in Postgres
```

```

## Loading Postgres with semi-structured data
```

```

## Converting tabular data to JSON
```

```

## Extracting keys from JSON objects with Postgres
```

```

## Querying JSON data using Postgres
```

```

## Querying top-level JSON data
```

```

## Finding the type of JSON data
```

```

## Working with nested JSON objects
```

```

## Filtering document databases with Postgres JSON
```

```

## Advanced Postgres JSON query techniques
```

```

## Arrow and hash arrow operators
```

```

## #> and #>>
```

```

## Extracting document data
```

```

## Manipulating document data
```

```





# 4. Key-value and Graph Databases
## Introduction to key-value databases
```

```

## Key-value databases
```

```

## Key-value vs. document databases
```

```

## Connecting to a Redis cluster
```

```

## Storing and retrieving key-value data
```

```

## Storing key-value data with Redis
```

```

## Retrieving key-value data with Redis
```

```

## Storing Python dictionaries with Redis
```

```

## Graph databases
```

```

## Understanding graph databases
```

```

## Graph database providers
```

```

## Using graph databases
```

```

## Wrapping up!
```

```
