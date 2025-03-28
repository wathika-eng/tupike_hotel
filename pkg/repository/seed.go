package repository

import (
	"context"
	"log"
	"math/rand"
	"time"
	"tupike_hotel/pkg/types"

	"github.com/google/uuid"
)

func (r *FoodRepo) SeedFoodItems() {
	foodItems := []types.FoodItem{
		{ID: uuid.New(), Item: "Burger", Description: "Cheese Burger with lettuce",
			ImageURL: "https://via.placeholder.com/151", Quantity: 20, Price: 5.99},
		{ID: uuid.New(), Item: "Pizza", Description: "Pepperoni Pizza with extra cheese",
			ImageURL: "https://via.placeholder.com/152", Quantity: 15, Price: 8.99},
		{ID: uuid.New(), Item: "Pasta", Description: "Creamy Alfredo Pasta",
			ImageURL: "https://via.placeholder.com/153", Quantity: 12, Price: 7.49},
		{ID: uuid.New(), Item: "Sushi", Description: "Fresh salmon sushi rolls",
			ImageURL: "https://via.placeholder.com/154", Quantity: 10, Price: 12.99},
		{ID: uuid.New(), Item: "Salad", Description: "Healthy green salad",
			ImageURL: "https://via.placeholder.com/155", Quantity: 18, Price: 4.99},
		{ID: uuid.New(), Item: "Steak", Description: "Grilled beef steak",
			ImageURL: "https://via.placeholder.com/156", Quantity: 8, Price: 15.99},
		{ID: uuid.New(), Item: "Fries", Description: "Crispy French fries",
			ImageURL: "https://via.placeholder.com/157", Quantity: 25, Price: 2.99},
		{ID: uuid.New(), Item: "Tacos", Description: "Mexican beef tacos",
			ImageURL: "https://via.placeholder.com/158", Quantity: 14, Price: 6.99},
		{ID: uuid.New(), Item: "Soup", Description: "Hot chicken soup",
			ImageURL: "https://via.placeholder.com/159", Quantity: 9, Price: 3.99},
		{ID: uuid.New(), Item: "Ice Cream", Description: "Vanilla ice cream scoop",
			ImageURL: "https://via.placeholder.com/150", Quantity: 22, Price: 2.49},
	}

	ctx := context.Background()
	for _, food := range foodItems {
		food.CreatedAt = time.Now()
		food.OrderFreq = rand.Intn(10)

		_, err := r.db.DB.NewInsert().Model(&food).Exec(ctx)
		if err != nil {
			log.Printf("Failed to insert food item: %v", err)
		}
	}

	log.Println("✅ Successfully inserted 10 food items!")
}
