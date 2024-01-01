-- select bill_details as json
SELECT
	t.id AS bill_id,
	to_char(t.bill_date, 'DD/MM/YYYY') AS bill_date,
	to_char(t.entry_date, 'DD/MM/YYYY') AS entry_date,
	to_char(t.finish_date, 'DD/MM/YYYY') AS finish_date,
	e.id AS employee_id,
	e.name AS employee_name,
	e.phone_number AS employee_phone_number,
	e.address AS employee_address,
	c.id AS customer_id,
	c.name AS customer_name,
	c.phone_number AS customer_phone_number,
	c.address AS customer_address,
	array_agg(
		jsonb_build_object(
			'id', bd.id,
			'product_id', bd.product_id,
			'product_name', p.name,
			'product_price', p.price,
			'product_unit', p.unit,
			'quantity', bd.quantity
		)
	) AS bill_details
FROM
	bill t
JOIN
	employee e ON t.employee_id = e.id
JOIN
	customer c ON t.customer_id = c.id
JOIN
	bill_detail bd ON t.id = bd.bill_id
JOIN
	product p ON bd.product_id = p.id
GROUP BY
	t.id, t.bill_date, t.entry_date, t.finish_date, e.id, e.name, e.phone_number, e.address, c.id, c.name, c.phone_number, c.address
ORDER BY
	t.id ASC;

-- normal select
SELECT t.id, to_char(t.bill_date, 'DD/MM/YYYY'), to_char(t.entry_date, 'DD/MM/YYYY'), to_char(t.finish_date, 'DD/MM/YYYY'),
      e.id, e.name, e.phone_number, e.address,
      c.id, c.name, c.phone_number, c.address,
      bd.id, bd.product_id, p.name, p.price, p.unit, bd.quantity
FROM bill t
JOIN employee e ON t.employee_id = e.id
JOIN customer c ON t.customer_id = c.id
JOIN bill_detail bd ON t.id = bd.bill_id
JOIN product p ON bd.product_id = p.id
ORDER BY t.id ASC;
