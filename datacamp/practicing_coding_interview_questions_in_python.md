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
def retrieve_character_indices(string):
    character_indices = dict()
    # Define the 'for' loop
    for index, character in enumerate(string):
        # Update the dictionary if the key already exists
        if character in character_indices:
            character_indices[character].append(index)
        # Update the dictionary if the key is absent
        else:
            character_indices[character] = [index]
            
    return character_indices
  
print(retrieve_character_indices('enumerate an Iterable'))
```

## Traversing a DataFrame
```python
column_counts = dict()

# Traverse through the columns in the heroes DataFrame
for column_name, series in heroes.iteritems():
    # Retrieve the values stored in series in a list form
    values = list(series)
    category_counts = dict()  
    # Traverse through unique categories in values
    for category in set(values):
        # Count the appearance of category in values
        category_counts[category] = values.count(category)
    
    column_counts[column_name] = category_counts
    
print(column_counts)
```

## What is a list comprehension?
```python
# Convert the text to lower case and create a word list
words = create_word_list(spam.lower())

# Create a set storing only unique words
word_set = set(words)

# Create a dictionary that counts each word in the list
tuples = [(word, words.count(word)) for word in word_set]
word_counter = dict(tuples)

# Printing words that appear more than once
for (key, value) in word_counter.items():
    if value > 1:
        print("{}: {}".format(key, value))
```

## Prime number sequence
```python
##
def is_prime(n):    
    # Define the initial check
    if n < 2:
       return True
    # Define the loop checking if a number is not prime
    for i in range(n, 2):
        if i % 2 > 0:
            return False
    return True

##
def is_prime(n):
    # Define the initial check
    if n < 2:
       return False
    # Define the loop checking if a number is not prime
    for i in range(2, int(math.sqrt(n)) + 1):
        if n % i == 0:
            return False
    return True
    
# Filter prime numbers into the new list
primes = [num for num in cands if is_prime(num)]
print("primes = " + str(primes))
```

## Coprime number sequence
```python
##
def gcd(a, b):
    # Define the while loop as described
    while b != 0:
        temp_a = a
        a = b
        b = temp_a % b  
    # Complete the return statement
    return a

##
def gcd(a, b):
    # Define the while loop as described
    while b != 0:
        temp_a = a
        a = b
        b = temp_a % b    
    # Complete the return statement
    return a
    
# Create a list of tuples defining pairs of coprime numbers
coprimes = [(i, j) for i in list1 
                   for j in list2 if gcd(i, j) == 1]
print(coprimes)
```

## Combining iterable objects
```python
# Define a function searching for the longest word
def get_longest_word(words):
    longest_word = ''
    for word in words:
        if len(word) > len(longest_word):
            longest_word = word
    return longest_word

# Create lists with the lengths and longest words
lengths = [len(item) for item in wlist]
words = [get_longest_word(item) for item in wlist]

# Combine the resulting data into one iterable object
for item in zip(wlist, lengths, words):
    print(item)
```

## Extracting tuples
```python
# Create a list of tuples with lengths and longest words
result = [
    (len(item), get_longest_word(item)) for item in wlist
]

# Unzip the result    
lengths, words = zip(*result)

for item in zip(wlist, lengths, words):
    print(item) 
```

## Creating a DataFrame
```python
# Create a list of tuples with words and their lengths
word_lengths = [
    (item, len(item)) for items in wlist for item in items
]

# Unwrap the word_lengths
words, lengths = zip(*word_lengths)

# Create a zip object
col_names = ['word', 'length']
result = zip(col_names, [words, lengths])

# Convert the result to a dictionary and build a DataFrame
data_frame = pd.DataFrame(dict(result))
print(data_frame)
```

## Shift a string
```python
def shift_string(string, shift):
    len_string = len(string)
    # Define a for loop with the yield statement
    for idx in range(0, len_string):
        yield string[(idx - shift) % len_string]
       
# Create a generator
gen = shift_string('DataCamp', 5)

# Create a new string using the generator and print it out
string_shifted = ''.join(gen)
print(string_shifted)
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
