-- Create a new conversation
-- name: CreateConversation :one
INSERT INTO Conversation (name, bio)
VALUES ($1, $2)
RETURNING id,
    name,
    bio;
-- Add a new text message to a conversation
-- name: AddTextMessage :one
INSERT INTO Message (
        conversation_id,
        sender,
        recipient,
        content,
        message_type
    )
VALUES ($1, $2, $3, $4, $5)
RETURNING id,
    conversation_id,
    sender,
    recipient,
    content,
    message_type,
    timestamp;
-- Add a new media message to a conversation
-- name: AddMediaMessage :one
INSERT INTO Message (conversation_id, sender, recipient, message_type)
VALUES ($1, $2, $3, $4)
RETURNING id,
    conversation_id,
    sender,
    recipient,
    content,
    message_type,
    timestamp;
-- Add media details for a message
-- name: AddMedia :exec
INSERT INTO Media (message_id, media_url, media_type, media_size)
VALUES ($1, $2, $3, $4);
-- Get all messages in a conversation
-- name: GetMessagesByConversation :many
SELECT id,
    conversation_id,
    sender,
    recipient,
    content,
    message_type,
    timestamp
FROM Message
WHERE conversation_id = $1
ORDER BY timestamp;
-- Get all media for a specific message
-- name: GetMediaByMessage :many
SELECT id,
    message_id,
    media_url,
    media_type,
    media_size
FROM Media
WHERE message_id = $1;
-- Get all conversations
-- name: GetAllConversations :many
SELECT id,
    name,
    bio
FROM Conversation;
-- Get a specific conversation by ID
-- name: GetConversationByID :one
SELECT id,
    name,
    bio
FROM Conversation
WHERE id = $1;
-- Delete a message by ID
-- name: DeleteMessage :exec
DELETE FROM Message
WHERE id = $1;
-- Delete a conversation by ID (and cascade delete messages and media)
-- name: DeleteConversation :exec
DELETE FROM Conversation
WHERE id = $1;
-- Get all media messages in a specific conversation
-- name: GetMediaMessagesByConversation :many
SELECT m.id AS message_id,
    m.sender,
    m.recipient,
    m.timestamp,
    media.media_url,
    media.media_type,
    media.media_size
FROM Message m
    JOIN Media media ON m.id = media.message_id
WHERE m.conversation_id = $1
ORDER BY m.timestamp;