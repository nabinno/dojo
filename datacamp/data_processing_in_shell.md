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
## Installation and documentation for csvkit
```sh
# Upgrade csvkit using pip  
pip install --upgrade csvkit

# Print manual for in2csv
in2csv -h

# Print manual for csvlook
csvlook -h
```

## Converting and previewing data with csvkit
```sh
# Use ls to find the name of the zipped file
ls

# Use Linux's built in unzip tool to unpack the zipped file 
unzip SpotifyData.zip

# Check to confirm name and location of unzipped file
ls

# Convert SpotifyData.xlsx to csv
in2csv SpotifyData.xlsx > SpotifyData.csv

# Print a preview in console using a csvkit suite command 
csvlook SpotifyData.csv
```

## File conversion and summary statistics with csvkit
```sh
##
# Check to confirm name and location of the Excel data file
ls

# Convert sheet "Worksheet1_Popularity" to CSV
in2csv SpotifyData.xlsx --sheet "Worksheet1_Popularity" > Spotify_Popularity.csv

# Check to confirm name and location of the new CSV file
ls

# Print high level summary statistics for each column
csvstat Spotify_Popularity.csv 

##
# Check to confirm name and location of the Excel data file
ls

# Convert sheet "Worksheet2_MusicAttributes" to CSV
in2csv SpotifyData.xlsx --sheet "Worksheet2_MusicAttributes" > Spotify_MusicAttributes.csv

# Check to confirm name and location of the new CSV file
ls

# Print preview of Spotify_MusicAttributes
csvlook Spotify_MusicAttributes.csv
```

## Printing column headers with csvkit
```sh
# Check to confirm name and location of data file
ls

# Print a list of column headers in data file 
csvcut -n Spotify_MusicAttributes.csv
```

## Filtering data by column with csvkit
```sh
##
# Print a list of column headers in the data 
csvcut -n Spotify_MusicAttributes.csv

# Print the first column, by position
csvcut -c 1 Spotify_MusicAttributes.csv

##
# Print a list of column headers in the data 
csvcut -n Spotify_MusicAttributes.csv

# Print the first, third, and fifth column, by position
csvcut -c 1,3,5 Spotify_MusicAttributes.csv

##
# Print a list of column headers in the data 
csvcut -n Spotify_MusicAttributes.csv

# Print the first column, by name
csvcut -c "track_id" Spotify_MusicAttributes.csv

##
# Print a list of column headers in the data 
csvcut -n Spotify_MusicAttributes.csv

# Print the track id, song duration, and loudness, by name 
csvcut -c "track_id","duration_ms","loudness" Spotify_MusicAttributes.csv
```

## Filtering data by row with csvkit
```sh
##
# Print a list of column headers in the data 
csvcut -n Spotify_MusicAttributes.csv

# Filter for row(s) where track_id = 118GQ70Sp6pMqn6w1oKuki
csvgrep -c "track_id" -m 118GQ70Sp6pMqn6w1oKuki Spotify_MusicAttributes.csv

##
# Print a list of column headers in the data 
csvcut -n Spotify_MusicAttributes.csv

# Filter for row(s) where danceability = 0.812
csvgrep -c "danceability" -m 0.812 Spotify_MusicAttributes.csv
```

## Stacking files with csvkit
```sh
# Stack the two files and save results as a new file
csvstack SpotifyData_PopularityRank6.csv SpotifyData_PopularityRank7.csv > SpotifyPopularity.csv

# Preview the newly created file 
csvlook SpotifyPopularity.csv
```

## Chaining commands using operators
```sh
##
# If csvlook succeeds, then run csvstat 
csvlook Spotify_Popularity.csv && csvstat Spotify_Popularity.csv

##
# Use the output of csvsort as input to csvlook
csvsort -c 2 Spotify_Popularity.csv | csvlook

##
# Take top 15 rows from sorted output and save to new file
csvsort -c 2 Spotify_Popularity.csv | head -n 15 > Spotify_Popularity_Top15.csv

# Preview the new file 
csvlook Spotify_Popularity_Top15.csv
```

## Data processing with csvkit
```sh
# Convert the Spotify201809 tab into its own csv file 
in2csv Spotify_201809_201810.xlsx --sheet "Spotify201809" > Spotify201809.csv

# Check to confirm name and location of data file
ls

# Preview file preview using a csvkit function
csvlook Spotify201809.csv

# Create a new csv with 2 columns: track_id and popularity
csvcut -c "track_id","popularity" Spotify201809.csv > Spotify201809_subset.csv

# While stacking the 2 files, create a data source column
csvstack -g "Sep2018","Oct2018" Spotify201809_subset.csv Spotify201810_subset.csv > Spotify_all_rankings.csv
```


# 3. Database Operations on the Command Line
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


