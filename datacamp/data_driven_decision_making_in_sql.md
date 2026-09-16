---
title: Data-Driven Decision Making in SQL
tags: structured-query-language, analytics
url: https://campus.datacamp.com/courses/data-driven-decision-making-in-sql
---

# 1 Introduction to business intelligence for a online movie rental database
## Exploring the database
```
Explore the tables and its columns. Which of the following quantities can't be computed?

[ ]The number of customers from each country.
[x]The number of movies with an international award.
[ ]The average rating of a movie.
[ ]The number of movies with the actor Daniel Radcliffe.
```

## Exploring the table renting
```sql
-- 1)
SELECT *  -- Select all
FROM renting;        -- From table renting

-- 2)
SELECT movie_id,  -- Select all columns needed to compute the average rating per movie
       rating
FROM renting;
```

## Working with dates
```sql
-- 1)
SELECT *
FROM renting
WHERE date_renting = '2018-10-09'; -- Movies rented on October 9th, 2018

-- 2)
SELECT *
FROM renting
WHERE date_renting BETWEEN '2018-04-01' AND '2018-08-31'; -- from beginning April 2018 to end August 2018

-- 3)
SELECT *
FROM renting
WHERE date_renting BETWEEN '2018-04-01' AND '2018-08-31'
ORDER BY date_renting ASC; -- Order by recency in decreasing order
```

## Selecting movies
```sql
-- 1)
SELECT *
FROM movies
WHERE genre <> 'Drama'; -- All genres except drama

-- 2)
SELECT *
FROM movies
WHERE title IN ('Showtime', 'Love Actually', 'The Fighter'); -- Select all movies with the given titles

-- 3)
SELECT *
FROM movies
ORDER BY renting_price ASC ; -- Order the movies by increasing renting price
```

## Select from renting
```sql
SELECT *
FROM renting
WHERE date_renting BETWEEN '2018-01-01' AND '2018-12-31' -- Renting in 2018
AND rating IS NOT NULL; -- Rating exists
```

## Summarizing customer information
```sql
-- 1)
SELECT COUNT(*) -- Count the total number of customers
FROM customers
WHERE date_of_birth BETWEEN '1980-01-01' AND '1989-12-31'; -- Select customers born between 1980-01-01 and 1989-12-31

-- 2)
SELECT COUNT(*)   -- Count the total number of customers
FROM customers
WHERE country = 'Germany'; -- Select all customers from Germany

-- 3)
SELECT COUNT(DISTINCT country)   -- Count the number of countries
FROM customers;
```

## Ratings of movie 25
```sql
SELECT MIN(rating) min_rating, -- Calculate the minimum rating and use alias min_rating
       MAX(rating) max_rating, -- Calculate the maximum rating and use alias max_rating
       AVG(rating) avg_rating, -- Calculate the average rating and use alias avg_rating
       COUNT(rating) number_ratings -- Count the number of ratings and use alias number_ratings
FROM renting
WHERE movie_id = 25; -- Select all records of the movie with ID 25
```

## Examining annual rentals
```sql
-- 1)
SELECT * -- Select all records of movie rentals since January 1st 2019
FROM renting
WHERE date_renting >= '2019-01-01';

-- 2)
SELECT
    COUNT(*), -- Count the total number of rented movies
    AVG(rating) -- Add the average rating
FROM renting
WHERE date_renting >= '2019-01-01';

-- 3)
SELECT
    COUNT(*) AS number_renting, -- Give it the column name number_renting
    AVG(rating) AS average_rating  -- Give it the column name average_rating
FROM renting
WHERE date_renting >= '2019-01-01';

-- 4)
SELECT
    COUNT(*) AS number_renting,
    AVG(rating) AS average_rating,
    COUNT(rating) AS number_ratings -- Add the total number of ratings here.
FROM renting
WHERE date_renting >= '2019-01-01';
```




# 2 Decision Making with simple SQL queries
## First account for each country.
```sql
SELECT country, -- For each country report the earliest date when an account was created
    MIN(date_account_start) AS first_account
FROM customers
GROUP BY country
ORDER BY first_account;
```

## Average movie ratings
```sql
-- 1)
SELECT movie_id,
       AVG(rating)    -- Calculate average rating per movie
FROM renting
GROUP BY movie_id;

-- 2)
SELECT movie_id,
       AVG(rating) AS avg_rating, -- Use as alias avg_rating
       COUNT(rating) AS number_rating,                -- Add column for number of ratings with alias number_rating
       COUNT(*) AS number_renting                 -- Add column for number of movie rentals with alias number_renting
FROM renting
GROUP BY movie_id;

-- 3)
SELECT movie_id,
       AVG(rating) AS avg_rating,
       COUNT(rating) AS number_ratings,
       COUNT(*) AS number_renting
FROM renting
GROUP BY movie_id
ORDER BY avg_rating DESC; -- Order by average rating in decreasing order
```

## Average rating per customer
```sql
SELECT customer_id,  -- Report the customer_id
       AVG(rating), -- Report the average rating per customer
       COUNT(rating), -- Report the number of ratings per customer
       COUNT(*) -- Report the number of movie rentals per customer
FROM renting
GROUP BY customer_id
HAVING COUNT(*) > 7 -- Select only customers with more than 7 movie rentals
ORDER BY AVG(rating); -- Order by the average rating in ascending order
```

## Join renting and customers
```sql
-- 1)
SELECT * -- Join renting with customers
FROM renting AS r
LEFT JOIN customers AS c
ON r.customer_id = c.customer_id;

-- 2)
SELECT *
FROM renting AS r
LEFT JOIN customers AS c
ON r.customer_id = c.customer_id
WHERE c.country = 'Belgium'; -- Select only records from customers coming from Belgium

-- 3)
SELECT AVG(r.rating) -- Average ratings of customers from Belgium
FROM renting AS r
LEFT JOIN customers AS c
ON r.customer_id = c.customer_id
WHERE c.country = 'Belgium';
```

## Aggregating revenue, rentals and active customers
```sql
-- 1)
SELECT *
FROM renting AS r
LEFT JOIN movies AS m -- Choose the correct join statment
ON r.movie_id = m.movie_id;

-- 2)
SELECT
    SUM(m.renting_price), -- Get the revenue from movie rentals
    COUNT(*), -- Count the number of rentals
    COUNT(DISTINCT r.customer_id)  -- Count the number of customers
FROM renting AS r
LEFT JOIN movies AS m
ON r.movie_id = m.movie_id;

-- 3)
SELECT
    SUM(m.renting_price),
    COUNT(*),
    COUNT(DISTINCT r.customer_id)
FROM renting AS r
LEFT JOIN movies AS m
ON r.movie_id = m.movie_id
-- Only look at movie rentals in 2018
WHERE date_renting BETWEEN '2018-01-01' AND '2018-12-31';
```

## Movies and actors
```sql
SELECT DISTINCT m.title, -- Create a list of movie titles and actor names
                a.name
FROM actsin AS ai
LEFT JOIN movies AS m
ON m.movie_id = ai.movie_id
LEFT JOIN actors AS a
ON a.actor_id = ai.actor_id;
```

## Income from movies
```sql
-- 1)
SELECT m.title, -- Use a join to get the movie title and price for each movie rental
       m.renting_price
FROM renting AS r
LEFT JOIN movies AS m
ON r.movie_id = m.movie_id;

-- 2)
SELECT title, -- Report the income from movie rentals for each movie
       SUM(renting_price) AS income_movie
FROM
       (SELECT m.title,
               m.renting_price
       FROM renting AS r
       LEFT JOIN movies AS m
       ON r.movie_id = m.movie_id) AS rm
GROUP BY title
ORDER BY income_movie DESC; -- Order the result by decreasing income
```

## Age of actors from the USA
```sql
SELECT gender, -- Report for male and female actors from the USA
       MIN(year_of_birth), -- The year of birth of the oldest actor
       MAX(year_of_birth) -- The year of birth of the youngest actor
FROM
   (SELECT *
    FROM actors
    WHERE nationality = 'USA') -- Use a subsequent SELECT to get all information about actors from the USA
   AS a -- Give the table the name a
GROUP BY gender;
```

## Identify favorite movies for a group of customers
```sql
SELECT m.title,
COUNT(*),
AVG(r.rating)
FROM renting AS r
LEFT JOIN customers AS c
ON c.customer_id = r.customer_id
LEFT JOIN movies AS m
ON m.movie_id = r.movie_id
WHERE c.date_of_birth BETWEEN '1970-01-01' AND '1979-12-31'
GROUP BY m.title
HAVING COUNT(*) > 1 -- Remove movies with only one rental
ORDER BY AVG(r.rating); -- Order with highest rating first
```

## Identify favorite actors for Spain
```sql
SELECT a.name,  c.gender,
       COUNT(*) AS number_views,
       AVG(r.rating) AS avg_rating
FROM renting as r
LEFT JOIN customers AS c
ON r.customer_id = c.customer_id
LEFT JOIN actsin as ai
ON r.movie_id = ai.movie_id
LEFT JOIN actors as a
ON ai.actor_id = a.actor_id
WHERE c.country = 'Spain' -- Select only customers from Spain
GROUP BY a.name, c.gender
HAVING AVG(r.rating) IS NOT NULL
  AND COUNT(*) > 5
ORDER BY avg_rating DESC, number_views DESC;
```

## KPIs per country
```sql
-- 1)
SELECT *
FROM renting r -- Augment the table renting with information about customers
LEFT JOIN customers c
ON r.customer_id = c.customer_id
LEFT JOIN movies m -- Augment the table renting with information about movies
ON r.movie_id = m.movie_id
WHERE r.date_renting >= '2019-01-01'; -- Select only records about rentals since the beginning of 2019

-- 2)
SELECT
    c.country,                     -- For each country report
    COUNT(*) AS number_renting,    -- The number of movie rentals
    AVG(r.rating) AS average_rating, -- The average rating
    SUM(m.renting_price) AS revenue  -- The revenue from movie rentals
FROM renting AS r
LEFT JOIN customers AS c
ON c.customer_id = r.customer_id
LEFT JOIN movies AS m
ON m.movie_id = r.movie_id
WHERE date_renting >= '2019-01-01'
GROUP BY c.country;
```





# 3 Data Driven Decision Making with advanced SQL queries
## Nested query
```sql

```

## Often rented movies
```sql

```

## Frequent customers
```sql

```

## Movies with rating above average
```sql

```

## Correlated nested queries
```sql

```

## Analyzing customer behavior
```sql

```

## Customers who gave low ratings
```sql

```

## Movies and ratings with correlated queries
```sql

```

## Queries with EXISTS
```sql

```

## Customers with at least one rating
```sql

```

## Actors in comedies
```sql

```

## Queries with UNION and INTERSECT
```sql

```

## Young actors not coming from the USA
```sql

```

## Dramas with high ratings
```sql

```




# 4 Data Driven Decision Making with OLAP SQL queries
## OLAP: CUBE operator
```sql

```

## Groups of customers
```sql

```

## Categories of movies
```sql

```

## Analyzing average ratings
```sql

```

## ROLLUP
```sql

```

## Number of customers
```sql

```

## Analyzing preferences of genres across countries
```sql

```

## GROUPING SETS
```sql

```

## Queries with GROUPING SETS
```sql

```

## Exploring nationality and gender of actors
```sql

```

## Exploring rating by country and gender
```sql

```

## Bringing it all together
```sql

```

## Customer preference for genres
```sql

```

## Customer preference for actors
```sql

```
