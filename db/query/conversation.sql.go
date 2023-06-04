// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: conversation.sql

package query

import (
	"context"
)

const conversationCount = `-- name: ConversationCount :one
select count(*) from conversation
`

func (q *Queries) ConversationCount(ctx context.Context) (int64, error) {
	row := q.queryRow(ctx, q.conversationCountStmt, conversationCount)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createConversation = `-- name: CreateConversation :one
insert into conversation (name) values (null)
returning id, name, protected, selected
`

func (q *Queries) CreateConversation(ctx context.Context) (Conversation, error) {
	row := q.queryRow(ctx, q.createConversationStmt, createConversation)
	var i Conversation
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Protected,
		&i.Selected,
	)
	return i, err
}

const deleteConversation = `-- name: DeleteConversation :one
delete from conversation where id = ? returning id, name, protected, selected
`

func (q *Queries) DeleteConversation(ctx context.Context, id int64) (Conversation, error) {
	row := q.queryRow(ctx, q.deleteConversationStmt, deleteConversation, id)
	var i Conversation
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Protected,
		&i.Selected,
	)
	return i, err
}

const getActiveConversation = `-- name: GetActiveConversation :one
select id, name, protected, selected from conversation where selected=true
`

func (q *Queries) GetActiveConversation(ctx context.Context) (Conversation, error) {
	row := q.queryRow(ctx, q.getActiveConversationStmt, getActiveConversation)
	var i Conversation
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Protected,
		&i.Selected,
	)
	return i, err
}

const getConversations = `-- name: GetConversations :many
SELECT id, name, protected, selected FROM conversation order by id
`

func (q *Queries) GetConversations(ctx context.Context) ([]Conversation, error) {
	rows, err := q.query(ctx, q.getConversationsStmt, getConversations)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Conversation
	for rows.Next() {
		var i Conversation
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Protected,
			&i.Selected,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const nextConversation = `-- name: NextConversation :one
select id, name, protected, selected from conversation
where id > (
	select id from conversation where selected = true
)
order by id
limit 1
`

func (q *Queries) NextConversation(ctx context.Context) (Conversation, error) {
	row := q.queryRow(ctx, q.nextConversationStmt, nextConversation)
	var i Conversation
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Protected,
		&i.Selected,
	)
	return i, err
}

const previousConversation = `-- name: PreviousConversation :one
select id, name, protected, selected from conversation
where id < (
	select id from conversation where selected = true
)
order by id desc
limit 1
`

func (q *Queries) PreviousConversation(ctx context.Context) (Conversation, error) {
	row := q.queryRow(ctx, q.previousConversationStmt, previousConversation)
	var i Conversation
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Protected,
		&i.Selected,
	)
	return i, err
}

const setSelectedConversation = `-- name: SetSelectedConversation :exec
update conversation
set selected = true
where id=?
`

func (q *Queries) SetSelectedConversation(ctx context.Context, id int64) error {
	_, err := q.exec(ctx, q.setSelectedConversationStmt, setSelectedConversation, id)
	return err
}

const unsetSelectedConversation = `-- name: UnsetSelectedConversation :exec
update conversation
set selected = false
`

func (q *Queries) UnsetSelectedConversation(ctx context.Context) error {
	_, err := q.exec(ctx, q.unsetSelectedConversationStmt, unsetSelectedConversation)
	return err
}
