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

## Survey Distinct Values
```python

```

## Categorical data validation
```python

```

## Never from there, but sometimes there at last
```python

```

## Countries of affiliation
```python

```

## Distinct Values Given Filters
```python

```

## Born here, went there
```python

```

## Triple plays (mostly) all around
```python

```

## Filter Arrays using Distinct Values
```python

```

## Sharing in physics after World War II
```python

```

## Meanwhile, in other categories...
```python

```

## Organizations and prizes over time
```python

```

## Distinct As You Like It
```python

```

## Glenn, George, and others in the G.B. crew
```python

```

## Germany, then and now
```python

```

## The prized transistor
```python

```





# 3. Get Only What You Need, and Fast

## Projection
```python

```

## Shares of the 1903 Prize in Physics
```python

```

## Rounding up the G.S. crew
```python

```

## Doing our share of data validation
```python

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

