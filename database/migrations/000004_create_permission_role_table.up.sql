CREATE TABLE permission_role (
  permission_id CHAR(26) NOT NULL,
  role_id CHAR(26) NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITH TIME ZONE,
  
  PRIMARY KEY (permission_id, role_id),
  CONSTRAINT fk_permission_role_permission
    FOREIGN KEY (permission_id)
    REFERENCES permissions(id)
    ON DELETE CASCADE,
  CONSTRAINT fk_permission_role_role
    FOREIGN KEY (role_id)
    REFERENCES roles(id)
    ON DELETE CASCADE
);

CREATE INDEX idx_permission_role_permission_id ON permission_role(permission_id);
CREATE INDEX idx_permission_role_role_id ON permission_role(role_id);