CREATE TABLE students (
  student_id SERIAL PRIMARY KEY,
  first_name VARCHAR(100) NOT NULL,
  last_name VARCHAR(100) NOT NULL,
  email VARCHAR(100) NOT NULL UNIQUE,
  enrollment_date DATE 
);

CREATE INDEX idx_last_name ON students (last_name);
