package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	r "github.com/dancannon/gorethink"
)

var (
	address, db, table string
	session            *r.Session
)

func init() {
	flag.StringVar(&address, "address", "localhost:28015", "The address to connect to a RethinkDB instance.")
	flag.StringVar(&db, "database", "", "The RethinkDB database to use.")
	flag.StringVar(&table, "table", "", "The RethinkDB database table to use.")
	flag.Parse()

	if len(db) == 0 {
		fmt.Println("You must pass the name of a database to use.\n")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if len(table) == 0 {
		fmt.Println("You must pass the name of database table to use.\n")
		flag.PrintDefaults()
		os.Exit(1)
	}

	//session, err := r.Connect(r.ConnectOpts{
	//	Address:  address,
	//	Database: db,
	//	MaxIdle:  10,
	//	MaxOpen:  10,
	//})
	//if err != nil {
	//	log.Fatalln(err.Error())
	//}
}

func main() {
	saveData(getData())
}

func saveData(data []map[string]interface{}) {
	fmt.Println(data)
}

// Reads from stdin, then marshals the line into a Go value
func getData() []map[string]interface{} {
	reader := bufio.NewReader(os.Stdin)

	var line []byte
	for {
		read, isPrefix, err := reader.ReadLine()
		if err != nil {
			log.Fatalf("Error reading from Stdin: %s", err.Error())
		}
		line = append(line, read...)
		if !isPrefix {
			break
		}
	}

	var data []map[string]interface{}
	if err := json.Unmarshal(line, &data); err != nil {
		log.Fatalf("Could not unmarshal JSON. Error: %s", err.Error())
	}

	return data
}
