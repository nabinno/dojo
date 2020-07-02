---
title: Machine Learning for Marketing in Python
tags: machine-learning, marketing, python
url: https://campus.datacamp.com/courses/machine-learning-for-marketing-in-python
---

# 1. Machine learning for marketing basics
## Supervised vs. unsupervised learning
```python
##
# Print header of telco dataset
print(telco.head())
```

## Investigate the data
```python
# Print the data types of telco_raw dataset
print(telco_raw.dtypes)

# Print the header of telco_raw dataset
print(telco_raw.head())

# Print the number of unique values in each telco_raw column
print(telco_raw.nunique())
```

## Separate numerical and categorical columns
```python
# Store customerID and Churn column names
custid = ['customerID']
target = ['Churn']

# Store categorical column names
categorical = telco_raw.nunique()[telco_raw.nunique() < 5].keys().tolist()

# Remove target from the list of categorical variables
categorical.remove(target[0])

# Store numerical column names
numerical = [x for x in telco_raw.columns if x not in custid + target + categorical]
```

## Encode categorical and scale numerical variables
```python
# Perform one-hot encoding to categorical variables 
telco_raw = pd.get_dummies(data = telco_raw, columns = categorical, drop_first=True)

# Initialize StandardScaler instance
scaler = StandardScaler()

# Fit and transform the scaler on numerical columns
scaled_numerical = scaler.fit_transform(telco_raw[numerical])

# Build a DataFrame from scaled_numerical
scaled_numerical = pd.DataFrame(scaled_numerical, columns=numerical)
```

## Split data to training and testing
```python
# Split X and Y into training and testing datasets
train_X, test_X, train_Y, test_Y = train_test_split(X, Y, test_size=0.25)

# Ensure training dataset has only 75% of original X data
print(train_X.shape[0] / X.shape[0])

# Ensure testing dataset has only 25% of original X data
print(test_X.shape[0] / X.shape[0])
```

## Fit a decision tree
```python
# Initialize the model with max_depth set at 5
mytree = tree.DecisionTreeClassifier(max_depth = 5)

# Fit the model on the training data
treemodel = mytree.fit(train_X, train_Y)

# Predict values on the testing data
pred_Y = treemodel.predict(test_X)

# Measure model performance on testing data
accuracy_score(test_Y, pred_Y)
```

## Predict churn with decision tree
```python
# Initialize the Decision Tree
clf = tree.DecisionTreeClassifier(max_depth = 7, 
                                  criterion = 'gini', 
                                  splitter  = 'best')

# Fit the model to the training data
clf = clf.fit(train_X, train_Y)

# Predict the values on test dataset
pred_Y = clf.predict(test_X)

# Print accuracy values
print("Training accuracy: ", np.round(clf.score(train_X, train_Y), 3)) 
print("Test accuracy: ", np.round(accuracy_score(test_Y, pred_Y), 3))
```



# 2. Churn prediction and drivers
## Explore churn rate and split data
```python
# Print the unique Churn values
print(set(telcom['Churn']))

# Calculate the ratio size of each churn group
telcom.groupby(['Churn']).size() / telcom.shape[0] * 100

# Import the function for splitting data to train and test
from sklearn.model_selection import train_test_split

# Split the data into train and test
train, test = train_test_split(telcom, test_size = .25)
```

## Separate features and target variable
```python
# Store column names from `telcom` excluding target variable and customer ID
cols = [col for col in telcom.columns if col not in custid + target]

# Extract training features
train_X = train[cols]

# Extract training target
train_Y = train[target]

# Extract testing features
test_X = test[cols]

# Extract testing target
test_Y = test[target]
```

## Fit logistic regression model
```python
# Fit logistic regression on training data
logreg.fit(train_X, train_Y)

# Predict churn labels on testing data
pred_test_Y = logreg.predict(test_X)

# Calculate accuracy score on testing data
test_accuracy = accuracy_score(test_Y, pred_test_Y)

# Print test accuracy score rounded to 4 decimals
print('Test accuracy:', round(test_accuracy, 4))
```

## Fit logistic regression with L1 regularization
```python

```

## Identify optimal L1 penalty coefficient
```python

```

## Predict churn with decision trees
```python

```

## Fit decision tree model
```python

```

## Identify optimal tree depth
```python

```

## Identify and interpret churn drivers
```python

```

## Explore logistic regression coefficients
```python

```

## Break down decision tree rules
```python

```



# 3. Customer Lifetime Value (CLV) prediction
## Customer Lifetime Value (CLV) basics
```python

```

## Build retention and churn tables
```python

```

## Explore retention and churn
```python

```

## Calculating and projecting CLV
```python

```

## Calculate basic CLV
```python

```

## Calculate granular CLV
```python

```

## Calculate traditional CLV
```python

```

## Data preparation for purchase prediction
```python

```

## Build features
```python

```

## Define target variable
```python

```

## Split data to training and testing
```python

```

## Predicting customer transactions
```python

```

## Predict next month transactions
```python

```

## Measure model fit
```python

```

## Explore model coefficients
```python

```



# 4. Customer segmentation
## Customer and product segmentation basics
```python

```

## Explore customer product purchase dataset
```python

```

## Understand differences in variables
```python

```

## Data preparation for segmentation
```python

```

## Unskew the variables
```python

```

## Normalize the variables
```python

```

## Build customer and product segmentation
```python

```

## Determine the optimal number of clusters
```python

```

## Build segmentation with k-means clustering
```python

```

## Alternative segmentation with NMF
```python

```

## Visualize and interpret segmentation solutions
```python

```

## K-means segmentation averages
```python

```

## NMF segmentation averages
```python

```

## Congratulations!
```python

```

