# StockSymbolLoader
EOD Data Symbol List Loader

## Overview

Small program written in Golang that loads symbol list files made up of CSV data.  You can get symbol lists for free with an account sign-up at EODData.com

Current state, the program has only been tested with NASDAQ, NYSE, TSX and TSXV.  Shouldn't be a problem to run the other exchanges as well, code may need to be modified to handle other symbols with suffix exchange indicators.  Open a pull request or an issue, more then happy to update the code.

## Prerequisite 

* GOPATH is set correctly
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

It is assumed you know what to do here, this is really just a reminder you need to apply the symbols DDL file under /etc/ the file named 
symbols.sql to your database.

#### Update Config file

At this point in time the config file contains configuration for database access, of your you need to apply your settings here.

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

For testing purposes only you can use the existing data files under the *dat* folder othewise you will need to replace the files
with actual files from EODData.com

#### Compile and Run the application

```bash
go build github.com/jasonlam604/StockSymbolLoader/
```
```bash
./StockSymbolLoader
```

## Debugging

Log file is found under *log* directory, this is where you will find messages like duplicate entry...
