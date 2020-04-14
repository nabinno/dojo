---
title: Practicing Coding Interview Questions in Python
tags: python
url: https://www.datacamp.com/courses/practicing-coding-interview-questions-in-python
---

# 1. Python Data Structures and String Manipulation
## List methods
```python
# Remove fruits from basket2 that are present in basket1
for item in basket1:
    if item in basket2:
        basket2.remove(item)

print('Basket 1: ' + str(basket1))
print('Basket 2: ' + str(basket2))

# Transfer fruits from basket1 to basket2
while len(basket1) > len(basket2):
    item_to_transfer = basket1.pop()
    basket2.append(item_to_transfer)

print('Basket 1: ' + str(basket1))
print('Basket 2: ' + str(basket2))
```

## Operations on sets
```python
(A | (B & C)) - (D & E)
```

## Storing data in a dictionary
```python
##
circ_parab = dict()

for x in range_x:
    for y in range_y:       
        # Calculate the value for z
        z = x**2 + y**2
        # Create a new key for the dictionary
        key = (x, y)
        # Create a new key-value pair
        circ_parab[(x, y)] = z
```

## String indexing and concatenation
```python
##
def encrypt(text, key):
  
    encrypted_text = ''

    # Fill in the blanks to create an encrypted text
    for char in text.lower():
        idx = (alphabet.index(char) + key) % len(alphabet)
        encrypted_text = encrypted_text + alphabet[idx]

    return encrypted_text

# Check the encryption function with the shift equals to 10
print(encrypt("datacamp", 10))
```

## Operations on strings
```python
##
# Create a word list from the string stored in text
word_list = text.split()

##
# Create a word list from the string stored in text
word_list = text.split()

# Make every other word uppercased; otherwise - lowercased
for i in range(len(word_list)):
    if (i + 1) % 2 == 0:
        word_list[i] = word_list[i].upper()
    else:
        word_list[i] = word_list[i].lower()

##
# Create a word list from the string stored in 'text'
word_list = text.split()

# Make every other word uppercased; otherwise - lowercased
for i in range(len(word_list)):
    if (i + 1) % 2 == 0:
        word_list[i] = word_list[i].upper()
    else:
        word_list[i] = word_list[i].lower()
        
# Join the words back and form a new string
new_text = " ".join(word_list)
print(new_text)
```

## Fixing string errors in a DataFrame
```python
# Make all the values in the 'Hair color' column lowercased
heroes['Hair color'] = heroes['Hair color'].str.lower()
  
# Check the values in the 'Hair color' column
print(heroes['Hair color'].value_counts())

# Substitute 'Fmale' with 'Female' in the 'Gender' column
heroes['Gender'] = heroes['Gender'].str.replace('Fmale', 'Female')

# Check if there is no occurences of 'Fmale'
print(heroes['Gender'].value_counts())
```

## Write a regular expression
```python
# Define the pattern to search for valid temperatures
pattern = re.compile(r'[+-]?\d+\.?\d* [CF]')

# Print the temperatures out
print(re.findall(pattern, text))

# Create an object storing the matches using 'finditer()'
matches_storage = re.finditer(pattern, text)

# Loop over matches_storage and print out item properties
for match in matches_storage:
    print('matching sequence = ' + match.group(0))
    print('start index = ' + str(match.start()))
    print('end index = ' + str(match.end()))
```

## Splitting by a pattern
```python
# Compile a regular expression
pattern = re.compile(r', \d+, ')

movies_without_year = []
for movie in movies:
    # Retrieve a movie name and its director
    split_result = re.split(pattern, movie)
    # Create a new string with a movie name and its director
    movie_without_year = ', '.join(split_result)
    # Append the resulting string to movies_without_year
    movies_without_year.append(movie_without_year)
    
for movie in movies_without_year:
    print(movie)
```

# 2. Iterable objects and representatives
## enumerate()
```python

```

## Iterators
```python

```

## Traversing a DataFrame
```python

```

## What is a list comprehension?
```python

```

## Basic list comprehensions
```python

```

## Prime number sequence
```python

```

## Coprime number sequence
```python

```

## What is a zip object?
```python

```

## Combining iterable objects
```python

```

## Extracting tuples
```python

```

## Creating a DataFrame
```python

```

## What is a generator and how to create one?
```python

```

## Shift a string
```python

```

## Throw a dice
```python

```

## Generator comprehensions
```python

```


# 3. Functions and lambda expressions
## How to pass a variable number of arguments to a function?
```python

```

## Positional arguments of variable size
```python

```

## Keyword arguments of variable size
```python

```

## Combining argument types
```python

```

## What is a lambda expression?
```python

```

## Define lambda expressions
```python

```

## Converting functions to lambda expressions
```python

```

## Using a lambda expression as an argument
```python

```

## What are the functions map(), filter(), reduce()?
```python

```

## The map() function
```python

```

## The filter() function
```python

```

## The reduce() function
```python

```

## What is recursion?
```python

```

## Calculate the number of function calls
```python

```

## Calculate an average value
```python

```

## Approximate Pi with recursion
```python

```


# 4. Python for scientific computing
## What is the difference between a NumPy array and a list?
```python

```

## Incorrect array initialization
```python

```

## Accessing subarrays
```python

```

## Operations with NumPy arrays
```python

```

## How to use the .apply() method on a DataFrame?
```python

```

## Simple use of .apply()
```python

```

## Additional arguments
```python

```

## Functions with additional arguments
```python

```

## How to use the .groupby() method on a DataFrame?
```python

```

## Standard DataFrame methods
```python

```

## BMI of villains
```python

```

## NaN value imputation
```python

```

## How to visualize data in Python?
```python

```

## Explore feature relationships
```python

```

## Plot a histogram
```python

```

## Creating boxplots
```python

```

## Final thoughts
```python

```
