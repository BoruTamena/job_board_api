
CREATE TABLE users (
  user_id SERIAL PRIMARY KEY,
  email TEXT UNIQUE,
  password VARCHAR(255),
  role_id INTEGER REFERENCES users_type(id) ON DELETE SET NULL,
  created_at TIMESTAMP,
  updated_at TIMESTAMP
);
