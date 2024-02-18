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
##
SELECT SUM(duration)
FROM films;

##
SELECT AVG(duration)
FROM films;

##
SELECT MIN(duration)
FROM films;

##
SELECT MAX(duration)
FROM films;
```

## Aggregate functions practice
```sql
##
SELECT SUM(gross)
FROM films;

##
SELECT AVG(gross)
FROM films;

##
SELECT MIN(gross)
FROM films;

##
SELECT MAX(gross)
FROM films;
```

## Combining aggregate functions with WHERE
```sql
##
SELECT SUM(gross)
FROM films
WHERE release_year >= 2000;

##
SELECT AVG(gross)
FROM films
WHERE title LIKE 'A%';

##
SELECT MIN(gross)
FROM films
WHERE release_year = 1994;

##
SELECT MAX(gross)
FROM films
WHERE release_year BETWEEN 2000 AND 2012;
```

## A note on arithmetic
```sql
SELECT (10/3);
-- 3
```

## It's AS simple AS aliasing
```sql
##
SELECT title, gross - budget AS net_profit
FROM films;

##
SELECT title, duration/60.0 AS duration_hours
FROM films;

##
SELECT AVG(duration)/60.0 AS avg_duration_hours
FROM films;
```

## Even more aliasing
```sql
##
-- get the count(deathdate) and multiply by 100.0
-- then divide by count(*)
SELECT COUNT(deathdate)*100.0/COUNT(*) AS percentage_dead
FROM people;

##
SELECT MAX(release_year) - MIN(release_year) AS difference
FROM films;

##
SELECT (MAX(release_year) - MIN(release_year)) / 10.0 AS number_of_decades
FROM films;
```




# 4. Sorting and grouping
## ORDER BY
```sql
##
SELECT name
FROM people
ORDER BY name ASC;

##
SELECT name
FROM people
ORDER BY birthdate ASC;

##
SELECT birthdate, name
FROM people
ORDER BY birthdate ASC;
```

## Sorting single columns (2)
```sql
##
SELECT title
FROM films
WHERE release_year = 2000 OR release_year = 2012;

##
SELECT *
FROM films
WHERE release_year <> 2015
ORDER BY duration ASC;

##
SELECT title, gross
FROM films
WHERE title LIKE 'M%'
ORDER BY title ASC;
```

## Sorting single columns (DESC)
```sql
##
SELECT imdb_score, film_id
FROM reviews
ORDER BY imdb_score DESC;

##
SELECT title
FROM films
ORDER BY title DESC;

##
SELECT title, duration
FROM films
ORDER BY duration DESC;
```

## Sorting multiple columns
```sql
##
SELECT birthdate, name
FROM people
ORDER BY birthdate ASC, name ASC;

##
SELECT release_year, duration, title
FROM films
ORDER BY release_year ASC, duration ASC;

##
SELECT certification, release_year, title
FROM films
ORDER BY certification ASC, release_year ASC;

##
SELECT name, birthdate
FROM people
ORDER BY name ASC, birthdate ASC;
```

## GROUP BY practice
```sql
##
SELECT release_year, COUNT(*)
FROM films
GROUP BY release_year;

##
SELECT release_year, AVG(duration)
FROM films
GROUP BY release_year;

##
SELECT release_year, MAX(budget)
FROM films
GROUP BY release_year;

##
SELECT imdb_score, COUNT(*)
FROM reviews
GROUP BY imdb_score;
```

## GROUP BY practice (2)
```sql
##
SELECT release_year, MIN(gross)
FROM films
GROUP BY release_year;

##
SELECT language, SUM(gross)
FROM films
GROUP BY language;

##
SELECT country, SUM(budget)
FROM films
GROUP BY country;

##
SELECT release_year, country, MAX(budget)
FROM films
GROUP BY release_year, country
ORDER BY release_year, country;

##
SELECT country, release_year, MIN(gross)
FROM films
GROUP BY release_year, country
ORDER BY country, release_year
```

## HAVING a great time
```sql
SELECT COUNT(t.c1) FROM (
  SELECT release_year AS c1
  FROM films
  GROUP BY release_year
  HAVING COUNT(title) > 200
) AS t;
```

## All together now
```sql
##
SELECT release_year, AVG(budget) AS avg_budget, AVG(gross) AS avg_gross
FROM films
WHERE release_year > 1990
GROUP BY release_year;

##
SELECT release_year, AVG(budget) AS avg_budget, AVG(gross) AS avg_gross
FROM films
GROUP BY release_year
HAVING AVG(budget) >= 60000000;

##
SELECT release_year, AVG(budget) AS avg_budget, AVG(gross) AS avg_gross
FROM films
WHERE release_year > 1990
GROUP BY release_year
HAVING AVG(budget) > 60000000
ORDER BY AVG(gross) DESC;
```

## All together now (2)
```sql
-- select country, average budget, 
--     and average gross
SELECT country, AVG(budget) AS avg_budget, AVG(gross) AS avg_gross
-- from the films table
FROM films
-- group by country 
GROUP BY country
-- where the country has more than 10 titles
HAVING COUNT(title) > 10
-- order by country
ORDER BY country
-- limit to only show 5 results
LIMIT 5;
```

## A taste of things to come
```sql
##
SELECT title, imdb_score
FROM films
JOIN reviews
ON films.id = reviews.film_id
WHERE title = 'To Kill a Mockingbird';
```

----------------------------------------------------------------------

# 1. Relational Databases
## What are the advantages of databases?
```
[ ]More storage
[ ]Many people can use at once
[ ]Can be secured with encryption
[x]All of the above

NOTE: Advantages of storing data in a database, rather than using traditional formats like spreadsheets, include:
- More storage: Databases can typically handle much larger volumes of data compared to spreadsheets, making them suitable for storing large datasets efficiently.
- Many people can use at once: Databases support concurrent access by multiple users, allowing for collaborative work and simultaneous data retrieval and manipulation without conflicts.
- Can be secured with encryption: Databases offer features for securing data through encryption techniques, ensuring that sensitive information remains protected from unauthorized access or breaches.
```

## Data organization
```
[ ]This is a table containing three relational databases: employees, job_levels, and departments.
[x]This is a relational database containing three tables: employees, job_levels, and departments.
[ ]This is a database, but it is not relational, because no relationship exists between job levels and departments.
[ ]This is not a database because there is no SQL code shown.
```

## Picking a unique ID
```
[ ]name
[ ]dept_id
[ ]year_hired
[x]id
```

## Setting the table in style
```
customer:
| id  | customer | phone_num    | zip_code |
| --- | -------- | ------------ | -------- |
| 567 | Jele     | 781-765-2395 | 02476    |
| 568 | Raushon  | 617-356-7772 | 01132    |
```

## Our very own table
```
SELECT *
FROM books;
```

## At your service
```
[ ]Servers can be used for storing website information as well as databases.
[ ]A server can handle requests from many computers at once.
[x]Servers are usually personal computers such as laptops.
[ ]Data from a database is physically stored on a server.

NOTE: Servers are not usually personal computers such as laptops. Servers are typically powerful computers specifically designed to handle requests and store data for multiple clients or users. They often have specialized hardware configurations optimized for performance, reliability, and scalability.
```

## Finding data types
```
[ ]You can find this information by looking at each table in the database.
[ ]You can find this information by looking at a diagram of relationships between tables.
[ ]You can find this information by looking at the values in each field for each table.
[s]You can find this information by looking at a database schema.
```

## Choice of type
```
VARCHAR
- Phone numbers such as 321-123-5555
- Product reviews written by customers

INT
- Model year such as 2004
- Number of mailing list subscribers such as 9782

NUMERIC
- Product prices in dollars such as 65.75
- Weight in tons such as 5.57
```




# 2. Querying
## SQL strengths
```
[ ]All data needed to answer the business question is presented in a spreadsheet, and no complicated relationships exist between different data points.
[x]Large amounts of data about many different but related areas of a business are housed in a relational database.
[ ]The data needed to answer the business question doesn't exist yet.

NOTE: SQL (Structured Query Language) is particularly well-suited for querying and manipulating data stored in relational databases, where data is organized into tables with defined relationships between them. Therefore, when dealing with large amounts of data stored in a structured manner within a relational database, SQL provides an efficient and effective means to retrieve and analyze the required information.
```

## Developing SQL style
```

```

## Querying the books table
```

```

## Writing queries
```

```

## Making queries DISTINCT
```

```

## Aliasing
```

```

## VIEWing your query
```

```

## SQL flavors
```

```

## Comparing flavors
```

```

## Limiting results
```

```

## Translating between flavors
```

```

## Congratulations!
```

```

