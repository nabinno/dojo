---
title: Introduction to Relational Databases in SQL
tags: sql,database
url: https://www.datacamp.com/courses/introduction-to-relational-databases-in-sql
---

# 1. Your first database
## Query information_schema with SELECT
```sql
-- Query the right table in information_schema
SELECT table_name 
FROM information_schema.tables
-- Specify the correct table_schema value
WHERE table_schema = 'public';

-- Query the right table in information_schema to get columns
SELECT column_name, data_type 
FROM information_schema.columns
WHERE table_name = 'university_professors' AND table_schema = 'public';

-- Query the first five rows of our table
SELECT * 
FROM university_professors 
LIMIT 5;
```

## CREATE your first few TABLEs
```sql
-- Create a table for the professors entity type
CREATE TABLE professors (
 firstname text,
 lastname text
);

-- Print the contents of this table
SELECT * 
FROM professors

-- Create a table for the universities entity type
CREATE TABLE universities (
  university_shortname text,
  university text,
  university_city text
);

-- Print the contents of this table
SELECT * 
FROM universities
```

## ADD a COLUMN with ALTER TABLE
```sql
-- Add the university_shortname column
ALTER TABLE professors
ADD COLUMN university_shortname text;

-- Print the contents of this table
SELECT * 
FROM professors
```

## RENAME and DROP COLUMNs n affiliations
```sql
-- Rename the organisation column
ALTER TABLE affiliations
RENAME COLUMN organisation TO organization;

-- Delete the university_shortname column
ALTER TABLE affiliations
DROP COLUMN university_shortname;
```

## Migrate data with INSERT INTO SELECT DISTINCT
```sql
-- Insert unique professors into the new table
INSERT INTO professors 
SELECT DISTINCT firstname, lastname, university_shortname 
FROM university_professors;

-- Doublecheck the contents of professors
SELECT * 
FROM professors;

-- Insert unique affiliations into the new table
INSERT INTO affiliations 
SELECT DISTINCT firstname, lastname, function, organization 
FROM university_professors;

-- Doublecheck the contents of affiliations
SELECT * 
FROM affiliations;
```

## Delete tables with DROP TABLE
```sql
-- Delete the university_professors table
DROP TABLE university_professors;
```

# 2. Enforce data consistency with attribute constraints
## Conforming with data types
```sql
-- Let's add a record to the table
INSERT INTO transactions (transaction_date, amount, fee) 
VALUES ('2018-09-24', 5454, '30');

-- Doublecheck the contents
SELECT *
FROM transactions;
```

## Type CASTs
```sql
-- Calculate the net amount as amount + fee
SELECT
  transaction_date,
  CAST(amount AS integer) + CAST(fee AS integer) AS net_amount
FROM transactions;
```

## Change types with ALTER COLUMN
```sql
-- Select the university_shortname column
SELECT DISTINCT(university_shortname) 
FROM professors;

-- Specify the correct fixed-length character type
ALTER TABLE professors
ALTER COLUMN university_shortname
TYPE char(3);

-- Change the type of firstname
ALTER TABLE professors
ALTER COLUMN firstname
TYPE varchar(64);
```

## Convert types USING a function
```sql
-- Convert the values in firstname to a max. of 16 characters
ALTER TABLE professors 
ALTER COLUMN firstname 
TYPE varchar(16)
USING SUBSTRING(firstname FROM 1 FOR 16)
```

## Disallow NULL values with SET NOT NULL
```sql
-- Disallow NULL values in firstname
ALTER TABLE professors 
ALTER COLUMN firstname SET NOT NULL;

-- Disallow NULL values in lastname
ALTER TABLE professors
ALTER COLUMN lastname SET NOT NULL;
```

# 3. Uniquely identify records with key constraints
## Make your columns UNIQUE with ADD CONSTRAINT
```sql
-- Make universities.university_shortname unique
ALTER TABLE universities
ADD CONSTRAINT university_shortname_unq UNIQUE(university_shortname);

-- Make organizations.organization unique
ALTER TABLE organizations
ADD CONSTRAINT organization_unq UNIQUE(organization);
```

## Get to know SELECT COUNT DISTINCT
```sql
-- Count the number of rows in universities
SELECT COUNT(DISTINCT(university_shortname, university, university_city))
FROM universities;

-- Count the number of distinct values in the university_city column
SELECT COUNT(DISTINCT(university_city)) 
FROM universities;
```

## Identify keys with SELECT COUNT DISTINCT
```sql
-- Try out different combinations
SELECT COUNT(DISTINCT(firstname, lastname, university_shortname)) 
FROM professors;

-- Try out different combinations
SELECT COUNT(DISTINCT(firstname, lastname)) 
FROM professors;
```

## ADD key CONSTRAINTs to the tables
```sql
-- Rename the organization column to id
ALTER TABLE organizations
RENAME COLUMN organization TO id;

-- Make id a primary key
ALTER TABLE organizations
ADD CONSTRAINT organization_pk PRIMARY KEY (id);

-- Rename the university_shortname column to id
ALTER TABLE universities
RENAME COLUMN university_shortname TO id;

-- Make id a primary key
ALTER TABLE universities
ADD CONSTRAINT university_pk PRIMARY KEY (id);
```

## Add a SERIAL surrogate key
```sql
-- Add the new column to the table
ALTER TABLE professors 
ADD COLUMN id serial;

-- Make id a primary key
ALTER TABLE professors 
ADD CONSTRAINT professors_pkey PRIMARY KEY (id);

-- Have a look at the first 10 rows of professors
SELECT *
FROM professors
LIMIT 10;
```

## CONCATenate columns to a surrogate key
```sql
-- Count the number of distinct rows with columns make, model
SELECT COUNT(DISTINCT(make, model)) 
FROM cars;

-- Add the id column
ALTER TABLE cars
ADD COLUMN id varchar(128);

-- Update id with make + model
UPDATE cars
SET id = CONCAT(make, model);

-- Make id a primary key
ALTER TABLE cars
ADD CONSTRAINT id_pk PRIMARY KEY(id);

-- Have a look at the table
SELECT * FROM cars;
```

## Test your knowledge before advancing
```sql
-- Create the table
CREATE TABLE students (
  last_name varchar(128) NOT NULL,
  ssn integer PRIMARY KEY,
  phone_no char(12)
);
```

# 4. Glue together tables with foreign keys
## REFRENCE a table with a FOREIGN KEY
```sql
-- Rename the university_shortname column
ALTER TABLE professors
RENAME COLUMN university_shortname TO university_id;

-- Add a foreign key on professors referencing universities
ALTER TABLE professors
ADD CONSTRAINT professors_fkey FOREIGN KEY (university_id) REFERENCES universities (id);
```

## Explore foreign key constraints
```sql
-- Try to insert a new professor
INSERT INTO professors (firstname, lastname, university_id)
VALUES ('Albert', 'Einstein', 'UZH');
```

## JOIN tables linked by a foreign key
```sql
-- Select all professors working for universities in the city of Zurich
SELECT professors.lastname, universities.id, universities.university_city
FROM professors
JOIN universities
ON professors.university_id = universities.id
WHERE universities.university_city = 'Zurich';
```

## Add foreign keys to the "affiliations" table
```sql
-- Add a professor_id column
ALTER TABLE affiliations
ADD COLUMN professor_id integer REFERENCES professors (id);

-- Rename the organization column to organization_id
ALTER TABLE affiliations
RENAME organization TO organization_id;

-- Add a foreign key on organization_id
ALTER TABLE affiliations
ADD CONSTRAINT affiliations_organization_fkey FOREIGN KEY (organization_id) REFERENCES organizations (id);
```

## Populate the "professor_id" column
```sql
-- Have a look at the 10 first rows of affiliations
SELECT *
FROM affiliations
LIMIT 10;

-- Update professor_id to professors.id where firstname, lastname correspond to rows in professors
UPDATE affiliations
SET professor_id = professors.id
FROM professors
WHERE affiliations.firstname = professors.firstname AND affiliations.lastname = professors.lastname;

-- Have a look at the 10 first rows of affiliations again
SELECT *
FROM affiliations
LIMIT 10;
```

## Drop "firstname" and "lastname"
```sql
-- Drop the firstname column
ALTER TABLE affiliations
DROP COLUMN firstname;

-- Drop the lastname column
ALTER TABLE affiliations
DROP COLUMN lastname;
```

## Change the referential integrity behavior of a key
```sql
-- Identify the correct constraint name
SELECT constraint_name, table_name, constraint_type
FROM information_schema.table_constraints
WHERE constraint_type = 'FOREIGN KEY';

-- Drop the right foreign key constraint
ALTER TABLE affiliations
DROP CONSTRAINT affiliations_organization_id_fkey;

-- Add a new foreign key constraint from affiliations to organizations which cascades deletion
ALTER TABLE affiliations
ADD CONSTRAINT affiliations_organization_id_fkey FOREIGN KEY (organization_id) REFERENCES organizations (id) ON DELETE CASCADE;

-- Delete an organization 
DELETE FROM organizations 
WHERE id = 'CUREM';

-- Check that no more affiliations with this organization exist
SELECT * FROM affiliations
WHERE organization_id = 'CUREM';
```

## Count affiliations per university
```sql
-- Count the total number of affiliations per university
SELECT COUNT(*), professors.university_id 
FROM affiliations
JOIN professors
ON affiliations.professor_id = professors.id
-- Group by the ids of professors
GROUP BY professors.university_id
ORDER BY count DESC;
```

## Join all the tables together
```sql
-- Join all tables
SELECT *
FROM affiliations
JOIN professors
ON affiliations.professor_id = professors.id
JOIN organizations
ON affiliations.organization_id = organizations.id
JOIN universities
ON professors.university_id = universities.id;

-- Group the table by organization sector, professor and university city
SELECT
  COUNT(*),
  organizations.organization_sector,
  professors.id,
  universities.university_city
FROM affiliations
JOIN professors
ON affiliations.professor_id = professors.id
JOIN organizations
ON affiliations.organization_id = organizations.id
JOIN universities
ON professors.university_id = universities.id
GROUP BY
  organizations.organization_sector, 
  professors.id,
  universities.university_city;

-- Filter the table and sort it
SELECT
  COUNT(*),
  organizations.organization_sector, 
  professors.id,
  universities.university_city
FROM affiliations
JOIN professors
ON affiliations.professor_id = professors.id
JOIN organizations
ON affiliations.organization_id = organizations.id
JOIN universities
ON professors.university_id = universities.id
WHERE organizations.organization_sector = 'Media & communication'
GROUP BY
  organizations.organization_sector, 
  professors.id,
  universities.university_city
ORDER BY count DESC;
```
