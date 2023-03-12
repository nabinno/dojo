---
title: Machine Learning with Tree-Based Models in Python
tags: python,machine-learning
url: https://app.datacamp.com/learn/courses/machine-learning-with-tree-based-models-in-python
---

# 1. Classification and Regression Trees
## Train your first classification tree
```python
# Import DecisionTreeClassifier from sklearn.tree
from sklearn.tree import DecisionTreeClassifier

# Instantiate a DecisionTreeClassifier 'dt' with a maximum depth of 6
dt = DecisionTreeClassifier(max_depth=6, random_state=SEED)

# Fit dt to the training set
dt.fit(X_train, y_train)

# Predict test set labels
y_pred = dt.predict(X_test)
print(y_pred[0:5])
```

## Evaluate the classification tree
```python
# Import accuracy_score
from sklearn.metrics import accuracy_score

# Predict test set labels
y_pred = dt.predict(X_test)

# Compute test set accuracy  
acc = accuracy_score(y_test, y_pred)
print("Test set accuracy: {:.2f}".format(acc))
```

## Logistic regression vs classification tree
```python
# Import LogisticRegression from sklearn.linear_model
from sklearn.linear_model import  LogisticRegression

# Instatiate logreg
logreg = LogisticRegression(random_state=1)

# Fit logreg to the training set
logreg.fit(X_train, y_train)

# Define a list called clfs containing the two classifiers logreg and dt
clfs = [logreg, dt]

# Review the decision regions of the two classifiers
plot_labeled_decision_regions(X_test, y_test, clfs)
```

## Using entropy as a criterion
```python
# Import DecisionTreeClassifier from sklearn.tree
from sklearn.tree import DecisionTreeClassifier

# Instantiate dt_entropy, set 'entropy' as the information criterion
dt_entropy = DecisionTreeClassifier(max_depth=8, criterion='entropy', random_state=1)

# Fit dt_entropy to the training set
dt_entropy.fit(X_train, y_train)
```

## Entropy vs Gini index
```python
# Import accuracy_score from sklearn.metrics
from sklearn.metrics import accuracy_score

# Use dt_entropy to predict test set labels
y_pred = dt_entropy.predict(X_test)

# Evaluate accuracy_entropy
accuracy_entropy = accuracy_score(y_test, y_pred)

# Print accuracy_entropy
print(f'Accuracy achieved by using entropy: {accuracy_entropy:.3f}')

# Print accuracy_gini
print(f'Accuracy achieved by using the gini index: {accuracy_gini:.3f}')
```

## Train your first regression tree
```python
# Import DecisionTreeRegressor from sklearn.tree
from sklearn.tree import DecisionTreeRegressor

# Instantiate dt
dt = DecisionTreeRegressor(max_depth=8,
             min_samples_leaf=0.13,
            random_state=3)

# Fit dt to the training set
dt.fit(X_train, y_train)
```

## Evaluate the regression tree
```python
# Import mean_squared_error from sklearn.metrics as MSE
from sklearn.metrics import mean_squared_error as MSE

# Compute y_pred
y_pred = dt.predict(X_test)

# Compute mse_dt
mse_dt = MSE(y_test, y_pred)

# Compute rmse_dt
rmse_dt = mse_dt**(1/2)

# Print rmse_dt
print("Test set RMSE of dt: {:.2f}".format(rmse_dt))
```

## Linear regression vs regression tree
```python
# Predict test set labels 
y_pred_lr = lr.predict(X_test)

# Compute mse_lr
mse_lr = MSE(y_test, y_pred_lr)

# Compute rmse_lr
rmse_lr = mse_lr**(1/2)

# Print rmse_lr
print('Linear Regression test set RMSE: {:.2f}'.format(rmse_lr))

# Print rmse_dt
print('Regression Tree test set RMSE: {:.2f}'.format(rmse_dt))
```



# 2. The Bias-Variance Tradeoff
## Complexity, bias and variance
```python
Complexity, bias and variance
- As the complexity of f increases, the bias term decreases while the variance term increases.
```

## Instantiate the model
```python
# Import train_test_split from sklearn.model_selection
from sklearn.model_selection import train_test_split

# Set SEED for reproducibility
SEED = 1

# Split the data into 70% train and 30% test
X_train, X_test, y_train, y_test = train_test_split(X, y, test_size=0.3, random_state=SEED)

# Instantiate a DecisionTreeRegressor dt
dt = DecisionTreeRegressor(max_depth=4, min_samples_leaf=0.26, random_state=SEED)
```

## Evaluate the 10-fold CV error
```python
# Compute the array containing the 10-folds CV MSEs
MSE_CV_scores = - cross_val_score(dt, X_train, y_train, cv=10, 
                                  scoring='neg_mean_squared_error', 
                                  n_jobs=-1) 

# Compute the 10-folds CV RMSE
RMSE_CV = (MSE_CV_scores.mean())**(1/2)

# Print RMSE_CV
print('CV RMSE: {:.2f}'.format(RMSE_CV))
```

## Evaluate the training error
```python
# Import mean_squared_error from sklearn.metrics as MSE
from sklearn.metrics import mean_squared_error as MSE

# Fit dt to the training set
dt.fit(X_train, y_train)

# Predict the labels of the training set
y_pred_train = dt.predict(X_train)

# Evaluate the training set RMSE of dt
RMSE_train = (MSE(y_train, y_pred_train))**(1/2)

# Print RMSE_train
print('Train RMSE: {:.2f}'.format(RMSE_train))
```

## High bias or high variance?
```
## Problem
dt suffers from high bias because RMSE_CV = RMSE_train and both scores are greater than baseline_RMSE.

## IPython Shell
In [2]: RMSE_train
Out[2]: 5.15
In [3]: RMSE_CV
Out[3]: 5.14
In [4]: baseline_RMSE
Out[4]: 5.1
```

## Define the ensemble
```python
# Set seed for reproducibility
SEED=1

# Instantiate lr
lr = LogisticRegression(random_state=SEED)

# Instantiate knn
knn = KNN(n_neighbors=27)

# Instantiate dt
dt = DecisionTreeClassifier(min_samples_leaf=0.13, random_state=SEED)

# Define the list classifiers
classifiers = [('Logistic Regression', lr), ('K Nearest Neighbours', knn), ('Classification Tree', dt)]
```

## Evaluate individual classifiers
```python

```

## Better performance with a Voting Classifier
```python

```




# 3. Bagging and Random Forests
## Bagging
```python

```

## Define the bagging classifier
```python

```

## Evaluate Bagging performance
```python

```

## Out of Bag Evaluation
```python

```

## Prepare the ground
```python

```

## OOB Score vs Test Set Score
```python

```

## Random Forests (RF)
```python

```

## Train an RF regressor
```python

```

## Evaluate the RF regressor
```python

```

## Visualizing features importances
```python

```




# 4. Boosting
0%
Boosting refers to an ensemble method in which several models are trained sequentially with each model learning from the errors of its predecessors. In this chapter, you'll be introduced to the two boosting methods of AdaBoost and Gradient Boosting.

## Adaboost
```python

```

## Define the AdaBoost classifier
```python

```

## Train the AdaBoost classifier
```python

```

## Evaluate the AdaBoost classifier
```python

```

## Gradient Boosting (GB)
```python

```

## Define the GB regressor
```python

```

## Train the GB regressor
```python

```

## Evaluate the GB regressor
```python

```

## Stochastic Gradient Boosting (SGB)
```python

```

## Regression with SGB
```python

```

## Train the SGB regressor
```python

```

## Evaluate the SGB regressor
```python

```




# 4. Model Tuning
## Tuning a CART's Hyperparameters
```python

```

## Tree hyperparameters
```python

```

## Set the tree's hyperparameter grid
```python

```

## Search for the optimal tree
```python

```

## Evaluate the optimal tree
```python

```

## Tuning a RF's Hyperparameters
```python

```

## Random forests hyperparameters
```python

```

## Set the hyperparameter grid of RF
```python

```

## Search for the optimal forest
```python

```

## Evaluate the optimal forest
```python

```

## Congratulations!
```python

```

