-- migrate:up
CREATE TABLE IF NOT EXISTS refresh_token (
  id INT auto_increment PRIMARY KEY,
  user_id BIGINT NOT NULL,
  refresh_token TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,  
  constraint fk_user_id_refresh_token Foreign Key (user_id) REFERENCES users(id)  
)

-- migrate:down
DROP TABLE IF EXISTS refresh_token
