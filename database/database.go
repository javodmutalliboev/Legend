package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func DB() *sql.DB {
	//portStr := os.Getenv("DB_PORT")
	//port, err := strconv.Atoi(portStr)
	//if err != nil {
	//	log.Println(err)
	//}

	/*
		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			os.Getenv("DB_HOST"), port, os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"))
	*/

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db
}
