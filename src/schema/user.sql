CREATE TABLE users (
	user_id BIGSERIAL PRIMARY KEY,
	username VARCHAR(32) NOT NULL,
	hashed_password TEXT NOT NULL,
	user_token UUID DEFAULT gen_random_uuid() UNIQUE
);
