---
title: Customer Segmentation in Python
tags: customer-segmentation, python
url: https://campus.datacamp.com/courses/customer-segmentation-in-python
---

# 1. Cohort Analysis
## Assign daily acquisition cohort
```python
# Define a function that will parse the date
def get_day(x): return dt.datetime(x.year, x.month, x.day)

# Create InvoiceDay column
online['InvoiceDay'] = online['InvoiceDate'].apply(get_day)

# Group by CustomerID and select the InvoiceDay value
grouping = online.groupby('CustomerID')['InvoiceDay'] 

# Assign a minimum InvoiceDay value to the dataset
online['CohortDay'] = grouping.transform(min)

# View the top 5 rows
print(online.head())
```

## Calculate time offset in days - part 1
```python
# Get the integers for date parts from the `InvoiceDay` column
invoice_year, invoice_month, invoice_day = get_date_int(online, 'InvoiceDay')

# Get the integers for date parts from the `CohortDay` column
cohort_year, cohort_month, cohort_day = get_date_int(online, 'CohortDay')
```

## Calculate time offset in days - part 2
```python
# Calculate difference in years
years_diff = invoice_year - cohort_year

# Calculate difference in months
months_diff = invoice_month - cohort_month

# Calculate difference in days
days_diff = invoice_day - cohort_day

# Extract the difference in days from all previous values
online['CohortIndex'] = years_diff * 365 + months_diff * 30 + days_diff + 1
print(online.head())
```

## Customer retention
```python

```

## Calculate retention rate from scratch
```python

```

## Calculate average price
```python

```

## Visualizing cohort analysis
```python

```

## Visualize average quantity metric
```python

```


# 2. Recency, Frequency, Monetary Value analysis
## Recency, Frequency, Monetary (RFM) segmentation
```python

```

## Calculate Spend quartiles (q=4)
```python

```

## Calculate Recency deciles (q=4)
```python

```

## Calculating RFM metrics
```python

```

## Largest Frequency value
```python

```

## Calculate RFM values
```python

```

## Building RFM segments
```python

```

## Calculate 3 groups for Recency and Frequency
```python

```

## Calculate RFM Score
```python

```

## Analyzing RFM table
```python

```

## Find average value for RFM Score segment
```python

```

## Creating custom segments
```python

```

## Analyzing custom segments
```python

```


# 3. Data pre-processing for clustering
## Data pre-processing
```python

```

## Assumptions of k-means
```python

```

## Calculate statistics of variables
```python

```

## Managing skewed variables
```python

```

## Detect skewed variables
```python

```

## Manage skewness
```python

```

## Centering and scaling data
```python

```

## Center and scale manually
```python

```

## Center and scale with StandardScaler()
```python

```

## Pre-processing pipeline
```python

```

## Visualize RFM distributions
```python

```

## Pre-process RFM data
```python

```

## Visualize the normalized variables
```python

```


# 4. Customer Segmentation with K-means
## Practical implementation of k-means clustering
```python

```

## Run KMeans
```python

```

## Assign labels to raw data
```python

```

## Choosing the number of clusters
```python

```

## Calculate sum of squared errors
```python

```

## Plot sum of squared errors
```python

```

## Profile and interpret segments
```python

```

## Prepare data for the snake plot
```python

```

## Visualize snake plot
```python

```

## Calculate relative importance of each attribute
```python

```

## Plot relative importance heatmap
```python

```

## End-to-end segmentation solution
```python

```

## Pre-process data
```python

```

## Calculate and plot sum of squared errors
```python

```

## Build 4-cluster solution
```python

```

## Analyze the segments
```python

```

## Final thoughts
```python

```

