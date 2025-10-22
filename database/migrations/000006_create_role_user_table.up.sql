CREATE TABLE role_user (
  role_id CHAR(26) NOT NULL,
  user_id CHAR(26) NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITH TIME ZONE,
  
  PRIMARY KEY (role_id, user_id),
  CONSTRAINT fk_role_user_role
    FOREIGN KEY (role_id)
    REFERENCES roles(id)
    ON DELETE CASCADE,
  CONSTRAINT fk_role_user_user
    FOREIGN KEY (user_id)
    REFERENCES users(id)
    ON DELETE CASCADE
);

CREATE INDEX idx_role_user_role_id ON role_user(role_id);   
CREATE INDEX idx_role_user_user_id ON role_user(user_id);