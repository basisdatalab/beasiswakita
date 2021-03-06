CREATE TABLE students (
	id int(11) NOT NULL AUTO_INCREMENT,
	name varchar(100) NOT NULL DEFAULT '',
	birthdate date NOT NULL DEFAULT '',
	address varchar(255) NOT NULL DEFAULT '',
	city varchar(50) NOT NULL DEFAULT '',
	region varchar(50) NOT NULL DEFAULT '',
	country varchar(50) NOT NULL DEFAULT '',
	zipcode varchar(50),
	education varchar(50) NOT NULL DEFAULT '',
	school_name varchar(100) NOT NULL DEFAULT '',
	school_address varchar(100) NOT NULL DEFAULT '',
	school_city varchar(50) NOT NULL DEFAULT '',
	school_region varchar(50) NOT NULL DEFAULT '',
	school_zipcode varchar(50),
	user_id int(11) NOT NULL,
	created_at datetime NOT NULL,
	updated_at datetime NOT NULL,
	PRIMARY KEY (id),
	FOREIGN KEY (user_id) REFERENCES users (id)
);
