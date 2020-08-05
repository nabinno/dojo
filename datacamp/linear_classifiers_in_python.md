---
title: Linear Classifiers in Python
tags: python, machine-learning
url: https://www.datacamp.com/courses/linear-classifiers-in-python
---

# 1. Applying logistic regression and SVM
## KNN classification
```python
from sklearn.neighbors import KNeighborsClassifier

# Create and fit the model
knn = KNeighborsClassifier()
knn.fit(X_train, y_train)

# Predict on the test features, print the results
pred = knn.predict(X_test)[0]
print("Prediction for test example 0:", pred)
```

## Comparing models
```python

```

## Overfitting
```python

```

## Applying logistic regression and SVM
```python

```

## Running LogisticRegression and SVC
```python

```

## Sentiment analysis for movie reviews
```python

```

## Linear classifiers
```python

```

## Which decision boundary is linear?
```python

```

## Visualizing decision boundaries
```python

```


# 2. Loss functions
## Linear classifiers: the coefficients
```python

```

## How models make predictions
```python

```

## Changing the model coefficients
```python

```

## What is a loss function?
```python

```

## The 0-1 loss
```python

```

## Minimizing a loss function
```python

```

## Loss function diagrams
```python

```

## Classification loss functions
```python

```

## Comparing the logistic and hinge losses
```python

```

## Implementing logistic regression
```python

```



# 3. Logistic regression
## Logistic regression and regularization
```python

```

## Regularized logistic regression
```python

```

## Logistic regression and feature selection
```python

```

## Identifying the most positive and negative words
```python

```

## Logistic regression and probabilities
```python

```

## Getting class probabilities
```python

```

## Regularization and probabilities
```python

```

## Visualizing easy and difficult examples
```python

```

## Multi-class logistic regression
```python

```

## Counting the coefficients
```python

```

## Fitting multi-class logistic regression
```python

```

## Visualizing multi-class logistic regression
```python

```

## One-vs-rest SVM
```python

```


# 4. Support Vector Machines
## Support vectors
```python

```

## Support vector definition
```python

```

## Effect of removing examples
```python

```

## Kernel SVMs
```python

```

## GridSearchCV warm-up
```python

```

## Jointly tuning gamma and C with GridSearchCV
```python

```

## Comparing logistic regression and SVM (and beyond)
```python

```

## An advantage of SVMs
```python

```

## An advantage of logistic regression
```python

```

## Using SGDClassifier
```python

```

## Conclusion
```python

```


