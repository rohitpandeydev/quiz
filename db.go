package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func Connect() (*pgx.Conn, error) {
	// Debugging: Print environment variables
	if err := godotenv.Load(); err != nil {
		log.Printf("Error when loading the env variable")
	}
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	conn, err := pgx.Connect(context.Background(), fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, host, port, dbname))
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func QueryData(conn *pgx.Conn) {
	rows, err := conn.Query(context.Background(), "SELECT id, question FROM quiz")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var quest []byte
		var que question
		err := rows.Scan(&id, &quest)
		if err != nil {
			log.Fatal(err)
		}
		err = json.Unmarshal(quest, &que)
		if err != nil {
			log.Fatal(err)
		}
		questions = append(questions, que)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
