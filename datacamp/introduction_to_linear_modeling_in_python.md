---
title: Introduction to Linear Modeling in Python
tags: statistics, python
url: https://www.datacamp.com/courses/introduction-to-linear-modeling-in-python
---

# 1. Exploring Linear Trends
## Reasons for Modeling: Interpolation
```python
# Compute the total change in distance and change in time
total_distance = distances[-1] - distances[0]
total_time = times[-1] - times[0]

# Estimate the slope of the data from the ratio of the changes
average_speed = total_distance / total_time

# Predict the distance traveled for a time not measured
elapse_time = 2.5
distance_traveled = average_speed * elapse_time
print("The distance traveled is {}".format(distance_traveled))
```

## Reasons for Modeling: Extrapolation
```python
# Select a time not measured.
time = 8

# Use the model to compute a predicted distance for that time.
distance = model(time)

# Inspect the value of the predicted distance traveled.
print(distance)

# Determine if you will make it without refueling.
answer = (distance <= 400)
print(answer)
```

## Reasons for Modeling: Estimating Relationships
```python
# Complete the function to model the efficiency.
def efficiency_model(miles, gallons):
   return np.mean(miles/gallons)

# Use the function to estimate the efficiency for each car.
car1['mpg'] = efficiency_model(car1['miles'] , car1['gallons'] )
car2['mpg'] = efficiency_model(car2['miles'] , car2['gallons'] )

# Finish the logic statement to compare the car efficiencies.
if car1['mpg'] > car2['mpg'] :
    print('car1 is the best')
elif car1['mpg'] < car2['mpg'] :
    print('car2 is the best')
else:
    print('the cars have the same efficiency')
```

## Visualizing Linear Relationships
```python
# Create figure and axis objects using subplots()
fig, axis = plt.subplots()

# Plot line using the axis.plot() method
line = axis.plot(times ,distances , linestyle=" ", marker="o", color="red")

# Use the plt.show() method to display the figure
plt.show()
```

## Plotting the Model on the Data
```python
# Pass times and measured distances into model
model_distances = model(times, measured_distances)

# Create figure and axis objects and call axis.plot() twice to plot data and model distances versus times
fig, axis = plt.subplots()
axis.plot(times, measured_distances, linestyle=" ", marker="o", color="black", label="Measured")
axis.plot(times, model_distances, linestyle="-", marker=None, color="red", label="Modeled")

# Add grid lines and a legend to your plot, and then show to display
axis.grid(True)
axis.legend(loc="best")
plt.show()
```

## Visually Estimating the Slope & Intercept
```python
# Look at the plot data and guess initial trial values
trial_slope = 1
trial_intercept = 2

# input thoses guesses into the model function to compute the model values.
xm, ym = model(trial_intercept, trial_slope)

# Compare your your model to the data with the plot function
fig = plot_data_and_model(xd, yd, xm, ym)
plt.show()

# Repeat the steps above until your slope and intercept guess makes the model line up with the data.
final_slope = 1
final_intercept = 2
```

## Mean, Deviation, & Standard Deviation
```python
# Compute the deviations by subtracting the mean offset
dx = x - np.mean(x)
dy = y - np.mean(y)

# Normalize the data by dividing the deviations by the standard deviation
zx = dx / np.std(x)
zy = dy / np.std(y)

# Plot comparisons of the raw data and the normalized data
fig = plot_cdfs(dx, dy, zx, zy)
```

## Covariance vs Correlation
```python
# Compute the covariance from the deviations.
dx = x - np.mean(x)
dy = y - np.mean(y)
covariance = np.mean(dx * dy)
print("Covariance: ", covariance)

# Compute the correlation from the normalized deviations.
zx = dx / np.std(x)
zy = dy / np.std(y)
correlation = np.std(zx * zy)
print("Correlation: ", correlation)

# Plot the normalized deviations for visual inspection. 
fig = plot_normalized_deviations(zx, zy)
```

## Correlation Strength
```python
# Complete the function that will compute correlation.
def correlation(x,y):
    x_dev = x - np.mean(x)
    y_dev = y - np.mean(y)
    x_norm = x_dev / np.std(x)
    y_norm = y_dev / np.std(y)
    return np.mean(x_norm * y_norm)

# Compute and store the correlation for each data set in the list.
for name, data in data_sets.items():
    data['correlation'] = correlation(data['x'], data['y'])
    print('data set {} has correlation {:.2f}'.format(name, data['correlation']))

# Assign the data set with the best correlation.
best_data = data_sets['A']
```


# 2. Building Linear Models
## What makes a model linear
```python

```

## Terms in a Model
```python

```

## Model Components
```python

```

## Model Parameters
```python

```

## Interpreting Slope and Intercept
```python

```

## Linear Proportionality
```python

```

## Slope and Rates-of-Change
```python

```

## Intercept and Starting Points
```python

```

## Model Optimization
```python

```

## Residual Sum of the Squares
```python

```

## Minimizing the Residuals
```python

```

## Visualizing the RSS Minima
```python

```

## Least-Squares Optimization
```python

```

## Least-Squares with `numpy`
```python

```

## Optimization with Scipy
```python

```

## Least-Squares with `statsmodels`
```python

```




# 3. Making Model Predictions
## Modeling Real Data
```python

```

## Linear Model in Anthropology
```python

```

## Linear Model in Oceanography
```python

```

## Linear Model in Cosmology
```python

```

## The Limits of Prediction
```python

```

## Interpolation: Inbetween Times
```python

```

## Extrapolation: Going Over the Edge
```python

```

## Goodness-of-Fit
```python

```

## RMSE Step-by-step
```python

```

## R-Squared
```python

```

## Standard Error
```python

```

## Variation Around the Trend
```python

```

## Variation in Two Parts
```python

```




# 4. Estimating Model Parameters
## Inferential Statistics Concepts
```python

```

## Sample Statistics versus Population
```python

```

## Variation in Sample Statistics
```python

```

## Visualizing Variation of a Statistic
```python

```

## Model Estimation and Likelihood
```python

```

## Estimation of Population Parameters
```python

```

## Maximizing Likelihood, Part 1
```python

```

## Maximizing Likelihood, Part 2
```python

```

## Model Uncertainty and Sample Distributions
```python

```

## Bootstrap and Standard Error
```python

```

## Estimating Speed and Confidence
```python

```

## Visualize the Bootstrap
```python

```

## Model Errors and Randomness
```python

```

## Test Statistics and Effect Size
```python

```

## Null Hypothesis
```python

```

## Visualizing Test Statistics
```python

```

## Visualizing the P-Value
```python

```

## Course Conclusion
```python

```

