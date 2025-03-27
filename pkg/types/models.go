package types

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Customer struct {
	bun.BaseModel `bun:"table:customers"`
	ID            uuid.UUID `json:"id" bun:",pk,type:uuid,default:gen_random_uuid()"`
	//UserName      string    `json:"username" validate:"required" bun:"user_name,notnull"`
	Email       string    `json:"email" validate:"required,email" bun:"email,unique,notnull"`
	PhoneNumber string    `json:"phone_number" validate:"required,e164" bun:"phone_number,unique,notnull"`
	OTP         string    `json:"otp" bun:"otp"`
	IsAdmin     bool      `json:"is_admin" bun:"is_admin,notnull,default:false"`
	Password    string    `json:"password" validate:"required,min=8" bun:"password,notnull"`
	Verified    bool      `json:"verified" bun:"verified,default:false"`
	CreatedAt   time.Time `json:"created_at" bun:"created_at,notnull,default:current_timestamp"`
	LastLogin   time.Time `json:"last_login" bun:"last_login,notnull,default:current_timestamp"`
	Orders      []Order   `json:"orders" bun:"rel:has-many,join:id=customer_id"`
}

type FoodItem struct {
	bun.BaseModel `bun:"table:food"`
	ID            uuid.UUID `json:"id" bun:",pk,type:uuid,default:gen_random_uuid()"`
	Item          string    `json:"item" validate:"required" bun:"item,notnull,unique"`
	Description   string    `json:"description" validate:"required" bun:"description,notnull"`
	ImageURL      string    `json:"image_url" validate:"required,url" bun:"image_url,notnull,unique"`
	OrderFreq     int       `json:"order_freq" bun:"order_freq,notnull,default:0"`
	Quantity      int       `json:"quantity" validate:"required" bun:"quantity,notnull,default:0"`
	Price         float64   `json:"price" validate:"required,gt=0" bun:"price,notnull"`
	CreatedAt     time.Time `json:"created_at" bun:"created_at,notnull,default:current_timestamp"`
	Orders        []Order   `bun:"rel:has-many,join:id=food_id"`
}

type Order struct {
	bun.BaseModel  `bun:"table:orders"`
	ID             uuid.UUID `json:"id" bun:",pk,type:uuid,default:gen_random_uuid()"`
	CustomerID     uuid.UUID `json:"customer_id" validate:"required" bun:"customer_id,notnull"`
	FoodID         uuid.UUID `json:"food_id" validate:"required" bun:"food_id,notnull"`
	DeliveryStatus string    `json:"delivery_status" validate:"required" bun:"delivery_status,notnull,default:'pending'"`
	PaymentStatus  bool      `json:"payment_status" bun:"payment_status,notnull,default:false"`
	OrderedAt      time.Time `json:"ordered_at" bun:"ordered_at,notnull,default:current_timestamp"`
	Quantity       int       `json:"order_quantity" validate:"required" bun:"order_quantity,default:1"`
	AmountTotal    float64   `json:"amount_total" validate:"required,gt=0" bun:"amount_total,notnull"`
	Discount       float64   `json:"discount" bun:"discount,notnull,default:0"`
	Customer       *Customer `bun:"rel:belongs-to,join:customer_id=id"`
	FoodItem       *FoodItem `bun:"rel:belongs-to,join:food_id=id"`
}
