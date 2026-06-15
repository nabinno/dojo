---
title: Exploratory Data Analysis in SQL
tags: statistics, structured-query-language
url: https://campus.datacamp.com/courses/exploratory-data-analysis-in-sql/whats-in-the-database
---

# 1 What's in the Database?
## Explore table sizes
```txt
Which table has the most rows?

[ ]fortune500 has the most rows 500
[ ]tag_company has the most rows => 56
[x]stackoverflow has the most rows => 45238
[ ]tag_type has the most rows => 61
```

## Count missing values
```sql
-- Count the number of null values in the ticker column
SELECT count(*) - count(ticker) AS missing
  FROM fortune500;

-- Count the number of null values in the industry column
SELECT COUNT(*) - COUNT(industry) AS missing
FROM fortune500;
```

## Join tables
```sql
SELECT company.name
-- Table(s) to select from
  FROM company
       INNER JOIN fortune500
       ON company.ticker=fortune500.ticker;
```
## Foreign keys
```
[ ]stackoverflow.tag is not a primary key
[ ]tag_type.tag contains NULL values
[x]stackoverflow.tag contains duplicate values
[ ]tag_type.tag does not contain all the values in stackoverflow.tag

A foreign key must reference a unique column. Since `stackoverflow.tag` has duplicate values, it cannot be a foreign key target.
```

## Read an entity relationship diagram
```sql
-- Count the number of tags with each type
SELECT type, COUNT(*) AS count
  FROM tag_type
 -- To get the count for each type, what do you need to do?
 GROUP BY type
 -- Order the results with the most common tag types listed first
ORDER BY count DESC;

-- Select the 3 columns desired
SELECT company.name, tag_type.tag, tag_type.type
  FROM company
       -- Join to the tag_company table
       INNER JOIN tag_company
       ON company.id = tag_company.company_id
       -- Join to the tag_type table
       INNER JOIN tag_type
       ON tag_company.tag = tag_type.tag
  -- Filter to most common type
  WHERE type='cloud';
```

## Coalesce
```sql
-- Use coalesce
SELECT COALESCE(industry, sector, 'Unknown') AS industry2,
       -- Don't forget to count!
       COUNT(*)
  FROM fortune500
-- Group by what? (What are you counting by?)
 GROUP BY industry2
-- Order results to see most common first
 ORDER BY COUNT(*) DESC
-- Limit results to get just the one value you want
 LIMIT 1;
```

## Effects of casting
```sql
-- Select the original value
SELECT profits_change,
       -- Cast profits_change
       CAST(profits_change AS integer) AS profits_change_int
  FROM fortune500;

-- Divide 10 by 3
SELECT 10/3,
       -- Cast 10 as numeric and divide by 3
       10::numeric/3;

SELECT '3.2'::numeric,
       '-123'::numeric,
       '1e3'::numeric,
       '1e-3'::numeric,
       '02314'::numeric,
       '0002'::numeric;
```

## Summarize the distribution of numeric values
```sql
-- Select the count of each value of revenues_change
SELECT revenues_change, count(*)
  FROM fortune500
 GROUP BY revenues_change
 -- order by the values of revenues_change
 ORDER BY revenues_change;

-- Select the count of each revenues_change integer value
SELECT revenues_change::integer, count(*)
  FROM fortune500
 GROUP BY revenues_change::integer
 -- order by the values of revenues_change
 ORDER BY revenues_change::integer;

-- Count rows
SELECT count(*)
  FROM fortune500
 -- Where...
 WHERE revenues_change > 0;
```




# 2 Summarizing and Aggregating Numeric Data
## Numeric data types and summary functions
```

```

## Division
```

```

## Explore with division
```

```

## Summarize numeric columns
```

```

## Summarize group statistics
```

```

## Exploring distributions
```

```

## Truncate
```

```

## Generate series
```

```

## More summary functions
```

```

## Correlation
```

```

## Mean and Median
```

```

## Creating temporary tables
```

```

## Create a temp table
```

```

## Create a temp table to simplify a query
```

```

## Insert into a temp table
```

```




# 3 Exploring Categorical Data and Unstructured Text
## Character data types and common issues
```

```

## Count the categories
```

```

## Spotting character data problems
```

```

## Cases and spaces
```

```

## Trimming
```

```

## Exploring unstructured text
```

```

## Splitting and concatenating text
```

```

## Concatenate strings
```

```

## Split strings on a delimiter
```

```

## Shorten long strings
```

```

## Strategies for multiple transformations
```

```

## Create an "other" category
```

```

## Group and recode values
```

```

## Create a table with indicator variables
```

```




# 4 Working with Dates and Timestamps
## Date/time types and formats
```

```

## ISO 8601
```

```

## Date comparisons
```

```

## Date arithmetic
```

```

## Completion time by category
```

```

## Date/time components and aggregation
```

```

## Date parts
```

```

## Variation by day of week
```

```

## Date truncation
```

```

## Aggregating with date/time series
```

```

## Find missing dates
```

```

## Custom aggregation periods
```

```

## Monthly average with missing dates
```

```

## Time between events
```

```

## Longest gap
```

```

## Rats!
```

```

## Wrap-up
```

```
