---
title: Joining Data in SQL
tags: database,structured-query-language,join
url: https://campus.datacamp.com/courses/joining-data-in-sql
---

## Your first join
```sql
-- Select all columns from cities
SELECT *
FROM cities;

SELECT *
FROM cities
-- Inner join to countries
INNER JOIN countries
-- Match on country codes
ON countries.code = cities.country_code;

-- Select name fields (with alias) and region
SELECT cities.name AS city, countries.name AS country, countries.region
FROM cities
INNER JOIN countries
ON cities.country_code = countries.code;
```

## Joining with aliased tables
```sql
-- Select fields with aliases
SELECT c.code AS country_code, c.name, e.year, e.inflation_rate
FROM countries AS c
-- Join to economies (alias e)
INNER JOIN economies AS e
-- Match on code field using table aliases
ON c.code = e.code;
```

## USING in action
```sql
SELECT c.name AS country, l.name AS language, official
FROM countries AS c
INNER JOIN languages AS l
-- Match using the code column
USING(code);
```

## Relationships in our database
```sql
SELECT cities.name, countries.code
FROM countries
INNER JOIN cities
ON cities.country_code = countries.code;

SELECT languages.name, countries.code
FROM countries
INNER JOIN languages
ON languages.code = countries.code;
```

## Inspecting a relationship
```sql
-- Select country and language names, aliased
SELECT c.name country, l.name language
-- From countries (aliased)
FROM countries c
-- Join to languages (aliased)
INNER JOIN languages l
-- Use code as the joining field with the USING keyword
USING(code);

-- Rearrange SELECT statement, keeping aliases
SELECT c.name AS country, l.name AS language
FROM countries AS c
INNER JOIN languages AS l
USING(code)
-- Order the results by language
ORDER BY language;

-- Rearrange SELECT statement, keeping aliases
SELECT c.name AS country, l.name AS language
FROM countries AS c
INNER JOIN languages AS l
USING(code)
-- Order the results by language
ORDER BY country;
```

## Joining multiple tables
```sql
-- Select fields
SELECT name, year, fertility_rate
FROM countries AS c
INNER JOIN populations AS p
ON c.code = p.country_code;

-- Select fields
SELECT name, e.year, fertility_rate, unemployment_rate
FROM countries AS c
INNER JOIN populations AS p
ON c.code = p.country_code
-- Join to economies (as e)
INNER JOIN economies AS e
-- Match on country code
ON c.code = e.code;
```



## # 2. Outer Joins, Cross Joins and Self Joins
```sql
SELECT name, e.year, fertility_rate, unemployment_rate
FROM countries AS c
INNER JOIN populations AS p
ON c.code = p.country_code
INNER JOIN economies AS e
ON c.code = e.code
-- Add an additional joining condition such that you are also joining on year
    AND e.year = p.year;
```

## Remembering what is LEFT
```sql
SELECT c.name AS country, local_name, l.name AS language, percent
FROM countries AS c
LEFT JOIN languages AS l
USING(code)
ORDER BY country DESC;
```

## This is a LEFT JOIN, right?
```sql

```

## Building on your LEFT JOIN
```sql

```

## Is this RIGHT?
```sql

```

## FULL JOINs
```sql

```

## Comparing joins
```sql

```

## Chaining FULL JOINs
```sql

```

## Crossing into CROSS JOIN
```sql

```

## Histories and languages
```sql

```

## Choosing your join
```sql

```

## Self joins
```sql

```

## Comparing a country to itself
```sql

```

## All joins on deck
```sql

```




## # 3. Set Theory for SQL Joins
```sql

```

## Set theory for SQL Joins
```sql

```

## UNION vs. UNION ALL
```sql

```

## Comparing global economies
```sql

```

## Comparing two set operations
```sql

```

## At the INTERSECT
```sql

```

## INTERSECT
```sql

```

## Review UNION and INTERSECT
```sql

```

## EXCEPT
```sql

```

## You've got it, EXCEPT...
```sql

```

## Calling all set operators
```sql

```





## # 4. Subqueries
```sql

```

## Subquerying with semi joins and anti joins
```sql

```

## Multiple WHERE clauses
```sql

```

## Semi join
```sql

```

## Diagnosing problems using anti join
```sql

```

## Subqueries inside WHERE and SELECT
```sql

```

## Subquery inside WHERE
```sql

```

## WHERE do people live?
```sql

```

## Subquery inside SELECT
```sql

```

## Subqueries inside FROM
```sql

```

## Subquery inside FROM
```sql

```

## Subquery challenge
```sql

```

## Final challenge
```sql

```

## The finish line
```sql

```
