---
title: Manipulating Time Series Data in Python
tags: python,time-series,data-pre-processing
url: https://www.datacamp.com/courses/manipulating-time-series-data-in-python
---

# 1. Working with Time Series in Pandas
## Your first time series
```python
# Create the range of dates here
seven_days = pd.date_range('2017-1-1', '2017-1-7')

# Iterate over the dates and print the number and name of the weekday
for day in seven_days:
    print(day.dayofweek, day.weekday_name)
```

## Create a time series of air quality data
```python
data = pd.read_csv('nyc.csv')

# Inspect data
print(data.info())

# Convert the date column to datetime64
data['date'] = pd.to_datetime(data['date'])

# Set date column as index
data.set_index('date', inplace=True)

# Inspect data 
print(data.info())

# Plot data
data.plot(subplots=True)
plt.show()
```

## Compare annual stock price trends
```python

```

## Set and change time series frequency
```python

```

## Lags, changes, and returns for stock price series
```python

```

## Shifting stock prices across time
```python

```

## Calculating stock price changes
```python

```

## Plotting multi-period returns
```python

```

# 2. Basic Time Series Metrics & Resampling
## Compare time series growth rates
```python

```

## Compare the performance of several asset classes
```python

```

## Comparing stock prices with a benchmark
```python

```

## Plot performance difference vs benchmark index
```python

```

## Changing the time series frequency: resampling
```python

```

## Convert monthly to weekly data
```python

```

## Create weekly from monthly unemployment data
```python

```

## Upsampling & interpolation with `.resample()`
```python

```

## Use interpolation to create weekly employment data
```python

```

## Interpolate debt/GDP and compare to unemployment
```python

```

## Downsampling & aggregation
```python

```

## Compare weekly, monthly and annual ozone trends for NYC & LA
```python

```

## Compare monthly average stock prices for Facebook and Google
```python

```

## Compare quarterly GDP growth rate and stock returns
```python

```

## Visualize monthly mean, median and standard deviation of S&P500 returns
```python

```

# 3. Window Functions: Rolling & Expanding Metrics
## Rolling window functions with pandas
```python

```

## Rolling average air quality since 2010 for new york city
```python

```

## Rolling 360-day median & std. deviation for nyc ozone data since 2000
```python

```

## Rolling quantiles for daily air quality in nyc
```python

```

## Expanding window functions with pandas
```python

```

## Cumulative sum vs .diff()
```python

```

## Cumulative return on $1,000 invested in google vs apple I
```python

```

## Cumulative return on $1,000 invested in google vs apple II
```python

```

## Case study: S&P500 price simulation
```python

```

## Random walk I
```python

```

## Random walk II
```python

```

## Random walk III
```python

```

## Relationships between time series: correlation
```python

```

## Annual return correlations among several stocks
```python

```

# 4. Putting it all together: Building a value-weighted index
## Select index components & import data
```python

```

## Explore and clean company listing information
```python

```

## Select and inspect index components
```python

```

## Import index component price information
```python

```

## Build a market-cap weighted index
```python

```

## Calculate number of shares outstanding
```python

```

## Create time series of market value
```python

```

## Calculate & plot the composite index
```python

```

## Evaluate index performance
```python

```

## Calculate the contribution of each stock to the index
```python

```

## Compare index performance against benchmark I
```python

```

## Compare index performance against benchmark II
```python

```

## Index correlation & exporting to excel
```python

```

## Visualize your index constituent correlations
```python

```

## Save your analysis to multiple excel worksheets
```python

```

## Congratulations!
```python

```


