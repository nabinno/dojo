---
title: Optimizing Python Code with pandas
tags: python,pandas
url: https://www.datacamp.com/courses/optimizing-python-code-with-pandas
---

# 1. Selecting columns and rows efficiently
## Measuring time I
```python
# Calculate the result of the problem using formula() and print the time required
N = 1000000
fm_start_time = time.time()
first_method = formula(N)
print("Time using formula: {} sec".format(time.time() - fm_start_time))

# Calculate the result of the problem using brute_force() and print the time required
sm_start_time = time.time()
second_method = brute_force(N)
print("Time using the brute force: {} sec".format(time.time() - sm_start_time))
```

## Measuring time II
```python
# Store the time before the execution
start_time = time.time()

# Execute the operation
letlist = [wrd for wrd in words if wrd.startswith('b')]

# Store and print the difference between the start and the current time
total_time_lc = time.time() - start_time
print('Time using list comprehension: {} sec'.format(total_time_lc))

# Store the time before the execution
start_time = time.time()

# Execute the operation
letlist = []
for wrd in words:
    if wrd.startswith('b'):
        letlist.append(wrd)
        
# Print the difference between the start and the current time
total_time_fl = time.time() - start_time
print('Time using for loop: {} sec'.format(total_time_fl))
```

## Row selection: loc[] vs iloc[]
```python
# Define the range of rows to select: row_nums
row_nums = range(0, 1000)

# Select the rows using .loc[] and row_nums and record the time before and after
loc_start_time = time.time()
rows = poker_hands.loc[row_nums]
loc_end_time = time.time()

# Print the time it took to select the rows using .loc
print("Time using .loc[]: {} sec".format(loc_end_time - loc_start_time))

# Select the rows using .iloc[] and row_nums and record the time before and after
iloc_start_time = time.time()
rows = poker_hands.iloc[row_nums]
iloc_end_time = time.time()

# Print the time it took to select the rows using .iloc
print("Time using .iloc[]: {} sec".format(iloc_end_time - iloc_start_time))
```

## Column selection: .iloc[] vs by name
```python
# Use .iloc to select the first 6 columns and record the times before and after
iloc_start_time = time.time()
cols = poker_hands.iloc[:,0:6]
iloc_end_time = time.time()

# Print the time it took
print("Time using .iloc[] : {} sec".format(iloc_end_time - iloc_start_time))

# Use simple column selection to select the first 6 columns 
names_start_time = time.time()
cols = poker_hands[['S1', 'R1', 'S2', 'R2', 'S3', 'R3']]
names_end_time = time.time()

# Print the time it took
print("Time using selection by name : {} sec".format(names_end_time - names_start_time))
```

## Random row selection
```python
# Extract number of rows in dataset
N=poker_hands.shape[0]

# Select and time the selection of the 75% of the dataset's rows
rand_start_time = time.time()
poker_hands.iloc[np.random.randint(low=0, high=N, size=int(0.75 * N))]
print("Time using Numpy: {} sec".format(time.time() - rand_start_time))

# Select and time the selection of the 75% of the dataset's rows using sample()
samp_start_time = time.time()
poker_hands.sample(int(0.75 * N), axis=0, replace = True)
print("Time using .sample: {} sec".format(time.time() - samp_start_time))
```

## Random column selection
```python
# Extract number of columns in dataset
D=poker_hands.shape[1]

# Select and time the selection of 4 of the dataset's columns using NumPy
np_start_time = time.time()
poker_hands.iloc[:,np.random.randint(low=0, high=D, size=4)]
print("Time using NymPy's random.randint(): {} sec".format(time.time() - np_start_time))

# Select and time the selection of 4 of the dataset's columns using pandas
pd_start_time = time.time()
poker_hands.sample(4, axis=1)
print("Time using panda's .sample(): {} sec".format(time.time() - pd_start_time))
```

# 2. Replacing values in a DataFrame
## Replacing scalar values I
```python
# Replace Class 1 to -2 
poker_hands['Class'].replace(1, -2, inplace=True)
# Replace Class 2 to -3
poker_hands['Class'].replace(2, -3, inplace=True)

print(poker_hands[['Class', 'Explanation']])
```

## Replace scalar values II
```python
##
start_time = time.time()

# Replace all the entries that has 'FEMALE' as a gender with 'GIRL'
names['Gender'].loc[names['Gender'] == 'FEMALE'] = 'GIRL'

print("Time using .loc[]: {} sec".format(time.time() - start_time))

##
start_time = time.time()

# Replace all the entries that has 'FEMALE' as a gender with 'GIRL'
names['Gender'].replace('FEMALE', 'GIRL', inplace=True)

print("Time using .replace(): {} sec".format(time.time() - start_time))
```

## Replace multiple values I
```python
##
start_time = time.time()

# Replace all non-Hispanic ethnicities with 'NON HISPANIC'
names['Ethnicity'].loc[(names["Ethnicity"] == 'BLACK NON HISPANIC') | 
                      (names["Ethnicity"] == 'BLACK NON HISP') | 
                      (names["Ethnicity"] == 'WHITE NON HISPANIC') | 
                      (names["Ethnicity"] == 'WHITE NON HISP')] = 'NON HISPANIC'

print("Time using .loc[]: sec".format(time.time() - start_time))

##
start_time = time.time()

# Replace all non-Hispanic ethnicities with 'NON HISPANIC'
names['Ethnicity'].replace(['BLACK NON HISPANIC', 'BLACK NON HISP', 'WHITE NON HISPANIC', 'WHITE NON HISP'], 'NON HISPANIC', inplace=True)

print("Time using .replace(): {} sec".format(time.time() - start_time))
```

## Replace multiple values II
```python
start_time = time.time()

# Replace ethnicities as instructed
names['Ethnicity'].replace(['ASIAN AND PACI','BLACK NON HISP', 'WHITE NON HISP'], ['ASIAN AND PACIFIC ISLANDER','BLACK NON HISPANIC','WHITE NON HISPANIC'], inplace=True)

print("Time using .replace(): {} sec".format(time.time() - start_time))
```

## Replace single values I
```python
# Replace Royal flush or Straight flush to Flush
poker_hands.replace({'Royal flush':'Flush', 'Straight flush':'Flush'}, inplace=True)
print(poker_hands['Explanation'].head())
```

## Replace single values II
```python
# Replace the number rank by a string
names['Rank'].replace({1:'FIRST', 2:'SECOND', 3:'THIRD'}, inplace=True)
print(names.head())
```

## Replace multiple values III
```python

```

## Most efficient method for scalar replacement
```python

```

# 3. Efficient iterating
## Looping using the .iterrows() function
```python

```

## Create a generator for a pandas DataFrame
```python

```

## The iterrows() function for looping
```python

```

## Looping using the .apply() function
```python

```

## .apply() function in every cell
```python

```

## .apply() for rows iteration
```python

```

## Vectorization over pandas series
```python

```

## Why vectorization in pandas is so fast?
```python

```

## pandas vectorization in action
```python

```

## Vectorization with NumPy arrays using .values()
```python

```

## Best method of vectorization
```python

```

## Vectorization methods for looping a DataFrame
```python

```

# 4. Data manipulation using .groupby()
## Data transformation using .groupby().transform
```python

```

## The min-max normalization using .transform()
```python

```

## Transforming values to probabilities
```python

```

## Validation of normalization
```python

```

## When to use transform()?
```python

```

## Missing value imputation using transform()
```python

```

## Identifying missing values
```python

```

## Missing value imputation
```python

```

## Data filtration using the filter() function
```python

```

## When to use filtration?
```python

```

## Data filtration
```python

```

## Congratulations!
```python

```
