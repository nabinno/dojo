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

## Different breeds
```python
## Data engineers
Modifying the whole songs table to remove trailing spaces entered by mistake in front of the title.
Creating a new table to store the songs customers listened to the most over the past year.
Updating an artist's table after they edited their biography.

## Data scientist
Querying the top songs of the past year to identify which genre dominated.
Querying the lyrics table to find all the songs that have 'data' in the title.
Querying the artist table to find all the bands that come from France.
```

## Our warehouse (in the middle of our street)
```python
## A data lake
Is optimized for cost efficiency.
Stores raw data.
Is mainly used by data scientist and engineers.
Can store structured, semi-structured and unstructured data.

## A data warehouse
Stores mainly structured data.
Usually stores smaller amounts of data than the other.
Is optimized for analysis.
Is mainly used by data analysts, business analysts, data scientists and machine learning engineers.
```



# 3. Moving and processing data
## Connect the dots
```python
## Extract
Pulling the top 20 songs users have been listening to on a loop.
Collecting data from Google Analytics about our web-markting promotion offering 3 months of access to the premium tier.

## Transform
Summarizing the yearly listening activity to tell users how many hours they've listened to music on Spotflix this year.
Sorting a playlist's songs based on the date they were added.

## Load
Saving the new order of a playlist that was sorted based on the date songs were added, so that it remains that way the next time the user connects.
Writing all the followers of a user in a table.
```

## Schedules
```python
## Manual
Running the pipeline processing sign ups because in the past 10 minutes, 100 new users complained to support that they can't connect.
Running the song encoding pipeline, because engineering changed the encoder and wants to make sure they still pass the validation check.

## Time
Collecting data from Google Analytics every morning to check how the promotion campaign is going.
Processing music videos uploaded by artists every hour.
Generating the Spotflix Weekly Playlist from Chapter 1 every Monday at 00:00 AM.

## Condition
Running validation checks if new videos are being collected.
Updating the number of followers in a playlist table after a user subscribed to it.
```

## One or the other
```python
## Batch
Loading new employees to Spotflix's employee table.
Reducing access to premium features when someone unsubscribes.

## Stream
Updating the count of followers in a playlist when a user subscribes to it.
When a user listens to songs that are being recommended in real time, loading his upvotes and downvotes on each song.
```

## Parallel universe
```python
## Right
Parallel computing is used to provide extra processing power.
It's a good idea to use parallel computing to encode songs uploaded by artists to the `.ogg` format that Spotflix prefers.
Parallel computing relies on processing units.

## Wrong
It's a good idea to use parallel computing to update the employees table every morning.
Parallel computing can't be used to optimize for memory usage.
Parallel computing will always make things faster.
```

## Obscured by clouds
```python
## Right
Cloud computing encompasses storage, database and computing solutions.
A multicloud solution reduces reliance on a single vendor.
Leveraging the cloud instead of having our own on-premises data center allows us to use just the resources we need, when we need them.

## Wrong
EC2, S3 and RDS are solutions offered by Microsoft Azure.
Multicloud solutions reduce security and governance concerns.
Cloud computing reduces all kinds of risk.
```

## Somewhere I belong
```python

```

## We are the champions
```python

```
