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
## Getting started with correlated subqueries
```sql
SELECT
    main.country_id,
    main.date,
    main.home_goal,
    main.away_goal
FROM match AS main
WHERE
    -- Filter the main query by the subquery
    (home_goal + away_goal) >
        (SELECT AVG((sub.home_goal + sub.away_goal) * 3)
         FROM match AS sub
         -- Join the main query to the subquery with country_id in WHERE
         WHERE main.country_id = sub.country_id);
```

## Correlated subquery with multiple conditions
```sql
SELECT
    main.country_id,
    main.date,
    main.home_goal,
    main.away_goal
FROM match AS main
WHERE
    -- Filter for matches with the maximum number of total goals scored
    (home_goal + away_goal) =
        (SELECT MAX(sub.home_goal + sub.away_goal)
         FROM match AS sub
         -- Join the main query to the subquery in WHERE
         WHERE main.country_id = sub.country_id
               AND main.season = sub.season);
```

## Getting started with nested subqueries
```sql
SELECT
    season,
    MAX(home_goal + away_goal) AS max_goals,
    (SELECT MAX(home_goal + away_goal)
     FROM match
     WHERE season = main.season
     -- Subquery to get the max goals in an 'England Premier League' match for the same season
     AND country_id IN (1729)
    ) AS pl_max_goals
FROM match AS main
GROUP BY season;
```

## Nest a subquery in FROM
```sql
-- Select matches where a team scored 5+ goals
SELECT
    country_id,
    season,
    id
FROM match
WHERE home_goal >= 5 OR away_goal >= 5;

-- Count match ids
SELECT
    country_id,
    season,
    COUNT(id) AS matches
-- Set up and alias the subquery
FROM (
    SELECT
        country_id,
        season,
        id
    FROM match
    WHERE home_goal >= 5 OR away_goal >= 5)
    AS subquery
GROUP BY country_id, season;

SELECT
    c.name AS country,
    -- Calculate the average matches per season
    AVG(outer_s.matches) AS avg_seasonal_high_scores
FROM country AS c
-- Left join outer_s to country
LEFT JOIN (
  SELECT country_id, season,
         COUNT(id) AS matches
  FROM (
    SELECT country_id, season, id
    FROM match
    WHERE home_goal >= 5 OR away_goal >= 5) AS inner_s
  -- Close parentheses and alias the subquery
  GROUP BY country_id, season) AS outer_s
ON c.id = outer_s.country_id
GROUP BY country;
```

## Clean up with CTEs
```sql
-- Set up your CTE
WITH match_list AS (
    SELECT
        country_id,
        id
    FROM match
    WHERE (home_goal + away_goal) >= 10)
-- Select league name and count of matches from the CTE
SELECT
    l.name AS league,
    COUNT(match_list.id) AS matches
FROM league AS l
-- Join the CTE to the league table using country_id
LEFT JOIN match_list ON l.id = match_list.country_id
GROUP BY l.name;
```

## Organizing with CTEs
```sql
-- Set up your CTE
WITH match_list AS (
  -- Select the league name, date, home, and away goals
    SELECT
        l.name AS league,
        m.date,
        m.home_goal,
        m.away_goal,
       (m.home_goal + m.away_goal) AS total_goals
    FROM match AS m
    LEFT JOIN league as l ON m.country_id = l.id)
-- Select the league, date, home, and away goals from the CTE
SELECT league, date, home_goal, away_goal
FROM match_list
-- Filter by total goals
WHERE total_goals >= 10;
```

## CTEs with nested subqueries
```sql
-- Set up your CTE
WITH match_list AS (
    SELECT
        country_id,
       (home_goal + away_goal) AS goals
    FROM match
    -- Create a list of match IDs to filter data in the CTE
    WHERE id IN (
       SELECT id
       FROM match
       WHERE season = '2013/2014' AND EXTRACT(MONTH FROM date) = 8))
-- Select the league name and average of goals in the CTE
SELECT
    l.name,
    AVG(match_list.goals)
FROM league AS l
-- Join the CTE onto the league table using country_id
LEFT JOIN match_list ON l.country_id = match_list.country_id
GROUP BY l.name;
```

## Get team names with a subquery
```sql
SELECT
    m.id,
    t.team_long_name AS hometeam
-- Left join match and team
FROM match AS m
LEFT JOIN team as t
ON m.hometeam_id = team_api_id;

SELECT
    m.date,
    -- Get the home team name and away team name from the subqueries
    home.hometeam_name,
    away.awayteam_name,
    m.home_goal,
    m.away_goal
FROM match AS m

-- Join the home subquery to the match table
LEFT JOIN (
  SELECT match.id, team.team_long_name AS hometeam_name
  FROM match
  LEFT JOIN team
  ON match.hometeam_id = team.team_api_id) AS home
ON home.id = m.id

-- Join the away subquery to the match table
LEFT JOIN (
  SELECT match.id, team.team_long_name AS awayteam_name
  FROM match
  LEFT JOIN team
  -- Get the away team ID in the subquery
  ON match.awayteam_id = team.team_api_id) AS away
ON away.id = m.id;
```

## Get team names with correlated subqueries
```sql
SELECT
    m.date,
   (SELECT team_long_name
    FROM team AS t
    -- Connect the team's team_api_id to the match's hometeam_id
    WHERE t.team_api_id = m.hometeam_id) AS hometeam
FROM match AS m;

SELECT
    m.date,
    (SELECT team_long_name
     FROM team AS t
     WHERE t.team_api_id = m.hometeam_id) AS hometeam,
    -- Connect the team to the match table
    (SELECT team_long_name
     FROM team AS t
     WHERE t.team_api_id = m.awayteam_id) AS awayteam,
    -- Select home_goal and away_goal
     m.home_goal,
     m.away_goal
FROM match AS m;
```

## Get team names with CTEs
```sql
SELECT
    -- Select match id and team long name
    m.id,
    t.team_long_name AS hometeam
FROM match AS m
-- Join team to match using team_api_id and hometeam_id
LEFT JOIN team AS t
ON t.team_api_id = m.hometeam_id;

-- Declare the home CTE
WITH home AS (
    SELECT m.id, t.team_long_name AS hometeam
    FROM match AS m
    LEFT JOIN team AS t
    ON m.hometeam_id = t.team_api_id)
-- Select everything from home
SELECT *
FROM home;

WITH home AS (
  SELECT m.id, m.date,
         t.team_long_name AS hometeam, m.home_goal
  FROM match AS m
  LEFT JOIN team AS t
  ON m.hometeam_id = t.team_api_id),
-- Declare and set up the away CTE
away AS (
  SELECT m.id, m.date,
         t.team_long_name AS awayteam, m.away_goal
  FROM match AS m
  LEFT JOIN team AS t
  ON m.awayteam_id = t.team_api_id)
-- Select date, home_goal, and away_goal
SELECT
    home.date,
    home.hometeam,
    away.awayteam,
    home.home_goal,
    away.away_goal
-- Join away and home on the id column
FROM home
INNER JOIN away
ON home.id = away.id;
```

## Which technique to use?
```sql
Which of the following statements is false regarding differences in the use and performance of multiple/nested subqueries, correlated subqueries, and common table expressions?

[ ]Correlated subqueries can allow you to circumvent multiple, complex joins.
[ ]Common table expressions are declared first, improving query run time.
[x]Correlated subqueries can reduce the length of your query, which improves query run time.
[ ]Multiple or nested subqueries are processed first, before your main query.
```




# 4 Window Functions
## The match is OVER
```sql
SELECT
    -- Select the match id, country name, season, home, and away goals
    m.id,
    c.name AS country,
    m.season,
    m.home_goal,
    m.away_goal,
    -- Use a window to include the aggregate average in each row
    AVG(m.home_goal + m.away_goal) OVER() AS overall_avg
FROM match AS m
LEFT JOIN country AS c ON m.country_id = c.id;
```

## What's OVER here?
```sql
SELECT
    -- Select the league name and average goals scored
    l.name AS league,
    AVG(m.home_goal + m.away_goal) AS avg_goals,
    -- Rank each league over the average goals
    RANK() OVER(ORDER BY AVG(m.home_goal + m.away_goal)) AS league_rank
FROM league AS l
LEFT JOIN match AS m
ON l.id = m.country_id
WHERE m.season = '2011/2012'
GROUP BY l.name
-- Order the query by the rank you created
ORDER BY avg_goals;
```

## Flip OVER your results
```sql
SELECT
    -- Select the league name and average goals scored
    l.name AS league,
    AVG(m.home_goal + m.away_goal) AS avg_goals,
    -- Rank leagues in descending order by average goals
    RANK() OVER(ORDER BY AVG(m.home_goal + m.away_goal) DESC) AS league_rank
FROM league AS l
LEFT JOIN match AS m
ON l.id = m.country_id
WHERE m.season = '2011/2012'
GROUP BY l.name
-- Order the query by the rank you created
ORDER BY league_rank;
```

## OVER with a PARTITION
```sql
SELECT
    date,
    season,
    home_goal,
    away_goal,
    CASE WHEN hometeam_id = 8673 THEN 'home'
         ELSE 'away' END AS warsaw_location,
    -- Calculate the average goals scored partitioned by season
    AVG(home_goal) OVER(PARTITION BY season) AS season_homeavg,
    AVG(away_goal) OVER(PARTITION BY season) AS season_awayavg
FROM match
-- Filter the data set for Legia Warszawa (id 8673) matches only
WHERE
    hometeam_id = 8673
    OR awayteam_id = 8673
ORDER BY (home_goal + away_goal) DESC;
```

## PARTITION BY multiple columns
```sql
SELECT
    date,
    season,
    home_goal,
    away_goal,
    CASE WHEN hometeam_id = 8673 THEN 'home'
         ELSE 'away' END AS warsaw_location,
    -- Calculate average goals partitioned by season and month
    AVG(home_goal) OVER(PARTITION BY season,
            EXTRACT(month FROM date)) AS season_mo_home,
    AVG(away_goal) OVER(PARTITION BY season,
            EXTRACT(month FROM date)) AS season_mo_away
FROM match
WHERE
    hometeam_id = 8673
    OR awayteam_id = 8673
ORDER BY (home_goal + away_goal) DESC;
```

## Slide to the left
```sql
SELECT
    date,
    home_goal,
    away_goal,
    -- Create a running total and running average of home goals
    SUM(home_goal) OVER(ORDER BY date
         ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW) AS running_total,
    AVG(home_goal) OVER(ORDER BY date
         ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW) AS running_avg
FROM match
WHERE
    hometeam_id = 9908
    AND season = '2011/2012';
```

## Slide to the right
```sql
SELECT
    -- Select the date and away goals
    date,
    away_goal,
    -- Create a running total sum and running average of away goals
    SUM(away_goal) OVER(ORDER BY date DESC
         ROWS BETWEEN CURRENT ROW AND UNBOUNDED FOLLOWING) AS running_total,
    AVG(away_goal) OVER(ORDER BY date DESC
         ROWS BETWEEN CURRENT ROW AND UNBOUNDED FOLLOWING) AS running_avg
FROM match
WHERE
    awayteam_id = 9908
    AND season = '2011/2012';
```

## Setting up the home team CTE
```sql
SELECT
    m.id,
    t.team_long_name,
    -- Identify matches as home/away wins or ties
    CASE WHEN m.home_goal > m.away_goal THEN 'MU Win'
        WHEN m.home_goal < m.away_goal THEN 'MU Loss'
        ELSE 'Tie' END AS outcome
FROM match AS m
-- Left join team on the hometeam_ID and team_API_id
LEFT JOIN team AS t
ON m.hometeam_id = t.team_api_id
WHERE
    -- Filter for 2014/2015 and Manchester United as the home team
    m.season = '2014/2015'
    AND t.team_long_name = 'Manchester United';
```

## Setting up the away team CTE
```sql
SELECT
    m.id,
    t.team_long_name,
    -- Identify matches as home/away wins or ties
    CASE WHEN m.home_goal > m.away_goal THEN 'MU Loss'
        WHEN m.home_goal < m.away_goal THEN 'MU Win'
        ELSE 'Tie' END AS outcome
-- Join the match table's awayteam_id to the team table
FROM match AS m
LEFT JOIN team AS t
ON m.awayteam_id = t.team_api_id
WHERE
    -- Filter for 2014/2015 and Manchester United as the away team
    m.season = '2014/2015'
    AND t.team_long_name = 'Manchester United';
```

## Putting the CTEs together
```

```

## Add a window function
```

```
