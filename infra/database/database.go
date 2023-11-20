package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

var (
	host     = os.Getenv("DB_HOST")
	port     = os.Getenv("DB_PORT")
	user     = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	dbname   = os.Getenv("DB_NAME")
)

var (
	db  *sql.DB
	err error
)

func handlerDatabaseConnection() {
	sqlinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err = sql.Open("postgres", sqlinfo)

	if err != nil {

		log.Panic("Error Validate database", err)
	}

	err = db.Ping()

	if err != nil {

		log.Panic("Error Connect to database", err)
	}

}

func handlerDatabaseInstance() {
	todoTable := `CREATE TABLE IF NOT EXISTS "todo" ( todo_id serial primary key, title varchar (225) not null, completed bool, Created_at timestamptz default now(), Updated_at timestamptz default now() );`

	_, err := db.Exec(todoTable)

	if err != nil {
		log.Panic("Error create table", err)
	}
}

func InitDatabase() {
	handlerDatabaseConnection()
	handlerDatabaseInstance()
}

func GetDatabaseInstance() *sql.DB {

	if db == nil {
		log.Panic("Erorr saat get db")
	}
	return db
}
