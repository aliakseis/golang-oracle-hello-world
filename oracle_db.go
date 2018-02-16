package main
    
import (
    "fmt"
    "database/sql"
    _ "github.com/mattn/go-oci8"
)
    
func main(){
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
    
    rows,err := db.Query("select 2+2 from dual")
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
