CREATE TABLE students (
	id int(11) NOT NULL AUTO_INCREMENT,
	name varchar(100) NOT NULL,
	birthdate date NOT NULL,
	address varchar(255) NOT NULL,
	city varchar(50) NOT NULL,
	region varchar(50) NOT NULL,
	country varchar(50) NOT NULL,
	zipcode varchar(50),
	education varchar(50) NOT NULL,
	school_name varchar(100) NOT NULL,
	school_address varchar(100) NOT NULL,
	school_city varchar(50) NOT NULL,
	school_region varchar(50) NOT NULL,
	school_zipcode varchar(50),
	user_id int(11) NOT NULL,
	PRIMARY KEY (id),
	FOREIGN KEY (user_id) REFERENCES users (id)
);
