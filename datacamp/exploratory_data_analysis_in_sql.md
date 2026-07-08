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
## Division
```sql
-- Select average revenue per employee by sector
SELECT sector,
       avg(revenues/employees::numeric) AS avg_rev_employee
  FROM fortune500
 GROUP BY sector
 -- Use the column alias to order the results
 ORDER BY avg_rev_employee;
```

## Explore with division
```sql
-- Divide unanswered_count by question_count
SELECT unanswered_count/question_count::numeric AS computed_pct,
       -- What are you comparing the above quantity to?
       unanswered_pct
  FROM stackoverflow
 -- Select rows where question_count is not 0
 WHERE question_count <> 0
 LIMIT 10;
```

## Summarize numeric columns
```sql
-- Select min, avg, max, and stddev of fortune500 profits
SELECT min(profits),
       avg(profits),
       max(profits),
       stddev(profits)
  FROM fortune500;

-- Select sector and summary measures of fortune500 profits
SELECT sector,
       min(profits),
       avg(profits),
       max(profits),
       stddev(profits)
  FROM fortune500
 -- What to group by?
 GROUP BY sector
 -- Order by the average profits
 ORDER BY avg
```

## Summarize group statistics
```sql
-- Compute standard deviation of maximum values
SELECT stddev(maxval),
       -- min
       min(maxval),
       -- max
       max(maxval),
       -- avg
       avg(maxval)
  -- Subquery to compute max of question_count by tag
  FROM (SELECT max(question_count) AS maxval
          FROM stackoverflow
         -- Compute max by...
         GROUP BY tag) AS max_results; -- alias for subquery
```

## Truncate
```sql
-- Truncate employees
SELECT trunc(employees, -5) AS employee_bin,
       -- Count number of companies with each truncated value
       count(*)
  FROM fortune500
 -- Use alias to group
 GROUP BY employee_bin
 -- Use alias to order
 ORDER BY employee_bin;

-- Truncate employees
SELECT trunc(employees, -4) AS employee_bin,
       -- Count number of companies with each truncated value
       count(*)
  FROM fortune500
 -- Limit to which companies?
 WHERE employees < 100000
 -- Use alias to group
 GROUP BY employee_bin
 -- Use alias to order
 ORDER BY employee_bin;
```

## Generate series
```sql
-- Select the min and max of question_count
SELECT MIN(question_count),
       MAX(question_count)
  -- From what table?
  FROM stackoverflow
 -- For tag dropbox
 WHERE tag = 'dropbox';

-- Create lower and upper bounds of bins
SELECT generate_series(2200, 3050, 50) AS lower,
       generate_series(2250, 3100, 50) AS upper;

-- Bins created in Step 2
WITH bins AS (
      SELECT generate_series(2200, 3050, 50) AS lower,
             generate_series(2250, 3100, 50) AS upper),
     -- Subset stackoverflow to just tag dropbox (Step 1)
     dropbox AS (
      SELECT question_count
        FROM stackoverflow
       WHERE tag='dropbox')
-- Select columns for result
-- What column are you counting to summarize?
SELECT lower, upper, count(question_count)
  FROM bins  -- Created above
       -- Join to dropbox (created above), keeping all rows from the bins table in the join
       LEFT JOIN dropbox
       -- Compare question_count to lower and upper
         ON question_count >= lower
        AND question_count < upper
 -- Group by lower and upper to count values in each bin
 GROUP BY lower, upper
 -- Order by lower to put bins in order
 ORDER BY lower;
```

## Correlation
```sql
-- Correlation between revenues and profit
SELECT corr(revenues, profits) AS rev_profits,
       -- Correlation between revenues and assets
       corr(revenues, assets) AS rev_assets,
       -- Correlation between revenues and equity
       corr(revenues, equity) AS rev_equity
FROM fortune500;
```

## Mean and Median
```sql
-- What groups are you computing statistics by?
SELECT sector,
       -- Select the mean of assets with the avg function
       avg(assets) AS mean,
       -- Select the median
       percentile_disc(0.5) WITHIN GROUP (ORDER BY assets) AS median
  FROM fortune500
 -- Computing statistics for each what?
 GROUP BY sector
 -- Order results by a value of interest
 ORDER BY mean;
```

## Create a temp table
```sql
-- To clear table if it already exists; fill in name of temp table
DROP TABLE IF EXISTS profit80;

-- Create the temporary table
CREATE TEMP TABLE profit80 AS
  -- Select the two columns you need; alias as needed
  SELECT sector,
         percentile_disc(0.8) WITHIN GROUP (ORDER BY profits) AS pct80
    -- What table are you getting the data from?
    FROM fortune500
   -- What do you need to group by?
   GROUP BY sector;

-- See what you created: select all columns and rows from the table you created
SELECT *
  FROM profit80;

---

-- Code from previous step
DROP TABLE IF EXISTS profit80;

CREATE TEMP TABLE profit80 AS
  SELECT sector,
         percentile_disc(0.8) WITHIN GROUP (ORDER BY profits) AS pct80
    FROM fortune500
   GROUP BY sector;

-- Select columns, aliasing as needed
SELECT title, fortune500.sector,
       profits, profits/pct80 AS ratio
-- What tables do you need to join?
  FROM fortune500
       LEFT JOIN profit80
-- How are the tables joined?
       ON fortune500.sector=profit80.sector
-- What rows do you want to select?
 WHERE profits > pct80;
```

## Create a temp table to simplify a query
```sql
-- 1)
-- To clear table if it already exists
DROP TABLE IF EXISTS startdates;

-- Create temp table syntax
CREATE TEMP TABLE startdates AS
-- Compute the minimum date for each what?
SELECT tag,
       MIN(date) AS mindate
  FROM stackoverflow
 -- What do you need to compute the min date for each tag?
 GROUP BY tag;

-- Look at the table you created
SELECT *
  FROM startdates;

-- 2)
-- To clear table if it already exists
DROP TABLE IF EXISTS startdates;

CREATE TEMP TABLE startdates AS
SELECT tag, min(date) AS mindate
  FROM stackoverflow
 GROUP BY tag;

-- Select tag (Remember the table name!) and mindate
SELECT startdates.tag,
       startdates.mindate,
       -- Select question count on the min and max days
       so_min.question_count AS min_date_question_count,
       so_max.question_count AS max_date_question_count,
       -- Compute the change in question_count (max - min)
       so_max.question_count - so_min.question_count AS change
  FROM startdates
       -- Join startdates to stackoverflow with alias so_min
       INNER JOIN stackoverflow AS so_min
          -- What needs to match between tables?
          ON startdates.tag = so_min.tag
         AND startdates.mindate = so_min.date
       -- Join to stackoverflow again with alias so_max
       INNER JOIN stackoverflow AS so_max
          -- Again, what needs to match between tables?
          ON startdates.tag = so_max.tag
         AND so_max.date = '2018-09-25';
```

## Insert into a temp table
```sql
-- 1)
DROP TABLE IF EXISTS correlations;

-- Create temp table
CREATE TEMP TABLE correlations AS
-- Select each correlation
SELECT 'profits'::varchar AS measure,
       -- Compute correlations
       corr(profits, profits) AS profits,
       corr(profits, profits_change) AS profits_change,
       corr(profits, revenues_change) AS revenues_change
FROM fortune500;

-- 2)
DROP TABLE IF EXISTS correlations;

CREATE TEMP TABLE correlations AS
SELECT 'profits'::varchar AS measure,
       corr(profits, profits) AS profits,
       corr(profits, profits_change) AS profits_change,
       corr(profits, revenues_change) AS revenues_change
FROM fortune500;

-- Add a row for profits_change
INSERT INTO correlations
SELECT 'profits_change'::varchar AS measure,
       corr(profits_change, profits) AS profits,
       corr(profits_change, profits_change) AS profits_change,
       corr(profits_change, revenues_change) AS revenues_change
FROM fortune500;

-- Repeat the above, but for revenues_change
INSERT INTO correlations
SELECT 'revenues_change'::varchar AS measure,
       corr(revenues_change, profits) AS profits,
       corr(revenues_change, profits_change) AS profits_change,
       corr(revenues_change, revenues_change) AS revenues_change
FROM fortune500;

-- 3)
DROP TABLE IF EXISTS correlations;

CREATE TEMP TABLE correlations AS
SELECT 'profits'::varchar AS measure,
       corr(profits, profits) AS profits,
       corr(profits, profits_change) AS profits_change,
       corr(profits, revenues_change) AS revenues_change
  FROM fortune500;

INSERT INTO correlations
SELECT 'profits_change'::varchar AS measure,
       corr(profits_change, profits) AS profits,
       corr(profits_change, profits_change) AS profits_change,
       corr(profits_change, revenues_change) AS revenues_change
  FROM fortune500;

INSERT INTO correlations
SELECT 'revenues_change'::varchar AS measure,
       corr(revenues_change, profits) AS profits,
       corr(revenues_change, profits_change) AS profits_change,
       corr(revenues_change, revenues_change) AS revenues_change
  FROM fortune500;

-- Select each column, rounding the correlations
SELECT measure,
       round(profits::numeric, 2) AS profits,
       round(profits_change::numeric, 2) AS profits_change,
       round(revenues_change::numeric, 2) AS revenues_change
FROM correlations;
```




# 3 Exploring Categorical Data and Unstructured Text
## Count the categories
```sql
-- 1)
-- Select the count of each level of priority
SELECT priority, COUNT(*)
  FROM evanston311
 GROUP BY priority;

-- 2)
-- Find values of zip that appear in at least 100 rows
-- Also get the count of each value
SELECT zip, COUNT(*)
  FROM evanston311
 GROUP BY zip
HAVING COUNT(*) >= 100;

-- 3)
-- Find values of source that appear in at least 100 rows
-- Also get the count of each value
SELECT source, COUNT(*)
  FROM evanston311
 GROUP BY source
HAVING COUNT(*) >= 100;

-- 4)
-- Find the 5 most common values of street and the count of each
SELECT street, COUNT(*)
  FROM evanston311
 GROUP BY street
 ORDER BY COUNT(*) DESC
 LIMIT 5;
```

## Spotting character data problems
```
Spotting character data problems
Explore the distinct values of the street column. Select each street value and the count of the number of rows with that value. Sort the results by street to see similar values near each other.

Which of the following is NOT an issue you see with the values of street?
[ ]The street suffix (e.g. Street, Avenue) is sometimes abbreviated
[x]There are sometimes extra spaces at the beginning and end of values
[ ]House/street numbers sometimes appear in the column
[ ]Capitalization is not consistent across values
[ ]All of the above are potential problems

---

1. Display unique street values.
SELECT street, COUNT(*)
FROM evanston311
GROUP BY street
ORDER BY street;

2. Look for data quality issues:
* Abbreviations (`St` vs `Street`)
* Inconsistent capitalization (`MAIN ST` vs `Main St`)
* House numbers mistakenly included in `street` (e.g. `12A DODGE AVE`, `8 1/2 ASBURY AVE`)
* Extra leading/trailing spaces

3. In this dataset:
* ✔ Abbreviations exist.
* ✔ Capitalization is inconsistent.
* ✔ Some house numbers are incorrectly stored in `street`.
* ✘ No leading/trailing spaces were found.
```

## Trimming
```sql
SELECT distinct street,
       -- Trim off unwanted characters from street
       TRIM(street, '0123456789#/. ') AS cleaned_street
  FROM evanston311
 ORDER BY street;
```

## Exploring unstructured text
```sql
-- 1)
-- Count rows
SELECT COUNT(*)
  FROM evanston311
 -- Where description includes trash or garbage
 WHERE description ILIKE '%trash%'
    OR description ILIKE '%garbage%';

-- 2)
-- Select categories containing Trash or Garbage
SELECT category
  FROM evanston311
 -- Use LIKE
 WHERE category LIKE '%Trash%'
    OR category LIKE '%Garbage%';

-- 3)
-- Count rows
SELECT COUNT(*)
  FROM evanston311
 -- description contains trash or garbage (any case)
 WHERE (description ILIKE '%trash%'
    OR description ILIKE '%garbage%')
 -- category does not contain Trash or Garbage
   AND category NOT LIKE '%Trash%'
   AND category NOT LIKE '%Garbage%';

-- 4)
-- Count rows with each category
SELECT category, COUNT(*)
  FROM evanston311
 WHERE (description ILIKE '%trash%'
    OR description ILIKE '%garbage%')
   AND category NOT LIKE '%Trash%'
   AND category NOT LIKE '%Garbage%'
 -- What are you counting?
 GROUP BY category
 ORDER BY count DESC
 LIMIT 10;
```

## Concatenate strings
```sql
-- Concatenate house_num, a space, and street and trim spaces from the start of the result
SELECT TRIM(LEADING FROM CONCAT(house_num, ' ', street)) AS address
FROM evanston311;
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
