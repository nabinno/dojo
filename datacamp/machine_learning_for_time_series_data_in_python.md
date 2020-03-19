---
title: Machine Learning for Time Series Data in Python
tags: python, time-series, machine-learning
url: https://www.datacamp.com/courses/machine-learning-for-time-series-data-in-python
---

# 1. Time Series and Machine Learning Primer
## Plotting a time series (I)
```python
##
# Print the first 5 rows of data
print(data.head())

##
# Print the first 5 rows of data2
print(data2.head())

##
# Plot the time series in each dataset
fig, axs = plt.subplots(2, 1, figsize=(5, 10))
data.iloc[:1000].plot(y='data_values', ax=axs[0])
data2.iloc[:1000].plot(y='data_values', ax=axs[1])
plt.show()
```

## Plotting a time series (II)
```python
# Plot the time series in each dataset
fig, axs = plt.subplots(2, 1, figsize=(5, 10))
data.iloc[:1000].plot(x='time', y='data_values', ax=axs[0])
data2.iloc[:1000].plot(x='time', y='data_values', ax=axs[1])
plt.show()
```

## Machine learning basics
```python
##
# Print the first 5 rows for inspection
print(data.head())

##
from sklearn.svm import LinearSVC

# Construct data for the model
X = data[['petal length (cm)', 'petal width (cm)']]
y = data[['target']]

# Fit the model
model = LinearSVC()
model.fit(X, y)
```

## Predicting using a classification model
```python
# Create input array
X_predict = targets[['petal length (cm)', 'petal width (cm)']]

# Predict with the model
predictions = model.predict(X_predict)
print(predictions)

# Visualize predictions and actual values
plt.scatter(X_predict['petal length (cm)'], X_predict['petal width (cm)'],
            c=predictions, cmap=plt.cm.coolwarm)
plt.title("Predicted class values")
plt.show()
```

## Fitting a simple model: regression
```python

```

## Predicting using a regression model
```python

```

## Machine learning and time series data
```python

```

## Inspecting the classification data
```python

```

## Inspecting the regression data
```python

```


# 2. Time Series as Inputs to a Model
## Classifying a time series
```python

```

## Many repetitions of sounds
```python

```

## Invariance in time
```python

```

## Build a classification model
```python

```

## Improving features for classification
```python

```

## Calculating the envelope of sound
```python

```

## Calculating features from the envelope
```python

```

## Derivative features: The tempogram
```python

```

## The spectrogram
```python

```

## Spectrograms of heartbeat audio
```python

```

## Engineering spectral features
```python

```

## Combining many features in a classifier
```python

```


# 3. Predicting Time Series Data
## Predicting data over time
```python

```

## Introducing the dataset
```python

```

## Fitting a simple regression model
```python

```

## Visualizing predicted values
```python

```

## Advanced time series prediction
```python

```

## Visualizing messy data
```python

```

## Imputing missing values
```python

```

## Transforming raw data
```python

```

## Handling outliers
```python

```

## Creating features over time
```python

```

## Engineering multiple rolling features at once
```python

```

## Percentiles and partial functions
```python

```

## Using "date" information
```python

```


# 4. Validating and Inspecting Time Series Models
## Creating features from the past
```python

```

## Creating time-shifted features
```python

```

## Special case: Auto-regressive models
```python

```

## Visualize regression coefficients
```python

```

## Auto-regression with a smoother time series
```python

```

## Cross-validating time series data
```python

```

## Cross-validation with shuffling
```python

```

## Cross-validation without shuffling
```python

```

## Time-based cross-validation
```python

```

## Stationarity and stability
```python

```

## Stationarity
```python

```

## Bootstrapping a confidence interval
```python

```

## Calculating variability in model coefficients
```python

```

## Visualizing model score variability over time
```python

```

## Accounting for non-stationarity
```python

```

## Wrap-up
```python

```

