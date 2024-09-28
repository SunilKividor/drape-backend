package main

import (
	"log"
	"os"

	"github.com/SunilKividor/drape/internal/configs"
	"github.com/SunilKividor/drape/internal/database"
	"github.com/SunilKividor/drape/internal/server"
	"github.com/joho/godotenv"
)

func main() {
	//loading .env file
	err := godotenv.Load("/Users/sunil/go/src/go-drape/cmd/server/.env")
	if err != nil {
		log.Printf("Error type: %T", err)
		log.Fatalf("error : Error loading .env")
	}

	//db server
	sqlConfig := configs.DBConfig{
		DBDriver:   "sql",
		DBHost:     os.Getenv("DBHost"),
		DBUser:     os.Getenv("DBUser"),
		DBPassword: os.Getenv("DBPassword"),
		DBPort:     os.Getenv("DBPort"),
		DBname:     os.Getenv("DBname"),
	}
	database := database.NewDbServer(sqlConfig)
	_, err = database.StartDatabase()
	if err != nil {
		log.Println("db server error")
		log.Fatal(err)
	}

	//server
	s := server.GetServer(":3000")
	err = s.StartServer()
	if err != nil {
		log.Fatalf("error: Error starting the server")
	}
}
