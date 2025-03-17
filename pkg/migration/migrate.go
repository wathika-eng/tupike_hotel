package migrations

import (
	"context"
	"log"
	"time"
	"tupike_hotel/pkg/types"

	"github.com/uptrace/bun"
)

func Migrate(db *bun.DB) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*60))
	defer cancel()
	_, err := db.ExecContext(ctx, `CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`)
	if err != nil {
		log.Fatalf("❌ Failed to enable uuid-ossp extension: %v", err)
	}

	// Customers table
	_, err = db.NewCreateTable().IfNotExists().
		Model((*types.Customer)(nil)).Exec(ctx)
	if err != nil {
		log.Fatalf("❌ Failed to create customers table: %v", err)
	}
	log.Println("✅ Customers table created successfully!")

	// FoodItems table
	_, err = db.NewCreateTable().IfNotExists().
		Model((*types.FoodItem)(nil)).
		ForeignKey(`("customer_id") REFERENCES "customers" ("id") ON DELETE CASCADE`).
		Exec(ctx)
	if err != nil {
		log.Fatalf("❌ Failed to create food_items table: %v", err)
	}
	log.Println("✅ FoodItems table created successfully!")

	// Orders table
	_, err = db.NewCreateTable().IfNotExists().
		Model((*types.Order)(nil)).
		ForeignKey(`("customer_id") REFERENCES "customers" ("id") ON DELETE CASCADE`).
		Exec(ctx)
	if err != nil {
		log.Fatalf("❌ Failed to create orders table: %v", err)
	}
	log.Println("✅ Orders table created successfully!")
}

func Drop(db *bun.DB) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*60))
	defer cancel()
	err := db.ResetModel(ctx, (*types.Customer)(nil), (*types.FoodItem)(nil), (*types.Order)(nil))
	if err != nil {
		log.Fatalf("❌ Failed to drop tables: %v", err)
	}
	log.Println("✅ Tables dropped successfully!")
}
