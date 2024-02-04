---
title: Data Science for Business
tags: data-science, analytics
url: https://campus.datacamp.com/courses/data-science-for-business/introduction-to-data-science
---

# 1. Introduction to Data Science
## Customer Segmentation Workflow
```
1. Use SQL to download the delivery date of each box
2. Create a line chart that shows decay in subscriptions by cohort
3. Cluster the users into different personas and perform a regression to predict churn for each cluster

NOTE: This order allows the team to first gather the necessary data, then visualize the subscription decay, and finally, perform customer segmentation and regression analysis to understand and predict churn.
```

## Building a customer service chatbot
```
Data Collection:
- Gather corresponding customer information for each conversation
- Input the timestamps for each transcript

Exploration and Visualization:
- Create a bar chart of the number of conversations of each type
- Plot the number of conversations vs. the time of day

Experimentation and Prediction:
- Use a Markov model to predict possible responses for each question
- Create an algorithm that classifies the initial customer question
```

## Improving OKRs
```
[x]Use a linear regression to estimate a user's Net Promoter Score (NPS) based on their previous purchases, social media interactions, and cohort.
[ ]Collect data on which users are most active on the website.
[ ]Create a dashboard with a histogram showing the number of users with different levels of engagement.

NOTE: This Key Result involves applying predictive modeling techniques (linear regression) to estimate NPS based on various user data, which aligns with the concept of predicting outcomes using data. It also involves experimentation by building a predictive model to assess user behavior.
```

## Assigning data science projects
```
## Traditional Machine Learning
- Cluster customers into different segments
- Dynamic pricing of products

## Internet of Things
- Automate building cooling using temperature sensors
- Detect machinery failure with vibration detectors

## Deep Learning
- Automatically summarize text from meeting notes
- Flag images that contain a safety violation
```

## Building a data science team
```
[ ]1. Traditional machine learning: Traditional machine learning techniques involve training models on structured data to make predictions or classifications. In the case of the startup, while they may use some machine learning algorithms to process the sensor data and make predictions about traffic conditions, the primary focus appears to be on sensor technology and data collection rather than machine learning. Therefore, this category may not be the best fit.
[ ]2. Deep learning: Deep learning involves using neural networks to analyze and extract patterns from data. If the startup is using deep learning models to process the sensor data and provide traffic insights, then this category could be relevant. However, based on the provided information, it seems that the startup's core innovation is the sensor technology itself, with data collection being a key component. Deep learning might come into play for more advanced analysis, but it's not the primary focus.
[x]3. Internet of things (IoT): IoT refers to the network of interconnected physical devices that collect and exchange data. In this case, the startup's product involves installing vibration sensors to collect data from bridges and highways. This aligns closely with the IoT concept, as it involves sensor deployment, data collection, and potentially remote monitoring. Therefore, categorizing the startup under the IoT category would be appropriate.
[ ]4. Natural language processing (NLP): NLP deals with the interaction between computers and human language. Given that the startup's focus is on vibration sensors and traffic measurement, it does not seem relevant to the NLP category, as it doesn't involve language processing or text analysis.
```

## Interpreting a team sprint
```
[x]Hybrid
[ ]Embedded
[ ]Isolated
```

## Editing a job posting
```
[x]Expert user of Excel or Google Sheets, including VLOOKUP and pivot tables
[ ]Some experience with Deep Learning and neural nets
[ ]Basic proficiency in either Java, Scala, or Python for database operations

# Data Analyst Job Posting
Job Id: 8675309
Location: New York
Business: Cool Company
Job Summary:  Join our awesome company and do data analysis for us!
Basic Qualifications:
- Proficient in SQL for data analysis (including JOINs, WINDOW functions)
- Familiarity with Tableau or Power BI for building dashboards
- Expert user of Excel or Google Sheets, including VLOOKUP and pivot tables
```

## Matching skills to jobs
```
Data Engineer:
- Strong Java skills
- Expert at building and maintaining SQL databases

Machine Learning Scientist:
- Experience with Python, especially frediction and modeling
```

## Classifying data tasks
```
Data Engineer:
- Create a new table in SQL database
- Give new team members database access

Data Analyst:
- Create a dashboard for Marketing team
- Update Excel spreadsheet with new graphs

Machine Learning Scientist:
- Train anormaly detector algorithm
- Build image classifier for self-driving car
```




# 2. Data Collection and Storage
## Classifying data for security
```
PII:
- Work address
- Customer name
- Email address

Not PII:
- Customer's language
- Delivery date
- Product ordered
```

## Creating web data events
```
[x] Liking a story: This event is valuable for tracking user engagement with the content on the website. It can help understand which stories or posts are popular and which ones are not resonating with users.
[x] Sharing a story: Tracking when users share a story is important for understanding the virality and reach of the content on the website. It can provide insights into the effectiveness of social sharing features.
[x] Clicking a navigation link: This event is crucial for monitoring user navigation behavior within the website. It helps determine which sections or pages of the website are most frequently visited and can identify potential usability issues if users are frequently clicking on links that don't work.

NOTE: All of the above events should be stored in Jacinda's data warehouse because they provide valuable data points for analyzing user interactions and improving the website's performance. These events can help Jacinda's team make data-driven decisions to enhance the user experience and content engagement on SkimmedIt.
```

## Protecting PII
```
[ ] Everyone should have access to both tables because transparency is a core company value: This approach might be risky from a data privacy perspective, especially if the tables contain Personally Identifiable Information (PII) or sensitive customer data. Access should generally be restricted to authorized personnel to protect the data.
[ ] The entire analytics team should have access to Customers, and everyone at the company should have access to Purchases: This approach may be reasonable if the Customers table contains non-sensitive information that the analytics team needs for their work, while Purchases data is considered less sensitive. However, it's still important to ensure that sensitive PII is appropriately protected within the Customers table, and access should be controlled based on specific job responsibilities and data protection regulations.
[x] Only people with a demonstrated need to contact customers should have access to Customers, and everyone can have access to Purchases: This approach is more in line with data protection best practices. Access to the Customers table is limited to those who need it for their job functions, reducing the risk of unauthorized access to sensitive customer information. Purchases, which may contain less sensitive data, can be made more widely available.
[ ] Only managers should have access to either table: Restricting access to only managers may be too limiting, as it could hinder operational efficiency and prevent authorized personnel from performing their tasks. Access should generally be granted based on job roles and responsibilities, rather than just managerial positions.

NOTE: In practice, a more balanced approach is often recommended, where access is granted on a need-to-know basis, and security and privacy concerns are carefully considered. Data access should be aligned with data protection laws, such as GDPR or CCPA, and organizational policies to ensure the appropriate level of security and compliance. A Data Access Policy should be developed that outlines who can access what data and under what circumstances.
```

## Identifying question purpose
```
Create marketing collateral:
- What makes "Wine Not?" your favorite wine subscription service?
- What is the best bottle of wine you have received from "Wine Not?"

De-risk decision making:
- For a friend's birthday, rank these gifts: gift certificate for dinner at a nice restaurant, gift subscription to "Wine Not?", $50 Amazon gift card.
- On a scale of 1-5, how likely would you be to gift a subscription to "Wine Not?" to a friend colleague?

Monitor quality:
- Please rate your satisfaction with your last "Wine Not?" box on a scale of 1-5, where 1 is extremely dissatisfied and 5 is extremely satisfied.
```

## Validating focus group feedback
```
[ ]"Do you wish Foodwheel had delivery estimates? Select 'yes' or 'no'." - This question directly addresses the focus group's discussion about delivery time estimates and can provide quantitative data on whether this is an important feature for a larger sample of users.
[ ]"What features do you want from Foodwheel? Write in anything you like!" - This open-ended question allows users to provide feedback on a wide range of features they desire. It will help capture additional insights and potentially reveal if high-res pictures of food or corporate accounts are also important to a broader user base.
[x]"Rank the following features from most important to least important: delivery time estimate, high-res pictures of food, corporate accounts." - This question allows users to prioritize these specific features, providing a quantitative measure of their importance relative to each other. It can help validate the focus group's hypotheses about feature importance.
[ ]"Which would you rather have: corporate accounts so that your company can pay for your lunch or delivery time estimates?" - This question directly addresses the trade-off between corporate accounts and delivery time estimates, which was discussed during the focus group. It can help validate whether users prefer one feature over the other.
```

## Net Promoter Score
```
[ ]Quantitative method of measuring revealed preference
[ ]Qualitative method of measuring revealed preference
[x]Quantitative method of measuring stated preference
[ ]Qualitative method of measuring stated preference

NOTE: Net Promoter Score (NPS) is a quantitative method of measuring stated preference, as it involves asking users to provide a numerical rating (on a scale of 0-10) to indicate how likely they are to recommend a product or website.
```

## Sorting data sources
```
APIs:
- Number of Facebook Likes for a new beauty product from users: Explanation: Lucinda can use APIs provided by Facebook or the beauty product's social media platform to access real-time data on the number of likes for the product.
- Track stock prices of different beauty brands over time: Explanation: Lucinda can use financial market APIs (e.g., stock market data providers) to access historical and real-time stock price information for different beauty brands.

Public Records:
- Number of women between the ages of 30 - 45 living in Miami-Dade: Explanation: Public records, such as census data or government demographic information, would be the most appropriate source for population statistics like the number of women in a specific age group living in a particular area.
- Average household income in Miami-Dade County: Explanation: Public records or government data sources would provide information on average household income in a specific geographic area like Miami-Dade County.

Mechanical Turk (MTurk):
- Build a training set where each customer review is labeled as either: Explanation: Using Mechanical Turk, Lucinda can crowdsource the labeling of customer reviews by hiring human annotators to classify them.
```

## Asthma frequency
```
[ ]APIs:  Many health-related organizations and government agencies provide APIs (Application Programming Interfaces) that allow developers to access and retrieve health data, including asthma hospitalization rates by state. Examples include APIs provided by the Centers for Disease Control and Prevention (CDC) or the World Health Organization (WHO).
[x]Public Records: Government health departments and agencies often publish health statistics, including asthma hospitalization rates, as public records. These records can be accessed and used for analysis and visualization.
[ ]SQL: atabase: Some organizations maintain their own databases of health-related data, which can be queried using SQL (Structured Query Language) to extract relevant information. These databases may include hospitalization records and state-specific health data.
[ ]Mechanical Turk: While Mechanical Turk could be used for data collection, it is less likely to be the source for this specific type of data as asthma hospitalization rates are typically collected and reported by official health agencies or organizations with access to health-related datasets.

NOTE: The team member who created the visualization of asthma hospitalization rates for different states for the health insurance company likely used public records or APIs to collect this data.
```

## Cloud platforms
```
[ ]Google Cloud
[ ]Amazon Web Services
[ ]Microsoft Azure
[x]MongoDB

NOTE: MongoDB is NOT an example of a Cloud provider. MongoDB is a NoSQL database management system, while Google Cloud, Amazon Web Services, and Microsoft Azure are all cloud service providers that offer a wide range of cloud computing services, including infrastructure as a service (IaaS), platform as a service (PaaS), and software as a service (SaaS) offerings.
```

## Querying a database
```
[ ]Relational Database, NoSQL
[x]Relational Database, SQL
[ ]Document Database, NoSQL
[ ]Document Database, SQL
```

## Which type of database?
```
Relational Database
- Customer information for all customers, such as name, phone number, user_id, and location.
- The dates, times, subjects, and recipient addresses for all emails sent by the sales team

Document Database
- Text from various emails sendt and received by the sales team
- Images of different products, including metadata about the image's contents
```




# 3. Analysis and Visualization
## Classifying dashboard elements
```
[ ](A) Page Visit Length By Age
[ ](B) Visits Today
[X](C) Monthly Active Users
[ ](D) Web Traffic Source
```

## Improving a dashboard
```
[ ]Nothing, it's fine as it is.
[ ]Change it to a bar chart as in "categorical comparisons" to make it clear how many stocks are from each industry.
[x]Change it to a stacked bar chart as in "tracking composition over time" to make it easier to visually compare different segments.
[ ]Change it to a single number as in "highlighting a single number", to emphasize the total value of the portfolio.
```

## Choosing the right dashboard
```
Line Chart: Number of clients served per month ... A Line chart is ideal for tracking a value over time, such as the number of clients served per month. It allows Kevin to visualize the trend and fluctuations in the number of clients served over different months.
Bar Chart: Number of trades made by each analyst in the past 30 days ... A Bar chart is suitable for categorical comparison, and in this case, it can be used to compare the number of trades made by each analyst in the past 30 days. Each analyst's name can be on the x-axis, and the number of trades they made can be represented by the height of the bars.
Table: A list of the names and ticker symbols for the top 5 stocks over the past 10 days ... A Table is the best choice for displaying information in a structured manner. It allows Kevin to present a list of names and ticker symbols for the top 5 stocks over the past 10 days in a clear and organized format.
```

## Filling out an ad hoc request
```
[ ]A SQL command that she thinks might get the results; it will speed up the process.
[ ]The name of the analyst who she wants to work on the project; that way she knows it will be done correctly.
[x]The reason why she needs this data; the analysts might suggest other data that will help her with her problem.
```

## Classifying requests
```
Dashboard request
- CEO wants daily updates on how many new subscriptions we received.
- Growth team wants to track weekly progress on their new sign-up workflow.

Ad hoc request:
- Legal department wants to know if any of our subscribers are from a sactioned country. They need to remove those users ASAP!
- Marketing wants to know what percent of subscribers have a `.edu` email address. They're investigating a student discount.
```

## Creating an A/B testing workflow
```
1. Picking a metric to track: Determine the specific metric(s) you want to measure to evaluate the performance of your ad copies. This should be the first step to define your experiment's goals.
2. Calculating sample size: After selecting the metric, calculate the sample size required for your A/B test to ensure statistical significance. This step is crucial to ensure that your results are reliable.
3. Running the experiment: Once you have determined the sample size and have set up the necessary infrastructure for the test, you can begin running the A/B test by randomly assigning your audience to two groups: one exposed to ad copy A and the other exposed to ad copy B.
4. Checking for significance: After running the experiment for an appropriate duration, collect and analyze the data to determine if the differences in performance between the two ad copies are statistically significant. This step is essential to draw valid conclusions from your A/B test results.
```

## Sample size
```
[ ]increases/increases
[ ]decreases/increases
[ ]increases/decreases
[x]decreases/decreases

NOTE: As baseline conversion increases, sample size decreases. As minimum detectable difference increases, sample size decreases.
```

## Intermediate results
```
[ ]It looks like the bar representing the new workflow has a slightly lower cancellation rate than the bar representing the old workflow, so Alondra should stop the test and switch to the new workflow.
[x]The number of users exposed to the test has not yet reached the desired sample size, so Alondra should wait and make her judgment after the desired sample size has been reached.
[ ]The results are not yet significant, so Alondra should cancel the test and keep using the old workflow.

NOTE: Based on the information provided, Alondra should conclude that the results are not yet significant. Therefore, she should not make any final judgments or decisions based on the current state of the experiment. She should wait until the desired sample size has been reached and conduct a proper statistical analysis to determine the significance of the results before making any decisions about whether to continue with the new workflow or stick with the old one.
```




# 4.
## When to use Supervised Learning
```
A/B Test:
- Growth team wants to know which barner ad gets the most clicks.
- Sales team wants to know which script maximized purchases from calls with new clients.

Supervised Machine Learning:
- Customer Success team wants to know which customers are likely to churn based on previous purchases and social media interactions.
- Finance team wants to know which customer accounts are likely to miss a pyment based on previous payments and demographics.

Dashboard:
- CEO wants a chart correlating customer age with lifetime value.
```

## Features and labels
```
[ ]50XP
[ ]Age
[ ]Income
[x]Loan Type

NOTE: In the context of building a recommendation system to predict the type of loan that a customer is most likely to use, the "Loan Type" column is the label. The label is the target variable that the machine learning model will predict. In this case, Juan's team wants to predict the loan type, so "Loan Type" is the label, and the other columns like "Age" and "Income" are likely to be features or attributes used to make predictions about the label.
```

## Model evaluation
```
[x]Decision tree - it has high rates of accuracy for both valid transactions and fraudulent transactions.
[ ]Logistic classifier - overall, it is accurate in 97% of cases.
[ ]Random forest - it has equal accuracy for both fraudulent and valid transactions.

NOTE: Decision Tree:
- Accuracy for Valid Transactions: High
- Accuracy for Fraudulent Transactions: High
- Precision for Fraudulent Transactions: High
- Recall for Fraudulent Transactions: High
The decision tree model appears to perform well on both valid and fraudulent transactions based on the provided data. It has high accuracy for both classes, which suggests that it correctly classifies a significant portion of transactions.
```

## Supervised vs. unsupervised
```
Supervised:
- Predict whether a new clothing style will be successful based on previous season's trends - This involves predicting a binary outcome (successful or not) based on historical data, making it a classic supervised learning problem.
- Predict if a customer will purchase an item based on what previous customers purchased - Again, this problem involves predicting a binary outcome (purchase or not) based on historical customer behavior, making it a supervised learning task.

Unsupervised:
- Divide customers into different categories based on their shopping habits - This task involves clustering customers into different categories based on their shopping habits. Since there are no predefined labels, it's an unsupervised learning problem, specifically clustering.
- Create groupings of clothing items that have similar features - This problem involves finding patterns and similarities among clothing items without predefined labels, which makes it an unsupervised learning task, particularly clustering or dimensionality reduction.

NOTE: In summary, both tasks involving prediction with historical data are Supervised, while tasks involving grouping or clustering without predefined labels are Unsupervised.
```

## Cluster size selection
```
[ ]2
[ ]3
[x]4
[ ]5
[ ]6
```

## Special topics in Machine Learning
```

```

## Classifying machine learning tasks
```

```

## Sentiment Analysis
```

```

## Deep Learning and Explainable AI
```

```

## Finding the correct solution
```

```

## Should I use Deep Learning?
```

```
