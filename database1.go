package main

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// NOTE: loading the sql driver anonymously, aliasing its 
//       package qualifier to _ so none of its exported 
//       names are visible to my code.

func main() {
	fmt.Printf("In Toy Database Program...\n\n")

	db, err := sql.Open("mysql",
		"root:password@tcp(127.0.0.1:3306)/mediaq_test")
	if err != nil {
		//log.Fatal(err)
		fmt.Printf("ERROR!\n\n")
	}
	// looks like the db actually opened, so defer the close
	// from here to be sure it happens before a return
	defer db.Close()


	// Just ping the db to see what's up...
	err = db.Ping()
	if err != nil {
		// do something here
		fmt.Printf("Ping error!!\n\n")
	}

	

	var (
		id int
		name string
		link string
	)

	rows, err := db.Query("select id, name, link from items where id = ?", 1)
	if err != nil {
		//log.Fatal(err)
		fmt.Printf("Query error!\n\n")
	}
	// looks like the query worked, defer closing rows
	defer rows.Close()
	
	for rows.Next() {
		err := rows.Scan(&id, &name, &link)
		if err != nil {
			//log.Fatal(err)
			fmt.Printf("Scanning rows error!!\n\n")
		}
		//log.Println(id, name)
		fmt.Printf(name + "\n\n")
	}

	err = rows.Err()
	if err != nil {
		//log.Fatal(err)
		fmt.Printf("rows.error\n\n")
	}

}