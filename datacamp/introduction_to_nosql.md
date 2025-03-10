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
```python
# Build a query to extract nested location coordinates
query = """
SELECT
  city_meta:coordinates.lat,
  city_meta:coordinates.long
FROM host_cities;
"""

# Execute the query and output the results
results = conn.cursor().execute(query).fetch_pandas_all()
print(results)
```





# 3. Document Databases
## Executing queries with sqlalchemy and pandas
```python
import pandas as pd
import sqlalchemy

# Create a connection to the reviews database
db_engine = sqlalchemy.create_engine("postgresql+psycopg2://repl:password@localhost:5432/disneyland")

# Execute a query against the nested_reviews table
results = pd.read_sql("SELECT * FROM nested_reviews;", db_engine)
print(results)
```

## Querying JSON and JSONB data from Postgres
```python
import pandas as pd
import sqlalchemy

# Create a connection to the reviews database
db_engine = sqlalchemy.create_engine("postgresql+psycopg2://repl:password@localhost:5432/disneyland")

query = """SELECT * FROM nested_reviews;"""

# Execute the query, check out the results
results = pd.read_sql(query, db_engine)

# Print the review column from the results DataFrame
print(results["review"])
```

## Loading Postgres with semi-structured data
```
[ ]INSERT INTO and COPY ... FROM can be used interchangeably to transfer data from a semi-structured database to Postgres.
[ ]INSERT INTO loads the records from a file to a table, while COPY ... FROM inserts an individual record, or set of records, into a table in Postgres.
[ ]COPY ... FROM can only duplicate data from an existing Postgres table, while INSERT INTO can pull data from any semi-structured data store into Postgres.
[x]INSERT INTO is used to add rows to a Postgres table, while COPY ... FROM allows for all records in a file to populate a table.
```

## Converting tabular data to JSON
```python
# Build a query to create a JSON-object
query = """
SELECT
  row_to_json(row(review_id, rating, year_month))
FROM reviews;
"""

# Execute the query, and output the results
results = pd.read_sql(query, db_engine)
print(results.head(10))
```

## Extracting keys from JSON objects with Postgres
```python
# Build a query to find the unique keys in the review column
query = """
SELECT
  DISTINCT json_object_keys(review)
FROM nested_reviews;
"""

# Execute the query, show the results
unique_keys = pd.read_sql(query, db_engine)
print(unique_keys)
```

## Querying top-level JSON data
```python
# Build the query to select the review_id and rating fields
query = """
SELECT
  review -> 'location' AS location,
  review -> 'statement' AS statement
FROM nested_reviews;
"""

# Execute the query, render results
data = pd.read_sql(query, db_engine)
print(data)
```

## Finding the type of JSON data
```python
# Find the data type of the location field
query = """
SELECT
  json_typeof(review -> 'location') AS location_type
FROM nested_reviews;
"""

# Execute the query, render results
data = pd.read_sql(query, db_engine)
print(data)
```

## Working with nested JSON objects
```python
# Update the query to select the nested reviewer field
query = """
SELECT
    review -> 'location' ->> 'branch' AS branch,
    review -> 'location' ->> 'reviewer' AS reviewer
FROM nested_reviews;
"""

# Execute the query, render results
data = pd.read_sql(query, db_engine)
print(data)
```

## Filtering document databases with Postgres JSON
```python
# Build the query to select the rid and rating fields
query = """
SELECT
    review -> 'statement' AS customer_review
FROM nested_reviews
WHERE review -> 'location' ->> 'branch' = 'Disneyland_California';
"""

# Execute the query, render results
data = pd.read_sql(query, db_engine)
print(data)
```

## Arrow and hash arrow operators
```txt
[x]The arrow operators can be chained together to query nested data, but the resulting statement can become long.
[x]Using the hash arrow operators, a column is provided along with a string array, specifying the field to be queried.
[ ]Neither the arrow or hash arrow operators can be used to query nested document data.
[ ]The arrow operators can take a string array of fields, while the hash arrow operators can only take a single field, as a string.
```

## #> and #>>
```python
# Attempt to query the statement, nested branch, and nested
# zipcode fields from the review column
query = """
SELECT
    json_typeof(review #> '{statement}'),
    review #>> '{location,branch}' AS branch,
    review #>> '{location,zipcode}' AS zipcode
FROM nested_reviews;
"""

# Execute the query, render results
data = pd.read_sql(query, db_engine)
print(data)
```

## Extracting document data
```python
# Return the statement and reviewer fields, filter by the
# nested branch field
query = """
    SELECT
        json_extract_path(review, 'statement'),
        json_extract_path_text(review, 'location', 'reviewer')
    FROM nested_reviews
    WHERE json_extract_path_text(review, 'location', 'branch') = 'Disneyland_California';
"""

data = pd.read_sql(query, db_engine)
print(data)
```

## Manipulating document data
```python
# Extract fields from JSON, and filter by reviewer location
query = """
    SELECT
        review_id,
        review #> '{location, branch}' AS branch,
        review ->> 'statement' AS statement,
        rating
    FROM nested_reviews
    WHERE json_extract_path_text(review, 'location', 'reviewer') = 'Australia'
    ORDER BY rating DESC;
"""

data = pd.read_sql(query, db_engine)
print(data)
```





# 4. Key-value and Graph Databases
## Key-value databases
```
[x]Key-value database can only be queried by key, not by value.
[x]Redis is a popular tool to use when configuring a key-value database.
[ ]Key-value databases can be used to store large amounts of relational data to source dashboard.
[x]Its common to use key-value databases for caching or session management, especially in web applications.
```

## Key-value vs. document databases
```
Key-value databases:
- Commonly used when for caching or session management in web applications.
- Typically stores data in-memory, rather than on disk.
- Values can only be retrieved by querying keys, leading to fast read and write operations.

Document database:
- Data is stored on disk, rather than being stored in memory.
- Values stored within documents can be queried directly.
```

## Connecting to a Redis cluster
```python
import redis

# Create a connection to Redis cluster
r = redis.Redis(
host="localhost",
    port=6379,
    decode_responses=True
)
```

## Storing key-value data with Redis
```python
# Store the city key-value pair
redis_conn.set("city", "London")

# Store the sunshine key-value pair
redis_conn.set("sunshine", "7")

# Retrieve values stored at the city and sunshine keys
city = redis_conn.get("city")
sunshine = redis_conn.get("sunshine")

print(city)
print(sunshine)
```

## Retrieving key-value data with Redis
```python
# Loop through each of the cities
for city in cities:
    # Grab the temperature
    temperature = redis_conn.get(f"{city}_temp")

    # Check if the temperature is None
    if temperature is None:
        # Store an unknown temperature
        redis_conn.set(f"{city}_temp", "unknown temperature")
        print(f"Unknown temperature in {city}")
    else:
        # Otherwise, print the temperature
        print(f"The temperature in {city} is {temperature}")
```

## Storing Python dictionaries with Redis
```python
# Create a dictionary containing weather data
london_weather_mapping = {
    "temperature": 42,
    "humidity": 88,
    "visibility": "low"
}

# Store the london_weather key-value pair
redis_conn.hset(
    "london_weather",
    mapping=london_weather_mapping
)

# Retrieve and print the london_weather key-value pair
print(redis_conn.hgetall("london_weather"))
```

## Understanding graph databases
```
[ ]Databases used to represent KPIs such as item sales or inventory in tools such as Tableau.
[x]NoSQL data stores that persist data in a network of nodes and edges, where each node represents an entity, and each edge represents a relationship between those entities.
[ ]Data stores that are created and maintained by data visualization engineers to store metadata about graphs built for executive dashboards.
[ ]A tool used in graph theory to store information about all the different ways to traverse a graph.
```

## Graph database providers
```
[x]Neo4j
[x]Amazon (AWS) Neptune
[ ]Snowflake
[x]ArangoDB
[ ]Postgres JSON
```

## Using graph databases
```
Appropriate use-case for graph database:
- Identifying fraudulent employee activity
- Creating a data model for a dating app
- Building a data storage architecture for a homegrown search engine

Inappropriate use-case for graph abs:
- Working with large, analytics datasets
- Storing simple, key-value data in a web application
```
