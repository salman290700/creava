-- migrate:up
CREATE TABLE IF NOT EXISTS images(
  id BIGINT PRIMARY KEY AUTO_INCREMENT,
  image longtext not NULL,  
  created_at DATE
)

-- migrate:down
DROP TABLE IF EXISTS images

