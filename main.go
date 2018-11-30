package main

import (
	"database/sql"
  "fmt"
  "flag"
  "strconv"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
  host := flag.String("host", "localhost", "The address of MySQL server")
  port := flag.Int("port", 3306, "Port of MySQL service")
  database := flag.String("database", "", "Database to connect")
  user := flag.String("user", "root", "MySQL user")
  password := flag.String("password", "", "Password for MySQL user")
  query := flag.String("query", "", "SQL query to execute")
  flag.Parse()

  if *query == "" {
    panic("Please specify SQL query to execute.")
  }
	db, err := sql.Open("mysql", *user + ":" + *password + "@tcp("+ *host + ":" + strconv.Itoa(*port) +")/" + *database)
	if err != nil {
		panic("Can not open connection to MySQL service by address " + *host)
	}
	defer db.Close()

  rows, err := db.Query(*query)
	if err != nil {
		panic(err.Error())
	}

	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}

	values := make([]sql.RawBytes, len(columns))

	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}

		var value string
		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			fmt.Println(columns[i], ": ", value)
		}
	}
	if err = rows.Err(); err != nil {
		panic(err.Error())
	}
}
