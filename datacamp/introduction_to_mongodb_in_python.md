---
title: Introduction to MongoDB in Python
tags: python,database
url: https://www.datacamp.com/courses/introduction-to-using-mongodb-for-data-science-with-python
---

# 1. Flexibly Structured Data
## Count documents in a collection
```python
client.nobel.prizes.count_documents({})
client.nobel.laureates.count_documents({})
```

## Listing databases and collections
```python
# Save a list of names of the databases managed by client
db_names = client.list_database_names()
print(db_names)

# Save a list of names of the collections managed by the "nobel" database
nobel_coll_names = client.nobel.list_collection_names()
print(nobel_coll_names)
```

## List fields of a document
```python
# Connect to the "nobel" database
db = client.nobel

# Retrieve sample prize and laureate documents
prize = db.prizes.find_one()
laureate = db.laureates.find_one()

# Print the sample prize and laureate documents
print(prize)
print(laureate)
print(type(laureate))

# Get the fields present in each type of document
prize_fields = list(prize.keys())
laureate_fields = list(laureate.keys())

print(prize_fields)
print(laureate_fields)
```

## "born" approximation
```python
In: db.laureates.count_documents({"born": {"$lt": "1800"}})
38

In: db.laureates.count_documents({"born": {"$lt": "1700"}})
38
```

## Composing filters
```python
##
# Create a filter for laureates who died in the USA
criteria = {"diedCountry": "USA"}

# Save the count of these laureates
count = db.laureates.count_documents(criteria)
print(count)

##
# Create a filter for laureates who died in the USA but were born in Germany
criteria = {"diedCountry": "USA",
            "bornCountry": "Germany"}

# Save the count
count = db.laureates.count_documents(criteria)
print(count)

##
# Create a filter for Germany-born laureates who died in the USA and with the first name "Albert"
criteria = {"diedCountry": "USA",
            "bornCountry": "Germany",
            "firstname": "Albert"}

# Save the count
count = db.laureates.count_documents(criteria)
print(count)
```

## We've got options
```python
##
# Save a filter for laureates born in the USA, Canada, or Mexico
criteria = { "bornCountry":
                { "$in": ["USA", "Canada", "Mexico"]}
             }

# Count them and save the count
count = db.laureates.count_documents(criteria)
print(count)

##
# Save a filter for laureates who died in the USA and were not born there
criteria = { "diedCountry": "USA",
               "bornCountry": { "$ne": "USA"},
             }

# Count them
count = db.laureates.count_documents(criteria)
print(count)
```

## Starting our ascent
```python
# Filter for laureates born in Austria with non-Austria prize affiliation
criteria = {"bornCountry": "Austria",
              "prizes.affiliations.country": {"$ne": "Austria"}}

# Count the number of such laureates
count = db.laureates.count_documents(criteria)
print(count)
```

## Our 'born' approximation, and a special laureate
```python
##
# Filter for documents without a "born" field
criteria = {"born": {"$exists": False}}

# Save count
count = db.laureates.count_documents(criteria)
print(count)

##
# Filter for laureates with at least three prizes
criteria = {"prizes.2": {"$exists": True}}

# Find one laureate with at least three prizes
doc = db.laureates.find_one(criteria)

# Print the document
print(doc)
```




# 2. Working with Distinct Values and Sets
## Categorical data validation
```python
In [1]: db.prizes.find_one()
In [2]: db.laureates.find_one()
In [3]: set(db.prizes.distinct("category")) == set(db.laureates.distinct("prizes.category"))
True
In [4]: assert set(db.prizes.distinct("category")) == set(db.laureates.distinct("prizes.category"))
```

## Never from there, but sometimes there at last
```python
# Countries recorded as countries of death but not as countries of birth
countries = set(db.laureates.distinct("diedCountry")) - set(db.laureates.distinct("bornCountry"))
print(countries)
```

## Countries of affiliation
```python
# The number of distinct countries of laureate affiliation for prizes
count = len(db.laureates.distinct("prizes.affiliations.country"))
print(count)
```

## Born here, went there
```python
db.laureates.distinct("prizes.affiliations.country", {"bornCountry": "USA"})
```

## Triple plays (mostly) all around
```python
# Save a filter for prize documents with three or more laureates
criteria = {"laureates.2": {"$exists": True}}

# Save the set of distinct prize categories in documents satisfying the criteria
triple_play_categories = set(db.prizes.distinct("category", criteria))
assert set(db.prizes.distinct("category")) - triple_play_categories == {"literature"}
```

## Sharing in physics after World War II
```python
In [1]: db.laureates.count_documents({
            "prizes": {"$elemMatch": {
                "category": "physics",
                "share": {"$ne": "1"},
                "year": {"$gte": "1945"}}}})
Out[1]: 143
In [2]: db.laureates.count_documents({
            "prizes": {"$elemMatch": {
                "category": "physics",
                "share": "1",
                "year": {"$gte": "1945"}}}})
Out[2]: 18
In [3]: 18/143
Out[3]: 0.1258741258741259
```

## Meanwhile, in other categories...
```python
# Save a filter for laureates with unshared prizes
unshared = {
    "prizes": {"$elemMatch": {
        "category": {"$nin": ["physics", "chemistry", "medicine"]},
        "share": "1",
        "year": {"$gte": "1945"},
    }}}

# Save a filter for laureates with shared prizes
shared = {
    "prizes": {"$elemMatch": {
        "category": {"$nin": ["physics", "chemistry", "medicine"]},
        "share": {"$ne": "1"},
        "year": {"$gte": "1945"},
    }}}

ratio = db.laureates.count_documents(unshared) / db.laureates.count_documents(shared)
print(ratio)
```

## Organizations and prizes over time
```python
# Save a filter for organization laureates with prizes won before 1945
before = {
    "gender": "org",
    "prizes.year": {"$lt": "1945"},
    }

# Save a filter for organization laureates with prizes won in or after 1945
in_or_after = {
    "gender": "org",
    "prizes.year": {"$gte": "1945"},
    }

n_before = db.laureates.count_documents(before)
n_in_or_after = db.laureates.count_documents(in_or_after)
ratio = n_in_or_after / (n_in_or_after + n_before)
print(ratio)
```

## Glenn, George, and others in the G.B. crew
```python
In [1]: db.laureates.count_documents({"firstname": Regex("^G.+"), "surname": Regex("^S.+")})
Out[1]: 9
```

## Germany, then and now
```python
##
from bson.regex import Regex

# Filter for laureates with "Germany" in their "bornCountry" value
criteria = {"bornCountry": Regex("Germany", 0)}
print(set(db.laureates.distinct("bornCountry", criteria)))

##
from bson.regex import Regex

# Filter for laureates with a "bornCountry" value starting with "Germany"
criteria = {"bornCountry": Regex("^Germany", 0)}
print(set(db.laureates.distinct("bornCountry", criteria)))

##
from bson.regex import Regex

# Fill in a string value to be sandwiched between the strings "^Germany " and "now"
criteria = {"bornCountry": Regex("^Germany " + "\\(" + "now", 0)}
print(set(db.laureates.distinct("bornCountry", criteria)))

##
from bson.regex import Regex

#Filter for currently-Germany countries of birth. Fill in a string value to be sandwiched between the strings "now" and "$"
criteria = {"bornCountry": Regex("now" + " Germany\\)" + "$", 0)}
print(set(db.laureates.distinct("bornCountry", criteria)))
```

## The prized transistor
```python
from bson.regex import Regex

# Save a filter for laureates with prize motivation values containing "transistor" as a substring
criteria = {"prizes.motivation": Regex("transistor", 0)}

# Save the field names corresponding to a laureate's first name and last name
first, last = ["firstname", "surname"]
print([(laureate[first], laureate[last]) for laureate in db.laureates.find(criteria)])
```




# 3. Get Only What You Need, and Fast
## Shares of the 1903 Prize in Physics
```python
In [1]: db.laureates.find_one({"prizes": {"$elemMatch": {"category": "physics", "year": "1903"}}})
```

## Rounding up the G.S. crew
```python
# Use projection to select only firstname and surname
docs = db.laureates.find(
       filter= {"firstname" : {"$regex" : "^G"},
                "surname" : {"$regex" : "^S"}  },
   projection= ["firstname", "surname"]  )

# Iterate over docs and concatenate first name and surname
full_names = [doc["firstname"] + " " + doc["surname"] for doc in docs]

# Print the full names
print(full_names)
```

## Doing our share of data validation
```python
# Save documents, projecting out laureates share
prizes = db.prizes.find({}, ["laureates.share"])

# Iterate over prizes
for prize in prizes:
    # Initialize total share
    total_share = 0
    
    # Iterate over laureates for the prize
    for laureate in prize["laureates"]:
        # add the share of the laureate to total_share
        total_share += 1 / float(laureate['share'])
        
    # Print the total share
    print(total_share)
```

## Sorting
```python

```

## What the sort?
```python

```

## Sorting together: MongoDB + Python
```python

```

## Gap years
```python

```

## What are indexes?
```python

```

## High-share categories
```python

```

## Recently single?
```python

```

## Born and affiliated
```python

```

## Limits
```python

```

## Setting a new limit?
```python

```

## The first five prizes with quarter shares
```python

```

## Pages of particle-prized people
```python

```





# 4. Aggregation Pipelines: Let the Server Do It For You

## Intro to Aggregation
```python

```

## Sequencing stages
```python

```

## Aggregating a few individuals' country data
```python

```

## Passing the aggregation baton to Python
```python

```

## Aggregation Operators and Grouping
```python

```

## Field Paths and Sets
```python

```

## Organizing prizes
```python

```

## Gap years, aggregated
```python

```

## Zoom into Array Fields
```python

```

## Embedding aggregation expressions
```python

```

## Here and elsewhere
```python

```

## Countries of birth by prize category
```python

```

## Something Extra: $addFields to Aid Analysis
```python

```

## "...it's the life in your years"
```python

```

## How many prizes were awarded to immigrants?
```python

```

## Refinement: filter out "unaffiliated" people
```python

```

## Wrap-Up
```python

```

