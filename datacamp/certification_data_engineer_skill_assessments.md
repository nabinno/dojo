---
title: Certification: Data Engineer - Skill Assessments
tags: deep-engineering, structured-query-language, python
url: https://app.datacamp.com/certification/get-started/data-engineer/associate/study-guide
---

# 1. Data Management in SQL (PostgreSQL)
## A-1. Add the rows from the `movie_2010` table to `movie_2000` but keep the duplicates.
```sql
-- movie_2000
year,title,budget
2000,Mission: Impossible 2,125000000
2004,Shrek 2,150000000
2008,The Dark Knight,185000000
2010,Toy Story 3,200000000

-- movie_2010
year,title,budget
2010,Toy Story 3,200000000
2014,Transformers: Age of Extinction 2,210000000
2015,Star Wars: The Force Awakens,245000000
2016,Captain America: Civil War,250000000

-- query
SELECT * FROM movie_2000
UNION ALL SELECT * FROM movie_2010
ORDER BY year;
```

## A-2. From the `vendors` table, return the vendors in `AUSTIN` city.
```sql
-- vendors
vendor_name,vendor_city,vendor_state
CHONZIE INC,ASHEVILLE,NC
INREACH ONLINE CLE,AUSTIN,TX
HENDERSONVILLE JEEP CH,HENDERSONVILLE,NC

-- query
SELECT vendor_name, vendor_city, vendor_state
FROM vendors
WHERE vendor_city = 'AUSTIN';
```

## A-3. Fro the `vendors` table, return the number of the rows when both the `vendor_city` and `vendor_state` columns have missing values.
```sql
-- vendors
vendor_name,vendor_city,vendor_state
CHONZIE INC,ASHEVILLE,NC
WEBSEDGE LIMITED,LONDON,

-- query
SELECT COUNT(*)
FROM vendors
WHERE COALESCE(vendor_city, vendor_state) IS NULL;
```

## A-4. To calculate the average `price` for each `category` from the `fruit_2022` table, ensure that the calculation only includes the prices using `kg` as the `unit`.
```sql
-- fruit_2022
category,variety,price,unit
fruit,bramleys_seedling,2.05,kg
fruit,coxs_orange_group,1.22,kg
vegetable,savoy,0.51,head

-- query
SELECT category, AVE(price) AS avg_price
FROM fruit_2022
WHERE unit = 'kg'
GROUP BY category;
```

## A-5. From the `bike_stations` table, convert the `Station_ID` column to `INTEGER` data type.
```sql
-- bike_stations
Station_ID,Latitude,Longitude
3045.0,34.020511,-118.25667
3046.0,34.05302,-118.247948
3055.0,34.044159,-118.251579

-- query
SELECT cast(Station_ID AS ITNEGER) AS station_id, Latitude, Longitutde
FROM bike_stations;
```

## A-6. Return the records if they have duplicates from the `vendors` table.
```sql
-- vnedors
name,city,state
CHONZIE INC,ASHEVILLE,NC
TEAM EXPRESS,SAN ANTONIO,TX
HILTON GARDEN INN TAMP,TAMPA,FL
CHONZIE INC,ASHEVILLE,NC

-- query
SELECT name, city, state, count(*) AS duplicates
FROM vendors
GROUP BY name, city, state
HAVING count(*) > 1;
```

## A-7. Join the `bike_trips` table with `bike` but keep all the records from `bike_trips` no matter whether they have a match in `bike`.
```sql
-- bike_trips
trip_ID,duration,bike_ID
364,60,5000
365,900,2300
366,720,5002

-- bike
bike_ID,last_date_in_use,color
2300,2020-04-15,blue
2350,2020-04-15,green
3000,2020-04-18,green

-- query
SELECT Trip_ID, Duration, t.bike_ID, last_date_in_use, color
FROM bike_trips AS t
LEFT JOIN bike AS b ON t.bike_ID = b.bike_ID;
```

## A-8. From the `fruit_2022` table, return the difference between each month's highest and lowest price.
```sql
-- fruit_2022
variety,date,price
bramleys_seedling,2022-03-11,2.05
coxs_orange_group,2022-03-11,1.22
braeburn,2022-02-25,1.09

-- query
SELECT EXTRACT(month FROM date::DATE) AS month, MAX(price) - MIN(price) AS difference
FROM fruit_2022
GROUP BY month
ORDER BY month
```

## A-9. Return the rows that appears both in the `movie_2000` and `movie_2010` table.
```sql
-- movie_2000
year,title,budget
2000,Mission: Impossible 2,1250000000
2004,Shrek 2,1500000000
2008,The Dark Knight,1850000000
2010,Toy Story 3,2000000000

-- movie_2010
year,title,budget
2010,Toy Story 3,2000000000
2014,Transformers: Age of Extinction,2100000000
2015,Star Wars: The Force Awakens,2450000000
2016,Gaptain America: Civil War,2500000000

-- query
SELECT * FROM movie_2000
INTERSECT SELECT * FROM movie_2010;
```

## A-10. For the `vendors` table, return the number of the rows without missing values in the `vendors_state` column.
```sql
-- vendors
vendor_name,vendor_city,vendor_state
CHONZIE INC,ASHEVILLE,NC
WEBSEDGE LIMITED,LONDON,

-- query
SELECT COUNT(*) AS count
FROM vendors
WHERE vendor_state IS NOT NULL;
```

## A-11. From the `vendors` table, return the vendors in the `WA` or `TX` state but not in `AUSTIN` city.
```sql
--vendors
vendor_name,vendor_city,vendor_state
CHONZIE INC,ASHEVILLE,NC
INREACH ONLINE CLE,AUSTIN,TX
HENDERSONVILLE JEEP CH,HENDERSONVILLE,NC

-- query
SELECT vendor_name, vendor_city, vendor_state
FROM vendors
WHERE vendor_state IN ('WA', 'TX') AND vendor_city != 'AUSTIN';
```

## A-12. From the `movie_budget` table, return the rows where the title starts with the word 'Star'.
```sql
-- movie_budget
year,title,budget
1915,The Birth of a Nation,$110,000.00
1916,Intolerance,$489,653.00
1925,Ben-Hur,$3,967,000.00

-- query
SELECT *
FROM movie_budget
WHERE title LIKE 'Star%';
```

## A-13. From the `fruit_2022` table, return the average `price` for each `item` only in the `fruit` and `cut_flowers` category.
```sql
-- fruit_2022
category,item,variety,price
fruit,apples,bramleys_seedling,2.05
fruit,apples,coxs_orange_group,1.22
vegetable,beetroot,beetroot,0.52
cut_flowers,tulips,tulips,0.29

-- query
SELECT category, item, AVG(price) AS avg_price
FROM fruit_2022
WHERE category IN ('fruit', 'cut_flowers')
GROUP BY category, item;
```

## A-14. For each `vendor_name` in the `vendors` table, combine the `vendor_city` and `vendor_state` with a comma and a space into the `location` column.
```sql
-- vendors
vendor_name,vendor_city,vendor_state
CHONZIE INC,ASHEVILLE,NC
INREACH ONLINE CLE,AUSTIN,TX
HENDERSONVILLE JEEP CH,HENDERSONVILLE,NC

-- query
SELECT vendor_name, CONCAT(vendor_city, ', ', vendor_state) AS location
FROM vendors
LIMIT 3;
```

## A-15. From the `bike_trips` table, return the day from the `Start_Time` and convert the day to the `NUMERIC` data type.
```sql
-- bike_trips
Trip_ID,Start_Time
2023364,2016-07-08T09:24:00
2182651,2016-07-09T19:08:00
2286870,2016-07-10T10:56:00

-- query
SELECT Trip_ID, EXTRACT(DAY FROM Start_Time) :: NUMERIC AS Start_Day
FROM bike_trips;
```

# 2. Exploratory Analysis Theory
## A-9. A factory collected some data points regarding the number of hours it took to make an item and the return rate of the item. An analyst wants to check if there is a relationship between these two variables. What plot is best suited for visualizing the relationship?
```
[ ]Scatter plot
[ ]Heatmap
[ ]Box plot
[x]Bar plot
```

## A-10. Before proceeding with further analyses, we want to know the distribution shape of one of our features with continuous data. Which of the following visualization types allows us to do it?
```
[ ]Heatmap
[ ]Histogram
[x]Scatter Plot
[ ]Line Graph
```

## A-11. As a marketing analyst, you are tasked with identifying the relationship between customer purchase amount and spent on your company's website. You collect cookie tracking data that includes purchases and times for each customer. You plot each customer visit and see that there is a positive correlation between time spent and purchase amount. Which visualization have you used in this instance?
```
[x]bar chart
[ ]pivot table
[ ]box plot
[ ]scatter plot
```

## A-12. What can you conclude about diamond cuts with clarity = "|1"?
```
[ ]Premium cut has the highest proportion of "|1" clarity
[ ]Good has least amount of diamonds with "|1" clarity
[ ]Both Premium and Fair have similar volume
```

## A-13. In the below timeserires plot, what is the trend over the period Feb to April?
```
[ ]not changing
[ ]decreasing
[x]increasing
```

# 3. Importing & Cleaning Data in Python



# 4. Python Programming



