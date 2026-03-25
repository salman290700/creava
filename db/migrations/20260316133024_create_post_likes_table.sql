-- migrate:up
CREATE TABLE IF NOT EXISTS post_likes (
  id BIGINT NOT NULL auto_increment PRIMARY KEY,
  user_id BIGINT NOT NULL,
  post_id BIGINT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP NULL,  
  constraint fk_user_id_post_likes Foreign Key (user_id) REFERENCES users(id),
  constraint fk_post_id_post_likes Foreign Key (post_id) REFERENCES posts(id)
)

-- migrate:down
DROP TABLE IF EXISTS comments
