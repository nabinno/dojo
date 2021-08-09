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
## Left Join
```sql
##
-- Select the city name (with alias), the country code,
-- the country name (with alias), the region,
-- and the city proper population
SELECT c1.name AS city, code, c2.name AS country,
       region, city_proper_pop
-- From left table (with alias)
FROM cities AS c1
  -- Join to right table (with alias)
  INNER JOIN countries AS c2
    -- Match on country code
    ON c1.country_code = c2.code
-- Order by descending country code
ORDER BY code DESC;

##
SELECT c1.name AS city, code, c2.name AS country,
       region, city_proper_pop
FROM cities AS c1
  -- Join right table (with alias)
  LEFT OUTER JOIN countries AS c2
    -- Match on country code
    ON c1.country_code = c2.code
-- Order by descending country code
ORDER BY c2.code DESC;
```

## Left join (2)
```sql
##
/*
Select country name AS country, the country's local name,
the language name AS language, and
the percent of the language spoken in the country
*/
SELECT c.name AS country, local_name, l.name AS language, percent
-- From left table (alias as c)
FROM countries AS c
  -- Join to right table (alias as l)
  INNER JOIN languages AS l
    -- Match on fields
    ON c.code = l.code
-- Order by descending country
ORDER BY country DESC;

##
/*
Select country name AS country, the country's local name,
the language name AS language, and
the percent of the language spoken in the country
*/
SELECT c.name AS country, local_name, l.name AS language, percent
-- From left table (alias as c)
FROM countries AS c
  -- Join to right table (alias as l)
  LEFT OUTER JOIN languages AS l
    -- Match on fields
    ON c.code = l.code
-- Order by descending country
ORDER BY country DESC;
```

## Left join (3)
```sql
##
-- Select name, region, and gdp_percapita
SELECT c.name, c.region, e.gdp_percapita
-- From countries (alias as c)
FROM countries AS c
  -- Left join with economies (alias as e)
  LEFT OUTER JOIN economies AS e
    -- Match on code fields
    ON c.code = e.code
-- Focus on 2010
WHERE e.year = 2010;

##
-- Select fields
SELECT c.region, AVG(e.gdp_percapita) AS avg_gdp
-- From countries (alias as c)
FROM countries AS c
  -- Left join with economies (alias as e)
  LEFT OUTER JOIN economies AS e
    -- Match on code fields
    ON e.code = c.code
-- Focus on 2010
WHERE e.year = 2010
-- Group by region
GROUP BY c.region;

##
-- Select fields
SELECT region, AVG(gdp_percapita) AS avg_gdp
-- From countries (alias as c)
FROM countries AS c
  -- Left join with economies (alias as e)
  LEFT OUTER JOIN economies AS e
    -- Match on code fields
    ON e.code = c.code
-- Focus on 2010
WHERE e.year = 2010
-- Group by region
GROUP BY c.region
-- Order by descending avg_gdp
ORDER BY avg_gdp DESC;
```

## Right join
```sql
-- convert this code to use RIGHT JOINs instead of LEFT JOINs
/*
SELECT cities.name AS city, urbanarea_pop, countries.name AS country,
       indep_year, languages.name AS language, percent
FROM cities
  LEFT JOIN countries
    ON cities.country_code = countries.code
  LEFT JOIN languages
    ON countries.code = languages.code
ORDER BY city, language;
*/

SELECT cities.name AS city, urbanarea_pop, countries.name AS country,
       indep_year, languages.name AS language, percent
FROM languages
  RIGHT JOIN countries
    ON languages.code = countries.code
  RIGHT JOIN cities
    ON countries.code = cities.country_code
ORDER BY city, language;
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

