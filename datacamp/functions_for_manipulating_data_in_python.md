---
title: Functions for Manipulating Data in Python
tags: python,data-engineering
url: https://campus.datacamp.com/courses/functions-for-manipulating-data-in-postgresql/overview-of-common-data-types
---

# 1 Overview of Common Data Types
## Text data types
```
Which of the following is not a valid text data type in PostgreSQL?
[ ]TEXT
[x]STRING
[ ]CHAR
[ ]VARCHAR
```

## Getting information about your database
```sql
-- Select all columns from the TABLES system database
SELECT *
FROM INFORMATION_SCHEMA.TABLES
-- Filter by schema
WHERE table_schema = 'public';

-- Select all columns from the COLUMNS system database
SELECT *
FROM INFORMATION_SCHEMA.COLUMNS
WHERE table_name = 'actor';
```

## Determining data types
```sql
-- Get the column name and data type
SELECT
    column_name,
    data_type
-- From the system database information schema
FROM INFORMATION_SCHEMA.COLUMNS
-- For the customer table
WHERE table_name = 'customer';
```

## Properties of date and time data types
```
Which of the following is NOT correct?

[ ]TIMESTAMP data types contain both date and time values.
[ ]DATE data types use an yyyy-mm-dd format.
[ ]INTERVAL types are representations of periods of time.
[x]TIME data types are stored with a timezone by default.
```

## Interval data types
```sql
SELECT
    -- Select the rental and return dates
    rental_date,
    return_date,
    -- Calculate the expected_return_date
    rental_date + INTERVAL '3 days' AS expected_return_date
FROM rental;
```

## Accessing data in an ARRAY
```sql
-- 1) Select the title and special features column
SELECT
  title,
  special_features
FROM film;

-- 2) Select the title and special features column
SELECT
  title,
  special_features
FROM film
-- Use the array index of the special_features column
WHERE special_features[1] = 'Trailers';

-- 3) Select the title and special features column
SELECT
  title,
  special_features
FROM film
-- Use the array index of the special_features column
WHERE special_features[2] = 'Deleted Scenes';
```

## Searching an ARRAY with ANY
```sql
SELECT
  title,
  special_features
FROM film
-- Modify the query to use the ANY function
WHERE 'Trailers' = ANY (special_features);
```

## Searching an ARRAY with @>
```sql
SELECT
  title,
  special_features
FROM film
-- Filter where special_features contains 'Deleted Scenes'
WHERE special_features @> ARRAY['Deleted Scenes'];
```




# 2 Working with DATE/TIME Functions and Operators
## Overview of basic arithmetic operators
```

```

## Adding and subtracting date and time values
```

```

## INTERVAL arithmetic
```

```

## Calculating the expected return date
```

```

## Functions for retrieving current date/time
```

```

## Current timestamp functions
```

```

## Working with the current date and time
```

```

## Manipulating the current date and time
```

```

## Extracting and transforming date/ time data
```

```

## Using EXTRACT
```

```

## Using DATE_TRUNC
```

```

## Putting it all together
```

```




# 3 Parsing and Manipulating Text
## Reformatting string and character data
```

```

## Concatenating strings
```

```

## Changing the case of string data
```

```

## Replacing string data
```

```

## Parsing string and character data
```

```

## Determining the length of strings
```

```

## Truncating strings
```

```

## Extracting substrings from text data
```

```

## Combining functions for string manipulation
```

```

## Truncating and padding string data
```

```

## Padding
```

```

## The TRIM function
```

```

## Putting it all together
```

```




# 4 Full-text Search and PostgresSQL Extensions
## Introduction to full-text search
```

```

## A review of the LIKE operator
```

```

## What is a tsvector?
```

```

## Basic full-text search
```

```

## Extending PostgreSQL
```

```

## User-defined data types
```

```

## Getting info about user-defined data types
```

```

## User-defined functions in Sakila
```

```

## Intro to PostgreSQL extensions
```

```

## Enabling extensions
```

```

## Measuring similarity between two strings
```

```

## Levenshtein distance examples
```

```

## Putting it all together
```

```

## Wrap Up
```

```
