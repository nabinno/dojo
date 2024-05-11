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

## Comparing a country to itself
```sql
-- Select aliased fields from populations as p1
SELECT p1.country_code, p1.size AS size2010, p2.size AS size2015
FROM populations AS p1
-- Join populations as p1 to itself, alias as p2, on country code
INNER JOIN populations AS p2 ON p1.country_code = p2.country_code;

SELECT
    p1.country_code,
    p1.size AS size2010,
    p2.size AS size2015
FROM populations AS p1
INNER JOIN populations AS p2
    ON p1.country_code = p2.country_code
WHERE p1.year = 2010
-- Filter such that p1.year is always five years before p2.year
    AND p2.year = 2015;
```

## All joins on deck
```sql
INNER JOIN:
You sell houses and have two tables, `listing_prices` and `price_solid`. You want a table with sale prices and listing prices, only if you know both.

LEFT JOIN:
You run a pizza delivery service with loyal clients. You want a table of clients and their weekly orders, with nulls if here are not orders.

FULL JOIN:
You want a report of whether your patients have reached out to you, or you have reached out to them. You are fine with nulls for either condition.
```




# 3. Set Theory for SQL Joins
## UNION vs. UNION ALL
```sql
SELECT code
FROM languages
UNION
SELECT curr_id
FROM currencies;
```

## Comparing global economies
```sql
 -- Select all fields from economies2015
 SELECT * FROM economies2015
 -- Set operation
 UNION
 -- Select all fields from economies2019
 SELECT * FROM economies2019
 ORDER BY code, year;
```

## Comparing two set operations
```sql
-- Query that determines all pairs of code and year from economies and populations, without duplicates
SELECT code, year
FROM economies
UNION
SELECT country_code, year
FROM populations;

SELECT code, year
FROM economies
-- Set theory clause
UNION ALL
SELECT country_code, year
FROM populations
ORDER BY code, year;
```

## INTERSECT
```sql
-- Return all cities with the same name as a country
SELECT cities.name
FROM cities
INTERSECT
SELECT countries.name
FROM countries;
```

## You've got it, EXCEPT...
```sql
-- Return all cities that do not have the same name as a country
SELECT name FROM cities
EXCEPT
SELECT name FROM countries
ORDER BY name;
```

## Calling all set operators
```sql
UNION or UNION ALL:
- You are a school teacher teaching multiple classes. You would like to comibe the grades of all students into are consolidated sheet.

INTERSECT:
- A residence hall has asked students to rank their preferences to be assigned a room. Tehy now want to pair students based on common preferences.

EXCEPT:
- You run a music streaming service and have a list of songs a user has listened to. You want to show them new songs they haven't heard before.
```



# 4. Subqueries
## Multiple WHERE clauses
```sql
SELECT *
FROM economiies2019
WHERE code in (
    SELECT code
    FROM economies2015
    WHERE gross_savings < 22.5
);
```

## Semi join
```sql
-- Select country code for countries in the Middle East
SELECT code
FROM countries
WHERE region = 'Middle East';

-- Select unique language names
SELECT DISTINCT name
FROM languages
-- Order by the name of the language
ORDER BY name ASC;

SELECT DISTINCT name
FROM languages
-- Add syntax to use bracketed subquery below as a filter
WHERE code IN
    (SELECT code
     FROM countries
     WHERE region = 'Middle East')
ORDER BY name;
```

## Diagnosing problems using anti join
```sql
-- Select code and name of countries from Oceania
SELECT code, name
FROM countries
WHERE continent = 'Oceania';

SELECT code, name
FROM countries
WHERE continent = 'Oceania'
  -- Filter for countries not included in the bracketed subquery
  AND code NOT IN
      (SELECT code
       FROM currencies);
```

## Subquery inside WHERE
```sql
SELECT *
FROM populations
-- Filter for only those populations where life expectancy is 1.15 times higher than average
WHERE life_expectancy > 1.15 *
  (SELECT AVG(life_expectancy)
   FROM populations
   WHERE year = 2015)
  AND year = 2015;
```

## WHERE do people live?
```sql
-- Select relevant fields from cities table
SELECT name, country_code, urbanarea_pop
FROM cities
-- Filter using a subquery on the countries table
WHERE name IN (
    SELECT capital
    FROM countries
)
ORDER BY urbanarea_pop DESC;
```

## Subquery inside SELECT
```sql
-- Find top nine countries with the most cities
SELECT countries.name AS country, COUNT(*) AS cities_num
FROM countries
LEFT JOIN cities
ON countries.code = cities.country_code
GROUP BY country
-- Order by count of cities as cities_num
ORDER BY cities_num DESC, country
LIMIT 9;

SELECT countries.name AS country,
-- Subquery that provides the count of cities
  (SELECT COUNT(*)
   FROM cities
   WHERE cities.country_code = countries.code) AS cities_num
FROM countries
ORDER BY cities_num DESC, country
LIMIT 9;
```

## Subqueries inside FROM
```sql
-- Select code, and language count as lang_num
SELECT code, COUNT(*) lang_num
FROM languages
GROUP BY code;

-- Select local_name and lang_num from appropriate tables
SELECT local_name, lang_num
FROM countries,
  (SELECT code, COUNT(*) AS lang_num
   FROM languages
   GROUP BY code) AS sub
-- Where codes match
WHERE countries.code = sub.code
ORDER BY lang_num DESC;
```

## Subquery challenge
```sql
-- Select relevant fields
SELECT code, inflation_rate, unemployment_rate
FROM economies
WHERE year = 2015
  AND code NOT IN
-- Subquery returning country codes filtered on gov_form
  (SELECT code
   FROM countries
   WHERE gov_form LIKE '%Republic%'
     OR gov_form LIKE '%Monarchy%')
ORDER BY inflation_rate;
```

## Final challenge
```sql
-- Select fields from cities
SELECT
    name,
    country_code,
    city_proper_pop,
    metroarea_pop,
    city_proper_pop / metroarea_pop * 100 AS city_perc
FROM cities
-- Use subquery to filter city name
WHERE name IN
    (SELECT capital
     FROM countries
     WHERE (continent = 'Europe'
         OR continent LIKE '%America'))
-- Add filter condition such that metroarea_pop does not have null values
    AND metroarea_pop IS NOT NULL
-- Sort and limit the result
ORDER BY city_perc DESC
LIMIT 10;
```
