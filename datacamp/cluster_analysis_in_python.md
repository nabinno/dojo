---
title: Cluster Analysis in Python
tags: cluster-analysis, analytics, python
url: https://campus.datacamp.com/courses/cluster-analysis-in-python
---

# 1. Introduction to Clustering
## Unsupervised learning in real world
```txt
Segmentation of learners at DataCamp based on courses they complete. The training data has no labels.
```

## Pokémon sightings
```python
# Import plotting class from matplotlib library
from matplotlib import pyplot as plt

# Create a scatter plot
plt.scatter(x, y)

# Display the scatter plot
plt.show()
```

## Pokémon sightings: hierarchical clustering
```python
# Import linkage and fcluster functions
from scipy.cluster.hierarchy import linkage, fcluster

# Use the linkage() function to compute distance
Z = linkage(df, 'ward')

# Generate cluster labels
df['cluster_labels'] = fcluster(Z, 2, criterion='maxclust')

# Plot the points with seaborn
sns.scatterplot(x='x', y='y', hue='cluster_labels', data=df)
plt.show()
```

## Pokémon sightings: k-means clustering
```python
# Import kmeans and vq functions
from scipy.cluster.vq import kmeans, vq

# Compute cluster centers
centroids,_ = kmeans(df, 2)

# Assign cluster labels
df['cluster_labels'], _ = vq(df, centroids)

# Plot the points with seaborn
sns.scatterplot(x='x', y='y', hue='cluster_labels', data=df)
plt.show()
```

## Normalize basic list data
```python
# Import the whiten function
from scipy.cluster.vq import whiten

goals_for = [4,3,2,3,1,1,2,0,1,4]

# Use the whiten() function to standardize the data
scaled_data = whiten(goals_for)
print(scaled_data)
```

## Visualize normalized data
```python
# Plot original data
plt.plot(goals_for, label='original')

# Plot scaled data
plt.plot(scaled_data, label='scaled')

# Show the legend in the plot
plt.legend()

# Display the plot
plt.show()
```

## Normalization of small numbers
```python

```

## FIFA 18: Normalize data
```python

```




# 2. Hierarchical Clustering
## Basics of hierarchical clustering
```python

```

## Hierarchical clustering: ward method
```python

```

## Hierarchical clustering: single method
```python

```

## Hierarchical clustering: complete method
```python

```

## Visualize clusters
```python

```

## Visualize clusters with matplotlib
```python

```

## Visualize clusters with seaborn
```python

```

## How many clusters?
```python

```

## Create a dendrogram
```python

```

## How many clusters in comic con data?
```python

```

## Limitations of hierarchical clustering
```python

```

## Timing run of hierarchical clustering
```python

```

## FIFA 18: exploring defenders
```python

```




# 3. K-Means Clustering
## Basics of k-means clustering
```python

```

## K-means clustering: first exercise
```python

```

## Runtime of k-means clustering
```python

```

## How many clusters?
```python

```

## Elbow method on distinct clusters
```python

```

## Elbow method on uniform data
```python

```

## Limitations of k-means clustering
```python

```

## Impact of seeds on distinct clusters
```python

```

## Uniform clustering patterns
```python

```

## FIFA 18: defenders revisited
```python

```




# 4. Clustering in Real World
## Dominant colors in images
```python

```

## Extract RGB values from image
```python

```

## How many dominant colors?
```python

```

## Display dominant colors
```python

```

## Document clustering
```python

```

## TF-IDF of movie plots
```python

```

## Top terms in movie clusters
```python

```

## Clustering with multiple features
```python

```

## Clustering with many features
```python

```

## Basic checks on clusters
```python

```

## FIFA 18: what makes a complete player?
```python

```

## Farewell!
```python

```




