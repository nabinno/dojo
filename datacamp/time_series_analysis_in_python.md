---
title: Time Series Analysis in Python
tags: python,time-series,statistics
url: https://www.datacamp.com/courses/introduction-to-time-series-analysis-in-python
---

# 1. Correlation and Autocorrelation
## A "Thin" Application of Time Series
```python
##
# Import pandas and plotting modules
import pandas as pd
import matplotlib.pyplot as plt

# Convert the date index to datetime
diet.index = pd.to_datetime(diet.index)

##
# Plot the entire time series diet and show gridlines
diet.plot(grid=True)
plt.show()

##
# Slice the dataset to keep only 2012
diet2012 = diet['2012']

# Plot 2012 data
diet2012.plot(grid=True)
plt.show()
```

## Merging Time Series With Different Dates
```python
# Import pandas
import pandas as pd

# Convert the stock index and bond index into sets
set_stock_dates = set(stocks.index)
set_bond_dates = set(bonds.index)

# Take the difference between the sets and print
print(set_stock_dates - set_bond_dates)

# Merge stocks and bonds DataFrames using join()
stocks_and_bonds = stocks.join(bonds, how='inner')
```

## Correlation of Stocks and Bonds
```python
# Compute percent change using pct_change()
returns = stocks_and_bonds.pct_change()

# Compute correlation using corr()
correlation = returns['SP500'].corr(returns['US10Y'])
print("Correlation of stocks and interest rates: ", correlation)

# Make scatter plot
plt.scatter(returns['SP500'], returns['US10Y'])
plt.show()
```

## Flying Saucers Aren't Correlated to Flying Markets
```python

```

## Simple Linear Regression
```python

```

## Looking at a Regression's R-Squared
```python

```

## Match Correlation with Regression Output
```python

```

## Autocorrelation
```python

```

## A Popular Strategy Using Autocorrelation
```python

```

## Are Interest Rates Autocorrelated?
```python

```

# 2. Some Simple Time Series
## Autocorrelation Function
```python

```

## Taxing Exercise: Compute the ACF
```python

```

## Are We Confident This Stock is Mean Reverting?
```python

```

## White Noise
```python

```

## Can't Forecast White Noise
```python

```

## Random Walk
```python

```

## Generate a Random Walk
```python

```

## Get the Drift
```python

```

## Are Stock Prices a Random Walk?
```python

```

## How About Stock Returns?
```python

```

## Stationarity
```python

```

## Is it Stationary?
```python

```

## Seasonal Adjustment During Tax Season
```python

```

# 3. Autoregressive (AR) Models
## Describe AR Model
```python

```

## Simulate AR(1) Time Series
```python

```

## Compare the ACF for Several AR Time Series
```python

```

## Match AR Model with ACF
```python

```

## Estimating and Forecasting AR Model
```python

```

## Estimating an AR Model
```python

```

## Forecasting with an AR Model
```python

```

## Let's Forecast Interest Rates
```python

```

## Compare AR Model with Random Walk
```python

```

## Choosing the Right Model
```python

```

## Estimate Order of Model: PACF
```python

```

## Estimate Order of Model: Information Criteria
```python

```

# 4. Moving Average (MA) and ARMA Models
## Describe Model
```python

```

## Simulate MA(1) Time Series
```python

```

## Compute the ACF for Several MA Time Series
```python

```

## Match ACF with MA Model
```python

```

## Estimation and Forecasting an MA Model
```python

```

## Estimating an MA Model
```python

```

## Forecasting with MA Model
```python

```

## ARMA models
```python

```

## High Frequency Stock Prices
```python

```

## More Data Cleaning: Missing Data
```python

```

## Applying an MA Model
```python

```

## Equivalence of AR(1) and MA(infinity)
```python

```

# 5. Putting It All Together
## Cointegration Models
```python

```

## A Dog on a Leash? (Part 1)
```python

```

## A Dog on a Leash? (Part 2)
```python

```

## Are Bitcoin and Ethereum Cointegrated?
```python

```

## Case Study: Climate Change
```python

```

## Is Temperature a Random Walk (with Drift)?
```python

```

## Getting "Warmed" Up: Look at Autocorrelations
```python

```

## Which ARMA Model is Best?
```python

```

## Don't Throw Out That Winter Coat Yet
```python

```

## Congratulations
```python

```

