---
title: Software Engineering for Data Scientists in Python
tags: python
url: https://www.datacamp.com/courses/software-engineering-for-data-scientists-in-python
---

# 1. Software Engineering & Data Science
## Python modularity in the wild
```python
# import the numpy package
import numpy as np

# create an array class object
arr = np.array([8, 6, 7, 5, 3, 0, 9])

# use the sort method
arr.sort()

# print the sorted array
print(arr)
```

## Leveraging documentation
```python
# load the Counter function into our environment
from collections import Counter

# View the documentation for Counter.most_common
help(Counter.most_common)

# use Counter to find the top 5 most common words
top_5_words = Counter(words).most_common(5)

# display the top 5 most common words
print(top_5_words)
```

## Using pycodestyle
```python
# Import needed package
import pycodestyle

# Create a StyleGuide instance
style_checker = pycodestyle.StyleGuide()

# Run PEP 8 check on multiple files
result = style_checker.check_files(['nay_pep8.py', 'yay_pep8.py'])

# Print result of PEP 8 style check
print(result.messages)
```

## Conforming to PEP 8
```python
# Assign data to x
x = [8, 3, 4]

# Print the data
print(x)
```

## PEP 8 in documentation
```python
def print_phrase(phrase, polite=True, shout=False):
    if polite:  # It's generally polite to say please
        phrase = 'Please ' + phrase

    if shout:  # All caps looks like a written shout
        phrase = phrase.upper() + '!!'

    print(phrase)


# Politely ask for help
print_phrase('help me', polite=True)
# Shout about a discovery
print_phrase('eureka', shout=True)
```

# 2. Writing a Python Module
## Minimal package requirements
```python
# Import the package with a name that follows PEP 8
import text_analyzer, textAnalyzer, TextAnalyzer, __text_analyzer__
```

## Naming packages
```python
# Import local packages
import package
import py_package

# View the help for each package
help(package)
help(py_package)
```

## Adding functionality to your package
```python
##
# Import needed functionality
from collections import Counter

def plot_counter(counter, n_most_common=5):
  # Subset the n_most_common items from the input counter
  top_items = counter.most_common(n_most_common)
  # Plot `top_items`
  plot_counter_most_common(top_items)

##
# Import needed functionality
from collections import Counter

def sum_counters(counters):
  # Sum the inputted counters
  return sum(counters, Counter())
```

## Using your package's new functionality
```python
# Import local package
import text_analyzer

# Sum word_counts using sum_counters from text_analyzer
word_count_totals = text_analyzer.sum_counters(word_counts)

# Plot word_count_totals using plot_counter from text_analyzer
text_analyzer.plot_counter(word_count_totals)
```

## Writing requirements.txt
```python
requirements = """
matplotlib>=3.0.0
numpy==1.15.4
pandas<=0.22.0
pycodestyle
"""
```

## Creating setup.py
```python
# Import needed function from setuptools
from setuptools import setup

# Create proper setup to be used by pip
setup(name='text_analyzer',
      version='0.0.1',
      description='Perform and visualize a text anaylsis.',
      author='nabinno',
      packages=['text_analyzer'])
```

## Listing requirements in setup.py
```python
# Import needed function from setuptools
from setuptools import setup

# Create proper setup to be used by pip
setup(name='text_analyzer',
      version='0.0.1',
      description='Perform and visualize a text anaylsis.',
      author='nabinno',
      packages=['text_analyzer'],
      install_requires=['matplotlib>=3.0.0'])
```

# 3. Utilizing Classes
## Writing a class for your package
```python
# Define Document class
class Document:
    """A class for text analysis
    
    :param text: string of text to be analyzed
    :ivar text: string of text to be analyzed; set by `text` parameter
    """
    # Method to create a new instance of MyClass
    def __init__(self, text):
        # Store text parameter to the text attribute
        self.text = text
```

## Using your package's class
```python
# Import custom text_analyzer package
import text_analyzer

# Create an instance of Document with datacamp_tweet
my_document = text_analyzer.Document(text=datacamp_tweet)

# Print the text attribute of the Document instance
print(my_document.text)
```

## Writing a non-public method
```python
class Document:
  def __init__(self, text):
    self.text = text
    # Tokenize the document with non-public tokenize method
    self.tokens = self._tokenize()
    # Perform word count with non-public count_words method
    self.word_counts = self._count_words()

  def _tokenize(self):
    return tokenize(self.text)
	
  # non-public method to tally document's word counts with Counter
  def _count_words(self):
    return Counter(self.tokens)
```

## Using your class's functionality
```python
# create a new document instance from datacamp_tweets
datacamp_doc = Document(datacamp_tweets)

# print the first 5 tokens from datacamp_doc
print(datacamp_doc.tokens[:5])

# print the top 5 most used words in datacamp_doc
print(datacamp_doc.word_counts.most_common(5))
```

## Using inheritance to create a class
```python
# Define a SocialMedia class that is a child of the `Document class`
class SocialMedia(Document):
    def __init__(self, text):
        Document.__init__(self, text)
```

## Adding functionality to a child class
```python
# Define a SocialMedia class that is a child of the `Document class`
class SocialMedia(Document):
    def __init__(self, text):
        Document.__init__(self, text)
        self.hashtag_counts = self._count_hashtags()
        self.mention_counts = self._count_mentions()
        
    def _count_hashtags(self):
        # Filter attribute so only words starting with '#' remain
        return filter_word_counts(self.word_counts, first_char='#')      
    
    def _count_mentions(self):
        # Filter attribute so only words starting with '@' remain
        return filter_word_counts(self.word_counts, first_char='@')
```

## Using your child class
```python
# Import custom text_analyzer package
import text_analyzer

# Create a SocialMedia instance with datacamp_tweets
dc_tweets = text_analyzer.SocialMedia(text=datacamp_tweets)

# Print the top five most most mentioned users
print(dc_tweets.mention_counts.most_common(5))

# Plot the most used hashtags
text_analyzer.plot_counter(dc_tweets.hashtag_counts)
```

## Exploring with dir and help
```python
# Import needed package
import text_analyzer

# Create instance of document
my_doc = text_analyzer.Document(datacamp_tweets)

# Run help on my_doc's plot method
help(my_doc.plot_counts)

# Plot the word_counts of my_doc
my_doc.plot_counts()
```

## Creating a grandchild class
```python
# Define a Tweet class that inherits from SocialMedia
class Tweets(SocialMedia):
    def __init__(self, text):
        # Call parent's __init__ with super()
        super.__init__(self, text)
        # Define retweets attribute with non-public method
        self.retweets = self._process_retweets()

    def _process_retweets(self):
        # Filter tweet text to only include retweets
        retweet_text = filter_lines(self.text, first_chars='RT')
        # Return retweet_text as a SocialMedia object
        return SocialMedia(retweet_text)
```

## Using inherited methods
```python
##
# Import needed package
import text_analyzer

# Create instance of Tweets
my_tweets = text_analyzer.Tweets(datacamp_tweets)

# Plot the most used hashtags in the tweets
my_tweets.plot_counts('hashtag_counts')

##
# Import needed package
import text_analyzer

# Create instance of Tweets
my_tweets = text_analyzer.Tweets(datacamp_tweets)

# Plot the most used hashtags in the retweets
my_tweets.retweets.plot_counts('hashtag_counts')
```

# 4. Maintainability
## Documentation
```python

```

## Identifying good comments
```python

```

## Identifying proper docstrings
```python

```

## Writing docstrings
```python

```

## Readability counts
```python

```

## Using good function names
```python

```

## Using good variable names
```python

```

## Refactoring for readability
```python

```

## Unit testing
```python

```

## Using doctest
```python

```

## Using pytest
```python

```

## Documentation & testing in practice
```python

```

## Documenting classes for Sphinx
```python

```

## Identifying tools
```python

```

## Final Thoughts
```python

```
