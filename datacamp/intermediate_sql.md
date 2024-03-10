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

## SQL best practices
```
Best Practice:
- Captialize keywords
- Use underscores in field names rather than spaces
- End queries with a semicolon

Poor Practice:
- Write the query on one line
- Don't capitalize keywords
- Write lots of queries with no semicolon
```

## Formatting
```
-- Rewrite this query
SELECT person_id, role
FROM roles
LIMIT 10;
```

## Non-standard fields
```
SELECT film_id, "facebook likes"
FROM reviews;
```



# 2. Filtering Records
## Filtering results
```
SELECT title
FROM films
WHERE release_year > 2000;

[ ]Films released before the year 2000
[x]Films released after the year 2000
[ ]Films released after the year 2001
[ ]Films released in 2000
```

## Using WHERE with numbers
```
-- Select film_ids and imdb_score with an imdb_score over 7.0
SELECT film_id, imdb_score
FROM reviews
WHERE imdb_score > 7.0;

-- Select film_ids and facebook_likes for ten records with less than 1000 likes
SELECT film_id, facebook_likes
FROM reviews
WHERE facebook_likes < 1000
LIMIT 10;

-- Count the records with at least 100,000 votes
SELECT COUNT(num_votes) AS films_over_100K_votes
FROM reviews
WHERE num_votes > 100000;
```

## Using WHERE with text
```
-- Count the Spanish-language films
SELECT COUNT(language) AS count_spanish
FROM films
WHERE language = 'Spanish';
```

## Using AND
```
-- Select the title and release_year for all German-language films released before 2000
SELECT title, release_year
FROM films
WHERE language = 'German'
    AND release_year < 2000;

-- Update the query to see all German-language films released after 2000
SELECT title, release_year
FROM films
WHERE release_year > 2000
    AND language = 'German';

-- Select all records for German-language films released after 2000 and before 2010
SELECT *
FROM films
WHERE language = 'German'
    AND release_year > 2000
    AND release_year < 2010;
```

## Using OR
```
-- Find the title and year of films from the 1990 or 1999
SELECT title, release_year
FROM films
WHERE release_year = 1990
    OR release_year = 1999;

SELECT title, release_year
FROM films
WHERE (release_year = 1990 OR release_year = 1999)
-- Add a filter to see only English or Spanish-language films
    AND language IN ('English', 'Spanish');

SELECT title, release_year
FROM films
WHERE (release_year = 1990 OR release_year = 1999)
    AND (language = 'English' OR language = 'Spanish')
    -- Filter films with more than $2,000,000 gross
    AND gross > 2000000;
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
