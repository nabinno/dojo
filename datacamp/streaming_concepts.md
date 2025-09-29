---
title: Streaming Concepts
tags: snowflake,analytics,data-modeling,data-engineering
url: https://campus.datacamp.com/courses/streaming-concepts
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
## In the event of...
```
True:
- Event-based processes what for something to occur.
- Event-based programs can contain batch components.

False:
- Event-based processes wait until all data is present.
- Event-based processes always run at a certain time or date.
```

## Welcome to the event!
```
Event:
- Button or link click
- File upload
- New user created
- The clock roled to 12:00am Tuesday

Not an event:
```

## Queue characteristics
```
Describes a queue:
- First-In First-Out
- Similar to a line
- Sometimes called a buffer
- Can be disconnected from a data pipeline

Does not describe a queue:
- Last-In First-Out
- Can retrieve items in any order
```

## To queue or not to queue
```
Use a queue:
- A client requires an order handling system for digital items that are always instock. They want each sale fulfilled in the order it was received.
- An online game that can only support a single user, but your manager asks to allow each user to click a `I ant to play` button and wait their turn.

Don't use a queue:
- Develop an image processing system, where you must process as many items as possible at a time. Ordering does not matter in this case.
- A client wants to run a raffle where user usere information is centrally stored. Once the raffle starts, they want to abtain the winner info on demand.
```

## Log stream order
```
1 Listener waits for events.
2 Parser reads the information in the event.
3 The process logic determines what to do with the data.
4 THe log writer stores the information in a file as required.
```

## Log options
```
Logging:
- A list of read queries run against a database.
- A list of database updates and deletions.
- Debugging information about the tasks performed in a program.
- Web page requests.

Something else:
```

## Batch, queue, or stream?
```
Batch:
- Processing all items received by 1pm daily.
- Creating groups of 10 random images every 2 hours.

Queue:
- Processing user comments in the order they're received, but only when the system has spare processing capacity.

Stream:
- Processing user song requests as soon as possible, without delay.
- Storing user clicks on a web application as they arrive.
```

## Log stream processor
```
An event occurs that enters the logging system.
Parse the information in the event.
Classify the type of event.
Send alerts out for appropriate event types.
Write the event to storage for later processing / review.
```




# 3 Streaming Systems
## Real-time?
```
Real-time:
- Guarantee transport time
- Higher cost
- Defined latency

Non-realtime:
- Best effort delivery
- Lowest cost
```

## Is it real this time?
```
Real-time:
- Vehicle ABS system must update within 30ms continuously.
- User website orders should be accepted within 10 seconds, otherwise display an error.
- System security logs must be written to permanent storage within 15s of receipt.

Non-real time:
- Customer data should be converted to various image sizes as cheaply as possible.
- Files should be replicated between multiple systems, based on availability.
```

## Scaling reasons
```
[x]Your streaming system currently meets all SLAs with at least a 20% time window to spare.
[ ]A processing pipeline needs to process 30% more data by next year.
[ ]You expect to add 15 new customers to the same system within 3 months.
[ ]The sales team has sold new stringent SLA agreements on a near fully loaded processing system.
```

## To vertically scale...?
```
Would improve:
- Swap solid-state drives in place of spinning hard disks
- Install faster CPUs
- Lower SLA expectation at a cheaper maintenance cost

Would not improve:
- Move to a batch process
- Ignore the SLA
```

## Upscaled out
```
Vertical:
- Replacing hard disks with solid store drives
- Swapping the CPU for a faster model

Horizontal:
- Purchasing 4 more systems for a total of 13 processing systems
- Adding more CPUs to add more processing lines
```

## SLA guarantees
```
Would help meet SLA:
- A customer needs to reduce the time required to process a given item. You swap in a faster CPU.
- A client would like to improve a CPU bound process. This system already has the fastest CPU available. You add an SSD.

Would not help meet SLA:
- A client needs to process more events simultaneously. You double the number of systems being used.
- A coworker needs to reduce the processing time for a single event. You add 3 systems to the cluster.
```

## Streaming attributes
```
True:
- Streaming data usaually represents events that have completed
- Streaming processes often require extra memory to maintain state
- Streaming processes are often implemented as immutable logs

False:
- Streaming data has a known ending point
- Streaming processes must store data on disk prior to processing
```

## Issue types
```
Missing:
- 1,2,-,4,5
- Meesage A,-,Message C,Message D,Message E

Out of order:
- Msg1,Msg2,Msg3,Msg5,Msg4
- A,B,D,C,E

Delayed:
- 1,2,-,-,3
- A,B,C,-<D
```


## Streaming challenges
```
Options:

Issues for streaming:
- Repeat messages
- Delayed messages
- Out of order messages

Issues for batching:
- Data received after cutoff interval
- Not finishing within processing interval
```




# 4 Real-World Use Cases
## Streaming truths
```
Statements:

True:
- Kafka can store messages indefinitely, assuming there is space to do so
- Spark streaming can be used to transition from batch workloads
- Celery is a queue-based system

False:
- Spark streaming can store messages indefinitely
- Celery only works on a single system
- Kafka is useful for batch processing
```

## Crossing the streams...
```
Celery:
- Resizing a user's profile image to the common sizes
- Sending a welcome email when a user signs up for a new class

Kafka:
- Storing user interactions indefinitely
- Passing emails to a group of systems all requiring a copy

Spark streaming:
- Converting a large quantity of messages from one format to another
- Processing a large quantity of messages quickly and predicting if the content is positive or negative
```

## Message components
```
Required:
- Timestamp
- Button pressed
- Client identifier

Not required:
- Background color of user's profile
- Video being played by a friend at the some time as event
- Application version
```

## Answer me this...
```
Table columns:
- client_id
- timestamp
- action_taken: Video liked, Video skipped, Video rewatched, Genre liked
- action_details

Can answer:
- How many times has a user skipped a given video?
- What are the three most active times for all users?
- How many games does the user like?

Cannot answer:
- Did a user skip a video because they don't like it?
- Does a user like the same videos as his first friend?
- Does a user like all videos from a given creator?
```

## Great order of the SLAs
```
Doorbell button is pressed
The onboard sensor detects movement
The video data is uploaded for further analysis
Current temperature data is provided
```

## Sensor scaling considerations
```
Vertical:
- A group of traffic sensors, doubling in quantity, but should process data within the same SLA. For budget reasons, only a single server can be used.
- A set of temperature sensors for a single data center, where the data should be processed more quickly

Horizontal:
- A worldwide set of connected doorbells, which should be balanced based on location.
- Power utilization sensors across a large metro area which should have no more than 10 sensors per central server.
```

## Vaccination clinic - classify areas
```
Batch process:
- Monitoring

Queued process:
- Arrival / entrance

Streaming process:
- Departure
- Vaccine administration
- Registration
```

## A new problem...
```
Batch:
- Process all data from the past hour for analytic consolidation.
- Generate customer billing statements monthly.

Queue:
- Contact customer directly to verify purchase attempts.
- Alert the customer to any issue via email / text message.

Stream:
- Write all transaction information to a log upon receipt.
- Run transactions for fraud prediction against a machine learning model.
```
