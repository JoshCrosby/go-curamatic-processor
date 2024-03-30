package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"go-curamatic-processor/internal/db"
	"go-curamatic-processor/internal/notification"
	"log"
	"os"
)

func main() {
	// TODO: move this to another function outside of Main
	// Initialize the database connection.
	dbPool, err := pgxpool.NewWithConfig(context.Background(), db.Config())
	if err != nil {
		notification.Send("Error connecting to the database.")
		// for debugging
		log.Fatalf("Unable to connect to database: %v", err)
	}
	connection, err := dbPool.Acquire(context.Background())
	if err != nil {
		notification.Send("Error acquiring connection from the pool: " + err.Error())
	}
	defer connection.Release()

	err = connection.Ping(context.Background())
	if err != nil {
		notification.Send("Error pinging the database: " + err.Error())
	}

	fmt.Println("Connected to the database.")

	// Logic - grab fhir data from source, validate, and insert into the database
	// ProcessFHIRData(dbPool)
	db.CreateTableQuery(dbPool)
	db.InsertQuery(dbPool)
	db.SelectQuery(dbPool)

	// Close the database connection when the application exits
	defer dbPool.Close()
}

func processFile(filePath string) ([]Patient, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	var patients []Patient
	var totalRecords, validRecords int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var p Patient
		totalRecords++

		if err := json.Unmarshal(scanner.Bytes(), &p); err != nil {
			fmt.Printf("Warning: skipping line due to error: %v\n", err)
			continue // Skip lines that can't be unmarshaled into the Patient struct
		}

		// Add your validation logic here.
		// For example, check if required fields are present and valid.
		// If the record is not valid according to your criteria, continue to the next iteration.
		// if !isValid(p) {
		//     continue
		// }

		validRecords++
		patients = append(patients, p)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	// Check if at least 50% of the records are valid
	if float64(validRecords)/float64(totalRecords) < 0.5 {
		// Trigger a notification or handle this condition as needed
		notification.Send("Less than 50% of the records are valid. Processing stopped.")
		return nil, fmt.Errorf("less than 50% of the records are valid")
	}

	return patients, nil
}

func processPatientFile(filePath string) ([]Patient, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	var patients []Patient
	var totalRecords, validRecords int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var p Patient
		totalRecords++

		if err := json.Unmarshal(scanner.Bytes(), &p); err != nil {
			fmt.Printf("Warning: skipping line due to error: %v\n", err)
			continue // Skip lines that can't be unmarshaled into the Patient struct
		}

		// Add your validation logic here.
		// For example, check if required fields are present and valid.
		// If the record is not valid according to your criteria, continue to the next iteration.
		// if !isValid(p) {
		//     continue
		// }

		validRecords++
		patients = append(patients, p)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	// Check if at least 50% of the records are valid
	if float64(validRecords)/float64(totalRecords) < 0.5 {
		// Trigger a notification or handle this condition as needed
		sendNotification("Less than 50% of the records are valid. Processing stopped.")
		return nil, fmt.Errorf("less than 50% of the records are valid")
	}

	return patients, nil
}

//// Weekly task to process FHIR data.
//scheduler.ScheduleTask(168*time.Hour, func() {
//	err := ProcessFHIRData(dbPool)
//	if err != nil {
//		notification.Send("Error processing FHIR data: " + err.Error())
//	}
//})

//// Keep the application running.
//select {} // In a real-world application, use a more robust mechanism to keep the service running or to handle graceful shutdown.

//// ProcessFHIRData encapsulates the logic to fetch, validate, and insert FHIR data.
//// This is a simplified version. Implement according to your project requirements.
//func ProcessFHIRData(dbPool *dbPool) error {
//	// Example processing logic. You would replace this with actual calls to fetch and process data.
//	log.Println("Starting to process FHIR data...")
//
//	// Example: Process patients.
//	// This would involve fetching data, validating it, and inserting valid records into the database.
//	// You would implement this logic in detail, likely in a separate package/function.
//	err := processPatients(dbPool)
//	if err != nil {
//		log.Printf("Error processing patients: %v", err)
//		return err
//	}
//
//	// Similar for claims or any other FHIR resources you need to process.
//
//	log.Println("Finished processing FHIR data.")
//	return nil
//}

//// processPatients is a placeholder for the actual data processing logic for patients.
//func processPatients(dbPool *DBPool) error {
//	// Placeholder for fetch, validate, and insert logic.
//	// In a real application, you'd fetch the data, validate it against your criteria,
//	// and insert the valid records into the database using the db package.
//	return nil
//}
