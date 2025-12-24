CREATE TABLE material_test_services (
  id CHAR(26) PRIMARY KEY NOT NULL,
  machine_id CHAR(26) NOT NULL,
  method_id CHAR(26) NOT NULL,
  test_name VARCHAR(255) NOT NULL UNIQUE,
  service_type VARCHAR(50) NOT NULL,
  service_code VARCHAR(50) NOT NULL UNIQUE,
  unit VARCHAR(50) NOT NULL,
  price DECIMAL(10, 2) NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITH TIME ZONE,
  deleted_at TIMESTAMP WITH TIME ZONE,

  CONSTRAINT fk_material_test_services_machine_id FOREIGN KEY (machine_id) REFERENCES material_test_machines(id) ON DELETE CASCADE,
  CONSTRAINT fk_material_test_services_method_id FOREIGN KEY (method_id) REFERENCES material_test_methods(id) ON DELETE CASCADE

);

--   CREATE INDEX idx_mtos_latest_string_based
-- ON material_test_order_services (
--     service_id,
--     (SPLIT_PART(SPLIT_PART(sample_number, '/', 2), '.', 1)::int),
--     (RIGHT(sample_number, 3)::int)
-- );