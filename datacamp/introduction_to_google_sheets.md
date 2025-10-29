---
title: Introduction to Google Sheets
tags: google-spreadsheets,analytics
url: https://campus.datacamp.com/courses/introduction-to-google-sheets/cells-and-formulas
---

# 1 Cells and Formulas
## Importing data into Google Sheets
```
[ ]Data can only be imported from other spreadsheet files, such as .xlsx files.
[ ]To import data into a spreadsheet, navigate to File > Open.
[ ]When importing data, you have to create a new spreadsheet.
[x]Data from .csv, .txt, .tsv, and .xlsx files can all be imported into Google Sheets.
```

## Changing cell contents
```
Name	Weight (g)	KCal	Expires on	Price
Eggs	121	200	2018-01-21	$0.40
Sugar	125	484	2020-01-01	$2.50
Butter	65	466	2018-01-20	$2.50
Flour	190	692	2018-09-01	$0.30
Baking powder	30	10	2018-12-31	$0.10
```

## Cell ranges
```
Name	Weight (g)	KCal	Expires on	Price
Eggs	130	200	2018-01-21	$0.40
Sugar	135	500	2020-01-01	$0.20
Butter	65	466	2018-01-20	$2.50
Flour	190	692	2018-09-01	$0.25
```

## Formulas
```
Name	Weight (g)	KCal	Expires on	Price	Quantity	Total Price
Eggs	121	186	2018-01-21	$0.40	2	0.8
Sugar	125	484	2020-01-01	$0.20	3	0.6
Butter	65	466	2018-01-20	$2.00	2	4
Flour	190	692	2018-09-01	$0.30	4	1.2
```

## Exponents and parentheses
```
Name	Weight (g)	KCal	Expires on	Price	New Price
Eggs	121	186	2018-01-21	$0.40	0.36
Sugar	125	484	2020-01-01	$0.20	0.16
Butter	65	466	2018-01-20	$2.00	4.84
Flour	190	692	2018-09-01	$0.30	0.25
```

## Percentages
```
Name	Weight (g)	KCal	Expires on	Price	Sales tax	Total
Eggs	121	186	2018-01-21	$0.40	$0.08	$0.48
Sugar	125	484	2020-01-01	$0.20	$0.04	$0.24
Butter	65	466	2018-01-20	$2.00	$0.42	$2.42
Flour	190	692	2018-09-01	$0.30	$0.06	$0.36
```

## Data types and formatting
```
Name	Weight (g)	KCal	Expires on	Price
Eggs	121	186	2018-01-21	$0.40	21%
Sugar	125	484	2020-01-01	$0.20
Butter	65	466	2018-01-20	$2.00
Flour	190	692	2018-09-01	$0.30
```

## Working with currencies and dates
```
Name	Weight (g)	KCal	Expires on	Price
Eggs	121	186	1/21/2018	$0
Sugar	125	484	1/1/2020	$0
Butter	65	466	1/20/2018	$2
Flour	190	692	9/1/2018	$0
```

## Comparison operators and logicals
```
Name	Weight (g)	KCal	Expires on	Price	Less than $1	Energy dense
Eggs	121	186	2018-01-21	$0.40	TRUE	FALSE
Sugar	125	484	2020-01-01	$0.20	TRUE	TRUE
Butter	65	466	2018-01-20	$2.00	FALSE	TRUE
Flour	190	692	2018-09-01	$0.30	TRUE	TRUE
```





# 2 Cell References
## Using cell references
```
Country	Continent	Population	Land Area (km2)
China	Europe	1,409,517,397	9,326,410
India	Asia	1,339,180,127	2,973,190
United States	North America	324,459,463	9,147,593
Indonesia	Asia	263,991,379	1,811,569
Brazil	South America	209,288,278	8,460,415
Pakistan	Asia	197,015,955	881,912
Nigeria	Africa	190,886,311	910,768
Bangladesh	Asia	164,669,751	130,168
Russia	Europe	143,989,754	16,377,742
Mexico	North America	129,163,276	1,943,945
```

## Mathematical operators and references
```
Country	Population	Land Area (km2)	Land Area (mi2)
China	1409517397	9,326,410	3,600,931
India	1339180127	2,973,190	1,147,950
United States	324459463	9,147,593	3,531,889
Indonesia	263991379	1,811,569	699,447
Brazil	209288278	8,460,415	3,266,569
Pakistan	197015955	881,912	340,507
Nigeria	190886311	910,768	351,648
Bangladesh	164669751	130,168	50,258
Russia	143989754	16,377,742	6,323,453
Mexico	129163276	1,943,945	750,558
```

## Comparison operators and references
```
Country	Population	Land Area (km2)	Growth	Density
China	1,409,517,397	9,326,410	15,786,595	151
India	1,339,180,127	2,973,190	14,998,817	450
United States	324,459,463	9,147,593	3,633,946	35
Indonesia	263,991,379	1,811,569	2,956,703	146
Brazil	209,288,278	8,460,415	2,344,029	25
Pakistan	197,015,955	881,912	2,206,579	223
Nigeria	190,886,311	910,768	2,137,927	210
Bangladesh	164,669,751	130,168	1,844,301	1,265
Russia	143,989,754	16,377,742	1,612,685	9
Mexico	129,163,276	1,943,945	1,446,629	66
```

## Comparison operators and references
```
Country	Continent	Population	Land Area (km2)	Density	Over Avg	North American
China	Asia	1,409,517,397	9,326,410	151	TRUE	FALSE
India	Asia	1,339,180,127	2,973,190	450	TRUE	FALSE
United States	North America	324,459,463	9,147,593	35	FALSE	TRUE
Indonesia	Asia	263,991,379	1,811,569	146	TRUE	FALSE
Brazil	South America	209,288,278	8,460,415	25	FALSE	FALSE
Pakistan	Asia	197,015,955	881,912	223	TRUE	FALSE
Nigeria	Africa	190,886,311	910,768	210	TRUE	FALSE
Bangladesh	Asia	164,669,751	130,168	1,265	TRUE	FALSE
Russia	Europe/Asia	143,989,754	16,377,742	9	FALSE	FALSE
Mexico	North America	129,163,276	1,943,945	66	TRUE	TRUE
```

## Using absolute references
```
Country	Population	Land Area (km2)	Relative population
China	1,409,517,397	9,326,410	100
India	1,339,180,127	2,973,190	95
United States	324,459,463	9,147,593	23
Indonesia	263,991,379	1,811,569	19
Brazil	209,288,278	8,460,415	15
Pakistan	197,015,955	881,912	14
Nigeria	190,886,311	910,768	14
Bangladesh	164,669,751	130,168	12
Russia	143,989,754	16,377,742	10
Mexico	129,163,276	1,943,945	9
World	7,550,262,101	148,940,000	536
```

## Absolute references: fixing rows
```
Country	Population	Land Area (km2)	Relative population	Relative area
China	1,409,517,397	9,326,410	19	6
India	1,339,180,127	2,973,190	18	2
United States	324,459,463	9,147,593	4	6
Indonesia	263,991,379	1,811,569	3	1
Brazil	209,288,278	8,460,415	3	6
Pakistan	197,015,955	881,912	3	1
Nigeria	190,886,311	910,768	3	1
Bangladesh	164,669,751	130,168	2	0
Russia	143,989,754	16,377,742	2	11
Mexico	129,163,276	1,943,945	2	1
World	7,550,262,101	148,940,000
```

## Absolute references: fixing columns
```
Country	Population	Land Area (km2)	Land Area (mi2)	Density (km2)	Density (mi2)
China	1,409,517,397	9,326,410	3,600,931	151	391
India	1,339,180,127	2,973,190	1,147,950	450	1,167
United States	324,459,463	9,147,593	3,531,889	35	92
Indonesia	263,991,379	1,811,569	699,447	146	377
Brazil	209,288,278	8,460,415	3,266,569	25	64
Pakistan	197,015,955	881,912	340,507	223	579
Nigeria	190,886,311	910,768	351,648	210	543
Bangladesh	164,669,751	130,168	50,258	1,265	3,276
Russia	143,989,754	16,377,742	6,323,453	9	23
Mexico	129,163,276	1,943,945	750,558	66	172
World	7,550,262,101	148,940,000
```

## References everywhere!
```
Country	Population	Land Area (km2)	Growth	Density (km2)	Density Growth
China	1,409,517,397	9,326,410	28,190,348	151	3
India	1,339,180,127	2,973,190	26,783,603	450	9
United States	324,459,463	9,147,593	6,489,189	35	1
Indonesia	263,991,379	1,811,569	5,279,828	146	3
Brazil	209,288,278	8,460,415	4,185,766	25	0
Pakistan	197,015,955	881,912	3,940,319	223	4
Nigeria	190,886,311	910,768	3,817,726	210	4
Bangladesh	164,669,751	130,168	3,293,395	1,265	25
Russia	143,989,754	16,377,742	2,879,795	9	0
Mexico	129,163,276	1,943,945	2,583,266	66	1
World	7,550,262,101	148,940,000	151,005,242	51	1

Growth Index	2.00%
```
