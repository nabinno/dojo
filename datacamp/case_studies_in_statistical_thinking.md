---
title: Case Studies in Statistical Thinking
tags: statistics, python
url: https://www.datacamp.com/courses/case-studies-in-statistical-thinking
---

# 1. Fish sleep and bacteria growth: A review of Statistical Thinking I and II
## EDA: Plot ECDFs of active bout length
```python
# Import the dc_stat_think module as dcst
import dc_stat_think as dcst

# Generate x and y values for plotting ECDFs
x_wt, y_wt = dcst.ecdf(bout_lengths_wt)
x_mut, y_mut = dcst.ecdf(bout_lengths_mut)

# Plot the ECDFs
_ = plt.plot(x_wt, y_wt, marker='.', linestyle='none')
_ = plt.plot(x_mut, y_mut, marker='.', linestyle='none')

# Make a legend, label axes, and show plot
_ = plt.legend(('wt', 'mut'))
_ = plt.xlabel('active bout length (min)')
_ = plt.ylabel('ECDF')
plt.show()
```

## Interpreting ECDFs and the story
```python

```

## Bootstrap confidence intervals
```python

```

## Parameter estimation: active bout length
```python

```

## Permutation and bootstrap hypothesis tests
```python

```

## Permutation test: wild type versus heterozygote
```python

```

## Bootstrap hypothesis test
```python

```

## Linear regressions and pairs bootstrap
```python

```

## Assessing the growth rate
```python

```

## Plotting the growth curve
```python

```



# 2. Analysis of results of the 2015 FINA World Swimming Championships
## Introduction to swimming data
```python

```

## Graphical EDA of men's 200 free heats
```python

```

## 200 m free time with confidence interval
```python

```

## Do swimmers go faster in the finals?
```python

```

## EDA: finals versus semifinals
```python

```

## Parameter estimates of difference between finals and semifinals
```python

```

## How to do the permutation test
```python

```

## Generating permutation samples
```python

```

## Hypothesis test: Do women swim the same way in semis and finals?
```python

```

## How does the performance of swimmers decline over long events?
```python

```

## EDA: Plot all your data
```python

```

## Linear regression of average split time
```python

```

## Hypothesis test: are they slowing down?
```python

```




# 3. The "Current Controversy" of the 2013 World Championships
## Introduction to the current controversy
```python

```

## A metric for improvement
```python

```

## ECDF of improvement from low to high lanes
```python

```

## Estimation of mean improvement
```python

```

## How should we test the hypothesis?
```python

```

## Hypothesis test: Does lane assignment affect performance?
```python

```

## Did the 2015 event have this problem?
```python

```

## The zigzag effect
```python

```

## Which splits should we consider?
```python

```

## EDA: mean differences between odd and even splits
```python

```

## How does the current effect depend on lane position?
```python

```

## Hypothesis test: can this be by chance?
```python

```

## Recap of swimming analysis
```python

```




# 4. Statistical seismology and the Parkfield region
## Introduction to statistical seismology and the Parkfield experiment
```python

```

## Parkfield earthquake magnitudes
```python

```

## Computing the b-value
```python

```

## The b-value for Parkfield
```python

```

## Timing of major earthquakes and the Parkfield sequence
```python

```

## Interearthquake time estimates for Parkfield
```python

```

## When will the next big Parkfield quake be?
```python

```

## How are the Parkfield interearthquake times distributed?
```python

```

## Computing the value of a formal ECDF
```python

```

## Computing the K-S statistic
```python

```

## Drawing K-S replicates
```python

```

## The K-S test for Exponentiality
```python

```



# 5. Earthquakes and oil mining in Oklahoma
## Variations in earthquake frequency and seismicity
```python

```

## EDA: Plotting earthquakes over time
```python

```

## Estimates of the mean interearthquake times
```python

```

## Hypothesis test: did earthquake frequency change?
```python

```

## How to display your analysis
```python

```

## Earthquake magnitudes in Oklahoma
```python

```

## EDA: Comparing magnitudes before and after 2010
```python

```

## Quantification of the b-values
```python

```

## How should we do a hypothesis test on differences of the b-value?
```python

```

## Hypothesis test: are the b-values different?
```python

```

## What can you conclude from this analysis?
```python

```

## Closing comments
```python

```
