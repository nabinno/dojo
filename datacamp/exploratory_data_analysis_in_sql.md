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
```

```

## The keys to the database
```

```

## Foreign keys
```

```

## Read an entity relationship diagram
```

```

## Coalesce
```

```

## Column types and constraints
```

```

## Effects of casting
```

```

## Summarize the distribution of numeric values
```

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
