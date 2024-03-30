package db

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"os"
)

func InitDB() (*pgxpool.Pool, error) {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable is required")
	}

	dbPool, err := pgxpool.Connect(context.Background(), dbURL)
	if err != nil {
		return nil, err
	}

	return dbPool, nil
}

func CloseDB(dbPool *pgxpool.Pool) {
	dbPool.Close()
}

func insertPatient(dbPool *pgxpool.Pool, patient Patient) error {
	// Insert patient into database. Simplified for demonstration.
	_, err := dbPool.Exec(context.Background(), "INSERT INTO patients (id, full_name, birth_date) VALUES ($1, $2, $3) ON CONFLICT (id) DO NOTHING", patient.ID, patient.FullName, patient.BirthDate)
	return err
}

func insertClaim(dbPool *pgxpool.Pool, claim Claim) error {
	// Insert claim into database. Simplified for demonstration.
	_, err := dbPool.Exec(context.Background(), "INSERT INTO claims (id, patient_id, type, amount) VALUES ($1, $2, $3, $4) ON CONFLICT (id) DO NOTHING", claim.ID, claim.PatientID, claim.Type, claim.Amount)
	return err
}
