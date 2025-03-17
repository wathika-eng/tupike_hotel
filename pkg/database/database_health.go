package database

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"
)

func (s *service) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// Configure database connection settings
	//maxLifeConns, _ := strconv.Atoi(envs.DB_MAX_LIFETIME_CONNS)
	//s.db.SetConnMaxLifetime(time.Duration(maxLifeConns) * time.Second)
	maxIdleConns, _ := strconv.Atoi(envs.DbMaxIdleConns)
	s.db.SetMaxIdleConns(maxIdleConns)
	maxOpenConns, _ := strconv.Atoi(envs.DbMaxOpenConns)
	s.db.SetMaxOpenConns(maxOpenConns)

	stats := make(map[string]string)

	// Measure ping latency
	start := time.Now()
	err := s.db.PingContext(ctx)
	pingLatency := time.Since(start).Milliseconds() // Latency in milliseconds

	if err != nil {
		stats["status"] = "down"
		stats["error"] = fmt.Sprintf("db down: %v", err)
		log.Fatalf("db down: %v", err) // to fix
		return stats
	}

	// Database is up, add more statistics
	stats["status"] = "up"
	stats["message"] = "It's healthy"
	stats["ping_latency_ms"] = fmt.Sprintf("%d", pingLatency) // Add ping latency to stats

	// Get database stats (like open connections, in use, idle, etc.)
	DBService := s.db.Stats()
	stats["open_connections"] = strconv.Itoa(DBService.OpenConnections)
	stats["in_use"] = strconv.Itoa(DBService.InUse)
	stats["idle"] = strconv.Itoa(DBService.Idle)
	stats["wait_count"] = strconv.FormatInt(DBService.WaitCount, 10)
	stats["wait_duration"] = DBService.WaitDuration.String()
	stats["max_idle_closed"] = strconv.FormatInt(DBService.MaxIdleClosed, 10)
	stats["max_lifetime_closed"] = strconv.FormatInt(DBService.MaxLifetimeClosed, 10)

	// Evaluate stats to provide a health message
	if DBService.OpenConnections > 40 { // Assuming 50 is the max for this example
		stats["message"] = "The database is experiencing heavy load."
	}

	if DBService.WaitCount > 1000 {
		stats["message"] = "The database has a high number of wait events, indicating potential bottlenecks."
	}

	if DBService.MaxIdleClosed > int64(DBService.OpenConnections)/2 {
		stats["message"] = "Many idle connections are being closed, consider revising the connection pool settings."
	}

	if DBService.MaxLifetimeClosed > int64(DBService.OpenConnections)/2 {
		stats["message"] = "Many connections are being closed due to max lifetime, consider increasing max lifetime or revising the connection usage pattern."
	}

	// Add additional metrics (e.g., connection speed)
	// Note: Connection speed is not directly measurable in Go's database/sql package,
	// but you can estimate it by measuring the time taken to execute a simple query.
	startQuery := time.Now()
	_, err = s.db.ExecContext(ctx, "SELECT 1") // Simple query to measure speed
	queryLatency := time.Since(startQuery).Milliseconds()
	if err != nil {
		stats["query_speed_ms"] = "failed"
	} else {
		stats["query_speed_ms"] = fmt.Sprintf("%d", queryLatency)
	}

	return stats
}
