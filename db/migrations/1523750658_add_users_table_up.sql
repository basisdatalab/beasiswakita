CREATE TABLE users (
	id int(11) NOT NULL AUTO_INCREMENT,
	email_address varchar(50) NOT NULL UNIQUE,
	password varchar(255) NOT NULL,
	role varchar(50) NOT NULL DEFAULT "",
	created_at datetime NOT NULL,
	updated_at datetime NOT NULL,
	PRIMARY KEY (id),
	INDEX email_INDEX (email_address ASC)
);
