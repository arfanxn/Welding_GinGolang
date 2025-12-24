CREATE TABLE material_test_order_user (
  order_id CHAR(26) NOT NULL,
  user_id CHAR(26) NOT NULL,
  type VARCHAR(50) NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITH TIME ZONE,
  
  PRIMARY KEY (order_id, user_id, type),
  CONSTRAINT fk_material_test_order_user_order
    FOREIGN KEY (order_id)
    REFERENCES material_test_orders(id)
    ON DELETE CASCADE,
  CONSTRAINT fk_material_test_order_user_user
    FOREIGN KEY (user_id)
    REFERENCES users(id)
    ON DELETE CASCADE
);

CREATE INDEX idx_material_test_order_user_order_id ON material_test_order_user(order_id);   
CREATE INDEX idx_material_test_order_user_user_id ON material_test_order_user(user_id);