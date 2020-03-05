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
# Set the date column as the index of your DataFrame discoveries
discoveries = discoveries.set_index('date')

# Plot the time series in your DataFrame
ax = discoveries.plot(color='blue')

# Specify the x-axis label in your plot
ax.set_xlabel('Date')

# Specify the y-axis label in your plot
ax.set_ylabel('Number of great discoveries')

# Show plot
plt.show()
```

## Specify plot styles
```python
##
# Import the matplotlib.pyplot sub-module
import matplotlib.pyplot as plt

# Use the fivethirtyeight style
plt.style.use('fivethirtyeight')

# Plot the time series
ax1 = discoveries.plot()
ax1.set_title('FiveThirtyEight Style')
plt.show()

##
# Import the matplotlib.pyplot sub-module
import matplotlib.pyplot as plt

# Use the ggplot style
plt.style.use('ggplot')
ax2 = discoveries.plot()

# Set the title
ax2.set_title('ggplot Style')
plt.show()
```

## Display and label plots
```python
# Plot a line chart of the discoveries DataFrame using the specified arguments
ax = discoveries.plot(color='blue', figsize=(8, 3), linewidth=2, fontsize=6)

# Specify the title in your plot
ax.set_title('Number of great inventions and scientific discoveries from 1860 to 1959', fontsize=8)

# Show plot
plt.show()
```

## Subset time series data
```python
##
# Select the subset of data between 1945 and 1950
discoveries_subset_1 = discoveries['1945-01-01':'1950-01-01']

# Plot the time series in your DataFrame as a blue area chart
ax = discoveries_subset_1.plot(color='blue', fontsize=15)

# Show plot
plt.show()

##
# Select the subset of data between 1939 and 1958
discoveries_subset_2 = discoveries['1939-01-01':'1958-01-01']

# Plot the time series in your DataFrame as a blue area chart
ax = discoveries_subset_2.plot(color='blue', fontsize=15)

# Show plot
plt.show()
```

## Add vertical and horizontal markers
```python
# Plot your the discoveries time series
ax = discoveries.plot(color='blue', fontsize=6)

# Add a red vertical line
ax.axvline('1939-01-01', color='red', linestyle='--')

# Add a green horizontal line
ax.axhline(4, color='green', linestyle='--')

plt.show()
```

## Add shaded regions to your plot
```python
# Plot your the discoveries time series
ax = discoveries.plot(color='blue', fontsize=6)

# Add a vertical red shaded region
ax.axvspan('1900-01-01', '1915-01-01', color='red', alpha=.3)

# Add a horizontal green shaded region
ax.axhspan(6, 8, color='green', alpha=.3)

plt.show()
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


