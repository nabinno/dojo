---
tags: seaborn, python, data-visualization
title: Introduction to Data Visualization with Seaborn
url: https://campus.datacamp.com/courses/introduction-to-data-visualization-with-seaborn/
---

# 1. Introduction to Seaborn
## Making a scatter plot with lists
```python
# Import Matplotlib and Seaborn
import matplotlib.pyplot as plt
import seaborn as sns

# Change this scatter plot to have percent literate on the y-axis
sns.scatterplot(x=gdp, y=percent_literate)

# Show plot
plt.show()
```

## Making a count plot with a list
```python
# Import Matplotlib and Seaborn
import matplotlib.pyplot as plt
import seaborn as sns

# Create count plot with region on the y-axis
sns.countplot(y=region)

# Show plot
plt.show()
```

## "Tidy" vs. "untidy" data
```python
# Import pandas
import pandas as pd

# Create a DataFrame from csv file
df = pd.read_csv(csv_filepath)

# Print the head of df
print(df.head())
```

## Making a count plot with a DataFrame
```python
# Import Matplotlib, pandas, and Seaborn
import matplotlib.pyplot as plt
import pandas as pd
import seaborn as sns

# Create a DataFrame from csv file
df = pd.read_csv(csv_filepath)

# Create a count plot with "Spiders" on the x-axis
sns.countplot(x="Spiders", data=df)

# Display the plot
plt.show()
```

## Adding a third variable with hue
```python
# Import Matplotlib and Seaborn
import matplotlib.pyplot as plt
import seaborn as sns

# Change the legend order in the scatter plot
sns.scatterplot(x="absences", y="G3", 
                data=student_data, 
                hue="location", hue_order=["Rural", "Urban"])

# Show plot
plt.show()
```

## Hue and count plots
```python
# Import Matplotlib and Seaborn
import matplotlib.pyplot as plt
import seaborn as sns

# Create a dictionary mapping subgroup values to colors
palette_colors = {"Rural": "green", "Urban": "blue"}

# Create a count plot of school with location subgroups
sns.countplot(x="school", data=student_data, hue="location", palette=palette_colors)

# Display plot
plt.show()
```




# 2. Visualizing Two Quantitative Variables
## Creating subplots with col and row
```python
##
# Change this scatter plot to arrange the plots in rows instead of columns
sns.relplot(x="absences", y="G3", 
            data=student_data,
            kind="scatter", 
            col="study_time")

# Show plot
plt.show()

##
# Change this scatter plot to arrange the plots in rows instead of columns
sns.relplot(x="absences", y="G3", 
            data=student_data,
            kind="scatter", 
            row="study_time")

# Show plot
plt.show()
```

## Creating two-factor subplots
```python
# Adjust further to add subplots based on family support
sns.relplot(x="G1", y="G3", 
            data=student_data,
            kind="scatter", 
            col="schoolsup",
            col_order=["yes", "no"],
            row="famsup", row_order=["yes", "no"])

# Show plot
plt.show()
```

## Changing the size of scatter plot points
```python
# Import Matplotlib and Seaborn
import matplotlib.pyplot as plt
import seaborn as sns

# Create scatter plot of horsepower vs. mpg
sns.relplot(x="horsepower", y="mpg", 
            data=mpg, kind="scatter", 
            size="cylinders", hue="cylinders")

# Show plot
plt.show()
```

## Changing the style of scatter plot points
```python
# Import Matplotlib and Seaborn
import matplotlib.pyplot as plt
import seaborn as sns

# Create a scatter plot of acceleration vs. mpg
sns.relplot(x="acceleration", y="mpg", data=mpg, kind="scatter", style="origin", hue="origin")

# Show plot
plt.show()
```

## Interpreting line plots
```python
##
# Import Matplotlib and Seaborn
import matplotlib.pyplot as plt
import seaborn as sns

# Create line plot
sns.relplot(x="model_year", y="mpg", data=mpg, kind="line")

# Show plot
plt.show()
```

## Visualizing standard deviation with line plots
```python
# Make the shaded area show the standard deviation
sns.relplot(x="model_year", y="mpg",
            data=mpg, kind="line", ci="sd")

# Show plot
plt.show()
```

## Plotting subgroups in line plots
```python
# Import Matplotlib and Seaborn
import matplotlib.pyplot as plt
import seaborn as sns

# Add markers and make each line have the same style
sns.relplot(x="model_year", y="horsepower", 
            data=mpg, kind="line", 
            ci=None, style="origin", 
            hue="origin", dashes=False, markers=True)

# Show plot
plt.show()
```




# 3. Visualizing a Categorical and a Quantitative Variable
## Count plots and bar plots
```python

```

## Count plots
```python

```

## Bar plots with percentages
```python

```

## Customizing bar plots
```python

```

## Box plots
```python

```

## Create and interpret a box plot
```python

```

## Omitting outliers
```python

```

## Adjusting the whiskers
```python

```

## Point plots
```python

```

## Customizing point plots
```python

```

## Point plots with subgroups
```python

```




# 4. Customizing Seaborn Plots
## Changing plot style and color
```python

```

## Changing style and palette
```python

```

## Changing the scale
```python

```

## Using a custom palette
```python

```

## Adding titles and labels: Part 1
```python

```

## FacetGrids vs. AxesSubplots
```python

```

## Adding a title to a FacetGrid object
```python

```

## Adding titles and labels: Part 2
```python

```

## Adding a title and axis labels
```python

```

## Rotating x-tick labels
```python

```

## Putting it all together
```python

```

## Box plot with subgroups
```python

```

## Bar plot with subgroups and subplots
```python

```

## Well done! What's next?
```python

```

