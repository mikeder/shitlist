CREATE TABLE IF NOT EXISTS clicks(
    id SERIAL PRIMARY KEY,
    user_id UUID NOT NULL,
    clicks NUMERIC,
    CONSTRAINT fk_user
      FOREIGN KEY(user_id) 
        REFERENCES users(user_id)
);
