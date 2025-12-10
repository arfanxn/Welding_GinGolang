CREATE TABLE material_test_machines (
  id CHAR(26) PRIMARY KEY NOT NULL,
  name VARCHAR(50) NOT NULL,
  description TEXT,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITH TIME ZONE,
  deleted_at TIMESTAMP WITH TIME ZONE,

  CONSTRAINT uc_material_test_machines_name UNIQUE (name)
);