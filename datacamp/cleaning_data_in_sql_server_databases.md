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

## Using CONVERT()
```sql
SELECT 
    airport_code,
    carrier_code,
    canceled,
    -- Convert the registration_date to a DATE and print it in mm/dd/yyyy format
    CONVERT(VARCHAR(10), CAST(registration_date AS DATE), 101) AS registration_date
FROM flight_statistics 
-- Convert the registration_date to mm/dd/yyyy format
WHERE CONVERT(VARCHAR(10), CAST(registration_date AS DATE), 101) 
	-- Filter the first six months of 2014 in mm/dd/yyyy format 
	BETWEEN '01/01/2014' AND '2014/06/30'
```

## Using FORMAT()
```sql
SELECT 
	pilot_code,
	pilot_name,
	pilot_surname,
	carrier_code,
    -- Convert the entry_date to a DATE and print it in dd/MM/yyyy format
	FORMAT(CAST(entry_date AS DATE), 'dd/MM/yyyy') AS entry_date
from pilots
```



# 3. Dealing with out of range values, different data types, and pattern matching
## Detecting out of range values
```sql
##
SELECT * FROM series
-- Detect the out of range values
WHERE num_ratings NOT BETWEEN 0 AND 5000

##
SELECT * FROM series
-- Detect the out of range values
WHERE num_ratings < 0 OR num_ratings > 5000
```

## Excluding out of range values
```sql
##
SELECT * FROM series
-- Exclude the out of range values
WHERE num_ratings BETWEEN 0 AND 5000

#
SELECT * FROM series
-- Exclude the out of range values
WHERE num_ratings >= 0 AND num_ratings <= 5000
```

## Detecting and excluding inaccurate data
```sql
##
SELECT * FROM series
-- Detect series for adults
WHERE is_adult = 1
-- Detect series with the minimum age smaller than 18
AND min_age < 18

##
SELECT * FROM series
-- Filter series for adults
WHERE is_adult = 1
-- Exclude series with the minimum age greater or equals to 18
AND min_age >= 18
```

## Using CAST() and CONVERT()
```sql
##
-- Use CAST() to convert the num_ratings column
SELECT AVG(CAST(num_ratings AS INT))
FROM series
-- Use CAST() to convert the num_ratings column
WHERE CAST(num_ratings AS INT) BETWEEN 0 AND 5000

##
-- Use CONVERT() to convert the num_ratings column
SELECT AVG(CONVERT(INT, num_ratings))
FROM series
-- Use CONVERT() to convert the num_ratings column
WHERE CONVERT(INT, num_ratings) BETWEEN 0 AND 5000
```

## The series with most episodes
```sql
SELECT s.name series_name, COUNT(e.name) episodes_count
FROM episodes e
INNER JOIN series s ON s.id = e.series_id
GROUP BY s.name
ORDER BY episodes_count DESC;
```

## Matching urls
```sql
SELECT 
	name,
    -- URL of the official site
	official_site
FROM series
-- Get the URLs that don't match the pattern
WHERE official_site NOT LIKE
	-- Write the pattern
	'www.%'
```

## Checking phone numbers
```sql
SELECT 
	name, 
    -- Contact number
    contact_number
FROM series
-- Get the numbers that don't match the pattern
WHERE contact_number NOT LIKE
	-- Write the pattern
	'555-___-____'
```




# 4. Combining, splitting, and transforming data
## Combining cities and states using +
```sql
##
SELECT 
	client_name,
	client_surname,
    -- Concatenate city with state
    city + ', ' + state AS city_state
FROM clients

##
SELECT 
	client_name,
	client_surname,
    -- Consider the NULL values
	ISNULL(city, '') + ISNULL(', ' + state, '') AS city_state
FROM clients
```

## Concatenating cities and states
```sql
SELECT 
		client_name,
		client_surname,
    -- Use the function to concatenate the city and the state
		CONCAT(
				city,
				CASE WHEN state IS NULL THEN '' 
				ELSE CONCAT(', ', state) END) AS city_state
FROM clients
```

## Working with DATEFROMPARTS()
```sql
##
SELECT 
	product_name,
	units,
    -- Use the function to concatenate the different parts of the date
	DATEFROMPARTS(
      	year_of_sale,
      	month_of_sale, 
      	day_of_sale) AS complete_date
FROM paper_shop_daily_sales
```

## Using SUBSTRING() and CHARINDEX()
```sql
SELECT 
	client_name,
	client_surname,
    -- Extract the name of the city
	SUBSTRING(city_state, 1, CHARINDEX(', ', city_state) - 1) AS city,
    -- Extract the name of the state
    SUBSTRING(city_state, CHARINDEX(', ', city_state) + 1, LEN(city_state)) AS state
FROM clients_split
```

## Using RIGHT() , LEFT() and REVERSE()
```sql
SELECT
	client_name,
	client_surname,
    -- Extract the name of the city
	LEFT(city_state, CHARINDEX(', ', city_state) - 1) AS city,
    -- Extract the name of the state
    RIGHT(city_state, CHARINDEX(' ,', REVERSE(city_state)) - 1) AS state
FROM clients_split
```

## Turning rows into columns
```sql
SELECT
	year_of_sale,
    -- Select the pivoted columns
	notebooks, 
	pencils, 
	crayons
FROM
   (SELECT 
		SUBSTRING(product_name_units, 1, charindex('-', product_name_units)-1) product_name, 
		CAST(SUBSTRING(product_name_units, charindex('-', product_name_units)+1, len(product_name_units)) AS INT) units,	
    	year_of_sale
	FROM paper_shop_monthly_sales) sales
-- Sum the units for column that contains the values that will be column headers
PIVOT (SUM(units) FOR product_name IN (notebooks, pencils, crayons))
-- Give the alias name
AS paper_shop_pivot
```

## Turning columns into rows
```sql

```

## Congratulations!
```sql

```


