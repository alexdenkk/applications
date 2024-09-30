package main

import (
	"log"
	"alexdenkk/applications/internal/server"
	"alexdenkk/applications/internal/writer"
	"strconv"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	godotenv.Load()

	chat_id, err := strconv.ParseInt(os.Getenv("CHAT"), 10, 64)

	if err != nil {
		log.Fatal(err)
	}

	w := writer.New(os.Getenv("TOKEN"), chat_id)

	srv := server.New(w)

	srv.Run()
}
