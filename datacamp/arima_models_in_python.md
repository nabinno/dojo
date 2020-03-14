---
title: ARIMA Models in Python
tags: arima, autoregressive-model, moving-average, time-series, statistics, python
url: https://campus.datacamp.com/courses/arima-models-in-python/chapter-1-arma-models
---

# 1. ARMA Models
## Exploration
```python
# Import modules
import pandas as pd
import matplotlib.pyplot as plt

# Load in the time series
candy = pd.read_csv('candy_production.csv', 
                 index_col='date', 
                 parse_dates=True)

# Plot and show the time series on axis ax
fig, ax = plt.subplots()
candy.plot(ax=ax)
plt.show()
```

## Train-test splits
```python
# Split the data into a train and test set
candy_train = candy.loc[:'2006']
candy_test = candy.loc['2007':]

# Create an axis
fig, ax = plt.subplots()

# Plot the train and test sets on the axis ax
candy_train.plot(ax=ax)
candy_test.plot(ax=ax)
plt.show()
```

## Is it stationary
```python
##
df1.plot()
plt.show()

##
df2.plot()
plt.show()

##
df3.plot()
plt.show()
```

## Making time series stationary
```python

```

## Augmented Dicky-Fuller
```python

```

## Taking the difference
```python

```

## Other tranforms
```python

```

## Intro to AR, MA and ARMA models
```python

```

## Model order
```python

```

## Generating ARMA data
```python

```

## Fitting Prelude
```python

```


# 2. Fitting the Future
## Fitting time series models
```python

```

## Fitting AR and MA models
```python

```

## Fitting an ARMA model
```python

```

## Fitting an ARMAX model
```python

```

## Forecasting
```python

```

## Generating one-step-ahead predictions
```python

```

## Plotting one-step-ahead predictions
```python

```

## Generating dynamic forecasts
```python

```

## Plotting dynamic forecasts
```python

```

## Intro to ARIMA models
```python

```

## Differencing and fitting ARMA
```python

```

## Unrolling ARMA forecast
```python

```

## Fitting an ARIMA model
```python

```

## Choosing ARIMA model
```python

```


# 3. The Best of the Best Models
## Intro to ACF and PACF
```python

```

## AR or MA
```python

```

## Order of earthquakes
```python

```

## Intro to AIC and BIC
```python

```

## Searching over model order
```python

```

## Choosing order with AIC and BIC
```python

```

## AIC and BIC vs ACF and PACF
```python

```

## Model diagnostics
```python

```

## Mean absolute error
```python

```

## Diagnostic summary statistics
```python

```

## Plot diagnostics
```python

```

## Box-Jenkins method
```python

```

## Identification
```python

```

## Identification II
```python

```

## Estimation
```python

```

## Diagnostics
```python

```


# 4. Seasonal ARIMA Models
## Seasonal time series
```python

```

## Seasonal decompose
```python

```

## Seasonal ACF and PACF
```python

```

## SARIMA models
```python

```

## Fitting SARIMA models
```python

```

## Choosing SARIMA order
```python

```

## SARIMA vs ARIMA forecasts
```python

```

## Automation and saving
```python

```

## Automated model selection
```python

```

## Saving and updating models
```python

```

## SARIMA and Box-Jenkins
```python

```

## Multiplicative vs additive seasonality
```python

```

## SARIMA model diagnostics
```python

```

## SARIMA forecast
```python

```

## Congratulations!
```python

```
