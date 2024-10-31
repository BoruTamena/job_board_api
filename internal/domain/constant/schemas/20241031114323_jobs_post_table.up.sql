
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