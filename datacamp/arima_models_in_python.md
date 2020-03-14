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

## Augmented Dicky-Fuller
```python
# Import augmented dicky-fuller test function
from statsmodels.tsa.stattools import adfuller

# Run test
result = adfuller(earthquake['earthquakes_per_year'])

# Print test statistic
print(result[0])

# Print p-value
print(result[1])

# Print critical values
print(result[4])
```

## Taking the difference
```python
##
# Run the ADF test on the time series
result = adfuller(city['city_population'])

# Plot the time series
fig, ax = plt.subplots()
city.plot(ax=ax)
plt.show()

# Print the test statistic and the p-value
print('ADF Statistic:', result[0])
print('p-value:', result[1])

##
# Calculate the first difference of the time series
city_stationary = city.diff().dropna()

# Run ADF test on the differenced time series
result = adfuller(city_stationary['city_population'])

# Plot the differenced time series
fig, ax = plt.subplots()
city_stationary.plot(ax=ax)
plt.show()

# Print the test statistic and the p-value
print('ADF Statistic:', result[0])
print('p-value:', result[1])

##
# Calculate the second difference of the time series
city_stationary = city.diff().diff().dropna()

# Run ADF test on the differenced time series
result = adfuller(city_stationary['city_population'])

# Plot the differenced time series
fig, ax = plt.subplots()
city_stationary.plot(ax=ax)
plt.show()

# Print the test statistic and the p-value
print('ADF Statistic:', result[0])
print('p-value:', result[1])
```

## Other tranforms
```python
##
# Calculate the first difference and drop the nans
amazon_diff = amazon.diff().dropna()

# Run test and print
result_diff = adfuller(amazon_diff['close'])
print(result_diff)

##
# Calculate the first difference and drop the nans
amazon_diff = amazon.diff().dropna()

# Run test and print
result_diff = adfuller(amazon_diff['close'])
print(result_diff)

# Calculate log-return and drop nans
amazon_log = np.log((amazon/amazon.shift(1)).dropna())

# Run test and print
result_log = adfuller(amazon_log['close'])
print(result_log)
```

## Generating ARMA data
```python
##
# Import data generation function and set random seed
from statsmodels.tsa.arima_process import arma_generate_sample
np.random.seed(1)

# Set coefficients
ar_coefs = [1]
ma_coefs = [1, -0.7]

# Generate data
y = arma_generate_sample(ar_coefs, ma_coefs, nsample=100, sigma=0.5, )

plt.plot(y)
plt.ylabel(r'$y_t$')
plt.xlabel(r'$t$')
plt.show()

##
# Import data generation function and set random seed
from statsmodels.tsa.arima_process import arma_generate_sample
np.random.seed(2)

# Set coefficients
ar_coefs = [1, -0.3, -0.2]
ma_coefs = [1]

# Generate data
y = arma_generate_sample(ar_coefs, ma_coefs, nsample=100, sigma=0.5, )

plt.plot(y)
plt.ylabel(r'$y_t$')
plt.xlabel(r'$t$')
plt.show()

##
# Import data generation function and set random seed
from statsmodels.tsa.arima_process import arma_generate_sample
np.random.seed(3)

# Set coefficients
ar_coefs = [1, 0.2]
ma_coefs = [1, 0.3, 0.4]

# Generate data
y = arma_generate_sample(ar_coefs, ma_coefs, nsample=100, sigma=0.5, )

plt.plot(y)
plt.ylabel(r'$y_t$')
plt.xlabel(r'$t$')
plt.show()
```

## Fitting Prelude
```python
# Import the ARMA model
from statsmodels.tsa.arima_model import ARMA

# Instantiate the model
model = ARMA(y, order=(1, 1))

# Fit the model
results = model.fit()
```

# 2. Fitting the Future
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
