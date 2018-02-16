package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"

	_ "github.com/mattn/go-oci8"
)

func helloWorld(db *sql.DB) {
	rows, err := db.Query("select 2+2 from dual")
	if err != nil {
		fmt.Println("Error fetching addition")
		fmt.Println(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var sum int
		rows.Scan(&sum)
		fmt.Printf("2 + 2 always equals: %d\n", sum)
	}
}

func getJSON(db *sql.DB, sqlString string) (string, error) {
	rows, err := db.Query(sqlString)
	if err != nil {
		return "", err
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return "", err
	}
	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			val := values[i]
			if s, ok := val.(string); ok {
				var v interface{}
				if i, err := strconv.Atoi(s); err == nil {
					v = i
				} else {
					v = val
				}
				entry[col] = v
			} else {
				entry[col] = val
			}
		}
		tableData = append(tableData, entry)
	}
	jsonData, err := json.Marshal(tableData)
	if err != nil {
		return "", err
	}
	//fmt.Println(string(jsonData))
	return string(jsonData), nil
}

func main() {
	// db, err := sql.Open("oci8", "username/password@localhost:1521/xe")
	db, err := sql.Open("oci8", "system/adm123@localhost:1521/xe")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		fmt.Printf("Error connecting to the database: %s\n", err)
		return
	}

	helloWorld(db)

	if jsonString, err := getJSON(db, "SELECT * FROM HR.DEPARTMENTS"); err == nil {
		fmt.Println(jsonString)
	}
}
