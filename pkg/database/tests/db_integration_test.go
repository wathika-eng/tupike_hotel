package database

// import (
// 	"context"
// 	"database/sql"
// 	"fmt"
// 	"log"
// 	"testing"
// 	"time"

// 	"github.com/stretchr/testify/assert"
// 	"github.com/testcontainers/testcontainers-go"
// 	"github.com/testcontainers/testcontainers-go/wait"
// 	"github.com/uptrace/bun"
// 	"github.com/uptrace/bun/dialect/pgdialect"
// 	"github.com/uptrace/bun/driver/pgdriver"
// )

// // setupTestDB starts a PostgreSQL container for testing and returns its DSN
// func setupTestDB(ctx context.Context, t *testing.T) (string, func()) {
// 	t.Helper() // Marks this function as a test helper

// 	// Define the PostgreSQL container
// 	req := testcontainers.ContainerRequest{
// 		Image:        "postgres:latest",
// 		ExposedPorts: []string{"5432/tcp"},
// 		Env: map[string]string{
// 			"POSTGRES_USER":     "testuser",
// 			"POSTGRES_PASSWORD": "testpassword",
// 			"POSTGRES_DB":       "testdb",
// 		},
// 		WaitingFor: wait.ForLog("database system is ready to accept connections").
// 			WithStartupTimeout(30 * time.Second),
// 	}

// 	// Start the container
// 	postgresContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
// 		ContainerRequest: req,
// 		Started:          true,
// 	})
// 	if err != nil {
// 		t.Fatalf("❌ Failed to start PostgreSQL container: %v", err)
// 	}

// 	// Get container's host and mapped port
// 	host, err := postgresContainer.Host(ctx)
// 	if err != nil {
// 		t.Fatalf("❌ Failed to get container host: %v", err)
// 	}
// 	port, err := postgresContainer.MappedPort(ctx, "5432")
// 	if err != nil {
// 		t.Fatalf("❌ Failed to get container port: %v", err)
// 	}

// 	// Build DSN
// 	dsn := fmt.Sprintf("postgres://testuser:testpassword@%s:%s/testdb?sslmode=disable", host, port.Port())

// 	// Cleanup function to terminate the container
// 	cleanup := func() {
// 		log.Println("🛑 Stopping PostgreSQL test container...")
// 		if err := postgresContainer.Terminate(ctx); err != nil {
// 			t.Errorf("⚠️ Failed to terminate PostgreSQL container: %v", err)
// 		}
// 	}

// 	return dsn, cleanup
// }

// // setupDB initializes a Bun database instance
// func setupDB(dsn string, t *testing.T) (*service, func()) {
// 	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
// 	db := bun.NewDB(sqldb, pgdialect.New())

// 	// Check if the database is reachable
// 	if err := db.Ping(); err != nil {
// 		t.Fatalf("❌ Database connection failed: %v", err)
// 	}

// 	log.Println("✅ Test database connected successfully")

// 	svc := &service{db: db}

// 	// Cleanup function to close DB connection
// 	cleanup := func() {
// 		if err := svc.Close(); err != nil {
// 			t.Errorf("⚠️ Failed to close database connection: %v", err)
// 		}
// 	}

// 	return svc, cleanup
// }

// // TestIntegration runs the database integration test
// func TestIntegration(t *testing.T) {
// 	ctx := context.Background()

// 	// Setup test database container
// 	dsn, cleanupContainer := setupTestDB(ctx, t)
// 	defer cleanupContainer()

// 	// Setup database instance
// 	svc, cleanupDB := setupDB(dsn, t)
// 	defer cleanupDB()

// 	// Run the Health check
// 	stats := svc.Health()
// 	assert.Equal(t, "up", stats["status"], "Expected database status to be 'up'")
// 	assert.Equal(t, "It's healthy", stats["message"], "Expected health message to indicate 'healthy'")

// 	// Test closing the database
// 	err := svc.Close()
// 	assert.NoError(t, err, "Expected database to close without errors")
// }
