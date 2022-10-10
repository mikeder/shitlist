CREATE TABLE IF NOT EXISTS authentications(
    authentication_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    authentication_provider VARCHAR (50) NOT NULL,
    CONSTRAINT fk_authentication_user FOREIGN KEY(user_id) REFERENCES users(user_id),
    UNIQUE (authentication_provider, fk_authentication_user)
);