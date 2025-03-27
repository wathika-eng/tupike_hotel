CREATE TABLE IF NOT EXISTS customers (
  id SERIAL PRIMARY KEY,
  username VARCHAR(255) NOT NULL,
  email VARCHAR(55) NOT NULL UNIQUE,
  password VARCHAR(255) NOT NULL,
  phone_number VARCHAR(20) NOT NULL UNIQUE,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  is_admin BOOLEAN NOT NULL DEFAULT TRUE
);

-- Add indexes for faster lookups
CREATE INDEX idx_customers_email ON customers(email);
CREATE INDEX idx_customers_phone ON customers(phone_number);

CREATE TABLE IF NOT EXISTS food (
  id SERIAL PRIMARY KEY,
  item VARCHAR(255) NOT NULL UNIQUE,
  description TEXT NOT NULL,
  image_url VARCHAR(255) NOT NULL UNIQUE,
  order_freq INT NOT NULL DEFAULT 0,
  price FLOAT NOT NULL
);

-- Index for faster searches by food item
CREATE INDEX idx_food_item ON food(item);

CREATE TABLE IF NOT EXISTS orders (
  id SERIAL PRIMARY KEY,
  customer_id INT NOT NULL,
  food_id INT NOT NULL,
  delivery_status VARCHAR(50) NOT NULL DEFAULT 'pending',
  payment_status BOOLEAN NOT NULL DEFAULT FALSE,
  ordered_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  amount_total INT NOT NULL,
  discount FLOAT NOT NULL DEFAULT 0,
  CONSTRAINT fk_customer FOREIGN KEY (customer_id) REFERENCES customers(id) ON DELETE CASCADE,
  CONSTRAINT fk_food FOREIGN KEY (food_id) REFERENCES food(id) ON DELETE CASCADE
);

-- Indexes for faster queries by customer and food
CREATE INDEX idx_orders_customer ON orders(customer_id);
CREATE INDEX idx_orders_food ON orders(food_id);
CREATE INDEX idx_orders_delivery_status ON orders(delivery_status);
