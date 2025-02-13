-- Create a table to store warehouse information
CREATE TABLE warehouses (
  id INT AUTO_INCREMENT PRIMARY KEY,
  address VARCHAR(255),
  telephone VARCHAR(255),
  warehouse_code VARCHAR(255) UNIQUE,
  minimum_capacity INT,
  minimum_temperature FLOAT
);

-- Create a table to store employee card information
CREATE TABLE buyers (
  id INT AUTO_INCREMENT PRIMARY KEY,
  card_number_id VARCHAR(255) UNIQUE,
  first_name VARCHAR(255),
  last_name VARCHAR(255)
);

-- Create a table to store employee information
CREATE TABLE employees (
  id INT AUTO_INCREMENT PRIMARY KEY,
  card_number_id VARCHAR(255) UNIQUE,
  first_name VARCHAR(255),
  last_name VARCHAR(255),
  warehouse_id INT,
  FOREIGN KEY (warehouse_id) REFERENCES warehouses(id)
);

-- Create a table to store section information
CREATE TABLE sections (
  id INT AUTO_INCREMENT PRIMARY KEY,
  section_number VARCHAR(255) UNIQUE,
  current_temperature DECIMAL(19,2),
  minimum_temperature DECIMAL(19,2),
  current_capacity INT,
  minimum_capacity INT,
  maximum_capacity INT,
  warehouse_id INT,
  product_type_id INT,
  FOREIGN KEY (warehouse_id) REFERENCES warehouses(id)
);

-- Create a table to store locality information
CREATE TABLE localities (
  id INT AUTO_INCREMENT PRIMARY KEY,
  locality_name VARCHAR(255),
  province_name VARCHAR(255),
  country_name VARCHAR(255),
  UNIQUE (locality_name, province_name, country_name)
);

-- Create a table to store seller information
CREATE TABLE sellers (
  id INT AUTO_INCREMENT PRIMARY KEY,
  cid INT UNIQUE,
  company_name VARCHAR(255),
  address VARCHAR(255),
  telephone VARCHAR(255),
  locality_id INT,
  FOREIGN KEY (locality_id) REFERENCES localities(id)
);

-- Create a table to store products information
CREATE TABLE products (
  id INT AUTO_INCREMENT PRIMARY KEY,
  product_code VARCHAR(255) NOT NULL UNIQUE,
  description VARCHAR(255),
  height DECIMAL(19,2),
  length DECIMAL(19,2),
  width DECIMAL(19,2),
  net_weight DECIMAL(19,2),
  expiration_rate DECIMAL(19,2),
  freezing_rate DECIMAL(19,2),
  recommended_freezing_temperature DECIMAL(19,2),
  seller_id INT,
  product_type_id INT,
  FOREIGN KEY (seller_id) REFERENCES sellers(id)
  -- FOREIGN KEY (product_type_id) REFERENCES product_types(id)
);

-- Create a table to store carrier information
CREATE TABLE carries (
  id INT AUTO_INCREMENT PRIMARY KEY,
  cid VARCHAR(255) UNIQUE,
  company_name VARCHAR(255),
  address VARCHAR(255),
  phone_number VARCHAR(255),
  locality_id INT,
  FOREIGN KEY (locality_id) REFERENCES localities(id)
);

-- Create a table to store product batch information
CREATE TABLE product_batches (
  id INT AUTO_INCREMENT PRIMARY KEY,
  batch_number VARCHAR(255) UNIQUE,
  initial_quantity INT,
  current_quantity INT,
  current_temperature DECIMAL(19,2),
  minimum_temperature DECIMAL(19,2),
  due_date DATE,
  manufacturing_date DATE,
  manufacturing_hour TIME,
  product_id INT,
  section_id INT,
  FOREIGN KEY (product_id) REFERENCES products(id),
  FOREIGN KEY (section_id) REFERENCES sections(id)
);

-- Create a table to store product record information
CREATE TABLE product_records (
  id INT AUTO_INCREMENT PRIMARY KEY,
  last_update_date DATETIME(6),
  purchase_price DECIMAL(19, 2),
  sale_price DECIMAL(19, 2),
  product_id INT,
  FOREIGN KEY (product_id) REFERENCES products(id)
);

-- Create a table to store inbound order information
CREATE TABLE inbound_orders (
  id INT AUTO_INCREMENT PRIMARY KEY,
  order_date DATE,
  order_number INT UNIQUE,
  employee_id INT,
  product_batch_id INT,
  warehouse_id INT,
  FOREIGN KEY (employee_id) REFERENCES employees(id),
  FOREIGN KEY (product_batch_id) REFERENCES product_batches(id),
  FOREIGN KEY (warehouse_id) REFERENCES warehouses(id)
);

-- Create a table to store purchase order information
CREATE TABLE purchase_orders (
  id INT AUTO_INCREMENT PRIMARY KEY,
  order_number VARCHAR(255) UNIQUE,
  order_date DATE,
  tracking_code VARCHAR(255),
  buyer_id INT,
  product_record_id INT,
  FOREIGN KEY (buyer_id) REFERENCES buyers(id),
  FOREIGN KEY (product_record_id) REFERENCES product_records(id)
);
