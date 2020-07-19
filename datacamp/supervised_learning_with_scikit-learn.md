---
title: Supervised Learning with scikit-learn
tags: python,machine-learning
url: https://www.datacamp.com/courses/supervised-learning-with-scikit-learn
---

# 1. Classification
## k-Nearest Neighbors: Fit
```python
# Import KNeighborsClassifier from sklearn.neighbors
from sklearn.neighbors import KNeighborsClassifier

# Create arrays for the features and the response variable
y = df['party'].values
X = df.drop('party', axis=1).values

# Create a k-NN classifier with 6 neighbors
knn = KNeighborsClassifier(n_neighbors=6)

# Fit the classifier to the data
knn.fit(X, y)
```

## k-Nearest Neighbors: Predict
```python

```

## Measuring model performance
```python

```

## The digits recognition dataset
```python

```

## Train/Test Split + Fit/Predict/Accuracy
```python

```

## Overfitting and underfitting
```python

```



# 2. Regression
## Introduction to regression
```python

```

## Which of the following is a regression problem?
```python

```

## Importing data for supervised learning
```python

```

## Exploring the Gapminder data
```python

```

## The basics of linear regression
```python

```

## Fit & predict for regression
```python

```

## Train/test split for regression
```python

```

## Cross-validation
```python

```

## 5-fold cross-validation
```python

```

## K-Fold CV comparison
```python

```

## Regularized regression
```python

```

## Regularization I: Lasso
```python

```

## Regularization II: Ridge
```python

```



# 3. Fine-tuning your model
## How good is your model?
```python

```

## Metrics for classification
```python

```

## Logistic regression and the ROC curve
```python

```

## Building a logistic regression model
```python

```

## Plotting an ROC curve
```python

```

## Precision-recall Curve
```python

```

## Area under the ROC curve
```python

```

## AUC computation
```python

```

## Hyperparameter tuning
```python

```

## Hyperparameter tuning with GridSearchCV
```python

```

## Hyperparameter tuning with RandomizedSearchCV
```python

```

## Hold-out set for final evaluation
```python

```

## Hold-out set reasoning
```python

```

## Hold-out set in practice I: Classification
```python

```

## Hold-out set in practice II: Regression
```python

```



# 4. Preprocessing and pipelines
## Preprocessing data
```python

```

## Exploring categorical features
```python

```

## Creating dummy variables
```python

```

## Regression with categorical features
```python

```

## Handling missing data
```python

```

## Dropping missing data
```python

```

## Imputing missing data in a ML Pipeline I
```python

```

## Imputing missing data in a ML Pipeline II
```python

```

## Centering and scaling
```python

```

## Centering and scaling your data
```python

```

## Centering and scaling in a pipeline
```python

```

## Bringing it all together I: Pipeline for classification
```python

```

## Bringing it all together II: Pipeline for regression
```python

```

## Final thoughts
```python

```


