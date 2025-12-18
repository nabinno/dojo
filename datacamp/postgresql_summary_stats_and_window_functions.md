---
title: PostgreSQL Summary Stats and Window Functions
tags: postgresql,structured-query-language
url: https://campus.datacamp.com/courses/postgresql-summary-stats-and-window-functions
---

# 1 Introduction to window functions
## Window functions vs GROUP BY
```
Which of the following is FALSE?
[ ]Unlike GROUP BY results, window functions don't reduce the number of rows in the output.
[ ]Window functions can fetch values from other rows into the table, whereas GROUP BY functions cannot.
[x]Window functions can open a "window" to another table, whereas GROUP BY functions cannot.
[ ]Window functions can calculate running totals and moving averages, whereas GROUP BY functions cannot.
```

## Numbering rows
```sql
SELECT
  *,
  -- Assign numbers to each row
  ROW_NUMBER() OVER() AS Row_N
FROM Summer_Medals
ORDER BY Row_N ASC;
```

## Numbering Olympic games in ascending order
```sql
SELECT
  Year,

  -- Assign numbers to each year
  ROW_NUMBER() OVER() AS Row_N
FROM (
  SELECT DISTINCT Year
  FROM Summer_Medals
  ORDER BY Year ASC
) AS Years
ORDER BY Year ASC;
```

## ORDER BY
```sql

```

## Numbering Olympic games in descending order
```sql

```

## Numbering Olympic athletes by medals earned
```sql

```

## Reigning weightlifting champions
```sql

```

## PARTITION BY
```sql

```

## Reigning champions by gender
```sql

```

## Reigning champions by gender and event
```sql

```

## Row numbers with partitioning
```sql

```




# 2 Fetching, ranking, and paging
## Fetching
```sql

```

## Future gold medalists
```sql

```

## First athlete by name
```sql

```

## Last country by name
```sql

```

## Ranking
```sql

```

## Ranking athletes by medals earned
```sql

```

## Ranking athletes from multiple countries
```sql

```

## DENSE_RANK's output
```sql

```

## Paging
```sql

```

## Paging events
```sql

```

## Top, middle, and bottom thirds
```sql

```




# 3 Aggregate window functions and frames
## Aggregate window functions
```sql

```

## Running totals of athlete medals
```sql

```

## Maximum country medals by year
```sql

```

## Minimum country medals by year
```sql

```

## Frames
```sql

```

## Number of rows in a frame
```sql

```

## Moving maximum of Scandinavian athletes' medals
```sql

```

## Moving maximum of Chinese athletes' medals
```sql

```

## Moving averages and totals
```sql

```

## Moving average's frame
```sql

```

## Moving average of Russian medals
```sql

```

## Moving total of countries' medals
```sql

```




# 4 Beyond window functions
## Pivoting
```sql

```

## A basic pivot
```sql

```

## Pivoting with ranking
```sql

```

## ROLLUP and CUBE
```sql

```

## Country-level subtotals
```sql

```

## All group-level subtotals
```sql

```

## A survey of useful functions
```sql

```

## Cleaning up results
```sql

```

## Summarizing results
```sql

```
