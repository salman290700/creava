-- migrate:up
CREATE TABLE IF NOT EXISTS comment_likes (
  id BIGINT NOT NULL auto_increment PRIMARY KEY,
  user_id BIGINT NOT NULL,
  comment_id BIGINT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP NULL,  
  constraint fk_user_id_comment_likes Foreign Key (user_id) REFERENCES users(id),
  constraint fk_comment_id_comment_likes Foreign Key (comment_id) REFERENCES comments(id)
)

-- migrate:down
DROP TABLE IF EXISTS comment_likes
