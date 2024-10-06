CREATE TABLE IF NOT EXISTS users (
    username TEXT PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS messages (
    id UUID PRIMARY KEY,
    username TEXT NOT NULL,
    content TEXT NOT NULL,
    is_liked BOOLEAN NOT NULL,

    FOREIGN KEY (username) REFERENCES users(username) ON DELETE CASCADE
);