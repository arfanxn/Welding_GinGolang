CREATE TABLE material_test_services (
  id CHAR(26) PRIMARY KEY NOT NULL,
  machine_id CHAR(26) NOT NULL,
  method_id CHAR(26) NOT NULL,
  service_type VARCHAR(50) NOT NULL,
  service_code VARCHAR(50) NOT NULL,
  unit VARCHAR(50) NOT NULL,
  price DECIMAL(10, 2) NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITH TIME ZONE,
  deleted_at TIMESTAMP WITH TIME ZONE,

  CONSTRAINT fk_material_test_services_machine_id FOREIGN KEY (machine_id) REFERENCES material_test_machines(id) ON DELETE CASCADE,
  CONSTRAINT fk_material_test_services_method_id FOREIGN KEY (method_id) REFERENCES material_test_methods(id) ON DELETE CASCADE
);