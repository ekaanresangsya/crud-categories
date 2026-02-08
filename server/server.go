package server

import (
	"fmt"
	"kasir-api/internal/database"
	"log"
	"time"
)

func Start() {
	log.Print("Starting server ...")

	fmt.Println(time.Now())

	config := LoadConfig()

	db, err := database.InitDB(config.DBConn)
	if err != nil {
		log.Fatalf("error connecting to database, got %v", err)
	}
	defer db.Close()

	router := InitRouter(db)

	port := config.ServerPort
	err = router.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		log.Printf("error running server, got %v", err)
	}
}
