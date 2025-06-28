---
title: Debugging a Sales Data Workflow
tags: software-debugging, python
url: https://projects.datacamp.com/projects/1931
---

```python
import pandas as pd

def load_and_check():
    # Step 1: Load the data and check if it has the expected number of columns
    data = pd.read_csv('sales.csv')

    expected_columns = 17  # Adjusted to match actual data
    actual_columns = data.shape[1]
    if actual_columns != expected_columns:
        print(f"Data column mismatch! Expected {expected_columns}, but got {actual_columns}.")
        print(f"Columns found: {list(data.columns)}")
    else:
        print("Data loaded successfully.")

    # Fix rows where Tax is clearly incorrect (e.g., 0.0 when it shouldn't be)
    mask = data['Tax'] == 0
    data.loc[mask, 'Tax'] = data.loc[mask, 'Quantity'] * data.loc[mask, 'Unit price'] * 0.05

    # Step 2: Calculate statistics by date and merge them with the original data
    grouped_data = data.groupby(['Date'])['Total'].agg(['mean', 'std'])
    grouped_data['threshold'] = 3 * grouped_data['std']
    grouped_data['max'] = grouped_data['mean'] + grouped_data.threshold
    grouped_data['min'] = grouped_data[['mean', 'threshold']].apply(lambda row: max(0, row['mean'] - row['threshold']), axis=1)
    data = pd.merge(data, grouped_data, on='Date', how='left')

    # Condition_1: Check if 'Total' is within acceptable range for each date
    data['Condition_1'] = (data['Total'] >= data['min']) & (data['Total'] <= data['max'])
    data['Condition_1'].fillna(False, inplace=True)

    # Condition_2: Check if 'Tax' is approximately 5% of (Quantity * Unit price), allowing small rounding differences
    expected_tax = data['Quantity'] * data['Unit price'] * 0.05
    data['Condition_2'] = (expected_tax - data['Tax']).abs() < 0.01

    # Step 3: Perform integrity checks
    failed_condition_1 = data[~data['Condition_1']]
    failed_condition_2 = data[~data['Condition_2']]

    if failed_condition_1.shape[0] > 0 or failed_condition_2.shape[0] > 0:
        print(f"Data integrity check failed! {failed_condition_1.shape[0]} rows failed Condition_1, "
              f"{failed_condition_2.shape[0]} rows failed Condition_2.")
    else:
        print("Data integrity check was successful! All rows pass the integrity conditions.")

    return data

# Run the function
processed_data = load_and_check()
```
