---
title: Customer Analytics and A/B Testing in Python
tags: ab-testing, python
url: https://campus.datacamp.com/courses/customer-analytics-and-ab-testing-in-python/
---

# 1. Key Performance Indicators: Measuring Business Success
## Loading & examining our data
```python
# Import pandas 
import pandas as pd

# Load the customer_data
customer_data = pd.read_csv('customer_data.csv')

# Load the app_purchases
app_purchases = pd.read_csv('inapp_purchases.csv')

# Print the columns of customer data
print(customer_data.columns)

# Print the columns of app_purchases
print(app_purchases.columns)
```

## Merging on different sets of fields
```python
# Merge on the 'uid' and 'date' field
uid_date_combined_data = app_purchases.merge(customer_data, on=['uid', 'date'], how='inner')

# Examine the results 
print(uid_date_combined_data.head())
print(len(uid_date_combined_data))
```

## Practicing aggregations
```python
# Calculate the mean and median of price and age
purchase_summary = purchase_data.agg({'price': ['mean', 'median'], 'age': ['mean', 'median']})

# Examine the output 
print(purchase_summary)
```

## Grouping & aggregating
```python
# Group the data 
grouped_purchase_data = purchase_data.groupby(by = ['device', 'gender'])

# Aggregate the data
purchase_summary = grouped_purchase_data.agg({'price': ['mean', 'median', 'std']})

# Examine the results
print(purchase_summary)
```

## Calculating KPIs
```python
# Compute max_purchase_date
max_purchase_date = current_date - timedelta(days=28)

# Filter to only include users who registered before our max date
purchase_data_filt = purchase_data[purchase_data.reg_date < max_purchase_date]

# Filter to contain only purchases within the first 28 days of registration
purchase_data_filt = purchase_data_filt[(purchase_data_filt.date <= 
                        purchase_data_filt.reg_date + timedelta(days=28))]

# Output the mean price paid per purchase
print(purchase_data_filt.price.mean())
```

## Average purchase price by cohort
```python
# Set the max registration date to be one month before today
max_reg_date = current_date - timedelta(days=28)

# Find the month 1 values
month1 = np.where((purchase_data.reg_date < max_reg_date) &
                 (purchase_data.date < purchase_data.reg_date + timedelta(days=28)),
                  purchase_data.price, 
                  np.NaN)
                 
# Update the value in the DataFrame
purchase_data['month1'] = month1

# Group the data by gender and device 
purchase_data_upd = purchase_data.groupby(by=['gender', 'device'], as_index=False) 

# Aggregate the month1 and price data 
purchase_summary = purchase_data_upd.agg(
                        {'month1': ['mean', 'median'],
                        'price': ['mean', 'median']})

# Examine the results 
print(purchase_summary)
```


# 2. Exploring and Visualizing Customer Behavior
## Parsing dates
```python

```

## Creating time series graphs with matplotlib
```python

```

## Plotting time series data
```python

```

## Pivoting our data
```python

```

## Examining the different cohorts
```python

```

## Understanding and visualizing trends
```python

```

## Seasonality and moving averages
```python

```

## Exponential rolling average & over/under smoothing
```python

```

## Events and releases
```python

```

## Visualizing user spending
```python

```

## Looking more closely at revenue
```python

```



# 3. The Design and Application of A/B Testing
## Introduction to A/B testing
```python

```

## Good applications of A/B testing
```python

```

## General properties of an A/B Test
```python

```

## A/B test generalizability
```python

```

## Initial A/B test design
```python

```

## Experimental units: Revenue per user day
```python

```

## Preparing to run an A/B test
```python

```

## Conversion rate sensitivities
```python

```

## Sensitivity
```python

```

## Standard error
```python

```

## Calculating sample size
```python

```

## Exploring the power calculation
```python

```

## Calculating the sample size
```python

```



# 4. Analyzing A/B Testing Results
## Analyzing the A/B test results
```python

```

## Confirming our test results
```python

```

## Thinking critically about p-values
```python

```

## Understanding statistical significance
```python

```

## Intuition behind statistical significance
```python

```

## Checking for statistical significance
```python

```

## Understanding confidence intervals
```python

```

## Calculating confidence intervals
```python

```

## Interpreting your test results
```python

```

## Plotting the distribution
```python

```

## Plotting the difference distribution
```python

```

## Finale
```python

```

