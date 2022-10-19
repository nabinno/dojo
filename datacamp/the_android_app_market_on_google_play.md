---
title: The Android App Market on Google Play
tags: google-play, android, statistics
url: https://projects.datacamp.com/projects/619
---

# 1. Google Play Store apps and reviews
```python
# Read in dataset
import pandas as pd
apps_with_duplicates = pd.read_csv('datasets/apps.csv')

# Drop duplicates from apps_with_duplicates
apps = apps_with_duplicates.drop_duplicates()

# Print the total number of apps
print('Total number of apps in the dataset = ', apps_with_duplicates)

# Have a look at a random sample of 5 rows
print(apps)
```


# 2. Data cleaning
```python

```


# 3. Correcting data types
```python

```


# 4. Exploring app categories
```python

```


# 5. Distribution of app ratings
```python

```


# 6. Size and price of an app
```python

```


# 7. Relation between app category and app price
```python

```


# 8. Filter out "junk" apps
```python

```


# 9. Popularity of paid apps vs free apps
```python

```


# 10. Sentiment analysis of user reviews
```python

```


