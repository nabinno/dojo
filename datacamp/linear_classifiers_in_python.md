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

## Running LogisticRegression and SVC
```python
from sklearn import datasets
digits = datasets.load_digits()
X_train, X_test, y_train, y_test = train_test_split(digits.data, digits.target)

# Apply logistic regression and print scores
lr = LogisticRegression()
lr.fit(X_train, y_train)
print(lr.score(X_train, y_train))
print(lr.score(X_test, y_test))

# Apply SVM and print scores
svm = SVC()
svm.fit(X_train, y_train)
print(svm.score(X_train, y_train))
print(svm.score(X_test, y_test))
```

## Sentiment analysis for movie reviews
```python
# Instantiate logistic regression and train
lr = LogisticRegression()
lr.fit(X, y)

# Predict sentiment for a glowing review
review1 = "LOVED IT! This movie was amazing. Top 10 this year."
review1_features = get_features(review1)
print("Review:", review1)
print("Probability of positive review:", lr.predict_proba(review1_features)[0,1])

# Predict sentiment for a poor review
review2 = "Total junk! I'll never watch a film by that director again, no matter how good the reviews."
review2_features = get_features(review2)
print("Review:", review2)
print("Probability of positive review:", lr.predict_proba(review2_features)[0,1])
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


