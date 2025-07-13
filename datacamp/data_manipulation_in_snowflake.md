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

## Applying conditional logic in Snowflake
```

```

## Comparing invoice totals
```

```

## Validating data quality
```

```

## How many protected files?
```

```




# 2 Manipulating Data with Subqueries and Common Table Expressions
## Subqueries
```

```

## Building subqueries
```

```

## Are jazz songs long?
```

```

## Identifying large transactions
```

```

## Common Table Expressions
```

```

## Analyzing track length
```

```

## Finding the most efficient composer
```

```

## Where are customers buying?
```

```

## Advanced Common Table Expressions
```

```

## Building a detailed invoice
```

```

## Finding the most popular artists
```

```

## Wrapping up!
```

```

## Albums driving sales
```

```

## Hot genres
```

```
