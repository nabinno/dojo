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

## Filtering and ordering
```sql

```

## Working with dates
```sql

```

## Selecting movies
```sql

```

## Select from renting
```sql

```

## Aggregations - summarizing data
```sql

```

## Summarizing customer information
```sql

```

## Ratings of movie 25
```sql

```

## Examining annual rentals
```sql

```




# 2 Decision Making with simple SQL queries
## Grouping movies
```sql

```

## First account for each country.
```sql

```

## Average movie ratings
```sql

```

## Average rating per customer
```sql

```

## Joining movie ratings with customer data
```sql

```

## Join renting and customers
```sql

```

## Aggregating revenue, rentals and active customers
```sql

```

## Movies and actors
```sql

```

## Money spent per customer with sub-queries
```sql

```

## Income from movies
```sql

```

## Age of actors from the USA
```sql

```

## Identify favorite actors of customer groups
```sql

```

## Identify favorite movies for a group of customers
```sql

```

## Identify favorite actors for Spain
```sql

```

## KPIs per country
```sql

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
