package server

import (
	"crud-categories/internal/database"
	"fmt"
	"log"
	"os"
)

func Start() {
	log.Print("Starting server ...")

	config := LoadConfig()

	fmt.Println("config")
	fmt.Println(config.DBConn)
	fmt.Println(config.ServerPort)
	fmt.Println("------------------------------")
	fmt.Println("env")
	fmt.Println(os.Getenv("DB_CONN"))
	fmt.Println(os.Getenv("SERVER_PORT"))

	db, err := database.InitDB(os.Getenv("DB_CONN"))
	if err != nil {
		log.Fatalf("error connecting to database, got %v", err)
	}
	defer db.Close()

	router := InitRouter(db)

	port := os.Getenv("SERVER_PORT")
	err = router.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		log.Printf("error running server, got %v", err)
	}
}
