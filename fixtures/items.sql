INSERT INTO materials (name, rarity, created_at, updated_at) VALUES ('Common Dust', 'common', now(), now());
INSERT INTO materials (name, rarity, created_at, updated_at) VALUES ('Rare Dust', 'rare', now(), now());
INSERT INTO materials (name, rarity, created_at, updated_at) VALUES ('Super Dust', 'super', now(), now());
INSERT INTO materials (name, rarity, created_at, updated_at) VALUES ('Ultra Dust', 'ultra', now(), now());
INSERT INTO materials (name, rarity, created_at, updated_at) VALUES ('Ancient Dust', 'ancient', now(), now());


INSERT INTO equipment(id, name, rarity, attributes, slot_type, created_at, updated_at) VALUES(
  'a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 'Manticore', 'rare', '{ "type": "strength", "min": 5, "max": 10 }', 'shoulder', now(), now()
);

INSERT INTO plans(name, rarity, equipment_id, created_at, updated_at) VALUES(
  'Manticore Plan', 'rare', 'a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', now(), now()
);

INSERT INTO users(oauth_identifier, created_at, updated_at) VALUES ('DDAWDW-AWDWAD-AWDWD-AWDWD', now(), now())

