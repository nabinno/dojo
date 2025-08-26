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

## Preparing dimensions
```sql
--Modify entity
ALTER TABLE IF EXISTS employees
RENAME TO dim_employees;

ALTER TABLE IF EXISTS departments
RENAME TO dim_departments;

ALTER TABLE IF EXISTS trainings
RENAME TO dim_trainings;
```

## Creating dimensions
```sql
-- Create new entity
CREATE OR REPLACE TABLE dim_date (
  	-- Add unique identifier
    date_id NUMBER(10,0) PRIMARY KEY,
  	-- Add new attributes to register date
    year NUMBER(4,0),
    month NUMBER(2,0)
);
```

## Retrieving data from dimensional model
```sql
SELECT
	dim_employees.*,
    dim_trainings.avg_training_score,
    dim_departments.department_name
FROM fact_employee_trainings
	JOIN dim_employees
    ON fact_employee_trainings.employee_id = dim_employees.employee_id
    JOIN dim_trainings
    ON fact_employee_trainings.training_id = dim_trainings.training_id
    JOIN dim_departments
    ON fact_employee_trainings.department_id = dim_departments.department_id
    -- Add dimension needed
    JOIN dim_date
    ON fact_employee_trainings.date_id = dim_date.date_id
WHERE dim_trainings.avg_training_score < 100
    -- Add extra filter
    AND dim_date.year = 2023;
```

## Creating hubs
```sql
-- Create a new hub entity
CREATE OR REPLACE TABLE hub_employee (
	-- Assign automated values to the hub key
	hub_employee_key NUMBER(10,0) AUTOINCREMENT PRIMARY KEY,
	employee_id NUMBER(38,0),
	-- Add attributes for historical tracking
	load_date TIMESTAMP,
	record_source VARCHAR(255)
);

CREATE OR REPLACE TABLE hub_department (
	-- Assign automated values to the hub key
	hub_department_id NUMBER(10,0) AUTOINCREMENT PRIMARY KEY,
  	-- Add hubs key reference
	department_id NUMBER(38,0),
	-- Add attributes for historical tracking
	load_date TIMESTAMP,
	record_source VARCHAR(255)
);

CREATE OR REPLACE TABLE hub_training (
	-- Add hub key
	hub_training_key NUMBER(10,0) AUTOINCREMENT PRIMARY KEY,
    -- Add the key attribute of trainings
	training_id NUMBER(38,0),
	-- Add history tracking attributes
	load_date TIMESTAMP,
	record_source VARCHAR(255)
);
```

## Creating satellites
```sql
-- Create a new satellite
CREATE OR REPLACE TABLE sat_employee (
	sat_employee_key NUMBER(10,0) AUTOINCREMENT PRIMARY KEY,
	hub_employee_key NUMBER(10,0),
   	employee_name VARCHAR(255),
    gender CHAR(1),
    age NUMBER(3,0),
	-- Add history tracking attributes
	load_date TIMESTAMP,
    record_source VARCHAR(255),
	-- Add a reference to foreign hub
    FOREIGN KEY (hub_employee_key) REFERENCES hub_employee(hub_employee_key)
);

CREATE OR REPLACE TABLE sat_department (
	-- Add the satellite's unique identifier
	sat_department_key NUMBER(10,0) AUTOINCREMENT PRIMARY KEY,
	-- Add the hub's key attribute
	hub_department_key NUMBER(10,0),
	department_name VARCHAR(255),
	region VARCHAR(255),
    -- Add history tracking attributes
    load_date TIMESTAMP,
    record_source VARCHAR(255),
	-- Add a reference to foreign hub
	FOREIGN KEY (hub_department_key) REFERENCES hub_department(hub_department_key)
);

CREATE OR REPLACE TABLE sat_training (
	sat_training_key NUMBER(10,0) AUTOINCREMENT PRIMARY KEY,
	-- Add the hub's key reference
	hub_training_key NUMBER(10,0),
	training_type VARCHAR(255),
    duration NUMBER(4,0),
    trainer_name VARCHAR(255),
    load_date TIMESTAMP,
    record_source VARCHAR(255),
	-- Add a reference to foreign hub
	FOREIGN KEY (hub_training_key) REFERENCES hub_training(hub_training_key)
);
```

## Creating links
```sql
CREATE OR REPLACE TABLE link_all (
	link_key NUMBER(10,0) AUTOINCREMENT PRIMARY KEY,
	hub_employee_key NUMBER(10,0),
  	-- Add the hub's key attributes
  	hub_training_key NUMBER(10,0),
  	hub_department_key NUMBER(10,0),
  	load_date TIMESTAMP,
    record_source VARCHAR(255),
	FOREIGN KEY (hub_employee_key) REFERENCES hub_employee(hub_employee_key),
  	-- Add a relationship with the foreign hubs
  	FOREIGN KEY (hub_training_key) REFERENCES hub_training(hub_training_key),
  	FOREIGN KEY (hub_department_key) REFERENCES hub_department(hub_department_key)
);
```

## Classifying data modeling techniques
```
Entity-relationship:
Daily routine business data

Dimensional:
Two major component: dimensions and facts.
Easy for analysis and reporting

Data vault:
Keeps long-term historical data tracking
Require managing multiple elements.
Components are hubs, links and satelites.
```

## Mastering data retrieval
```sql
SELECT
    hub_e.hub_employee_key,
    sat_d.department_name,
    -- Aggregate the attribute
    MAX(sat_t.avg_training_score) AS average_training
FROM hub_employee hub_e
	JOIN link_all AS li
    ON hub_e.hub_employee_key = li.hub_employee_key
    JOIN sat_department AS sat_d
    ON li.hub_department_key = sat_d.hub_department_key
    LEFT JOIN sat_training AS sat_t
    ON li.hub_training_key = sat_t.hub_training_key
WHERE sat_t.awards_won = 1
-- Group the results
GROUP BY hub_e.hub_employee_key, sat_d.department_name;
```




# 4 Snowflake Components
## Snowflake's storage method
```
Row-Based Storage:
Scans a list from top to bottom for details.
Time-consuming to search specific data if the dataset is large.
PostgreSQL is a databases system that use this storage method.
Optimized for analytical queries targeting specific data

Columnar Storage:
columns.
Fetches needed data quickly by going straight down the relevant column.
```

## Snowflake's advantages
```
[x]Snowflake uses micro-partitions to enable efficient large-volume processing.
[x]Snowflake's MPP system operates on a single server to process queries.
[ ]Snowflake accesses only necessary micro-partitions for queries using an organized index.
[x]Traditional systems like PostgreSQL may scan larger data sections even for small subsets.
```

## Virtual data warehouses
```
[x]Virtual warehouses in Snowflake allow for dynamic allocation of computing resources for data processing tasks.
[ ]In Snowflake, the data warehouse is a physical location where databases are stored.
[x]Snowflake's virtual warehouse is like having multiple dedicated rooms, each handling different tasks simultaneously without interference.
[ ]Virtual warehouses in Snowflake require manual intervention to scale resources based on the query load.
```

## Snowflake components
```
Tables:
- There are made up of rows and columns.

Schemas:
- It defines containers for database objects.
- They act like folders for better arganization and data retrieval.
- They can group multiple tables.

Databases:
- They can provides administrative boundaries far access control, resource allocation, and usage management.
```

## Implementing views
```sql
-- Create a view customer_financial_summary
CREATE OR REPLACE VIEW customer_financial_summary AS
SELECT c.customerid,
	c.estimatedsalary,
	cp.productid
FROM customers AS c
	-- Merge entity
	LEFT JOIN customerproducts AS cp
	ON c.customerid = cp.customerid;

CREATE OR REPLACE VIEW customer_financial_summary AS
SELECT c.customerid,
	-- Create a new conditional attribute
    CASE
        WHEN AVG(c.estimatedsalary) > 150000 THEN 'Top Income'
        WHEN AVG(c.estimatedsalary) > 90000 THEN 'High Income'
        WHEN AVG(c.estimatedsalary) > 20000 THEN 'Average Income'
        ELSE 'Low Income'
    END AS customer_category,
    -- Add aggregation to the attributes
    COUNT(DISTINCT cp.productid) AS product_count
FROM customers AS c
	LEFT JOIN customerproducts AS cp
	ON c.customerid = cp.customerid
-- Group the results
GROUP BY c.customerid;
```

## Order of execution
```
FROM
WHERE
GROUP BY
HAVING
ORDER BY
```

## Subquery mastery
```sql
WITH customer_status AS (
	SELECT c.customerid,
  		c.age,
        c.tenure,
        CASE
            WHEN ch.customerid IS NOT NULL THEN 'Churned'
            ELSE 'Active'
        END AS status
    FROM customers AS c
    	LEFT JOIN churn AS ch
  		ON c.customerid = ch.customerid
    GROUP BY c.customerid, c.age, c.tenure, status
)
SELECT status,
	COUNT(customerid) AS unique_customers,
    -- Calculate averages
    AVG(age) AS average_age,
    AVG(tenure) AS average_tenure
FROM customer_status
WHERE customerid IN (SELECT customerid
                     FROM customers
                     WHERE estimatedsalary > 175000)
GROUP BY status
-- Filter data
HAVING average_tenure > 2;
```
