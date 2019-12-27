---
title: Exploring Google Ngrams with Amazon EMR
tags: amazon-emr,apache-hadoop
url: https://www.qwiklabs.com/focuses/346?parent=catalog
---

# Goal
- Create an Amazon EMR cluster running Hive and Ganglia
- Use Hive statements to create tables from Google Ngram input data stored in Amazon S3
- Run Hive queries to drill-down and analyze data
- Use Ganglia to monitor an AMR cluster

# Task
- [x] Overview
- [x] Start Lab
- [x] Task 1: Launch an Amazon EMR cluster
- [x] Task 2: Connect to Your Cluster
- [x] Task 3: Analyze Data
- [x] Achievements
- [x] End Lab
- [x] References

# Supplement
## Task 2: Connect to Your Cluster
```sh
ssh -i qwikLABS-L179-10866319.pem hadoop@ec2-3-231-225-54.compute-1.amazonaws.com
```

## Task 3: Analyze Data
### Execute Hive
```sh
hive
set hive.base.inputformat=org.apache.hadoop.hive.ql.io.HiveInputFormat;
set mapred.min.split.size=67108864;
set mapred.max.split.size=536870912;
```
Specify input data with using Hive

### Select the input data
```sql
set lang=eng-1M;
set ngram=1gram;

CREATE EXTERNAL TABLE IF NOT EXISTS ngrams (
  gram string,
  year int,
  occurrences bigint,
  pages bigint,
  books bigint
)
ROW FORMAT DELIMITED FIELDS TERMINATED BY '\t'
STORED AS SEQUENCEFILE LOCATION 's3://datasets.elasticmapreduce/ngrams/books/20090715/${hiveconf:lang}/${hiveconf:ngram}/';

DESCRIBE ngrams;
SELECT * FROM ngrams LIMIT 10;
```

### Standardise the input data
```sql
CREATE TABLE IF NOT EXISTS normalized (
  gram string,
  year int,
  occurrences bigint
);

set min_year=1990;
set max_year=2005;

INSERT OVERWRITE TABLE normalized
SELECT lower(gram), year, occurrences
FROM ngrams
WHERE
  year BETWEEN (${hiveconf:min_year} - 1) AND ${hiveconf:max_year} AND
  gram REGEXP "^[A-Za-z+\'-]{3,}$";

SELECT *
FROM normalized
LIMIT 20;

-- Display the 50 most-used words
SELECT
  gram,
  sum(occurrences) as total_occurrences
FROM normalized
GROUP BY gram
ORDER BY total_occurrences DESC
LIMIT 50;

-- Display the 50 most-used words longer than 10 characters
SELECT
  gram,
  sum(occurrences) as total_occurrences
FROM normalized
WHERE length(gram) > 10
GROUP BY gram
ORDER BY total_occurrences DESC
LIMIT 50;
```

### Monitoring with Ganglia
```
http://ec2-3-231-225-54.compute-1.amazonaws.com/ganglia
```

### Create a ratio table
```sql
CREATE TABLE IF NOT EXISTS ratios (
  gram string,
  year int,
  occurrences bigint,
  ratio double
);

INSERT OVERWRITE TABLE ratios
SELECT
  a.gram,
  a.year,
  sum(a.occurrences) AS occurrences,
  sum(a.occurrences) / b.total AS ratio
FROM normalized a
JOIN (
  SELECT year, sum(occurrences) AS total
  FROM normalized
  GROUP BY year
) b ON a.year = b.year
GROUP BY a.gram, a.year, b.total;

-- Words that most increased in popularity each year
SELECT 
  year,
  gram,
  occurrences,
  CONCAT(CAST(increase AS INT), 'x increase') AS increase
FROM (
  SELECT
    y2.gram,
    y2.year,
    y2.occurrences,
    y2.ratio / y1.ratio as increase,
    rank() OVER (PARTITION BY y2.year ORDER BY y2.ratio / y1.ratio DESC) AS rank
  FROM ratios y2
  JOIN ratios y1 ON y1.gram = y2.gram and y2.year = y1.year + 1
  WHERE
    y2.year BETWEEN 1991 and 2005 AND
    y1.occurrences > 1000 AND
    y2.occurrences > 1000
) grams
WHERE rank = 1
ORDER BY year;

-- Occurrences of 'internet' in books by year?
SELECT
  year,
  occurrences
FROM ratios
WHERE gram = 'internet'
ORDER BY year;

-- Most popular words of each length
SELECT DISTINCT length, gram
FROM
(
  SELECT length(gram) AS length,
  gram,
  rank() OVER (partition by length(gram) order by occurrences desc) AS rank
  FROM ratios
) x
WHERE rank = 1
ORDER BY length;
```

### Calculate the year-over-year ratio change
```sql
set outputbucket=s3n://[your bucket]/output;

CREATE EXTERNAL TABLE IF NOT EXISTS output_table (
  gram string,
  year int,
  ratio double,
  increase double
)
ROW FORMAT DELIMITED FIELDS TERMINATED BY '\t' 
STORED AS TEXTFILE LOCATION '${hiveconf:outputbucket}';

INSERT OVERWRITE TABLE output_table
SELECT a.gram as gram, a.year as year, a.ratio as ratio, a.ratio / b.ratio as increase
FROM ratios a
JOIN ratios b ON a.gram = b.gram AND a.year - 1 = b.year
WHERE a.ratio > 0.000001 AND
  a.year >= ${hiveconf:min_year} AND
  a.year <= ${hiveconf:max_year}
DISTRIBUTE BY year
SORT BY year ASC, increase DESC;

SELECT year, gram, increase
FROM output_table
WHERE year = 1977
LIMIT 100;
```

# References
## Amazon EMR
- https://aws.amazon.com/jp/blogs/big-data/
- https://docs.aws.amazon.com/emr/latest/ReleaseGuide/emr-release-5x.html

## Apache Hadoop
- https://cwiki.apache.org/confluence/display/HADOOP2/Home

## Apache Hive
- https://cwiki.apache.org/confluence/display/Hive/Home

## Ganglia
- http://ganglia.sourceforge.net/

