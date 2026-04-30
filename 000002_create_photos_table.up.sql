CREATE TABLE photos (
    id SERIAL PRIMARY KEY,

    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,

    url TEXT NOT NULL
);