CREATE TABLE scholarships (
    id int(11) NOT NULL AUTO_INCREMENT,
    name varchar(100) NOT NULL DEFAULT '',
    country varchar(50) NOT NULL DEFAULT '',
    flag varchar(20) NOT NULL DEFAULT '',
    state int(2) NOT NULL,
    start_date datetime NOT NULL,
    end_date datetime NOT NULL,
    description varchar(255) NOT NULL DEFAULT '',
    requirement varchar(255) NOT NULL DEFAULT '',
    organization_id int(11) NOT NULL,
    created_at datetime NOT NULL,
    updated_at datetime NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (organization_id) REFERENCES organizations (id),
    INDEX organization_INDEX (organization_id ASC)
)
