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

## Self-join
```sql
##
-- Select fields with aliases
SELECT p1.country_code,
p1.size AS size2010,
p2.size AS size2015
-- From populations (alias as p1)
FROM populations AS p1
  -- Join to itself (alias as p2)
  INNER JOIN populations as p2
    -- Match on country code
    USING(country_code)

##
-- Select fields with aliases
SELECT p1.country_code,
       p1.size AS size2010,
       p2.size AS size2015
-- From populations (alias as p1)
FROM populations as p1
  -- Join to itself (alias as p2)
  INNER JOIN populations as p2
    -- Match on country code
    ON p1.country_code = p2.country_code
        -- and year (with calculation)
        AND p1.year = p2.year - 5;

##
-- Select fields with aliases
SELECT p1.country_code,
       p1.size AS size2010, 
       p2.size AS size2015,
       -- Calculate growth_perc
       ((p2.size - p1.size)/p1.size * 100.0) AS growth_perc
-- From populations (alias as p1)
FROM populations AS p1
  -- Join to itself (alias as p2)
  INNER JOIN populations AS p2
    -- Match on country code
    ON p1.country_code = p2.country_code
        -- and year (with calculation)
        AND p1.year = p2.year - 5;
```

## Case when and then
```sql
SELECT name, continent, code, surface_area,
    -- First case
    CASE WHEN surface_area > 2000000 THEN 'large'
        -- Second case
        WHEN surface_area > 350000 THEN 'medium'
        -- Else clause + end
        ELSE 'small' END
        -- Alias name
        AS geosize_group
-- From table
FROM countries;
```

## Inner challenge
```sql
##
SELECT country_code, size,
    -- First case
    CASE WHEN size > 50000000 THEN 'large'
        -- Second case
        WHEN size > 1000000 THEN 'medium'
        -- Else clause + end
        ELSE 'small' END
        -- Alias name (popsize_group)
        AS popsize_group
-- From table
FROM populations
-- Focus on 2015
WHERE year = 2015;

##
SELECT country_code, size,
    CASE WHEN size > 50000000 THEN 'large'
        WHEN size > 1000000 THEN 'medium'
        ELSE 'small' END
        AS popsize_group
-- Into table
INTO pop_plus
FROM populations
WHERE year = 2015;

-- Select all columns of pop_plus
SELECT * FROM pop_plus;

##
SELECT country_code, size,
  CASE WHEN size > 50000000
            THEN 'large'
       WHEN size > 1000000
            THEN 'medium'
       ELSE 'small' END
       AS popsize_group
INTO pop_plus       
FROM populations
WHERE year = 2015;

-- Select fields
SELECT c.name, c.continent, c.geosize_group, p.popsize_group
-- From countries_plus (alias as c)
FROM countries_plus AS c
  -- Join to pop_plus (alias as p)
  INNER JOIN pop_plus AS p
    -- Match on country code
    ON c.code = p.country_code
-- Order the table    
ORDER BY c.geosize_group;
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

