CREATE SEQUENCE customer_seq START WITH 1;
CREATE SEQUENCE product_seq START WITH 1;
CREATE SEQUENCE employee_seq START WITH 1;
CREATE SEQUENCE bill_seq START WITH 1;
CREATE SEQUENCE bill_detail_seq START WITH 1;

CREATE TABLE customer ( 
  id VARCHAR(100) NOT NULL DEFAULT ('C_' || nextval('customer_seq')),
  name VARCHAR(100) NOT NULL,
  phone_number VARCHAR(15) NOT  NULL,
  address VARCHAR(255) NOT NULL,
  CONSTRAINT "customer_pkey" PRIMARY KEY ("id")
);

CREATE TABLE employee ( 
  id VARCHAR(100) NOT NULL DEFAULT ('E_' || nextval('employee_seq')),
  name VARCHAR(100) NOT NULL,
  phone_number VARCHAR(15) NOT  NULL,
  address VARCHAR(255) NOT NULL,
  CONSTRAINT "employee_pkey" PRIMARY KEY ("id")
);

CREATE TABLE product ( 
  id VARCHAR(100) NOT NULL DEFAULT ('P_' || nextval('product_seq')),
  name VARCHAR(100) NOT NULL,
  price INTEGER NOT NULL,
  CONSTRAINT "product_pkey" PRIMARY KEY ("id")
);

CREATE TABLE bill ( 
  id VARCHAR(100) NOT NULL DEFAULT ('B_' || nextval('bill_seq')),
  bill_date DATE NOT NULL,
  entry_date DATE NOT NULL,
  finish_date DATE NOT NULL,
  employee_id VARCHAR(100) NOT NULL,
  customer_id VARCHAR(100) NOT NULL,
  CONSTRAINT "bill_pkey" PRIMARY KEY ("id")
);

CREATE TABLE bill_detail ( 
  id VARCHAR(100) NOT NULL DEFAULT ('BD_' || nextval('bill_detail_seq')),
  bill_id VARCHAR(100) NOT NULL,
  product_id VARCHAR(100) NOT NULL,
  quantity INTEGER NOT NULL,
  CONSTRAINT "bill_detail_pkey" PRIMARY KEY ("id")
);

ALTER TABLE bill ADD CONSTRAINT "bill_employee_id_fkey" FOREIGN KEY ("employee_id") REFERENCES employee ("id");
ALTER TABLE bill ADD CONSTRAINT "bill_customer_id_fkey" FOREIGN KEY ("customer_id") REFERENCES customer ("id");

ALTER TABLE bill_detail ADD CONSTRAINT "bill_detail_bill_id_fkey" FOREIGN KEY ("bill_id") REFERENCES bill ("id");
ALTER TABLE bill_detail ADD CONSTRAINT "bill_detail_product_id_fkey" FOREIGN KEY ("product_id") REFERENCES product ("id");
