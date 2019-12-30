---
title: Image Processing in Python
tags: python,image-processing
url: https://www.datacamp.com/courses/image-processing-in-python
---

# 1. Introducing Image Processing and scikit-image
## Is this gray or full of color?
```python
np.shape(coffee_image)
np.shape(coins_image)
```

## RGB to grayscale
```python
# Import the modules from skimage
from skimage import data, color

# Load the rocket image
rocket = data.rocket()

# Convert the image to grayscale
gray_scaled_rocket = color.rgb2grey(rocket)

# Show the original image
show_image(rocket, 'Original RGB image')

# Show the grayscale image
show_image(gray_scaled_rocket, 'Grayscale image')
```

## Flipping out
```python
# Flip the image vertically
seville_vertical_flip = np.flipud(flipped_seville)

# Flip the image horizontally
seville_horizontal_flip = np.fliplr(seville_vertical_flip)

# Show the resulting image
show_image(seville_horizontal_flip, 'Seville')
```

## Histograms
```python
# Obtain the red channel
red_channel = image[:, :, 0]

# Plot the red histogram with bins in a range of 256
plt.hist(red_channel.ravel(), bins=256)

# Set title and show
plt.title('Red Histogram')
plt.show()
```

## Apply global thresholding
```python
# Import the otsu threshold function
from skimage.filters import threshold_otsu

# Make the image grayscale using rgb2gray
chess_pieces_image_gray = rgb2gray(chess_pieces_image)

# Obtain the optimal threshold value with otsu
thresh = threshold_otsu(chess_pieces_image_gray)

# Apply thresholding to the image
binary = chess_pieces_image_gray > thresh

# Show the image
show_image(binary, 'Binary image')
```

## When the background isn't that obvious
```python
# Import the otsu threshold function
from skimage.filters import threshold_otsu

# Obtain the optimal otsu global thresh value
global_thresh = threshold_otsu(page_image)

# Obtain the binary image by applying global thresholding
binary_global = page_image > global_thresh

# Show the binary image obtained
show_image(binary_global, 'Global thresholding')

##
# Import the local threshold function
from skimage.filters import threshold_local

# Set the block size to 35
block_size = 35

# Obtain the optimal local thresholding
local_thresh = threshold_local(page_image, block_size, offset=10)

# Obtain the binary image by applying local thresholding
binary_local = page_image > local_thresh

# Show the binary image
show_image(binary_local, 'Local thresholding')
```

## Trying other methods
```python
# Import the try all function
from skimage.filters import try_all_threshold

# Import the rgb to gray convertor function 
from skimage.color import rgb2gray

# Turn the fruits image to grayscale
grayscale = rgb2gray(fruits_image)

# Use the try all method on the grayscale image
fig, ax = try_all_threshold(grayscale, verbose=False)

# Show the resulting plots
plt.show()
```

## Apply thresholding
```python
# Import threshold and gray convertor functions
from skimage.filters import threshold_otsu
from skimage.color import rgb2gray

# Turn the image grayscale
tools_image = rgb2gray(tools_image)

# Obtain the optimal thresh
thresh = threshold_otsu(tools_image)

# Obtain the binary image by applying thresholding
binary_image = tools_image > thresh

# Show the resulting binary image
show_image(binary_image, 'Binarized image')
```

# 2. Filters, Contrast, Transformation and Morphology
## Jump into filtering
```python

```

## Edge detection
```python

```

## Blurring to reduce noise
```python

```

## Contrast enhancement
```python

```

## What's the contrast of this image?
```python

```

## Medical images
```python

```

## Aerial image
```python

```

## Let's add some impact and contrast
```python

```

## Transformations
```python

```

## Aliasing, rotating and rescaling
```python

```

## Enlarging images
```python

```

## Proportionally resizing
```python

```

## Morphology
```python

```

## Handwritten letters
```python

```

## Improving thresholded image
```python

```

# 3. Image restoration, Noise, Segmentation and Contours
## Image restoration
```python

```

## Let's restore a damaged image
```python

```

## Removing logos
```python

```

## Noise
```python

```

## Let's make some noise!
```python

```

## Reducing noise
```python

```

## Reducing noise while preserving edges
```python

```

## Superpixels & segmentation
```python

```

## Number of pixels
```python

```

## Superpixel segmentation
```python

```

## Finding contours
```python

```

## Contouring shapes
```python

```

## Find contours of an image that is not binary
```python

```

## Count the dots in a dice's image
```python

```

# 4. Advanced Operations, Detecting Faces and Features
## Finding the edges with Canny
```python

```

## Edges
```python

```

## Less edgy
```python

```

## Right around the corner
```python

```

## Perspective
```python

```

## Less corners
```python

```

## Face detection
```python

```

## Is someone there?
```python

```

## Multiple faces
```python

```

## Segmentation and face detection
```python

```

## Real-world applications
```python

```

## Privacy protection
```python

```

## Help Sally restore her graduation photo
```python

```

## Amazing work!
```python

```


