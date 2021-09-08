---
title: Introduction to Bash Scripting
tag: bash
url: https://campus.datacamp.com/courses/introduction-to-bash-scripting/
---

# 1. From Command-Line to Bash Script
## Extracting scores with shell
```bash
cat start_dir/second_dir/soccer_scores.csv | grep 1959
```

## Searching a book with shell
```bash
cat two_cities.txt | egrep 'Sydney Carton|Charles Darnay' | wc -l
```

## A simple Bash script
```bash
#!/bin/bash

# Concatenate the file
cat *.txt

# Now save and run!
```

## Shell pipelines to Bash scripts
```bash
#!/bin/bash

# Create a single-line pipe
cat soccer_scores.csv | cut -d "," -f 2 | tail -n +2 | sort | uniq -c

# Now save and run!
```

## Extract and edit using Bash scripts
```bash
#!/bin/bash

# Create a sed pipe to a new file
cat soccer_scores.csv | sed 's/Cherno/Cherno City/g' | sed 's/Arda/Arda United/g' > soccer_scores_edited.csv

# Now save and run!
```

## Using arguments in Bash scripts
```bash
# Echo the first and second ARGV arguments
echo $1
echo $2

# Echo out the entire ARGV array
echo $*

# Echo out the size of ARGV
echo $#
```

## Using arguments with HR data
```bash
# Echo the first ARGV argument
echo $1 

# Cat all the files
# Then pipe to grep using the first ARGV argument
# Then write out to a named csv using the first ARGV argument
cat hire_data/* | grep "$1" > "$1".csv
```




# 2. Variables in Bash Scripting
## Using variables in Bash
```bash
# Create the required variable
yourname="Sam"

# Print out the assigned name (Help fix this error!)
echo "Hi there ${yourname}, welcome to the website!"
```

## Shell within a shell
```bash
echo "Right now it is `date`"
```

## Converting Fahrenheit to Celsius
```bash
$ cat script.sh
# Get first ARGV into variable
temp_f=$1

# Subtract 32
temp_f2=$(echo "scale=2; $temp_f - 32" | bc)

# Multiply by 5/9 and print
temp_c=$(echo "scale=2; $temp_f2 * 5 / 9" | bc)

# Print the celsius temp
echo $temp_c

$ bash script.sh 108
```

## Extracting data from files
```bash
# Create three variables from the temp data files' contents
temp_a=$(cat temps/region_A)
temp_b=$(cat temps/region_B)
temp_c=$(cat temps/region_C)

# Print out the three variables
echo "The three temperatures were $temp_a, $temp_b, and $temp_c"
```

## Creating an array
```bash
##
# Create a normal array with the mentioned elements
capital_cities=("Sydney" "New York" "Paris")

##
# Create a normal array with the mentioned elements using the declare method
declare capital_cities=()

# Add (append) the elements
capital_cities+="Sydney"
capital_cities+="New York"
capital_cities+="Paris"

##
# The array has been created for you
capital_cities=("Sydney" "New York" "Paris")

# Print out the entire array
echo ${capital_cities[@]}

# Print out the array length
echo ${#capital_cities[@]}
```

## Creating associative arrays
```bash
##
# Create empty associative array
declare -A model_metrics

# Add the key-value pairs
model_metrics[model_accuracy]=98
model_metrics[model_name]="knn"
model_metrics[model_f1]=0.82

##
# Declare associative array with key-value pairs on one line
declare -A model_metrics=([model_accuracy]=98 [model_name]="knn" [model_f1]=0.82)

# Print out the entire array
echo ${model_metrics[@]}

##
# An associative array has been created for you
declare -A model_metrics=([model_accuracy]=98 [model_name]="knn" [model_f1]=0.82)

# Print out just the keys
echo ${!model_metrics[@]}
```

## Climate calculations in Bash
```bash
# Create variables from the temperature data files
temp_b="$(cat temps/region_B)"
temp_c="$(cat temps/region_C)"

# Create an array with these variables as elements
region_temps=($temp_b $temp_c)

# Call an external program to get average temperature
average_temp=$(echo "scale=2; (${region_temps[0]} + ${region_temps[1]}) / 2" | bc)

# Append average temp to the array
region_temps+=($average_temp)

# Print out the whole array
echo ${region_temps[@]}
```




# 3. Control Statements in Bash Scripting
## IF statements
```bash

```

## Sorting model results
```bash

```

## Moving relevant files
```bash

```

## FOR loops & WHILE statements
```bash

```

## A simple FOR loop
```bash

```

## Correcting a WHILE statement
```bash

```

## Cleaning up a directory
```bash

```

## CASE statements
```bash

```

## Days of the week with CASE
```bash

```

## Moving model results with CASE
```bash

```

## Finishing a CASE statement
```bash

```



# 4. Functions and Automation
## Basic functions in Bash
```bash

```

## Uploading model results to the cloud
```bash

```

## Get the current day
```bash

```

## Arguments, return values, and scope
```bash

```

## A percentage calculator
```bash

```

## Sports analytics function
```bash

```

## Summing an array
```bash

```

## Scheduling your scripts with Cron
```bash

```

## Analyzing a crontab schedule
```bash

```

## Creating cronjobs
```bash

```

## Scheduling cronjobs with crontab
```bash

```

## Thanks and wrap up
```bash

```
