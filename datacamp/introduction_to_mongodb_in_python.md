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

## What the sort?
```python
In [1]: docs = list(db.laureates.find(
    {"born": {"$gte": "1900"}, "prizes.year": {"$gte": "1954"}},
    {"born": 1, "prizes.year": 1, "_id": 0},
    sort=[("prizes.year", 1), ("born", -1)]))
In [2]: for doc in docs[:5]:
    print(doc)
{'born': '1916-08-25', 'prizes': [{'year': '1954'}]}
{'born': '1915-06-15', 'prizes': [{'year': '1954'}]}
{'born': '1901-02-28', 'prizes': [{'year': '1954'}, {'year': '1962'}]}
{'born': '1913-07-12', 'prizes': [{'year': '1955'}]}
{'born': '1911-01-26', 'prizes': [{'year': '1955'}]}
```

## Sorting together: MongoDB + Python
```python
from operator import itemgetter

def all_laureates(prize):  
  # sort the laureates by surname
  sorted_laureates = sorted(prize["laureates"], key=itemgetter("surname"))
  
  # extract surnames
  surnames = [laureate["surname"] for laureate in sorted_laureates]
  
  # concatenate surnames separated with " and " 
  all_names = " and ".join(surnames)
  
  return all_names

# find physics prizes, project year and name, and sort by year
docs = db.prizes.find(
           filter= {"category": "physics"}, 
           projection= ["year", "laureates.firstname", "laureates.surname"], 
           sort= [("year", 1)])

# print the year and laureate names (from all_laureates)
for doc in docs:
  print("{year}: {names}".format(year=doc["year"], names=all_laureates(doc)))
```

## Gap years
```python
# original categories from 1901
original_categories = db.prizes.distinct("category", {"year": "1901"})
print(original_categories)

# project year and category, and sort
docs = db.prizes.find(
        filter={},
        projection={"year":1, "category":1, "_id":0},
        sort=[("year", -1), ("category", 1)]
)

#print the documents
for doc in docs:
  print(doc)
```

## High-share categories
```python
In [1]: db.prizes.distinct("category", {"laureates.share": {"$gt": "3"}})
In [2]: db.prizes.find(
           filter={},
           projection={"year":1, "category":1, "_id":0},
           sort=[("laureates.share", 1), ("category", 1)]
        )
```

## Recently single?
```python
# Specify an index model for compound sorting
index_model = [("category", 1), ("year", -1)]
db.prizes.create_index(index_model)

# Collect the last single-laureate year for each category
report = ""
for category in sorted(db.prizes.distinct("category")):
    doc = db.prizes.find_one(
        {"category": category, "laureates.share": "1"},
        sort=[("year", -1)]
    )
    report += "{category}: {year}\n".format(**doc)

print(report)
```

## Born and affiliated
```python
from collections import Counter

# Ensure an index on country of birth
db.laureates.create_index([("bornCountry", 1)])

# Collect a count of laureates for each country of birth
n_born_and_affiliated = {
    country: db.laureates.count_documents({
        "bornCountry": country,
        "prizes.affiliations.country": country
    })
    for country in db.laureates.distinct("bornCountry")
}

five_most_common = Counter(n_born_and_affiliated).most_common(5)
print(five_most_common)
```

## Setting a new limit?
```python
In [1]: list(db.prizes.find({"category": "economics"},
                            {"year": 1, "_id": 0})
            .sort("year")
            .limit(3)
            .limit(5))
Out[1]: [{'year': '1969'},
         {'year': '1970'},
         {'year': '1971'},
         {'year': '1972'},
         {'year': '1973'}]
```

## The first five prizes with quarter shares
```python
from pprint import pprint

# Fetch prizes with quarter-share laureate(s)
filter_ = {"laureates.share": "4"}

# Save the list of field names
projection = ["category", "year", "laureates.motivation"]

# Save a cursor to yield the first five prizes
cursor = db.prizes.find(filter_, projection).sort("year", 1).limit(5)
pprint(list(cursor))
```

## Pages of particle-prized people
```python
from pprint import pprint

# Write a function to retrieve a page of data
def get_particle_laureates(page_number=1, page_size=3):
    if page_number < 1 or not isinstance(page_number, int):
        raise ValueError("Pages are natural numbers (starting from 1).")
    particle_laureates = list(
        db.laureates.find(
            {"prizes.motivation": {"$regex": "particle"}},
            ["firstname", "surname", "prizes"])
        .sort([("prizes.year", 1), ("surname", 1)])
        .skip(page_size * (page_number - 1))
        .limit(page_size))
    return particle_laureates

# Collect and save the first nine pages
pages = [get_particle_laureates(page_number=page) for page in range(1,9)]
pprint(pages[0])
```



# 4. Aggregation Pipelines: Let the Server Do It For You
## Aggregating a few individuals' country data
```python
# Translate cursor to aggregation pipeline
pipeline = [
    {"$match": {"gender": {"$ne": "org"}}},
    {"$project": {"bornCountry": 1, "prizes.affiliations.country": 1}},
    {"$limit": 3}
]

for doc in db.laureates.aggregate(pipeline):
    print("{bornCountry}: {prizes}".format(**doc))
```

## Passing the aggregation baton to Python
```python
from collections import OrderedDict
from itertools import groupby
from operator import itemgetter

original_categories = set(db.prizes.distinct("category", {"year": "1901"}))

# Save an pipeline to collect original-category prizes
pipeline = [
    {"$match": {"category": {"$in": list(original_categories)}}},
    {"$project": {"category": 1, "year": 1}},
    {"$sort": OrderedDict([("year", -1)])}
]
cursor = db.prizes.aggregate(pipeline)
for key, group in groupby(cursor, key=itemgetter("year")):
    missing = original_categories - {doc["category"] for doc in group}
    if missing:
        print("{year}: {missing}".format(year=key, missing=", ".join(sorted(missing))))
```

## Field Paths and Sets
```python
list(db.prizes.aggregate([
    {"$project": {"allThree": {"$setEquals": ["$laureates.share", ["3"]]},
                  "noneThree": {"$not": {"$setIsSubset": [["3"], "$laureates.share"]}}}},
    {"$match": {"$nor": [{"allThree": True}, {"noneThree": True}]}}]))
```

## Organizing prizes
```python
# Count prizes awarded (at least partly) to organizations as a sum over sizes of "prizes" arrays.
pipeline = [
    {"$match": {"gender": "org"}},
    {"$project": {"n_prizes": {"$size": "$prizes"}}},
    {"$group": {"_id": None, "n_prizes_total": {"$sum": "$n_prizes"}}}
]

print(list(db.laureates.aggregate(pipeline)))
```

## Gap years, aggregated
```python
from collections import OrderedDict

original_categories = sorted(set(db.prizes.distinct("category", {"year": "1901"})))
pipeline = [
    {"$match": {"category": {"$in": original_categories}}},
    {"$project": {"category": 1, "year": 1}},
    
    # Collect the set of category values for each prize year.
    {"$group": {"_id": "$year", "categories": {"$addToSet": "$category"}}},
    
    # Project categories *not* awarded (i.e., that are missing this year).
    {"$project": {"missing": {"$setDifference": [original_categories, "$categories"]}}},
    
    # Only include years with at least one missing category
    {"$match": {"missing.0": {"$exists": True}}},
    
    # Sort in reverse chronological order. Note that "_id" is a distinct year at this stage.
    {"$sort": OrderedDict([("_id", -1)])},
]
for doc in db.prizes.aggregate(pipeline):
    print("{year}: {missing}".format(year=doc["_id"],missing=", ".join(sorted(doc["missing"]))))
```

## Embedding aggregation expressions
```python
In [1]: db.laureates.count_documents({"bornCountry": {"$in": db.laureates.distinct("bornCountry")}})
Out[1]: 901

In [2]: db.laureates.count_documents({"$expr": {"$in": ["$bornCountry", db.laureates.distinct("bornCountry")]}})
Out[2]: 901

In [3]: db.laureates.count_documents({"$expr": {"$eq": [{"$type": "$bornCountry"}, "string"]}})
Out[3]: 901

In [4]: db.laureates.count_documents({"bornCountry": {"$type": "string"}})
Out[4]: 901
```

## Here and elsewhere
```python
key_ac = "prizes.affiliations.country"
key_bc = "bornCountry"
pipeline = [
    {"$project": {key_bc: 1, key_ac: 1}},

    # Ensure a single prize affiliation country per pipeline document
    {"$unwind": "$prizes"},
    {"$unwind": "$prizes.affiliations"},

    # Ensure values in the list of distinct values (so not empty)
    {"$match": {key_ac: {"$in": db.laureates.distinct(key_ac)}}},
    {"$project": {"affilCountrySameAsBorn": {
        "$gte": [{"$indexOfBytes": ["$"+key_ac, "$"+key_bc]}, 0]}}},

    # Count by "$affilCountrySameAsBorn" value (True or False)
    {"$group": {"_id": "$affilCountrySameAsBorn",
                "count": {"$sum": 1}}},
]
for doc in db.laureates.aggregate(pipeline): print(doc)
```

## Countries of birth by prize category
```python
pipeline = [
    # Unwind the laureates array
    {"$unwind": "$laureates"},
    {"$lookup": {
        "from": "laureates", "foreignField": "id",
        "localField": "laureates.id", "as": "laureate_bios"}},

    # Unwind the new laureate_bios array
    {"$unwind": "$laureate_bios"},
    {"$project": {"category": 1,
                  "bornCountry": "$laureate_bios.bornCountry"}},

    # Collect bornCountry values associated with each prize category
    {"$group": {"_id": "$category",
                "bornCountries": {"$addToSet": "$bornCountry"}}},

    # Project out the size of each category's (set of) bornCountries
    {"$project": {"category": 1,
                  "nBornCountries": {"$size": "$bornCountries"}}},
    {"$sort": {"nBornCountries": -1}},
]
for doc in db.prizes.aggregate(pipeline): print(doc)
```

## "...it's the life in your years"
```python
{"$project": {"years": 1, "firstname": 1, "surname": 1, "_id": 0}}
```

## How many prizes were awarded to immigrants?
```python
pipeline = [
    # Limit results to people; project needed fields; unwind prizes
    {"$match": {"gender": {"$ne": "org"}}},
    {"$project": {"bornCountry": 1, "prizes.affiliations.country": 1}},
    {"$unwind": "$prizes"},
  
    # Count prizes with no country-of-birth affiliation
    {"$addFields": {"bornCountryInAffiliations": {"$in": ["$bornCountry", "$prizes.affiliations.country"]}}},
    {"$match": {"bornCountryInAffiliations": False}},
    {"$count": "awardedElsewhere"},
]

print(list(db.laureates.aggregate(pipeline)))
```

## Refinement: filter out "unaffiliated" people
```python
pipeline = [
    {"$match": {"gender": {"$ne": "org"}}},
    {"$project": {"bornCountry": 1, "prizes.affiliations.country": 1}},
    {"$unwind": "$prizes"},
    {"$addFields": {"bornCountryInAffiliations": {"$in": ["$bornCountry", "$prizes.affiliations.country"]}}},
    {"$match": {"bornCountryInAffiliations": False}},
    {"$count": "awardedElsewhere"},
]

# Construct the additional filter stage
added_stage = {"$match": {"prizes.affiliations.country": {"$in": db.laureates.distinct("prizes.affiliations.country")}}}

# Insert this stage into the pipeline
pipeline.insert(3, added_stage)
print(list(db.laureates.aggregate(pipeline)))
```
