---
title: Introduction to Databricks
tags: databricks, data-science, data-engineering, analytics
url: https://app.datacamp.com/learn/courses/introduction-to-databricks
---

# 1. Introduction to Databricks
## Why pick a Lakehouse?
```
Which of the following is not a reason to implement a lakehouse architecture?

[ ]You want great performance for both BI and machine learning workloads.
[ ]You want a single architecture to support all of your data workloads.
[x]You want to store all of your data in a proprietory data format.
[ ]You want the most cost-effective architectural solution for your organization.
```

## Benefits of the Databricks Lakehouse
As part of your presentation to the CIO, you would like to list why you should want to switch to the lakehouse architecture. Which of the following are key benefits of the Databricks platform? Select all that apply.
```
[x]Built on open-source technologies
[x]Databricks can run on any major cloud provider.
[x]Databricks enables data teams to be more collaborative.
[ ]Only teams that heavily use Python will benefit from Databricks
[ ]Database administrators can separate their data warehousing workloads away from other languages and personas.
```

## Architectural Decisions
```
## Data Warehouse
Typically very expensive
Performant for structured data only

## Data Lake
Data becomes very messy
Sub-optimal performance for different workloads

## Lakehouse
Great performance across all workloads
Very cost-effective for any data workload
```

## Why Delta?
```
The benefits of storing data in the Delta Lake table format include:

1. Fully ACID-compliant transactions: Delta Lake provides support for ACID (Atomicity, Consistency, Isolation, Durability) transactions, which ensures data integrity and consistency even in the face of concurrent data updates and failures.
2. Unified batch and streaming data sources: Delta Lake is designed to handle both batch and streaming data sources, making it suitable for real-time data processing scenarios.
3. Time travel and table history: Delta Lake allows you to access and query previous versions of the data (table history) through time travel capabilities, which can be useful for auditing, data recovery, and analysis purposes.

So, the correct answers are:

[x]Fully ACID-compliant transactions
[x]Unified batch and streaming data sources
[ ]Rigid table schema limitations
[x]Time travel and table history
```

## Databricks for different personas
```
One of the concerns of your CIO is that your future data architecture must satisfy all the different teams who work with your company's data. While you feel confident that Databricks can, you want to provide a list that clearly outlines supported personas.
Which of the following is not a dedicated persona in the Databricks Lakehouse platform?

[ ]Data engineer
[x]Software developer
[ ]Data scientist
[ ]SQL analyst
```

## Capabilities for each data persona
```
## SQL analyst
Dashboards
SQL Editor
SQL Warehouses

## Data Engineer
Delta Live Tables

## Data Scientist
Model Registry
Model Serving
```

## Managing and adding users
```
As your organization starts to adopt Databricks more, you have deployed multiple different workspaces for your lines of business. Each workspace is designed for a particular team and should only have access to their respective datasets. Your CIO has expressed concerns about how you can control access to the underlying datasets and the workspaces themselves as adoptions grow. They have asked you how Databricks and your architecture address this concern.
Which of the following statements is true with regard to user access within a Databricks account?

[ ]There is no way to easily ensure that users are assigned to the correct workspace, and must assign users manually.
[x]You can automatically sync users from your IdP to your Databricks account and then assign them to the correct workspaces based on their line of business.
[ ]You can automatically sync users from your IdP to your Databricks workspaces with no configuration
[ ]You do not need to assign users to Databricks, and can just pull their access information from your IdP.

Note: This means that while Databricks supports automatic synchronization of users from your Identity Provider, there is still a requirement for configuration or manual intervention to ensure that users are assigned to the correct workspaces according to their specific roles or lines of business. This approach helps maintain secure and organized access control as your organization scales up its use of Databricks.
```

## Control Plane vs. Data Plane
```
As you deploy your Databricks workspaces, you ensure your security team has approved the platform. In particular, the security team is curious about what information Databricks has access to and what you control.
Which of the following is true regarding the relationship between the Databricks Control Plane and the customer Data Plane?

[ ]The Control Plane stores customer data and the Databricks application, and the Data Plane only stores data.
[x]User interactions with the Databricks platform exist within the Control Plane, and is responsible for tasks such as launching clusters and initiating jobs.
[ ]Databricks creates a duplicate copy of everything between the Control and Data Plane.
[x]The Data Plane is where customers store their data in the data lake, ensuring better data security.
```

## Configure your Databricks workspace
```
Databricks:
Set up Unity Catalog to govern data.
Add users and groups to the Databricks environment.
Create a starter cluster and SQL Warehouse for your users.

Cloud Service Provider (CSP):
Collect and store data from source systems.
Set up user security groups across cloud services.

NOTE: Databricks focuses mainly on setting up data management and analytical environments, while CSP involves the collection, storage, and security management of data.
```



# 2. Data Engineering
## Data Explorer capabilities
```
You are a data engineer working for the centralized IT department at Sierra Publishing. On a day-to-day basis, you will get a variety of requests from different parts of your organization regarding different datasets. Historically, you have explored data programmatically by reading in different datasets one at a time.
Which of the following will be the benefits of using the Data Explorer instead of your previous approach?

[x]Ability to see sample data for a particular table.
[x]See the lineage for your data assets.
[ ]Provides the code to query a specific table.
[x]Ability to share data with other users.
[ ]Ability to automatically clean your data.
```

## Setting up a notebook
```
[ ]Create a SQL notebook and require all code to be in SQL.
[x]Create a Python notebook for your data engineering pipelines, and use %sql in the cells for your data analysts.
[ ]Create a Python notebook and require all code to be in Python.
[ ]There is no way for you to create a setup that supports both languages.

NOTE: This approach allows you to write in Python, which is your primary language, and enables your data analysts to write SQL queries within the same notebook by using the `%sql` magic command. Databricks notebooks support multi-language capabilities, making it possible to use different programming languages within the same notebook. This setup fosters collaboration and ensures that both you and your data analyst team can work efficiently in your preferred languages within the same data asset.
```

## Cluster configurations
```
## Valid Cluster Configuration
Databricks Runtime: Valid. The Databricks Runtime is a crucial part of the cluster configuration, encompassing the Apache Spark cluster manager and optimized versions of Hadoop components.
Node instance types: Valid. You can choose the types of nodes (instances) to use in your cluster. Different instance types offer varying combinations of CPU, memory, storage, and networking capacity, suitable for different workloads.
Auto termination time: Valid. This setting allows clusters to automatically terminate after a specified period of inactivity, which is an important feature for cost savings.
User access to cluster: Partially Valid. While there isn't a direct setting to control user access to a cluster, managing user permissions and access controls at the workspace level indirectly relates to cluster usage.

## Invalid Cluster Configuration
How fast workloads run on the cluster: Invalid. There's no direct option to set the "speed" of workload execution. However, the performance of workloads is indirectly influenced by various settings like node types and cluster size.
Cluster monthly budget: Invalid. Databricks doesn’t have a direct feature to set a "monthly budget" for clusters. However, cost management tools can be used to track and limit expenses, though this falls under organizational management rather than being a part of cluster configuration.
```

## Data Engineering foundations in Databricks
```
[x]DataFrames already exist in both pandas and SQL: This is correct. DataFrames are a common data structure used in both pandas (a popular Python library for data manipulation) and SQL. Your team's familiarity with these tools will be an advantage when working with DataFrames in Databricks.
[ ]DataFrames are complex and difficult to work with: DataFrames are a powerful and flexible way to work with data, but they are not inherently complex or difficult to work with. In fact, they provide a structured and efficient way to perform data operations, and their syntax is often more intuitive than writing raw SQL queries or using lower-level Spark RDDs (Resilient Distributed Datasets). Your team should not be discouraged by the perception that DataFrames are overly complex.
[x]Databricks will create the DataFrame, regardless of the language you use in your programming: This is also correct. Databricks supports multiple programming languages like Python, Scala, and R. DataFrames can be created and manipulated using these languages, so your team can continue using their preferred language while working with DataFrames.
[ ]DataFrames only work when writing code in Python
```

## DataFrames
```
[x] DataFrames already exist in both pandas and SQL: This is correct. DataFrames are a tabular data structure that exists in both pandas (a popular Python library for data manipulation) and SQL (Structured Query Language), so your team is likely already familiar with the concept of DataFrames.
[ ] DataFrames are complex and difficult to work with: This is not true. While DataFrames can handle complex data operations, they are designed to provide a convenient and intuitive way to work with structured data. They are not inherently difficult to work with, and with some training and practice, your team should be able to use them effectively.
[x] Databricks will create the DataFrame, regardless of the language you use in your programming: This is also correct. Databricks supports multiple programming languages, including Python, Scala, and R, for working with DataFrames. Regardless of the language your team chooses to use, Databricks will allow them to create and manipulate DataFrames.
[ ] DataFrames only work when writing code in Python: This is not true either. DataFrames can be used in multiple programming languages within the Databricks environment, as mentioned earlier. They are not limited to Python and can be used in Scala, R, and more.
```

## Reading from a database
```
Use spark.read() to use a JDBC connection and read data from the database.
```

## Write an external table
```
CREATE TABLE postgresql_table
USING postgresql
OPTIONS (
  dbtable '<table-name>',
  host '<database-host-url>',
  port '5432',
  database '<database-name>',
  user '<username>',
  password '<password>'
)
LOCATION <path-to-data-lake>
```

## Loading in hosted files
```
[x]Use Auto Loader to automatically ingest new data files as they come in: Auto Loader is a feature in Databricks that can automatically detect and ingest new data files as they are added to a specified directory. This would address the issue of not knowing when new data files come in and simplify the process of keeping your tables accurate by automatically incorporating new data.
[ ]Scheduling a job to check the directory every hour: this option is a manual approach and may not be as efficient or timely as using Auto Loader.
[ ]Creating a Kafka topic and continuously streaming the data: this option is a different approach that involves streaming data continuously and might be suitable for certain use cases, but it may not be necessary or the most straightforward solution for this specific problem.
```

## Selecting the right language
```
## Databricks Supported
Language: SQL - Joining two tables together
Language: R - Performing advanced statistical techniques on your data
Language: Python - Creating a custom function for a business-specific calculation

## Not Supported
Language: Rust - Creating a definition for a data schema
Language: C# - Crafting an In-depth testing process
Language: Ruby - Creating a complex function that Iterates over each row
```

## Data pipeline steps
```
1. Gather information about the Kafka topic: Before you can start processing the data, you need to gather information about the Kafka topic, such as the topic name, partitions, and offsets. This information will be essential for setting up the data stream.
2. Read the data into a Structured DataFrame: Once you have the necessary information about the Kafka topic, you can use Databricks to read the streaming data from Kafka into a Structured DataFrame. This step involves configuring the Kafka source and defining the schema for your DataFrame.
3. Clean the streaming data: After reading the data into a DataFrame, you may need to perform data cleaning and transformation operations. This could include handling missing values, data type conversions, or any other necessary data preparation tasks.
4. Join our other datasets into the stream: Next, you can join the streaming data with other datasets as required for your analytics. This step allows you to enrich the streaming data with additional context or information from your static datasets.
5. Write the stream as a Delta table in a new data lake bucket: Once you've processed and enriched the streaming data, you can write the result as a Delta table to a new data lake bucket. Delta tables in Databricks provide a reliable and efficient way to manage and version your data, making it available for downstream consumers.
```

## Possible automations in Databricks
```
## Can be a Databricks Job
- Python scripts: Databricks supports Python, and Python scripts can be executed as part of a Databricks Job.
- Databricks Notebooks: These are native to the Databricks environment and can be scheduled and automated as Databricks Jobs.
- SQL queries: Databricks supports the execution of SQL queries, which can be part of a Databricks Job.
- JAR files: Databricks can run JAR files as Jobs, especially those written for Apache Spark.
- dbt Tasks: Databricks can integrate with dbt (data build tool) for running transformation tasks as part of a Databricks Job.

## Cannot be a Databricks Job
- Cloud automations: This typically refers to broader cloud management tasks that are outside the scope of what Databricks Jobs are designed for.
- pgSQL scripts: While Databricks supports PostgreSQL syntax to some extent, running specific PostgreSQL (pgSQL) scripts directly as a Job might not be feasible without modifications or additional integrations.
- Excel macros: Databricks cannot execute Excel macros as they are specific to Microsoft Excel and require the Excel environment.
```

## Benefits of Delta Live Tables
```
[x]Combined process for batch and streaming datasets
[ ]Reduced need for data pipelines
[ ]Data is never lost or corrupted
[x]Maintain better data quality with set expectations
```

## Data pipeline steps
```
1. Ingest raw data files into a Delta table with Autoloader.
2. Create a Delta Live Table pipeline to clean and join datasets.
3. Create a Delta Live Table pipeline to aggregate datasets for BI applications.
4. Orchestrate all the Data Engineering tasks with a Databricks Workflow.
5. Monitor the data quality of the pipelines to keep track of the accuracy of your tables.
```




# 3. Databricks SQL and Data Warehousing
## Benefits of Databricks SQL
```
[x] Access to more advanced capabilities as part of a broader platform: Databricks provides a unified platform that integrates various advanced capabilities like machine learning, data engineering, and data science, in addition to its SQL capabilities. This integration can be highly beneficial for complex data analytics workflows.
[x] Flexible, open access to data and the SQL language: Databricks SQL offers flexible and open access to data, allowing users to query data across different data sources using the familiar SQL language. This flexibility can be crucial in a data-driven organization where data resides in various formats and locations.
[ ] Leverage a proprietary system that restricts access to different tools: This is not a benefit. In fact, this would be a limitation, as proprietary systems often restrict flexibility and integration with other tools and platforms.
[x] Use of a scalable, optimized compute engine for SQL workloads: Databricks SQL is built on a highly optimized and scalable compute engine designed to handle SQL workloads efficiently. This means that you can expect improved performance, especially for large-scale data processing tasks, which is a critical factor for analytics-driven organizations like yours.
```

## Databricks SQL in the data workflow
```
1. Ingest Raw Data into the Delta Lake: The first step is to ingest raw data into Delta Lake. Delta Lake provides a storage layer that brings reliability to Data Lakes. It allows you to store both structured and unstructured data and is an essential part of the Databricks platform.
2. Clean / Transform the Data with Delta Live Tables: Once the data is in Delta Lake, the next step is to clean and transform it. This step is crucial for ensuring data quality and usability. Delta Live Tables provide an easy and reliable way to transform and clean your data.
3. Aggregate the Data Based on Business Needs: After the data is cleaned and transformed, the next step is to aggregate it according to your business requirements. This involves summarizing your data to make it more suitable for analysis, such as calculating totals, averages, or other statistical measures.
4. Query the Data Using Databricks SQL: With your data aggregated, you can now query it using Databricks SQL. This SQL analytics service on Databricks allows you to run quick and efficient queries on your data, enabling you to extract meaningful insights.
5. Visualize Your Queries Using Your BI Tool: Finally, the insights gained from your queries can be visualized using a Business Intelligence (BI) tool. This step is about representing your data in a more accessible and understandable format, which is crucial for making informed business decisions.
```

## Databricks SQL vs. other databases
```
### Databricks SQL
- Operates on Open Data Formats: Databricks SQL is designed to work seamlessly with open data formats like Parquet, Delta Lake, etc. This is a distinguishing feature of Databricks SQL, as it allows for more flexibility and interoperability with different data sources and tools.
- Leverage ML Models: Databricks SQL integrates closely with machine learning capabilities, leveraging Apache Spark’s MLlib or Databricks’ own machine learning features. This enables users to easily incorporate machine learning models into their SQL queries and analytics workflows.
- Built on ANSI SQL and Spark: Databricks SQL is built upon Apache Spark and supports ANSI SQL, which is a standard SQL dialect. This ensures compatibility with traditional SQL queries while also benefiting from Spark's distributed processing capabilities.

### Other Data Warehousing Tools
- Tech-Specific SQL: Traditional data warehousing tools often use proprietary or tech-specific SQL dialects. This might limit interoperability and could require more specialized knowledge compared to the more universally known ANSI SQL.
- SQL Workloads Only: Many conventional data warehousing solutions are primarily designed for SQL workloads and might not natively support machine learning or operations on open data formats. They focus more on structured data processing and analysis.
```

## Choosing your SQL warehouse
```
[ ]Pro: Typically offers more advanced features and customization options. However, it might not be as quick to spin up compared to a serverless option, which can be a drawback for users who need immediate access to data.
[x]Serverless: This type of warehouse is designed for ease of use and quick accessibility. It generally requires less management and can automatically scale to meet demand, meaning users won't have to wait for the warehouse to spin up. This aligns well with the users' expectation of quickly accessing data. Additionally, Databricks has been known to integrate the latest features into their serverless offerings.
[ ]Classic: Usually the most basic option with standard features. It might not include the latest updates from Databricks as promptly as the serverless option and could have longer spin-up times compared to serverless warehouses.
```

## SQL Editor vs. notebooks
```
[X]A dedicated data explorer pane to see catalogs, schema, and tables - This feature can be a significant benefit of the SQL Editor. A dedicated data explorer pane provides an easy and organized way to view and manage databases, schemas, and tables, which is essential for efficient SQL query writing and data exploration.
[X]Dedicated area for writing SQL queries - The SQL Editor typically offers a streamlined and specialized interface for writing SQL queries. This focused environment can enhance productivity and user experience, especially for those primarily working with SQL.
[ ]SQL queries always run faster in the SQL Editor - This statement is misleading. The speed of SQL query execution depends on various factors such as the query's complexity, the amount of data being processed, and the underlying database's performance. The interface used for writing the query (SQL Editor vs. Notebook) does not inherently affect the execution speed.
[ ]Databricks will write SQL queries for you in the SQL Editor - This is not a feature of the SQL Editor. While some tools may offer query suggestions or autocomplete features, the actual writing of SQL queries is still a manual process that requires user input and decision-making.
```

## Creating the usSales table
```
CREATE TABLE usSales
USING delta AS (
    SELECT *
    FROM globalSales
    WHERE country = 'United States
)
```

## Understanding Databricks SQL assets
```
[ ]Dashboards can have multiple visualizations attached to them, and queries can show results from multiple dashboards or visualizations.
[ ]Visualizations can have multiple queries attached to them, and dashboards can show results from multiple queries or visualizations.
[x]Queries can have multiple visualizations attached to them, and dashboards can show results from multiple queries or visualizations.
[ ]There is no way to relate queries, visualizations, or dashboards together in Databricks SQL.

NOTE:
In the Databricks SQL environment, the typical workflow involves creating queries to extract and process data. These queries can then be used to generate visualizations, such as charts or graphs, to represent the data in a more digestible and visually appealing format. Multiple visualizations can be created from a single query, depending on the different aspects of the data you want to highlight.
Furthermore, these visualizations can be compiled together into dashboards. A dashboard is a collection of visualizations that provides a comprehensive view of the data insights. Each dashboard can contain results from multiple queries, allowing for a broad overview of different data points and trends. This setup is useful for monitoring, reporting, and data analysis purposes, where you might need to observe various metrics and KPIs simultaneously.
```

## Using parameters in queries
```
[ ]SELECT * FROM sales_distribution WHERE distributor = companyName AND num_transactions >= transactionCount
[x]SELECT * FROM sales_distribution WHERE distributor = {{companyName}} AND num_transactions >= {{transactionCount}}
[ ]SELECT * FROM sales_distribution WHERE {{company_name}} AND {{transactionCount}}
[ ]SELECT company_name, transactionCount FROM sales_distribution

NOTE: This query uses the parameters {{companyName}} and {{transactionCount}} in the WHERE clause to filter the results based on the selected distributor and the minimum number of transactions specified. This allows the user to interactively change the criteria for analysis through the parameter widgets in the dashboard.
```

## Create a user review query
```
USE book_catalog.distribution;
CREATE VIEW user_demographics AS (
    SELECT Location,
        avg(Age) AS averageAge,
        count(User-ID) AS totalReviews,
        round(avg(Book-Rating), 1) AS averageRating
    FROM user_reviews
    WHERE Location IN ('california', 'new york', 'quebec', 'alberta', 'ontario')
    GROUP BY Location
);
```




# 4. Databricks for Large-scale Applications and Machine Learning
## Lakehouse benefits to ML
```
[x]Access to a scalable and performant compute engine: This is a core benefit of Databricks. The platform provides a scalable and high-performance computing environment, which is essential for handling large datasets and complex machine learning models.
[x]Support for all the open source libraries and frameworks your team needs: Databricks is known for its extensive support of open-source libraries and frameworks. It integrates well with popular machine learning and data processing frameworks, making it a flexible choice for teams used to working with a variety of tools.
[ ]Create and tune models for you with a drag-and-drop UI: While Databricks provides an interactive workspace and supports various tools for machine learning, it doesn't create and tune models automatically through a drag-and-drop UI. It's more about providing the environment and tools for data scientists to build, tune, and deploy models rather than doing it automatically.
[x]Allow your team to collaborate with data engineers and data analysts: This is another significant benefit of Databricks. Its collaborative environment is designed to bridge the gap between data engineers, data scientists, and analysts, facilitating seamless collaboration and integration of workflows.
[ ]Automatically deploy your model anywhere you need it to go: Databricks streamlines the deployment process, but it doesn't fully automate deployment to any destination. It provides tools and integrations to facilitate deployment, but this typically requires some configuration and management from the team.
```

## MLOps tasks in Databricks
```
## DataOps
Removing NULL values from your silver table.
Adding a column as feature in a data table.
Ingesting row data int oDelta tables.

## ModelOps
Generate on example model using AutoML.
Training a SparkML model.

## DevOps
Governing access to models with Unity Catalog.
Deploying a model on a Serving Endpoint.
```

## EDA in Databricks
```
[ ]Yes, they can have Databricks write code for them to perform EDA.
[x]Yes, they can use the built-in visualization tools in Databricks Notebooks for EDA.
[ ]No, there is no way for them to perform EDA without knowing pandas.

NOTE: Databricks provides various visualization tools and capabilities that can help junior data analysts explore and analyze data without requiring extensive knowledge of pandas or coding. These tools allow them to create visualizations, generate summary statistics, and gain insights from the data, making it easier for them to contribute to machine learning workloads.
```

## Why the ML Runtime?
```
[x]ML Runtime comes with the leading libraries pre-installed: ML Runtime comes with the leading libraries pre-installed: The Databricks ML Runtime is pre-configured with popular machine learning libraries like scikit-learn, TensorFlow, PyTorch, and more. This eliminates the need for manually managing and installing these libraries, saving time and effort for data scientists.
[ ]ML Runtime comes with all the libraries you'll ever need from the online respositories.
[x]ML Runtime is optimized for machine learning workloads: ML Runtime is optimized for machine learning workloads: ML Runtime is specifically designed and optimized for machine learning workloads. It includes performance improvements, hardware acceleration support, and other optimizations that can significantly enhance the efficiency and speed of machine learning tasks when compared to manual library management on local laptops or non-specialized environments.
[ ]ML Runtime is no different than installing your libraries onto each cluster.
```

## Exploring data in a notebook
```
df = spark.read.table('review_data')
dbutils.data.summarize(df)
feature_df = (
    df.withColumn('feature1', ...)
      .withColumn('feature2', ...)
)
from databricks import feature_store
fs = feature_store.FeatureStoreClient()
fs.create_table(
    name=review_features,
    df=feature_df,
    description="review features"
)
```

## Single node vs. multi node ML
```
## Single Node
Quick, easy, local setup
scikit-learn
Great for play around with

## Multi node
Highly scalable
Sparl ML
Great for production
```

## Databricks for citizen data scientists
```
[ ]The junior data scientists should use an external "black box" auto ML tool and give those models to you.
[ ]The junior data scientists should force themselves to use specific frameworks and learn independently.
[x]The junior data scientists can create models with AutoML, and you can review or tune the code from the generated notebooks.
[ ]The junior data scientists will not be able to contribute until they learn how to program an end-to-end pipeline.

NOTE: This approach allows the junior data scientists to contribute effectively using AutoML tools, which are more accessible for those not yet comfortable with building machine-learning pipelines from scratch. At the same time, it provides an opportunity for you to review, understand, and improve upon the work they start, ensuring quality and consistency in your production models.
```

## Using MLFlow for Tracking
### 1. Define an Experiment
An experiment in MLflow is a way to organize and keep track of your machine learning runs. Each experiment can contain multiple runs, and each run represents a single trial of training a model. When you define an experiment, you essentially create a workspace where all related runs are stored and tracked.

### 2. Start the Machine Learning Run
To start tracking a machine learning run in MLflow, you use the `mlflow.start_run()` context manager. This initiates a new run in the current active experiment. If no experiment is active, it creates a run in a default experiment.

```python
import mlflow

with mlflow.start_run() as run:
# Your training code goes here
```

### 3. Train Your Model and Track Your Information
While the run is active, you can log parameters, metrics, models, and artifacts. There are two primary ways to log this information:
- **`mlflow.autolog()`**: This automatically logs metrics, parameters, and models without needing explicit log statements. It's useful for standard training scenarios.
```python
mlflow.autolog()
# Training code here
```
- **`mlflow.log_param()` and `mlflow.log_metric()`**: These are used for manual logging. `log_param` records a single key-value pair of a parameter, and `log_metric` records a key-value pair of a metric.
```python
mlflow.log_param("param_name", param_value)
mlflow.log_metric("metric_name", metric_value)
```

### 4. Compare Your Model Runs in the MLFlow Experiment
After conducting several runs, you can use the MLflow UI to compare these runs. The MLflow UI provides a visual interface where you can view, compare, and analyze different runs. You can sort and filter runs based on metrics and parameters, which is very useful for identifying the best performing models.

### 5. Select the Model from the Best Run
Once you've compared the runs, you can select the best model based on your criteria (e.g., highest accuracy, lowest error rate). MLflow allows you to register this model in its model registry, where you can version, manage, and deploy models.
```python
best_run_id = "your_best_run_id_here"
model_uri = f"runs:/{best_run_id}/model"
mlflow.register_model(model_uri, "model_name")
```
This process simplifies the experimentation and model selection phase of machine learning projects, making it easier to track, compare, and manage different models and their respective runs.

## Models and the Model Registry
```
[x]1. Easily see current and past versions of the same model. - This is true. The Model Registry allows you to track and view different versions of a model, making it easy to compare and manage them over time.
[x]2. Ability to push a model into either Staging or Production. - This is true. The Model Registry typically provides a mechanism to deploy models to different environments, such as staging or production, making it easier to manage the deployment process.
[ ]3. Ability to automatically select the best version of a model for Production. - This can be true in some cases, depending on the specific implementation of the Model Registry. Some Model Registries offer features for automated model selection based on criteria like performance metrics.
[ ]4. Automatically push the latest version of a model into Production. - This may or may not be a feature of the Model Registry. It depends on the configuration and automation set up for deployment. Not all Model Registries have this capability by default.
[x]5. Ability for others to discover their version of the model. - This is true. The Model Registry allows team members to discover and access different versions of a model, facilitating collaboration and knowledge sharing among data science teams.
```

## Why Databricks for model deployment?
```
[ ]Model Serving Endpoints can be deployed in both on-prem and cloud environments.
[ ]Model Serving Endpoints are deployed in our environment, and we can control them however we want.
[ ]Model Serving Endpoints do not have any appreciable benefit over hosting models with on-prem compute.
[x]Model Serving Endpoints simplify the overhead of managing compute resources and provide built-in capabilities for model monitoring.

NOTE: Using Databricks Model Serving Endpoints can simplify the deployment process for machine learning models by abstracting away much of the complexity of managing compute resources in the cloud. Additionally, Databricks provides built-in capabilities for model monitoring, which can be valuable for tracking model performance and ensuring that it continues to perform well in production. This can save time and effort compared to traditional deployment methods with on-prem compute resources, which often require more manual configuration and monitoring setup.
```

## End-to-end ML pipeline
```
1. Transform datasets, create features, and store them in the Feature Store. - This step involves data preprocessing and feature engineering, which should be done before model training.
2. Train your regression model using the SparkML library. - Once you have prepared the datasets and features, you can proceed with training your regression model.
3. Test out your model against another dataset and track the results with MLFlow. - After training, it's important to evaluate your model's performance on a separate test dataset and track the results using MLFlow for monitoring and experimentation.
4. Store your trained model in the Model Registry. - Once your model is trained and evaluated, you can store it in a model registry for versioning and easy access.
5. Serve your model with a Databricks serving endpoint for your end users to consume. - Finally, you can deploy your trained model to a serving endpoint, making it accessible for end users to consume and make predictions.
```

## Wrap Up
```

```
