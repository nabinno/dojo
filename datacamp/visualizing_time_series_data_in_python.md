---
title: Visualizing Time Series Data in Python
tags: python,time-series,data-visualization
url: https://www.datacamp.com/courses/visualizing-time-series-data-in-python
---

# 1. Introduction
## Load your time series data
```python
# Import pandas
import pandas as pd

# Read in the file content in a DataFrame called discoveries
discoveries = pd.read_csv(url_discoveries)

# Display the first five lines of the DataFrame
print(discoveries.head())
```

## Test whether your data is of the correct type
```python
# Print the data type of each column in discoveries
print(discoveries.dtypes)

# Convert the date column to a datestamp type
discoveries['date'] = pd.to_datetime(discoveries['date'])

# Print the data type of each column in discoveries, again
print(discoveries.dtypes)
```

## Your first plot!
```python

```

## Specify plot styles
```python

```

## Display and label plots
```python

```

## Customize your time series plot
```python

```

## Subset time series data
```python

```

## Add vertical and horizontal markers
```python

```

## Add shaded regions to your plot
```python

```

# 2. Summary Statistics and Diagnostics
## Clean your time series data
```python

```

## Find missing values
```python

```

## Handle missing values
```python

```

## Plot aggregates of your data
```python

```

## Display rolling averages
```python

```

## Display aggregated values
```python

```

## Summarize the values in your time series data
```python

```

## Compute numerical summaries
```python

```

## Boxplots and Histograms
```python

```

## Density plots
```python

```

# 3. Seasonality, Trend and Noise
## Autocorrelation and Partial autocorrelation
```python

```

## Autocorrelation in time series data
```python

```

## Interpret autocorrelation plots
```python

```

## Partial autocorrelation in time series data
```python

```

## Interpret partial autocorrelation plots
```python

```

## Seasonality, trend and noise in time series data
```python

```

## Time series decomposition
```python

```

## Plot individual components
```python

```

## A quick review
```python

```

## Visualize the airline dataset
```python

```

## Analyze the airline dataset
```python

```

## Time series decomposition of the airline dataset
```python

```

# 4. Work with Multiple Time Series
## Working with more than one time series
```python

```

## Load multiple time series
```python

```

## Visualize multiple time series
```python

```

## Statistical summaries of multiple time series
```python

```

## Plot multiple time series
```python

```

## Define the color palette of your plots
```python

```

## Add summary statistics to your time series plot
```python

```

## Plot your time series on individual plots
```python

```

## Find relationships between multiple time series
```python

```

## Compute correlations between time series
```python

```

## Visualize correlation matrices
```python

```

## Clustered heatmaps
```python

```

# 5. Case Study
## Apply your knowledge to a new dataset
```python

```

## Explore the Jobs dataset
```python

```

## Describe time series data with boxplots
```python

```

## Beyond summary statistics
```python

```

## Plot all the time series in your dataset
```python

```

## Annotate significant events in time series data
```python

```

## Plot monthly and yearly trends
```python

```

## Decompose time series data
```python

```

## Apply time series decomposition to your dataset
```python

```

## Visualize the seasonality of multiple time series
```python

```

## Compute correlations between time series
```python

```

## Correlations between multiple time series
```python

```

## Interpret correlations
```python

```

## Congratulations!
```python

```


