-- Create client table
CREATE TABLE IF NOT EXISTS client (
  id SERIAL PRIMARY KEY,
  identity_number VARCHAR,
  first_name VARCHAR,
  last_name VARCHAR,
  email VARCHAR,
  date_of_birth TIMESTAMP,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create product table
CREATE TABLE IF NOT EXISTS product (
  id SERIAL PRIMARY KEY,
  name VARCHAR,
  sku VARCHAR,
  category VARCHAR,
  price FLOAT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create sale table
CREATE TABLE IF NOT EXISTS sale (
  id SERIAL PRIMARY KEY,
  client_id INTEGER,
  shipped BOOLEAN,
  products JSONB,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (client_id) REFERENCES client(id) -- Reference to the client table
);