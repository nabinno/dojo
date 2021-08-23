---
title: Data Engineering for Everyone
tags: data-engineering
url: https://learn.datacamp.com/courses/data-engineering-for-everyone
---

# 1. What is data engineering?
## Go with the flow
```txt
1. Data collection and storage
2. Data preparation
3. Exploration and visualization
4. Experimentation and prediction
```

## Not responsible
```txt
## Data engineering tasks
Optimizing the customers databases for analysis.
Ensuring corrupted, unreadable music tracks are removed and don't end up facing customers.
Gathering music consumption data from desktop and mobile sources.

## Not data engineering tasks
Building a visualization to understand listening patterns by city.
Running an experiment to identify the optimal search bar positioning in the app.
Based on their listening behavior, predict which songs customers are likely to enjoy.
```

## Big time
```txt
# True
Data types refer to the variety of the data.
Value refers to how actionable the data is.

## False
Volume has to do with how trustworthy the data is.
Veracity refers to how frequently the data is generated.
Velocity refers to how big the data is.
```

## Who is it
```txt
## Data engineer
Provide listening sessions data so it can be analyzed with minimal preparation work.
Ensure that people who use the databases can't erase music videos by mistake.
Use Java to build a pipeline collecting album covers and storing them.

## Data scientist
Find out in which countries certain artists are popular to give them insights on where to tour.
Identify which customers are likely to end their Spotflix subscriptions, so marketing can target them and encourage them to renew.
Use Python to run an analysis and understand whether users prefer having the search bar on the top left or the top right of the Spotflix desktop app.
```

## Pipeline
```txt
1. Extract the songs Julian listened to the most over the past month.
2. Find other users who listened to these same songs a lot as well.
3. Load only the 10 top songs these users listened to the most over the past week into a table called "Similar profiles".
4. Extract only songs these other users listen to that are of the same genre as the ones in Julian's listening sessions. These are our recommendations.
5. Load the recommended songs into a new table. That's Julian's Weekly Playlist!
```


# 2. Storing data
## What's the difference
```python
## Structured
Corresponds to data in tabular format.
Is created and queried using SQL.
Is easy to search and organize.

## Semi-structured
Follows a model while allowing more flexibility than structured data.
Is moderately easy to easy and organize.
Is stored in XML or JSON format, or in NoSQL databases.

## Unstructured
Is difficult to search and organize.
Is usually stored in data lakes.
Stores images, pictures, videos and text.
```

## SQL databases
```python

```

## We can work it out
```python

```

## Columns
```python

```

## Different breeds
```python

```

## Data warehouses and data lakes
```python

```

## Tell the truth
```python

```

## Our warehouse (in the middle of our street)
```python

```




# 3. Moving and processing data
## Processing data
```python

```

## Connect the dots
```python

```

## Scheduling data
```python

```

## Schedules
```python

```

## One or the other
```python

```

## Parallel computing
```python

```

## Whenever, whenever
```python

```

## Parallel universe
```python

```

## Cloud computing
```python

```

## Obscured by clouds
```python

```

## Somewhere I belong
```python

```

## We are the champions
```python

```
