CREATE TABLE organizations (
	id int(11) NOT NULL AUTO_INCREMENT,
	name varchar(100) NOT NULL,
	position varchar(50) NOT NULL,
	organization_name varchar(100) NOT NULL,
	organization_email varchar(50) NOT NULL,
	address varchar(255) NOT NULL,
	city varchar(50) NOT NULL,
	region varchar(50) NOT NULL,
	country varchar(50) NOT NULL,
	zipcode varchar(50),
	website varchar(100),
	user_id int(11) NOT NULL,
	created_at datetime NOT NULL,
	updated_at datetime NOT NULL,
	PRIMARY KEY (id),
	FOREIGN KEY (user_id) REFERENCES users (id)
)
