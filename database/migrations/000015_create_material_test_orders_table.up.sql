CREATE TABLE material_test_orders (
  id CHAR(26) PRIMARY KEY NOT NULL,
  number VARCHAR(50) NOT NULL,
  work_category_id CHAR(26) NOT NULL,
  work_package_id CHAR(26) NOT NULL,
  applicant_name VARCHAR(255) NOT NULL,
  applicant_phone_number VARCHAR(25) NOT NULL,
  applicant_email VARCHAR(50) NOT NULL,
  applicant_full_address VARCHAR(512) NOT NULL,
  applicant_note TEXT,
  recipent_name VARCHAR(255) NOT NULL,
  tester_note TEXT,
  
  sub_total DECIMAL(10, 2) NOT NULL,
  tax DECIMAL(10, 2) NOT NULL,
  discount DECIMAL(10, 2) NOT NULL,
  total DECIMAL(10, 2) NOT NULL,

  entered_at TIMESTAMP WITH TIME ZONE NOT NULL,

  status VARCHAR(50) NOT NULL,

  payment_submitted_at TIMESTAMP WITH TIME ZONE,
  payment_rejected_at TIMESTAMP WITH TIME ZONE,
  payment_approved_at TIMESTAMP WITH TIME ZONE,

  testing_at TIMESTAMP WITH TIME ZONE,
  completed_at TIMESTAMP WITH TIME ZONE,

  cancelled_at TIMESTAMP WITH TIME ZONE,
  rejected_at TIMESTAMP WITH TIME ZONE,
  refunded_at TIMESTAMP WITH TIME ZONE,
  
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITH TIME ZONE,

  CONSTRAINT fk_material_test_orders_work_category_id FOREIGN KEY (work_category_id) REFERENCES material_test_work_categories(id) ON DELETE CASCADE,
  CONSTRAINT fk_material_test_orders_work_package_id FOREIGN KEY (work_package_id) REFERENCES material_test_work_packages(id) ON DELETE CASCADE,
  CONSTRAINT fk_material_test_orders_customer_id FOREIGN KEY (customer_id) REFERENCES customers(id) ON DELETE CASCADE
);