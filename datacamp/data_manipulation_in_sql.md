---
title: Data Manipulatiion in SQL
tags: structured-query-language
url: https://campus.datacamp.com/courses/data-manipulation-in-sql/well-take-the-case
---

# 1 We'll take the CASE
## Basic CASE statements
```sql
-- Identify the home team as Bayern Munich, Schalke 04, or neither
SELECT
    CASE WHEN hometeam_id = 10189 THEN 'FC Schalke 04'
        WHEN hometeam_id = 9823 THEN 'FC Bayern Munich'
         ELSE 'Other' END AS home_team,
    COUNT(id) AS total_matches
FROM matches_germany
GROUP BY home_team;
```

## CASE statements comparing column values
```sql
SELECT
    date,
    -- Identify home wins, losses, or ties
    CASE WHEN home_goal > away_goal THEN 'Home win!'
        WHEN home_goal < away_goal THEN 'Home loss :('
        ELSE 'Tie' END AS outcome
FROM matches_spain;
```

## CASE statements comparing two column values part 2
```sql
-- Select matches where Barcelona was the away team
SELECT
    m.date,
    t.team_long_name AS opponent,
    CASE WHEN m.home_goal < m.away_goal THEN 'Barcelona win!'
        WHEN m.home_goal > m.away_goal THEN 'Barcelona loss :('
        ELSE 'Tie' END AS outcome
FROM matches_spain AS m
LEFT JOIN teams_spain AS t
ON m.hometeam_id = t.team_api_id
-- Filter for Barcelona
WHERE m.awayteam_id = 8634;
```

## In CASE things get more complex
```

```

## In CASE of rivalry
```

```

## Filtering your CASE statement
```

```

## CASE WHEN with aggregate functions
```

```

## COUNT using CASE WHEN
```

```

## Filtering and totaling using CASE WHEN
```

```

## Calculating percent with CASE and AVG
```

```




# 2 Short and Simple Subqueries
## WHERE are the Subqueries?
```

```

## Filtering using scalar subqueries
```

```

## Filtering using a subquery with a list
```

```

## Filtering with more complex subquery conditions
```

```

## Subqueries in FROM
```

```

## Joining Subqueries in FROM
```

```

## Building on Subqueries in FROM
```

```

## Subqueries in SELECT
```

```

## Add a subquery to the SELECT clause
```

```

## Subqueries in Select for Calculations
```

```

## Subqueries everywhere! And best practices!
```

```

## ALL the subqueries EVERYWHERE
```

```

## Add a subquery in FROM
```

```

## Add a subquery in SELECT
```

```




# 3 Correlated Queries, Nested Queries, and Common Table Expressions
## Correlated subqueries
```

```

## Getting started with correlated subqueries
```

```

## Correlated subquery with multiple conditions
```

```

## Nested subqueries
```

```

## Getting started with nested subqueries
```

```

## Nest a subquery in FROM
```

```

## Common Table Expressions
```

```

## Clean up with CTEs
```

```

## Organizing with CTEs
```

```

## CTEs with nested subqueries
```

```

## Deciding on techniques to use
```

```

## Get team names with a subquery
```

```

## Get team names with correlated subqueries
```

```

## Get team names with CTEs
```

```

## Which technique to use?
```

```




# 4 Window Functions
## It's OVER
```

```

## The match is OVER
```

```

## What's OVER here?
```

```

## Flip OVER your results
```

```

## OVER with a PARTITION
```

```

## PARTITION BY a column
```

```

## PARTITION BY multiple columns
```

```

## Sliding windows
```

```

## Slide to the left
```

```

## Slide to the right
```

```

## Bringing it all together
```

```

## Setting up the home team CTE
```

```

## Setting up the away team CTE
```

```

## Putting the CTEs together
```

```

## Add a window function
```

```
