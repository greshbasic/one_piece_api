INSERT INTO characters (id, name, bounty, haki, affiliation, origin, status, age) VALUES
  (1, 'Gol D. Roger', 5564800000, '{"armament":{"awakened":true},"conqueror":{"awakened":true},"observation":{"awakened":true}}', 'Roger Pirates', 'Loguetown', 'Deceased', 53),
  (2, 'Monkey D. Luffy', 3000000000, '{"armament":{"awakened":true},"conqueror":{"awakened":true},"observation":{"awakened":true}}', 'Straw Hat Pirates', 'Fuschia Village', 'Alive', 19),
  (3, 'Charlotte Katakuri', 1057000000, '{"armament":{"awakened":true},"observation":{"awakened":true},"conqueror":{"awakened":false}}', 'Big Mom Pirates', 'Whole Cake Island', 'Alive', 48),
  (4, 'Gaimon', NULL, NULL, NULL, 'Island of Rare Animals, East Blue', 'Alive', NULL),
  (5, 'Kizaru', 3000000000, '{"armament":{"awakened":true},"observation":{"awakened":true}}', 'Marines', 'Unknown', 'Alive', NULL);

INSERT INTO devil_fruits (id, name, type, awakened, user_id) VALUES
  (1, 'Hito Hito no Mi, Model: Nika', 'mythical zoan', TRUE, 2),
  (2, 'Mochi Mochi no Mi', 'paramecia', TRUE, 3),
  (3, 'Mera Mera no Mi', 'logia', FALSE, NULL),
  (4, 'Pika Pika no Mi', 'logia', FALSE, 5);

INSERT INTO locations (name, description, affiliation) VALUES
  ('Wano', 'A closed country known for its samurai and isolationist policies.', 'Straw Hat Pirates'),
  ('God Valley', 'No longer exists; was a mysterious island involved with the Rocks Pirates and the World Government.', NULL),
  ('Mary Geoise', 'The Holy Land and capital of the World Government, located on the Red Line.', 'World Government');

INSERT INTO artifacts (id, name, description, origin) VALUES
  (1, 'The One Piece', 'The legendary treasure said to be of unimaginable value, sought by pirates worldwide.', 'Laugh Tale'),
  (2, 'Poseidon', 'The ancient weapon and mermaid princess with the power to communicate with Sea Kings.', 'Fish-Man Island'),
  (3, 'Treasure Tree Adam', 'A rare and extremely durable tree whose wood is highly prized for shipbuilding.', 'Elbaph');
