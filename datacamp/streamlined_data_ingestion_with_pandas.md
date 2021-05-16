---
title: Streamlined Data Ingestion with pandas
tags: python,pandas
url: https://www.datacamp.com/courses/streamlined-data-ingestion-with-pandas
---

# 1. Importing Data from Flat Files
## Get data from CSVs
```python
# Import pandas as pd
import pandas as pd

# Read the CSV and assign it to the variable data
data = pd.read_csv("vt_tax_data_2016.csv")

# View the first few lines of data
print(data.head())
```

## Get data from other flat files
```python
# Import pandas with the alias pd
import pandas as pd

# Load TSV using the sep keyword argument to set delimiter
data = pd.read_csv("vt_tax_data_2016.tsv", sep="\t")

# Plot the total number of tax returns by income group
counts = data.groupby("agi_stub").N1.sum()
counts.plot.bar()
plt.show()
```

## Import a subset of columns
```python
# Create list of columns to use
cols = ["zipcode", "agi_stub", "mars1", "MARS2", "NUMDEP"]

# Create data frame from csv using only selected columns
data = pd.read_csv("vt_tax_data_2016.csv", usecols=cols)

# View counts of dependents and tax returns by income level
print(data.groupby("agi_stub").sum())
```

## Import a file in chunks
```python
# Create data frame of next 500 rows with labeled columns
vt_data_next500 = pd.read_csv("vt_tax_data_2016.csv", 
                       		  nrows=500,
                       		  skiprows=500,
                       		  header=None,
                       		  names=list(vt_data_first500))

# View the Vermont data frames to confirm they're different
print(vt_data_first500.head())
print(vt_data_next500.head())
```

## Specify data types
```python
# Create dict specifying data types for agi_stub and zipcode
data_types = {'agi_stub': 'category',
			  'zipcode': str}

# Load csv using dtype to set correct data types
data = pd.read_csv("vt_tax_data_2016.csv", dtype=data_types)

# Print data types of resulting frame
print(data.dtypes.head())
```

## Set custom NA values
```python

```

## Skip bad data
```python

```




# 2. Importing Data From Excel Files
## Introduction to spreadsheets
```python

```

## Get data from a spreadsheet
```python

```

## Load a portion of a spreadsheet
```python

```

## Getting data from multiple worksheets
```python

```

## Select a single sheet
```python

```

## Select multiple sheets
```python

```

## Work with multiple spreadsheets
```python

```

## Modifying imports: true/false data
```python

```

## Set Boolean columns
```python

```

## Set custom true/false values
```python

```

## Modifying imports: parsing dates
```python

```

## Parse simple dates
```python

```




# 3. Importing Data from Databases
## Introduction to databases
```python

```

## Connect to a database
```python

```

## Load entire tables
```python

```

## Refining imports with SQL queries
```python

```

## Selecting columns with SQL
```python

```

## Selecting rows
```python

```

## Filtering on multiple conditions
```python

```

## More complex SQL queries
```python

```

## Getting distinct values
```python

```

## Counting in groups
```python

```

## Working with aggregate functions
```python

```

## Loading multiple tables with joins
```python

```

## Joining tables
```python

```

## Joining and filtering
```python

```

## Joining, filtering, and aggregating
```python

```




# 4. Importing JSON Data and Working with APIs
## Introduction to JSON
```python

```

## Load JSON data
```python

```

## Work with JSON orientations
```python

```

## Introduction to APIs
```python

```

## Get data from an API
```python

```

## Set API parameters
```python

```

## Set request headers
```python

```

## Working with nested JSONs
```python

```

## Flatten nested JSONs
```python

```

## Handle deeply nested data
```python

```

## Combining multiple datasets
```python

```

## Append data frames
```python

```

## Merge data frames
```python

```

## Wrap-up
```python

```

