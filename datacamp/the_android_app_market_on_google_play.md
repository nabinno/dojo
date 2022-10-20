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
# List of characters to remove
chars_to_remove = ["+",",","$"]
# List of column names to clean
cols_to_clean = ["Installs","Price"]

# Loop for each column in cols_to_clean
for col in cols_to_clean:
    # Loop for each char in chars_to_remove
    for char in chars_to_remove:
        # Replace the character with an empty string
        apps[col] = apps[col].apply(lambda x: x.replace(char, ''))
        
# Print a summary of the apps dataframe
print(apps.info())
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


