CREATE TABLE demo (
	id INTEGER NOT NULL PRIMARY KEY,
    device INTEGER,
    value INTEGER
);

CREATE TABLE demo_users (
	id INTEGER NOT NULL AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(255),
    password VARCHAR(255)
);