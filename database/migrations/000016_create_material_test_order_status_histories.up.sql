-- CREATE TABLE material_test_orders (
--   id CHAR(26) PRIMARY KEY NOT NULL,
--   order_id CHAR(26) NOT NULL,
--   status VARCHAR(50) NOT NULL,

--   payment_submitted_at TIMESTAMP WITH TIME ZONE,
--   payment_rejected_at TIMESTAMP WITH TIME ZONE,
--   payment_approved_at TIMESTAMP WITH TIME ZONE,

--   testing_at TIMESTAMP WITH TIME ZONE,
--   completed_at TIMESTAMP WITH TIME ZONE,

--   cancelled_at TIMESTAMP WITH TIME ZONE,
--   rejected_at TIMESTAMP WITH TIME ZONE,
--   refunded_at TIMESTAMP WITH TIME ZONE,
  
--   created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
--   updated_at TIMESTAMP WITH TIME ZONE,

--   CONSTRAINT fk_material_test_orders_work_category_id FOREIGN KEY (work_category_id) REFERENCES material_test_work_categories(id) ON DELETE CASCADE
-- );