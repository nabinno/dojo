---
title: Introduction to Scala
tags: scala
url: https://www.datacamp.com/courses/introduction-to-scala
---

# 1. A Scalable Language
## What is Scala?
```scala
## True
Many of Scala's design decisions aimed to address criticisms of Java.
Scala source code is intended to be compiled to Java bytecode.
Scala is a general-purpose programming language.
Scala powers some of the world's largest websites, applications, and data engineering infrastructures.
Scala executable code runs on a Java virtual machine.

## False
Scala means "fudgesicle" in Swedish.
```

## Why use Scala?
```scala
## Scalable
Lego
Scala
Farmers' market
Bazaar

## Not scalable
The Taj Mahl
Empire State Building
Cathedral
```

## What makes Scala scalable?
```scala
## Object-oriented
Every operation is a mtehotd call.
Every value is an object.

## Functional
Operations of a program should map input values to output values rather than change data in place.
Functions are first-class values.
```

## Scala is object-oriented
```scala
// lable?// Calculate the difference between 8 and 5
val difference = 8 - 5

// Print the difference
println(difference)
```

## Define immutable variables (val)
```scala
// Define immutable variables for clubs 2♣ through 4♣
val twoClubs: Int = 2
val threeClubs: Int = 3
val fourClubs: Int = 4
```

## Don't try to change me
```scala
// Define immutable variables for player names
val playerA: String = "Alex"
val playerB: String = "Chen"
var playerC: String = "Marta"

// Change playerC from Marta to Umberto
playerC = "Umberto"
```

## Define mutable variables (var)
```scala
// Define mutable variables for all aces
var aceClubs: Int = 1
var aceDiamonds: Int = 1
var aceHearts: Int = 1
var aceSpades: Int = 1
```

## You can change me
```scala
// Create a mutable variable for Alex as player A
var playerA: String = "Alex"

// Change the point value of A♦ from 1 to 11
aceDiamonds = 11

// Calculate hand value for J♣ and A♦
println(jackClubs + aceDiamonds)
```

## Pros and cons of immutability
```txt
## Pros
You have to write fewer tests.
Your code is easier to reason about.
Your data won't be changed inadvertently.

## Cons
More memory required due to data copying.
```


# 2. Workflows, Functions, Collections
## Create and run a Scala script
```txt
Open a bank file in your text editor of choice.
Write one line of code in the file: `println("Let's play Twenty-One!")`.
Save the file in your desired working directory with the name `game.scala`.
Open a command prompt. Navigate to your desired working directory, then type `scala game.scala` and click enter.
Observe `Let's play Twenty-One!` printed to output.
```

## What do functions do?
```
Functions are invoked with a list of arguments to produce a result.
```

## Call a function
```scala
// Calculate hand values
var handPlayerA: Int = queenDiamonds + threeClubs + aceHearts + fiveSpades
var handPlayerB: Int = kingHearts + jackHearts

// Find and print the maximum hand value
println(maxHand(handPlayerA, handPlayerB))
```

## Arrays
```scala

```

## Create and parameterize an array
```scala

```

## Initialize an array
```scala

```

## An array, all at once
```scala

```

## Updating arrays
```scala

```

## Lists
```scala

```

## Initialize and prepend to a list
```scala

```

## Initialize a list using cons and Nil
```scala

```

## Concatenate Lists
```scala

```



# 3. Type Systems, Control Structures, Style
## Scala's static type system
```scala

```

## Static typing vs. dynamic typing
```scala

```

## Pros and cons of static type systems
```scala

```

## Make decisions with if and else
```scala

```

## if and printing
```scala

```

## if expressions result in a value
```scala

```

## if and else inside of a function
```scala

```

## while and the imperative style
```scala

```

## A simple while loop
```scala

```

## Loop over a collection with while
```scala

```

## foreach and the functional style
```scala

```

## Is Scala purely a functional language?
```scala

```

## Converting while to foreach
```scala

```

## Signs of style
```scala

```

## The essence of Scala
```scala

```

