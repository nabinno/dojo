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
# Import pyplot
import matplotlib.pyplot as plt

# Assign the columns of new_points: xs and ys
xs = new_points[:,0]
ys = new_points[:,1]

# Make a scatter plot of xs and ys, using labels to define the colors
plt.scatter(xs, ys, c=labels, alpha=0.5)

# Assign the cluster centers: centroids
centroids = model.cluster_centers_

# Assign the columns of centroids: centroids_x, centroids_y
centroids_x = centroids[:,0]
centroids_y = centroids[:,1]

# Make a scatter plot of centroids_x and centroids_y
plt.scatter(centroids_x, centroids_y, marker='D', s=50)
plt.show()
```

## How many clusters of grain?
```python
ks = range(1, 6)
inertias = []

for k in ks:
    # Create a KMeans instance with k clusters: model
    model = KMeans(n_clusters=k)
    
    # Fit model to samples
    model.fit(samples)
    
    # Append the inertia to the list of inertias
    inertias.append(model.inertia_)
    
# Plot ks vs inertias
plt.plot(ks, inertias, '-o')
plt.xlabel('number of clusters, k')
plt.ylabel('inertia')
plt.xticks(ks)
plt.show()
```

## Evaluating the grain clustering
```python
# Create a KMeans model with 3 clusters: model
model = KMeans(n_clusters=3)

# Use fit_predict to fit model and obtain cluster labels: labels
labels = model.fit_predict(samples)

# Create a DataFrame with labels and varieties as columns: df
df = pd.DataFrame({'labels': labels, 'varieties': varieties})

# Create crosstab: ct
ct = pd.crosstab(df['labels'], df['varieties'])

# Display ct
print(ct)
```

## Scaling fish data for clustering
```python
# Perform the necessary imports
from sklearn.pipeline import make_pipeline
from sklearn.preprocessing import StandardScaler
from sklearn.cluster import KMeans

# Create scaler: scaler
scaler = StandardScaler()

# Create KMeans instance: kmeans
kmeans = KMeans(n_clusters=4)

# Create pipeline: pipeline
pipeline = make_pipeline(scaler, kmeans)
```

## Clustering the fish data
```python
# Import pandas
import pandas as pd

# Fit the pipeline to samples
pipeline.fit(samples)

# Calculate the cluster labels: labels
labels = pipeline.predict(samples)

# Create a DataFrame with labels and species as columns: df
df = pd.DataFrame({'labels': labels, 'species': species})

# Create crosstab: ct
ct = pd.crosstab(df['labels'], df['species'])

# Display ct
print(ct)
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


