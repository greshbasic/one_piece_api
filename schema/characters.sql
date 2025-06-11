DROP TABLE IF EXISTS characters;

CREATE TABLE characters (
  id VARCHAR(36) NOT NULL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  age INT NULL,
  status ENUM('alive', 'dead', 'missing', 'unknown') NULL,
  race VARCHAR(100) NULL,
  origin VARCHAR(100) NULL,
  first_appearance VARCHAR(255) NULL,
  bounty VARCHAR(50) NULL,
  affiliation VARCHAR(255) NULL,
  haki JSON NULL,
  devil_fruit_id VARCHAR(36) NULL,
  CONSTRAINT fk_devil_fruit FOREIGN KEY (devil_fruit_id) REFERENCES devil_fruits(id)
);
