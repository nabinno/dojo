---
title: "Using Tableau Desktop with Amazon Redshift"
tags: amazon-web-services, amazon-redshift, tableau
url: https://www.qwiklabs.com/focuses/401
---

# Goal
- Load data into Amazon Redshift from Amazon S3
- Connect to Amazon Redshift cluster from Tableau
- Create a dashboard using Tableau

# Task
- [x] Connect to Your Amazon EC2 instance
- [x] Load data in Amazon Redshift from Amazon S3
- [x] Install and Activate Tableau trial for use with the lab
- [x] Connect to Amazon Redshift from Tableau
- [x] Create a Dashboard in Tableau

# Supplement
## Load data in Amazon Redshift from Amazon S3
**run-sql.ps**
```powershell
Set-AWSCredentials -AccessKey "xxxxxxxxxxxxxxxxxxxx" -SecretKey "xxxxxxxxxxxxxxxxxxxx" -StoreAs default
Initialize-AWSDefaults -ProfileName default -Region us-west-2
$rsServer=(Get-RSClusters).Endpoint.Address
$rsPassword="x4SPp9FKjGhS"
$accessKey="xxxxxxxxxxxxxxxxxxxx"
$secretKey="xxxxxxxxxxxxxxxxxxxx"
$rsDatabase="faa"
$rsUser="administrator"
$rsPort="5439"

try
{
  $rsConnectionString="Driver={Amazon Redshift (x64)}; Server=$rsServer; Database=$rsDatabase; UID=$rsUser; PWD=$rsPassword; Port=$rsPort"
  $rsConnection=New-Object System.Data.Odbc.OdbcConnection
  $rsConnection.ConnectionString=$rsConnectionString

  $rsCommand=New-Object System.Data.Odbc.OdbcCommand
  $rsConnection.Open()
  $rsCommand.Connection=$rsConnection

  write-Host -NoNewline "Creating tables........"

  $rsLoadStatement = Get-Content C:\Users\Administrator\Downloads\create-table.sql -Raw
  $rsLoadStatement = $rsLoadStatement -replace "access-key", $accessKey
  $rsLoadStatement = $rsLoadStatement -replace "secret-key", $secretKey
  $rsCommand.CommandText=$rsLoadStatement
  $rsResult=$rsCommand.ExecuteNonQuery()

  write-Host "Success"
  write-Host ""

  write-Host -NoNewline "Loading data.........."

  $rsLoadStatement = Get-Content C:\Users\Administrator\Downloads\load-data.sql -Raw
  $rsLoadStatement = $rsLoadStatement -replace "access-key", $accessKey
  $rsLoadStatement = $rsLoadStatement -replace "secret-key", $secretKey
  $rsCommand.CommandText=$rsLoadStatement
  $rsResult=$rsCommand.ExecuteNonQuery()

  write-Host "Success"
  write-Host ""

}
catch
{
  write-Host "An error occurred while attempting to initialize Amazon Redshift. Please end the lab and email to aws-course-feedback@amazon.com"
}
finally
{
  if ($rsConnection.State -eq "Open")
  {
    $rsConnection.Close()
  }
}

Write-Host "Press any key to continue...."

$x = $host.UI.RawUI.ReadKey("NoEcho,IncludeKeyDown")
```

**create-table.sql**
```sql
CREATE TABLE public.dim_carriers (
    "carrier code" character(2) NOT NULL,
    "carrier name" character varying(100)  NOT NULL,
PRIMARY KEY("carrier code")
)
DISTSTYLE ALL;

CREATE TABLE public.dim_airports (
    "airport code" character(3)  NOT NULL,
    "airport name" character varying(50)  NOT NULL,
    city character varying(50) NULL,
    state character varying(2) NULL,
    country character varying(3) NOT NULL,
    latitude numeric(9,6)  NOT NULL,
    longitude numeric(9,6) NOT NULL,
 PRIMARY KEY ("airport code")
)
DISTSTYLE ALL;

CREATE TABLE public.fact_flights (
    "flight date" timestamp without time zone NOT NULL,
    "carrier code" character(2)  NOT NULL,
    "flight code" numeric(4,0)  NOT NULL,
    "tail number" character varying (6)  NOT NULL,
    "origin airport code" character(3)  NOT NULL,
    "destination airport code" character(3)  NOT NULL,
    "flight distance" numeric(4,0)  NOT NULL,
    "flight airtime" numeric(4,0)  NOT NULL,
    "scheduled departure time" numeric(4,0)  NOT NULL,
    "scheduled arrival time" numeric(4,0)  NOT NULL,
    "scheduled elasped time" numeric(4,0)  NOT NULL,
    "estimated departure time" numeric(4,0)  NULL,
    "estimated arrival time" numeric(4,0)  NULL,
    "estimated elapsed time" numeric(4,0)  NULL,
    "cancelled flag" character(1)  NOT NULL DEFAULT 'N',
    "diverted flag" character(1)  NOT NULL DEFAULT 'N',
    "aircraft delay" numeric(4,0)  NOT NULL DEFAULT 0,
    "airtraffic delay" numeric(4,0)  NOT NULL DEFAULT 0,
    "security delay" numeric(4,0) NOT  NULL DEFAULT 0,
    "weather delay" numeric(4,0)  NOT NULL DEFAULT 0,
 foreign key("origin airport code") references "dim_airports"("airport code"),
 foreign key("destination airport code") references "dim_airports"("airport code"),
 foreign key("carrier code") references "dim_carriers"("carrier code")
)
DISTSTYLE EVEN
INTERLEAVED SORTKEY ("flight date","origin airport code", "destination airport code");
```

**load-data.sql**
```sql
COPY dim_airports
FROM 's3://us-west-2-aws-training/awsu-spl/spl114-working-with-tableau/data/airports'
CREDENTIALS 'aws_access_key_id=access-key;aws_secret_access_key=secret-key'
REGION 'us-west-2'
GZIP;

COPY dim_carriers
FROM 's3://us-west-2-aws-training/awsu-spl/spl114-working-with-tableau/data/carriers'
CREDENTIALS 'aws_access_key_id=access-key;aws_secret_access_key=secret-key'
REGION 'us-west-2'
GZIP;

COPY fact_flights
FROM 's3://us-west-2-aws-training/awsu-spl/spl114-working-with-tableau/data/flights'
CREDENTIALS 'aws_access_key_id=access-key;aws_secret_access_key=secret-key'
REGION 'us-west-2'
GZIP;
```
