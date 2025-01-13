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
    batch_number VARCHAR(255),
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
    FOREIGN KEY (employee_id) REFERENCES employees(id),
    FOREIGN KEY (product_batch_id) REFERENCES product_batches(id)
);

-- Create a table to store purchase order information
CREATE TABLE purchase_orders (
    id INT AUTO_INCREMENT PRIMARY KEY,
    order_number VARCHAR(255),
    order_date DATE,
    tracking_code VARCHAR(255),
    buyer_id INT,
    product_record_id INT,
    FOREIGN KEY (buyer_id) REFERENCES buyers(id),
    FOREIGN KEY (product_record_id) REFERENCES product_records(id)
);

-- Insert data into warehouses
INSERT INTO warehouses (address, telephone, warehouse_code, minimum_capacity, minimum_temperature) VALUES ('123 Main St', '555-1234', 'WH001', 100, -10.0);
INSERT INTO warehouses (address, telephone, warehouse_code, minimum_capacity, minimum_temperature) VALUES ('456 Elm St', '555-5678', 'WH002', 200, -20.0);
INSERT INTO warehouses (address, telephone, warehouse_code, minimum_capacity, minimum_temperature) VALUES ('789 Oak St', '555-9012', 'WH003', 300, -30.0);
INSERT INTO warehouses (address, telephone, warehouse_code, minimum_capacity, minimum_temperature) VALUES ('101 Pine St', '555-3456', 'WH004', 400, -40.0);
INSERT INTO warehouses (address, telephone, warehouse_code, minimum_capacity, minimum_temperature) VALUES ('202 Maple St', '555-7890', 'WH005', 500, -50.0);

-- Insert data into buyers
INSERT INTO buyers (card_number_id, first_name, last_name) VALUES ('CARD001', 'John', 'Doe');
INSERT INTO buyers (card_number_id, first_name, last_name) VALUES ('CARD002', 'Jane', 'Smith');
INSERT INTO buyers (card_number_id, first_name, last_name) VALUES ('CARD003', 'Alice', 'Johnson');
INSERT INTO buyers (card_number_id, first_name, last_name) VALUES ('CARD004', 'Bob', 'Brown');
INSERT INTO buyers (card_number_id, first_name, last_name) VALUES ('CARD005', 'Charlie', 'Davis');

-- Insert data into employees
INSERT INTO employees (card_number_id, first_name, last_name, warehouse_id) VALUES ('EMP001', 'John', 'Doe', 1);
INSERT INTO employees (card_number_id, first_name, last_name, warehouse_id) VALUES ('EMP002', 'Jane', 'Smith', 2);
INSERT INTO employees (card_number_id, first_name, last_name, warehouse_id) VALUES ('EMP003', 'Alice', 'Johnson', 3);
INSERT INTO employees (card_number_id, first_name, last_name, warehouse_id) VALUES ('EMP004', 'Bob', 'Brown', 4);
INSERT INTO employees (card_number_id, first_name, last_name, warehouse_id) VALUES ('EMP005', 'Charlie', 'Davis', 5);

-- Insert data into sections
INSERT INTO sections (section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, maximum_capacity, warehouse_id, product_type_id) VALUES ('SC-001', -10.0, -20.0, 50, 10, 100, 1, 1);
INSERT INTO sections (section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, maximum_capacity, warehouse_id, product_type_id) VALUES ('SC-002', -15.0, -25.0, 60, 20, 200, 2, 2);
INSERT INTO sections (section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, maximum_capacity, warehouse_id, product_type_id) VALUES ('SC-003', -20.0, -30.0, 70, 30, 300, 3, 3);
INSERT INTO sections (section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, maximum_capacity, warehouse_id, product_type_id) VALUES ('SC-004', -25.0, -35.0, 80, 40, 400, 4, 4);
INSERT INTO sections (section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, maximum_capacity, warehouse_id, product_type_id) VALUES ('SC-005', -30.0, -40.0, 90, 50, 500, 5, 5);

-- Insert data into localities
INSERT INTO localities (locality_name, province_name, country_name) VALUES ('Locality1', 'Province1', 'Country1');
INSERT INTO localities (locality_name, province_name, country_name) VALUES ('Locality2', 'Province2', 'Country2');
INSERT INTO localities (locality_name, province_name, country_name) VALUES ('Locality3', 'Province3', 'Country3');
INSERT INTO localities (locality_name, province_name, country_name) VALUES ('Locality4', 'Province4', 'Country4');
INSERT INTO localities (locality_name, province_name, country_name) VALUES ('Locality5', 'Province5', 'Country5');

-- Insert data into sellers
INSERT INTO sellers (cid, company_name, address, telephone, locality_id) VALUES (1, 'Company1', '123 Main St', '555-1234', 1);
INSERT INTO sellers (cid, company_name, address, telephone, locality_id) VALUES (2, 'Company2', '456 Elm St', '555-5678', 2);
INSERT INTO sellers (cid, company_name, address, telephone, locality_id) VALUES (3, 'Company3', '789 Oak St', '555-9012', 3);
INSERT INTO sellers (cid, company_name, address, telephone, locality_id) VALUES (4, 'Company4', '101 Pine St', '555-3456', 4);
INSERT INTO sellers (cid, company_name, address, telephone, locality_id) VALUES (5, 'Company5', '202 Maple St', '555-7890', 5);

-- Insert data into products
INSERT INTO products (product_code, description, height, length, width, net_weight, expiration_rate, freezing_rate, recommended_freezing_temperature, product_type_id, seller_id) VALUES ('P001', 'Product 1', 10.0, 20.0, 30.0, 40.0, 0.1, 0.2, -10.0, 1, 1);
INSERT INTO products (product_code, description, height, length, width, net_weight, expiration_rate, freezing_rate, recommended_freezing_temperature, product_type_id, seller_id) VALUES ('P002', 'Product 2', 15.0, 25.0, 35.0, 45.0, 0.2, 0.3, -15.0, 2, 2);
INSERT INTO products (product_code, description, height, length, width, net_weight, expiration_rate, freezing_rate, recommended_freezing_temperature, product_type_id, seller_id) VALUES ('P003', 'Product 3', 20.0, 30.0, 40.0, 50.0, 0.3, 0.4, -20.0, 3, 3);
INSERT INTO products (product_code, description, height, length, width, net_weight, expiration_rate, freezing_rate, recommended_freezing_temperature, product_type_id, seller_id) VALUES ('P004', 'Product 4', 25.0, 35.0, 45.0, 55.0, 0.4, 0.5, -25.0, 4, 4);
INSERT INTO products (product_code, description, height, length, width, net_weight, expiration_rate, freezing_rate, recommended_freezing_temperature, product_type_id, seller_id) VALUES ('P005', 'Product 5', 30.0, 40.0, 50.0, 60.0, 0.5, 0.6, -30.0, 5, 5);

-- Insert data into carries
INSERT INTO carries (cid, company_name, address, phone_number, locality_id) VALUES (1, 'Carrier1', '123 Main St', '555-1234', 1);
INSERT INTO carries (cid, company_name, address, phone_number, locality_id) VALUES (2, 'Carrier2', '456 Elm St', '555-5678', 2);
INSERT INTO carries (cid, company_name, address, phone_number, locality_id) VALUES (3, 'Carrier3', '789 Oak St', '555-9012', 3);
INSERT INTO carries (cid, company_name, address, phone_number, locality_id) VALUES (4, 'Carrier4', '101 Pine St', '555-3456', 4);
INSERT INTO carries (cid, company_name, address, phone_number, locality_id) VALUES (5, 'Carrier5', '202 Maple St', '555-7890', 5);

-- Insert data into product_batches
INSERT INTO product_batches (batch_number, initial_quantity, current_quantity, current_temperature, minimum_temperature, due_date, manufacturing_date, manufacturing_hour, product_id, section_id) VALUES ('1', 200, 100, -10, -20, '2023-12-31', '2023-01-01', '08:00:00', 1, 1);
INSERT INTO product_batches (batch_number, initial_quantity, current_quantity, current_temperature, minimum_temperature, due_date, manufacturing_date, manufacturing_hour, product_id, section_id) VALUES ('2', 250, 150, -15, -25, '2023-11-30', '2023-02-01', '09:00:00', 2, 2);
INSERT INTO product_batches (batch_number, initial_quantity, current_quantity, current_temperature, minimum_temperature, due_date, manufacturing_date, manufacturing_hour, product_id, section_id) VALUES ('3', 300, 200, -20, -30, '2023-10-31', '2023-03-01', '10:00:00', 3, 3);
INSERT INTO product_batches (batch_number, initial_quantity, current_quantity, current_temperature, minimum_temperature, due_date, manufacturing_date, manufacturing_hour, product_id, section_id) VALUES ('4', 350, 250, -25, -35, '2023-09-30', '2023-04-01', '11:00:00', 4, 4);
INSERT INTO product_batches (batch_number, initial_quantity, current_quantity, current_temperature, minimum_temperature, due_date, manufacturing_date, manufacturing_hour, product_id, section_id) VALUES ('5', 400, 300, -30, -40, '2023-08-31', '2023-05-01', '12:00:00', 5, 5);

-- Insert data into product_records
INSERT INTO product_records (last_update_date, purchase_price, sale_price, product_id) VALUES ('2023-01-01', 10.00, 20.00, 1);
INSERT INTO product_records (last_update_date, purchase_price, sale_price, product_id) VALUES ('2023-02-01', 15.00, 25.00, 2);
INSERT INTO product_records (last_update_date, purchase_price, sale_price, product_id) VALUES ('2023-03-01', 20.00, 30.00, 3);
INSERT INTO product_records (last_update_date, purchase_price, sale_price, product_id) VALUES ('2023-04-01', 25.00, 35.00, 4);
INSERT INTO product_records (last_update_date, purchase_price, sale_price, product_id) VALUES ('2023-05-01', 30.00, 40.00, 5);

-- Insert data into inbound_orders
INSERT INTO inbound_orders (order_date, order_number, employee_id, product_batch_id) VALUES ('2023-01-01', 1001, 1, 1);
INSERT INTO inbound_orders (order_date, order_number, employee_id, product_batch_id) VALUES ('2023-02-01', 1002, 2, 2);
INSERT INTO inbound_orders (order_date, order_number, employee_id, product_batch_id) VALUES ('2023-03-01', 1003, 3, 3);
INSERT INTO inbound_orders (order_date, order_number, employee_id, product_batch_id) VALUES ('2023-04-01', 1004, 4, 4);
INSERT INTO inbound_orders (order_date, order_number, employee_id, product_batch_id) VALUES ('2023-05-01', 1005, 5, 5);

-- Insert data into purchase_orders
INSERT INTO purchase_orders (order_number, order_date, tracking_code, buyer_id, product_record_id) VALUES (2001, '2023-01-01', 3001, 1, 1);
INSERT INTO purchase_orders (order_number, order_date, tracking_code, buyer_id, product_record_id) VALUES (2002, '2023-02-01', 3002, 2, 2);
INSERT INTO purchase_orders (order_number, order_date, tracking_code, buyer_id, product_record_id) VALUES (2003, '2023-03-01', 3003, 3, 3);
INSERT INTO purchase_orders (order_number, order_date, tracking_code, buyer_id, product_record_id) VALUES (2004, '2023-04-01', 3004, 4, 4);
INSERT INTO purchase_orders (order_number, order_date, tracking_code, buyer_id, product_record_id) VALUES (2005, '2023-05-01', 3005, 5, 5);