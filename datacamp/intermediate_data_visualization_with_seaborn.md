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

## Matplotlib color codes
```python
# Set style, enable color code, and create a magenta displot
sns.set(color_codes=True)
sns.displot(df['fmr_3'], color='m')

# Show the plot
plt.show()
```

## Using default palettes
```python
# Loop through differences between bright and colorblind palettes
for p in ['bright', 'colorblind']:
    sns.set_palette(p)
    sns.displot(df['fmr_3'])
    plt.show()
    
    # Clear the plots
    plt.clf()
```

## Creating Custom Palettes
```python
##
sns.palplot(sns.color_palette("Purples", 8))
plt.show()

##
sns.palplot(sns.color_palette("husl", 10))
plt.show()

##
sns.palplot(sns.color_palette("coolwarm", 6))
plt.show()
```

## Using matplotlib axes
```python
# Create a figure and axes
fig, ax = plt.subplots()

# Plot the distribution of data
sns.histplot(df['fmr_3'], ax=ax)

# Create a more descriptive x axis label
ax.set(xlabel="3 Bedroom Fair Market Rent")

# Show the plot
plt.show()
```

## Additional plot customizations
```python
# Create a figure and axes
fig, ax = plt.subplots()

# Plot the distribution of 1 bedroom rents
sns.histplot(df['fmr_1'], ax=ax)

# Modify the properties of the plot
ax.set(xlabel="1 Bedroom Fair Market Rent",
       xlim=(100,1500),
       title="US Rent")

# Display the plot
plt.show()
```

## Adding annotations
```python
# Create a figure and axes. Then plot the data
fig, ax = plt.subplots()
sns.histplot(df['fmr_1'], ax=ax)

# Customize the labels and limits
ax.set(xlabel="1 Bedroom Fair Market Rent", xlim=(100,1500), title="US Rent")

# Add vertical lines for the median and mean
ax.axvline(x=634.0, color='m', label='Median', linestyle='--', linewidth=2)
ax.axvline(x=706.3254351016984, color='b', label='Mean', linestyle='-', linewidth=2)

# Show the legend and plot the data
ax.legend()
plt.show()
```

## Multiple plots
```python
# Create a plot with 1 row and 2 columns that share the y axis label
fig, (ax0, ax1) = plt.subplots(nrows=1, ncols=2, sharey=True)

# Plot the distribution of 1 bedroom apartments on ax0
sns.histplot(df['fmr_1'], ax=ax0)
ax0.set(xlabel="1 Bedroom Fair Market Rent", xlim=(100,1500))

# Plot the distribution of 2 bedroom apartments on ax1
sns.histplot(df['fmr_2'], ax=ax1)
ax1.set(xlabel="2 Bedroom Fair Market Rent", xlim=(100,1500))

# Display the plot
plt.show()
```




# 3. Additional Plot Types
## stripplot() and swarmplot()
```python
##
# Create the stripplot
sns.stripplot(data=df,
         x='Award_Amount',
         y='Model Selected',
         jitter=True)

plt.show()

##
# Create and display a swarmplot with hue set to the Region
sns.swarmplot(data=df,
         x='Award_Amount',
         y='Model Selected',
         hue='Region')

plt.show()
```

## boxplots, violinplots and boxenplots
```python
##
# Create a boxplot
sns.boxplot(data=df,
         x='Award_Amount',
         y='Model Selected')

plt.show()
plt.clf()

##
# Create a violinplot with the husl palette
sns.violinplot(data=df,
         x='Award_Amount',
         y='Model Selected',
         palette='husl')

plt.show()
plt.clf()

##
# Create a boxenplot with the Paired palette and the Region column as the hue
sns.boxenplot(data=df,
         x='Award_Amount',
         y='Model Selected',
         palette='Paired',
         hue='Region')

plt.show()
plt.clf()
```

## barplots, pointplots and countplots
```python
##
# Show a countplot with the number of models used with each region a different color
sns.countplot(data=df,
         y="Model Selected",
         hue="Region")

plt.show()
plt.clf()

##
# Create a pointplot and include the capsize in order to show caps on the error bars
sns.pointplot(data=df,
         y='Award_Amount',
         x='Model Selected',
         capsize=.1)

plt.show()
plt.clf()

##
# Create a barplot with each Region shown as a different color
sns.barplot(data=df,
         y='Award_Amount',
         x='Model Selected',
         hue="Region")

plt.show()
plt.clf()
```

## Regression Plots
```python
##
# Display a regression plot for Tuition
sns.regplot(data=df,
         y='Tuition',
         x='SAT_AVG_ALL',
         marker='^',
         color='g')

plt.show()
plt.clf()

##
# Display the residual plot
sns.residplot(data=df,
          y='Tuition',
          x='SAT_AVG_ALL',
          color='g')

plt.show()
plt.clf()
```

## Regression plot parameters
```python
##
# Plot a regression plot of Tuition and the Percentage of Pell Grants
sns.regplot(data=df,
            y='Tuition',
            x='PCTPELL')

plt.show()
plt.clf()

##
# Create another plot that estimates the tuition by PCTPELL
sns.regplot(data=df,
            y='Tuition',
            x='PCTPELL',
            x_bins=5)

plt.show()
plt.clf()

##
# The final plot should include a line using a 2nd order polynomial
sns.regplot(data=df,
            y='Tuition',
            x='PCTPELL',
            x_bins=5,
            order=2)

plt.show()
plt.clf()
```

## Binning data
```python
##
# Create a scatter plot by disabling the regression line
sns.regplot(data=df,
            y='Tuition',
            x='UG',
            fit_reg=False)

plt.show()
plt.clf()

##
# Create a scatter plot and bin the data into 5 bins
sns.regplot(data=df,
            y='Tuition',
            x='UG',
            x_bins=5)

plt.show()
plt.clf()

##
# Create a regplot and bin the data into 8 bins
sns.regplot(data=df,
         y='Tuition',
         x='UG',
         x_bins=8)

plt.show()
plt.clf()
```

## Creating heatmaps
```python
# Create a crosstab table of the data
pd_crosstab = pd.crosstab(df["Group"], df["YEAR"])
print(pd_crosstab)

# Plot a heatmap of the table
sns.heatmap(pd_crosstab)

# Rotate tick marks for visibility
plt.yticks(rotation=0)
plt.xticks(rotation=90)

plt.show()
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

