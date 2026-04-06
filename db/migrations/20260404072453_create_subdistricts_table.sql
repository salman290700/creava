-- migrate:up
CREATE TABLE IF NOT EXISTS subdistricts (
  id BIGINT PRIMARY KEY auto_increment,
  code VARCHAR(255),
  name VARCHAR(255),
  district_code VARCHAR(255)
)

-- migrate:down
DORP TABLE IF EXISTS subdistricts
