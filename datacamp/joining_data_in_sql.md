---
title: Joining Data in SQL
tags: database,structured-query-language,join
url: https://campus.datacamp.com/courses/joining-data-in-postgresql
---

# 1. Introduction to joins
## Introduction to INNER JOIN
```sql
##
-- Select all columns from cities
SELECT * FROM cities;

##
SELECT * 
FROM cities
  -- Inner join to countries
  INNER JOIN countries
    -- Match on the country codes
    ON cities.country_code = countries.code;

##
-- Select name fields (with alias) and region 
SELECT cities.name AS city, countries.name AS country, countries.region
FROM cities
  INNER JOIN countries
    ON cities.country_code = countries.code;
```

## Inner join
```sql
-- Select fields with aliases
SELECT c.code AS country_code, name, year, inflation_rate
FROM countries AS c
  -- Join to economies (alias e)
  INNER JOIN economies AS e
    -- Match on code
    ON e.code = c.code;
```

## Inner join (3)
```sql
##
-- Select fields
SELECT c.code, c.name, c.region, p.year, p.fertility_rate
  -- From countries (alias as c)
  FROM countries AS c
  -- Join with populations (as p)
  INNER JOIN populations AS p
    -- Match on country code
    ON c.code = p.country_code

##
-- Select fields
SELECT c.code, c.name, c.region, e.year, p.fertility_rate, e.unemployment_rate
  -- From countries (alias as c)
  FROM countries AS c
  -- Join to populations (as p)
  INNER JOIN populations AS p
    -- Match on country code
    ON c.code = p.country_code
  -- Join to economies (as e)
  INNER JOIN economies AS e
    -- Match on country code
    ON e.code = c.code;

##
-- Select fields
SELECT c.code, c.name, c.region, e.year, p.fertility_rate, e.unemployment_rate
  -- From countries (alias as c)
  FROM countries AS c
  -- Join to populations (as p)
  INNER JOIN populations AS p
    -- Match on country code
    ON c.code = p.country_code
  -- Join to economies (as e)
  INNER JOIN economies AS e
    -- Match on country code and year
    ON c.code = e.code AND p.year = e.year;
```

## Inner join with using
```sql
-- Select fields
SELECT c.name AS country, c.continent, l.name AS language, l.official
  -- From countries (alias as c)
  FROM countries AS c
  -- Join to languages (as l)
  INNER JOIN languages AS l
    -- Match using code
    USING(code)
```

## Self-ish joins, just in CASE
```sql

```

## Self-join
```sql

```

## Case when and then
```sql

```

## Inner challenge
```sql

```



# 2. Outer joins and cross joins
## LEFT and RIGHT JOINs
```sql

```

## Left Join
```sql

```

## Left join (2)
```sql

```

## Left join (3)
```sql

```

## Right join
```sql

```

## FULL JOINs
```sql

```

## Full join
```sql

```

## Full join (2)
```sql

```

## Full join (3)
```sql

```

## Review outer joins
```sql

```

## CROSSing the rubicon
```sql

```

## A table of two cities
```sql

```

## Outer challenge
```sql

```




# 3. Set theory clauses
## State of the UNION
```sql

```

## Union
```sql

```

## Union (2)
```sql

```

## Union all
```sql

```

## INTERSECTional data science
```sql

```

## Intersect
```sql

```

## Intersect (2)
```sql

```

## Review union and intersect
```sql

```

## EXCEPTional
```sql

```

## Except
```sql

```

## Except (2)
```sql

```

## Semi-joins and Anti-joins
```sql

```

## Semi-join
```sql

```

## Relating semi-join to a tweaked inner join
```sql

```

## Diagnosing problems using anti-join
```sql

```

## Set theory challenge
```sql

```




# 4. Subqueries
## Subqueries inside WHERE and SELECT clauses
```sql

```

## Subquery inside where
```sql

```

## Subquery inside where (2)
```sql

```

## Subquery inside select
```sql

```

## Subquery inside FROM clause
```sql

```

## Subquery inside from
```sql

```

## Advanced subquery
```sql

```

## Subquery challenge
```sql

```

## Subquery review
```sql

```

## Course review
```sql

```

## Final challenge
```sql

```

## Final challenge (2)
```sql

```

## Final challenge (3)
```sql

```

