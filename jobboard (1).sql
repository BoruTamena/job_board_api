-- Create tables with foreign keys directly
CREATE TABLE users_type (
  id SERIAL PRIMARY KEY,
  role_name VARCHAR(50)
);

CREATE TABLE job_type (
  id SERIAL PRIMARY KEY,
  category VARCHAR(100)
);

CREATE TABLE users (
  user_id SERIAL PRIMARY KEY,
  email TEXT UNIQUE,
  password VARCHAR(255),
  role_id INTEGER REFERENCES users_type(id) ON DELETE SET NULL,
  created_at TIMESTAMP,
  updated_at TIMESTAMP
);

CREATE TABLE jobs_post (
  id SERIAL PRIMARY KEY,
  title VARCHAR(255),
  job_type_id INTEGER REFERENCES job_type(id) ON DELETE SET NULL,
  posted_by_id INTEGER REFERENCES users(user_id) ON DELETE SET NULL,
  description TEXT,
  skill_set_id INTEGER,
  is_active BOOLEAN,
  created_at TIMESTAMP,
  dead_line TIMESTAMP
);

CREATE TABLE job_post_skill (
  id SERIAL PRIMARY KEY,
  job_post_id INTEGER REFERENCES jobs_post(id) ON DELETE CASCADE,
  skill_level VARCHAR(50),
  description TEXT
);

CREATE TABLE jobs_application (
  id SERIAL PRIMARY KEY,
  user_id INTEGER REFERENCES users(user_id) ON DELETE CASCADE,
  job_post_id INTEGER REFERENCES jobs_post(id) ON DELETE CASCADE,
  apply_date TIMESTAMP
);
