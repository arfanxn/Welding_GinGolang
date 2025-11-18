CREATE TYPE code_type_enum AS ENUM ('user_register_invitation', 'user_email_verification', 'user_reset_password');

CREATE TABLE codes (
  id CHAR(26) PRIMARY KEY NOT NULL,
  codeable_id VARCHAR(50),
  codeable_type VARCHAR(50),
  type code_type_enum NOT NULL,
  value CHAR(6) NOT NULL,
  meta JSONB,
  used_at TIMESTAMP WITH TIME ZONE,
  expired_at TIMESTAMP WITH TIME ZONE NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITH TIME ZONE,
  
  CONSTRAINT codes_type_check CHECK (type IN ('user_register_invitation', 'user_email_verification', 'user_reset_password')),
  CONSTRAINT codes_value_check CHECK (value ~ '^[0-9]{6}$'),
  CONSTRAINT codes_used_at_check CHECK (used_at <= expired_at),
  CONSTRAINT codes_expired_at_check CHECK (expired_at > CURRENT_TIMESTAMP),
  CONSTRAINT codes_type_value_unique UNIQUE (type, value)
);

-- Add indexes separately to avoid syntax issues
CREATE INDEX codes_codeable_id_index ON codes (codeable_id);
CREATE INDEX codes_type_index ON codes (type);
CREATE INDEX codes_used_at_index ON codes (used_at);
CREATE INDEX codes_expired_at_index ON codes (expired_at);
