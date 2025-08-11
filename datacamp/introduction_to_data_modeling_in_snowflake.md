---
title: Introduction to Data Modeling in Snowflake
tags: snowflake,analytics,data-modeling,data-engineering
url: https://campus.datacamp.com/courses/introduction-to-data-modeling-in-snowflake
---

# 1 Fundamentals of Data Modeling
## Components of data modeling
```
Entity:
- Data about an object, concept, or place

Attribute:
- They help distinguish one entity from another
- Property or characteristic

Relationship:
- Describes how two or more things one connected
- Shows the association between multiple entities
```

## Describing an entity
```sql
DESC TABLE ecommerceonlineretail;
```

## Implementing conceptual data model
```sql
-- Create a new products entity
CREATE OR REPLACE TABLE products(
	-- List the entity's attributes
	stockcode VARCHAR(255),
    description VARCHAR(255)
);

-- Create a new orders entity
CREATE OR REPLACE TABLE orders (
	-- List the invoice attributes
	invoiceno VARCHAR(10),
  	invoicedate TIMESTAMP_NTZ(9),
  	-- List the attributes related to price and quantity
  	unitprice NUMBER(10, 2),
  	quantity NUMBER(38, 0)
);
```

## Creating an entity
```sql
-- Create customers table
CREATE OR REPLACE TABLE customers (
  -- Define unique identifier
  customerid NUMBER(38) PRIMARY KEY,
  country VARCHAR(255)
);
```

## Building the physical data model
```sql
CREATE OR REPLACE TABLE orders (
  	invoiceno VARCHAR(10) PRIMARY KEY,
  	invoicedate TIMESTAMP_NTZ(9),
  	unitprice NUMBER(10,2),
  	quantity NUMBER(38,0),
  	customerid NUMBER(38,0),
  	stockcode VARCHAR(255),
  	-- Add foreign key refering to the foreign tables
	FOREIGN KEY (customerid) REFERENCES customers(customerid),
  	FOREIGN KEY (stockcode) REFERENCES products(stockcode)
);
```



# 2 Data Relationships and Normalization
## Altering an entity
```sql
-- Alter suppliers table
ALTER TABLE suppliers
-- Add new column
ADD COLUMN IF NOT EXISTS region VARCHAR(255);

-- Alter suppliers table
ALTER TABLE suppliers
-- Add the new column
ADD COLUMN IF NOT EXISTS contact VARCHAR(255);

-- Alter suppliers table
ALTER TABLE suppliers
-- Assign the unique identifier
ADD PRIMARY KEY (supplier_id);
```

## Adjusting the model
```sql
-- Create entity
CREATE OR REPLACE TABLE batchdetails (
	-- Add numerical attribute
	batch_id NUMBER(10,0),
	-- Add characters attributes
    batch_number VARCHAR(255),
    production_notes VARCHAR(255)
);

ALTER TABLE batchdetails ADD PRIMARY KEY (batch_id);

-- Modify the entity
ALTER TABLE productqualityrating
-- Add new column
ADD COLUMN IF NOT EXISTS batch_id NUMBER(10,0);
```

## Identifying Data Redundancy
```sql
SELECT manufacturer,
	company_location,
	COUNT(*) AS product_count
FROM productqualityrating
GROUP BY manufacturer,
	company_location
-- Add a filter for occurrence count greater than 1
HAVING product_count > 1;
```

## Spotting Anomalies
```sql
SELECT manufacturer,
	COUNT(DISTINCT cocoa_percent, ingredients) AS distinct_combinations
FROM productqualityrating
WHERE bar_name = 'Arriba'
    AND year_reviewed > 2006
GROUP BY manufacturer
-- Add the clause to filter
HAVING distinct_combinations > 1;
```

## Creating 1NF entities
```sql
-- Create a new entity
CREATE OR REPLACE TABLE ingredients (
	-- Add unique identifier
    ingredient_id NUMBER(10,0) PRIMARY KEY,
  	-- Add other attributes
    ingredient VARCHAR(255)
);

-- Create a new entity
CREATE OR REPLACE TABLE reviews (
	-- Add unique identifier
    review_id NUMBER(10,0) PRIMARY KEY,
  	-- Add other attributes
    review VARCHAR(255)
);
```

## Applying 1NF
```sql
-- Add command to insert data
INSERT INTO ingredients(ingredient_id, ingredient)
SELECT
	ROW_NUMBER() OVER (ORDER BY TRIM(f.value)),
	TRIM(f.value)
FROM productqualityrating,
LATERAL FLATTEN(INPUT => SPLIT(productqualityrating.ingredients, ';')) f
GROUP BY TRIM(f.value);

-- Modify script for review
INSERT INTO reviews(review_id, review)
SELECT
	ROW_NUMBER() OVER (ORDER BY TRIM(f.value)),
	TRIM(f.value)
FROM productqualityrating,
LATERAL FLATTEN(INPUT => SPLIT(productqualityrating.review, ';')) f
GROUP BY TRIM(f.value);
```

## 2NF and 3NF
```sql
-- Add new entity
CREATE OR REPLACE TABLE manufacturers (
  	-- Assign unique identifier
  	manufacturer_id NUMBER(10,0) PRIMARY KEY,
  	--Add other attributes
  	manufacturer VARCHAR(255),
  	company_location VARCHAR(255)
);

-- Add values to manufacturers
INSERT INTO manufacturers (manufacturer_id, manufacturer, company_location)
SELECT
	-- Generate a sequential number
	ROW_NUMBER() OVER (ORDER BY manufacturer, company_location),
	manufacturer,
	company_location
FROM productqualityrating
-- Aggregate data by the other attributes
GROUP BY manufacturer,
	company_location;
```

## Applying 3NF
```sql
-- Create entity
CREATE OR REPLACE TABLE locations (
	-- Add unique identifier
  	location_id NUMBER(10,0) PRIMARY KEY,
  	-- Add main attribute
  	location VARCHAR(255)
);

-- Populate entity from other entity's data
INSERT INTO locations (location_id, location)
SELECT
	-- Generate unique sequential number
	ROW_NUMBER() OVER (ORDER BY company_location),
    -- Select the main attribute
	company_location
FROM manufacturers
-- Aggregate data by main attribute
GROUP BY company_location;

-- Modify entity
ALTER TABLE manufacturers
-- Remove attribute
DROP COLUMN IF EXISTS company_location;
```




# 3 Data Modeling Techniques for Data Warehouse
## Creating entities for ER model
```sql
-- Create new entity
CREATE OR REPLACE TABLE employee_training_details (
  	-- Assign a unique identifier for the entity
	employee_training_id NUMBER(10,0) PRIMARY KEY,
  	-- Add new attribute
    year NUMBER(4,0),
  	-- Add new attributes to reference foreign entities
  	employee_id NUMBER(38,0),
    training_id NUMBER(38,0),
    FOREIGN KEY (employee_id) REFERENCES employees(employee_id),
    FOREIGN KEY (training_id) REFERENCES trainings(training_id)
);
```

## Retrieving data from ER model
```sql
SELECT
	employees.employee_id,
    trainings.avg_training_score
FROM employees
	JOIN trainings
	ON employees.employee_id = trainings.employee_id
    -- Merge new entity
    JOIN departments
    ON employees.department_id = departments.department_id
WHERE trainings.avg_training_score > 65
	-- Add extra filter
	AND departments.department_name = 'Operations'
LIMIT 50;
```

## Dimensional Modeling
```

```

## Preparing dimensions
```

```

## Creating dimensions
```

```

## Retrieving data from dimensional model
```

```

## Data Vault
```

```

## Creating hubs
```

```

## Creating satellites
```

```

## Creating links
```

```

## Choosing the Right Approach
```

```

## Classifying data modeling techniques
```

```

## Mastering data retrieval
```

```




# 4 Snowflake Components
## Query Performance in Snowflake
```

```

## Snowflake's storage method
```

```

## Snowflake's advantages
```

```

## Snowflake Data Objects
```

```

## Virtual data warehouses
```

```

## Snowflake components
```

```

## Implementing views
```

```

## Query Optimization
```

```

## Order of execution
```

```

## Subquery mastery
```

```

## Wrap-up
```

```

