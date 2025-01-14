// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package wago

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const addMedia = `-- name: AddMedia :exec
INSERT INTO Media (message_id, media_url, media_type, media_size)
VALUES ($1, $2, $3, $4)
`

type AddMediaParams struct {
	MessageID int64
	MediaUrl  string
	MediaType string
	MediaSize int64
}

// Add media details for a message
func (q *Queries) AddMedia(ctx context.Context, arg AddMediaParams) error {
	_, err := q.db.Exec(ctx, addMedia,
		arg.MessageID,
		arg.MediaUrl,
		arg.MediaType,
		arg.MediaSize,
	)
	return err
}

const addMediaMessage = `-- name: AddMediaMessage :one
INSERT INTO Message (conversation_id, sender, recipient, message_type)
VALUES ($1, $2, $3, $4)
RETURNING id,
    conversation_id,
    sender,
    recipient,
    content,
    message_type,
    timestamp
`

type AddMediaMessageParams struct {
	ConversationID int64
	Sender         string
	Recipient      string
	MessageType    string
}

// Add a new media message to a conversation
func (q *Queries) AddMediaMessage(ctx context.Context, arg AddMediaMessageParams) (Message, error) {
	row := q.db.QueryRow(ctx, addMediaMessage,
		arg.ConversationID,
		arg.Sender,
		arg.Recipient,
		arg.MessageType,
	)
	var i Message
	err := row.Scan(
		&i.ID,
		&i.ConversationID,
		&i.Sender,
		&i.Recipient,
		&i.Content,
		&i.MessageType,
		&i.Timestamp,
	)
	return i, err
}

const addTextMessage = `-- name: AddTextMessage :one
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
    timestamp
`

type AddTextMessageParams struct {
	ConversationID int64
	Sender         string
	Recipient      string
	Content        pgtype.Text
	MessageType    string
}

// Add a new text message to a conversation
func (q *Queries) AddTextMessage(ctx context.Context, arg AddTextMessageParams) (Message, error) {
	row := q.db.QueryRow(ctx, addTextMessage,
		arg.ConversationID,
		arg.Sender,
		arg.Recipient,
		arg.Content,
		arg.MessageType,
	)
	var i Message
	err := row.Scan(
		&i.ID,
		&i.ConversationID,
		&i.Sender,
		&i.Recipient,
		&i.Content,
		&i.MessageType,
		&i.Timestamp,
	)
	return i, err
}

const createConversation = `-- name: CreateConversation :one
INSERT INTO Conversation (name, bio)
VALUES ($1, $2)
RETURNING id,
    name,
    bio
`

type CreateConversationParams struct {
	Name string
	Bio  pgtype.Text
}

// Create a new conversation
func (q *Queries) CreateConversation(ctx context.Context, arg CreateConversationParams) (Conversation, error) {
	row := q.db.QueryRow(ctx, createConversation, arg.Name, arg.Bio)
	var i Conversation
	err := row.Scan(&i.ID, &i.Name, &i.Bio)
	return i, err
}

const deleteConversation = `-- name: DeleteConversation :exec
DELETE FROM Conversation
WHERE id = $1
`

// Delete a conversation by ID (and cascade delete messages and media)
func (q *Queries) DeleteConversation(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteConversation, id)
	return err
}

const deleteMessage = `-- name: DeleteMessage :exec
DELETE FROM Message
WHERE id = $1
`

// Delete a message by ID
func (q *Queries) DeleteMessage(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteMessage, id)
	return err
}

const getAllConversations = `-- name: GetAllConversations :many
SELECT id,
    name,
    bio
FROM Conversation
`

// Get all conversations
func (q *Queries) GetAllConversations(ctx context.Context) ([]Conversation, error) {
	rows, err := q.db.Query(ctx, getAllConversations)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Conversation
	for rows.Next() {
		var i Conversation
		if err := rows.Scan(&i.ID, &i.Name, &i.Bio); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getConversationByID = `-- name: GetConversationByID :one
SELECT id,
    name,
    bio
FROM Conversation
WHERE id = $1
`

// Get a specific conversation by ID
func (q *Queries) GetConversationByID(ctx context.Context, id int64) (Conversation, error) {
	row := q.db.QueryRow(ctx, getConversationByID, id)
	var i Conversation
	err := row.Scan(&i.ID, &i.Name, &i.Bio)
	return i, err
}

const getMediaByMessage = `-- name: GetMediaByMessage :many
SELECT id,
    message_id,
    media_url,
    media_type,
    media_size
FROM Media
WHERE message_id = $1
`

// Get all media for a specific message
func (q *Queries) GetMediaByMessage(ctx context.Context, messageID int64) ([]Medium, error) {
	rows, err := q.db.Query(ctx, getMediaByMessage, messageID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Medium
	for rows.Next() {
		var i Medium
		if err := rows.Scan(
			&i.ID,
			&i.MessageID,
			&i.MediaUrl,
			&i.MediaType,
			&i.MediaSize,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getMediaMessagesByConversation = `-- name: GetMediaMessagesByConversation :many
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
ORDER BY m.timestamp
`

type GetMediaMessagesByConversationRow struct {
	MessageID int64
	Sender    string
	Recipient string
	Timestamp pgtype.Timestamp
	MediaUrl  string
	MediaType string
	MediaSize int64
}

// Get all media messages in a specific conversation
func (q *Queries) GetMediaMessagesByConversation(ctx context.Context, conversationID int64) ([]GetMediaMessagesByConversationRow, error) {
	rows, err := q.db.Query(ctx, getMediaMessagesByConversation, conversationID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetMediaMessagesByConversationRow
	for rows.Next() {
		var i GetMediaMessagesByConversationRow
		if err := rows.Scan(
			&i.MessageID,
			&i.Sender,
			&i.Recipient,
			&i.Timestamp,
			&i.MediaUrl,
			&i.MediaType,
			&i.MediaSize,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getMessagesByConversation = `-- name: GetMessagesByConversation :many
SELECT id,
    conversation_id,
    sender,
    recipient,
    content,
    message_type,
    timestamp
FROM Message
WHERE conversation_id = $1
ORDER BY timestamp
`

// Get all messages in a conversation
func (q *Queries) GetMessagesByConversation(ctx context.Context, conversationID int64) ([]Message, error) {
	rows, err := q.db.Query(ctx, getMessagesByConversation, conversationID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Message
	for rows.Next() {
		var i Message
		if err := rows.Scan(
			&i.ID,
			&i.ConversationID,
			&i.Sender,
			&i.Recipient,
			&i.Content,
			&i.MessageType,
			&i.Timestamp,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
