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

## Standard streams & arguments
```bash

```

## Using arguments in Bash scripts
```bash

```

## Using arguments with HR data
```bash

```




# 2. Variables in Bash Scripting
## Basic variables in Bash
```bash

```

## Using variables in Bash
```bash

```

## Shell within a shell
```bash

```

## Numeric variables in Bash
```bash

```

## Converting Fahrenheit to Celsius
```bash

```

## Extracting data from files
```bash

```

## Arrays in Bash
```bash

```

## Creating an array
```bash

```

## Creating associative arrays
```bash

```

## Climate calculations in Bash
```bash

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

