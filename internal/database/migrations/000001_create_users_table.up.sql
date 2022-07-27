CREATE EXTENSION "pgcrypto";

CREATE TABLE IF NOT EXISTS users(
   -- https://www.starkandwayne.com/blog/uuid-primary-keys-in-postgresql/
   user_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
   user_name VARCHAR (50) UNIQUE NOT NULL,
   user_email VARCHAR (300) UNIQUE NOT NULL
);
