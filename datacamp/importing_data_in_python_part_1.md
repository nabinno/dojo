---
title: Importing Data in Python (Part 1)
tags: python,data-pre-processing
url: https://campus.datacamp.com/courses/importing-data-in-python-part-1/introduction-and-flat-files-1
---

# 1. Introduction and flat files
## Importing entire text files
```python
# Open a file: file
file = open("moby_dick.txt", "r")

# Print it
print(file.read())

# Check whether file is closed
print(file.closed)

# Close file
file.close()

# Check whether file is closed
print(file.closed)
```

## Importing text files line by line
```python
# Read & print the first 3 lines
with open('moby_dick.txt') as file:
    print(file.readline())
    print(file.readline())
    print(file.readline())
```

## Using NumPy to import flat files
```python
# Import package
import numpy as np

# Assign filename to variable: file
file = 'digits.csv'

# Load file as array: digits
digits = np.loadtxt(file, delimiter=",")

# Print datatype of digits
print(type(digits))

# Select and reshape a row
im = digits[21, 1:]
im_sq = np.reshape(im, (28, 28))

# Plot reshaped data (matplotlib.pyplot already loaded as plt)
plt.imshow(im_sq, cmap='Greys', interpolation='nearest')
plt.show()
```

## Customizing your NumPy import
```python
# Import numpy
import numpy as np

# Assign the filename: file
file = 'digits_header.txt'

# Load the data: data
data = np.loadtxt(file, delimiter="\t", skiprows=1, usecols=(0,2))

# Print data
print(data)
```

## Importing different datatypes
```python
# Assign filename: file
file = 'seaslug.txt'

# Import file: data
data = np.loadtxt(file, delimiter='\t', dtype=str)

# Print the first element of data
print(data[0])

# Import data as floats and skip the first row: data_float
data_float = np.loadtxt(file, delimiter="\t", dtype=float, skiprows=1)

# Print the 10th element of data_float
print(data_float[9])

# Plot a scatterplot of the data
plt.scatter(data_float[:, 0], data_float[:, 1])
plt.xlabel('time (min.)')
plt.ylabel('percentage of larvae')
plt.show()
```

## Working with mixed datatypes (2)
```python
# Assign the filename: file
file = 'titanic.csv'

# Import file using np.recfromcsv: d
d = np.recfromcsv(file)

# Print out first three entries of d
print(d[:3])
```

## Using pandas to import flat files as DataFrames (1)
```python
# Import pandas as pd
import pandas as pd

# Assign the filename: file
file = 'titanic.csv'

# Read the file into a DataFrame: df
df = pd.read_csv(file)

# View the head of the DataFrame
print(df.head())
```

## Using pandas to import flat files as DataFrame (2)
```python
# Assign the filename: file
file = 'digits.csv'

# Read the first 5 rows of the file into a DataFrame: data
data = pd.read_csv(file, nrows=5, header=None)

# Build a numpy array from the DataFrame: data_array
data_array = np.array(data)

# Print the datatype of data_array to the shell
print(type(data_array))
```

## Customizing your pandas import
```python
# Import matplotlib.pyplot as plt
import matplotlib.pyplot as plt

# Assign filename: file
file = 'titanic_corrupt.txt'

# Import file: data
data = pd.read_csv(file, sep="\t", comment="#", na_values="Nothing")

# Print the head of the DataFrame
print(data.head())

# Plot 'Age' variable in a histogram
pd.DataFrame.hist(data[['Age']])
plt.xlabel('Age (years)')
plt.ylabel('count')
plt.show()
```

# 2. Importing data from other file types
## Loading a pickled file
```python
# Import pickle package
import pickle

# Open pickle file and load data: d
with open('data.pkl', "rb") as file:
    d = pickle.load(file)

# Print d
print(d)

# Print datatype of d
print(type(d))
```

## Listing sheets in Excel files
```python
# Import pandas
import pandas as pd

# Assign spreadsheet filename: file
file = "battledeath.xlsx"

# Load spreadsheet: xls
xls = pd.ExcelFile(file)

# Print sheet names
print(xls.sheet_names)
```

## Importing sheets from Excel files
```python
# Load a sheet into a DataFrame by name: df1
df1 = xls.parse("2004")

# Print the head of the DataFrame df1
print(df1.head())

# Load a sheet into a DataFrame by index: df2
df2 = xls.parse(0)

# Print the head of the DataFrame df2
print(df2.head())
```

## Customizing your spreadsheet import
```python
# Parse the first sheet and rename the columns: df1
df1 = xls.parse(0, skiprows=[0], names=["Country", "AAM due to War (2002)"])

# Print the head of the DataFrame df1
print(df1.head())

# Parse the first column of the second sheet and rename the column: df2
df2 = xls.parse(1, usecols=0, skiprows=[0], names=["Country"])

# Print the head of the DataFrame df2
print(df2.head())
```

## Importing SAS files
```python

# Import sas7bdat package
from sas7bdat import SAS7BDAT

# Save file to a DataFrame: df_sas
with SAS7BDAT('sales.sas7bdat') as file:
    df_sas = file.to_data_frame()

# Print head of DataFrame
print(df_sas.head())

# Plot histogram of DataFrame features (pandas and pyplot already imported)
pd.DataFrame.hist(df_sas[['P']])
plt.ylabel('count')
plt.show()
```

## Importing Stata files
```python
# Import pandas
import pandas as pd

# Load Stata file into a pandas DataFrame: df
df = pd.read_stata("disarea.dta")

# Print the head of the DataFrame df
print(df.head())

# Plot histogram of one column of the DataFrame
pd.DataFrame.hist(df[['disa10']])
plt.xlabel('Extent of disease')
plt.ylabel('Number of countries')
plt.show()
```

## Using h5py to import HDF5 files
```python
# Import packages
import numpy as np
import h5py

# Assign filename: file
file = "LIGO_data.hdf5"

# Load file: data
data = h5py.File(file, "r")

# Print the datatype of the loaded file
print(type(data))

# Print the keys of the file
for key in data.keys():
    print(key)
```

## Extracting data from your HDF5 file
```python
# Get the HDF5 group: group
group = data["strain"]

# Check out keys of group
for key in group.keys():
    print(key)

# Set variable equal to time series data: strain
strain = data["strain"]["Strain"].value

# Set number of time points to sample: num_samples
num_samples = 10000

# Set time vector
time = np.arange(0, 1, 1/num_samples)

# Plot data
plt.plot(time, strain[:num_samples])
plt.xlabel('GPS Time (s)')
plt.ylabel('strain')
plt.show()
```

## Loading .mat files
```python
# Import package
import scipy.io

# Load MATLAB file: mat
mat = scipy.io.loadmat("albeck_gene_expression.mat")

# Print the datatype type of mat
print(type(mat))
```

## The structure of .mat in Python
```python
# Print the keys of the MATLAB dictionary
print(mat.keys())

# Print the type of the value corresponding to the key 'CYratioCyt'
print(type(mat["CYratioCyt"]))

# Print the shape of the value corresponding to the key 'CYratioCyt'
print(np.shape(mat["CYratioCyt"]))

# Subset the array and plot it
data = mat['CYratioCyt'][25, 5:]
fig = plt.figure()
plt.plot(data)
plt.xlabel('time (min.)')
plt.ylabel('normalized fluorescence (measure of expression)')
plt.show()
```

# 3. Working with relational databases in Python
##
```python

```
