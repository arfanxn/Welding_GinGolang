CREATE TABLE employees (
  user_id CHAR(26) PRIMARY KEY NOT NULL,
  employment_identity_number VARCHAR(50) NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITH TIME ZONE,
  
  CONSTRAINT uc_employment_identity_number UNIQUE (employment_identity_number),
  CONSTRAINT fk_employees_user_id
    FOREIGN KEY (user_id)
    REFERENCES users(id)
    ON DELETE CASCADE
);