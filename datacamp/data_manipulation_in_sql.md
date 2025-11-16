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
## Filtering using scalar subqueries
```sql
SELECT
    -- Select the date, home goals, and away goals scored
    date,
    home_goal,
    away_goal
FROM matches_2013_2014
-- Filter for matches where total goals exceeds 3x the average
WHERE (home_goal + away_goal) >
       (SELECT 3 * AVG(home_goal + away_goal)
        FROM matches_2013_2014);
```

## Filtering using a subquery with a list
```sql
SELECT
    -- Select the team long and short names
    team_long_name,
    team_short_name
FROM team
-- Exclude all values from the subquery
WHERE team_api_id NOT IN
     (SELECT DISTINCT awayteam_id FROM match);
```

## Filtering with more complex subquery conditions
```
SELECT
    -- Select the team long and short names
    team_long_name,
    team_short_name
FROM team
-- Filter for teams with 8 or more home goals
WHERE team_api_id IN
      (SELECT hometeam_id
       FROM match
       WHERE home_goal >= 8);
```

## Joining Subqueries in FROM
```sql
SELECT
    -- Select the country ID and match ID
    country_id,
    id
FROM match
-- Filter for matches with 10 or more goals in total
WHERE (home_goal + away_goal) >= 10;

SELECT
    -- Select country name and the count match IDs
    name AS country_name,
    COUNT(*) AS matches
FROM country AS c
-- Inner join the subquery onto country
-- Select the country id and match id columns
INNER JOIN (SELECT id, country_id
           FROM match
           -- Filter the subquery by matches with 10+ goals
           WHERE (home_goal + away_goal) >= 10) AS sub
ON c.id = sub.country_id
GROUP BY country_name;
```

## Building on Subqueries in FROM
```sql
SELECT
    -- Select country, date, home, and away goals from the subquery
    country,
    date,
    home_goal,
    away_goal
FROM
    -- Select country name, date, home_goal, away_goal, and total goals in the subquery
    (SELECT c.name AS country,
            m.date,
            m.home_goal,
            m.away_goal,
           (m.home_goal + m.away_goal) AS total_goals
    FROM match AS m
    LEFT JOIN country AS c
    ON m.country_id = c.id) AS subquery
-- Filter by total goals scored in the main query
WHERE total_goals >= 10;
```

## Add a subquery to the SELECT clause
```sql
SELECT
    l.name AS league,
    -- Round the average of the league's total goals
    ROUND(AVG(m.home_goal + m.away_goal), 2) AS avg_goals,
    -- Select and round the average total goals for the season
    (SELECT ROUND(AVG(home_goal + away_goal), 2)
     FROM match
     WHERE season = '2013/2014') AS overall_avg
FROM league AS l
LEFT JOIN match AS m
ON l.country_id = m.country_id
-- Filter for the 2013/2014 season
WHERE m.season = '2013/2014'
GROUP BY l.name;
```

## Subqueries in Select for Calculations
```sql
SELECT
    -- Select the league name and average goals scored
    l.name AS league,
    ROUND(AVG(m.home_goal + m.away_goal),2) AS avg_goals,
    -- Subtract the overall average from the league average
    ROUND(AVG(m.home_goal + m.away_goal) -
        (SELECT AVG(home_goal + away_goal)
         FROM match
         WHERE season = '2013/2014'),2) AS diff
FROM league AS l
LEFT JOIN match AS m
ON l.country_id = m.country_id
-- Only include 2013/2014 results
WHERE m.season = '2013/2014'
GROUP BY l.name;
```

## ALL the subqueries EVERYWHERE
```sql
SELECT
    -- Select the stage and average goals for each stage
    m.stage,
    ROUND(AVG(m.home_goal + m.away_goal),2) AS avg_goals,
    -- Select the average overall goals for the 2012/2013 season
    ROUND((SELECT AVG(home_goal + away_goal)
           FROM match
           WHERE season = '2012/2013'),2) AS overall
FROM match AS m
-- Filter for the 2012/2013 season
WHERE m.season = '2012/2013'
-- Group by stage
GROUP BY m.stage;
```

## Add a subquery in FROM
```sql
SELECT
    -- Select the stage and average goals from the subquery
    s.stage,
    ROUND(avg_goals,2) AS avg_goals
FROM
    -- Select the stage and average goals in 2012/2013
    (SELECT
         stage,
         AVG(home_goal + away_goal) AS avg_goals
     FROM match
     WHERE season = '2012/2013'
     GROUP BY stage) AS s
WHERE
    -- Filter the main query using the subquery
    s.avg_goals > (SELECT AVG(home_goal + away_goal)
                    FROM match WHERE season = '2012/2013');
```

## Add a subquery in SELECT
```sql
SELECT
    -- Select the stage and average goals from s
    s.stage,
    ROUND(s.avg_goals,2) AS avg_goal,
    -- Select the overall average for 2012/2013
    (SELECT AVG(home_goal + away_goal) FROM match WHERE season = '2012/2013') AS overall_avg
FROM
    -- Select the stage and average goals in 2012/2013 from match
    (SELECT
         stage,
         AVG(home_goal + away_goal) AS avg_goals
     FROM match
     WHERE season = '2012/2013'
     GROUP BY stage) AS s
WHERE
    -- Filter the main query using the subquery
    s.avg_goals > (SELECT AVG(home_goal + away_goal)
                    FROM match WHERE season = '2012/2013');
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
