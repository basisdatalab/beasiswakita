CREATE TABLE student_boards (
    id int(11) NOT NULL AUTO_INCREMENT,
    name varchar(100) NOT NULL,
    category int(2) NOT NULL,
    description longtext NOT NULL,
    state int(2) NOT NULL,
    user_id int(11) NOT NULL,
    created_at datetime NOT NULL,
	updated_at datetime NOT NULL,
    PRIMARY KEY(id),
    FOREIGN KEY (student_id) REFERENCES students (id)
);