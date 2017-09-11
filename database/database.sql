CREATE DATABASE "game-api";
REVOKE connect ON DATABASE "game-api" FROM PUBLIC;
CREATE USER "game-api" WITH PASSWORD '<Fetch Password>';
GRANT ALL PRIVILEGES ON DATABASE "game-api" to "game-api";

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
