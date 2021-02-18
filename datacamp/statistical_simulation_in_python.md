---
title: Statistical Simulation in Python
tags: statistics, python
url: https://www.datacamp.com/courses/statistical-simulation-in-python
---

# 1. Basics of randomness & simulation
## Poisson random variable
```python
# Initialize seed and parameters
np.random.seed(123) 
lam, size_1, size_2 = 5, 3, 1000  

# Draw samples & calculate absolute difference between lambda and sample mean
samples_1 = np.random.poisson(lam, size_1)
samples_2 = np.random.poisson(lam, size_2)
answer_1 = abs(lam - samples_1.mean())
answer_2 = abs(lam - samples_2.mean())

print("|Lambda - sample mean| with {} samples is {} and with {} samples is {}. ".format(size_1, answer_1, size_2, answer_2))
```

## Shuffling a deck of cards
```python

```

## Simulation basics
```python

```

## Throwing a fair die
```python

```

## Throwing two fair dice
```python

```

## Simulating the dice game
```python

```

## Using simulation for decision-making
```python

```

## Simulating one lottery drawing
```python

```

## Should we buy?
```python

```

## Calculating a break-even lottery price
```python

```



# 2. Probability & data generation process
## Probability basics
```python

```

## Queen and spade
```python

```

## Two of a kind
```python

```

## Game of thirteen
```python

```

## More probability concepts
```python

```

## The conditional urn
```python

```

## Birthday problem
```python

```

## Full house
```python

```

## Data generating process
```python

```

## Driving test
```python

```

## National elections
```python

```

## Fitness goals
```python

```

## eCommerce Ad Simulation
```python

```

## Sign up Flow
```python

```

## Purchase Flow
```python

```

## Probability of losing money
```python

```



# 3. Resampling methods
## Introduction to resampling methods
```python

```

## Sampling with replacement
```python

```

## Probability example
```python

```

## Bootstrapping
```python

```

## Running a simple bootstrap
```python

```

## Non-standard estimators
```python

```

## Bootstrapping regression
```python

```

## Jackknife resampling
```python

```

## Basic jackknife estimation - mean
```python

```

## Jackknife confidence interval for the median
```python

```

## Permutation testing
```python

```

## Generating a single permutation
```python

```

## Hypothesis testing - Difference of means
```python

```

## Hypothesis testing - Non-standard statistics
```python

```



# 4. Advanced Applications of Simulation
## Simulation for Business Planning
```python

```

## Modeling Corn Production
```python

```

## Modeling Profits
```python

```

## Optimizing Costs
```python

```

## Monte Carlo Integration
```python

```

## Integrating a Simple Function
```python

```

## Calculating the value of pi
```python

```

## Simulation for Power Analysis
```python

```

## Factors influencing Statistical Power
```python

```

## Power Analysis - Part I
```python

```

## Power Analysis - Part II
```python

```

## Applications in Finance
```python

```

## Portfolio Simulation - Part I
```python

```

## Portfolio Simulation - Part II
```python

```

## Portfolio Simulation - Part III
```python

```

## Wrap Up
```python

```

