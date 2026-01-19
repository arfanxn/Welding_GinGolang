CREATE TABLE material_test_order_service_evaluations (
  id CHAR(26) PRIMARY KEY NOT NULL,
  order_service_id CHAR(26) NOT NULL,
  is_equipment_available BOOLEAN NOT NULL DEFAULT FALSE,
  is_personnel_available BOOLEAN NOT NULL DEFAULT FALSE,
  is_time_available BOOLEAN NOT NULL DEFAULT FALSE,
  is_test_ready BOOLEAN NOT NULL DEFAULT FALSE,
  is_subcontract_lab_available BOOLEAN NOT NULL DEFAULT FALSE,

  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITH TIME ZONE,

  CONSTRAINT fk_material_test_order_service_evaluations_order_service_id FOREIGN KEY (order_service_id) REFERENCES material_test_order_services(id) ON DELETE CASCADE
);