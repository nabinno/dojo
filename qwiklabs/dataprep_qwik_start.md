---
title: Dataprep: Qwik Start
tags: google-cloud-dataprep,data-pre-processing,machine-learning
url: https://www.qwiklabs.com/focuses/584?parent=catalog
---

# Goal
- Google Cloud Dataprep

# Task
- [x] Setup and Requirements
- [x] Create a Cloud Storage bucket in your project
- [x] Initialize Cloud Dataprep
- [x] Create a flow
- [x] Import datasets
- [x] Prep the candidate file
- [x] Join the Contributions file
- [x] Summary of data
- [x] Rename columns
- [x] Congratulations!

# Supplement
## Import datasets
```
gs://dataprep-samples/us-fec

cn-2016.txt
itcont-2016.txt
```

## Prep the candidate file
1. Keep rows where(date(2016, 1, 1) <= column5) && (column5 < date(2018, 1, 1))
2. Change column6 type to String
3. where column7 is 'P'

## Join the Contributions file
1. Replace matches of `{start}"|"{end}` from all columns with ''
    - replacepatterns col: * with: '' on: ```{start}"|"{end}``` global: true`
2. Inner join with Candidate Master 2016 - 2 on column2 == column11

## Summary of data
3. Pivot and compute 3 functions grouped by 3 columns
    - `pivot value:sum(column16),average(column16),countif(column16 > 0) group: column2,column24,column8`

## Rename columns
4. Rename 6 columns
    - `rename type: manual mapping: [column24,'Candidate_Name'], [column2,'Candidate_ID'],[column8,'Party_Affiliation'], [sum_column16,'Total_Contribution_Sum'], [average_column16,'Average_Contribution_Sum'], [countif,'Number_of_Contributions']`
5. Set Average_Contribution_Sum to ROUND(Average_Contribution_Sum, 0)
    - `set col: Average_Contribution_Sum value: round(Average_Contribution_Sum)`
