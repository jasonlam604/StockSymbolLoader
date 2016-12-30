// Small simple App to load Symbols into a MySQL table, dependency file CSV format
// defined by www.eoddata.com.  That said this program can easily be modified 
// to support a different format. Authored by Jason Lam, jasonlam604@gmail.com.
// Code repository hosted on Github at https://github.com/jasonlam604/StockSymbolLoader
// Code is released under MIT License.
package main

import (
	"bufio"
	"database/sql"
	"encoding/csv"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pelletier/go-toml"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

// Data Structure to hold Instrument / Symbol info
type Symbol struct {
	code        string
	companyName string
	exchange    string
}

// Slice holding a number of 'symbol'
var symbols []Symbol

// Holder for Errors
var err error

// MySQL DB Handler
var db *sql.DB

// Config
var config *toml.TomlTree

// Path to Data (CSV symbol files)
var dataDir string

// Main Program Entry
func main() {
	loadFiles()
	dbBatchInsert()
	dbClose()
}

// Initial Program
func init() {

	// Locate Symobls CSV files, read from $GOPATH/dat (EOD Data files, sample files reside here, you
	// need to download the actual symbol files from Eoddata.com
	dataDir = path.Join(os.Getenv("GOPATH"), "dat")

	// Init Config
	configInit()

	// Init Logging
	logInit()

	// Init Database Connection
	dbInit()
	
	fmt.Println("Symbols Loaded.")
}

// Init Config, read configuration file from $GOPATH/config
func configInit() {
	config, err = toml.LoadFile(path.Join(os.Getenv("GOPATH"), "config", "stocksymbolloader.toml"))
	if err != nil {
		fmt.Println("Error ", err.Error())
	}
}

// Init Logging, set log file location under $GOPATH/config
func logInit() {
	f, err := os.OpenFile(path.Join(os.Getenv("GOPATH"), "log", "stocksymbolloader.log"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
	}
	log.SetOutput(f)
}

// Init Database Connection
func dbInit() {
	db, err = sql.Open("mysql", config.Get("database.username").(string)+":"+config.Get("database.password").(string)+"@tcp("+config.Get("database.host").(string)+":"+config.Get("database.port").(string)+")/"+config.Get("database.dbname").(string))

	if err != nil {
		log.Fatal(err)
	}

	// Test Open Connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}

// Close Database
func dbClose() {
	defer db.Close()
}

// Load symbols from text files (CSV formatted)
func loadFiles() {
	files, _ := ioutil.ReadDir(dataDir)
	for _, f := range files {
		readFile(f.Name())
	}
}

// Read file line by line
func readFile(filename string) {

	// Load a TXT file.
	f, _ := os.Open(path.Join(dataDir, filename))

	// Create a new reader.
	r := csv.NewReader(bufio.NewReader(f))
	r.Comma = '\t'

	for {
		record, err := r.Read()

		if err == io.EOF {
			break
		}

		if record[0] != "Symbol" {
			aSymbol := Symbol{sanitizeSymbol(record[0]), record[1], parseExchange(filename)}
			symbols = append(symbols, aSymbol)
		}
	}
}

// Drops Suffix TO or V (TSX opr TSXV)
func sanitizeSymbol(symbol string) string {

	sym := symbol
	symParts := strings.Split(symbol, ".")
	symSuffix := symParts[len(symParts)-1]

	if (symSuffix == "TO") || (symSuffix == "V") {
		symParts = symParts[:len(symParts)-1]
		sym = strings.Join(symParts, ".")
	}

	return sym
}

// Parse exchange from filename (drop extension)
func parseExchange(filename string) string {
	filenameParts := strings.Split(filename, ".")
	return filenameParts[0]
}

// Batch MySQL insert symbol data
func dbBatchInsert() {
	valueStrings := make([]string, 0, len(symbols))
	valueArgs := make([]interface{}, 0, len(symbols)*3)
	for _, symbol := range symbols {
		valueStrings = append(valueStrings, "(?, ?, ?)")
		valueArgs = append(valueArgs, symbol.exchange)
		valueArgs = append(valueArgs, symbol.code)
		valueArgs = append(valueArgs, symbol.companyName)
	}
	stmt := fmt.Sprintf("INSERT INTO symbols (exchange, code, company_name) VALUES %s", strings.Join(valueStrings, ","))
	_, err := db.Exec(stmt, valueArgs...)
	if err != nil {
		log.Fatal(err)
	}
}
