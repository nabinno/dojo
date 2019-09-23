---
title: "Working with Amazon Redshift"
tags: amazon-web-services, amazon-redshift
url: https://www.qwiklabs.com/focuses/332
---

# Goal
- Launch an Amazon Redshift cluster
- Connect to Amazon Redshift by using SQL client software
- Load data from Amazon S3 into Amazon Redshift
- Query data from Amazon Redshift
- Monitor Amazon Redshift performance

# Task
- [x] Launch your Amazon Redshift Cluster
- [x] Connect to Amazon Redshift
- [x] Load data
- [x] Run Queries
- [x] Joining tables
- [x] Analyze Performance
- [x] Explore the Amazon Redshift Console

# Supplement
## Load data
**Create table**
```sql
CREATE TABLE flights (
  year           smallint,
  month          smallint,
  day            smallint,
  carrier        varchar(80) DISTKEY,
  origin         char(3),
  dest           char(3),
  aircraft_code  char(3),
  miles          int,
  departures     int,
  minutes        int,
  seats          int,
  passengers     int,
  freight_pounds int
);
```

**Load data to table**
```sql
COPY flights
FROM 's3://us-west-2-aws-training/awsu-spl/spl17-redshift/static/data/flights-usa'
IAM_ROLE 'INSERT-YOUR-REDSHIFT-ROLE'
GZIP
DELIMITER ','
REMOVEQUOTES
REGION 'us-west-2';
```

## Run Queries
**Test table**
```sql
SELECT COUNT(*) FROM flights;
```

**Select from table**
```sql
SELECT *
FROM flights
ORDER BY random()
LIMIT 10;
```

**Execute query to data**
```sql
SELECT
  carrier,
  SUM(departures)
FROM flights
GROUP BY carrier
ORDER BY 2 DESC
LIMIT 10;
```

## Joining tables
**Load more data**
```sql
CREATE TABLE aircraft (
  aircraft_code CHAR(3) SORTKEY,
  aircraft      VARCHAR(50)
);
```

**Load more data**
```sql
COPY aircraft
FROM 's3://us-west-2-aws-training/awsu-spl/spl17-redshift/static/data/lookup_aircraft.csv'
IAM_ROLE 'INSERT-YOUR-REDSHIFT-ROLE'
IGNOREHEADER 1
DELIMITER ','
REMOVEQUOTES
TRUNCATECOLUMNS
REGION 'us-west-2';

SELECT *
FROM aircraft
ORDER BY random()
LIMIT 10;

SELECT
  aircraft,
  SUM(departures) AS trips
FROM flights
JOIN aircraft using (aircraft_code)
GROUP BY aircraft
ORDER BY trips DESC
LIMIT 10;
```

## Analyze Performance
**Explain query**
```sql
SET enable_result_cache_for_session TO OFF;

EXPLAIN
SELECT
  aircraft,
  SUM(departures) AS trips
FROM flights
JOIN aircraft using (aircraft_code)
GROUP BY aircraft
ORDER BY trips DESC
LIMIT 10;
```

**Compress data**
```sql
ANALYZE COMPRESSION flights;
```

**Create table from other table**
```sql
CREATE TABLE airports (
  airport_code CHAR(3) SORTKEY,
  airport      varchar(100)
);

COPY airports
FROM 's3://us-west-2-aws-training/awsu-spl/spl17-redshift/static/data/lookup_airports.csv'
IAM_ROLE 'INSERT-YOUR-REDSHIFT-ROLE'
IGNOREHEADER 1
DELIMITER ','
REMOVEQUOTES
TRUNCATECOLUMNS
REGION 'us-west-2';

CREATE TABLE vegas_flights
DISTKEY (origin)
SORTKEY (origin)
AS
SELECT flights.*, airport
FROM flights
JOIN airports ON origin = airport_code
WHERE dest = 'LAS';

SELECT
  airport,
  to_char(SUM(passengers), '999,999,999') as passengers
FROM vegas_flights
GROUP BY airport
ORDER BY SUM(passengers) desc
LIMIT 10;
```

**Confirm disk space and data distribution**
```sql
SELECT
  owner AS node,
  diskno,
  used,
  capacity,
  used/capacity::numeric * 100 as percent_used
FROM stv_partitions
WHERE host = node
ORDER BY 1, 2;

SELECT
  name,
  count(*)
FROM stv_blocklist
JOIN (SELECT DISTINCT name, id as tbl from stv_tbl_perm) USING (tbl)
GROUP BY name;
```
