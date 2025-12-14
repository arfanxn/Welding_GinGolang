CREATE TABLE material_test_order_services (
  id CHAR(26) PRIMARY KEY NOT NULL,
  order_id CHAR(26) NOT NULL,
  service_id CHAR(26) NOT NULL,
  sample_number VARCHAR(50) NOT NULL UNIQUE,
  sample_name VARCHAR(255) NOT NULL, -- the name of the material
  price DECIMAL(10, 2) NOT NULL, -- saved current price of the material test service
  quantity INT NOT NULL, -- the quantity of the material
  line_total DECIMAL(10, 2) NOT NULL, -- the line total of the material test order service (price * quantity)
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITH TIME ZONE,

  CONSTRAINT fk_material_test_order_services_order_id FOREIGN KEY (order_id) REFERENCES material_test_orders(id) ON DELETE CASCADE,
  CONSTRAINT fk_material_test_order_services_service_id FOREIGN KEY (service_id) REFERENCES material_test_services(id) ON DELETE CASCADE
);