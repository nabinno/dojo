---
title: Introduction to Data Warehousing
tags: data-warehouse,data-engineering
url: https://app.datacamp.com/learn/courses/introduction-to-data-warehousing
---

# 1. Data Warehouse Basics
## Knowing the what and why
```
In the last video, we discussed what is a data warehouse, what it does, and why it is valuable to the organization.

[ ]A data warehouse cannot integrate data from different areas of an organization but must gather data from only one area
[x]Data warehouses enable organizational analysis and decision-making
[ ]A data warehouse is not used for analysis but primarily to store and back up data from other critical systems
[ ]Data warehouses are designed only to store the AI models of a data scientist or analyst
```

## Possible use cases for a data warehouse for Zynga
```
Companies are willing to invest a large amount of money into developing a data warehouse for the potential insights they can bring. For example, Zynga, the social video game company, said the following when submitting to become a publicly traded company.

Let's imagine how Zynga might use its data warehouse. Select the answer that is not a use case for how Zynga may have used its data warehouse.

[x]For new game development, by allowing employees to store, track, and collaborate on the code for their latest game
[ ]For Marketing and measuring its sales campaign effectiveness by connecting data from internal and external systems such as web analytics platforms and advertising channels.
[ ]For team performance evaluations using metrics derived from the data warehouse to create customized dashboards or reports showing team performance.
```

## Data warehouses vs. data lakes
```
Data Lake:
Holds data the purpose of which is not yet determined
Is less organized
Can include unstructured data

Data Warehouse:
Contains only structured data
Is relatively more complicated to change because of upstream and downstream impacts
Holds data only where the purpose is known
```

## Data warehouses vs. data marts
```
In the last video you learned about some of the similarities and differences between a data warehouse and a data mart.
Select the statement that is true.

[ ]A data mart is an organization-wide repository of disparate data sources, while a data warehouse is a subset of data focused on one area.
[ ]A data mart typically stores data from multiple sources, compared to a data warehouse that typically stores data from just a few sources.
[x]A data mart size is typically less than 100 GB compared to a data warehouse that is often 100 GB or more.
```

## Deciding between a data lake, warehouse, and mart
```
You are performing some initial assessment for the hypothetical home office furniture company Bravo. The company is looking to invest in its data infrastructure.
From your initial interviews, you have learned that the company has multiple databases but no single data repository for analysis. Additionally, Bravo has confirmed that they are interested in storing unstructured data in this system, such as audio and video files. They expect many users of this data will have experience using sophisticated data tools to interface with the system. Finally, they estimate their current and future data needs will be about two terabytes (i.e.,>100 GB).
Based on this information, choose if Bravo should invest in a data warehouse, data mart, or data lake.

[x]Select a data lake since they need to integrate over a terabyte of structured and unstructured data covering many different departments.
[ ]Select a data mart since they need to integrate over a terabyte of structured and unstructured data covering many different departments.
[ ]Select a data warehouse since they need to integrate over a terabyte of structured and unstructured data covering many different departments.
```

## Data warehouse life cycle
```
Business Requirements
Data Modeling
ETL Design & Development
BI Application Development
Test & Deploy
```

## Support where needed
```
Different personas are needed to support the creation and deployment of a data warehouse. Select the true statement about a persona and their involvement in the data warehouse life cycle.

[ ]Data Scientists support the ETL Design and Development step using machine learning modeling skills.
[x]Data Analysts are needed in the Business Requirements step to help gather the organizational requirements of the data warehouse
[ ]Database Administrators are needed during the BI Application Development step to coordinate access to the transactional databases.

Data Analysts are close to the data and often perform different business analyses. Therefore, they are needed to assist with gathering the business requirements.
```

## Who does what?
```
Analyst:
Assist with collecting business requirements as they understand the organization's needs
Consult and help with the setup of BI reports during the BI Application Development step

Data Engineer:
Make changes to pipelines if needed during the Maintenance step
Assist with creating the data pipelines during the ETL Design and Development step

Both:
Are needed during the Data Modeling step
```



# 2. Warehouse Architectures and Properties
## Ordering data warehouse layers
```
1. Source Layer
- Overview: This is the origin of the data. The source layer includes transactional systems, IoT devices, APIs, and external data sources.
- Role: It generates and provides data. The data collected here is raw and usually not suitable for analysis in its initial form.
- Challenges: Ensuring data quality, handling diverse data formats, and maintaining data collection frequency or real-time capabilities are the main challenges.

2. Staging Layer
- Overview: The staging layer temporarily stores raw data collected from the source layer.
- Role: It performs data cleaning, transformation, and integration before the data moves to the storage layer.
- Challenges: Efficiently transforming and integrating data, detecting and correcting errors, and ensuring scalable processing are critical.

3. Storage Layer
- Overview: The storage layer is where cleaned data is stored long-term.
- Role: It stores data in a structured format, making it accessible for querying and analysis. Storage methods include databases, data lakes, and data warehouses.
- Challenges: Ensuring data scalability, optimizing performance, and maintaining data security and privacy are essential.

4. Presentation Layer
- Overview: The presentation layer is where data is delivered to the end consumers.
- Role: It provides data visualization, reporting, and business intelligence tools for users. Through this layer, users can access data for analysis and decision-making.
- Challenges: Ensuring accuracy and usability of data visualizations, designing user-friendly interfaces, and maintaining data access performance are important.
```

## Understanding ETL
```
[ ]The transformation process in ETL can only handle sorting data and no mathematical operations.
[x]Deduplication, or removing duplicate values, can be done in the ETL process.
[ ]The data warehouse ETL process can only work with structured data sources.
```

## Pick the correct layer
```
Source layer:
Might include JSON files
Comprises all of the input data sources used in the data  warehouse

Staging layer:
Includes on ETL process that transforms the input data into a structured form
Contains a database to store data temporarily during the ETL process
Includes the data warehosue

Storage layer:
Contains data marts
```

## Stepping into a consultant's shoes
```
[x]You suggest an automated reporting or dashboarding tool because the graphical interface is easier for those without a strong coding background. The automation capabilities can save significant time since it will be needed monthly.
[ ]You suggest a data mining tool because of their ability to explore the data and no coding experience.
[ ]You suggest the sales team learn SQL and run queries against the data warehouse.- The sales team lacks coding experience and needs a tool that makes it easy to use on a monthly basis.
```

## Supporting analysts and data scientist users
```
In this exercise, select the false statement.

] Support sophisticated tools such as R and Python to query the data warehouse directly.
]xSupport automated reporting by exporting business reports during the ETL process.
] Support data exploration and pattern discovery by making the data within the warehouse available to BI or business intelligence tools.
```

## Top-down vs bottom-up
```
[x]The top-down approach has longer upfront development time before users can deliver reports compared to the bottom-up approach due to aligning the organization on all data definitions and cleaning rules.
[ ]The top-down approach moves data into the data marts first compared to the bottom-up approach because this approach is focused on delivering organizational value as quickly as possible.
[ ]The top-down approach has a higher chance of duplicate data compared to the bottom-up approach because the data is modeled for each departmental data mart.
```

## Characteristics of top-down and bottom-up
```
Top-down
- Uses normalization to store data
- Has data flowing into the data warehouse and then to the data marts
- It was popularized by Bill Inmon and is often also called the Inmon approach

Bottom-up
- It has the disadvantage of increased ETL processing time because it uses denormalization
- Has the data flowing into the data marts and then to the data warehouse
- It was popularized by Ralph Kimball and is often also called the Kimmball approach
```

## Choosing a top-down approach
```
[x]The top-down approach requires the organization to create the data definitions first, allowing them to connect the data with different parts of the organization and avoid duplicate data.
[ ]Since they are building the data marts first, they can deliver reporting and business value quickly.
[ ]Due to the data being stored in a denormalized form, it will minimize the data storage size.
```

## The OLAP data cube
```
Select the true statement about the OLAP system data cube.

[ ]The OLAP data cube only stores data in two dimensions, in rows and columns, for fast processing.
[x]An OLAP data cube with user data within the New York Times newspaper could process a metric like the number of users by country, subscription type, and year.
[ ]The OLAP data cube only allows users to drill down on different data dimensions and not aggregate the values.
[ ]The OLAP data cube is limited to only three dimensions.
```

## OLAP vs. OLTP scenarios
```
OLAP:
Determining the total sales of the North-western sales department in the previous year
Finding which salesperson had the highest total sales last month
Summarizing each day last month,  what was the last order number shipped that day

OLTP:
Recording when an order has been picked up for a food delivery service
Tracking in the database when a customer places a new order
Updating a row of data in the database when a passenger wants to be picked up for a ride-sharing service
```

## Understanding OLTP
```
[ ]OLTP systems process data in multidimensional cubes to update the database quickly
[ ]Analysts often use OLTP systems to analyze vast amounts of data
[x]OLTP systems are designed to process large volumes of simple queries quickly
```





# 3.
## Understanding facts and dimensional tables
```
[ ]A fact table holds additional attributes/characteristics about an organizational process.
[ ]A dimensional table contains measurements and metrics for the organization.
[x]A fact table used within the Apple music service to track artist albums might include columns like ArtistID, GenreID, SongLength, and UnitPrice.
[ ]An example of a fact table used by a food manufacturer to track production is a table that only contains the manufacturing plant's ID, location, name, and size.
```

## One starry and snowy night
```
[ ]The relatively few joins for a star schema make queries easy to use by organizational users.
[x]When you must join the dimensional department table via the employee dimensional table to analyze an organization's total sales through the sales fact table, this is an example of a star schema.
[ ]In a snowflake schema, at least one-dimensional table cannot be joined directly to the fact table.
```

## Fact or dimension?
```

```

## Kimball's four step process
```

```

## Ordering Kimball's steps
```

```

## Deciding on the grain
```

```

## Selecting reasonable facts
```

```

## Slowly changing dimensions
```

```

## Pop-quiz on slow changes
```

```

## Difference between type I, II, and III
```

```

## Row vs. column data store
```

```

## Categorizing row and column store scenarios
```

```

## Why is column store faster?
```

```

## Which queries are faster?
```

```




# 4. Implementation and Data Prep
## ETL and ELT
```

```

## ETL compared to ELT
```

```

## Differences between ETL and ELT
```

```

## Selecting ELT
```

```

## Data cleaning
```

```

## Cleaning operations
```

```

## Finding truth in data transformations
```

```

## Understanding data governance
```

```

## On premise and cloud data warehouses
```

```

## Knowing the differences between on-premise and cloud
```

```

## Matching implementation to justification
```

```

## Data warehouse design example
```

```

## Connecting it all
```

```

## Selecting bottom-up
```

```

## Do you know it all?
```

```

## Wrap-up
```

```
