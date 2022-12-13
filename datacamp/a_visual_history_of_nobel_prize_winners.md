---
title: A Visual History of Nobel Prize Winners
tags: analytics, python
url: https://app.datacamp.com/learn/projects/nobel-winners/guided/Python
---

# 1. The most Nobel of Prizes
```python
# Loading in required libraries
import pandas as pd
import seaborn as sns
import numpy as np

# Reading in the Nobel Prize data
nobel = pd.read_csv('datasets/nobel.csv')

# Taking a look at the first several winners
nobel.head(n=6)
```

# 2. So, who gets the Nobel Prize?
```python
# Display the number of (possibly shared) Nobel Prizes handed
# out between 1901 and 2016
display(len(nobel))

# Display the number of prizes won by male and female recipients.
display(nobel['sex'].value_counts())

# Display the number of prizes won by the top 10 nationalities.
nobel['birth_country'].value_counts().head(10)
```


# 3. USA dominance
```python

```


# 4. USA dominance, visualized
```python

```


# 5. What is the gender of a typical Nobel Prize winner?
```python

```


# 6. The first woman to win the Nobel Prize
```python

```


# 7. Repeat laureates
```python

```


# 8. How old are you when you get the prize?
```python

```


# 9. Age differences between prize categories
```python

```


# 10. Oldest and youngest winners
```python

```


# 11. You get a prize!
```python

```


