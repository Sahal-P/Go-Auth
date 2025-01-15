package db

import (
	"context"
	"fmt"
	"log"

	"github.com/Sahal-P/Go-Auth/config"
	"github.com/jackc/pgx/v4"
)

type PostgreSQLStorage struct {
	Conn *pgx.Conn
}

func NewPostgreSQLStorage() *PostgreSQLStorage {

	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", 
		config.AppConfig.DBUser, 
		config.AppConfig.DBPassword, 
		config.AppConfig.DBHost, 
		config.AppConfig.DBPort, 
		config.AppConfig.DBName)

	// if connString == "" {
	// 	connString = "postgres://postgres:09876@localhost:5432/go_auth"
	// }

	fmt.Println("\033[33mConnecting to DB...\033[0m")
	conn, err := pgx.Connect(context.Background(), connString)

	if err != nil {
		log.Fatalf("Unable to Connect to DB: %v", err)
	}
	fmt.Println("\033[32mConnected to DB\033[0m") 


	return &PostgreSQLStorage{
		Conn: conn,
	}
}

func (s *PostgreSQLStorage) Ping() error {
	// Ping the database to check the connection
	err := s.Conn.Ping(context.Background())
	if err != nil {
		return fmt.Errorf("unable to ping DB: %w", err)
	}
	fmt.Println("\033[32mSuccessfully pinged the DB\033[0m")
	return nil
}

func (s *PostgreSQLStorage) Close() {
	if err := s.Conn.Close(context.Background()); err != nil {
		log.Fatalf("Unable to close DB connection: %v", err)
	}
	fmt.Println("DB Connection Closed")
}
