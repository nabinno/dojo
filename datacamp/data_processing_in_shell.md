---
title: Data Processing in Shell
tags: shell,data-engineering
url: https://www.datacamp.com/courses/data-processing-in-shell
---

# 1. Downlonading Data on the Command Line
## Using curl documentation
```sh
# Download and rename the file in the same step
curl -o Spotify201812.zip -L https://assets.datacamp.com/production/repositories/4180/datasets/eb1d6a36fa3039e4e00064797e1a1600d267b135/201812SpotifyData.zip
```

## Downloading multiple files using curl
```sh
# Download all 100 data files
curl -O https://s3.amazonaws.com/assets.datacamp.com/production/repositories/4180/datasets/files/datafile[001-100].txt

# Print all downloaded files to directory
ls datafile*.txt
```

## Downloading single file using wget
```sh
# Fill in the two option flags 
wget -c -b https://assets.datacamp.com/production/repositories/4180/datasets/eb1d6a36fa3039e4e00064797e1a1600d267b135/201812SpotifyData.zip

# Verify that the Spotify file has been downloaded
ls 

# Preview the log file 
cat wget-log
```

## Creating wait time using Wget
```sh
# View url_list.txt to verify content
cat url_list.txt

# Create a mandatory 1 second pause between downloading all files in url_list.txt
wget --wait=1 -i url_list.txt

# Take a look at all files downloaded
ls
```

## Data downloading with Wget and curl
```sh
# Use curl, download and rename a single file from URL
curl -o Spotify201812.zip -L https://assets.datacamp.com/production/repositories/4180/datasets/eb1d6a36fa3039e4e00064797e1a1600d267b135/201812SpotifyData.zip

# Unzip, delete, then re-name to Spotify201812.csv
unzip Spotify201812.zip && rm Spotify201812.zip
mv 201812SpotifyData.csv Spotify201812.csv

# View url_list.txt to verify content
cat url_list.txt

# Use Wget, limit the download rate to 2500 KB/s, download all files in url_list.txt
wget --limit=2500k -i url_list.txt

# Take a look at all files downloaded
ls
```



# 2. Data Cleaning and Munging on the Command Line
## Getting started with csvkit
```sh

```

## Installation and documentation for csvkit
```sh

```

## Converting and previewing data with csvkit
```sh

```

## File conversion and summary statistics with csvkit
```sh

```

## Filtering data using csvkit
```sh

```

## Printing column headers with csvkit
```sh

```

## Filtering data by column with csvkit
```sh

```

## Filtering data by row with csvkit
```sh

```

## Stacking data and chaining commands with csvkit
```sh

```

## Stacking files with csvkit
```sh

```

## Chaining commands using operators
```sh

```

## Data processing with csvkit
```sh

```




# 3. Database Operations on the Command Line

## Pulling data from database
```sh

```

## Using sql2csv documentation
```sh

```

## Understand sql2csv connectors
```sh

```

## Practice pulling data from database
```sh

```

## Manipulating data using SQL syntax
```sh

```

## Applying SQL to a local CSV file
```sh

```

## Cleaner scripting via shell variables
```sh

```

## Joining local CSV files using SQL
```sh

```

## Pushing data back to database
```sh

```

## Practice pushing data back to database
```sh

```

## Database and SQL with csvkit
```sh

```





# 4. Data Pipeline one the Command Line

## Sh on the command line
```sh

```

## Finding Sh version on the command line
```sh

```

## Executing Sh script on the command line
```sh

```

## Sh package installation with pip
```sh

```

## Understanding pip's capabilities
```sh

```

## Installing Sh dependencies
```sh

```

## Running a Sh model
```sh

```

## Data job automation with cron
```sh

```

## Understanding cron scheduling syntax
```sh

```

## Scheduling a job with crontab
```sh

```

## Model production on the command line
```sh

```

## Course recap
```sh

```


