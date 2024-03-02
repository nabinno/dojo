---
title: Intermediate SQL
tags: database,structured-query-language
url: https://campus.datacamp.com/courses/intermediate-sql
---

# 1. Selecting Data
## Learning to COUNT()
```
[ ]The number of unique films in the reviews table.
[x]The number of records containing a film_id.
[ ]The total number of records in the reviews table.
[ ]The sum of the film_id field.
```

## Practice with COUNT()
```
-- Count the number of records in the people table
SELECT COUNT(*) AS count_records
FROM people;

-- Count the number of birthdates in the people table
SELECT COUNT(birthdate) AS count_birthdate
FROM people;

-- Count the records for languages and countries represented in the films table
SELECT
  COUNT(country) AS count_countries,
  COUNT(language) AS count_languages
FROM films;
```

## SELECT DISTINCT
```
-- Return the unique countries from the films table
SELECT DISTINCT country
FROM films;

-- Count the distinct countries from the films table
SELECT COUNT(DISTINCT country) AS count_distinct_countries
FROM films;
```

## Order of execution
```
FROM
SELECT
LIMIT
```

## Debugging errors
```
-- Debug this code
SELECT certification
FROM films
LIMIT 5;

-- Debug this code
SELECT film_id, imdb_score, num_votes
FROM reviews;

-- Debug this code
SELECT COUNT(birthdate) AS count_birthdays
FROM people;
```

## SQL style
```

```

## SQL best practices
```

```

## Formatting
```

```

## Non-standard fields
```

```





# 2. Filtering Records
## Filtering numbers
```

```

## Filtering results
```

```

## Using WHERE with numbers
```

```

## Using WHERE with text
```

```

## Multiple criteria
```

```

## Using AND
```

```

## Using OR
```

```

## Using BETWEEN
```

```

## Filtering text
```

```

## LIKE and NOT LIKE
```

```

## WHERE IN
```

```

## Combining filtering and selecting
```

```

## NULL values
```

```

## What does NULL mean?
```

```

## Practice with NULLs
```

```




# 3.Aggregate Functions
## Summarizing data
```

```

## Aggregate functions and data types
```

```

## Practice with aggregate functions
```

```

## Summarizing subsets
```

```

## Combining aggregate functions with WHERE
```

```

## Using ROUND()
```

```

## ROUND() with a negative parameter
```

```

## Aliasing and arithmetic
```

```

## Using arithmetic
```

```

## Aliasing with functions
```

```

## Rounding results
```

```





# 4. Sorting and Grouping
## Sorting results
```

```

## Sorting text
```

```

## Sorting single fields
```

```

## Sorting multiple fields
```

```

## Grouping data
```

```

## GROUP BY single fields
```

```

## GROUP BY multiple fields
```

```

## Answering business questions
```

```

## Filtering grouped data
```

```

## Filter with HAVING
```

```

## HAVING and sorting
```

```

## All together now
```

```

## Congratulations!
```

```
