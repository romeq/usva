// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: files.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const deleteFile = `-- name: DeleteFile :exec
DELETE FROM file
WHERE file_uuid = $1
`

func (q *Queries) DeleteFile(ctx context.Context, fileUuid string) error {
	_, err := q.db.Exec(ctx, deleteFile, fileUuid)
	return err
}

const fileInformation = `-- name: FileInformation :one
SELECT file_uuid,
    title,
    upload_date,
    isencrypted,
    viewcount
FROM file
WHERE file_uuid = $1
`

type FileInformationRow struct {
	FileUuid    string
	Title       sql.NullString
	UploadDate  string
	Isencrypted bool
	Viewcount   int32
}

func (q *Queries) FileInformation(ctx context.Context, fileUuid string) (FileInformationRow, error) {
	row := q.db.QueryRow(ctx, fileInformation, fileUuid)
	var i FileInformationRow
	err := row.Scan(
		&i.FileUuid,
		&i.Title,
		&i.UploadDate,
		&i.Isencrypted,
		&i.Viewcount,
	)
	return i, err
}

const getLastSeenAll = `-- name: GetLastSeenAll :many
SELECT file_uuid,
    last_seen
FROM file
`

type GetLastSeenAllRow struct {
	FileUuid string
	LastSeen time.Time
}

func (q *Queries) GetLastSeenAll(ctx context.Context) ([]GetLastSeenAllRow, error) {
	rows, err := q.db.Query(ctx, getLastSeenAll)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetLastSeenAllRow
	for rows.Next() {
		var i GetLastSeenAllRow
		if err := rows.Scan(&i.FileUuid, &i.LastSeen); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPasswordHash = `-- name: GetPasswordHash :one
SELECT passwdhash
FROM file
WHERE file_uuid = $1
`

func (q *Queries) GetPasswordHash(ctx context.Context, fileUuid string) (sql.NullString, error) {
	row := q.db.QueryRow(ctx, getPasswordHash, fileUuid)
	var passwdhash sql.NullString
	err := row.Scan(&passwdhash)
	return passwdhash, err
}

const newFile = `-- name: NewFile :exec
INSERT INTO file(
        file_uuid,
        title,
        uploader,
        passwdhash,
        upload_date,
        isencrypted,
        last_seen,
        viewcount
    )
VALUES($1, $2, $3, $4, $5, $6, $7, 0)
`

type NewFileParams struct {
	FileUuid    string
	Title       sql.NullString
	Uploader    sql.NullString
	Passwdhash  sql.NullString
	UploadDate  string
	Isencrypted bool
	LastSeen    time.Time
}

func (q *Queries) NewFile(ctx context.Context, arg NewFileParams) error {
	_, err := q.db.Exec(ctx, newFile,
		arg.FileUuid,
		arg.Title,
		arg.Uploader,
		arg.Passwdhash,
		arg.UploadDate,
		arg.Isencrypted,
		arg.LastSeen,
	)
	return err
}

const updateLastSeen = `-- name: UpdateLastSeen :exec
UPDATE file
SET last_seen = $2
WHERE file_uuid = $1
`

type UpdateLastSeenParams struct {
	FileUuid string
	LastSeen time.Time
}

func (q *Queries) UpdateLastSeen(ctx context.Context, arg UpdateLastSeenParams) error {
	_, err := q.db.Exec(ctx, updateLastSeen, arg.FileUuid, arg.LastSeen)
	return err
}

const updateViewCount = `-- name: UpdateViewCount :exec
UPDATE file
SET viewcount = viewcount + 1
WHERE file_uuid = $1
`

func (q *Queries) UpdateViewCount(ctx context.Context, fileUuid string) error {
	_, err := q.db.Exec(ctx, updateViewCount, fileUuid)
	return err
}