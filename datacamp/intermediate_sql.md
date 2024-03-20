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
-- Select the title and release_year for films released between 1990 and 2000
SELECT title, release_year
FROM films
WHERE release_year BETWEEN 1990 AND 2000;

SELECT title, release_year
FROM films
WHERE release_year BETWEEN 1990 AND 2000
-- Narrow down your query to films with budgets > $100 million
    AND budget > 100000000;

SELECT title, release_year
FROM films
WHERE release_year BETWEEN 1990 AND 2000
    AND budget > 100000000
    -- Restrict the query to only Spanish-language films
    AND language = 'Spanish';

SELECT title, release_year
FROM films
WHERE release_year BETWEEN 1990 AND 2000
    AND budget > 100000000
    -- Amend the query to include Spanish or French-language films
    AND language IN ('Spanish', 'French');
```

## LIKE and NOT LIKE
```
 -- Select the names that start with B
 SELECT name
 FROM people
 WHERE name LIKE 'B%';

 SELECT name
 FROM people
 -- Select the names that have r as the second letter
 WHERE name LIKE '_r%';

 SELECT name
 FROM people
 -- Select names that don't start with A
 WHERE name NOT LIKE 'A%'
```

## WHERE IN
```
-- Find the title and release_year for all films over two hours in length released in 1990 and 2000
SELECT title, release_year
FROM films
WHERE release_year IN (1990, 2000)
    AND duration > 120;

-- Find the title and language of all films in English, Spanish, and French
SELECT title, language
FROM films
WHERE language IN ('English', 'Spanish', 'French');

-- Find the title, certification, and language all films certified NC-17 or R that are in English, Italian, or Greek
SELECT title, certification, language
FROM films
WHERE certification IN ('NC-17', 'R')
    AND language IN ('English', 'Italian', 'Greek');
```

## Combining filtering and selecting
```
-- Count the unique titles
SELECT COUNT(DISTINCT title) AS nineties_english_films_for_teens
FROM films
-- Filter to release_years to between 1990 and 1999
WHERE release_year BETWEEN 1990 AND 1999
-- Filter to English-language films
    AND language = 'English'
-- Narrow it down to G, PG, and PG-13 certifications
    AND certification IN ('G', 'PG', 'PG-13');
```

## Practice with NULLs
```
-- List all film titles with missing budgets
SELECT title AS no_budget_info
FROM films
WHERE budget IS NULL;

-- Count the number of films we have language data for
SELECT COUNT(*) AS count_language_known
FROM films
WHERE language IS NOT NULL;
```



# 3.Aggregate Functions
## Aggregate functions and data types
```
Numerical data only:
- SUM()
- AVG()

Various data types:
- COUNT()
- MIN()
- MAX()
```

## Practice with aggregate functions
```
-- Query the sum of film durations
SELECT SUM(duration) total_duration
FROM films;

-- Calculate the average duration of all films
SELECT AVG(duration) average_duration
FROM films;

-- Find the latest release_year
SELECT MAX(release_year) latest_year
FROM films;

-- Find the duration of the shortest film
SELECT MIN(duration) shortest_film
FROM films;
```

## Combining aggregate functions with WHERE
```
-- Calculate the sum of gross from the year 2000 or later
SELECT SUM(gross) total_gross
FROM films
WHERE release_year >= 2000;

-- Calculate the average gross of films that start with A
SELECT AVG(gross) avg_gross_A
FROM films
WHERE title LIKE 'A%';

-- Calculate the lowest gross film in 1994
SELECT MIN(gross) lowest_gross
FROM films
WHERE release_year = 1994;

-- Calculate the highest gross film released between 2000-2012
SELECT MAX(gross) highest_gross
FROM films
WHERE release_year BETWEEN 2000 AND 2012;
```

## Using ROUND()
```
-- Round the average number of facebook_likes to one decimal place
SELECT ROUND(AVG(facebook_likes),1) avg_facebook_likes
FROM reviews;
```

## Using arithmetic
```
SELECT CAST(discount AS FLOAT) / paid_price AS discount_percentage
FROM ticket_sales;
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
