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
## Find missing values
```python
##
# Display first seven rows of co2_levels
print(co2_levels.head(7))

##
# Set datestamp column as index
co2_levels = co2_levels.set_index('datestamp')

# Print out the number of missing values
print(co2_levels.isnull().sum())
```

## Handle missing values
```python
# Impute missing values with the next valid observation
co2_levels = co2_levels.fillna(method='bfill')

# Print out the number of missing values
print(co2_levels.isnull().sum())
```

## Display rolling averages
```python
# Compute the 52 weeks rolling mean of the co2_levels DataFrame
ma = co2_levels.rolling(window=52).mean()

# Compute the 52 weeks rolling standard deviation of the co2_levels DataFrame
mstd = co2_levels.rolling(window=52).std()

# Add the upper bound column to the ma DataFrame
ma['upper'] = ma['co2'] + (2 * mstd['co2'])

# Add the lower bound column to the ma DataFrame
ma['lower'] = ma['co2'] - (2 * mstd['co2'])

# Plot the content of the ma DataFrame
ax = ma.plot(linewidth=0.8, fontsize=6)

# Specify labels, legend, and show the plot
ax.set_xlabel('Date', fontsize=10)
ax.set_ylabel('CO2 levels in Mauai Hawaii', fontsize=10)
ax.set_title('Rolling mean and variance of CO2 levels\nin Mauai Hawaii from 1958 to 2001', fontsize=10)
plt.show()
```

## Display aggregated values
```python
# Get month for each dates in the index of co2_levels
index_month = co2_levels.index.month

# Compute the mean CO2 levels for each month of the year
mean_co2_levels_by_month = co2_levels.groupby(index_month).mean()

# Plot the mean CO2 levels for each month of the year
mean_co2_levels_by_month.plot(fontsize=6)

# Specify the fontsize on the legend
plt.legend(fontsize=10)

# Show plot
plt.show()
```

## Compute numerical summaries
```python
# Print out summary statistics of the co2_levels DataFrame
print(co2_levels.describe())

# Print out the minima of the co2 column in the co2_levels DataFrame
print(co2_levels.describe().loc['min'])

# Print out the maxima of the co2 column in the co2_levels DataFrame
print(co2_levels.describe().loc['max'])
```

## Boxplots and Histograms
```python
##
# Generate a boxplot
ax = co2_levels.boxplot()

# Set the labels and display the plot
ax.set_xlabel('CO2', fontsize=10)
ax.set_ylabel('Boxplot CO2 levels in Maui Hawaii', fontsize=10)
plt.legend(fontsize=10)
plt.show()

##
# Generate a histogram
ax = co2_levels.plot(kind='hist', bins=50, fontsize=6)

# Set the labels and display the plot
ax.set_xlabel('CO2', fontsize=10)
ax.set_ylabel('Histogram of CO2 levels in Maui Hawaii', fontsize=10)
plt.legend(fontsize=10)
plt.show()
```

## Density plots
```python
# Display density plot of CO2 levels values
ax = co2_levels.plot(kind='density', linewidth=4, fontsize=6)

# Annotate x-axis labels
ax.set_xlabel('CO2', fontsize=10)

# Annotate y-axis labels
ax.set_ylabel('Density plot of CO2 levels in Maui Hawaii', fontsize=10)

plt.show()
```

# 3. Seasonality, Trend and Noise
## Autocorrelation in time series data
```python
# Import required libraries
import matplotlib.pyplot as plt
plt.style.use('fivethirtyeight')
from statsmodels.graphics import tsaplots

# Display the autocorrelation plot of your time series
fig = tsaplots.plot_acf(co2_levels['co2'], lags=24)

# Show plot
plt.show()
```

## Partial autocorrelation in time series data
```python
# Import required libraries
import matplotlib.pyplot as plt
plt.style.use('fivethirtyeight')
from statsmodels.graphics import tsaplots

# Display the partial autocorrelation plot of your time series
fig = tsaplots.plot_pacf(co2_levels['co2'], lags=24)

# Show plot
plt.show()
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


