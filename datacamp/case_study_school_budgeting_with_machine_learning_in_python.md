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

## Exploring datatypes in pandas
```python

```

## Encode the labels as categorical variables
```python

```

## Counting unique labels
```python

```

## How do we measure success?
```python

```

## Penalizing highly confident wrong answers
```python

```

## Computing log loss with NumPy
```python

```



# 2. Creating a simple first model
## It's time to build a model
```python

```

## Setting up a train-test split in scikit-learn
```python

```

## Training a model
```python

```

## Making predictions
```python

```

## Use your model to predict values on holdout data
```python

```

## Writing out your results to a csv for submission
```python

```

## A very brief introduction to NLP
```python

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



