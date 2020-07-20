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
# Import KNeighborsClassifier from sklearn.neighbors
from sklearn.neighbors import KNeighborsClassifier 

# Create arrays for the features and the response variable
y = df['party'].values
X = df.drop('party', axis=1).values

# Create a k-NN classifier with 6 neighbors: knn
knn = KNeighborsClassifier(n_neighbors=6)

# Fit the classifier to the data
knn.fit(X, y)

# Predict the labels for the training data X: y_pred
y_pred = knn.predict(X)

# Predict and print the label for the new data point X_new
new_prediction = knn.predict(X_new)
print("Prediction: {}".format(new_prediction)) 
```

## The digits recognition dataset
```python
# Import necessary modules
from sklearn import datasets
import matplotlib.pyplot as plt

# Load the digits dataset: digits
digits = datasets.load_digits()

# Print the keys and DESCR of the dataset
print(digits.keys())
print(digits.DESCR)

# Print the shape of the images and data keys
print(digits.images.shape)
print(digits.data.shape)

# Display digit 1010
plt.imshow(digits.images[1010], cmap=plt.cm.gray_r, interpolation='nearest')
plt.show()
```

## Train/Test Split + Fit/Predict/Accuracy
```python
# Import necessary modules
from sklearn.neighbors import KNeighborsClassifier 
from sklearn.model_selection import train_test_split

# Create feature and target arrays
X = digits.data
y = digits.target

# Split into training and test set
X_train, X_test, y_train, y_test = train_test_split(X, y, test_size = 0.2, random_state=42, stratify=y)

# Create a k-NN classifier with 7 neighbors: knn
knn = KNeighborsClassifier(n_neighbors=7)

# Fit the classifier to the training data
knn.fit(X_train, y_train)

# Print the accuracy
print(knn.score(X_test, y_test))
```

## Overfitting and underfitting
```python
# Setup arrays to store train and test accuracies
neighbors = np.arange(1, 9)
train_accuracy = np.empty(len(neighbors))
test_accuracy = np.empty(len(neighbors))

# Loop over different values of k
for i, k in enumerate(neighbors):
    # Setup a k-NN Classifier with k neighbors: knn
    knn = KNeighborsClassifier(n_neighbors=k)

    # Fit the classifier to the training data
    knn.fit(X_train, y_train)
    
    #Compute accuracy on the training set
    train_accuracy[i] = knn.score(X_train, y_train)

    #Compute accuracy on the testing set
    test_accuracy[i] = knn.score(X_test, y_test)

# Generate plot
plt.title('k-NN: Varying Number of Neighbors')
plt.plot(neighbors, test_accuracy, label = 'Testing Accuracy')
plt.plot(neighbors, train_accuracy, label = 'Training Accuracy')
plt.legend()
plt.xlabel('Number of Neighbors')
plt.ylabel('Accuracy')
plt.show()
```


# 2. Regression
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


