// package main

// import "fmt"

// func main() {
// 	fmt.Printf("hello world")
// }

package main

import (
	//"bufio"

	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

/* type User struct {
	Name string `json:"name"`
} */

func main() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	
	// Open the file
	csvfile, err := os.Open("csv-file.csv")
	
	// Parse the file
	r := csv.NewReader(csvfile)
	r.Read()
	
	// Iterate through the records
	for {
		// Read each record from csv
		record, err := r.Read()
		// fmt.Println(record[1])

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		insert, err := db.Query("INSERT INTO persons(first_name, last_name, age, blood_group) VALUES('" + record[0] + "','" + record[1] + "','" + record[2] + "','" + record[3] + "')")

		if err != nil {
			panic(err.Error())
		}

		defer insert.Close()
		fmt.Println("Successfully inserted a data row")
	}
	// Establishing DB connection
	/* db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	fmt.Println("Successfully Connected") */
	// Inserting data in DB

	// insert, err := db.Query("INSERT INTO users(name) VALUES('Sohaib')")

	// if err != nil {
	// 	panic(err.Error())
	// }

	// defer insert.Close()

	// fmt.Println("Successfully inserted a data row")

	// Fetching Data from DB

	/* results, err := db.Query("SELECT name FROM users")

	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var user User

		err = results.Scan(&user.Name)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println(user.Name)
	} */
	// defer results.Close()

	// fmt.Println("Successfully inserted a data row")
	

	/* mysql.RegisterReaderHandler("data", func() io.Reader {
		var csvReader io.Reader // Some Reader that returns CSV data
		file, err := os.Open("input.csv")
		if err != nil {
			log.Fatalln("Couldn't open the csv file", err)
		}
		r := csv.NewReader(file)
		r.Read()
		return csvReader
	})
	ok, err := db.Exec("LOAD DATA LOCAL INFILE 'Reader::data' INTO TABLE users")

	fmt.Println(ok) */

	/* filePath := "input.csv"

	mysql.RegisterLocalFile(filePath)
	xyz, err := db.Exec("LOAD DATA LOCAL INFILE '" + filePath + "' INTO TABLE users") // FIELDS TERMINATED BY ',' LINES TERMINATED BY '\n' IGNORE 1 ROWS
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(xyz)
	fmt.Println(db.Exec("LOAD DATA LOCAL INFILE '" + filePath + "' INTO TABLE users")) */

	
	/* fileScanner := bufio.NewScanner(csvfile)
	lineCount := 0
	for fileScanner.Scan() {
		lineCount++
	}
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	} */

	
	//r := csv.NewReader(bufio.NewReader(csvfile))
	// var colIndex []int
	
}
