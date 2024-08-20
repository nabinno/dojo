---
title: Understanding Data Visualization
tags: data-visualization, analytics
url: https://campus.datacamp.com/courses/understanding-data-visualization/visualizing-distributions
---

# 1. Visualizing distributions
## Continuous vs. categorical variables
```
continuous
- Salary of employees
- Mass of squirrels
- Percentage of questions answered correctly
- Population of towns in Canada


categorical
- Provinces of towns in Canada
- Species of squirrels
- Was the exam passed or failed?
- Job title of employees
```

## Interpreting histograms
```
0True
- The histogram is unimodal.
- The histogram is right-skewed.
- The histogram is bimodal.

False
- The most popular salary bracket is $560k to $580k.
- The histogram is left-skewed.
- The most popular salary bracket is $40k to $60k.
```

## Adjusting bin width
```
[ ]The agouti had a high level of activity from 4am to 12pm, then moderate activity from 12pm to 8pm.
[x]The agouti were most active for a couple of hours after sunrise (6:30am to 8:30am), and before sunset (4pm to 6m).
[ ]The agouti showed a constant level of activity throughout sunlight hours.
[ ]The agouti activity was highly variable, with over a dozen peaks in activity throughout the day
```

## Interpreting box plots
```
True
- In 1990, three states were considered to have extreme values in the number of packets of cigarettes smoked per capita.
- The inter-quartile range of the number of packets of cigarettes smoked per capita decreased every year from 1985 to 1995.
- The median number of packets of cigarettes smoked per capita was below 100 from 1991 onwards.

False
- The lower quartile number of packets of cigarettes smoked per capita decreased every year from 1985 to 1995.
- The upper quartile number of packets of cigarettes smoked per capita decreased every year from 1985 to 1995.
- The inter-quartile range of the number of packets of cigarettes smoked per capita was smallest in 1992.
```



# 2. Visualizing two variables
## Interpreting scatter plots
```
True
- There is a positive correlation between the life expectancy and the length of schooling.
- Exactly one country averages more than 14 years schooling.
- As the average length of schooling increases, the average life expectancy typically increases too.
- Every country with an average life expectancy of less than 60 years has an average length of schooling less than 7 years.

False
- There is a negative correlation between the life expectancy and the length of schooling.
- No countries have an average length of schooling less than 6 years and an average life expectancy of more than 75 years.
- No countries have an average life expectancy of less than 55 years.
- If one country has a longer average length of schooling than another country, that country will also have a greater average life expectancy.
```

## Interpreting line plots
```
True
- In 1930, adoption of automobiles was greater than 50%.
- In 1945, two out of the four technologies had lower adoption than in 1940.
- After 1940, adoption of refrigerators was always higher than adoption of stoves.

False
- It took longer for refrigerators to go from 50% adoption to 75% adoption than it took vacuums.
- In 1940, adoption of stoves was greater than adoption of automobiles.
- After 1940, adoption of automobiles was always higher than adoption of vacuums.
```

## Logarithmic scales for line plots
```
[ ]On Feb 3, excluding mainland China, US had the most cumulative confirmed cases of COVID-19.
[x]On Feb 17, Germany had more cumulative confirmed cases of COVID-19 than France.
[ ]On Mar 02, Iran had less than 1000 cumulative confirmed cases of COVID-19.
[ ]On Mar 16, US had less than 4000 cumulative confirmed cases of COVID-19.
```

## Line plots without dates on the x-axis
```
[ ]2009
[ ]2010
[x]2011
[ ]2012
```

## Interpreting bar plots
```
[ ]Germany had the third most famous athletes.
[ ]Five sports had more than five famous athletes.
[x]Soccer players from the USA had more famous athletes than any other country/sport combination.
[ ]There were more famous cricketers on the list than famous French athletes.
```

## Interpreting stacked bar plots
```
R> show_plot

[ ]Less than half the women aged 80+ needed assistance for two or more activities.
[ ]The group with the smallest percentage of people needing assistance for exactly one activity was men aged 75-79.
[x]The group with the largest percentage of people needing no assistance was men aged 70-74.
[ ]More than half the men aged 80+ needed assistance for at least one activity.
```

## Interpreting dot plots
```
Which statement is false?
[ ]Basketball: Russell Westbrook has more Instagram followers than Carmelo Anthony.
[ ]Cricket: Virat Kohli has more followers on Facebook than the other platforms.
[x]Soccer: Cristiano Ronaldo has more Twitter followers than Marcelo Viera.
[ ]Tennis: Maria Sharapova has more Facebook followers than Roger Federer.
```

## Sorting dot plots
```
[ ]Ukraine has the fifth most expensive Big Macs by actual price.
[x]Two countries have Big Macs that cost over 100 USD after adjusting for GDP.
[ ]After adjusting for GDP, South Africa has the cheapest Big Macs.
[ ]Azerbaijan has the fifth most expensive Big Macs by actual price.
```




# 3. The color and the shape
## Another dimension for scatter plots
```
Explore different options for distinguishing points from the four cities, then determine which statement is false.

[ ]Using different sizes or transparencies makes it hard to distinguish points that overlap.
[ ]Using separate panels provides the best way to distinguish points from each city, but makes it harder to see if there is a single trend across the whole dataset.
[x]Using different shapes provides the best way to distinguish points from each city, but makes it harder to see if there is a single trend across the whole dataset.
[ ]Using different color provides a good way to distinguish points from each city, but lighter colors can be hard to see against a white background.
```

## Another dimension for line plots
```
Explore different options for distinguishing lines from the five companies, then determine which statement is false.

[ ]All five companies began 2018 with a higher price than they began 2017.
[ ]In 2018, Facebook's stock price decreased by a greater fraction than any of the other companies.
[ ]From the start of 2019 to the start of 2020, Apple's stock price more than doubled.
[x]All five companies began 2020 with a higher price than they had half way through 2019.
```

## Eye-catching colors
```
Which statement is true?

[ ]To ensure that all data points are equally perceivable, they should all have the same color.
[ ]To ensure that all data points are equally perceivable, they should all have the same chroma.
[ ]To ensure that all data points are equally perceivable, they should all have the same luminance.
[x]To ensure that all data points are equally perceivable, choose a qualitative, sequential, or diverging scale in hue-chroma-luminance colorspace.
[ ]To ensure that all data points are equally perceivable, they should all have the same hue.
```

## Qualitative, sequential, diverging
```
[ ]None
[ ]Qualitative: Distinguish unordered categories
[ ]Sequential: Show ordering
[x]Diverging: Show above or below a midpoint```
```

## Highlighting data
```

```

## Plotting many variables at once
```

```

## Interpreting pair plots
```

```

## Interpreting correlation heatmaps
```

```

## Interpreting parallel coordinates plots
```

```




# 4. 99 problems but a plot ain't one of them
## Polar coordinates
```

```

## Pie plots
```

```

## Rose plots
```

```

## Axes of evil
```

```

## Bar plot axes
```

```

## Dual axes
```

```

## Sensory overload
```

```

## Chartjunk
```

```

## Multiple plots
```

```

## Congratulations
```

```
