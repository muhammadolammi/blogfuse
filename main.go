package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/muhammadolammi/blogfuse/internal/database"
)

type Config struct {
	PORT string
	DB   *database.Queries
}

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
		return
	}
	port := os.Getenv("PORT")
	if port == "" {
		log.Println("empty port")
		return
	}
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Println("empty dbURL")
		return
	}
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		fmt.Println(err)
		return
	}
	dbQueries := database.New(db)

	apiConfig := Config{
		PORT: port,
		DB:   dbQueries,
	}

	// rssFeed, err := urlToRssFeed("https://blog.boot.dev/index.xml")
	// if err != nil {
	// 	log.Println("there is an error turning feed url to rssFeed, err:", err)
	// 	return
	// }
	// for _, item := range rssFeed.Channel.Item {
	// 	log.Println(item.PubDate)

	//}

	// if wait := true; wait {
	// 	return
	// }
	go startScrapping(dbQueries, 10, time.Minute)

	serverEntry(&apiConfig)
}
