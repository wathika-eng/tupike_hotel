package database

// import (
// 	"errors"
// 	"testing"

// 	"github.com/DATA-DOG/go-sqlmock"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/uptrace/bun"
// 	"github.com/uptrace/bun/dialect/pgdialect"
// )

// func TestHealth(t *testing.T) {
// 	// Create a mock SQL database with ping monitoring enabled
// 	sqldb, mock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
// 	if err != nil {
// 		t.Fatalf("❌ Error creating mock database: %v", err)
// 	}
// 	defer sqldb.Close()

// 	// Create a bun.DB instance with the mock database
// 	db := bun.NewDB(sqldb, pgdialect.New())

// 	// Create the service instance
// 	svc := &service{db: db}

// 	// Mock the PingContext call
// 	mock.ExpectPing()

// 	// Call the Health method
// 	stats := svc.Health()

// 	// Assert the results
// 	assert.Equal(t, "up", stats["status"], "Expected status to be 'up'")
// 	assert.Equal(t, "It's healthy", stats["message"], "Expected message to be 'It's healthy'")

// 	// Ensure all expectations were met
// 	assert.NoError(t, mock.ExpectationsWereMet(), "Unexpected mock behavior")
// }

// func TestHealth_Down(t *testing.T) {
// 	// Create a mock SQL database with ping monitoring enabled
// 	sqldb, mock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
// 	if err != nil {
// 		t.Fatalf("❌ Error creating mock database: %v", err)
// 	}
// 	defer sqldb.Close()

// 	// Create a bun.DB instance with the mock database
// 	db := bun.NewDB(sqldb, pgdialect.New())

// 	// Create the service instance
// 	svc := &service{db: db}

// 	// Mock the PingContext call to return an error
// 	mock.ExpectPing().WillReturnError(errors.New("database down"))

// 	// Call the Health method
// 	stats := svc.Health()

// 	// Assert the results
// 	assert.Equal(t, "down", stats["status"], "Expected status to be 'down'")
// 	assert.Contains(t, stats["error"], "database down", "Expected error message to contain 'database down'")

// 	// Ensure all expectations were met
// 	assert.NoError(t, mock.ExpectationsWereMet(), "Unexpected mock behavior")
// }

// func TestClose(t *testing.T) {
// 	// Create a mock SQL database
// 	sqldb, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("❌ Error creating mock database: %v", err)
// 	}
// 	defer sqldb.Close()

// 	// Create a bun.DB instance with the mock database
// 	db := bun.NewDB(sqldb, pgdialect.New())

// 	// Create the service instance
// 	svc := &service{db: db}

// 	// Mock the Close call
// 	mock.ExpectClose()

// 	// Call the Close method
// 	err = svc.Close()

// 	// Assert the results
// 	assert.NoError(t, err, "Expected Close() to succeed without errors")

// 	// Ensure all expectations were met
// 	assert.NoError(t, mock.ExpectationsWereMet(), "Unexpected mock behavior")
// }
