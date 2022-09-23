---
title: Introduction to Data Visualization with Matplotlib
tags: matplotlib, python, statistics, data-visualization
url: https://campus.datacamp.com/courses/introduction-to-data-visualization-with-matplotlib
---

# 1. Introduction to Matplotlib
## Using the matplotlib.pyplot interface
```python
# Import the matplotlib.pyplot submodule and name it plt
import matplotlib.pyplot as plt

# Create a Figure and an Axes with plt.subplots
fig, ax = plt.subplots()

# Call the show function to show the result
plt.show()
```

## Adding data to an Axes object
```python
# Import the matplotlib.pyplot submodule and name it plt
import matplotlib.pyplot as plt

# Create a Figure and an Axes with plt.subplots
fig, ax = plt.subplots()

# Plot MLY-PRCP-NORMAL from seattle_weather against the MONTH
ax.plot(seattle_weather["MONTH"], seattle_weather["MLY-PRCP-NORMAL"])

# Plot MLY-PRCP-NORMAL from austin_weather against MONTH
ax.plot(austin_weather["MONTH"], austin_weather["MLY-PRCP-NORMAL"])

# Call the show function
plt.show()
```

## Customizing data appearance
```python
# Plot Seattle data, setting data appearance
ax.plot(seattle_weather["MONTH"], seattle_weather["MLY-PRCP-NORMAL"], color='b', marker='o', linestyle='--')

# Plot Austin data, setting data appearance
ax.plot(austin_weather["MONTH"], austin_weather["MLY-PRCP-NORMAL"], color='r', marker='v', linestyle='--')

# Call show to display the resulting plot
plt.show()
```

## Customizing axis labels and adding titles
```python
ax.plot(seattle_weather["MONTH"], seattle_weather["MLY-PRCP-NORMAL"])
ax.plot(austin_weather["MONTH"], austin_weather["MLY-PRCP-NORMAL"])

# Customize the x-axis label
ax.set_xlabel("Time (months)")

# Customize the y-axis label
ax.set_ylabel("Precipitation (inches)")

# Add the title
ax.set_title("Weather patterns in Austin and Seattle")

# Display the figure
plt.show()
```

## Creating a grid of subplots
```python
fig, ax = plt.subplots(3, 2)
plt.show()
```

## Creating small multiples with plt.subplots
```python

```

## Small multiples with shared y axis
```python

```




# 2. Plotting time-series
## Plotting time-series data
```python

```

## Read data with a time index
```python

```

## Plot time-series data
```python

```

## Using a time index to zoom in
```python

```

## Plotting time-series with different variables
```python

```

## Plotting two variables
```python

```

## Defining a function that plots time-series data
```python

```

## Using a plotting function
```python

```

## Annotating time-series data
```python

```

## Annotating a plot of time-series data
```python

```

## Plotting time-series: putting it all together
```python

```




# 3. Quantitative comparisons and statistical visualizations
## Quantitative comparisons: bar-charts
```python

```

## Bar chart
```python

```

## Stacked bar chart
```python

```

## Quantitative comparisons: histograms
```python

```

## Creating histograms
```python

```

## "Step" histogram
```python

```

## Statistical plotting
```python

```

## Adding error-bars to a bar chart
```python

```

## Adding error-bars to a plot
```python

```

## Creating boxplots
```python

```

## Quantitative comparisons: scatter plots
```python

```

## Simple scatter plot
```python

```

## Encoding time by color
```python

```




# 4. Sharing visualizations with others
## Preparing your figures to share with others
```python

```

## Selecting a style for printing
```python

```

## Switching between styles
```python

```

## Saving your visualizations
```python

```

## Saving a file several times
```python

```

## Save a figure with different sizes
```python

```

## Automating figures from data
```python

```

## Unique values of a column
```python

```

## Automate your visualization
```python

```

## Where to go next
```python

```

