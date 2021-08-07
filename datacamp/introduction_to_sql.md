---
title: Introduction to SQL
tags: database,structured-query-language
url: https://campus.datacamp.com/courses/introduction-to-sql
---

# 1. Selecting columns
## Onboarding | Query Result
```sql
SELECT name FROM people;
```

## Onboarding | Errors
```sql
-- Try running me!
SELECT 'DataCamp <3 SQL' AS result;
```

## Onboarding | Multi-step Exercises
```sql
SELECT 'SQL is cool'
AS result;
```

## Beginning your SQL journey
```sql
##
SELECT title
FROM tilms;

##
SELECT release_year
FROM films;

##
SELECT name
FROM people;
```

## SELECTing multiple columns
```sql
##
SELECT * FROM films;

##
SELECT title, release_year
FROM films;

##
SELECT title, release_year, country
FROM films;
```

## SELECT DISTINCT
```sql
##
SELECT DISTINCT country FROM films;

##
SELECT DISTINCT certification FROM films;

##
SELECT DISTINCT role FROM roles;
```

## Learning to COUNT
```sql
SELECT COUNT(*) FROM reviews;
-- 4968
```

## Practice with COUNT
```sql
##
SELECT COUNT(*) FROM people;

##
SELECT COUNT(birthdate)
FROM people;

##
SELECT COUNT(DISTINCT birthdate)
FROM people;

##
SELECT COUNT(DISTINCT language) FROM films;

##
SELECT COUNT(DISTINCT country) FROM films;
```




# 2. Filtering rows
## Simple filtering of numeric values
```sql
##
SELECT *
FROM films
WHERE release_year = 2016;

##
SELECT COUNT(*)
FROM films
WHERE release_year < 2000;

##
SELECT title, release_year
FROM films
WHERE release_year > 2000;
```

## Simple filtering of text
```sql
##
SELECT *
FROM films
WHERE language = 'French';

##
SELECT name, birthdate
FROM people
WHERE birthdate = '1974-11-11';

##
SELECT COUNT(*)
FROM films
WHERE language = 'Hindi';

##
SELECT *
FROM films
WHERE certification = 'R';
```

## WHERE AND
```sql
##
SELECT title, release_year
FROM films
WHERE language = 'Spanish' AND release_year < 2000;

##
SELECT *
FROM films
WHERE language = 'Spanish' AND release_year > 2000;

##
SELECT *
FROM films
WHERE language = 'Spanish'
  AND release_year > 2000
  AND release_year < 2010;
```

## WHERE AND OR (2)
```sql
##
SELECT title, release_year
FROM films
WHERE release_year >= 1990 AND release_year < 2000;

##
SELECT title, release_year
FROM films
WHERE (release_year >= 1990 AND release_year < 2000)
AND (language = 'French' OR language = 'Spanish');

##
SELECT title, release_year
FROM films
WHERE (release_year >= 1990 AND release_year < 2000)
AND (language = 'French' OR language = 'Spanish')
AND (gross >= 2000000);
```

## BETWEEN (2)
```sql
SELECT title, release_year
FROM films
WHERE release_year BETWEEN 1990 AND 2000
AND budget > 100000000
AND (language = 'Spanish' OR language = 'French');
```

## WHERE IN
```sql
##
SELECT title, release_year
FROM films
WHERE release_year IN (1990, 2000)
AND duration > 120;

##
SELECT title, language
FROM films
WHERE language IN ('English', 'Spanish', 'French');

##
SELECT title, certification
FROM films
WHERE certification IN ('NC-17', 'R');
```

## NULL and IS NULL
```sql
##
SELECT name
FROM people
WHERE deathdate IS NULL;

##
SELECT title
FROM films
WHERE budget IS NULL;

##
SELECT COUNT(*)
FROM films
WHERE language IS NULL;
```

## LIKE and NOT LIKE
```sql
##
SELECT name
FROM people
WHERE name LIKE 'B%';

##
SELECT name
FROM people
WHERE name LIKE '_r%';

##
SELECT name
FROM people
WHERE name NOT LIKE 'A%';
```




# 3. Aggregate Functions
## Aggregate functions
```sql

```

## Aggregate functions practice
```sql

```

## Combining aggregate functions with WHERE
```sql

```

## A note on arithmetic
```sql

```

## It's AS simple AS aliasing
```sql

```

## Even more aliasing
```sql

```




# 4. Sorting and grouping
## ORDER BY
```sql

```

## Sorting single columns
```sql

```

## Sorting single columns (2)
```sql

```

## Sorting single columns (DESC)
```sql

```

## Sorting multiple columns
```sql

```

## GROUP BY
```sql

```

## GROUP BY practice
```sql

```

## GROUP BY practice (2)
```sql

```

## HAVING a great time
```sql

```

## All together now
```sql

```

## All together now (2)
```sql

```

## A taste of things to come
```sql

```



