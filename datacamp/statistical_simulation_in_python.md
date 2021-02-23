---
title: Statistical Simulation in Python
tags: statistics, python
url: https://www.datacamp.com/courses/statistical-simulation-in-python
---

# 1. Basics of randomness & simulation
## Poisson random variable
```python
# Initialize seed and parameters
np.random.seed(123) 
lam, size_1, size_2 = 5, 3, 1000  

# Draw samples & calculate absolute difference between lambda and sample mean
samples_1 = np.random.poisson(lam, size_1)
samples_2 = np.random.poisson(lam, size_2)
answer_1 = abs(lam - samples_1.mean())
answer_2 = abs(lam - samples_2.mean())

print("|Lambda - sample mean| with {} samples is {} and with {} samples is {}. ".format(size_1, answer_1, size_2, answer_2))
```

## Shuffling a deck of cards
```python
# Shuffle the deck
np.random.shuffle(deck_of_cards) 

# Print out the top three cards
card_choices_after_shuffle = deck_of_cards[0:3]
print(card_choices_after_shuffle)
```

## Throwing a fair die
```python
# Define die outcomes and probabilities
die, probabilities, throws = [1,2,3,4,5,6], [1/6, 1/6, 1/6, 1/6, 1/6, 1/6], 1

# Use np.random.choice to throw the die once and record the outcome
outcome = np.random.choice(die, size=1, p=probabilities)
print("Outcome of the throw: {}".format(outcome[0]))
```

## Throwing two fair dice
```python
# Initialize number of dice, simulate & record outcome
die, probabilities, num_dice = [1,2,3,4,5,6], [1/6, 1/6, 1/6, 1/6, 1/6, 1/6], 2
outcomes = np.random.choice(die, size=num_dice, p=probabilities)

# Win if the two dice show the same number
if len(outcomes) == 2:
    answer = 'win'
else:
    answer = 'lose'

print("The dice show {} and {}. You {}!".format(outcomes[0], outcomes[1], answer))
```

## Simulating the dice game
```python
# Initialize model parameters & simulate dice throw
die, probabilities, num_dice = [1,2,3,4,5,6], [1/6, 1/6, 1/6, 1/6, 1/6, 1/6], 2
sims, wins = 100, 0

for i in range(sims):
    outcomes = np.random.choice(die, size=num_dice, p=probabilities) 
    # Increment `wins` by 1 if the dice show same number
    if outcomes[0] == outcomes[1]:
        wins = wins + 1

print("In {} games, you win {} times".format(sims, wins))
```

## Simulating one lottery drawing
```python
# Pre-defined constant variables
lottery_ticket_cost, num_tickets, grand_prize = 10, 1000, 10000

# Probability of winning
chance_of_winning = 1/num_tickets

# Simulate a single drawing of the lottery
gains = [-lottery_ticket_cost, grand_prize-lottery_ticket_cost]
probability = [1-chance_of_winning, chance_of_winning]
outcome = np.random.choice(a=gains, size=1, p=probability, replace=True)

print("Outcome of one drawing of the lottery is {}".format(outcome))
```

## Should we buy?
```python
# Initialize size and simulate outcome
lottery_ticket_cost, num_tickets, grand_prize = 10, 1000, 10000
chance_of_winning = 1/num_tickets
size = 2000
payoffs = [-lottery_ticket_cost, grand_prize-lottery_ticket_cost]
probs = [1-chance_of_winning, chance_of_winning]

outcomes = np.random.choice(a=payoffs, size=size, p=probs, replace=True)

# Mean of outcomes.
answer = outcomes.mean()
print("Average payoff from {} simulations = {}".format(size, answer))
```

## Calculating a break-even lottery price
```python
# Initialize simulations and cost of ticket
sims, lottery_ticket_cost = 3000, 0

# Use a while loop to increment `lottery_ticket_cost` till average value of outcomes falls below zero
while 1:
    outcomes = np.random.choice([-lottery_ticket_cost, grand_prize-lottery_ticket_cost],
                 size=sims, p=[1-chance_of_winning, chance_of_winning], replace=True)
    if outcomes.mean() < 0:
        break
    else:
        lottery_ticket_cost += 1
answer = lottery_ticket_cost - 1

print("The highest price at which it makes sense to buy the ticket is {}".format(answer))
```



# 2. Probability & data generation process
## Two of a kind
```python
# Shuffle deck & count card occurrences in the hand
n_sims, two_kind = 10000, 0
for i in range(n_sims):
    np.random.shuffle(deck_of_cards)
    hand, cards_in_hand = deck_of_cards[0:5], {}
    for [suite, numeric_value] in hand:
        # Count occurrences of each numeric value
        cards_in_hand[numeric_value] = cards_in_hand.get(numeric_value, 0) + 1
    
    # Condition for getting at least 2 of a kind
    if max(cards_in_hand.values()) >=2: 
        two_kind += 1

print("Probability of seeing at least two of a kind = {} ".format(two_kind/n_sims))
```

## Game of thirteen
```python
# Pre-set constant variables
deck, sims, coincidences = np.arange(1, 14), 10000, 0

for _ in range(sims):
    # Draw all the cards without replacement to simulate one game
    draw = np.random.choice(deck, size=13, replace=False)
    # Check if there are any coincidences
    coincidence = (draw == list(np.arange(1, 14))).any()
    if coincidence == 1:
        coincidences += 1

# Calculate probability of winning
prob_of_winning = 1 - coincidences / sims
print("Probability of winning = {}".format(prob_of_winning))
```

## The conditional urn
```python

```

## Birthday problem
```python

```

## Full house
```python

```

## Data generating process
```python

```

## Driving test
```python

```

## National elections
```python

```

## Fitness goals
```python

```

## eCommerce Ad Simulation
```python

```

## Sign up Flow
```python

```

## Purchase Flow
```python

```

## Probability of losing money
```python

```



# 3. Resampling methods
## Introduction to resampling methods
```python

```

## Sampling with replacement
```python

```

## Probability example
```python

```

## Bootstrapping
```python

```

## Running a simple bootstrap
```python

```

## Non-standard estimators
```python

```

## Bootstrapping regression
```python

```

## Jackknife resampling
```python

```

## Basic jackknife estimation - mean
```python

```

## Jackknife confidence interval for the median
```python

```

## Permutation testing
```python

```

## Generating a single permutation
```python

```

## Hypothesis testing - Difference of means
```python

```

## Hypothesis testing - Non-standard statistics
```python

```



# 4. Advanced Applications of Simulation
## Simulation for Business Planning
```python

```

## Modeling Corn Production
```python

```

## Modeling Profits
```python

```

## Optimizing Costs
```python

```

## Monte Carlo Integration
```python

```

## Integrating a Simple Function
```python

```

## Calculating the value of pi
```python

```

## Simulation for Power Analysis
```python

```

## Factors influencing Statistical Power
```python

```

## Power Analysis - Part I
```python

```

## Power Analysis - Part II
```python

```

## Applications in Finance
```python

```

## Portfolio Simulation - Part I
```python

```

## Portfolio Simulation - Part II
```python

```

## Portfolio Simulation - Part III
```python

```

## Wrap Up
```python

```

