package main

import (
	"bufio"
	"database/sql" //Go's SQL package
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql" //Third party MySQL driver. We're aliasing the driver to _, so that its exported names are not visible to our code.
)

var (
	comment string
)

func main() {

	fmt.Println("Connecting to MySQL...")

	//Read username and password

	fmt.Print("Please enter the username:")
	reader := bufio.NewReader(os.Stdin)
	username, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Please enter the password:")
	reader = bufio.NewReader(os.Stdin)
	password, err1 := reader.ReadString('\n')
	if err1 != nil {
		log.Fatal(err1)
	}
	username = strings.Replace(username, "\n", "", -1)
	password = strings.Replace(password, "\n", "", -1)

	//Connecting to db
	dataSourceName := username + ":" + password + "@tcp(127.0.0.1:3306)/testdb"
	sqlHandler, err2 := sql.Open("mysql", dataSourceName)
	if err2 != nil {
		log.Fatal(err2)
	}

	defer sqlHandler.Close()

	//Creating a table within MySQL

	_, err3 := sqlHandler.Exec("CREATE TABLE random_strings (comments VARCHAR(100));")
	if err3 != nil {
		log.Fatal(err3)
	}

	// Inserting values into a table within MySQL

	exampleStringToStoreInDb := "Test string"
	_, err4 := sqlHandler.Exec("INSERT INTO random_strings (comments) VALUES (?);", exampleStringToStoreInDb) //? is a place holder.
	if err4 != nil {
		log.Fatal(err4)
	}

	// Read from the table

	rowsFromDb, err5 := sqlHandler.Query("SELECT * FROM random_strings;")
	if err5 != nil {
		log.Fatal(err5)
	}

	defer rowsFromDb.Close()

	for rowsFromDb.Next() {
		err6 := rowsFromDb.Scan(&comment)
		if err6 != nil {
			fmt.Println(err6)
		}
		fmt.Println(comment)
	}

}
