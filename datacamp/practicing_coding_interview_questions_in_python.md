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
def simulate_dice_throws():
    total, out = 0, dict([(i, [0, 0]) for i in range(1, 7)])
    while True:
        # Simulate a single toss to get a new number
        num = random.randint(1, 6)
        total += 1
        # Update the number and the ratio of realizations
        out[num][0] = out[num][0] + 1
        for j in range(1, 7):
        	out[j][1] = round(out[j][0]/total, 2)
        # Yield the updated dictionary
        yield out

# Create the generator and simulate 1000 tosses
dice_simulator = simulate_dice_throws()
for i in range(1, 1001):
    print(str(i) + ': ' + str(next(dice_simulator)))
```

## Generator comprehensions
```python
##
# Rewrite func1() as a generator comprehension
gen = (i**2 for i in range(0, 10))

for item in zip(gen, func1(10)):
    print(item)

##
# Rewrite func2() as a generator comprehension
gen = (2*i for i in range(0, 10) if i%2 == 0)

for item in zip(gen, func2(20)):
    print(item)

##
# Rewrite func3() as a generator comprehension
gen = (((i, j), i + j) for i in func1(8) for j in func2(10))

for item in zip(gen, func3(8, 10)):
    print(item)
```


# 3. Functions and lambda expressions
## Positional arguments of variable size
```python
# Define the function with an arbitrary number of arguments
def sort_types(*args):
    nums, strings = [], []   
    for arg in args:
        # Check if 'arg' is a number and add it to 'nums'
        if isinstance(arg, (int, float)):
            nums.append(arg)
        # Check if 'arg' is a string and add it to 'strings'
        elif isinstance(arg, str):
            strings.append(arg)
    
    return (nums, strings)
            
print(sort_types(1.0, 'car', 'hat', 4, 5, 'tree', 0.0))
```

## Keyword arguments of variable size
```python
# Define the function with an arbitrary number of arguments
def key_types(**kwargs):
    dict_type = dict()
    # Iterate over key value pairs
    for key, value in kwargs.items():
        # Update a list associated with a key
        if type(value) in dict_type:
            dict_type[type(value)].append(key)
        else:
            dict_type[type(value)] = [key]
            
    return dict_type
  
res = key_types(a=1, b=2, c=(1, 2), d=3.1, e=4.2)
print(res)
```

## Combining argument types
```python
# Define the arguments passed to the function
def sort_all_types(*args, **kwargs):

    # Find all the numbers and strings in the 1st argument
    nums1, strings1 = sort_types(*args)
    
    # Find all the numbers and strings in the 2nd argument
    nums2, strings2 = sort_types(*kwargs.values())
    
    return (nums1 + nums2, strings1 + strings2)
  
res = sort_all_types(
	1, 2.0, 'dog', 5.1, num1 = 0.0, num2 = 5, str1 = 'cat'
)
print(res)
```

## Define lambda expressions
```python
##
# Take x and return x squared if x > 0 and 0, otherwise
squared_no_negatives = lambda x: x**2 if x > 0 else 0
print(squared_no_negatives(2.0))
print(squared_no_negatives(-1))

##
# Take a list of integers nums and leave only even numbers
get_even = lambda nums: [num for num in nums if num % 2 == 0]
print(get_even([1, 2, 3, 4, 5, 6, 7, 8, 9, 10]))

##
# Take strings s1, s2 and list their common characters
common_chars = lambda s1, s2: set(s1).intersection(set(s2))
print(common_chars('pasta', 'pizza'))
```

## Converting functions to lambda expressions
```python
##
# Convert func1() to a lambda expression
lambda1 = lambda x, y: x if x >= y else y
print(str(func1(5, 4)) + ', ' + str(lambda1(5, 4)))
print(str(func1(4, 5)) + ', ' + str(lambda1(4, 5)))

##
# Convert func2() to a lambda expression
lambda2 = lambda s: dict([(c, s.count(c)) for c in set(s)])
print(func2('DataCamp'))
print(lambda2('DataCamp'))

##
# Convert func3() to a lambda expression
lambda3 = lambda *nums: math.sqrt(sum([n**2 for n in nums]))
print(str(func3(3, 4)) + ', ' + str(lambda3(3, 4)))
print(str(func3(3, 4, 5)) + ', ' + str(lambda3(3, 4, 5)))
```

## Using a lambda expression as an argument
```python
##
# Sort words by the string length
words.sort(key=lambda s: len(s))
print(words) 

##
# Sort words by the last character in a string
words.sort(key=lambda s: s[-1])
print(words)

##
# Sort words by the total amount of certain characters
words.sort(key=lambda s: s.count('a') + s.count('b') + s.count('c'))
print(words)
```

## The map() function
```python
def my_zip(*args):
    
    # Retrieve Iterable lengths and find the minimal length
    lengths = list(map(len, args))
    min_length = min(lengths)

    tuple_list = []
    for i in range(0, min_length):
        # Append new items to the 'tuple_list'
        tuple_list.append(tuple(map(lambda x: x[i], args)))

    return tuple_list

result = my_zip([1, 2, 3], ['a', 'b', 'c', 'd'], 'DataCamp')
print(result)
```

## The filter() function
```python
##
# Exclude all the numbers from nums divisible by 3 or 5
print(nums)
fnums = filter(lambda x: x % 3 != 0 and x % 5 != 0, nums)
print(list(fnums))

##
# Return the string without its vowels
print(string)
vowels = 'AEIOUaeiou'
fstring = filter(lambda x: x not in vowels, string)
print(''.join(fstring))

##
# Filter all the spells in spells with more than two 'a's
print(spells)
fspells = filter(lambda x: x.count('a') > 2, spells)
print(list(fspells))
```

## The reduce() function
```python
##
# Reverse a string using reduce()
string = 'DataCamp'
inv_string = reduce(lambda x, y: y + x, string)
print('Inverted string = ' + inv_string) 

##
# Find common items shared among all the sets in sets
sets = [{1, 4, 8, 9}, {2, 4, 6, 9, 10, 8}, {9, 0, 1, 2, 4}]
common_items = reduce(lambda x, y: x.intersection(y), sets)
print('common items = ' + str(common_items))

##
# Convert a number sequence into a single number
nums = [5, 6, 0, 1]
num = reduce(lambda x, y: 10*x + y, nums)
print(str(nums) + ' is converted to ' + str(num))
```

## Calculate an average value
```python
# Calculate an average value of the sequence of numbers
def average(nums):
  
    # Base case
    if len(nums) == 1:
        return nums[0]
    
    # Recursive call
    n = len(nums)
    return (nums[0] + (n - 1) * average(nums[1:])) / n

# Testing the function
print(average([1, 2, 3, 4, 5]))
```

## Approximate Pi with recursion
```python
# Write an expression to get the k-th element of the series 
get_elmnt = lambda k: ((-1)**k)/(2*k+1)

def calc_pi(n):
    curr_elmnt = get_elmnt(n)
    
    # Define the base case 
    if n == 0:
    	return 4
      
    # Make the recursive call
    return 4 * curr_elmnt + calc_pi(n-1)
  
# Compare the approximated Pi value to the theoretical one
print("approx = {}, theor = {}".format(calc_pi(500), math.pi))
```


# 4. Python for scientific computing
## Incorrect array initialization
```python
spiral = []

for i in range(0, size):
    # Convert each part marked by a red arrow to a list
    spiral += list(square[i, i:size-i])
    # Convert each part marked by a green arrow to a list
    spiral += list(square[i+1:size-i, size-i-1])
    # Convert each part marked by a blue arrow to a list
    spiral += list(reversed(square[size-i-1, i:size-i-1]))
    # Convert each part marked by a magenta arrow to a list
    spiral += list(reversed(square[i+1:size-i-1, i]))
        
print(spiral)
```

## Operations with NumPy arrays
```python
##
# Substitute the code in the block 1 given the input_array1
output_array1 = list(map(lambda x: [5*i for i in x], input_array1))
print(list(map(lambda x: [5*i for i in x], input_list1)))
print(output_array1)

##
# Substitute the code in the block 2 given the input_array2
output_array2 = list(filter(lambda x: x % 2 == 0, input_array2))
print(list(filter(lambda x: x % 2 == 0, input_list2)))
print(output_array2)

##
# Substitute the code in the block 3 given the input_array3
output_array3 = [[i*i for i in j] for j in input_array3]
print([[i*i for i in j] for j in input_list3])
print(output_array3)
```

## Simple use of .apply()
```python
def prevalence(series):
    vals = list(series)
    # Create a tuple list with unique items and their counts
    itms = [(x, vals.count(x)) for x in set(series)]
    # Extract a tuple with the highest counts using reduce()
    res = reduce(lambda x, y: x if x[1] > y[1] else y, itms)
    # Return the item with the highest counts
    return res[0]

# Apply the prevalence function on the scores DataFrame
result = scores[groups_to_consider].apply(prevalence)
print(result)
```

## Additional arguments
```python
def rank(series):
    # Calculate the mean of the input series
    mean = np.mean(series)
    # Return the mean and its rank as a list
    if mean > 90:
        return [mean, 'high']
    if mean > 60:
        return [mean, 'medium']
    return [mean, 'low']

# Insert the output of rank() into new columns of scores
cols = ['math score', 'reading score', 'writing score']
scores[['mean', 'rank']] = scores[cols].apply(rank, axis=1,
                                              result_type='expand')
print(scores[['mean', 'rank']].head())
```

## Functions with additional arguments
```python
##
def rescale(series, low, high):
   # Define the expression to rescale input series
   return series * (high - low)/100 + low

# Rescale the data in cols to lie between 1 and 10
cols = ['math score', 'reading score', 'writing score'] 
scores[cols] = scores[cols].apply(rescale, args=[1, 10])
print(scores[cols].head())

##
# Redefine the function to accept keyword arguments
def rescale(series, low=0, high=100):
   return series * (high - low)/100 + low

# Rescale the data in cols to lie between 1 and 10
cols = ['math score', 'reading score', 'writing score']
scores[cols] = scores[cols].apply(rescale, low=1, high=10)
print(scores[cols].head())
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
