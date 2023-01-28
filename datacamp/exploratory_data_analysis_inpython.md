---
title: Exploratory Data Analysis in Python
tags: python,analytics
url: https://app.datacamp.com/learn/courses/exploratory-data-analysis-in-python
---

# 1. Read, clean, and validate
## Exploring the NSFG data
```python
# Display the number of rows and columns
nsfg.shape

# Display the names of the columns
nsfg.columns

# Select column birthwgt_oz1: ounces
ounces = nsfg['birthwgt_oz1']

# Print the first 5 elements of ounces
print(ounces.head())
```

## Clean a variable
```python
# Replace the value 8 with NaN
nsfg['nbrnaliv'].replace([8], np.nan, inplace=True)

# Print the values and their frequencies
print(nsfg['nbrnaliv'].value_counts())
```

## Compute a variable
```python
# Select the columns and divide by 100
agecon = nsfg['agecon'] / 100
agepreg = nsfg['agepreg'] / 100

# Compute the difference
preg_length = agepreg - agecon

# Compute summary statistics
print(preg_length.describe())
```

## Make a histogram
```python
# Plot the histogram
plt.hist(agecon, bins=20, histtype='step')

# Label the axes
plt.xlabel('Age at conception')
plt.ylabel('Number of pregnancies')

# Show the figure
plt.show()
```

## Compute birth weight
```python
# Create a Boolean Series for full-term babies
full_term = nsfg['prglngth'] >= 37

# Select the weights of full-term babies
full_term_weight = birth_weight[full_term]

# Compute the mean weight of full-term babies
print(full_term_weight.mean())
```

## Filter
```python

```






# 2. Distributions
## Probability mass functions
```python

```

## Make a PMF
```python

```

## Plot a PMF
```python

```

## Cumulative distribution functions
```python

```

## Make a CDF
```python

```

## Compute IQR
```python

```

## Plot a CDF
```python

```

## Comparing distributions
```python

```

## Distribution of education
```python

```

## Extract education levels
```python

```

## Plot income CDFs
```python

```

## Modeling distributions
```python

```

## Distribution of income
```python

```

## Comparing CDFs
```python

```

## Comparing PDFs
```python

```




# 3. Relationships
## Exploring relationships
```python

```

## PMF of age
```python

```

## Scatter plot
```python

```

## Jittering
```python

```

## Visualizing relationships
```python

```

## Height and weight
```python

```

## Distribution of income
```python

```

## Income and height
```python

```

## Correlation
```python

```

## Computing correlations
```python

```

## Interpreting correlations
```python

```

## Simple regression
```python

```

## Income and vegetables
```python

```

## Fit a line
```python

```




# 4. Multivariate Thinking
## Limits of simple regression
```python

```

## Regression and causation
```python

```

## Using StatsModels
```python

```

## Multiple regression
```python

```

## Plot income and education
```python

```

## Non-linear model of education
```python

```

## Visualizing regression results
```python

```

## Making predictions
```python

```

## Visualizing predictions
```python

```

## Logistic regression
```python

```

## Predicting a binary variable
```python

```

## Next steps
```python

```

