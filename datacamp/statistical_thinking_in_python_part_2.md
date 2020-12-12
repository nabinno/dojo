---
title: Statistical Thinking in Python (Part 2)
tags: statistics, python
url: https://www.datacamp.com/courses/statistical-thinking-in-python-part-2
---

# 1. Parameter estimation by optimization
## How often do we get no-hitters?
```python
# Seed random number generator
np.random.seed(42)

# Compute mean no-hitter time: tau
tau = np.mean(nohitter_times)

# Draw out of an exponential distribution with parameter tau: inter_nohitter_time
inter_nohitter_time = np.random.exponential(tau, 100000)

# Plot the PDF and label axes
_ = plt.hist(inter_nohitter_time,
             bins=50, normed=True, histtype='step')
_ = plt.xlabel('Games between no-hitters')
_ = plt.ylabel('PDF')

# Show the plot
plt.show()
```

## Do the data follow our story?
```python
# Create an ECDF from real data: x, y
x, y = ecdf(nohitter_times)

# Create a CDF from theoretical samples: x_theor, y_theor
x_theor, y_theor = ecdf(inter_nohitter_time)

# Overlay the plots
plt.plot(x_theor, y_theor)
plt.plot(x, y, marker='.', linestyle='none')

# Margins and axis labels
plt.margins(0.02)
plt.xlabel('Games between no-hitters')
plt.ylabel('CDF')

# Show the plot
plt.show()
```

## How is this parameter optimal?
```python
# Plot the theoretical CDFs
plt.plot(x_theor, y_theor)
plt.plot(x, y, marker='.', linestyle='none')
plt.margins(0.02)
plt.xlabel('Games between no-hitters')
plt.ylabel('CDF')

# Take samples with half tau: samples_half
samples_half = np.random.exponential(tau / 2, 10000)

# Take samples with double tau: samples_double
samples_double = np.random.exponential(2 * tau, 10000)

# Generate CDFs from these samples
x_half, y_half = ecdf(samples_half)
x_double, y_double = ecdf(samples_double)

# Plot these CDFs as lines
_ = plt.plot(x_half, y_half)
_ = plt.plot(x_double, y_double)

# Show the plot
plt.show()
```

## EDA of literacy/fertility data
```python
# Plot the illiteracy rate versus fertility
_ = plt.plot(illiteracy, fertility, marker='.', linestyle='none')

# Set the margins and label axes
plt.margins(0.02)
_ = plt.xlabel('percent illiterate')
_ = plt.ylabel('fertility')

# Show the plot
plt.show()

# Show the Pearson correlation coefficient
print(pearson_r(illiteracy, fertility))
```

## Linear regression
```python
# Plot the illiteracy rate versus fertility
_ = plt.plot(illiteracy, fertility, marker='.', linestyle='none')
plt.margins(0.02)
_ = plt.xlabel('percent illiterate')
_ = plt.ylabel('fertility')

# Perform a linear regression using np.polyfit(): a, b
a, b = np.polyfit(illiteracy, fertility, 1)

# Print the results to the screen
print('slope =', a, 'children per woman / percent illiterate')
print('intercept =', b, 'children per woman')

# Make theoretical line to plot
x = np.array([0, 100])
y = a * x + b

# Add regression line to your plot
_ = plt.plot(x, y)

# Draw the plot
plt.show()
```

## How is it optimal?
```python
# Specify slopes to consider: a_vals
a_vals = np.linspace(0, 0.1, 200)

# Initialize sum of square of residuals: rss
rss = np.empty_like(a_vals)

# Compute sum of square of residuals for each value of a_vals
for i, a in enumerate(a_vals):
    rss[i] = np.sum((fertility - a*illiteracy - b)**2)

# Plot the RSS
plt.plot(a_vals, rss, '-')
plt.xlabel('slope (children per woman / percent illiterate)')
plt.ylabel('sum of square of residuals')

plt.show()
```

## Linear regression on appropriate Anscombe data
```python
# Perform linear regression: a, b
a, b = np.polyfit(x, y, 1)

# Print the slope and intercept
print(a, b)

# Generate theoretical x and y data: x_theor, y_theor
x_theor = np.array([3, 15])
y_theor = x_theor * a + b

# Plot the Anscombe data and theoretical line
_ = plt.plot(x, y, marker='.', linestyle='none')
_ = plt.plot(x_theor, y_theor, marker='.', linestyle='none')

# Label the axes
plt.xlabel('x')
plt.ylabel('y')

# Show the plot
plt.show()
```

## Linear regression on all Anscombe data
```python
# Iterate through x,y pairs
for x, y in zip(anscombe_x, anscombe_y):
    # Compute the slope and intercept: a, b
    a, b = np.polyfit(x, y, 1)

    # Print the result
    print('slope:', a, 'intercept:', b)
```


# 2. Bootstrap confidence intervals
## Visualizing bootstrap samples
```python
for _ in range(50):
    # Generate bootstrap sample: bs_sample
    bs_sample = np.random.choice(rainfall, size=len(rainfall))

    # Compute and plot ECDF from bootstrap sample
    x, y = ecdf(bs_sample)
    _ = plt.plot(x, y, marker='.', linestyle='none',
                 color='gray', alpha=0.1)

# Compute and plot ECDF from original data
x, y = ecdf(rainfall)
_ = plt.plot(x, y, marker='.')

# Make margins and label axes
plt.margins(0.02)
_ = plt.xlabel('yearly rainfall (mm)')
_ = plt.ylabel('ECDF')

# Show the plot
plt.show()
```

## Generating many bootstrap replicates
```python
def draw_bs_reps(data, func, size=1):
    """Draw bootstrap replicates."""

    # Initialize array of replicates: bs_replicates
    bs_replicates = np.empty(size)

    # Generate replicates
    for i in range(size):
        bs_replicates[i] = bootstrap_replicate_1d(data, func)

    return bs_replicates
```

## Bootstrap replicates of the mean and the SEM
```python
# Take 10,000 bootstrap replicates of the mean: bs_replicates
bs_replicates = draw_bs_reps(rainfall, np.mean, 10000)

# Compute and print SEM
sem = np.std(rainfall) / np.sqrt(len(rainfall))
print(sem)

# Compute and print standard deviation of bootstrap replicates
bs_std = np.std(bs_replicates)
print(bs_std)

# Make a histogram of the results
_ = plt.hist(bs_replicates, bins=50, normed=True)
_ = plt.xlabel('mean annual rainfall (mm)')
_ = plt.ylabel('PDF')

# Show the plot
plt.show()
```

## Confidence intervals of rainfall data
```python
np.percentile(bs_replicates, [2.5, 97.5])
```

## Bootstrap replicates of other statistics
```python

```

## Confidence interval on the rate of no-hitters
```python

```

## Pairs bootstrap
```python

```

## A function to do pairs bootstrap
```python

```

## Pairs bootstrap of literacy/fertility data
```python

```

## Plotting bootstrap regressions
```python

```




# 3. Introduction to hypothesis testing
## Formulating and simulating a hypothesis
```python

```

## Generating a permutation sample
```python

```

## Visualizing permutation sampling
```python

```

## Test statistics and p-values
```python

```

## Test statistics
```python

```

## What is a p-value?
```python

```

## Generating permutation replicates
```python

```

## Look before you leap: EDA before hypothesis testing
```python

```

## Permutation test on frog data
```python

```

## Bootstrap hypothesis tests
```python

```

## A one-sample bootstrap hypothesis test
```python

```

## A two-sample bootstrap hypothesis test for difference of means
```python

```




# 4. Hypothesis test examples
## A/B testing
```python

```

## The vote for the Civil Rights Act in 1964
```python

```

## What is equivalent?
```python

```

## A time-on-website analog
```python

```

## What should you have done first?
```python

```

## Test of correlation
```python

```

## Simulating a null hypothesis concerning correlation
```python

```

## Hypothesis test on Pearson correlation
```python

```

## Do neonicotinoid insecticides have unintended consequences?
```python

```

## Bootstrap hypothesis test on bee sperm counts
```python

```




# 5. Putting it all together: a case study
## Finch beaks and the need for statistics
```python

```

## EDA of beak depths of Darwin's finches
```python

```

## ECDFs of beak depths
```python

```

## Parameter estimates of beak depths
```python

```

## Hypothesis test: Are beaks deeper in 2012?
```python

```

## Variation in beak shapes
```python

```

## EDA of beak length and depth
```python

```

## Linear regressions
```python

```

## Displaying the linear regression results
```python

```

## Beak length to depth ratio
```python

```

## How different is the ratio?
```python

```

## Calculation of heritability
```python

```

## EDA of heritability
```python

```

## Correlation of offspring and parental data
```python

```

## Pearson correlation of offspring and parental data
```python

```

## Measuring heritability
```python

```

## Is beak depth heritable at all in G. scandens?
```python

```

## Final thoughts
```python

```


