CREATE TABLE IF NOT EXISTS authentications(
    authentication_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    authentication_user_id UUID NOT NULL,
    authentication_provider VARCHAR (50) NOT NULL,
    CONSTRAINT fk_authentication_user FOREIGN KEY(authentication_user_id) REFERENCES users(user_id),
    UNIQUE (authentication_user_id, authentication_provider)
);