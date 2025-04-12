---
title: DevOps Concepts
tags: devops
url: https://campus.datacamp.com/courses/devops-concepts/
---

# 1 Introduction to DevOps
## Use Cases for DevOps
```
True:
- DevOps helps organize engineering organizations in an improved way.
- DevOps uses MVPs to increase the pace of innovation.

False:
- Traditional change management models suggest an MVP release and product improvements.
- DevOps aims to go to the market with a finished product.
```

## Which engineer?
```
Infra Engineering:
- You want your engineers to setup an access management policy that ensure only the authorized personnel can access your databases
- You want to use a cloud solution for your website and you need engineers to do the configuration

Product Engineering:
- You want to add functionality to your website that allows people to book their hotels directly from your website
- You want to integrate a car rental agency's webpage to allow users to book car rentals through your website

Data & ML Engineering:
- You want to create intelligent systems to provide suggestions to your customers using their post behaviour
- You want to capture user data on your website and record that into a database in real-time
```

## Employing the correct model
```
1. DevOps
2. DataOps
3. MLOps
```

## DevOps implementations
```
Your engineering organization is growing, and now you have Data Engineers, Machine Learning Engineers, and Software Developers. You want all of the teams to work closely and collaboratively. Also, you want your teams to use the MVP approach (building the basic functionality first, adding functionality later). Which of the below methodologies should you employ?

[ ]DataOps
[ ]MLOps
[ ]DevOps
[x]All three: DataOps, MLOps, and DevOps
```

## Project management frameworks
```
True:
- Waterfall could be a viable option for small tasks that does not have much complexity
- Kanban is all about continuous improvement and using visualizations
- Scrum and Kanban are both great implementations of the Agile philosophy

False:
- Scrum does not define specific roles but ensures collective responsibility
- Agile is a definitive project management methodology with set practices
```

## Factors to decide the best project management framework
```
Important Factors:
- The size of the team
- The budget of the project
- The complexity of the project

Unimportant / Irrelevant Factors:
- Using cloud infrastructure
- The age of the organization
```

## Choosing the best project management model
```
You have decided to use DevOps for developing the minimum viable product (MVP) for the proposed new feature, it is time to choose the best project management framework for this. The product expert in your team suggests using a mind-mapping tool to visualize the tasks and ensure collaboration within your team. You agree with his suggestion. Which project management framework below could be the most suited for this purpose?

[ ]Waterfall
[ ]Agile
[x]Scrum
[ ]Kanban
```




# 2 DevOps Architecture
## DevOps cycle components
```
Defining the deliverables and goals related to the product, the software components, their functionality and relation to each other
Writing the code for the software components of the product
Check if the product functions as planned
Release the product for a limited set of users
Review the product one last time before the launch
Release the product for all of the users
```

## Data Engineering in the DevOps cycle
```
Design:
- Working with the product engineering teams to plan how data should flow between systems

Development:
- Implement the data pipelines that allows the data flow between systems

Testing:
- Provide mock data to the product to see if it behaves as planned
```

## Data in motion
```
Like a door is essential for a building to allow for people (and pets!) to go through, most software need incoming data to operate on and pass to its users. Which software component below is used to send and receive data between systems?

[ ]Microservices
[ ]Databases
[x]Application Programming Interfaces (APIs)
```

## Making architectural decisions
```
Relevant:
- Size of the team
- The complexity of the software
- The platform that the software is interacting with the users: mobile or web

Irrelevant:
- The sector of the organization
- The geolocation of software users
- The technology used for the data pipelines
```

## IT infrastructure components
```
Hardware:
- Running the online products on servers with significant computing power

Network:
- Making sure the product and infrastructure components are communicating with each other when necessary
- Allow customers to connect to your systems

Software:
- Developing APIs to exchange data
- Employing DevOps best practices
```

## Change management components
```
Developer Platform:
- Developing the tools that the developers need

DevOps:
- Continuously integrating the code changes witgh the version control software
- Arranging the software to start serving the users

Deployment:
- Automating software build, test, and deploy activities
```

## CI/CD pipelines
```
True:
- Data Engineers can use CI/CD pipeline to change the source code for their data tables
- Building the code is a part of CI/CD pipelines, and it means turning the code into machine executable files
- CI/CD pipelines use automation to increase the pace of software development

False:
- Continuous Delivery happens in the Design phase of the DevOps cycle
- Data Engineers can use CI/CD pipelines to configure their data pipelines
- Version control software works after the CI/CD pipelines
```

## DevOps concepts vs. generic infra concepts
```
The Infrastructure Concepts:
- Network
- Security
- Deployment

DevOps Concepts:
- CI/CD Pipelines
- Experimentation
- Minimum Viable Product
```




# 3 Implementation of DevOps for Data Engineering
## Microservices vs. monolithic architecture
```
Microservices:
- Software piece are deployed independently from each other
- Defferent teams specialize in different parts of the software

Monolithic:
- It can be a viable option for small scale backend systems
- There are only a few databases for the whole software
```

## Data Engineering for microservices
```
True:
- There are no central database in a microservices architecture
- Microservices have their own databases

False:
- Microservcies send data to each other via data pipelines
- Monolithic architecture has central databases that is used by the whole system
```

## Main data operations
```
1. Access the data contained within the microservices databases
2. Prepare the data for analytical purposes
3. Send the data to analytical database
```

## Batch vs. stream processing
```
Stream processing:
- For processing user payments
- For signing up users

Batch processing:
- For analytics and reporting
- Setting up a data pipeline that needs to be executed every day at midnight
```

## Change management metrics
```
Requirements:
- Number of new products and features

Design Develop, and Test:
- Time spent in development of the new products
- Pass rate of the DevOps tests

Deploy, Review, Launch:
- Amount of the time the product is unavailable to the users
- User satification rates
```

## Reporting architecture for DevOps
```
1. Identify the data sources within the CI/CD pipelines that have relevant information for your metrics
2. Ingest the data in the CI/CD pipelines into a data pipeline
3. Prepare the data for reporting
4. Move the data into an analytical database
5. Communicate and visualize the results, track the pace of innovation and other metrics
```

## Architecture systems and tools
```
Architecture systems and tools
You are developing a microservice that will provide data to ten other systems. However, the other systems will use your microservices outputs occasionally. Which architecture system and tool combination below could be a good option for your needs?

[ ]Monolithic Architecture | Analytical database
[ ]Microservices Architecture | Airflow
[ ]Monolithic Architecture | Kafka
[x]Microservices Architecture | Kafka
```

## Tools in DevOps cycle
```
Requirements:
- JIRA

Design Development, and Testing:
- GitHub
- Jenkins

Deploy, Review, and Launch:
- Docker
- Kubernetes
```

## Automating data pipelines
```
You want to...

1. establish data pipelines that will record the results of the API
2. automatically transform and record the data every night to another database

Which tool and what pipeline method can you use for these steps?

[ ]1. Jenkins: Batch; 2. Airflow: Batch
[ ]1. Airflow: Streaming; 2. Analytical databases: Batch
[x]1. Kafka: Streaming; 2. Airflow: Batch
[ ]1. Jenkins: Batch; 2. Kafka: Streaming
```




# 4 Accurate, Predictive, and Unbiased Data With DevOps
## Data quality
```
True:
- DevOps helps us have high data quality because it automates and ensures software testing
- Data quality referes to how trusted a set of information is

False:
- All data in an organization must be high quality
- Completeness and accuracy refers to the some elements of data quality
```

## Data quality elements
```
Accuracy or Timeliness:
- How recent and up to date the informaiton is
- Correctness of data in every detail

Consistency or Relevance:
- Holding and storing only the necessary information
- The reliability of the information; data should not contradict with other data

Completeness:
- Ensuring comprehensiveness of data to prove no part of the data was lost
```

## Observability vs. testing
```
Testing:
- Sending mock data to the software and checking how it behaves
- Sending an unexpected amount of data to check how it behaves under pressure

Observability:
- Monitoring how inner components of software behave when it interacts with the users
- CHecking how software interacts with users from different geolocations
```

## Increasing reliability
```
True:
- Observability is important to understand what is going on within software components
- DevOps helps increase reliability by automating testing

False:
- Microservices architecture is not reliable
- Observability is not necessary for increasing reliability
```

## Cultural concepts
```
True:
- When problems arise with the software, first solve it, then arrange a meeting to discuss and learn lessons
- DevOps teams should have representatives with different specialities

False:
- When developing a new product, a team member should wait for testing to finish, to put it into deployment
- DevOps teams should develop and then hand the product to another team for deploy and operate
```

## Reacting to an incident
```
1. Notice the alarm that your product is not working
2. Look at the observability results to identify the root cause of the problem
3. Solve the root cause of the problem and run the product again
4. Conduct a post mortem to see what went wrong and how it could have been prevented
5. Develop more testing so issues like this are more likely to be prevented
```

## Benefits of DevOps
```
[ ]Increasing reliability and quality of softwaere
[ ]AUtomate change management
[ ]Manage the costs when developing software products
[x]All of the above
```

## DevOps cycle
```
1. You collaborate with the business team to discuss requirements on JIRA.
2. You plan the product's appearance after the change and outline the necessary steps to implement it.
3. You develop the new version of the (software) product.
4. You test the new version and deploy it for your users through the CI/CD pipelines on Jenkins.
5. Lanunch the latest version of your product and monitor its performance.
```
