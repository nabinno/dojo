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

```

## Finding documents
```python

```

## "born" approximation
```python

```

## Composing filters
```python

```

## We've got options
```python

```

## Dot notation: reach into substructure
```python

```

## Choosing tools
```python

```

## Starting our ascent
```python

```

## Our 'born' approximation, and a special laureate
```python

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

