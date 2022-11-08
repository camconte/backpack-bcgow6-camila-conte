package db

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

var StorageDB *sql.DB

func init(){
	//Cargamos el archivo de variables
	/* err := godotenv.Load()
	if err != nil {
		log.Fatal("Error: Loading .env")
	} */

	//Seteamos la configuracion de la conexion a la base de datos
	configDB := mysql.Config{
		User:   "root",
		Passwd: "",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "storage",
	}

	var err error
	//abrimos la conexion
	StorageDB, err = sql.Open("mysql", configDB.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

}

/* func ConnectDatabase() (engine *gin.Engine, db *sql.DB) {
	//Cargamos el archivo de variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error: Loading .env")
	}

	//Seteamos la configuracion de la conexion a la base de datos
	configDB := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: os.Getenv("DBNAME"),
	}

	//abrimos la conexion
	db, err = sql.Open("mysql", configDB.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	engine = gin.Default()

	return engine, db
} */