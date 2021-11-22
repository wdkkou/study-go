USE sample_db;

CREATE TABLE users(
	id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	name text NOT NULL,
	email text NOT NULL
);