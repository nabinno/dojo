---
title: Database Design
tags: database,database-design
url: https://www.datacamp.com/courses/database-design
---

# 1. Processing, Storing, and Organizing Data
## OLAP vs. OLTP
```
OLAP
- Queries a larger ammount of data
- Helps businesses with decision making and problem solving
- Typically uses a data warehouse

OLTP
- Most likely to have data from the past hour
- Data is inserted and updated more often
- Typically uses an operational database
```

## Which is better?
```sql
SELECT * FROM Potholes;
```

## Name that data type!
```
Unstructured
- To-do notes in a text editor
- Images in your photo library
- Zip file of all text messages ever received

Semi-Structured
- JSON object of tweets outputted in real-time by the Twitter API
- XML
- CSVs of open data downloaded from your local government websites

Structured
- A relational database with latest withdrawals and deposits made by clients
```

## Ordering ETL Tasks
```
- eCommerce API outputs real time data of transactions
- Python script drops null rows and clean data into pre-determined columns
- Resulting dataframe is written into an AWS Redshift Warehouse
```

## Classifying data models
```
Conceptual Data Model
- Entities, attributes, and relationships
- Gathers business requirements

Logical Data Model
- Relational model
- Determining tables and columns

Physical Data Model
- File structure of data storage
```

## Deciding fact and dimension tables
```
-- Create a route dimension table
CREATE TABLE route(
	route_id INTEGER PRIMARY KEY,
    park_name VARCHAR(160) NOT NULL,
    city_name VARCHAR(160) NOT NULL,
    distance_km FLOAT NOT NULL,
    route_name VARCHAR(160) NOT NULL
);
-- Create a week dimension table
CREATE TABLE week(
	week_id INTEGER PRIMARY KEY,
    week INTEGER NOT NULL,
    month VARCHAR(160) NOT NULL,
    year INTEGER NOT NULL
);
```

## Querying the dimensional model
```
SELECT 
	-- Get the total duration of all runs
	SUM(duration_mins)
FROM 
	runs_fact
-- Get all the week_id's that are from July, 2019
INNER JOIN week_dim ON week_dim.week_id = runs_fact.week_id
WHERE week_dim.month = 'July' and week_dim.year = '2019';
```



# 2. Database Schemas and Normalization
## Adding foreign keys
```
-- Add the book_id foreign key
ALTER TABLE fact_booksales ADD CONSTRAINT sales_book
    FOREIGN KEY (book_id) REFERENCES dim_book_star (book_id);
    
-- Add the time_id foreign key
ALTER TABLE fact_booksales ADD CONSTRAINT sales_time
    FOREIGN KEY (time_id) REFERENCES dim_time_star (time_id);
    
-- Add the store_id foreign key
ALTER TABLE fact_booksales ADD CONSTRAINT sales_store
    FOREIGN KEY (store_id) REFERENCES dim_store_star (store_id);
```

## Extending the book dimension
```
-- Create a new table for dim_author with an author column
CREATE TABLE dim_author (
    author varchar(256)  NOT NULL
);

-- Insert authors 
INSERT INTO dim_author
SELECT DISTINCT author FROM dim_book_star;

-- Add a primary key 
ALTER TABLE dim_author ADD COLUMN author_id SERIAL PRIMARY KEY;

-- Output the new table
SELECT * FROM dim_author;
```

## Querying the star schema
```
-- Output each state and their total sales_amount
SELECT dim_store_star.state, sum(fact_booksales.sales_amount)
FROM fact_booksales
	-- Join to get book information
    JOIN dim_book_star on dim_book_star.book_id = fact_booksales.book_id
	-- Join to get store information
    JOIN dim_store_star on dim_store_star.store_id = fact_booksales.store_id
-- Get all books with in the novel genre
WHERE  
    dim_book_star.genre = 'novel'
-- Group results by state
GROUP BY
    dim_store_star.state;
```

## Querying the snowflake schema
```
-- Output each state and their total sales_amount
SELECT dim_state_sf.state, sum(fact_booksales.sales_amount)
FROM fact_booksales
    -- Joins for genre
    JOIN dim_book_sf on dim_book_sf.book_id = fact_booksales.book_id
    JOIN dim_genre_sf on dim_genre_sf.genre_id = dim_book_sf.genre_id
    -- Joins for state 
    JOIN dim_store_sf on dim_store_sf.store_id = fact_booksales.store_id 
    JOIN dim_city_sf on dim_city_sf.city_id = dim_store_sf.city_id
	JOIN dim_state_sf on dim_state_sf.state_id = dim_city_sf.state_id
-- Get all books with in the novel genre and group the results by state
WHERE  
    dim_genre_sf.genre = 'novel'
GROUP BY
    dim_state_sf.state;
```

## Updating countries
```
-- Output records that need to be updated in the star schema
SELECT * FROM dim_store_star
WHERE dim_store_star.country != 'USA' AND dim_store_star.country !='CA';
```

## Extending the snowflake schema
```
-- Add a continent_id column with default value of 1
ALTER TABLE dim_country_sf
ADD continent_id int NOT NULL DEFAULT(1);

-- Add the foreign key constraint
ALTER TABLE dim_country_sf ADD CONSTRAINT country_continent
   FOREIGN KEY (continent_id) REFERENCES dim_continent_sf(continent_id);
   
-- Output updated table
SELECT * FROM dim_country_sf;
```

## Converting to 1NF
```
-- Create a new table to hold the cars rented by customers
CREATE TABLE cust_rentals (
  customer_id INT NOT NULL,
  car_id VARCHAR(128) NULL,
  invoice_id VARCHAR(128) NULL
);

-- Drop column from customers table to satisfy 1NF
ALTER TABLE customers
DROP COLUMN cars_rented,
DROP COLUMN invoice_id;
```

## Converting to 2NF
```
-- Create a new table to satisfy 2NF
CREATE TABLE cars (
  car_id VARCHAR(256) NULL,
  model VARCHAR(128),
  manufacturer VARCHAR(128),
  type_car VARCHAR(128),
  condition VARCHAR(128),
  color VARCHAR(128)
);

-- Drop columns in customer_rentals to satisfy 2NF
ALTER TABLE customer_rentals
DROP COLUMN model,
DROP COLUMN manufacturer,
DROP COLUMN type_car,
DROP COLUMN condition,
DROP COLUMN color;
```

## Converting to 3NF
```
-- Create a new table to satisfy 3NF
CREATE TABLE car_model(
  model VARCHAR(128),
  manufacturer VARCHAR(128),
  type_car VARCHAR(128)
);

-- Drop columns in rental_cars to satisfy 3NF
ALTER TABLE rental_cars
DROP COLUMN manufacturer,
DROP COLUMN type_car;
```



# 3. Database Views
## Tables vs. views
```
Only Tables
- Part of the physical schema of a database

Views & Tables
- Contains rows and columns
- Can be queried
- Has access control

Only Views
- Takes up less memory
- Always defined by a query
```

## Viewing views
```
-- Get all non-systems views
SELECT * FROM information_schema.views
WHERE table_schema NOT IN ('pg_catalog', 'information_schema');
```

## Creating and querying a view
```
-- Create a view for reviews with a score above 9
CREATE VIEW high_scores AS
SELECT * FROM reviews
WHERE score > 9;

-- Count the number of self-released works in high_scores
SELECT COUNT(*) FROM high_scores
INNER JOIN labels ON labels.reviewid = high_scores.reviewid
WHERE label = 'self-released';
```

## Creating a view from other views
```
##
-- Create a view with the top artists in 2017
CREATE VIEW top_artists_2017 AS
-- with only one column holding the artist field
SELECT artist_title.artist FROM artist_title
INNER JOIN top_15_2017
ON top_15_2017.reviewid = artist_title.reviewid;

-- Output the new view
SELECT * FROM top_artists_2017;

##
DROP VIEW top_15_2017 CASCADE;
```

## Granting and revoking access
```
-- Revoke everyone's update and insert privileges
REVOKE UPDATE, INSERT ON long_reviews FROM PUBLIC;

-- Grant the editor update and insert privileges 
GRANT UPDATE, INSERT ON long_reviews TO editor;
```

## Updatable views
```
##
SELECT * FROM information_schema.views WHERE table_schema = 'public' AND is_updatable = 'YES';
```

## Redefining a view
```
-- Redefine the artist_title view to have a label column
CREATE OR REPLACE VIEW artist_title AS
SELECT reviews.reviewid, reviews.title, artists.artist, labels.label
FROM reviews
INNER JOIN artists
ON artists.reviewid = reviews.reviewid
INNER JOIN labels
ON labels.reviewid = reviews.reviewid;

SELECT * FROM artist_title;
```

## Materialized versus non-materialized
```
Non-Materialized Views
- Always returns up-to-date data
- Better to use on write-intensive databases

Non-Materialized & Materialized Views
- Helps reduce the overhead of writing queries
- Can be used in a data warehouse

Materialized Views
- Consumes more storage
- Stores the query result on disk
```

## Creating and refreshing a materialized view
```
-- Create a materialized view called genre_count 
CREATE MATERIALIZED VIEW genre_count AS
SELECT genre, COUNT(*) 
FROM genres
GROUP BY genre;

INSERT INTO genres
VALUES (50000, 'classical');

-- Refresh genre_count
REFRESH MATERIALIZED VIEW genre_count;

SELECT * FROM genre_count;
```



# 4. Database Mangement
## Create a role
```
##
-- Create a data scientist role
CREATE ROLE data_scientist;

##
-- Create a role for Marta
CREATE ROLE marta LOGIN;

##
-- Create an admin role
CREATE ROLE admin WITH CREATEDB CREATEROLE;
```

## GRANT privileges and ALTER attributes
```
-- Grant data_scientist update and insert privileges
GRANT UPDATE, INSERT ON long_reviews TO data_scientist;

-- Give Marta's role a password
ALTER ROLE marta WITH PASSWORD 's3cur3p@ssw0rd';
```

## Add a user role to a group role
```
-- Add Marta to the data scientist group
GRANT data_scientist TO marta;

-- Celebrate! You hired data scientists.

-- Remove Marta from the data scientist group
REVOKE data_scientist FROM marta;
```

## Partitioning and normalization
```
Normalization
- Reduce redundancy in table
- Changes the logical data model

Vertical Partitioning
- Move specific columns to slower medium
- (Example) Move the third and fourth column to separate table

Horizontal Partitioning
- Sharding is an extension on this, using multiple machines
- (Example) Use the timestamp to move rows from Q4 in a specific table
```

## Creating vertical partitions
```
-- Create a new table called film_descriptions
CREATE TABLE film_descriptions (
    film_id INT,
    long_description TEXT
);

-- Copy the descriptions from the film table
INSERT INTO film_descriptions
SELECT film_id, long_description FROM film;
    
-- Drop the column in the original table
ALTER TABLE film DROP COLUMN long_description;

-- Join to create the original table
SELECT * FROM film 
JOIN film_descriptions USING(film_id);
```

## Creating horizontal partitions
```
-- Create a new table called film_partitioned
CREATE TABLE film_partitioned (
  film_id INT,
  title TEXT NOT NULL,
  release_year TEXT
)
PARTITION BY LIST (release_year);

-- Create the partitions for 2019, 2018, and 2017
CREATE TABLE film_2019
	PARTITION OF film_partitioned FOR VALUES IN ('2019');

CREATE TABLE film_2018
	PARTITION OF film_partitioned FOR VALUES IN ('2018');

CREATE TABLE film_2017
	PARTITION OF film_partitioned FOR VALUES IN ('2017');

-- Insert the data into film_partitioned
INSERT INTO film_partitioned
SELECT film_id, title, release_year FROM film;

-- View film_partitioned
SELECT * FROM film_partitioned;
```

## Data integration do's and dont's
```
## False
- Automated testing and proactive alerts are not needed
- You should choose whichever solution is right for the job right now
- All your data has to be updated in real time in the final view
- Your data integration solution, hand-coded or ETL tool, should work once and then you can use the resulting view to run queries forever
- Everybody should have access to sensitive data in the final view
- After data integration all your data should be in a single table

## True
- Being able to access the desired data through a single view does not mean all data is sotred together
- My source data can be in different formats and database management systems
- My source data can be stored in different physical locations
- Data integration should be business driven, e.g. what combination of data will be useful for the business
- Data in the final view can be updated in different intervals
- You should be careful choosing a hand-coded solution because of maintenance cost
```

## SQL versus NoSQL
```

```

## Choosing the right DBMS
```

```


