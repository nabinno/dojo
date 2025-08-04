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

## Normalizing Relational Data
```

```

## Identifying Data Redundancy
```

```

## Spotting Anomalies
```

```

## The First Norm
```

```

## Creating 1NF entities
```

```

## Applying 1NF
```

```

## 2NF and 3NF
```

```

## Applying 2NF
```

```

## Applying 3NF
```

```




# 3 Data Modeling Techniques for Data Warehouse
## Entity–relationship model
```

```

## Creating entities for ER model
```

```

## Retrieving data from ER model
```

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

