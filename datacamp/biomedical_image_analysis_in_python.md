---
title: Biomedical Image Analysis in Python
tags: python,image-analysis,image-processing,physiology
url: https://www.datacamp.com/courses/biomedical-image-analysis-in-python
---

# 1. Exploration
## Load images
```python
# Import ImageIO
import imageio

# Load "chest-220.dcm"
im = imageio.imread("chest-220.dcm")

# Print image attributes
print('Image type:', type(im))
print('Shape of image array:', im.shape)
```

## Metadata
```python
# Import ImageIO
import imageio
im = imageio.imread('chest-220.dcm')

# Print the available metadata fields
print(im.meta.keys())
```

## Plot images
```python
# Import ImageIO and PyPlot 
import imageio
import matplotlib.pyplot as plt

##
# Read in "chest-220.dcm"
im = imageio.imread("chest-220.dcm")

# Draw the image in grayscale
plt.imshow(im, cmap='gray')

# Render the image
plt.show()

##
# Draw the image with greater contrast
plt.imshow(im, cmap='gray', vmin=-200, vmax=200)

# Render the image
plt.show()

##
# Remove axis ticks and labels
plt.axis('off')

# Render the image
plt.show()
```

## Stack images
```python
# Import ImageIO and NumPy
import imageio
import numpy as np

# Read in each 2D image
im1 = imageio.imread('chest-220.dcm')
im2 = imageio.imread('chest-221.dcm')
im3 = imageio.imread('chest-222.dcm')

# Stack images into a volume
vol = np.stack([im1, im2, im3])
print('Volume dimensions:', vol.shape)
```

## Load volumes
```python
# Import ImageIO
import imageio

# Load the "tcia-chest-ct" directory
vol = imageio.volread('tcia-chest-ct/')

# Print image attributes
print('Available metadata:', vol.meta.keys())
print('Shape of image array:', vol.shape)
```

## Generate subplots
```python
# Import PyPlot
import matplotlib.pyplot as plt

# Initialize figure and axes grid
fig, axes = plt.subplots(nrows=2, ncols=1)

# Draw an image on each subplot
axes[0].imshow(im1, cmap='gray')
axes[1].imshow(im2, cmap='gray')

# Remove ticks/labels and render
axes[0].axis('off')
axes[1].axis('off')
plt.show()
```

## Slice 3D images
```python
# Plot the images on a subplots array 
fig, axes = plt.subplots(nrows=1, ncols=4)

# Loop through subplots and draw image
for ii in range(4):
    im = vol[ii * 40]
    axes[ii].imshow(im, cmap='gray')
    axes[ii].axis('off')

# Render the figure
plt.show()
```

## Plot other views
```python
# Select frame from "vol"
im1 = vol[:, 256, :]
im2 = vol[:, :, 256]

# Compute aspect ratios
d0, d1, d2 = vol.meta['sampling']
asp1 = d0 / d2
asp2 = d0 / d1

# Plot the images on a subplots array 
fig, axes = plt.subplots(nrows=2, ncols=1)
axes[0].imshow(im1, cmap='gray', aspect=asp1)
axes[1].imshow(im2, cmap='gray', aspect=asp2)
plt.show()
```

# 2. Masks and Filters
## Intensity
```python

```

## Histograms
```python

```

## Masks
```python

```

## Create a mask
```python

```

## Apply a mask
```python

```

## Tune a mask
```python

```

## Filters
```python

```

## Filter convolutions
```python

```

## Filter functions
```python

```

## Smoothing
```python

```

## Feature detection
```python

```

## Detect edges (1)
```python

```

## Detect edges (2)
```python

```

# 3. Measurement
## Objects and labels
```python

```

## Segment the heart
```python

```

## Select objects
```python

```

## Extract objects
```python

```

## Measuring intensity
```python

```

## Measure variance
```python

```

## Separate histograms
```python

```

## Measuring morphology
```python

```

## Calculate volume
```python

```

## Calculate distance
```python

```

## Pinpoint center of mass
```python

```

## Measuring in time
```python

```

## Summarize the time series
```python

```

## Measure ejection fraction
```python

```

# 4. Image Comparison
## Spatial transformations
```python

```

## Translations
```python

```

## Rotations
```python

```

## Affine transform
```python

```

## Resampling and interpolation
```python

```

## Resampling
```python

```

## Interpolation
```python

```

## Comparing images
```python

```

## Mean absolute error
```python

```

## Intersection of the union
```python

```

## Normalizing measurements
```python

```

## Identifying potential confounds
```python

```

## Testing group differences
```python

```

## Normalizing metrics
```python

```


