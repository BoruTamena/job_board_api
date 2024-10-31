CREATE TABLE job_post_skill (
  id SERIAL PRIMARY KEY,
  job_post_id INTEGER REFERENCES jobs_post(id) ON DELETE CASCADE,
  skill_level VARCHAR(50),
  description TEXT
);