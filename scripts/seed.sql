INSERT INTO users (guid, display_name, first_name, last_name, created_at, updated_at)
SELECT '550e8400-e29b-41d4-a716-446655440000', 'GopherOne', 'Go', 'Pher', NOW(), NOW()
WHERE NOT EXISTS (SELECT 1 FROM users WHERE guid = '550e8400-e29b-41d4-a716-446655440000');

INSERT INTO users (guid, display_name, first_name, last_name, created_at, updated_at)
SELECT '550e8400-e29b-41d4-a716-446655440001', 'GopherTwo', 'Go', 'Lang', NOW(), NOW()
WHERE NOT EXISTS (SELECT 1 FROM users WHERE guid = '550e8400-e29b-41d4-a716-446655440001');
