DROP SCHEMA IF EXISTS myservice CASCADE;

-- CREATE SCHEMA IF NOT EXISTS myservice;

CREATE TABLE IF NOT EXISTS myservice.data (
    id SERIAL PRIMARY KEY,
    server_ip CIDR,
    client_ip CIDR,
    metadata JSONB,
    msg TEXT NOT NULL
);

INSERT INTO myservice.data (server_ip, client_ip, metadata, msg)
VALUES
    ('64.248.120.109', '187.210.255.49', '{"hello": "world"}', 'Est quibusdam qui numquam voluptas provident sit.'),
    ('16.98.253.18', '153.125.231.60', '{"hey": "there"}', 'voluptas laboriosam aut');
