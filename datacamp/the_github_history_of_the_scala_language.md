---
title: The GitHub History of the Scala Language
tags: scala, github, statistics
url: https://app.datacamp.com/learn/projects/163
---

# 1. Scala's real-world project repository data
```python
# Importing pandas
import pandas as pd

# Loading in the data
pulls_one = pd.read_csv('datasets/pulls_2011-2013.csv')
pulls_two = pd.read_csv('datasets/pulls_2014-2018.csv')
pull_files = pd.read_csv('datasets/pull_files.csv')
```


# 2. Preparing and cleaning the data
```python
# Append pulls_one to pulls_two
pulls = pulls_one.append(pulls_two)

# Convert the date for the pulls object
pulls['date'] = pd.to_datetime(pulls['date'], utc=True)
```


# 3. Merging the DataFrames
```python
# Merge the two DataFrames
data = pd.merge(pulls, pull_files)
```


# 4. Is the project still actively maintained?
```python
%matplotlib inline

# Create a column that will store the month
data['month'] = data['date'].dt.month

# Create a column that will store the year
data['year'] = data['date'].dt.year

# Group by the month and year and count the pull requests
counts = data.groupby(['year', 'month'])['pid'].count()

# Plot the results
counts.plot(kind='bar', figsize = (12,4))
```


# 5. Is there camaraderie in the project?
```python
# Required for matplotlib
%matplotlib inline

# Group by the submitter
by_user = data.groupby('user').agg({'date': 'max'})

# Plot the histogram
by_user.hist()
```


# 6. What files were changed in the last ten pull requests?
```python

```


# 7. Who made the most pull requests to a given file?
``python`

```


# 8. Who made the last ten pull requests on a given file?
``python`

```


# 9. The pull requests of two special developers
``python`

```


# 10. Visualizing the contributions of each developer
``python`

```



