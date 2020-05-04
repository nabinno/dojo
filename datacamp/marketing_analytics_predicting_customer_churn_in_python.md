---
title: Marketing Analytics: Predicting Customer Churn in Python
tagas: marketing, customer-development, python
url: https://campus.datacamp.com/courses/marketing-analytics-predicting-customer-churn-in-python
---

# 1. Exploratory Data Analysis
## Exploring customer churn
```python
telco['Churn].value_counts()
```

## Summary statistics for both classes
```python
##
# Group telco by 'Churn' and compute the mean
print(telco.groupby(['Churn']).mean())

##
# Adapt your code to compute the standard deviation
print(telco.groupby(['Churn']).std())
```

## Churn by State
```python
##
# Count the number of churners and non-churners by State
print(telco.groupby('State')['Churn'].value_counts())
```

## Exploring your data using visualizations
```python

```

## Exploring feature distributions
```python

```

## Customer service calls and churn
```python

```


# 2. Preprocessing for Churn Modeling
## Data preprocessing
```python

```

## Identifying features to convert
```python

```

## Encoding binary features
```python

```

## One hot encoding
```python

```

## Feature scaling
```python

```

## Feature selection and engineering
```python

```

## Dropping unnecessary features
```python

```

## Engineering a new column
```python

```


# 3. Churn Prediction
## Making Predictions
```python

```

## Predicting whether a new customer will churn
```python

```

## Training another scikit-learn model
```python

```

## Evaluating Model Performance
```python

```

## Creating training and test sets
```python

```

## Check each sets length
```python

```

## Computing accuracy
```python

```

## Model Metrics
```python

```

## Confusion matrix
```python

```

## Varying training set size
```python

```

## Computing precision and recall
```python

```

## Other model metrics
```python

```

## ROC curve
```python

```

## Area under the curve
```python

```

## Precision-recall curve
```python

```

## F1 score
```python

```


# 4. Model Tuning
## Tuning your model
```python

```

## Tuning the number of features
```python

```

## Tuning other hyperparameters
```python

```

## Randomized search
```python

```

## Feature importances
```python

```

## Visualizing feature importances
```python

```

## Improving the plot
```python

```

## Interpreting feature importances
```python

```

## Adding new features
```python

```

## Does model performance improve?
```python

```

## Computing other metrics
```python

```

## Final thoughts
```python

```


