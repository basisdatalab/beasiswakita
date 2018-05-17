CREATE TABLE student_boards (
    id int(11) NOT NULL AUTO_INCREMENT,
    name varchar(100) NOT NULL DEFAULT '',
    category int(2) NOT NULL,
    description varchar(255) NOT NULL DEFAULT '',
    state int(2) NOT NULL,
    user_id int(11) NOT NULL,
    created_at datetime NOT NULL,
    updated_at datetime NOT NULL,
    PRIMARY KEY(id),
    FOREIGN KEY (user_id) REFERENCES users (id),
    INDEX user_INDEX (user_id ASC)
);
