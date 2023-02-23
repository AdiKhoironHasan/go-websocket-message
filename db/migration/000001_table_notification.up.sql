CREATE TABLE users (
    id int NOT NULL AUTO_INCREMENT,
    FullName varchar(255) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE notifications (
    id int NOT NULL AUTO_INCREMENT,
	user_id int NOT NULL,
	is_read boolean NOT NULL DEFAULT 0,
	title varchar(255) NOT NULL,
    detail text NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES Users(id) ON DELETE CASCADE
);

