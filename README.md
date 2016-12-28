# StockSymbolLoader
EOD Data Symbol List Loader

## Overview

Small Program written in Go that loads a given symbol list file made up of CSV data, provided by EODData.com.  You can get symbol list are free with an account sign-up.

Current state the program has only been tested with NASDAQ, TSX and TSXV.  Shouldn't be problem to run the other exchanges as well, code may need to be modified to handle other symbols with suffix excchange indicators.  Open a pull request or an issue, more then happy to update the code.

## Prerequisite 

* [Glide](https://glide.sh/] a Go package manager

## Getting Started

#### Install Vendor Packages
Assuming you cloned the project and Go environment is configured correctly, first thing you need to do is install the vendor packages using glide:

```bash
> cd src
> glide install
```

Note due to Golang *vendor* directory requirements glide.yaml and glide must be ran under *src* directory.

#### Create MySQL DB Table

It's assumed you know what to do here, this really just a reminder you need to apply the symbols DDL file under /etc/ the file named 
symbols.sql

#### Update Config file

At this point in time the config file contains configuration for database access

```bash
[database]
username="mysql-username"
password="mysql-password"
dbname="mysql-database-name"
host="mysql-host-location"
port="mysql-host-port"
```

Be sure to rename the file *stocksymbolloader.example.toml* to *stocksymbolloader.toml*

#### Replace Data files

For testing purpose only you can use the existing data files under the *dat* folder othewise you will need to replace the files
with actual files from EODData.com

#### Compile and Run the application

```bash
go build github.com/jasonlam604/StockSymbolLoader/
```
```bash
./StockSymbolLoader
```
