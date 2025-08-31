---
title: Introduction to Data Modeling in Snowflake
tags: snowflake,analytics,data-modeling,data-engineering
url: https://campus.datacamp.com/courses/introduction-to-data-modeling-in-snowflake
---

# 1 Methods for Processing Data
## Batch ordering
```
Batch Process:
- Is run to completion
- Has a known size
- Runs on an interval

Other:
- Always listens for new data
- Runs forever
```

## On the scale
```
Vertical scaling:
- Faster CPU
- Faster memory
- Better networking

Horizontal scaling:
- More SSDs (Solid State Drive)
- More computers
- More CPUs
```

## Horizontally opposed
```
[x]Horizontal scaling is always more cost effective than vertical scaling.
[ ]Horizontal scaling is often managed with a processing framework.
[ ]Horizontal scaling requires connectivity between the processing systems.
[ ]Horizontal scaling is more complex than vertical scaling.
```

## Batch problems
```
[ ]There can be a noticeable delay between when data is delivered and when it is processed.
[ ]Given a large enough batch of data, it is possible to run out of storage space prior to processing.
[x]Batches of data cannot be processed in parallel.
[ ]Batches can take longer than the interval of time between processing instances.
```

## Batch scenarios
```
Appropriate for batching:
- Order processing
- Software compilation
- Machine learning training

Not appropriate for batching:
- Internet audio
- Fraud detection
```




# 2 Intro to Streaming
## Intro to event-based computing
```

```

## In the event of...
```

```

## Welcome to the event!
```

```

## Queuing
```

```

## Queue characteristics
```

```

## To queue or not to queue
```

```

## Single system data streaming
```

```

## Log stream order
```

```

## Log options
```

```

## Batching vs. streaming
```

```

## Batch, queue, or stream?
```

```

## Log stream processor
```

```




# 3 Streaming Systems
## Intro to real-time streaming
```

```

## Real-time?
```

```

## Is it real this time?
```

```

## Vertically scaling streaming systems
```

```

## Scaling reasons
```

```

## To vertically scale...?
```

```

## Horizontally scaling streaming systems
```

```

## Upscaled out
```

```

## SLA guarantees
```

```

## Streaming roadblocks
```

```

## Streaming attributes
```

```

## Issue types
```

```

## Streaming challenges
```

```




# 4 Real-World Use Cases
## Popular streaming systems
```

```

## Streaming truths
```

```

## Crossing the streams...
```

```

## Real-world use case: streaming music service
```

```

## Message components
```

```

## Answer me this...
```

```

## Real-world use case: sensor data
```

```

## Great order of the SLAs
```

```

## Sensor scaling considerations
```

```

## Real-world use case: vaccination clinic
```

```

## Vaccination clinic - classify areas
```

```

## A new problem...
```

```

## Congratulations!
```

```
