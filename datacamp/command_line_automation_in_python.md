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
## Permission check
```python
# Import packages
import subprocess
import os

# Setup
file_location = "/tmp/file.txt"
expected_uid = 1000

# Create a file
proc = subprocess.Popen(["touch", file_location])

# Check user permissions
stat = os.stat(file_location)
if stat.st_uid == expected_uid:
    print(f"File System exported properly: {expected_uid} == {stat.st_uid}")
else:
    print(f"File System NOT exported properly: {expected_uid} != {stat.st_uid}")
```

## Reading a creepy AI poem
```python
import subprocess

# Execute Unix command `head` safely as items in a list
with subprocess.Popen(["head", "/home/repl/workspace/poem.txt"], stdout=subprocess.PIPE) as head:
  
    # Print each line of list returned by `stdout.readlines()`
    for line in head.stdout.readlines():
        print(line)
    
# Execute Unix command `wc -c` safely as items in a list
with subprocess.Popen(["wc", "-c", "/home/repl/workspace/poem.txt"], stdout=subprocess.PIPE) as word_count:
  
    # Print the string output of standard out of `wc -c`
    print(word_count.stdout.read())
```

## Running processes script
```python
import subprocess

# Use subprocess to run the `ps aux` command that lists running processes
with subprocess.Popen(["ps", "aux"], stdout=subprocess.PIPE) as proc:
    process_output = proc.stdout.readlines()
    
# Look through each line in the output and skip it if it contains "python"
for line in process_output:
    if b"python" in line:
        continue
    print(line)
```

## Using subprocess Popen
```python
from subprocess import Popen, PIPE
import json
import pprint

# Use the with context manager to run subprocess.Popen()
with Popen(["pip","list","--format=json"], stdout=PIPE) as proc:
  # Pipe the output of subprocess.Popen() to stdout
  result = proc.stdout.readlines()
  
# Convert the JSON payload to a Python dictionary
# JSON is a datastructure similar to a Python dictionary
converted_result = json.loads(result[0])

# Display the result in the IPython terminal
pprint.pprint(converted_result)
```

## Waiting for processes
```python
from subprocess import (Popen, PIPE, TimeoutExpired)

# Start a long running process using subprocess.Popen()
proc = Popen(["sleep", "6"], stdout=PIPE, stderr=PIPE)

# Use subprocess.communicate() to create a timeout 
try:
    output, error = proc.communicate(timeout=5)
    
except TimeoutExpired:

	# Cleanup the process if it takes longer than the timeout
    proc.kill()
    
    # Read standard out and standard error streams and print
    output, error = proc.communicate()
    print(f"Process timed out with output: {output}, error: {error}")
```

## Detecting duplicate files with subprocess
```python
from subprocess import (Popen, PIPE)
import os

# Don't change these variables
checksums = {}
duplicates = []
root = "/home/repl/workspace/test_dir"
files = [os.path.join(root, f) for f in os.listdir(root)]

# Iterate over the list of files filenames
for filename in files:
    # Use Popen to call the md5sum utility
    with Popen(["md5sum", filename], stdout=PIPE) as proc:
        checksum, _ = proc.stdout.read().split()
        
        # Append duplicate to a list if the checksum is found
        if checksum in checksums:
            duplicates.append(filename)
        checksums[checksum] = filename

print(f"Found Duplicates: {duplicates}")
```

## Counting files in a directory tree
```python
import subprocess

# Runs find command to search for files
find = subprocess.Popen(
    ["find", "/home/repl/workspace/test_dir", "-type", "f", "-print"], stdout=subprocess.PIPE)

# Runs wc and counts the number of lines
word_count = subprocess.Popen(
    ["wc", "-l"], stdin=find.stdout, stdout=subprocess.PIPE)

# Print the decoded and formatted output
output = word_count.stdout.read()
print(output.decode("utf-8").strip())
```

## Running a health check
```python
import subprocess

# equivalent to 'echo "python3"'
echo = subprocess.Popen(
    ["echo", "python3"], stdout=subprocess.PIPE)

# equivalent to: echo "python3" | ./healthcheck.sh
path = subprocess.Popen(
    ["/home/repl/workspace/healthcheck.sh"], stdin=echo.stdout, stdout=subprocess.PIPE)

full_path = path.stdout.read().decode("utf-8")
print(f"...Health Check Output...\n\n {full_path}")

# The assertion will fail if python3 executable path is not found
assert "python3" in full_path
```

## Safely find directories
```python
import subprocess

#Accepts user input
print("Enter a path to search for directories: \n")
user_input = "/home/repl/workspace/file_system"
print(f"directory to process: {user_input}")

#Pass safe user input into subprocess
with subprocess.Popen(["find", user_input, "-type", "d"], stdout=subprocess.PIPE) as find:
    result = find.stdout.readlines()
    
    #Process each line and decode it and strip it
    for line in result:
        formatted_line = line.decode("utf-8").strip()
        print(f"Found Directory: {formatted_line}")
```

## Directory summarizer
```python
import subprocess
import shlex

print("Enter a list of directories to calculate storage total: \n")
user_input = "venus mars jupiter"

# Sanitize the user input
sanitized_user_input = shlex.split(user_input)
print(f"raw_user_input: {user_input} |  sanitized_user_input: {sanitized_user_input}")

# Safely Extend the command with sanitized input
cmd = ["du", "-sh", "--total"]
cmd.extend(["/home/repl/workspace/planets/" + planet for planet in sanitized_user_input])
print(f"cmd: {cmd}")

# Print the totals out
disk_total = subprocess.run(cmd, stdout=subprocess.PIPE)
print(disk_total.stdout.decode("utf-8"))
```


# 3. Walking the file system
## Double trouble
```python
import os

# Walk the filesystem starting at the test_dir
matches = []
for root, _, files in os.walk('/home/repl/workspace/test_dir'):
    for name in files:
      	# Create the full path to the file by using os.path.join()
        fullpath = os.path.join(root, name)
        print(f"Processing file: {fullpath}")
        # Split off the extension and discard the rest of the path
        _, ext = os.path.splitext(fullpath)
        # Match the extension pattern .csv
        if ext == ".csv":
            matches.append(fullpath)
            
# Print the matches you find
print(matches)
```

## Y'all got some renaming to do
```python
import pathlib
import os

# Walk the filesystem starting at the test_dir
for root, _, files in os.walk('cattle'):
    for name in files:
      	
        # Create the full path to the file by using os.path.join()
        fullpath = os.path.join(root, name)
        print(f"Processing file: {fullpath}")
        
        # Rename file
        if "shorthorn" in name:
            p = pathlib.Path(fullpath)
            shortname = name.split("_")[0] # You need to split the name by underscore
            new_name = f"{shortname}_longhorn"
            print(f"Renaming file {name} to {new_name}")
            p.rename(new_name)
```

## Sweet pickle
```python
import os
from my.models.pca import X_digits
from sklearn.externals import joblib

# Walk the filesystem starting at the my path
for root, _, files in os.walk('my'):
    for name in files:
      	# Create the full path to the file by using os.path.join()
        fullpath = os.path.join(root, name)
        print(f"Processing file: {fullpath}")
        _, ext = os.path.splitext(fullpath)
        # Match the extension pattern .joblib
        if ext == ".joblib":
            clf = joblib.load(fullpath)
            break

# Predict from pickled model
print(clf.transform(X_digits))
```

## Rogue founder code
```python
import pathlib
import os

path = pathlib.Path("/home/repl/workspace/prod")
matches = sorted(path.glob("*.jar"))
for match in matches:
  print(f"Found rogue .jar file in production: {match}")
```

## Is this pattern True?
```python
import fnmatch

# List of file names to process
files = ["data1.csv", "script.py", "image.png", "data2.csv", "all.py"]

# Function that returns 
def csv_matches(list_of_files):
    """Return matches for csv files"""

    matches = fnmatch.filter(list_of_files, "*.csv")
    return matches

# Call function to find matches
matches = csv_matches(files)
print(f"Found matches: {matches}")
```

## Goons over my shammy
```python
import tempfile
import os

# Create a self-destructing temporary file
with tempfile.NamedTemporaryFile() as exploding_file:
  	# This file will be deleted automatically after the with statement block
    print(f"Temp file created: {exploding_file.name}")
    exploding_file.write(b"This message will self-destruct in 5....4...\n")
    
    # Get to the top of the file
    exploding_file.seek(0)

    #Print the message
    print(exploding_file.read())

# Check to sure file self-destructed
if not os.path.exists(exploding_file.name): 
    print(f"self-destruction verified: {exploding_file.name}")
```

## Archive users
```python
from shutil import make_archive
import os

# Archive root
username = "user1"
root_dir = "/home/repl/workspace/"
apath = "/home/repl/workspace/archive"

# Archive base
final_archive_base = f"{apath}/{username}"

# Create tar and gzipped archive
make_archive(final_archive_base, "gztar", apath)

# Create zip archive
make_archive(final_archive_base, "zip", apath)

# Print out archives
print(os.listdir(apath))
```

## Does it even exist?
```python
import pathlib

# Read the index of social media posts
with open("/home/repl/workspace/posts_index.txt") as posts:
    for post in posts.readlines():
        
        # Create a pathlib object
        path = pathlib.Path(post.strip())

        # Check if the social media post still exists on disk
        if path.exists():
            print(f"Found active post: {post}")
        else:
            print(f"Post is missing: {post}")
```

## File writing one-liner
```python
from subprocess import run, PIPE
from pathlib import Path

# Find all the python files you created and print them out
for i in range(3):
    path = Path(f"/home/repl/workspace/test/file_{i}.py")
    path.write_text("#!/usr/bin/env python\n")
    path.write_text("import datetime;print(datetime.datetime.now())")
  

# Find all the python files you created and print them out
for file in Path("/home/repl/workspace/test/").glob("*.py"):
   # Gets the resolved full path
   fullpath = str(file.resolve())
   proc = run(["python3", fullpath], stdout=PIPE)
   print(proc)
```


# 4. Command line functions
## Funky clusters
```python
from sklearn.datasets.samples_generator import make_blobs
from sklearn.cluster import KMeans

# Create sample blobs from sklearn datasets
def blobs():
    X, y = make_blobs(n_samples=10, centers=3, n_features=2,random_state=0)
    return X,y
  
# Perform KMeans cluster
def cluster(X, random_state=170, num=2):
    return KMeans(n_clusters=num, random_state=random_state).fit_predict(X) # Returns cluster assignment

# Run everything:  Call both functions.  X creates the data and cluster clusters the data.
def main():
    X,_ = blobs()
    return cluster(X)
# Print the KMeans cluster assignments
print(main()) 
```

## Hello decorator
```python
from functools import wraps

def nothing(f):
    @wraps(f)
    def wrapper(*args, **kwds):
        return f(*args, **kwds)
    return wrapper

# Decorate first function
@nothing
def something():
    pass

# Decorate second function
@nothing
def another():
    pass

# Put uncalled function into a list and print name  
funcs = [something, another]
for func in funcs:
    print(f"function name: {func.__name__}")
```

## Debugging decorator
```python
from functools import wraps

# Create decorator
def debug(f):
	@wraps(f)
	def wrap(*args, **kw):
		result = f(*args, **kw)
		print(f"function name: {f.__name__}, args: [{args}], kwargs: [{kw}]")
		return result
	return wrap
  
# Apply decorator
@debug
def mult(x, y=10):
	return x*y
print(mult(5, y=5))
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

