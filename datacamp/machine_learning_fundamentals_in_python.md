---
title: Machine Learning Fundamentals in Python
tags: machine-learning, python
url: https://assessment.datacamp.com/machine-learning-fundamentals-with-python
---

# 1. Assessment
## RandomForestClassifier
```python
from sklearn.ensemble import RandomForestClassifier
from sklearn.metrics import accuracy_score

model = RandomForestClassifier(n_estimators=300, max_depth=1, random_state=1)
model.fit(X_train, y_train)

y_pred = model.predict(X_test)
accuracy_score(y_test, y_pred)
```

## create train
```python
from sklearn import model_selection

X_train, X_test, y_train, y_test = model_selection.train_test_split(X, y, split=0.8, random_state=42)
print("X_train shape: ", X_train.shape)
print("X_test shape: ", X_test.shape)
print("y_train shape: ",y_train.shape)
print("y_test shape: ",y_test.shape)
```

## pandas
```python
import pandas as pd

df['Age'] = pd.cut(df['Age'], bins=[20, 25, 35, 45, 60])

df
```

## GradientBoostingClassifier
```python
from sklearn.ensemble import GradientBoostingClassifier
from sklearn.metrics import accuracy_score

model = GradientBoostingClassifier(n_estimators=300, max_depth=1, random_state=1)
model.fit(X_train, y_train)

y_pred = model.predict(X_test)
accuracy_score(y_test, y_pred)
```

## kmeans
```python
from scipy.cluster.vq import kmeans, vq

kmeans(
	df[['x_scaled', 'y_scaled']],
	2
)
```

## GradientBoostingClassifier (2)
```python
from sklearn.ensemble import GradientBoostingClassifier
from sklearn.metrics import accuracy_score

model = GradientBoostingClassifier(n_estimators=300, max_depth=1, random_state=1)
model.fit(X_train, y_train)

y_pred = model.predict(X_test)
accuracy_score(y_test, y_pred)
```

## PowerTransformer
```python
from sklearn.preprocessing import PowerTransformer

log = PowerTransformer(method='box-cox')
df['log_x'] = log.fit_transform(df[['x']])
df['log_x'].head()
```

## StandardScaler
```python
from sklearn.preprocessing import StandardScaler

scaler = StandardScaler()
df_scaled = pd.DataFrame(scaler.fit_transform(df), columns=df.columns)
df_scaled.head()
```

## Lasso
```python
import numpy as np
from sklearn.metrics import mean_squared_error, r2_score
from sklearn.linear_model import Lasso

lasso_model = Lasso(alpha=0.01)
lasso_model.fit(X_train, y_train)
lasso_predictions = lasso_model.predict(X_test)
print("RMSE: ", np.sqrt(mean_squared_error(y_test, lasso_predictions)))
```

## RandomForestClassifier
```python
from sklearn.ensemble import RandomForestClassifier
from sklearn.metrics import accuracy_score

model = RandomForestClassifier(n_estimators=10, random_state=1)
model.fit(X_train, y_train)

y_pred = model.predict(X_test)
accuracy_score(y_test, y_pred)
```

## sklearn.decomposition.predict
```python
from sklearn.decomposition import PCA

pca = PCA(n_components=2)
pca.fit(scaled_samples)

pca_features = pca.transform(scaled_samples)
print(pca_features.shape)
```

## numpy
```python
import numpy as np

np.mean(x)
```

## sklearn.linear_model.LinearRegression
```python
import numpy as np
import pandas as pd
from sklearn.linear_model import LinearRegression

reg = LinearRegression()
reg.fit(x, y)

print("Regression coefficients: {}".format(reg.coef_))
print("Regression intercept: {}".format(reg.intercept_))
```

## pandas
```python
import pandas as pd

employee_churn.describe()
```


# 2. Assessment
## sklearn.linear_model.LogisticRegression
```python
from sklearn.linear_model import LogisticRegression

model = LogisticRegression(random_state=1)
model.fit(X_train, y_train)

y_pred = model.predict(X_test)
model.score(X_test, y_test)
```

## Linear regression
An algorithm that describes a continuous response variable as a function of one or more predictor variables.

## K-means clustering
K-means clustering groups data into relatively distinct groups by using a pre-determined number of clusters and iterating cluster assignments.

## numpy.mean
```python
import numpy as np

np.mean(x)
```

## Linear regression
A retail company that wants to predict their total sales to understand the impact of online promotions.

## Benefits of regularization
Regularization helps to address the problem of over-fitting training data by restricting the model's coefficients.

## sklearn.ensemble.RandomForestClassifier
```python
from sklearn.ensemble import RandomForestClassifier
from sklearn.metrics import accuracy_score

model = RandomForestClassifier(n_estimators=10, random_state=1)
model.fit(X_train, y_train)

y_pred = model.predict(X_test)
accuracy_score(y_test, y_pred)
```

## Classification model
Which of the following metrics would not be used when assessing the performance of a classification model?
- Median absolute error

## statsmodels.formula.api.glm
```python
import statsmodels.api as sm
from statsmodels.formula.api import glm

model = glm(
	'goal ~ player', 
	data = score,
	family = sm.families.Poisson()
).fit()

model.predict(test)
```

## sklearn.ensemble.RandomForestClassifier (2)
```python
from sklearn.ensemble import RandomForestClassifier
from sklearn.metrics import accuracy_score

model = RandomForestClassifier(n_estimators=300, max_depth=1, random_state=1)
model.fit(X_train, y_train)

y_pred = model.predict(X_test)
accuracy_score(y_test, y_pred)
```

## pandas
```python
import pandas as pd

pd.get_dummies(df, columns=['Animal'])
```

## sklearn.linear_model.LinearRegression
```python
import numpy as np
import pandas as pd
from sklearn.linear_model import LinearRegression

reg = LinearRegression()
reg.fit(x_train, y_train)

reg.predict(x_test)
```

## plot
Consider the following plot. Which of the statements below is not true?
- There are outliers in the data, which may be impacting results

## Transformation to a variable
Why might you consider applying a transformation to a variable?
- Because there is an outlier

## Unsupervised learning problem
Which one of the following statements describes an unsupervised learning problem?
- A machine learning problem where we seek to understand whether observations fit into distinct groups based on their similarities.


# 3. Assessment
## sklearn.linear_model.LogisticRegression
```python
from sklearn.linear_model import LogisticRegression

model = LogisticRegression(random_state=1)
model.fit(X_train, y_train)

model.score(X_test, y_test)
```

## sklearn.ensemble.RandomForestClassifier
```python
from sklearn.ensemble import RandomForestClassifier
from sklearn.metrics import accuracy_score

model = RandomForestClassifier(n_estimators=300, max_depth=1, random_state=1)
model.fit(X_train, y_train)

y_pred = model.predict(X_test)
accuracy_score(y_test, y_pred)
```

## Generalized linear model (GLM)
In which of the following situations would you recommend the use of a generalized linear model(GLM) with a non-Gaussian distribution rather than a simple linear model?
- To predict the number of cyclists that cross a bridge per day.

## statsmodels.formula.api.glm
```python
import statsmodels.api as sm
from statsmodels.formula.api import glm

model = glm(
	'goal ~ player', 
	data = score,
	family = sm.families.Poisson()
).fit()

model.params
```

## scipy.cluster.hierarchy.linkage
```python
from scipy.cluster.hierarchy import linkage, fcluster

distance_matrix = linkage(
	df[['x_scaled', 'y_scaled']],
	method = 'complete',
	metric = 'euclidean'
)

fcluster(
	distance_matrix,
	t = 2,
	criterion='maxclust'
)
```

## Bias-variance trade-off
When referring to the bias-variance trade-off, what does variance stand for and how can you diagnose it?
- Variance is when the model follows the training data too closely, resulting in high training error, and low testing errors.

## Not true regarding hierarchical clustering linkage methods
Given two different clusters, which of the following statements is not true regarding hierarchical clustering linkage methods?
- The complete method calculates pairwise similarity between all observations in both clusters. The sum of these distances is then used as the distance between the clusters.

## seaborn.boxplot
```python
import seaborn as sns
import matplotlib.pyplot as plt

sns.boxplot(x='Attrition', y='Age', data=churn)
plt.show()
```

## sklearn.metrics.auc
```python
from sklearn.linear_model import LogisticRegression
from sklearn import metrics

clf=LogisticRegression(solver='newton-cg', random_state=42)
clf.fit(X_train,y_train)

y_pred=clf.predict(X_test)

metrics.recall_score(y_test, y_pred)
```

## sklaern.preprocessing.MinMaxScaler
```python
from sklearn.preprocessing import MinMaxScaler

min_max = MinMaxScaler()
df_scaled = pd.DataFrame(min_max.fit_transform(df), columns=df.columns)
df_scaled.head()
```

## sklearn.metrics
```python
from sklearn.model_selection import train_test_split
from sklearn import metrics

X_train,X_test,y_train,y_test=train_test_split(X,y,test_size=0.25, random_state=42)
clf.fit(X_train,y_train)

y_pred=clf.predict(X_test)

metrics.accuracy_score(y_test, y_pred)
```

## sklearn.linear_model.LogisticRegression
```python
from sklearn.linear_model import LogisticRegression

model = LogisticRegression(random_state=1)
model.fit(X_train, y_train)

y_pred = model.predict(X_test)
model.score(X_test, y_test)
```

## sklearn.model_selection
```python
from sklearn import model_selection

X_train, X_test, y_train, y_test = model_selection.train_test_split(X, y, train_size=0.8, random_state=42)
print("X_train shape: ", X_train.shape)
print("X_test shape: ", X_test.shape)
print("y_train shape: ",y_train.shape)
print("y_test shape: ",y_test.shape)
```

## sklearn.model_selection (2)
```python
from sklearn import model_selection

scores = model_selection.cross_val_score(knn, X, y, cv=5, scoring='accuracy')
print(scores)
```

## sklearn.ensemble.GradientBoostingClassifier
```python
from sklearn.ensemble import GradientBoostingClassifier
from sklearn.metrics import accuracy_score

model = GradientBoostingClassifier(n_estimators=300, max_depth=1, random_state=1)
model.fit(X_train, y_train)

y_pred = model.predict(X_test)
accuracy_score(y_test, y_pred)
```

# 4. Assessment
## Random forest
How does a random forest improve upon a decision tree?
- have been trained on different bootstrap samples of the training set.

## sklearn.tree.DecisionTreeClassifier
```python
from sklearn.tree import DecisionTreeClassifier
from sklearn.metrics import accuracy_score

model = DecisionTreeClassifier(max_depth=4, random_state=1)
model.fit(X_train, y_train)

y_pred = model.predict(X_test)

accuracy_score(y_test, y_pred)
```

## Decision tree
In which of the following situations might you want to use a decision tree?
- When model output should be easily interpretable.


## sklearn.linear_model.LogisticRegression
```python
from sklearn.linear_model import LogisticRegression

model = LogisticRegression(random_state=1)
model.fit(X_train, y_train)

y_pred = model.predict(X_test)
model.score(X_test, y_test)
```
