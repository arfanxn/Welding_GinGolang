CREATE TABLE users (
  id CHAR(26) PRIMARY KEY NOT NULL,
  name VARCHAR(50) NOT NULL,
  email VARCHAR(50) NOT NULL,
  email_verified_at TIMESTAMP WITH TIME ZONE,
  phone_number VARCHAR(25),
  password VARCHAR(255) NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITH TIME ZONE,
  deleted_at TIMESTAMP WITH TIME ZONE,
  
  CONSTRAINT uc_users_email UNIQUE (email),
  CONSTRAINT uc_users_phone_number UNIQUE (phone_number)
);