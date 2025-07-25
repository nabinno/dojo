---
title: Data manipulation in Snowflake
tags: snowflake,analytics,data-engineering
url: https://campus.datacamp.com/courses/data-manipulation-in-snowflake/
---

# 1 Conditional Logic
## Explore the Store data
```
1. Take a peek in [Snowflake > Data > Databases COURSE_40931* > Schema Store].

2. Take a closer look at each of these tables, the columns, and the data types in  [schema STORE > tables].

3. Check the data type of [table TRACK > column track_id].
```

## Is this a premium song?
```sql
SELECT
    name,
    composer,
    -- Begin a CASE statement
    CASE
        -- A song priced at 0.99 should be a 'Standard Song'
        WHEN unit_price = 0.99 THEN 'Standard Song'
        -- Songs costing 1.99 should be denoted as 'Premium Song'
        WHEN unit_price = 1.99 THEN 'Premium Song'
    END AS song_description
FROM store.track;
```

## Inferring purchase quantity
```sql
SELECT
    customer_id,
    total,
    CASE
        -- Check if total is either 0.99 or 1.99 using IN
        WHEN total IN (0.99, 1.99) THEN '1 Song'
        -- Catch the scenarios when the above is not true
        ELSE '2+ Songs'
    -- End the CASE statement and name the new column
    END AS number_of_songs
FROM store.invoice;
```

## Is this a long song?
```sql
SELECT
    name,
    milliseconds,
    CASE
        WHEN milliseconds < 180000 THEN 'Short Song'
        WHEN milliseconds BETWEEN 180000 AND 300000 THEN 'Normal Length'
        ELSE 'Long Song'
    END AS song_length
FROM store.track;
```

## Determining buyer intent
```sql
SELECT
    name,
    unit_price,
    CASE
        -- Inexpensive Rock and Pop songs are always high-intent
        WHEN unit_price = 0.99 AND genre_id IN (5, 9) THEN 'High'
        -- Shorter, non-EDM tracks have neutral buyer intent
        WHEN milliseconds < 300000 AND genre_id != 15 THEN 'Neutral'
        -- Everything else is low
        ELSE 'Low'
    END AS buyer_intent
FROM store.track;
```

## Comparing invoice totals
```sql
-- Run this query without editing it.
SELECT
    customer_id,
    total,
    CASE
        WHEN total IN (0.99, 1.99) THEN '1 Song'
        ELSE '2+ Songs'
    END as number_of_songs
FROM store.invoice;

SELECT
    CASE
        WHEN total IN (0.99, 1.99) THEN '1 Song'
        ELSE '2+ Songs'
    END as number_of_songs,
    -- Find the average value of the total field
    AVG(total) AS average_total
FROM store.invoice
-- Group by the field you built using CASE
GROUP BY number_of_songs;
```

## Validating data quality
```sql
SELECT
    track.name,
    track.composer,
    artist.name,
    CASE
        -- A 'Track Lacks Detail' if the composer field is NULL
        WHEN track.composer IS NULL THEN 'Track Lacks Detail'
        -- Use the composer and artist name to determine if a match exists
        WHEN track.composer = artist.name THEN 'Matching Artist'
        ELSE 'Inconsistent Data'
    END AS data_quality
FROM store.track AS track
LEFT JOIN store.album AS album ON track.album_id = album.album_id
-- Join the album table to artist using the artist_id field
LEFT JOIN store.artist AS artist ON album.artist_id = artist.artist_id;
```

## How many protected files?
```sql
SELECT count(*)
FROM track
WHERE media_type IN (2, 3);
```




# 2 Manipulating Data with Subqueries and Common Table Expressions
## Building subqueries
```sql
SELECT
    invoice_id,
    SUM(quantity * unit_price)
FROM (
    SELECT invoice_id, quantity, unit_price FROM store.invoiceline
)
GROUP BY invoice_id;
```

## Are jazz songs long?
```sql
SELECT
    -- Find the genre name and average milliseconds
    genre_name,
    AVG(milliseconds) AS average_milliseconds
-- Retrieve records from the result of the subquery
FROM (
    SELECT
        genre.name AS genre_name,
        track.genre_id,
        track.milliseconds
    FROM store.track
    JOIN store.genre ON track.genre_id = genre.genre_id
)
-- Group the results by the genre name
GROUP BY genre_name;
```

## Identifying large transactions
```sql
SELECT
    invoice_id,
    COUNT(invoice_id) AS total_invoice_lines
FROM store.invoiceline
GROUP BY invoice_id
-- Only pull records with more than 10 total invoice lines
HAVING total_invoice_lines > 10;

SELECT
  billing_country,
  SUM(total) AS total_invoice_amount
FROM store.invoice
WHERE invoice_id IN (
  SELECT
      invoice_id,
  FROM store.invoiceline
  GROUP BY invoice_id
  HAVING COUNT(invoice_id) > 10
)
GROUP BY billing_country;
```

## Analyzing track length
```sql
-- Create a CTE named track_lengths
WITH track_lengths AS (
    SELECT
        genre.name,
        track.genre_id,
        track.milliseconds / 1000 AS num_seconds
    FROM store.track
    JOIN store.genre ON track.genre_id = genre.genre_id
)

SELECT
    track_lengths.name,
    -- Find the average length of each track in seconds
    AVG(track_lengths.num_seconds) AS avg_track_length
FROM track_lengths
GROUP BY track_lengths.name
-- Sort the results by average track_length
ORDER BY avg_track_length DESC;
```

## Finding the most efficient composer
```sql
-- Create a CTE called track_metrics, convert milliseconds to seconds
WITH track_metrics AS (
    SELECT
        composer,
        milliseconds / 1000 AS num_seconds,
        unit_price
    FROM store.track
    -- Retrieve records where composer is not NULL
    WHERE composer IS NOT NULL
)

SELECT
    composer,
    -- Find the average price-per-second
    AVG(unit_price / num_seconds) AS cost_per_second
FROM track_metrics
GROUP BY composer
ORDER BY cost_per_second DESC;
```

## Where are customers buying?
```sql
WITH usa_invoices AS (
  SELECT
    invoice_id,
    billing_state
  FROM invoice
  WHERE billing_country = 'USA'
)

SELECT
  COUNT(quantity)
FROM usa_invoices
JOIN invoiceline
GROUP BY billing_state
ORDER BY COUNT(quantity) DESC;
```

## Building a detailed invoice
```sql
WITH cleaned_invoices AS (
    SELECT
        invoice_id,
        invoice_date
    FROM store.invoice
    WHERE billing_country = 'Germany'
),

detailed_invoice_lines AS (
    SELECT
        invoiceline.invoice_id,
        invoiceline.invoice_line_id,
        track.name,
        invoiceline.unit_price,
        invoiceline.quantity,
    FROM store.invoiceline
    LEFT JOIN store.track ON invoiceline.track_id = track.track_id
)

SELECT
    ci.invoice_id,
    ci.invoice_date,
    dil.name,
    -- Find the total amount for the line
    dil.unit_price * dil.quantity AS line_amount
FROM detailed_invoice_lines AS dil

-- JOIN the cleaned_invoices and detailed_invoice_lines CTEs
LEFT JOIN cleaned_invoices AS ci ON dil.invoice_id = ci.invoice_id
ORDER BY ci.invoice_id, line_amount;
```

## Finding the most popular artists
```sql
-- Create an artist_info CTE, JOIN the artist and album tables
WITH artist_info AS (
    SELECT
        album.album_id,
        artist.name AS artist_name
    FROM store.album
    JOIN store.artist ON album.artist_id = artist.artist_id

-- Define a track_sales CTE to assign an album_id, name,
-- and number of seconds for each track
), track_sales AS (
    SELECT
        track.album_id,
        track.name,
        track.milliseconds / 1000 AS num_seconds
    FROM store.invoiceline
    JOIN store.track ON invoiceline.track_id = track.track_id
)

SELECT
    ai.artist_name,
    -- Calculate total minutes listed
    SUM(num_seconds) / 60 AS minutes_listened
FROM track_sales AS ts
JOIN artist_info AS ai ON ts.album_id = ai.album_id
-- Group the results by the non-aggregated column
GROUP BY ai.artist_name
ORDER BY minutes_listened DESC;
```

## Albums driving sales
```sql
-- Define an album_map CTE to combine albums and artists
WITH album_map AS (
    SELECT
        album.album_id, album.title AS album_name, artist.name AS artist_name,
        -- Determine if an album is a "Greatest Hits" album
        CASE
            WHEN album_name ILIKE '%greatest%' THEN TRUE
            ELSE FALSE
        END AS is_greatest_hits
    FROM store.album
    JOIN store.artist ON album.artist_id = artist.artist_id
), trimmed_invoicelines AS (
    SELECT
        invoiceline.invoice_id, track.album_id, invoice.total
    FROM store.invoiceline
    LEFT JOIN store.invoice ON invoiceline.invoice_id = invoice.invoice_id
    LEFT JOIN store.track ON invoiceline.track_id = track.track_id
)

SELECT
    album_map.album_name,
    album_map.artist_name,
    SUM(ti.total) AS total_sales_driven
FROM trimmed_invoicelines AS ti
JOIN album_map ON ti.album_id = album_map.album_id
-- Use a subquery to only "Greatest Hits" records
WHERE ti.album_id IN (SELECT album_id FROM album_map WHERE is_greatest_hits)
GROUP BY album_map.album_name, album_map.artist_name, is_greatest_hits
ORDER BY total_sales_driven DESC;
```

## Hot genres
```

```
