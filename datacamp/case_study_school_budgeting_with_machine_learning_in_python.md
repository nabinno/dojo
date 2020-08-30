---
title: Case Study: School Budgeting with Machine Learning in Python
tags: python, machine-learning
url: https://www.datacamp.com/courses/case-study-school-budgeting-with-machine-learning-in-python
---

# 1. Exploring the raw data
## Summarizing the data
```python
# Print the summary statistics
print(df.describe())

# Import matplotlib.pyplot as plt
import matplotlib.pyplot as plt

# Create the histogram
plt.hist(df['FTE'].dropna())

# Add title and labels
plt.title('Distribution of %full-time \n employee works')
plt.xlabel('% of full-time')
plt.ylabel('num employees')

# Display the histogram
plt.show()
```

## Encode the labels as categorical variables
```python
# Define the lambda function: categorize_label
categorize_label = lambda x: x.astype('category')

# Convert df[LABELS] to a categorical type
df[LABELS] = df[LABELS].apply(categorize_label, axis=0)

# Print the converted dtypes
print(df[LABELS].dtypes)
```

## Counting unique labels
```python
# Import matplotlib.pyplot
import matplotlib.pyplot as plt

# Calculate number of unique values for each label: num_unique_labels
num_unique_labels = df[LABELS].apply(pd.Series.nunique)

# Plot number of unique values for each label
num_unique_labels.plot(kind='bar')

# Label the axes
plt.xlabel('Labels')
plt.ylabel('Number of unique values')

# Display the plot
plt.show()
```

## Computing log loss with NumPy
```python
# Compute and print log loss for 1st case
correct_confident_loss = compute_log_loss(correct_confident, actual_labels)
print("Log loss, correct and confident: {}".format(correct_confident_loss)) 

# Compute log loss for 2nd case
correct_not_confident_loss = compute_log_loss(correct_not_confident, actual_labels)
print("Log loss, correct and not confident: {}".format(correct_not_confident_loss)) 

# Compute and print log loss for 3rd case
wrong_not_confident_loss = compute_log_loss(wrong_not_confident, actual_labels)
print("Log loss, wrong and not confident: {}".format(wrong_not_confident_loss)) 

# Compute and print log loss for 4th case
wrong_confident_loss = compute_log_loss(wrong_confident, actual_labels)
print("Log loss, wrong and confident: {}".format(wrong_confident_loss)) 

# Compute and print log loss for actual labels
actual_labels_loss = compute_log_loss(actual_labels, actual_labels)
print("Log loss, actual labels: {}".format(actual_labels_loss)) 
```



# 2. Creating a simple first model
## Setting up a train-test split in scikit-learn
```python
# Create the new DataFrame: numeric_data_only
numeric_data_only = df[NUMERIC_COLUMNS].fillna(-1000)

# Get labels and convert to dummy variables: label_dummies
label_dummies = pd.get_dummies(df[LABELS])

# Create training and test sets
X_train, X_test, y_train, y_test = multilabel_train_test_split(numeric_data_only,
                                                               label_dummies,
                                                               size=0.2,
                                                               seed=123)

# Print the info
print("X_train info:")
print(X_train.info())
print("\nX_test info:") 
print(X_test.info())
print("\ny_train info:") 
print(y_train.info())
print("\ny_test info:") 
print(y_test.info()) 
```

## Training a model
```python
# Import classifiers
from sklearn.linear_model import LogisticRegression
from sklearn.multiclass import OneVsRestClassifier

# Create the DataFrame: numeric_data_only
numeric_data_only = df[NUMERIC_COLUMNS].fillna(-1000)

# Get labels and convert to dummy variables: label_dummies
label_dummies = pd.get_dummies(df[LABELS])

# Create training and test sets
X_train, X_test, y_train, y_test = multilabel_train_test_split(numeric_data_only,
                                                               label_dummies,
                                                               size=0.2, 
                                                               seed=123)

# Instantiate the classifier: clf
clf = OneVsRestClassifier(LogisticRegression())

# Fit the classifier to the training data
clf.fit(X_train, y_train)

# Print the accuracy
print("Accuracy: {}".format(clf.score(X_test, y_test)))
```

## Use your model to predict values on holdout data
```python
# Instantiate the classifier: clf
clf = OneVsRestClassifier(LogisticRegression())

# Fit it to the training data
clf.fit(X_train, y_train)

# Load the holdout data: holdout
holdout = pd.read_csv('HoldoutData.csv', index_col=0)

# Generate predictions: predictions
predictions = clf.predict_proba(holdout.fillna(-1000)[NUMERIC_COLUMNS])
```

## Writing out your results to a csv for submission
```python
# Generate predictions: predictions
predictions = clf.predict_proba(holdout[NUMERIC_COLUMNS].fillna(-1000))

# Format predictions in DataFrame: prediction_df
prediction_df = pd.DataFrame(columns=pd.get_dummies(df[LABELS]).columns,
                             index=holdout.index,
                             data=predictions)


# Save prediction_df to csv
prediction_df.to_csv('predictions.csv')

# Submit the predictions for scoring: score
score = score_submission(pred_path='predictions.csv')

# Print score
print('Your model, trained with numeric data only, yields logloss score: {}'.format(score))
```

## Tokenizing text
```python

```

## Testing your NLP credentials with n-grams
```python

```

## Representing text numerically
```python

```

## Creating a bag-of-words in scikit-learn
```python

```

## Combining text columns for tokenization
```python

```

## What's in a token?
```python

```



# 3. Improving your model
## Pipelines, feature & text preprocessing
```python

```

## Instantiate pipeline
```python

```

## Preprocessing numeric features
```python

```

## Text features and feature unions
```python

```

## Preprocessing text features
```python

```

## Multiple types of processing: FunctionTransformer
```python

```

## Multiple types of processing: FeatureUnion
```python

```

## Choosing a classification model
```python

```

## Using FunctionTransformer on the main dataset
```python

```

## Add a model to the pipeline
```python

```

## Try a different class of model
```python

```

## Can you adjust the model or parameters to improve accuracy?
```python

```



# 4. Learning from the experts
## Learning from the expert: processing
```python

```

## How many tokens?
```python

```

## Deciding what's a word
```python

```

## N-gram range in scikit-learn
```python

```

## Learning from the expert: a stats trick
```python

```

## Which models of the data include interaction terms?
```python

```

## Implement interaction modeling in scikit-learn
```python

```

## Learning from the expert: the winning model
```python

```

## Why is hashing a useful trick?
```python

```

## Implementing the hashing trick in scikit-learn
```python

```

## Build the winning model
```python

```

## What tactics got the winner the best score?
```python

```

## Next steps and the social impact of your work
```python

```



