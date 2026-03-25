-- migrate:up
CREATE TABLE IF NOT EXISTS  posts (
  id BIGINT NOT NULL auto_increment PRIMARY KEY,
  user_id BIGINT NOT NULL,
  title VARCHAR(100),
  content LONGTEXT,  
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP NULL,
  constraint fk_user_id_post Foreign Key (user_id) REFERENCES users(id)
)

-- migrate:down
drop table if EXISTS posts;

