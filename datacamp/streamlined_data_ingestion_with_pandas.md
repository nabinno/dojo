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
# Create dict specifying that 0s in zipcode are NA values
null_values = {"zipcode": 0}

# Load csv using na_values keyword argument
data = pd.read_csv("vt_tax_data_2016.csv", 
                   na_values=null_values)

# View rows with NA ZIP codes
print(data[data.zipcode.isna()])
```

## Skip bad data
```python
try:
  # Set warn_bad_lines to issue warnings about bad records
  data = pd.read_csv("vt_tax_data_2016_corrupt.csv", 
                     error_bad_lines=False, 
                     warn_bad_lines=True)
  
  # View first 5 records
  print(data.head())
  
except pd.io.common.CParserError:
    print("Your data contained rows that could not be parsed.")
```


# 2. Importing Data From Excel Files
## Get data from a spreadsheet
```python
# Load pandas as pd
import pandas as pd

# Read spreadsheet and assign it to survey_responses
survey_responses = pd.read_excel("fcc_survey.xlsx")

# View the head of the data frame
print(survey_responses.head())
```

## Load a portion of a spreadsheet
```python
# Create string of lettered columns to load
col_string = "AD,AW:BA"

# Load data with skiprows and usecols set
survey_responses = pd.read_excel("fcc_survey_headers.xlsx", 
                        skiprows=2,
                        usecols=col_string)

# View the names of the columns selected
print(survey_responses.columns)
```

## Select a single sheet
```python
##
# Create df from second worksheet by referencing its position
responses_2017 = pd.read_excel("fcc_survey.xlsx",
                               sheet_name=1)

# Graph where people would like to get a developer job
job_prefs = responses_2017.groupby("JobPref").JobPref.count()
job_prefs.plot.barh()
plt.show()

##
# Create df from second worksheet by referencing its name
responses_2017 = pd.read_excel("fcc_survey.xlsx",
                               sheet_name='2017')

# Graph where people would like to get a developer job
job_prefs = responses_2017.groupby("JobPref").JobPref.count()
job_prefs.plot.barh()
plt.show()
```

## Select multiple sheets
```python
##
# Load both the 2016 and 2017 sheets by name
all_survey_data = pd.read_excel("fcc_survey.xlsx",
                                sheet_name=['2016', '2017'])

# View the data type of all_survey_data
print(type(all_survey_data))

##
# Load all sheets in the Excel file
all_survey_data = pd.read_excel("fcc_survey.xlsx",
                                sheet_name=[0, '2017'])

# View the sheet names in all_survey_data
print(all_survey_data.keys())

##
# Load all sheets in the Excel file
all_survey_data = pd.read_excel("fcc_survey.xlsx",
                                sheet_name=None)

# View the sheet names in all_survey_data
print(all_survey_data.keys())
```

## Work with multiple spreadsheets
```python
# Create an empty data frame
all_responses = pd.DataFrame()

# Set up for loop to iterate through values in responses
for df in responses.values():
  # Print the number of rows being added
  print("Adding {} rows".format(df.shape[0]))
  # Append df to all_responses, assign result
  all_responses = all_responses.append(df)

# Graph employment statuses in sample
counts = all_responses.groupby("EmploymentStatus").EmploymentStatus.count()
counts.plot.barh()
plt.show()
```

## Set Boolean columns
```python
##
# Load the data
survey_data = pd.read_excel("fcc_survey_subset.xlsx")

# Count NA values in each column
print(survey_data.isna().sum())

##
# Set dtype to load appropriate column(s) as Boolean data
survey_data = pd.read_excel("fcc_survey_subset.xlsx",
                            dtype={"HasDebt": bool})

# View financial burdens by Boolean group
print(survey_data.groupby('HasDebt').sum())
```

## Set custom true/false values
```python
# Load file with Yes as a True value and No as a False value
survey_subset = pd.read_excel("fcc_survey_yn_data.xlsx",
                              dtype={"HasDebt": bool,
                              "AttendedBootCampYesNo": bool},
                              true_values=["Yes"],
                              false_values=["No"])

# View the data
print(survey_subset.head())
```

## Parse simple dates
```python
# Load file, with Part1StartTime parsed as datetime data
survey_data = pd.read_excel("fcc_survey.xlsx",
                            parse_dates=["Part1StartTime"])

# Print first few values of Part1StartTime
print(survey_data.Part1StartTime.head())
```

## Get datetimes from multiple columns
```python

```

## Parse non-standard date formats
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

