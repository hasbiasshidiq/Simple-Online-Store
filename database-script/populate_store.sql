-- insert into sellers table
INSERT INTO sellers
	(seller_id, seller_name) 
VALUES
	('barokah_store'		, 'Barokah Store'				),
	('legenda_store'		, 'Legenda Store'				),
	('armani_store'			, 'Armani Store'				);

-- insert into customers table
INSERT INTO customers
	(customer_id, first_name, last_name) 
VALUES
	('bambang_pamungkas'	, 'Bambang'		, 'Pamungkas'	),
	('isnan_ali'			, 'Isnan'		, 'Ali'			),
	('imral_usman'			, 'Imral'		, 'Usman'		),
	('budi_sudarsono'		, 'Budi'		, 'Sudarsono'	),
	('hendro_kartiko'		, 'Hendro'		, 'Kartiko'		),
	('wawan_widiantoro'		, 'Wawan'		, 'Widiantoro'	),
	('hari_salisburi'		, 'Hari'		, 'Salisburi'	),
	('susi_susanti'			, 'Susi'		, 'Susanti'		),
	('alan_budi_kusuma'		, 'Alan'		, 'Budi Kusuma'	),
	('taufik_hidayat'		, 'Taufik' 		, 'Hidayat'		);

-- insert into products table
INSERT INTO products
	(product_name, category, price) 
VALUES
	('Honda Beat'			,'motorcycle'	, 15630000		),
	('Tesla Model 3'		,'car'			, 420000000		),
	('Flexamove'			,'herbal'		, 120000		),
	('Antapro Barata'		,'herbal'		, 200000		),
	('Daihatsu Xenia'		,'car'			, 150000000		),
	('Senator Lycan'		,'bicycle'		, 950000		);

-- insert into invetory table
INSERT INTO inventory
	(seller_id, product_id, quantity) 
VALUES
	('legenda_store'		,	1			, 10			),
	('legenda_store'		,	2			, 0				),
	('barokah_store'		,	3			, 100			),
	('barokah_store'		,	4			, 75			),
	('legenda_store'		,	5			, 20			),
	('legenda_store'		,	6			, 5				);