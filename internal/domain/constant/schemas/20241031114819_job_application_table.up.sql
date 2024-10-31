CREATE TABLE jobs_application (
  id SERIAL PRIMARY KEY,
  user_id INTEGER REFERENCES users(user_id) ON DELETE CASCADE,
  job_post_id INTEGER REFERENCES jobs_post(id) ON DELETE CASCADE,
  apply_date TIMESTAMP
);