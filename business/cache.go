package business

//This file is for demonstrating the cache system of the Zvuvon (question 8)
//There are 2 database table schemas: one is for storing the input history with its unique id and the other is for storing the results with the same id
//Each time the user calculates something, it will save the input and the result in the DB.
//If there inputs exists, it'll return the result right away by pulling it from the DB

/*
type cache interface {
	checkIfInputExists(input string) bool
	saveInput(input string) error
	saveOutput(output string) error
	getOutputById(id string) (string, error)
}

type DBContext struct {
	DB       *sql.DB
	DBHelper DBHelper
}
*/
