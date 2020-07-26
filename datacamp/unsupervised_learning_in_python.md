---
title: Unsupervised Learning in Python
tags: python,machine-learning
url: https://www.datacamp.com/courses/unsupervised-learning-in-python
---

# 1. Clustering for dataset exploration
## Clustering 2D points
```python
# Import KMeans
from sklearn.cluster import KMeans

# Create a KMeans instance with 3 clusters: model
model = KMeans(n_clusters=3)

# Fit model to points
model.fit(points)

# Determine the cluster labels of new_points: labels
labels = model.predict(new_points)

# Print cluster labels of new_points
print(labels)
```

## Inspect your clustering
```python

```

## Evaluating a clustering
```python

```

## How many clusters of grain?
```python

```

## Evaluating the grain clustering
```python

```

## Transforming features for better clusterings
```python

```

## Scaling fish data for clustering
```python

```

## Clustering the fish data
```python

```

## Clustering stocks using KMeans
```python

```

## Which stocks move together?
```python

```



# 2. Visualization with hierarchical clustering and t-SNE
## Visualizing hierarchies
```python

```

## How many merges?
```python

```

## Hierarchical clustering of the grain data
```python

```

## Hierarchies of stocks
```python

```

## Cluster labels in hierarchical clustering
```python

```

## Which clusters are closest?
```python

```

## Different linkage, different hierarchical clustering!
```python

```

## Intermediate clusterings
```python

```

## Extracting the cluster labels
```python

```

## t-SNE for 2-dimensional maps
```python

```

## t-SNE visualization of grain dataset
```python

```

## A t-SNE map of the stock market
```python

```



# 3. Decorrelating your data and dimension reduction
## Visualizing the PCA transformation
```python

```

## Correlated data in nature
```python

```

## Decorrelating the grain measurements with PCA
```python

```

## Principal components
```python

```

## Intrinsic dimension
```python

```

## The first principal component
```python

```

## Variance of the PCA features
```python

```

## Intrinsic dimension of the fish data
```python

```

## Dimension reduction with PCA
```python

```

## Dimension reduction of the fish measurements
```python

```

## A tf-idf word-frequency array
```python

```

## Clustering Wikipedia part I
```python

```

## Clustering Wikipedia part II
```python

```



# 4. Discovering interpretable features
## Non-negative matrix factorization (NMF)
```python

```

## Non-negative data
```python

```

## NMF applied to Wikipedia articles
```python

```

## NMF features of the Wikipedia articles
```python

```

## NMF reconstructs samples
```python

```

## NMF learns interpretable parts
```python

```

## NMF learns topics of documents
```python

```

## Explore the LED digits dataset
```python

```

## NMF learns the parts of images
```python

```

## PCA doesn't learn parts
```python

```

## Building recommender systems using NMF
```python

```

## Which articles are similar to 'Cristiano Ronaldo'?
```python

```

## Recommend musical artists part I
```python

```

## Recommend musical artists part II
```python

```

## Final thoughts
```python

```


