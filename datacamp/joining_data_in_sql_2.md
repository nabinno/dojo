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

```

## USING in action
```sql

```

## Defining relationships
```sql

```

## Relationships in our database
```sql

```

## Inspecting a relationship
```sql

```

## Multiple joins
```sql

```

## Joining multiple tables
```sql

```

## Checking multi-table joins
```sql

```





## # 2. Outer Joins, Cross Joins and Self Joins
```sql

```

## LEFT and RIGHT JOINs
```sql

```

## Remembering what is LEFT
```sql

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
