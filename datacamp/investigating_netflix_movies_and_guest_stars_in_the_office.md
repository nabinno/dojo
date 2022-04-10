---
title: Investigating Netflix Movies and Guest Stars in The Office
url: https://app.datacamp.com/learn/projects/entertainment-data/guided/Python
tags: python, data-engineer, analytics
---

# 1. Loading your friend's data into a dictionary
```python
# Create the years and durations lists
years = [2011, 2012, 2013, 2014, 2015, 2016, 2017, 2018, 2019, 2020]
durations = [103, 101, 99, 100, 100, 95, 95, 96, 93, 90]

# Create a dictionary with the two lists
movie_dict = {"years": years, "durations": durations}

# Print the dictionary
movie_dict
```


# 2. Creating a DataFrame from a dictionary
```python
# Import pandas under its usual alias
import pandas as pd

# Create a DataFrame from the dictionary
durations_df = pd.DataFrame(movie_dict)

# Print the DataFrame
durations_df
```


# 3. A visual inspection of our data
```python
# Import matplotlib.pyplot under its usual alias and create a figure
import matplotlib.pyplot as plt
fig = plt.figure()

# Draw a line plot of release_years and durations
plt.plot(durations_df['years'], durations_df['durations'])

# Create a title
plt.title("Netflix Movie Durations 2011-2020")

# Show the plot
plt.show()
```



# 4. Loading the rest of the data from a CSV
```python

```



# 5. Filtering for movies!
```python

```



# 6. Creating a scatter plot
```python

```



# 7. Digging deeper
```python

```



# 8. Marking non-feature films
```python

```



# 9. Plotting with color!
```python

```



# 10. What next?
```python

```

