DROP TABLE IF EXISTS users;

CREATE TABLE users (
  id VARCHAR(32) PRIMARY KEY,
  password varchar(255) NOT NULL,
  email varchar(255) NOT NULL,
  created_at timestamp NOT NULL DEFAULT NOW()
); 

DROP TABLE IF EXISTS posts;

CREATE TABLE posts (
  id VARCHAR(32) PRIMARY KEY,
  title varchar(255) NOT NULL,
  content text NOT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  user_id VARCHAR(32) NOT NULL,
  FOREIGN KEY (user_id) REFERENCES users(id)
);