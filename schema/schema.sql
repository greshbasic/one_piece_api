DROP TABLE IF EXISTS devil_fruits;
DROP TABLE IF EXISTS characters;
DROP TABLE IF EXISTS locations;
DROP TABLE IF EXISTS artifacts;

CREATE TABLE characters (
  id BIGINT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  bounty BIGINT DEFAULT NULL,
  haki JSON DEFAULT NULL,
  affiliation VARCHAR(255) DEFAULT NULL,
  origin VARCHAR(255) DEFAULT NULL,
  status VARCHAR(100) DEFAULT NULL,
  age INT DEFAULT NULL,

  INDEX idx_characters_name (name),
  INDEX idx_characters_affiliation (affiliation),
  INDEX idx_characters_status (status)
);

CREATE TABLE devil_fruits (
  id BIGINT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  type ENUM('paramecia', 'zoan', 'logia', 'ancient zoan', 'mythical zoan') NOT NULL,
  awakened TINYINT(1) NOT NULL DEFAULT 0,
  user_id BIGINT NULL,
  
  CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES characters(id) ON DELETE SET NULL,

  INDEX idx_devil_fruits_user_id (user_id),
  INDEX idx_devil_fruits_type (type),
  INDEX idx_devil_fruits_name (name)
);

CREATE TABLE locations (
  id BIGINT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  description TEXT DEFAULT NULL,
  affiliation VARCHAR(255) DEFAULT NULL,

  INDEX idx_locations_name (name),
  INDEX idx_locations_affiliation (affiliation)
);

CREATE TABLE artifacts (
  id BIGINT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  description TEXT DEFAULT NULL,
  origin VARCHAR(255) DEFAULT NULL,

  INDEX idx_artifacts_name (name),
  INDEX idx_artifacts_origin (origin)
);



