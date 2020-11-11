---
title: Statistical Thinking in Python (Part 1)
tags: statistics, python
url: https://www.datacamp.com/courses/statistical-thinking-in-python-part-1
---

# 1. Graphical exploratory data analysis
## Plotting a histogram of iris data
```python
# Import plotting modules
import matplotlib.pyplot as plt
import seaborn as sns

# Set default Seaborn style
sns.set()

# Plot histogram of versicolor petal lengths
plt.hist(versicolor_petal_length)

# Show histogram
plt.show()
```

## Axis labels!
```python
# Plot histogram of versicolor petal lengths
_ = plt.hist(versicolor_petal_length)

# Label axes
plt.xlabel('petal length (cm)')
plt.ylabel('count')

# Show histogram
plt.show()
```

## Adjusting the number of bins in a histogram
```python
# Import numpy
import numpy as np

# Compute number of data points: n_data
n_data = len(versicolor_petal_length)

# Number of bins is the square root of number of data points: n_bins
n_bins = np.sqrt(n_data)

# Convert number of bins to integer: n_bins
n_bins = int(n_bins)

# Plot the histogram
_ = plt.hist(versicolor_petal_length, bins=n_bins)

# Label axes
_ = plt.xlabel('petal length (cm)')
_ = plt.ylabel('count')

# Show histogram
plt.show()
```

## Bee swarm plot
```python
# Create bee swarm plot with Seaborn's default settings
sns.swarmplot(x='species', y='petal length (cm)', data=df)

# Label the axes
plt.xlabel('species')
plt.ylabel('petal length (cm)')

# Show the plot
plt.show()
```

## Computing the ECDF
```python
def ecdf(data):
    """Compute ECDF for a one-dimensional array of measurements."""
    # Number of data points: n
    n = len(data)

    # x-data for the ECDF: x
    x = np.sort(data)

    # y-data for the ECDF: y
    y = np.arange(1, n+1) / n

    return x, y
```

## Plotting the ECDF
```python
# Compute ECDF for versicolor data: x_vers, y_vers
x_vers, y_vers = ecdf(versicolor_petal_length)

# Generate plot
plt.plot(x_vers, y_vers, marker='.', linestyle='none')

# Label the axes
plt.xlabel('')
plt.ylabel('ECDF')

# Display the plot
plt.show()
```

## Comparison of ECDFs
```python
# Compute ECDFs
x_set, y_set = ecdf(setosa_petal_length)
x_vers, y_vers = ecdf(versicolor_petal_length)
x_virg, y_virg = ecdf(virginica_petal_length)

# Plot all ECDFs on the same plot
plt.plot(x_set, y_set)
plt.plot(x_vers, y_vers)
plt.plot(x_virg, y_virg)

# Annotate the plot
plt.legend(('setosa', 'versicolor', 'virginica'), loc='lower right')
_ = plt.xlabel('petal length (cm)')
_ = plt.ylabel('ECDF')

# Display the plot
plt.show()
```


# 2. Quantitative exploratory data analysis
## Introduction to summary statistics: The sample mean and median
```python

```

## Means and medians
```python

```

## Computing means
```python

```

## Percentiles, outliers, and box plots
```python

```

## Computing percentiles
```python

```

## Comparing percentiles to ECDF
```python

```

## Box-and-whisker plot
```python

```

## Variance and standard deviation
```python

```

## Computing the variance
```python

```

## The standard deviation and the variance
```python

```

## Covariance and the Pearson correlation coefficient
```python

```

## Scatter plots
```python

```

## Variance and covariance by looking
```python

```

## Computing the covariance
```python

```

## Computing the Pearson correlation coefficient
```python

```



# 3. Thinking probabilistically-- Discrete variables
## Probabilistic logic and statistical inference
```python

```

## What is the goal of statistical inference?
```python

```

## Why do we use the language of probability?
```python

```

## Random number generators and hacker statistics
```python

```

## Generating random numbers using the np.random module
```python

```

## The np.random module and Bernoulli trials
```python

```

## How many defaults might we expect?
```python

```

## Will the bank fail?
```python

```

## Probability distributions and stories: The Binomial distribution
```python

```

## Sampling out of the Binomial distribution
```python

```

## Plotting the Binomial PMF
```python

```

## Poisson processes and the Poisson distribution
```python

```

## Relationship between Binomial and Poisson distributions
```python

```

## How many no-hitters in a season?
```python

```

## Was 2015 anomalous?
```python

```



# 4. Thinking probabilistically-- Continuous variables
## Probability density functions
```python

```

## Interpreting PDFs
```python

```

## Interpreting CDFs
```python

```

## Introduction to the Normal distribution
```python

```

## The Normal PDF
```python

```

## The Normal CDF
```python

```

## The Normal distribution: Properties and warnings
```python

```

## Gauss and the 10 Deutschmark banknote
```python

```

## Are the Belmont Stakes results Normally distributed?
```python

```

## What are the chances of a horse matching or beating Secretariat's record?
```python

```

## The Exponential distribution
```python

```

## Matching a story and a distribution
```python

```

## Waiting for the next Secretariat
```python

```

## If you have a story, you can simulate it!
```python

```

## Distribution of no-hitters and cycles
```python

```

## Final thoughts
```python

```

