---
title: Data Analysis in SQL (PostgreSQL)
tags: postgresql
url: https://assessment.datacamp.com/data-analysis-in-sql
---

## JOIN
```sql
SELECT t.name, f.loudness
FROM tracks AS t
JOIN features AS f
    ON t.id = f.song_id
ORDER BY f.loudness, t.name
LIMIT 10;
```

## TRANSFORM
```sql
SELECT order_id, 
       quantity, 
       transform(quantity, 1) OVER() AS pre_quantity
FROM orders
ORDER BY order_id
```

## LOWER
```sql
SELECT user_id,
       LOWER(title) AS title_lower
FROM favorite
ORDER BY user_id, title_lower
LIMIT 5;
```

## COUNT, GROUP BY
```sql
SELECT artist_id, COUNT(id) AS artist_count
FROM tracks
GROUP BY artist_id
ORDER BY artist_count DESC, artist_id
LIMIT 5;
```

## Sub query
```sql
SELECT style, price
FROM wine_region
WHERE id IN(
	SELECT wine_id
	FROM pairing
)
ORDER BY price, style
LIMIT 5
```

## JOIN
```sql
SELECT awards.year, 
       artists.name 
FROM awards
JOIN artists
ON artists.id = awards.artist_id
ORDER BY awards.year DESC
LIMIT 5;
```

## CONCAT
```sql
SELECT CONCAT(first_name, last_name) AS full_name 
FROM directors 
ORDER BY full_name;
```

## Sub query
```sql
SELECT count_employees
FROM (SELECT manager_id,
            Count(employee_id) AS count_employees
           FROM employees
           GROUP BY manager_id) AS employee_summary
```

## MONTH
```sql
SELECT MONTH(month FROM timestamp '2005-01-24') AS month;
```

## Subquery
```sql
SELECT show_id, rating
FROM ratings
AS subquery
ORDER BY show_id
LIMIT 10;
```

## COUNT
```sql
SELECT COUNT(id) AS long_songs
FROM tracks
WHERE duration_ms > 100000;
```

## WHERE
```sql
SELECT name, nationality
FROM artists
WHERE nationality IN ('British', 'South African', 'Mexican')
ORDER BY name
LIMIT 10;
```

## ORDER BY
```sql
SELECT name, popularity 
FROM tracks
ORDER BY popularity DESC
LIMIT 10;
```
