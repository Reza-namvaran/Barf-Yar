CREATE TABLE IF NOT EXISTS activities (
    id SERIAL PRIMARY KEY,
    message_id BIGINT UNIQUE NOT NULL,
    title TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS activity_prompts (
    id SERIAL PRIMARY KEY,
    activity_id INTEGER UNIQUE NOT NULL REFERENCES activities(id) ON DELETE CASCADE,
    prompt_message_id BIGINT UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS activity_supporters (
    id SERIAL PRIMARY KEY,
    activity_id INTEGER NOT NULL REFERENCES activities(id) ON DELETE CASCADE,
    user_id BIGINT NOT NULL,                   
    joined_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(activity_id, user_id)
);

CREATE TABLE IF NOT EXISTS admins (
    id SERIAL PRIMARY KEY,
    username TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS sessions (
    token TEXT PRIMARY KEY,
    expires_at TIMESTAMP NOT NULL
);
