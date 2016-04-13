package main

import ( 
	"fmt"
	"os"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	if len(os.Args) != 3 {
		os.Exit(1) 
	} 
   	
   	//fmt.Println("It's over ", os.Args[1])
   	fmt.Println("Attempting Insert into items table:")
   	fmt.Println(os.Args[1])
   	fmt.Println(os.Args[2] + "\n")

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

	// prepare statement for insert
	stmt, err := db.Prepare("INSERT INTO items(name, link) VALUES(?, ?)")
	if err != nil {
		//log.Fatal(err)
	}
	//res, err := stmt.Exec("Dolly", "www.dolly-pants.com")
	res, err := stmt.Exec(os.Args[1], os.Args[2])
	if err != nil {
		//log.Fatal(err)
		fmt.Printf("Error inserting!!")
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		//log.Fatal(err)
		fmt.Printf("Error getting last insert id")
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		//log.Fatal(err)
		fmt.Printf("Error getting rows affected")
	}
	fmt.Printf("ID = %d, affected = %d\n", lastId, rowCnt)
}