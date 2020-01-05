---
title: Creating Robust Python Workflows
tags: python,data-pre-processing
url: https://www.datacamp.com/courses/creating-robust-python-workflows
---

# 1. Python Programming Principles
## Functions and iteration
```python
def print_files(filenames):
    # Set up the loop iteration instructions
    for name in filenames:
        # Use pathlib.Path to print out each file
        print(Path(name).read_text())
        
def list_files(filenames):
    # Use pathlib.Path to read the contents of each file
    return [Path(name).read_text()
            # Obtain each name from the list of filenames
            for name in filenames]

filenames = "diabetes.txt", "boston.txt", "digits.txt", "iris.txt", "wine.txt"
print_files(filenames)
pprint(list_files(filenames))
```

## Find matches
```python
def get_matches(filename, query):
    # Filter the list comprehension using an if clause
    return [line for line in Path(filename).open() if query in line]

# Iterate over files to find all matching lines
matches = [get_matches(name, "Number of") for name in filenames]
pprint(matches)
```

## Dataset dimensions
```python
def flatten(nested_list):
    return (item 
            # Obtain each list from the list of lists
            for sublist in nested_list
            # Obtain each element from each individual list
            for item in sublist)

number_generator = (int(substring) for string in flatten(matches)
                    for substring in string.split() if substring.isdigit())
pprint(dict(zip(filenames, zip(number_generator, number_generator))))
```

## Extract words
```python
def obtain_words(string):
    # Replace non-alphabetic characters with spaces
    return "".join(char if char.isalpha() else " " for char in string).split()

def filter_words(words, minimum_length=3):
    # Remove words shorter than 3 characters
    return [word for word in words if len(word) >= minimum_length]

words = obtain_words(Path("diabetes.txt").read_text().lower())
filtered_words = filter_words(words)
pprint(filtered_words)
```

## Most frequent words
```python
def count_words(word_list):
    # Count the words in the input list
    return {word: word_list.count(word) for word in word_list}

# Create the dictionary of words and word counts
word_count_dictionary = count_words(filtered_words)

(pd.DataFrame(word_count_dictionary.items())
 .sort_values(by=1, ascending=False)
 .head()
 .plot(x=0, kind="barh", xticks=range(5), legend=False)
 .set_ylabel("")
)
plt.show()
```

## Instance method
```python
# Fill in the first parameter in the pair_plot() definition
def pair_plot(self, vars=range(3), hue=None):
    return pairplot(pd.DataFrame(self.data), vars=vars, hue=hue, kind="reg")

ScikitData.pair_plot = pair_plot

# Create the diabetes instance of the ScikitData class
diabetes = ScikitData("diabetes")
diabetes.pair_plot(vars=range(2, 6), hue=1)._legend.remove()
plt.show()
```

## Class method
```python
# Fill in the decorator for the get_generator() definition
@classmethod
# Add the first parameter to the get_generator() definition
def get_generator(cls, dataset_names):
    return map(cls, dataset_names)

ScikitData.get_generator = get_generator
dataset_generator = ScikitData.get_generator(["diabetes", "iris"])
for dataset in dataset_generator:
    dataset.pair_plot()
    plt.show()
```

# 2. Documentation and Tests
## TextFile hints
```python
class TextFile:
  	# Add type hints to TextFile"s __init__() method
    def __init__(self, name: str) -> None:
        self.text = Path(name).read_text()

	# Type annotate TextFile"s get_lines() method
    def get_lines(self) -> List[str]:
        return self.text.split("\n")

help(TextFile)
```

## MatchFinder hints
```python
class MatchFinder:
  	# Add type hints to __init__()'s strings argument
    def __init__(self, strings: List[str]) -> None:
        self.strings = strings

	# Type annotate get_matches()'s query argument
    def get_matches(self, query: Optional[str] = None) -> List[str]:
        return [s for s in self.strings if query in s] if query else self.strings

help(MatchFinder)
```

## Docstrings
```python

```

## Get matches docstring
```python

```

## Obtain words docstring
```python

```

## Reports
```python

```

## Build notebooks
```python

```

## Convert notebooks
```python

```

## Pytest
```python

```

## Parametrize
```python

```

## Raises
```python

```

# 3. Shell superpowers
## Command-line interfaces
```python

```

## Argparse nbuild()
```python

```

## Docopt nbuild()
```python

```

## Git version control
```python

```

## Commit added files
```python

```

## Commit modified files
```python

```

## Virtual environments
```python

```

## List installed packages
```python

```

## Show package information
```python

```

## Persistence and packaging
```python

```

## Pickle dataframes
```python

```

## Pickle models
```python

```

# 4. Projects, pipelines, and parallelism
## Project templates
```python

```

## Set up template
```python

```

## Create project
```python

```

## Executable projects
```python

```

## Zipapp
```python

```

## Argparse main()
```python

```

## Notebook workflows
```python

```

## Parametrize notebooks
```python

```

## Summarize notebooks
```python

```

## Parallel computing
```python

```

## Dask dataframe
```python

```

## Joblib
```python

```

## Wrap-up
```python

```
