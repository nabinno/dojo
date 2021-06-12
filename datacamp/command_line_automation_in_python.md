---
title: Command Line Automation in Python
tags: python
url: https://www.datacamp.com/courses/command-line-automation-in-python
---

# 1. IPython shell commands
## Learn the Python interpreter
```sh
$ cat script.ipy
!python3 -c "from random import choices;days = ['Mo', 'Tu', 'We', 'Th', 'Fr'];print(choices(days))"

$ ipython script.ipy
```

## Execute Python commands
```ipython
var = !ls -h test_dir/*.csv
print(len(var))
```

## Capture IPython Shell output
```ipython
%%bash
echo "Running Directory Audit Script"
CSV=`ls -l test_dir/*.csv | wc -l`
TXT=`ls -l test_dir/*.csv | wc -l`
echo 'The directory contains this a total # *.csv files: ' $CSV
echo 'The directory contains this a total # *.txt files: ' $TXT
```

## Using Bash Magic command
```ipython
!ls -l | awk '{SUM+=$5} END {print SUM}'
```

## Use SList fields to parse shell output
```ipython
disk_space = !df -h
print(disk_space.fields(0))
```

## Find Python files using SLIST grep
```ipython
res = !ls src
print(res.grep('.py'))
```

## Using SList to grep
```python
import os
from slist_out import slist_out

# Save the name of the root directory
root = "test_dir"

# Find the backups with "_2" in slist_out
result = slist_out.grep('_2')

# Extract the filenames
for res in result:
	filename = res.split()[-1]
    
	# Create the full path
	fullpath = os.path.join(root, filename)
	print(f"fullpath of backup file: {fullpath}")
```



# 2. Shell commands with subprocess
## Execute shell commands in subprocess
```python

```

## Permission check
```python

```

## Reading a creepy AI poem
```python

```

## Running processes script
```python

```

## Capture output of shell commands
```python

```

## Using subprocess Popen
```python

```

## Waiting for processes
```python

```

## Detecting duplicate files with subprocess
```python

```

## Sending input to processes
```python

```

## Counting files in a directory tree
```python

```

## Running a health check
```python

```

## Passing arguments safely to shell commands
```python

```

## Safely find directories
```python

```

## Directory summarizer
```python

```






# 3. Walking the file system

## Dealing with file systems
```python

```

## Double trouble
```python

```

## Y'all got some renaming to do
```python

```

## Sweet pickle
```python

```

## Find files matching a pattern
```python

```

## Rogue founder code
```python

```

## Is this pattern True?
```python

```

## High-level file and directory operations
```python

```

## Goons over my shammy
```python

```

## Archive users
```python

```

## Using pathlib
```python

```

## Does it even exist?
```python

```

## File writing one-liner
```python

```






# 4. Command line functions

## Using functions for automation
```python

```

## Funky clusters
```python

```

## Hello decorator
```python

```

## Debugging decorator
```python

```

## Understand script input
```python

```

## Using python command-line tools
```python

```

## Backwards day
```python

```

## Introduction to Click
```python

```

## Simple yet true
```python

```

## Running a click application from subprocess
```python

```

## Using click to write command line tools
```python

```

## Got a ticket to write
```python

```

## Invoking command line tests
```python

```

## Course Summary
```python

```

