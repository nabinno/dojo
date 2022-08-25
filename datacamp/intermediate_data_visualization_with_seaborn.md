---
title: Intermediate Data Visualization with Seaborn
tags: seaborn, python, statistics, data-visualization
url: https://campus.datacamp.com/courses/intermediate-data-visualization-with-seaborn
---

# 1. Seaborn Introduction
## Reading a csv file
```python
# import all modules
import pandas as pd
import seaborn as sns
import matplotlib.pyplot as plt

# Read in the DataFrame
df = pd.read_csv(grant_file)
```

## Comparing a histogram and displot
```python
##
# Display pandas histogram
df['Award_Amount'].plot.hist()
plt.show()

# Clear out the pandas histogram
plt.clf()

##
# Display a Seaborn displot
sns.displot(df['Award_Amount'])
plt.show()

# Clear the displot
plt.clf()
```

## Plot a histogram
```python
# Create a displot
sns.displot(df['Award_Amount'],
             bins=20)

# Display the plot
plt.show()
```

## Rug plot and kde shading
```python
# Create a displot of the Award Amount
sns.displot(df['Award_Amount'],
             kind='kde',
             rug=True,
             fill=True)

# Plot the results
plt.show()
```

## Create a regression plot
```python
##
# Create a regression plot of premiums vs. insurance_losses
sns.regplot(data=df, x="insurance_losses", y="premiums")

# Display the plot
plt.show()

##
# Create an lmplot of premiums vs. insurance_losses
sns.lmplot(data=df, y="premiums", x="insurance_losses")

# Display the second plot
plt.show()
```

## Plotting multiple variables
```python
# Create a regression plot using hue
sns.lmplot(data=df,
           x="insurance_losses",
           y="premiums",
           hue="Region")

# Show the results
plt.show()
```

## Facetting multiple regressions
```python
# Create a regression plot with multiple rows
sns.lmplot(data=df,
           x="insurance_losses",
           y="premiums",
           row="Region")

# Show the plot
plt.show()
```





# 2. Customizing Seaborn Plots
## Setting the default style
```python
# Plot the pandas histogram
df['fmr_2'].plot.hist()
plt.show()
plt.clf()

# Set the default seaborn style
sns.set()

# Plot the pandas histogram again
df['fmr_2'].plot.hist()
plt.show()
plt.clf()
```

## Comparing styles
```python
##
sns.set_style("dark")
sns.displot(df["fmr_2"])
plt.show()
plt.clf()

##
sns.set_style("whitegrid")
sns.displot(df["fmr_2"])
plt.show()
plt.clf()
```

## Removing spines
```python
# Set the style to white
sns.set_style('white')

# Create a regression plot
sns.lmplot(data=df,
           x='pop2010',
           y='fmr_2')

# Remove the spines
sns.despine()

# Show the plot and clear the figure
plt.show()
plt.clf()
```

## Colors in Seaborn
```python

```

## Matplotlib color codes
```python

```

## Using default palettes
```python

```

## Color Palettes
```python

```

## Creating Custom Palettes
```python

```

## Customizing with matplotlib
```python

```

## Using matplotlib axes
```python

```

## Additional plot customizations
```python

```

## Adding annotations
```python

```

## Multiple plots
```python

```




# 3. Additional Plot Types
## Categorical Plot Types
```python

```

## stripplot() and swarmplot()
```python

```

## boxplots, violinplots and boxenplots
```python

```

## barplots, pointplots and countplots
```python

```

## Regression Plots
```python

```

## Regression and residual plots
```python

```

## Regression plot parameters
```python

```

## Binning data
```python

```

## Matrix plots
```python

```

## Creating heatmaps
```python

```

## Customizing heatmaps
```python

```




# 4. Creating Plots on Data Aware Grids
## Using FacetGrid, catplot and lmplot
```python

```

## Building a FacetGrid
```python

```

## Using a catplot
```python

```

## Using a lmplot
```python

```

## Using PairGrid and pairplot
```python

```

## Building a PairGrid
```python

```

## Using a pairplot
```python

```

## Additional pairplots
```python

```

## Using JointGrid and jointplot
```python

```

## Building a JointGrid and jointplot
```python

```

## Jointplots and regression
```python

```

## Complex jointplots
```python

```

## Selecting Seaborn Plots
```python

```

