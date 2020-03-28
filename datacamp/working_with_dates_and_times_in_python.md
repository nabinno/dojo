---
title: Working with Dates and Times in Python
tags: python,database
url: https://www.datacamp.com/courses/working-with-dates-and-times-in-python
---

# 1. Dates and Calendars
## Which day of the week?
```python
# Import date from datetime
from datetime import date

# Create a date object
hurricane_andrew = date(1992, 8, 24)

# Which day of the week is the date?
print(hurricane_andrew.weekday())
```

## How many hurricanes come early?
```python
# Counter for how many before June 1
early_hurricanes = 0

# We loop over the dates
for hurricane in florida_hurricane_dates:
  # Check if the month is before June (month number 6)
  if hurricane.month < 6:
    early_hurricanes = early_hurricanes + 1
    
print(early_hurricanes)
```

## Subtracting dates
```python
# Import date
from datetime import date

# Create a date object for May 9th, 2007
start = date(2007, 5, 9)

# Create a date object for December 13th, 2007
end = date(2007, 12, 13)

# Subtract the two dates and print the number of days
print((end - start).days)
```

## Counting events per calendar month
```python
# A dictionary to count hurricanes per calendar month
hurricanes_each_month = {1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6:0,
		  				 7: 0, 8:0, 9:0, 10:0, 11:0, 12:0}

# Loop over all hurricanes
for hurricane in florida_hurricane_dates:
  # Pull out the month
  month = hurricane.month
  # Increment the count in your dictionary by one
  hurricanes_each_month[month] += 1
  
print(hurricanes_each_month)
```

## Putting a list of dates in order
```python
# Print the first and last scrambled dates
print(dates_scrambled[0])
print(dates_scrambled[-1])

# Put the dates in order
dates_ordered = sorted(dates_scrambled)

# Print the first and last ordered dates
print(dates_ordered[0])
print(dates_ordered[-1])
```

## Printing dates in a friendly format
```python
# Assign the earliest date to first_date
first_date = min(florida_hurricane_dates)

# Convert to ISO and US formats
iso = "Our earliest hurricane date: " + first_date.isoformat()
us = "Our earliest hurricane date: " + first_date.strftime("%m/%d/%Y")

print("ISO: " + iso)
print("US: " + us)
```

## Representing dates in different ways
```python
##
# Import date
from datetime import date

# Create a date object
andrew = date(1992, 8, 26)

# Print the date in the format 'YYYY-MM'
print(andrew.strftime('%Y-%m'))

##
# Import date
from datetime import date

# Create a date object
andrew = date(1992, 8, 26)

# Print the date in the format 'MONTH (YYYY)'
print(andrew.strftime('%B (%Y)'))

##
# Import date
from datetime import date

# Create a date object
andrew = date(1992, 8, 26)

# Print the date in the format 'YYYY-DDD'
print(andrew.strftime('%Y-%j'))
```


# 2. Combining Dates and Times
## Creating datetimes by hand
```python
##
# Import datetime
from datetime import datetime

# Create a datetime object
dt = datetime(2017, 10, 1, 15, 26, 26)

# Print the results in ISO 8601 format
print(dt.isoformat())

##
# Import datetime
from datetime import datetime

# Create a datetime object
dt = datetime(2017, 12, 31, 15, 19, 13)

# Print the results in ISO 8601 format
print(dt.isoformat())

##
# Import datetime
from datetime import datetime

# Create a datetime object
dt = datetime(2017, 12, 31, 15, 19, 13)

# Replace the year with 1917
dt_old = dt.replace(year=1917)

# Print the results in ISO 8601 format
print(dt_old)
```

## Counting events before and after noon
```python

```

## Printing and parsing datetimes
```python

```

## Turning strings into datetimes
```python

```

## Parsing pairs of strings as datetimes
```python

```

## Recreating ISO format with strftime()
```python

```

## Unix timestamps
```python

```

## Working with durations
```python

```

## Turning pairs of datetimes into durations
```python

```

## Average trip time
```python

```

## The long and the short of why time is hard
```python

```


# 3. Time Zones and Daylight Saving
## UTC offsets
```python

```

## Creating timezone aware datetimes
```python

```

## Setting timezones
```python

```

## What time did the bike leave in UTC?
```python

```

## Time zone database
```python

```

## Putting the bike trips into the right time zone
```python

```

## What time did the bike leave? (Global edition)
```python

```

## Starting daylight saving time
```python

```

## How many hours elapsed around daylight saving?
```python

```

## March 29, throughout a decade
```python

```

## Ending daylight saving time
```python

```

## Finding ambiguous datetimes
```python

```

## Cleaning daylight saving data with fold
```python

```


# 4. Easy and Powerful: Dates and Times in Pandas
## Reading date and time data in Pandas
```python

```

## Loading a csv file in Pandas
```python

```

## Making timedelta columns
```python

```

## Summarizing datetime data in Pandas
```python

```

## How many joyrides?
```python

```

## It's getting cold outside, W20529
```python

```

## Members vs casual riders over time
```python

```

## Combining groupby() and resample()
```python

```

## Additional datetime methods in Pandas
```python

```

## Timezones in Pandas
```python

```

## How long per weekday?
```python

```

## How long between rides?
```python

```

## Wrap-up
```python

```

