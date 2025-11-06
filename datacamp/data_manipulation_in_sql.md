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

## Filtering your CASE statement
```sql
SELECT
    season,
    date,
    home_goal,
    away_goal
FROM matches_italy
WHERE
    -- Find games where home_goal is more than away_goal
    CASE WHEN hometeam_id = 9857 AND home_goal > away_goal THEN 'Bologna Win'
        -- Find games where away_goal is more than home_goal
        WHEN awayteam_id = 9857 AND away_goal > home_goal THEN 'Bologna Win'
        -- Exclude games not won by Bologna
        END IS NOT NULL;
```

## COUNT using CASE WHEN
```
SELECT
    c.name AS country,
    -- Count matches in 2012/13
    COUNT(CASE WHEN m.season = '2012/2013' THEN m.id END) AS matches_2012_2013,
    -- Count matches in 2013/14
    COUNT(CASE WHEN m.season = '2013/2014' THEN m.id END) AS matches_2013_2014
FROM country AS c
LEFT JOIN match AS m
ON c.id = m.country_id
GROUP BY country;
```

## Filtering and totaling using CASE WHEN
```
SELECT season,
    -- SUM the home goals
    SUM(CASE WHEN hometeam_id = 8560 THEN home_goal END) AS home_goals,
    -- SUM the away goals
    SUM(CASE WHEN awayteam_id = 8560 THEN away_goal END) AS away_goals
FROM match
-- Group the results by season
GROUP BY season
```

## Calculating percent with CASE and AVG
```sql
SELECT
    c.name AS country,
    -- Calculate the percentage of tied games in each season
    AVG(CASE WHEN m.season= '2013/2014' AND m.home_goal = m.away_goal THEN 1
             WHEN m.season= '2013/2014' AND m.home_goal != m.away_goal THEN 0
             END) AS ties_2013_2014,
    AVG(CASE WHEN m.season= '2014/2015' AND m.home_goal = m.away_goal THEN 1
             WHEN m.season= '2014/2015' AND m.home_goal != m.away_goal THEN 0
             END) AS ties_2014_2015
FROM country AS c
LEFT JOIN matches AS m
ON c.id = m.country_id
GROUP BY country;
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
