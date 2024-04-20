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
SELECT
    c1.name AS city,
    code,
    c2.name AS country,
    region,
    city_proper_pop
FROM cities AS c1
-- Perform an inner join with cities as c1 and countries as c2 on country code
INNER JOIN countries AS c2 ON c2.code = c1.country_code
ORDER BY code DESC;

SELECT
    c1.name AS city,
    code,
    c2.name AS country,
    region,
    city_proper_pop
FROM cities AS c1
-- Join right table (with alias)
LEFT JOIN countries AS c2
ON c1.country_code = c2.code
ORDER BY code DESC;
```

## Building on your LEFT JOIN
```sql
SELECT name, region, gdp_percapita
FROM countries AS c
LEFT JOIN economies AS e
-- Match on code fields
USING(code)
-- Filter for the year 2010
WHERE year = 2010;

-- Select region, and average gdp_percapita as avg_gdp
SELECT region, AVG(gdp_percapita) AS avg_gdp
FROM countries AS c
LEFT JOIN economies AS e
USING(code)
WHERE year = 2010
-- Group by region
GROUP BY region;

SELECT region, AVG(gdp_percapita) AS avg_gdp
FROM countries AS c
LEFT JOIN economies AS e
USING(code)
WHERE year = 2010
GROUP BY region
-- Order by descending avg_gdp
ORDER BY avg_gdp DESC
-- Return only first 10 records
LIMIT 10;
```

## Is this RIGHT?
```sql
-- Modify this query to use RIGHT JOIN instead of LEFT JOIN
SELECT countries.name AS country, languages.name AS language, percent
FROM languages
RIGHT JOIN countries
USING(code)
ORDER BY language;
```

## Comparing joins
```sql
SELECT name AS country, code, region, basic_unit
FROM countries
-- Join to currencies
FULL JOIN currencies
USING (code)
-- Where region is North America or name is null
WHERE region = 'North America' OR name IS NULL
ORDER BY region;

SELECT name AS country, code, region, basic_unit
FROM countries
-- Join to currencies
LEFT JOIN currencies
USING (code)
WHERE region = 'North America'
    OR name IS NULL
ORDER BY region;

SELECT name AS country, code, region, basic_unit
FROM countries
-- Join to currencies
INNER JOIN currencies
USING (code)
WHERE region = 'North America'
    OR name IS NULL
ORDER BY region;
```

## Chaining FULL JOINs
```sql
SELECT
    c1.name AS country,
    region,
    l.name AS language,
    basic_unit,
    frac_unit
FROM countries as c1
-- Full join with languages (alias as l)
FULL JOIN languages l USING(code)
-- Full join with currencies (alias as c2)
FULL JOIN currencies c2 USING(code)
WHERE region LIKE 'M%esia';
```

## Histories and languages
```sql
SELECT c.name AS country, l.name AS language
FROM languages l
-- Inner join countries as c with languages as l on code
INNER JOIN countries c USING(code)
WHERE c.code IN ('PAK','IND')
    AND l.code in ('PAK','IND');

SELECT c.name AS country, l.name AS language
FROM countries AS c
-- Perform a cross join to languages (alias as l)
CROSS JOIN languages AS l
WHERE c.code in ('PAK','IND')
    AND l.code in ('PAK','IND');```
```

## Choosing your join
```sql
SELECT
    c.name AS country,
    region,
    life_expectancy AS life_exp
FROM countries AS c
-- Join to populations (alias as p) using an appropriate join
JOIN populations AS p
    ON c.code = p.country_code
-- Filter for only results in the year 2010
WHERE p.year = 2010
-- Sort by life_exp
ORDER BY life_expectancy
-- Limit to five records
LIMIT 5;
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
