package database

// import (
// 	"database/sql"
// 	"log"

// 	_ "github.com/lib/pq" // PostgreSQL driver
// )

// // Define a global variable to hold the database connection.
// var db *sql.DB

// // InitializeDatabase initializes the database connection.
// func InitializeDatabase(connectionString string) error {
// 	var err error
// 	db, err = sql.Open("postgres", connectionString)
// 	if err != nil {
// 		return err
// 	}

// 	err = db.Ping()
// 	if err != nil {
// 		return err
// 	}

// 	log.Println("Connected to the database")
// 	return nil
// }

// // CloseDatabase closes the database connection.
// func CloseDatabase() {
// 	if db != nil {
// 		db.Close()
// 		log.Println("Database connection closed")
// 	}
// }

// // Example function to execute a query and retrieve data from the database.
// func GetUsers() ([]User, error) {
// 	// Define the User struct and any other necessary data structures.

// 	rows, err := db.Query("SELECT id, name, email FROM users")
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var users []User
// 	for rows.Next() {
// 		var user User
// 		err := rows.Scan(&user.ID, &user.Name, &user.Email)
// 		if err != nil {
// 			return nil, err
// 		}
// 		users = append(users, user)
// 	}

// 	if err = rows.Err(); err != nil {
// 		return nil, err
// 	}

// 	return users, nil
// }

// // Example function to insert a new user into the database.
// func InsertUser(user User) error {
// 	_, err := db.Exec("INSERT INTO users (name, email) VALUES ($1, $2)", user.Name, user.Email)
// 	return err
// }
