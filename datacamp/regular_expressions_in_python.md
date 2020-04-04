---
title: Regular Expressions in Python
tags: python, regular-expression
url: https://www.datacamp.com/courses/regular-expressions-in-python
---

# 1. Basic Concepts of String Manipulation
## First day!
```python
# Find characters in movie variable
length_string = len(movie)

# Convert to string
to_string = str(length_string)

# Predefined variable
statement = "Number of characters in this review:"

# Concatenate strings and print result
print(statement, to_string)
```

## Artificial reviews
```python
# Select the first 32 characters of movie1
first_part = movie1[:32]

# Select from 43rd character to the end of movie1
last_part = movie1[42:]

# Select from 33rd to the 42nd character
middle_part = movie2[32:42]

# Print concatenation and movie2 variable
print(first_part+last_part+middle_part) 
print(movie2)
```

## Palindromes
```python
# Get the word
movie_title = movie[11:30]

# Obtain the palindrome
palindrome = movie_title[::-1]

# Print the word if it's a palindrome
if movie_title == palindrome:
	print(movie_title)
```

## String operations
```python
# Convert to lowercase and print the result
movie_lower = movie.lower()
print(movie_lower)

# Remove whitespaces and print the result
movie_no_space = movie_lower.strip("$")
print(movie_no_space)

# Split the string into substrings and print the result
movie_split = movie_no_space.split()
print(movie_split)

# Select root word and print the result
word_root = movie_split[1][:-1]
print(word_root)
```

## Time to join!
```python
# Remove tags happening at the end and print results
movie_tag = movie.rstrip("<\i>")
print(movie_tag)

# Split the string using commas and print results
movie_no_comma = movie_tag.split(",")
print(movie_no_comma)

# Join back together and print results
movie_join = " ".join(movie_no_comma)
print(movie_join)
```

## Split lines or split the line?
```python
# Split string at line boundaries
file_split = file.splitlines()

# Print file_split
print(file_split)

# Complete for-loop to split by commas
for substring in file_split:
    substring_split = substring.split(',')
    print(substring_split)
```

## Finding a substring
```python
for movie in movies:
  	# Find if actor occurrs between 37 and 41 inclusive
    if movie.find("actor", 37, 42) == -1:
        print("Word not found")
    # Count occurrences and replace two by one
    elif movie.count("actor") == 2:  
        print(movie.replace("actor actor", "actor"))
    else:
        # Replace three occurrences by one
        print(movie.replace("actor actor actor", "actor"))
```

## Where's the word?
```python

```

## Replacing negations
```python

```


# 2. Formatting Strings
## Positional formatting
```python

```

## Put it in order!
```python

```

## Calling by its name
```python

```

## What day is today?
```python

```

## Formatted string literal
```python

```

## Literally formatting
```python

```

## Make this function
```python

```

## On time
```python

```

## Template method
```python

```

## Preparing a report
```python

```

## Identifying prices
```python

```

## Playing safe
```python

```


# 3. Regular Expressions for Pattern Matching
## Introduction to regular expressions
```python

```

## Are they bots?
```python

```

## Find the numbers
```python

```

## Match and split
```python

```

## Repetitions
```python

```

## Everything clean
```python

```

## Some time ago
```python

```

## Getting tokens
```python

```

## Regex metacharacters
```python

```

## Finding files
```python

```

## Give me your email
```python

```

## Invalid password
```python

```

## Greedy vs. non-greedy matching
```python

```

## Understanding the difference
```python

```

## Greedy matching
```python

```

## Lazy approach
```python

```
