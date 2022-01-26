---
title: Cleaning Data in SQL Server Databases
tags: sql-server
url: https://campus.datacamp.com/courses/cleaning-data-in-sql-server-databases/starting-with-cleaning-data
---

# 1. Starting with Cleaning Data
## Unifying flight formats I
```sql
SELECT 
	-- Concat the strings
	CONCAT(
		carrier_code, 
		' - ', 
      	-- Replicate zeros
		REPLICATE('0', 9 - LEN(registration_code)), 
		registration_code, 
		', ', 
		airport_code)
	AS registration_code
FROM flight_statistics
-- Filter registers with more than 100 delays
WHERE delayed > 100
```

## Unifying flight formats II
```sql
SELECT 
    -- Concat the strings
	CONCAT(
		carrier_code, 
		' - ', 
        -- Format the code
		FORMAT(CAST(registration_code AS INT), '0000000'),
		', ', 
		airport_code
	) AS registration_code
FROM flight_statistics
-- Filter registers with more than 100 delays
WHERE delayed > 100
```

## Trimming strings I
```sql
SELECT 
	airport_code,
	-- Use the appropriate function to remove the extra spaces
    TRIM(airport_name) AS airport_name,
	airport_city,
    airport_state
-- Select the source table
FROM airports
```

## Trimming strings II
```sql
SELECT 
	airport_code,
	-- Use the appropriate function to remove the extra spaces
    RTRIM(LTRIM(airport_name)) AS airport_name,
	airport_city,
    airport_state
-- Select the source table
FROM airports
```

## Unifying strings
```sql
##
SELECT 
	airport_code,
	airport_name,
    -- Use the appropriate function to unify the values
    REPLACE(airport_city, 'ch', 'Chicago') AS airport_city,
	airport_state
FROM airports  
WHERE airport_code IN ('ORD', 'MDW')

##
SELECT airport_code, airport_name, 
	-- Use the CASE statement
	CASE
    	-- Unify the values
		WHEN airport_city <> 'Chicago' THEN REPLACE(airport_city, 'ch', 'Chicago')
		ELSE airport_city 
	END AS airport_city,
    airport_state
FROM airports
WHERE airport_code IN ('ORD', 'MDW')

##
SELECT 
	airport_code, airport_name,
    	-- Convert to uppercase
    	UPPER(
            -- Replace 'Chicago' with 'ch'.
          	REPLACE(airport_city, 'Chicago', 'ch')
        ) AS airport_city,
    airport_state
FROM airports
WHERE airport_code IN ('ORD', 'MDW')
```

## Comparing names with SOUNDEX()
```sql
SELECT 
    -- First name and surname of the statisticians
	DISTINCT S1.statistician_name, S1.statistician_surname
-- Join flight_statistics with itself
FROM flight_statistics S1 INNER JOIN flight_statistics S2 
	-- The SOUNDEX result of the first name and surname have to be the same
	ON SOUNDEX(S1.statistician_name) = SOUNDEX(S2.statistician_name) 
	AND SOUNDEX(S1.statistician_surname) = SOUNDEX(S2.statistician_surname) 
-- The texts of the first name or the texts of the surname have to be different
WHERE S1.statistician_name <> S2.statistician_name
	OR S1.statistician_surname <> S2.statistician_surname
```

## Comparing names with DIFFERENCE()
```sql
SELECT 
    -- First name and surnames of the statisticians
	DISTINCT S1.statistician_name, S1.statistician_surname
-- Join flight_statistics with itself
FROM flight_statistics S1 INNER JOIN flight_statistics S2 
	-- The DIFFERENCE of the first name and surname has to be equals to 4
	ON DIFFERENCE(S1.statistician_name, S2.statistician_name) = 4
	AND DIFFERENCE(S1.statistician_surname, S2.statistician_surname) = 4
-- The texts of the first name or the texts of the surname have to be different
WHERE S1.statistician_name <> S2.statistician_name
	OR S1.statistician_surname <> S2.statistician_surname
```



# 2. Dealing with missing data, duplicate data, and different date formats
## Removing missing values
```sql
##
SELECT *
-- Select the appropriate table
FROM airports
-- Exclude the rows where airport_city is NULL
WHERE airport_city IS NOT NULL

##
SELECT *
-- Select the appropriate table
FROM airports
-- Return only the rows where airport_city is NULL
WHERE airport_city IS NULL
```

## Removing blank spaces
```sql
##
SELECT *
-- Select the appropriate table
FROM airports
-- Exclude the rows where airport_city is missing
WHERE airport_city <> ''

##
SELECT *
-- Select the appropriate table
FROM airports
-- Return only the rows where airport_city is missing
WHERE airport_city = ''
```

## Filling missing values using ISNULL()
```sql
SELECT
  airport_code,
  airport_name,
  -- Replace missing values for airport_city with 'Unknown'
  ISNULL(airport_city, 'Unknown') AS airport_city,
  -- Replace missing values for airport_state with 'Unknown'
  ISNULL(airport_state, 'Unknown') AS airport_state
FROM airports
```

## Filling missing values using COALESCE()
```sql
SELECT
airport_code,
airport_name,
-- Replace the missing values
COALESCE(airport_city, airport_state, 'Unknown') AS location
FROM airports
```

## Treating duplicates
```sql
##
SELECT *,
	   -- Apply ROW_NUMBER()
       ROW_NUMBER() OVER (
         	-- Write the partition
            PARTITION BY 
                airport_code,
                carrier_code,
                registration_date
			ORDER BY 
                airport_code, 
                carrier_code, 
                registration_date
        ) row_num
FROM flight_statistics

##
-- Use the WITH clause
WITH cte AS (
    SELECT *, 
        ROW_NUMBER() OVER (
            PARTITION BY 
                airport_code, 
                carrier_code, 
                registration_date
			ORDER BY 
                airport_code, 
                carrier_code, 
                registration_date
        ) row_num
    FROM flight_statistics
)
SELECT * FROM cte
-- Get only duplicates
WHERE row_num > 1;

##
WITH cte AS (
    SELECT *, 
        ROW_NUMBER() OVER (
            PARTITION BY 
                airport_code, 
                carrier_code, 
                registration_date
			ORDER BY 
                airport_code, 
                carrier_code, 
                registration_date
        ) row_num
    FROM flight_statistics
)
SELECT * FROM cte
-- Exclude duplicates
WHERE row_num = 1;
```

## Dealing with different date formats
```sql

```

## Using CONVERT()
```sql

```

## Using FORMAT()
```sql

```

## CONVERT() vs FORMAT()
```sql

```




# 3. Dealing with out of range values, different data types, and pattern matching
## Out of range values and inaccurate data
```sql

```

## Out of range values or inaccurate data?
```sql

```

## Detecting out of range values
```sql

```

## Excluding out of range values
```sql

```

## Detecting and excluding inaccurate data
```sql

```

## Converting data with different types
```sql

```

## Using CAST() and CONVERT()
```sql

```

## The series with most episodes
```sql

```

## Pattern matching
```sql

```

## Characters to specify a patterns
```sql

```

## Matching urls
```sql

```

## Checking phone numbers
```sql

```




# 4. Combining, splitting, and transforming data
## Combining data of some columns into one column
```sql

```

## Combining cities and states using +
```sql

```

## Concatenating cities and states
```sql

```

## Working with DATEFROMPARTS()
```sql

```

## Splitting data of one column into more columns
```sql

```

## Using SUBSTRING() and CHARINDEX()
```sql

```

## Using RIGHT() , LEFT() and REVERSE()
```sql

```

## SUBSTRING() or CHARINDEX()?
```sql

```

## Transforming rows into columns and vice versa
```sql

```

## PIVOT or UNPIVOT?
```sql

```

## Turning rows into columns
```sql

```

## Turning columns into rows
```sql

```

## Congratulations!
```sql

```


