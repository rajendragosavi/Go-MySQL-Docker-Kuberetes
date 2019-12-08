package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	// Create a handler - db
	db, err := sql.Open("mysql", "user:password@/demoDB")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//open connection for DB
	err = db.Ping()
	if err != nil {
		log.Fatalln("Error Connecting Database")
		panic(err.Error())
	}
	//select the database ..
	_, err = db.Exec("USE demoDB")
	if err != nil {
		panic(err.Error())
	} else {
		log.Println("Successfully created the table ...")
	}
	_, err = db.Exec("CREATE TABLE Voters(VoterID   INT  NOT NULL, LastName VARCHAR (255)  NOT NULL,FirstName VARCHAR(255)  NOT NULL,State  VARCHAR (255) ,City   VARCHAR (255),  PRIMARY KEY (VoterID) );")
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("new table crated successully")
	}

}
