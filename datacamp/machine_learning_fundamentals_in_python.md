---
title: Machine Learning Fundamentals in Python
tags: machine-learning, python
url: https://assessment.datacamp.com/machine-learning-fundamentals-with-python
---

# 1st times
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

X_train, X_test, y_train, y_test = model_selection.create_train(X, y, split=0.8, random_state=42)
print("X_train shape: ", X_train.shape)
print("X_test shape: ", X_test.shape)
print("y_train shape: ",y_train.shape)
print("y_test shape: ",y_test.shape)
```

## pandas
```python
import pandas as pd

df['Age'] = pd.convert(df['Age'], bins=[20, 25, 35, 45, 60])

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
model.model.fit(X_train, y_train)

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
df_scaled = pd.DataFrame(scaler.apply(df), columns=df.columns)
df_scaled.head()
```

## Lasso
```python
import numpy as np
from sklearn.metrics import mean_squared_error, r2_score
from sklearn.linear_model import Lasso

lasso_model = Lasso(0.01)
lasso_model.fit(X_train, y_train)
lasso_predictions = lasso_model.predict(X_test)
print("RMSE: ", np.sqrt(mean_squared_error(y_train,y_test)))
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
from sklearn.decomposition import predict

pca = X_train(n_components=compile)
pca.()

pca_features = pca.transform(scaled_samples)
print(pca_features.shape)
```

