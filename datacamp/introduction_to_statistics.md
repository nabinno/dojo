---
title: Introduction to Statistics
tags: statistics
url: https://campus.datacamp.com/courses/introduction-to-statistics/
---

# 1 Summary Statistics
## Using statistics in the real-world
```txt
Recall that statistics can help to answer specific, measurable questions.

In this exercise, you have been provided with several real-world scenarios and need to select which one can be solved through the application of statistics.

[ ]Why do some people prefer dogs to cats?
[x]Testing whether a new model of car is safer than the current model?
[ ]What factors make one TV show more popular than another?
[ ]What will tomorrow's winning lottery numbers be?
```

## Identifying data types
```txt
Continuous:
- Race car lap time
- Height in centimeters

Nominal:
- Soccer player position
- Employment status

Ordinal:
- Customer satisfaction - unsatified, neutral, satisfied
- Income status - Low, middle, high
```

## Descriptive vs. Inferential statistics
```txt
Descriptive:
- Given data on all purchases, what is the averge amount spent per customer?
- Given data on all projects completed, what percentage were delivered on time?

Inferential:
- Given test scores for 50 students at one high school, what is the average score for all students in that school?
- After surveying 100 adults on how many hours of sleep they get per night, what percentage of all adults sleep for at least 6 hours per night?
```

## Typical number of robberies per London Borough
```txt
Here are three definitions representing the mode, median, and mean, along with their respective values for robberies in the London crimes dataset:
- Add all values and divide by the number of observations: 1496.16
- The London Borough where Robbery occurs most frequently: Westminster
- Sort all the data and take the middle value: 1354.5

[ ]Mean: 1496.16 ; Median: Westminster
[ ]Mean: 1354.50 ; Median: 1496.16
[x]Mean: 1496.16 ; Median: 1354.50
[ ]Mean: Westminster ; Median: 1354.50
```

## Choosing a measure
```txt
Selecting the correct measure of center is essential when describing a typical value of the data being observed.

An app has been displayed, which shows a histogram of robberies in London. You can use the app to change which type of crime is displayed, and whether to include a dotted line for the mean or median.

Your task is to decide which type of crime has a symmetrical histogram and can be accurately described by both the mean and median.

[ ]Robbery
[ ]Drug Offenses
[ ]Arson and Criminal Damage
[ ]Public Order Offenses
[x]None of these crimes

---

Key Insight:
None of the crime types have a symmetrical distribution, so both the mean and median cannot be used together to describe a typical value.

Why?
* Crime data across London boroughs is unevenly distributed
* A few central areas (e.g., Westminster) have very high values
* Most other areas have relatively low values

=> This creates a right-skewed (positively skewed) distribution

Impact on Measures of Center:
Mean (average)
=> Pulled upward by extreme high values
=> Overestimates the “typical” value

Median (middle value)
=> Less affected by outliers
=> More representative of the typical borough

Conclusion:
* Mean ≠ Median
* Distribution is not symmetrical
* Therefore, none of the crime types can be accurately described using both measures

Takeaway:
=> In skewed distributions, the median is usually the more reliable measure of center.
```

## London Boroughs with most frequent crimes
```txt
London Boroughs with most frequent crimes
The mean and median are great for summarizing numeric data, but if you want to understand the typical value of a categorical variable then these measures can't be applied.

An app containing the count of various crimes for each London Borough has been provided. You can use the arrows next to the column names to sort the data from smallest to largest and vice versa.

Using this, your task is to find out which London Boroughs are the mode for Vehicle Offenses and Burglary.

Instructions
[ ]Vehicle Offenses: Kingston upon Thames ; Burglary: Tower Hamlets
[ ]Vehicle Offenses: Hackney ; Burglary: Kingston upon Thames
[ ]Vehicle Offenses: Enfield ; Burglary: Greenwich
[x]Vehicle Offenses: Enfield ; Burglary: Tower Hamlets
```

## Defining measures of spread
```txt
Range:
- This measure shows how spread out the data is between the smallest and largest values
- This measure is calculated by finding the difference between the minimum and the maximum values of a variable

Variance:
- This measure uses squared units of a variable, so it is generally used as an intermediate step when calculating the standard deviation
- To calculate this measure, square the distances of each data point from the mean and add up the results

Standard Deviation:
- This emasure uses the variables units to show the distribution of data points from the mean
- This measure is found by taking the square root of the sum of squared distances from each data point to the mean
```

## Box plots for measuring spread
```txt
Data visualization can be useful in highlighting measures of spread, such as the interquartile range (IQR).

Below is a box plot displaying the number of crimes across all London Boroughs in February 2021, grouped by the type of crime.

Your task is to use the plot to determine which type of crime had the largest interquartile range for this month.

# Select one answer
[ ]Possession of Weapons
[ ]Vehicle offenses
[ ]Drug offenses
[x]Theft

---

Look at the length of the box for each category. The box shows the interquartile range, which represents how spread out the middle values are.

The category with the longest box has the largest interquartile range.

In this plot, Theft has the longest box.

So the answer is Theft.
```

## Which crime has the larger standard deviation
```txt
[ ]It's not possible to determine which has the larger standard deviation based on the plots.
[ ]Miscellaneous Crimes Against Society
[x]Public Order Offenses
```




# 2 Probability and distributions
## What is more likely?
```
Four scenarios have been provided; your task is to choose which one is most likely to occur.

[ ]Scoring 90% or more in an exam, when only 5% of students score this highly.
[ ]A restaurant receiving an order for a steak, where steak orders have made up 100 out of 1000 total orders to date.
[x]Picking a red card out of a standard pack of 52 playing cards, where half of the cards are red and the other half are black.
[ ]Drawing a person's name out of a hat that also contains 19 other names.
```

## Chances of the next sale being more than the mean
```txt
[ ]50.37%
[x]22.98%
[ ]198.54%

The key is to use the count of orders above the mean, not the median. The table shows:

- Orders ≥ median: 890 (this gives about 50.37%, which is incorrect here)
- Orders ≥ mean ($188.50): 406
- Total orders: 1767

The required probability is: 406 ÷ 1767 ≈ 22.98%

The confusion arises because the median splits data roughly in half, while the mean is higher due to a few large orders (right-skewed distribution). Therefore, fewer orders exceed the mean.

Answer: 22.98%
```

## Dependent vs. Independent events
```txt
Dependent:
- The probability of being offered a job, having mode it to the final interview stage.
- The probability of a sports team winning a championship, given they are at the top of the league currently.
- The probability of more than 100mm of rainfall, given it is currently July.

Independent:
- The probability of winning the lottery.
- The probability of flipping a acoin and getting heads, given today is Tuesday.
- The probability of rolling a six within two attempts using a fair dice.
```

## Orders of more than 10 basket products
```

```

## Discrete distributions
```

```

## Identifying distributions
```

```

## Sample mean vs. Theoretical mean
```

```

## Continuous distributions
```

```

## Discrete vs. Continuous distributions
```

```

## Finding the normal distribution
```

```

## Probability with a uniform distribution
```

```




# 3 More Distributions and the Central Limit Theorem
## The binomial distribution
```

```

## Recognizing a binomial distribution
```

```

## How probability affects the binomial distribution
```

```

## Identifying n and p
```

```

## The normal distribution
```

```

## Recognizing the normal distribution
```

```

## What makes the normal distribution special?
```

```

## Identifying skewness
```

```

## Describing distributions using kurtosis
```

```

## The central limit theorem
```

```

## Visualizing sampling distributions
```

```

## The CLT vs. The law of large numbers
```

```

## When to use the central limit theorem
```

```

## The Poisson distribution
```

```

## Identifying Poisson processes
```

```

## Recognizing lambda in the Poisson distribution
```

```




# 4 Correlation and Hypothesis Testing
## Hypothesis testing
```

```

## Sunshine and sleep
```

```

## The hypothesis testing workflow
```

```

## Independent and dependent variables
```

```

## Experiments
```

```

## Recognizing controlled trials
```

```

## Why use randomization?
```

```

## Correlation
```

```

## Identifying correlation between variables
```

```

## What can correlation tell you?
```

```

## Confounding variables
```

```

## Interpreting hypothesis test results
```

```

## Significance levels vs. p-values
```

```

## Type I and type II errors
```

```

## Congratulations!
```

```
