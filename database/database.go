package database

import (
	"database/sql"
	"fmt"
	"log"

	"lantorabde.app/helper"
	
)

func ConnectDatabase() {

	const (
		host     = "ep-royal-wind-a1qt5wjl.ap-southeast-1.pg.koyeb.app"
		port     = 5432
		user     = "koyeb-adm"
		password = "npg_5ZybEfO7Tdoj"
		dbname   = "koyebdb"
	)

	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=require", host, port, user, password, dbname)
	var err error

	helper.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	if err = helper.DB.Ping(); err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	fmt.Println("Database connection established")

	// createTable := `CREATE TABLE IF NOT EXISTS bde_users(
	// id SERIAL PRIMARY KEY,
	// full_name TEXT NOT NULL,
	// email TEXT UNIQUE NOT NULL,
	// phone TEXT UNIQUE NOT NULL,
	// password_hash TEXT NOT NULL,
	// driving_license TEXT,
	// role TEXT DEFAULT 'BDE' CHECK(role IN ('BDE','Manager','Admin')),
	// join_date DATE DEFAULT CURRENT_DATE ,
	// status TEXT DEFAULT 'ACTIVE' CHECK(status IN ('ACTIVE','INACTIVE')),
    // created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    // updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP)`
	// _, err = helper.DB.Exec(createTable)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Table created successfully")

}
