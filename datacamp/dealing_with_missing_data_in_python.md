---
title: Dealing with Missing Data in Python
tags: python
url: https://www.datacamp.com/courses/dealing-with-missing-data-in-python
---

# 1. The Problem With Missing Data
## Null value operations
```python
##
try:
  # Print the sum of two None's
  print("Add operation output of 'None': ", None + None)

except TypeError:
  # Print if error
  print("'None' does not support Arithmetic Operations!!")

##
try:
  # Print the sum of two np.nan's
  print("Add operation output of 'np.nan': ", np.nan + np.nan)

except TypeError:
  # Print if error
  print("'np.nan' does not support Arithmetic Operations!!")

##
try:
  # Print the output of logical OR of two None's
  print("OR operation output of 'None': ", None or None)

except TypeError:
  # Print if error
  print("'None' does not support Logical Operations!!")

##
try:
  # Print the output of logical OR of two np.nan's
  print("OR operation output of 'np.nan': ", np.nan or np.nan)

except TypeError:
  # Print if error
  print("'np.nan' does not support Logical Operations!!")
```

## Finding Null values
```python
##
try:
  # Print the comparison of two 'None's
  print("'None' comparison output: ", None == None)

except TypeError:
  # Print if error
  print("'None' does not support this operation!!")

##
try:
  # Print the comparison of two 'np.nan's
  print("'np.nan' comparison output: ", np.nan == np.nan)

except TypeError:
  # Print if error  
  print("'np.nan' does not support this operation!!")

##
try:
  # Check if 'None' is 'NaN'
  print("Is 'None' same as nan? ", np.isnan(None))

except TypeError:
  # Print if error
  print("Function 'np.isnan()' does not support this Type!!")

##
try:
  # Check if 'np.nan' is 'NaN'
  print("Is 'np.nan' same as nan? ", np.isnan(np.nan))

except TypeError:
  # Print if error
  print("Function 'np.isnan()' does not support this Type!!")
```

## Detecting missing values
```python
# Read the dataset 'college.csv'
college = pd.read_csv('college.csv')
print(college.head())

# Print the info of college
print(college.info())

# Store unique values of 'csat' column to 'csat_unique'
csat_unique = college.csat.unique()

# Print the sorted values of csat_unique
print(np.sort(csat_unique))
```

## Replacing missing values
```python
# Read the dataset 'college.csv' with na_values set to '.'
college = pd.read_csv('college.csv', na_values='.')
print(college.head())

# Print the info of college
print(college.info())
```

## Replacing hidden missing values
```python
# Print the description of the data
print(diabetes.describe())

# Store all rows of column 'BMI' which are equal to 0 
zero_bmi = diabetes.BMI[diabetes.BMI == 0]
print(zero_bmi)

# Set the 0 values of column 'BMI' to np.nan
diabetes.BMI[diabetes.BMI == 0] = np.nan

# Print the 'NaN' values in the column BMI
print(diabetes.BMI[np.isnan(diabetes.BMI)])
```

## Analyzing missingness percentage
```python
# Load the airquality dataset
airquality = pd.read_csv('air-quality.csv', parse_dates=['Date'], index_col='Date')

# Create a nullity DataFrame airquality_nullity
airquality_nullity = airquality.isnull()
print(airquality_nullity.head())

# Calculate total of missing values
missing_values_sum = airquality_nullity.sum()
print('Total Missing Values:\n', missing_values_sum)

# Calculate percentage of missing values
missing_values_percent = airquality_nullity.mean() * 100
print('Percentage of Missing Values:\n', missing_values_percent)
```

## Visualize missingness
```python
##
# Import missingno as msno
import missingno as msno

# Plot amount of missingness
msno.bar(airquality)

# Display bar chart of missing values
display("/usr/local/share/datasets/bar_chart.png")

##
# Import missingno as msno
import missingno as msno

# Plot nullity matrix of airquality
msno.matrix(airquality)

# Display nullity matrix
display("/usr/local/share/datasets/matrix.png")

##
# Import missingno as msno
import missingno as msno

# Plot nullity matrix of airquality with frequency 'M'
msno.matrix(airquality, freq='M')

# Display nullity matrix
display("/usr/local/share/datasets/matrix_frequency.png")

##
# Import missingno as msno
import missingno as msno

# Plot the sliced nullity matrix of airquality with frequency 'M'
msno.matrix(airquality.loc['May-1976':'Jul-1976'], freq='M')

# Display nullity matrix
display("/usr/local/share/datasets/matrix_sliced.png")
```

# 2. Does Missingness Have A Pattern?
## Is the data missing at random?
```python

```

## Guess the missingness type
```python

```

## Deduce MNAR
```python

```

## Finding patterns in missing data
```python

```

## Finding correlations in your data
```python

```

## Identify the missingness type
```python

```

## Visualizing missingness across a variable
```python

```

## Fill dummy values
```python

```

## Generate scatter plot with missingness
```python

```

## When and how to delete missing data
```python

```

## Delete MCAR
```python

```

## Will you delete?
```python

```


# 3. Imputation Techniques
## Mean, median & mode imputations
```python

```

## Mean & median imputation
```python

```

## Mode and constant imputation
```python

```

## Visualize imputations
```python

```

## Imputing time-series data
```python

```

## Filling missing time-series data
```python

```

## Impute with interpolate method
```python

```

## Visualizing time-series imputations
```python

```

## Visualize forward fill imputation
```python

```

## Visualize backward fill imputation
```python

```

## Plot interpolations
```python

```


# 4. Advanced Imputation Techniques
## Imputing using fancyimpute
```python

```

## KNN imputation
```python

```

## MICE imputation
```python

```

## Imputing categorical values
```python

```

## Ordinal encoding of a categorical column
```python

```

## Ordinal encoding of a DataFrame
```python

```

## KNN imputation of categorical values
```python

```

## Evaluation of different imputation techniques
```python

```

## Analyze the summary of linear model
```python

```

## Comparing R-squared and coefficients
```python

```

## Comparing density plots
```python

```

## Conclusion
```python

```

