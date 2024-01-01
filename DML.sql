INSERT INTO "public"."product" ("id", "name", "unit", "price") VALUES ('P_1', 'Cuci + Setrika', 'KG', 7000);
INSERT INTO "public"."product" ("id", "name", "unit", "price") VALUES ('P_2', 'Laundy Bedcover', 'Buah', 50000);
INSERT INTO "public"."product" ("id", "name", "unit", "price") VALUES ('P_3', 'Laundy Boneka', 'Buah', 25000);

INSERT INTO "public"."customer" ("id", "name", "phone_number", "address") VALUES ('C_1', 'Bjarn Stroustroup', '081223234222', 'Jakarta');
INSERT INTO "public"."customer" ("id", "name", "phone_number", "address") VALUES ('C_2', 'Jessica', '0812654987', 'Jakarta');

INSERT INTO "public"."employee" ("id", "name", "phone_number", "address") VALUES ('E_1', 'Mirna', '081213336899', 'Jakarta');

INSERT INTO "public"."bill" ("id", "bill_date", "entry_date", "finish_date", "employee_id", "customer_id") VALUES ('B_8', '2023-12-31', '2023-12-31', '2024-01-02', 'E_1', 'C_1');

INSERT INTO "public"."bill_detail" ("id", "bill_id", "product_id", "quantity") VALUES ('BD_4', 'B_8', 'P_1', 5);
INSERT INTO "public"."bill_detail" ("id", "bill_id", "product_id", "quantity") VALUES ('BD_5', 'B_8', 'P_2', 1);
INSERT INTO "public"."bill_detail" ("id", "bill_id", "product_id", "quantity") VALUES ('BD_6', 'B_8', 'P_3', 2);

INSERT INTO "public"."bill" ("id", "bill_date", "entry_date", "finish_date", "employee_id", "customer_id") VALUES ('B_9', '2023-12-31', '2023-12-31', '2024-01-02', 'E_1', 'C_1');

INSERT INTO "public"."bill_detail" ("id", "bill_id", "product_id", "quantity") VALUES ('BD_7', 'B_9', 'P_1', 5);
INSERT INTO "public"."bill_detail" ("id", "bill_id", "product_id", "quantity") VALUES ('BD_8', 'B_9', 'P_2', 1);
INSERT INTO "public"."bill_detail" ("id", "bill_id", "product_id", "quantity") VALUES ('BD_9', 'B_9', 'P_3', 2);

INSERT INTO "public"."bill" ("id", "bill_date", "entry_date", "finish_date", "employee_id", "customer_id") VALUES ('B_10', '2023-12-31', '2023-12-31', '2024-01-02', 'E_1', 'C_2');


INSERT INTO "public"."bill_detail" ("id", "bill_id", "product_id", "quantity") VALUES ('BD_10', 'B_10', 'P_1', 5);
