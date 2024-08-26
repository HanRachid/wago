-- +goose Up
CREATE TABLE Conversation (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    bio TEXT
);
CREATE TABLE Message (
    id BIGSERIAL PRIMARY KEY,
    conversation_id BIGINT NOT NULL REFERENCES Conversation(id),
    sender TEXT NOT NULL,
    recipient TEXT NOT NULL,
    content TEXT,
    message_type TEXT NOT NULL CHECK (
        message_type IN ('text', 'image', 'video', 'audio', 'document')
    ),
    timestamp TIMESTAMP DEFAULT current_timestamp
);
CREATE TABLE Media (
    id BIGSERIAL PRIMARY KEY,
    message_id BIGINT NOT NULL REFERENCES Message(id),
    media_url TEXT NOT NULL,
    media_type TEXT NOT NULL CHECK (
        media_type IN ('image', 'video', 'audio', 'document')
    ),
    media_size BIGINT NOT NULL
);
CREATE INDEX idx_conversation_id ON Message (conversation_id);
CREATE INDEX idx_message_id ON Media (message_id);

-- +goose Down
DROP TABLE IF EXISTS Conversation;
DROP TABLE IF EXISTS Message;
DROP TABLE IF EXISTS Media;
