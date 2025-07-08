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

## Conditional logic with CASE statements
```

```

## Inferring purchase quantity
```

```

## Is this a long song?
```

```

## Determining buyer intent
```

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
