DROP TABLE IF EXISTS devil_fruits;

CREATE TABLE devil_fruits (
  id VARCHAR(36) NOT NULL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  type ENUM('paramecia', 'zoan', 'logia', 'ancient zoan', 'mythical zoan') NOT NULL,
  awakened BOOLEAN NOT NULL DEFAULT FALSE,
  user_id VARCHAR(36) NULL
);
